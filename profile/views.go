package profile

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
	if !utils.MustBindJSON(ctx, &profile) {
		return
	}

	exists, err := store.CallSignExists(ctx, profile.CallSign.CallSign)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	if exists {
		ctx.JSON(errs.CallSignDuplicateError.Code, errs.CallSignDuplicateError)
		return
	}

	p, err := store.CreateProfile(ctx, profile.CallSign.CallSign, profile.Title, profile.Category, account.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
	}
	ctx.JSON(http.StatusCreated, p)
}

func updateProfile(ctx *gin.Context, store *db.Store) {
	var id ID
	if !utils.MustBindUri(ctx, &id) {
		return
	}

	var profile Profile
	if !utils.MustBindJSON(ctx, &profile) {
		return
	}

	p, err := store.UpdateProfile(ctx, id.ID, profile.Title, profile.CallSign.CallSign, profile.Category)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
	}

	ctx.JSON(http.StatusOK, p)

}

func getProfile(ctx *gin.Context, store *db.Store) {
	var id ID
	if !utils.MustBindUri(ctx, &id) {
		return
	}

	ps, err := store.FindProfileByID(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
	}
	ctx.JSON(http.StatusOK, ps)
}

func accountProfiles(ctx *gin.Context, store *db.Store) {
	account := ctx.MustGet(value.StringAccount).(*ent.Account)

	ps, err := store.FindProfilesByAccountID(ctx, account.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
	}
	ctx.JSON(http.StatusOK, ps)
}
