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

// register godoc
// @Summary register an account
// @Tags account
// @Accept json
// @Produce json
// @Param accountInfo body EmailPassword true "account info"
// @Success 201 {object} PublicResponse
// @Failure 400 {object} errs.HTTPError
// @Router /account/register [post]
func register(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	var register EmailPassword
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	if err := utils.StrongPassword(register.Password); err != nil {
		return 0, nil, errs.WeakPasswordError
	}

	exists, err := store.DB.Account.EmailExists(ctx, register.Email)
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

	user, err := store.DB.Account.Create(ctx, register.Email, encryptedPassword)
	if err != nil {
		return 0, nil, err
	}

	var responseData = &PublicResponse{}
	responseData.Email = user.Email
	responseData.ID.ID = user.ID

	return http.StatusCreated, responseData, nil
}

// login godoc
// @Summary login an account
// @Tags account
// @Accept json
// @Produce json
// @Param accountInfo body EmailPassword true "account info"
// @Success 201 {object} ent.Token
// @Failure 400 {object} errs.HTTPError
// @Router /account/login [post]
func login(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	var login EmailPassword
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	account, err := store.DB.Account.Find(ctx, login.Email)
	if err != nil {
		return 0, nil, err
	}

	if !utils.Validate(login.Password, account.Password) {
		return 0, nil, errs.RecordNotMatchError
	}

	tokenBody := utils.RandomString(128)
	t, err := store.DB.Token.Create(ctx, account.ID, tokenBody, value.StringRefreshToken, store.SiteConfig.RefreshTokenExpireTime)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, t, nil
	//ctx.SetCookie(value.StringRefreshToken, t.Body, int(t.ExpireTime.Sub(time.Now()).Seconds()), "/account/refresh", store.Config().Domain, false, true)
}

type Identity struct {
	Account  *ent.Account   `json:"account"`
	Profiles []*ent.Profile `json:"profiles"`
}

// info godoc
// @Summary info an account
// @Tags account,profile
// @Produce json
// @Success 201 {object} Identity
// @Failure 400 {object} errs.HTTPError
// @Router /account [get]
func info(_ *gin.Context, _ *data.Store, identity *controller.Identity) (int, interface{}, error) {

	return http.StatusOK, &Identity{
		Account:  identity.Account(),
		Profiles: identity.AllProfiles(),
	}, nil
}

func TokenIsValid(token *ent.Token) bool {
	if token == nil {
		return false
	}
	if token.Expire.Before(time.Now()) {
		return false
	}
	return true
}
