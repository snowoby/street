package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errors"
	"street/utils"
	"time"
)

func AccessToken(ctx *gin.Context, s *db.Store) {
	at, err := cookieTokenValidate(ctx, s, value.StringAccessToken)
	if err != nil {
		return
	}

	ctx.Set(value.StringAccount, at.Edges.Account)
	ctx.Set(value.StringAccessToken, at)
	ctx.Next()
	return

}

func RefreshToken(ctx *gin.Context, s *db.Store) {
	rt, err := cookieTokenValidate(ctx, s, value.StringRefreshToken)
	if err != nil {
		return
	}

	tokenBody := utils.RandomString(128)
	t, err := s.CreateToken(ctx, rt.Edges.Account.ID, tokenBody, value.StringAccessToken, s.Config().AccessTokenExpireTime)
	ctx.SetCookie(value.StringAccessToken, t.Body, int(t.ExpireTime.Sub(time.Now()).Seconds()), "/", s.Config().Domain, false, true)
	ctx.Status(http.StatusNoContent)
	return

}

func cookieTokenValidate(ctx *gin.Context, store *db.Store, tokenType string) (*ent.Token, error) {
	tokenBody, err := ctx.Cookie(tokenType)
	if err != nil {
		ctx.JSON(errors.TokenNotExistsError.Code, errors.TokenNotExistsError)
		return nil, err
	}

	t, err := store.FindToken(ctx, tokenBody, tokenType, true)
	if err != nil {
		databaseError := errors.DatabaseError(err)
		ctx.JSON(databaseError.Code, databaseError)
		return nil, err
	}

	if t.ExpireTime.Before(time.Now()) {
		ctx.JSON(errors.TokenExpiredError.Code, errors.TokenExpiredError)
		return nil, err
	}

	return t, nil
}
