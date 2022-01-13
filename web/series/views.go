package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
	"street/pkg/utils"
)

type TitleContent struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func create(ctx *gin.Context, store *data.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)

	var series TitleContent
	if !utils.MustBindJSON(ctx, series) {
		return
	}

	ep, err := store.Series().Create(ctx, series.Title, series.Content, profile.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusCreated, ep)

}

func update(ctx *gin.Context, store *data.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)
	var id ID
	if !utils.MustBindUri(ctx, id) {
		return
	}

	var series TitleContent
	err := ctx.ShouldBindJSON(&series)
	if err != nil {
		e := errs.BindingError(err)
		ctx.JSON(e.Code, e)
		return
	}

	if !seriesMustBelong(ctx, store, profile.ID, id.ID) {
		return
	}

	ep, err := store.Series().Update(ctx, id.ID, series.Title, series.Content)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusOK, ep)

}

func get(ctx *gin.Context, store *data.Store) {
	var id ID
	if !utils.MustBindUri(ctx, id) {
		return
	}

	ep, err := store.Series().FindByID(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusOK, ep)

}

func del(ctx *gin.Context, store *data.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)
	var id ID
	if !utils.MustBindUri(ctx, id) {
		return
	}

	if !seriesMustBelong(ctx, store, profile.ID, id.ID) {
		return
	}

	err := store.Series().Delete(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.Status(http.StatusNoContent)

}

func seriesMustBelong(ctx *gin.Context, store *data.Store, profileID, seriesID uuid.UUID) bool {
	belongs, err := store.Series().IsOwner(ctx, profileID, seriesID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return false
	}

	if !belongs {
		ctx.JSON(errs.NotBelongsToOperator.Code, errs.NotBelongsToOperator)
		return false
	}
	return true
}
