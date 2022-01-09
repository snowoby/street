package account

import (
	"github.com/gin-gonic/gin"
	"street/db"
	"street/ent"
	"street/errors"
	"time"
)

const StringAccount = "account"

func AccessTokenMiddleware(dbStore db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := &store{dbStore}
		at, err := cookieTokenValidate(ctx, s, StringAccessToken)
		if err != nil {
			return
		}

		ctx.Set(StringAccount, at.Edges.Account)
		ctx.Set(StringAccessToken, at)
		ctx.Next()
		return
	}

}

func cookieTokenValidate(ctx *gin.Context, store *store, tokenType string) (*ent.Token, error) {
	tokenBody, err := ctx.Cookie(tokenType)
	if err != nil {
		ctx.AbortWithStatusJSON(TokenNotExistsError.Code, TokenNotExistsError)
		return nil, err
	}

	t, err := store.findToken(ctx, tokenBody, tokenType, true)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return nil, err
	}

	if t.ExpireAt.Before(time.Now()) {
		ctx.AbortWithStatusJSON(TokenExpiredError.Code, TokenExpiredError)
		return nil, err
	}

	return t, nil
}
