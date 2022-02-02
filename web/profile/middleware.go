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
		ID string `uri:"pid" binding:"uuid,required"`
	}

	ps, err := store.FindProfilesByAccountID(ctx, account.ID)
	if err != nil {
		ctx.Next()
		return
	}
	ctx.Set(value.StringAllProfiles, ps)

	var idString ID
	err = ctx.ShouldBindUri(&idString)
	if err != nil {
		ctx.Next()
		return
	}

	id, _ := uuid.Parse(idString.ID)
	for _, profile := range ps {
		if profile.ID == id {
			ctx.Set(value.StringProfile, profile)
		}
	}
	ctx.Next()
}

func MustProfile(ctx *gin.Context) {
	_, ok := ctx.Get(value.StringProfile)
	if !ok {
		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code, errs.ProfileIdentityError)
		return
	}
	ctx.Next()
	return

}
