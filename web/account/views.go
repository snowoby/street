package account

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
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

func register(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	var register EmailPassword
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	if err := utils.StrongPassword(register.Password); err != nil {
		return 0, nil, errs.WeakPasswordError
	}

	exists, err := store.EmailExists(ctx, register.Email)
	if err != nil {
		return 0, nil, err
	}

	if exists {
		return 0, nil, errs.DuplicateEmailError
	}

	encryptedPassword, err := utils.Encrypt(register.Password)
	if err != nil {
		return 0, nil, errs.PasswordHashError
	}

	user, err := store.CreateAccount(ctx, register.Email, encryptedPassword)
	if err != nil {
		return 0, nil, err
	}

	var responseData = &PublicResponse{}
	responseData.Email = user.Email
	responseData.ID.ID = user.ID

	return http.StatusCreated, responseData, nil
}

func login(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	var login EmailPassword
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	account, err := store.FindAccount(ctx, login.Email)
	if err != nil {
		return 0, nil, err
	}

	if !utils.Validate(login.Password, account.Password) {
		return 0, nil, errs.RecordNotMatchError
	}

	tokenBody := utils.RandomString(128)
	t, err := store.CreateToken(ctx, account.ID, tokenBody, value.StringRefreshToken, store.Config().RefreshTokenExpireTime)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, t, nil
	//ctx.SetCookie(value.StringRefreshToken, t.Body, int(t.ExpireTime.Sub(time.Now()).Seconds()), "/account/refresh", store.Config().Domain, false, true)
}

func info(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	return http.StatusOK, struct {
		Account  *ent.Account   `json:"account"`
		Profiles []*ent.Profile `json:"profiles"`
	}{
		Account:  identity.Account(),
		Profiles: identity.AllProfiles(),
	}, nil
}
