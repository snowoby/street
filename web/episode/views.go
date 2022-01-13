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

type Episode struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func create(ctx *gin.Context, store *data.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)

	var episode Episode
	if !utils.MustBindJSON(ctx, episode) {
		return
	}

	ep, err := store.Episode().Create(ctx, episode.Title, episode.Content, profile.ID)
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

	var episode Episode
	err := ctx.ShouldBindJSON(&episode)
	if err != nil {
		e := errs.BindingError(err)
		ctx.JSON(e.Code, e)
		return
	}

	if !episodeMustBelong(ctx, store, profile.ID, id.ID) {
		return
	}

	ep, err := store.Episode().Update(ctx, id.ID, episode.Title, episode.Content)
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

	ep, err := store.Episode().FindByID(ctx, id.ID)
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

	if !episodeMustBelong(ctx, store, profile.ID, id.ID) {
		return
	}

	err := store.Episode().Delete(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.Status(http.StatusNoContent)

}

func episodeMustBelong(ctx *gin.Context, store *data.Store, profileID, episodeID uuid.UUID) bool {
	belongs, err := store.Episode().IsOwner(ctx, profileID, episodeID)
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
