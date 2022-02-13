package series

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

type TitleContent struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type Series struct {
	TitleContent
	ProfileID string `json:"profileID" binding:"required,uuid"`
}

type ResponseSeries struct {
	*ent.Series
	value.NoEdges
}

// create godoc
// @Summary create a series
// @Tags series
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param series body TitleContent true "series info"
// @Success 201 {object} ResponseSeries
// @Failure 400 {object} errs.HTTPError
// @Router /series/{pid} [post]
func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {

	var series Series
	err := ctx.ShouldBindJSON(&series)
	if err != nil {
		return 0, nil, err
	}

	s, err := store.DB.Series.Create(ctx, series.Title, series.Content, uuid.MustParse(series.ProfileID))
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, &ResponseSeries{Series: s}, nil
}

// update godoc
// @Summary update a series
// @Tags series
// @Accept json
// @Produce json
// @Param series body TitleContent true "series info"
// @Param pid path string true "profile id"
// @Param id path string true "series id"
// @Success 200 {object} ResponseSeries
// @Failure 400 {object} errs.HTTPError
// @Router /series/{pid}/{id} [put]
func update(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	var series TitleContent
	err := ctx.ShouldBindJSON(&series)
	if err != nil {
		return 0, nil, err
	}

	s, err := store.DB.Series.Update(ctx, id, series.Title, series.Content)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, &ResponseSeries{Series: s}, nil

}

// get godoc
// @Summary get a series
// @Tags series
// @Produce json
// @Param id path string true "series id"
// @Success 200 {object} ResponseSeries
// @Failure 400 {object} errs.HTTPError
// @Router /series/{id} [get]
func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	s, err := store.DB.Series.FindByIDWithProfile(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, &ResponseSeries{Series: s}, nil

}

// get godoc
// @Summary get all series
// @Tags series
// @Produce json
// @Param id path string true "series id"
// @Success 200 {object} []ResponseSeries
// @Failure 400 {object} errs.HTTPError
// @Router /series/ [get]
func getAll(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {

	ss, err := store.DB.Series.FindSeriesByAccount(ctx, identity.Account().ID)
	if err != nil {
		return 0, nil, err
	}
	reps := make([]*ResponseSeries, len(ss))

	for i, s := range ss {
		reps[i] = &ResponseSeries{
			Series: s,
		}
	}
	return http.StatusOK, &reps, nil

}

// del godoc
// @Summary del a series
// @Tags series
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "series id"
// @Success 204
// @Failure 400 {object} errs.HTTPError
// @Router /series/{id} [delete]
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
	id, err := store.DB.Series.OwnerID(ctx, objectID)
	if err != nil {
		return err
	}
	if operator.HaveProfile(uuid.MustParse(id)) {
		return errs.NotBelongsToOperator
	}
	return nil
}
