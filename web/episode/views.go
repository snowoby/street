package episode

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

type Episode struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type ResponseEpisode struct {
	*ent.Episode
	value.NoEdges
}

// create godoc
// @Summary create episode
// @Tags episode
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param episode body Episode true "episode info"
// @Success 201 {object} ResponseEpisode
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{pid} [post]
func create(ctx *gin.Context, store *data.Store, identity *controller.Identity) (int, interface{}, error) {
	profile := identity.Profile()

	var episode Episode
	err := ctx.ShouldBindJSON(&episode)
	if err != nil {
		return 0, nil, err
	}
	ep, err := store.DB.Episode.Create(ctx, episode.Title, episode.Content, profile.ID)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusCreated, &ResponseEpisode{
		Episode: ep,
	}, nil

}

// update godoc
// @Summary update episode
// @Tags episode
// @Accept json
// @Produce json
// @Param id path string true "episode id"
// @Param pid path string true "profile id"
// @Param episode body Episode true "episode info"
// @Success 200 {object} ResponseEpisode
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{pid}/{id} [put]
func update(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	var episode Episode
	err := ctx.ShouldBindJSON(&episode)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	ep, err := store.DB.Episode.Update(ctx, id, episode.Title, episode.Content)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, &ResponseEpisode{
		Episode: ep,
	}, nil

}

// get godoc
// @Summary get episode
// @Tags episode
// @Accept json
// @Produce json
// @Param id path string true "episode id"
// @Success 200 {object} ResponseEpisode
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{id} [get]
func get(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	ep, err := store.DB.Episode.FindByID(ctx, id)
	if err != nil {
		return 0, nil, err

	}

	return http.StatusOK, &ResponseEpisode{
		Episode: ep,
	}, nil

}

// getAll godoc
// @Summary get all episodes
// @Tags episode
// @Produce json
// @Success 200 {object} []ResponseEpisode
// @Failure 400 {object} errs.HTTPError
// @Router /episode [get]
func getAll(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	eps, err := store.DB.Episode.All(ctx)
	reps := make([]*ResponseEpisode, len(eps))
	for i, ep := range eps {
		reps[i] = &ResponseEpisode{Episode: ep}
	}
	if err != nil {
		return 0, nil, err

	}

	return http.StatusOK, reps, nil

}

// del godoc
// @Summary delete one episode
// @Tags episode
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "episode id"
// @Success 204
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{pid}/{id} [delete]
func del(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	err := store.DB.Episode.Delete(ctx, id)
	if err != nil {
		return 0, nil, err

	}
	return http.StatusNoContent, nil, nil
}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ok, err := store.DB.Episode.IsOwner(ctx, operator.Profile().ID, objectID)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NotBelongsToOperator
	}
	return nil
}
