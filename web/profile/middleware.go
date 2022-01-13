package profile

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
)

func TryProfile(ctx *gin.Context, store *data.Store) {
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

func MustProfile(ctx *gin.Context, store *data.Store) {
	_, ok := ctx.Get(value.StringProfile)
	if !ok {
		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code, errs.ProfileIdentityError)
		return
	}
	ctx.Next()
	return

}
