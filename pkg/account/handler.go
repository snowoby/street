package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/ent"
	"street/ent/account"
	"street/ent/token"
	"street/errs"
	"street/pkg/auth"
	"street/pkg/composer"
	"street/pkg/d"
	"street/pkg/operator"
	"street/pkg/utils"
	"time"
)

type service struct {
	db     *ent.Client
	auth   auth.Service
	router *gin.RouterGroup
}

func (s *service) registerRouters() {
	s.router.POST("/refresh", s.refresh)
	s.router.POST("/login", composer.Bare(s.login))
	s.router.POST("/register", composer.Bare(s.register))
	s.router.GET("/", s.auth.MustLogin, composer.Authed(s.info))
}

func New(db *ent.Client, auth auth.Service, router *gin.RouterGroup) *service {
	s := &service{
		db:     db,
		auth:   auth,
		router: router,
	}
	s.registerRouters()
	return s
}

// Register godoc
// @Summary register an account
// @Tags account
// @Accept json
// @Produce json
// @Param accountInfo body d.AccountForm true "account info"
// @Success 201 {object} d.Account
// @Failure 400 {object} errs.HTTPError
// @Router /account/register [post]
func (s *service) register(ctx *gin.Context) (int, interface{}, error) {

	var accountForm d.AccountForm
	err := ctx.ShouldBindJSON(&accountForm)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	if err := utils.StrongPassword(accountForm.Password); err != nil {
		return 0, nil, errs.WeakPasswordError
	}

	exists, err := s.db.Account.Query().Where(account.EmailEQ(accountForm.Email)).Exist(ctx)
	if err != nil {
		return 0, nil, err
	}
	if exists {
		return 0, nil, errs.DuplicateEmailError
	}

	encryptedPassword, err := utils.Encrypt(accountForm.Password)
	if err != nil {
		return 0, nil, errs.PasswordHashError
	}

	user, err := s.db.Account.Create().SetEmail(accountForm.Email).SetPassword(encryptedPassword).Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, d.AccountFromEnt(user), nil
}

// Login godoc
// @Summary login an account
// @Tags account
// @Accept json
// @Produce json
// @Param accountInfo body d.AccountForm true "account info"
// @Success 201 {object} d.Token
// @Failure 400 {object} errs.HTTPError
// @Router /account/login [post]
func (s *service) login(ctx *gin.Context) (int, interface{}, error) {
	var accountForm d.AccountForm
	err := ctx.ShouldBindJSON(&accountForm)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	accountRecord, err := s.db.Account.Query().Where(account.EmailEQ(accountForm.Email)).Only(ctx)
	if err != nil {
		return 0, nil, err
	}

	if !utils.Validate(accountForm.Password, accountRecord.Password) {
		return 0, nil, errs.RecordNotMatchError
	}

	tokenBody := utils.RandomString(128)
	t, err := s.db.Token.Create().
		SetAccountID(accountRecord.ID).
		SetBody(tokenBody).
		SetType(d.StringRefreshToken).
		SetExpire(time.Now().Add(time.Hour * 24 * 7 * 4)).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.TokenFromEnt(t), nil
}

// Info godoc
// @Summary info an account
// @Tags account,profile
// @Produce json
// @Success 201 {object} d.Identity
// @Failure 400 {object} errs.HTTPError
// @Router /account [get]
func (s *service) info(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {

	profiles := d.ProfilesFromEnt(operator.AllProfiles(ctx))

	return http.StatusOK, &d.Identity{
		Account:  d.AccountFromEnt(operator.Account()),
		Profiles: profiles,
	}, nil
}

func (s *service) tryToken(ctx *gin.Context, tokenType string) *ent.Token {
	var tokenData d.TokenForm
	err := ctx.ShouldBindHeader(&tokenData)
	if err == nil {
		if len(tokenData.Token) > 7 {
			tokenBody := tokenData.Token[7:]
			tokenRecord, err := s.db.Token.Query().Where(token.Body(tokenBody)).Where(token.Type(tokenType)).WithAccount().Only(ctx)
			if err == nil {
				if tokenIsValid(tokenRecord) {
					return tokenRecord
				}
			}
		}
	}
	return nil
}

func tokenIsValid(token *ent.Token) bool {
	if token == nil {
		return false
	}
	if token.Expire.Before(time.Now()) {
		return false
	}
	return true
}

// Refresh godoc
// @Summary refresh token
// @Tags account,token
// @Produce json
// @Success 201 {object} d.Token
// @Failure 400 {object} errs.HTTPError
// @Router /account/refresh [post]
func (s *service) refresh(ctx *gin.Context) {
	tokenType := d.StringRefreshToken
	t := s.tryToken(ctx, tokenType)
	if t == nil {
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code, errs.UnauthorizedError)
		return
	}

	tokenBody := utils.RandomString(128)
	// Create access token
	t, err := s.db.Token.Create().
		SetAccountID(t.Edges.Account.ID).
		SetBody(tokenBody).
		SetType(tokenType).
		SetExpire(time.Now().Add(time.Hour * 24 * 7)).
		Save(ctx)
	if t == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, d.TokenFromEnt(t))
	return
}
