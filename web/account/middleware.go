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

func MustRefresh(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	tokenType := value.StringRefreshToken
	t := tryToken(ctx, store, tokenType)
	if t == nil {
		return 0, nil, errs.UnauthorizedError
	}

	tokenBody := utils.RandomString(128)
	// Create access token
	t, err := store.CreateToken(ctx, t.Edges.Account.ID, tokenBody, value.StringAccessToken, store.Config().RefreshTokenExpireTime)
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
			t, err := store.FindToken(ctx, tokenBody, tokenType, true)
			if err == nil {
				if utils.TokenIsValid(t) {
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
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code(), errs.UnauthorizedError)
		return
	}
	ctx.Next()
	return

}
