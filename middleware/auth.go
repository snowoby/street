package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errs"
	"street/utils"
)

func MustRefresh(ctx *gin.Context, store *db.Store) {
	tokenType := value.StringRefreshToken
	t := tryToken(ctx, store, tokenType)
	if t == nil {
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code, errs.UnauthorizedError)
		return
	}

	tokenBody := utils.RandomString(128)
	// Create an access token
	t, err := store.CreateToken(ctx, t.Edges.Account.ID, tokenBody, value.StringAccessToken, store.Config().RefreshTokenExpireTime)
	if err != nil {
		databaseError := errs.DatabaseError(err)
		ctx.AbortWithStatusJSON(databaseError.Code, databaseError)
		return
	}

	ctx.JSON(http.StatusCreated, t)
}

func TryAccessToken(ctx *gin.Context, store *db.Store) {
	t := tryToken(ctx, store, value.StringAccessToken)
	if t != nil {
		ctx.Set(value.StringAccount, t.Edges.Account)
		ctx.Set(value.StringAccessToken, t)
	}

	ctx.Next()
}

func tryToken(ctx *gin.Context, store *db.Store, tokenType string) *ent.Token {
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

type ID struct {
	ID uuid.UUID `binding:"uuid" header:"Profile"`
}

func TryProfile(ctx *gin.Context, store *db.Store) {
	var id ID
	err := ctx.ShouldBindHeader(&id)
	if err == nil {
		p, err := store.FindProfileByID(ctx, id.ID)
		if err == nil {
			a, ok := ctx.Get(value.StringAccount)
			if ok {
				account := a.(*ent.Account)
				if p.Edges.Account.ID == account.ID {
					ctx.Set(value.StringProfile, p)
				}
			}
		}
	}
	ctx.Next()
}

func MustLogin(ctx *gin.Context, store *db.Store) {
	_, ok := ctx.Get(value.StringAccount)
	if !ok {
		ctx.AbortWithStatusJSON(errs.UnauthorizedError.Code, errs.UnauthorizedError)
		return
	}
	ctx.Next()
	return

}

func MustProfile(ctx *gin.Context, store *db.Store) {
	_, ok := ctx.Get(value.StringProfile)
	if !ok {
		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code, errs.ProfileIdentityError)
		return
	}
	ctx.Next()
	return

}
