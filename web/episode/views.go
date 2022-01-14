package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

type Episode struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	profile := identity.Profile()

	var episode Episode
	err := ctx.ShouldBindJSON(&episode)
	if err != nil {
		return 0, nil, err
	}
	ep, err := store.Episode().Create(ctx, episode.Title, episode.Content, profile.ID)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusCreated, ep, nil

}

func update(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(*uuid.UUID)

	var episode Episode
	err := ctx.ShouldBindJSON(&episode)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	ep, err := store.Episode().Update(ctx, *id, episode.Title, episode.Content)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, ep, nil

}

func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(*uuid.UUID)

	ep, err := store.Episode().FindByID(ctx, *id)
	if err != nil {
		return 0, nil, err

	}

	return http.StatusOK, ep, nil

}

func del(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(*uuid.UUID)

	err := store.Episode().Delete(ctx, *id)
	if err != nil {
		return 0, nil, err

	}
	return http.StatusNoContent, nil, nil
}

//func episodeMustBelong(ctx *gin.Context, store *data.Store, profileID, episodeID uuid.UUID) bool {
//	belongs, err := store.Episode().IsOwner(ctx, profileID, episodeID)
//	if err != nil {
//		e := errs.DatabaseError(err)
//		ctx.JSON(e.Code, e)
//		return false
//	}
//
//	if !belongs {
//		ctx.JSON(errs.NotBelongsToOperator.Code, errs.NotBelongsToOperator)
//		return false
//	}
//	return true
//}
