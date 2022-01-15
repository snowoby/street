package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
)

func TryProfile(ctx *gin.Context, store *data.Store) {
	a, ok := ctx.Get(value.StringAccount)
	if !ok {
		ctx.Next()
		return
	}

	account := a.(*ent.Account)
	type ID struct {
		ID uuid.UUID `binding:"uuid" header:"Profile" uri:"id"`
	}

	var id ID
	err := ctx.ShouldBindHeader(&id)
	if err == nil {
		p, err := store.FindProfilesByAccountID(ctx, account.ID)
		if err == nil {
			for _, profile := range p {
				if profile.ID == id.ID {
					ctx.Set(value.StringProfile, profile)
				}
			}
			ctx.Set(value.StringAllProfiles, p)
		}
	}
	ctx.Next()
}

func MustProfile(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringProfile)
	if !ok {
		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code(), errs.ProfileIdentityError)
		return
	}
	ctx.Next()
	return

}
