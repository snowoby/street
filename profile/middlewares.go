package profile

import (
	"github.com/gin-gonic/gin"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errors"
)

func Middleware(ctx *gin.Context, s *db.Store) {
	var id ID
	err := ctx.ShouldBindJSON(&id)
	if err != nil {
		e := errors.BindingError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
		return
	}

	p, err := s.FindProfileByID(ctx, id.ID)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
		return
	}
	account := ctx.MustGet(value.StringAccount).(*ent.Account)
	if p.Edges.Account.ID != account.ID {
		ctx.AbortWithStatusJSON(ProfileIdentityError.Code, ProfileIdentityError)
		return
	}

	ctx.Set(value.StringProfile, p)
	ctx.Next()
}
