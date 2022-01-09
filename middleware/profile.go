package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errors"
)

type ID struct {
	ID uuid.UUID `binding:"uuid" json:"profile_id"`
}

func Profile(ctx *gin.Context, s *db.Store) {
	//TODO try to bind token and profile -> middleware handle

	var id ID
	err := ctx.ShouldBindJSON(&id)
	if err != nil {
		e := errors.BindingError(err)
		ctx.JSON(e.Code, e)
		return
	}

	p, err := s.FindProfileByID(ctx, id.ID)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}
	account := ctx.MustGet(value.StringAccount).(*ent.Account)
	if p.Edges.Account.ID != account.ID {
		ctx.JSON(errors.ProfileIdentityError.Code, errors.ProfileIdentityError)
		return
	}

	ctx.Set(value.StringProfile, p)
	ctx.Next()
}
