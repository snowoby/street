package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/ent"
	"street/errors"
	"time"
)

func AccessTokenMiddleware(ctx *gin.Context, s *store) {
	at, err := cookieTokenValidate(ctx, s, AccessToken)
	if err != nil {
		return
	}

	ctx.Set("account", at.Edges.Account)
	ctx.Set("accessToken", at)
	ctx.Next()
	return

}

func RefreshTokenMiddleware(ctx *gin.Context, s *store) {
	rt, err := cookieTokenValidate(ctx, s, RefreshToken)
	if err != nil {
		return
	}

	tokenBody := RandomString(128)
	t, err := s.createToken(ctx, rt.Edges.Account.ID, tokenBody, AccessToken, s.Config().AccessTokenExpireTime)
	ctx.SetCookie(AccessToken, t.Body, int(t.ExpireAt.Sub(time.Now()).Seconds()), "/", s.Config().Domain, false, true)
	ctx.AbortWithStatus(http.StatusCreated)
	return

}

func cookieTokenValidate(ctx *gin.Context, store *store, tokenType string) (*ent.Token, error) {
	tokenBody, err := ctx.Cookie(RefreshToken)
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
