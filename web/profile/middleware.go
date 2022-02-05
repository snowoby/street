package profile

import (
	"github.com/gin-gonic/gin"
	"street/ent"
	"street/errs"
	"street/pkg/data/value"
)

func TryProfile(ctx *gin.Context) {
	a := ctx.MustGet(value.StringAccount)

	account := a.(*ent.Account)

	ps, err := account.QueryProfile().All(ctx)
	if err != nil {
		ctx.Next()
		return
	}

	ctx.Set(value.StringAllProfiles, ps)
	ctx.Next()
}

func MustHaveProfile(ctx *gin.Context) {
	ps, ok := ctx.Get(value.StringAllProfiles)
	if !ok {
		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code, errs.ProfileIdentityError)
		return
	}

	profiles := ps.([]*ent.Profile)

	if len(profiles) == 0 {
		ctx.AbortWithStatusJSON(errs.NoProfiles.Code, errs.NoProfiles)
		return
	}

	ctx.Next()
	return

}

//func MustUseProfile(ctx *gin.Context) {
//	_, ok := ctx.Get(value.StringProfile)
//	if !ok {
//		ctx.AbortWithStatusJSON(errs.ProfileIdentityError.Code, errs.ProfileIdentityError)
//		return
//	}
//	ctx.Next()
//	return
//
//}
