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

type Episode struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func createEpisode(ctx *gin.Context, store *db.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)

	var episode Episode
	if !utils.MustBindJSON(ctx, episode) {
		return
	}

	ep, err := store.CreateEpisode(ctx, episode.Title, episode.Content, profile.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusCreated, ep)

}

func updateEpisode(ctx *gin.Context, store *db.Store) {
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

	if !episodeBelongs(ctx, store, profile.ID, id.ID) {
		return
	}

	ep, err := store.UpdateEpisode(ctx, id.ID, episode.Title, episode.Content)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusOK, ep)

}

func getEpisode(ctx *gin.Context, store *db.Store) {
	var id ID
	if !utils.MustBindUri(ctx, id) {
		return
	}

	ep, err := store.FindEpisode(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.JSON(http.StatusOK, ep)

}

func deleteEpisode(ctx *gin.Context, store *db.Store) {
	profile := ctx.MustGet(value.StringProfile).(*ent.Profile)
	var id ID
	if !utils.MustBindUri(ctx, id) {
		return
	}

	if !episodeBelongs(ctx, store, profile.ID, id.ID) {
		return
	}

	err := store.DeleteEpisode(ctx, id.ID)
	if err != nil {
		e := errs.DatabaseError(err)
		ctx.JSON(e.Code, e)
		return
	}

	ctx.Status(http.StatusNoContent)

}

func episodeBelongs(ctx *gin.Context, store *db.Store, profileID, episodeID uuid.UUID) bool {
	belongs, err := store.EpisodeBelongs(ctx, profileID, episodeID)
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
