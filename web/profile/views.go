package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

type Call struct {
	Call string `json:"call" binding:"required"`
}

type Profile struct {
	Title    string `json:"title" binding:"required"`
	Call     string `json:"call" binding:"required"`
	Category string `json:"category" binding:"required"`
	Avatar   string `json:"avatar"`
}

type ResponseProfile struct {
	*ent.Profile
	value.NoEdges
}

// create godoc
// @Summary create profile
// @Tags profile
// @Accept json
// @Produce json
// @Param profile body Profile true "profile info"
// @Success 201 {object} ResponseProfile
// @Failure 400 {object} errs.HTTPError
// @Router /profile [post]
func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	var profile Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	exists, err := store.DB.Profile.CallExists(ctx, profile.Call)
	if err != nil {
		return 0, nil, err
	}

	if exists {
		return 0, nil, errs.CallDuplicateError
	}

	p, err := store.DB.Profile.Create(ctx, profile.Call, profile.Title, profile.Category, profile.Avatar, identity.Account().ID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, &ResponseProfile{Profile: p}, nil
}

// create godoc
// @Summary update profile
// @Tags profile
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param profile body Profile true "profile info"
// @Success 201 {object} ResponseProfile
// @Failure 400 {object} errs.HTTPError
// @Router /profile/{pid} [put]
func update(ctx *gin.Context, store *data.Store, _ *controller.Identity) (int, interface{}, error) {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	var profile Profile
	err := ctx.ShouldBindJSON(&profile)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	p, err := store.DB.Profile.Update(ctx, objectID, profile.Title, profile.Call, profile.Category, profile.Avatar)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, &ResponseProfile{Profile: p}, nil

}

// get godoc
// @Summary get profile
// @Tags profile
// @Produce json
// @Param pid path string true "profile id"
// @Success 200 {object} ResponseProfile
// @Failure 400 {object} errs.HTTPError
// @Router /profile/{pid} [get]
func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	p, err := store.DB.Profile.FindByIDWithAccount(ctx, objectID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, &ResponseProfile{Profile: p}, nil
}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	id, err := store.DB.Profile.OwnerID(ctx, objectID)
	if err != nil {
		return err
	}
	if id != operator.Account().ID.String() {
		return errs.NotBelongsToOperator
	}
	return nil
}
