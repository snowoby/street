package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/db"
	"street/db/value"
	"street/ent"
	"street/errors"
)

type CallSign struct {
	CallSign string `json:"call_sign" binding:"required"`
}

type Profile struct {
	Title string `json:"title" binding:"required"`
	CallSign
	Category string `json:"category" binding:"required"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func createProfile(ctx *gin.Context, store *db.Store) {
	account := ctx.MustGet(value.StringAccount).(*ent.Account)
	var profile Profile
	err := ctx.ShouldBind(&profile)
	if err != nil {
		e := errors.BindingError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}

	p, err := store.CreateProfile(ctx, profile.CallSign.CallSign, profile.Title, profile.Category, account.ID)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}
	ctx.JSON(http.StatusCreated, p)
}

func updateProfile(ctx *gin.Context, store *db.Store) {
	var id ID
	err := ctx.ShouldBindUri(&id)
	if err != nil {
		e := errors.BindingError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}

	var profile Profile
	err = ctx.ShouldBind(&profile)
	if err != nil {
		e := errors.BindingError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}

	p, err := store.UpdateProfile(ctx, id.ID, profile.Title, profile.CallSign.CallSign, profile.Category)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}

	ctx.JSON(http.StatusCreated, p)

}

func getProfile(ctx *gin.Context, store *db.Store) {
	var id ID
	err := ctx.ShouldBindUri(&id)
	if err != nil {
		e := errors.BindingError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}

	ps, err := store.FindProfileByID(ctx, id.ID)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}
	ctx.JSON(http.StatusOK, ps)
}

func accountProfiles(ctx *gin.Context, store *db.Store) {
	account := ctx.MustGet(value.StringAccount).(*ent.Account)

	ps, err := store.FindProfilesByAccountID(ctx, account.ID)
	if err != nil {
		e := errors.DatabaseError(err)
		ctx.AbortWithStatusJSON(e.Code, e)
	}
	ctx.JSON(http.StatusCreated, ps)
}
