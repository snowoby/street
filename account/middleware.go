package account

import (
	"github.com/gin-gonic/gin"
	"street/db"
	"street/ent"
	"street/errors"
	"time"
)

const StringAccount = "account"

func AccessTokenMiddleware(ctx *gin.Context, s *db.Store) {
	at, err := cookieTokenValidate(ctx, s, StringAccessToken)
	if err != nil {
		return
	}

	ctx.Set(StringAccount, at.Edges.Account)
	ctx.Set(StringAccessToken, at)
	ctx.Next()
	return

}

func cookieTokenValidate(ctx *gin.Context, store *db.Store, tokenType string) (*ent.Token, error) {
	tokenBody, err := ctx.Cookie(tokenType)
	if err != nil {
		ctx.AbortWithStatusJSON(TokenNotExistsError.Code, TokenNotExistsError)
		return nil, err
	}

	t, err := store.DB().FindToken(ctx, tokenBody, tokenType, true)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return nil, err
	}

	if t.ExpireTime.Before(time.Now()) {
		ctx.AbortWithStatusJSON(TokenExpiredError.Code, TokenExpiredError)
		return nil, err
	}

	return t, nil
}
