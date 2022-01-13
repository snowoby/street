package account

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
	"street/pkg/utils"
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

func register(ctx *gin.Context, store *data.Store) (code int, responseData interface{}, err errs.ResponseError) {
	var register EmailPassword
	e := ctx.ShouldBindJSON(&register)
	if err != nil {
		return code, responseData, errs.BindingError(e)
	}

	if err := utils.StrongPassword(register.Password); err != nil {
		return code, responseData, errs.WeakPasswordError
	}

	exists, e := store.EmailExists(ctx, register.Email)
	if err != nil {
		databaseError := errs.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}

	if exists {
		ctx.JSON(errs.DuplicateEmailError.Code, errs.DuplicateEmailError)
		return
	}

	encryptedPassword, err := utils2.Encrypt(register.Password)
	if err != nil {
		ctx.JSON(errs.PasswordHashError.Code, errs.PasswordHashError)
	}

	user, err := store.CreateAccount(ctx, register.Email, encryptedPassword)
	if err != nil {
		databaseError := errs.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return
	}

	var responseData = &PublicResponse{}
	responseData.Email = user.Email
	responseData.ID.ID = user.ID
	ctx.JSON(http.StatusCreated, responseData)

}

func login(ctx *gin.Context, store *data.Store) (code int, responseData interface{}, err errs.ResponseError) {
	var login EmailPassword
	e := ctx.ShouldBindJSON(&login)
	if err != nil {
		return code, responseData, errs.BindingError(e)
	}

	account, e := store.FindAccount(ctx, login.Email)
	if err != nil {
		return code, responseData, errs.DatabaseError(e)
	}

	if !utils.Validate(login.Password, account.Password) {
		return code, responseData, errs.RecordNotMatchError
	}

	tokenBody := utils.RandomString(128)
	t, e := store.CreateToken(ctx, account.ID, tokenBody, value.StringRefreshToken, store.Config().RefreshTokenExpireTime)
	if err != nil {
		return code, responseData, errs.DatabaseError(e)
	}

	return http.StatusCreated, t, nil
	//ctx.SetCookie(value.StringRefreshToken, t.Body, int(t.ExpireTime.Sub(time.Now()).Seconds()), "/account/refresh", store.Config().Domain, false, true)
}

func info(ctx *gin.Context, s *data.Store, operator controller.Identity) {
	return
}
