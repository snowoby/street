package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/ent"
	"street/ent/account"
	"street/errs"
	"street/pkg/d"
	"street/pkg/data/value"
	"street/pkg/operator"
	"street/pkg/utils"
	"time"
)

type Service interface {
}

type service struct {
	db *ent.Client
}

func New() *service {
	return &service{}
}

func (s *service) Create(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {

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

func (s *service) Login(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
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
		SetType(value.StringRefreshToken).
		SetExpire(time.Now().Add(time.Hour * 24 * 7 * 4)).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.TokenFromEnt(t), nil
}
