package series

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

type TitleContent struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid" json:"id"`
}

func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	profile := identity.Profile()

	var series TitleContent
	err := ctx.ShouldBindJSON(&series)
	if err != nil {
		return 0, nil, err
	}

	ep, err := store.DB.Series.Create(ctx, series.Title, series.Content, profile.ID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, ep, nil
}

func update(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	var series TitleContent
	err := ctx.ShouldBindJSON(&series)
	if err != nil {
		return 0, nil, err
	}

	ep, err := store.DB.Series.Update(ctx, id, series.Title, series.Content)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, ep, nil

}

func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	ep, err := store.DB.Series.FindByID(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, ep, nil

}

func del(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	err := store.DB.Series.Delete(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusNoContent, nil, nil

}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ok, err := store.DB.Series.IsOwner(ctx, operator.Profile().ID, objectID)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NotBelongsToOperator
	}
	return nil
}
