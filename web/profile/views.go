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

type CallSign struct {
	CallSign string `json:"callSign" binding:"required"`
}

type Profile struct {
	Title    string `json:"title" binding:"required"`
	CallSign string `json:"callSign" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	var profile Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	exists, err := store.Profile.CallSignExists(ctx, profile.CallSign)
	if err != nil {
		return 0, nil, err
	}

	if exists {
		return 0, nil, errs.CallSignDuplicateError
	}

	p, err := store.Profile.Create(ctx, profile.CallSign, profile.Title, profile.Category, identity.Account().ID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, p, nil
}

func update(ctx *gin.Context, store *data.Store, _ *controller.Identity) (int, interface{}, error) {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	var profile Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	p, err := store.Profile.Update(ctx, objectID, profile.Title, profile.CallSign, profile.Category)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, p, nil

}

func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ps, err := store.Profile.FindByID(ctx, objectID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, ps, nil
}

func accountProfiles(_ *gin.Context, _ *data.Store, identity *controller.Identity) (int, interface{}, error) {
	return http.StatusOK, identity.AllProfiles(), nil
}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ok, err := store.Profile.IsOwner(ctx, operator.Profile().ID, objectID)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NotBelongsToOperator
	}
	return nil
}
