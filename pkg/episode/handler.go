package episode

import (
	"net/http"
	"street/ent"
	"street/ent/episode"
	"street/errs"
	"street/pkg/auth"
	"street/pkg/composer"
	"street/pkg/d"
	"street/pkg/operator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	db     *ent.Client
	auth   auth.Service
	router *gin.RouterGroup
}

func (s *service) registerRouters() {
	s.router.GET("/:id", composer.ID(s.get))
	s.router.GET("/", composer.Bare(s.getAll))
	s.router.Use(s.auth.MustLogin)
	s.router.POST("/", composer.Authed(s.create))
	s.router.PUT("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.update))
	s.router.DELETE("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.delete))

}

func New(db *ent.Client, auth auth.Service, router *gin.RouterGroup) *service {
	s := &service{
		db:     db,
		auth:   auth,
		router: router,
	}
	s.registerRouters()
	return s
}

// create godoc
// @Summary create episode
// @Tags episode
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param episode body d.EpisodeForm true "episode info"
// @Success 201 {object} d.Episode
// @Failure 400 {object} errs.HTTPError
// @Router /episode [post]
func (s *service) create(ctx *gin.Context, identity *operator.Identity) (int, interface{}, error) {

	var episodeForm d.EpisodeForm
	err := ctx.ShouldBindJSON(&episodeForm)
	if err != nil {
		return 0, nil, err
	}

	if !identity.HaveProfile(ctx, episodeForm.ProfileID) {
		return 0, nil, errs.UnauthorizedError
	}

	episodeBuilder := s.db.Episode.Create().
		SetContent(episodeForm.Content).
		SetProfileID(uuid.MustParse(episodeForm.ProfileID)).
		SetFiles(episodeForm.Files)

	if episodeForm.SeriesID != nil {
		episodeBuilder = episodeBuilder.SetSeriesID(uuid.MustParse(*episodeForm.SeriesID))
	}
	if episodeForm.Title != nil {
		episodeBuilder = episodeBuilder.SetTitle(*episodeForm.Title)
	}
	if episodeForm.Cover != nil {
		episodeBuilder = episodeBuilder.SetCover(*episodeForm.Cover)
	}

	ep, err := episodeBuilder.Save(ctx)

	if err != nil {
		return 0, nil, err
	}
	return http.StatusCreated, d.EpisodeFromEnt(ep), nil

}

// update godoc
// @Summary update episode
// @Tags episode
// @Accept json
// @Produce json
// @Param id path string true "episode id"
// @Param pid path string true "profile id"
// @Param episode body d.EpisodeForm true "episode info"
// @Success 200 {object} d.Episode
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{id} [put]
func (s *service) update(ctx *gin.Context, id string) (int, interface{}, error) {
	var episodeForm d.EpisodeForm
	err := ctx.ShouldBindJSON(&episodeForm)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	episodeBuilder := s.db.Episode.UpdateOneID(uuid.MustParse(id)).
		SetContent(episodeForm.Content).
		SetFiles(episodeForm.Files)

	if episodeForm.SeriesID != nil {
		episodeBuilder = episodeBuilder.SetSeriesID(uuid.MustParse(*episodeForm.SeriesID))
	}
	if episodeForm.Title != nil {
		episodeBuilder = episodeBuilder.SetTitle(*episodeForm.Title)
	}
	if episodeForm.Cover != nil {
		episodeBuilder = episodeBuilder.SetCover(*episodeForm.Cover)
	}
	ep, err := episodeBuilder.Save(ctx)

	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.EpisodeFromEnt(ep), nil
}

// get godoc
// @Summary get episode
// @Tags episode
// @Accept json
// @Produce json
// @Param id path string true "episode id"
// @Success 200 {object} d.Episode
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{id} [get]
func (s *service) get(ctx *gin.Context, id string) (int, interface{}, error) {

	ep, err := s.db.Episode.Query().
		Where(episode.ID(uuid.MustParse(id))).
		WithProfile().
		WithSeries().
		Only(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.EpisodeFromEnt(ep), nil

}

// getAll godoc
// @Summary get all episodes
// @Tags episode
// @Produce json
// @Success 200 {object} []d.Episode
// @Failure 400 {object} errs.HTTPError
// @Router /episode [get]
func (s *service) getAll(ctx *gin.Context) (int, interface{}, error) {
	eps, err := s.db.Episode.Query().
		WithProfile().
		WithSeries().
		Order(ent.Desc(episode.FieldCreateTime)).
		All(ctx)
	if err != nil {
		return 0, nil, err

	}

	return http.StatusOK, d.EpisodesFromEnt(eps), nil

}

// delete godoc
// @Summary delete one episode
// @Tags episode
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "episode id"
// @Success 204
// @Failure 400 {object} errs.HTTPError
// @Router /episode/{id} [delete]
func (s *service) delete(ctx *gin.Context, id string) (int, interface{}, error) {
	err := s.db.Episode.DeleteOneID(uuid.MustParse(id)).Exec(ctx)
	if err != nil {
		return 0, nil, err

	}
	return http.StatusNoContent, nil, nil
}

func (s *service) owned(ctx *gin.Context, operator *operator.Identity, objID string) error {
	id, err := s.db.Episode.Query().Where(episode.ID(uuid.MustParse(objID))).QueryProfile().OnlyID(ctx)
	if err != nil {
		return err
	}

	err = operator.HaveProfileX(ctx, id.String())
	return err
}
