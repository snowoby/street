package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/ent"
	"street/ent/token"
	"street/errs"
	"street/pkg/d"
	"street/pkg/data/value"
	"street/pkg/utils"
	"time"
)

type Service interface {
	MustLogin(c *gin.Context)
}

type service struct {
	db *ent.Client
}

func New() *service {
	return &service{}
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
// @Success 201 {object} ResponseToken
// @Failure 400 {object} errs.HTTPError
// @Router /account/refresh [post]
func (s *service) Refresh(ctx *gin.Context) {
	tokenType := value.StringRefreshToken
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

func (s *service) MustLogin(ctx *gin.Context) {
	t := s.tryToken(ctx, value.StringAccessToken)
	if t == nil {
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code, &errs.UnauthorizedError)
		return
	}
	ctx.Set(value.StringAccount, t.Edges.Account)
	ctx.Set(value.StringAccessToken, t)
	ctx.Next()
	return
}
