package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
	"street/pkg/utils"
)

type Token struct {
	Token string `header:"Authorization" binding:"required"`
}

// MustRefresh godoc
// @Summary refresh token
// @Tags account,token
// @Produce json
// @Success 201 {object} ent.Token
// @Failure 400 {object} errs.HTTPError
// @Router /account/refresh [post]
func MustRefresh(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	tokenType := value.StringRefreshToken
	t := tryToken(ctx, store, tokenType)
	if t == nil {
		return 0, nil, errs.UnauthorizedError
	}

	tokenBody := utils.RandomString(128)
	// Create access token
	t, err := store.DB.Token.Create(ctx, t.Edges.Account.ID, tokenBody, value.StringAccessToken, store.SiteConfig.RefreshTokenExpireTime)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, t, nil
}

func TryAccessToken(ctx *gin.Context, store *data.Store) {
	t := tryToken(ctx, store, value.StringAccessToken)
	if t != nil {
		ctx.Set(value.StringAccount, t.Edges.Account)
		ctx.Set(value.StringAccessToken, t)
	}

	ctx.Next()
}

func tryToken(ctx *gin.Context, store *data.Store, tokenType string) *ent.Token {
	var token Token
	err := ctx.ShouldBindHeader(&token)
	if err == nil {
		if len(token.Token) > 7 {
			tokenBody := token.Token[7:]
			t, err := store.DB.Token.Find(ctx, tokenBody, tokenType, true)
			if err == nil {
				if TokenIsValid(t) {
					return t
				}
			}
		}
	}
	return nil
}
func MustLogin(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringAccount)
	if !ok {
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code, &errs.UnauthorizedError)
		return
	}
	ctx.Next()
	return
}
