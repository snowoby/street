package account

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errors"
	"street/utils"
	"time"
)

type ID struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

type EmailPassword struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type PublicResponse struct {
	EmailPassword
	ID
	Password *struct{} `json:"password,omitempty"`
}

func register(ctx *gin.Context, store *db.Store) {
	var register EmailPassword
	if !utils.MustBindJSON(ctx, &register) {
		return
	}

	if err := utils.StrongPassword(register.Password); err != nil {
		ctx.JSON(errors.WeakPasswordError.Code, errors.WeakPasswordError)
		return
	}

	exists, err := store.EmailExists(ctx, register.Email)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}

	if exists {
		ctx.JSON(errors.DuplicateEmailError.Code, errors.DuplicateEmailError)
		return
	}

	encryptedPassword, err := utils.Encrypt(register.Password)
	if err != nil {
		ctx.JSON(errors.PasswordHashError.Code, errors.PasswordHashError)
	}

	user, err := store.CreateAccount(ctx, register.Email, encryptedPassword)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}

	var responseData = &PublicResponse{}
	responseData.Email = user.Email
	responseData.ID.ID = user.ID
	ctx.JSON(http.StatusCreated, responseData)

}

func login(ctx *gin.Context, store *db.Store) {
	var login EmailPassword
	if !utils.MustBindJSON(ctx, &login) {
		return
	}

	account, err := store.FindAccount(ctx, login.Email)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}

	if !utils.Validate(login.Password, account.Password) {
		ctx.JSON(errors.RecordNotMatchError.Code, errors.RecordNotMatchError)
		return
	}

	tokenBody := utils.RandomString(128)
	t, err := store.CreateToken(ctx, account.ID, tokenBody, value.StringRefreshToken, store.Config().RefreshTokenExpireTime)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}
	ctx.SetCookie(value.StringRefreshToken, t.Body, int(t.ExpireTime.Sub(time.Now()).Seconds()), "/account/refresh", store.Config().Domain, false, true)
	ctx.Status(http.StatusNoContent)
}

func info(ctx *gin.Context, s *db.Store) {
	account := ctx.MustGet(value.StringAccount).(*ent.Account)
	ctx.JSON(http.StatusOK, account)
}
