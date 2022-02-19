package series

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/ent/account"
	"street/ent/profile"
	"street/ent/series"
	"street/pkg/auth"
	"street/pkg/composer"
	"street/pkg/d"
	"street/pkg/operator"
)

type service struct {
	db     *ent.Client
	auth   auth.Service
	router *gin.RouterGroup
}

func (s *service) registerRouters() {
	s.router.GET("/:id", composer.ID(s.get))
	s.router.Use(s.auth.MustLogin)
	s.router.GET("/", composer.Authed(s.getAll))
	s.router.POST("/", composer.Authed(s.create))
	s.router.PUT("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.update))

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

func (s *service) get(ctx *gin.Context, id string) (int, interface{}, error) {
	se, err := s.db.Series.Query().Where(series.ID(uuid.MustParse(id))).WithOwner().Only(ctx)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, d.SeriesFromEnt(se), nil
}

func (s *service) getAll(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	ses, err := s.db.Series.Query().
		Where(
			series.HasOwnerWith(
				profile.HasAccountWith(
					account.ID(operator.Account().ID),
				),
			),
		).
		WithOwner().
		All(ctx)

	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.SeriesManyFromEnt(ses), nil
}

func (s *service) create(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	var seriesForm d.SeriesForm

	if err := ctx.ShouldBindJSON(&seriesForm); err != nil {
		return 0, nil, err
	}

	err := operator.HaveProfileX(ctx, seriesForm.ProfileID)
	if err != nil {
		return 0, nil, err
	}

	se, err := s.db.Series.Create().SetTitle(seriesForm.Title).SetOwnerID(uuid.MustParse(seriesForm.ProfileID)).SetType(seriesForm.Type).Save(ctx)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusCreated, d.SeriesFromEnt(se), nil
}

func (s *service) update(ctx *gin.Context, id string) (int, interface{}, error) {
	var seriesForm d.SeriesForm

	if err := ctx.ShouldBindJSON(&seriesForm); err != nil {
		return 0, nil, err
	}

	se, err := s.db.Series.UpdateOneID(uuid.MustParse(id)).SetTitle(seriesForm.Title).SetType(seriesForm.Type).Save(ctx)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusCreated, d.SeriesFromEnt(se), nil
}

func (s *service) owned(ctx *gin.Context, operator *operator.Identity, objID string) error {
	id, err := s.db.Series.Query().Where(series.ID(uuid.MustParse(objID))).QueryOwner().OnlyID(ctx)
	if err != nil {
		return err
	}

	err = operator.HaveProfileX(ctx, id.String())
	return err
}
