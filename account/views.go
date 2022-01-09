package account

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/errors"
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

func register(ctx *gin.Context, store *store) {
	var register EmailPassword
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		bindingError := errors.BindingError(err)
		ctx.AbortWithStatusJSON(bindingError.Code, bindingError)
		return
	}

	if err := StrongPassword(register.Password); err != nil {
		ctx.AbortWithStatusJSON(WeakPasswordError.Code, WeakPasswordError)
		return
	}

	exists := store.emailExists(ctx, register.Email)
	if exists {
		ctx.AbortWithStatusJSON(DuplicateEmailError.Code, DuplicateEmailError)
		return
	}

	encryptedPassword, err := Encrypt(register.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(PasswordHashError.Code, PasswordHashError)
	}

	user, err := store.createAccount(ctx, register.Email, encryptedPassword)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return
	}
	var responseData = &PublicResponse{}
	responseData.Email = user.Email
	responseData.ID.ID = user.ID
	ctx.JSON(http.StatusCreated, responseData)

}

func login(ctx *gin.Context, store *store) {
	var login EmailPassword
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		bindingError := errors.BindingError(err)
		ctx.AbortWithStatusJSON(bindingError.Code, bindingError)
		return
	}

	account, err := store.findAccount(ctx, login.Email)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return
	}

	if !Validate(login.Password, account.Password) {
		ctx.AbortWithStatusJSON(RecordNotMatchError.Code, RecordNotMatchError)
		return
	}

	tokenBody := RandomString(128)
	t, err := store.createToken(ctx, account.ID, tokenBody, StringRefreshToken, store.Config().RefreshTokenExpireTime)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return
	}
	ctx.SetCookie(StringRefreshToken, t.Body, int(t.ExpireAt.Sub(time.Now()).Seconds()), "/account/refresh", store.Config().Domain, false, true)
	ctx.AbortWithStatus(http.StatusNoContent)
}

func refreshToken(ctx *gin.Context, s *store) {
	rt, err := cookieTokenValidate(ctx, s, StringRefreshToken)
	if err != nil {
		return
	}

	tokenBody := RandomString(128)
	t, err := s.createToken(ctx, rt.Edges.Account.ID, tokenBody, StringAccessToken, s.Config().AccessTokenExpireTime)
	ctx.SetCookie(StringAccessToken, t.Body, int(t.ExpireAt.Sub(time.Now()).Seconds()), "/", s.Config().Domain, false, true)
	ctx.AbortWithStatus(http.StatusCreated)
	return

}

func info(ctx *gin.Context, s *store) {
	account := ctx.MustGet(StringAccount).(*ent.Account)
	ctx.JSON(http.StatusOK, account)
}
