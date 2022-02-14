package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/ent/profile"
	"street/errs"
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

// create godoc
// @Summary create profile
// @Tags profile
// @Accept json
// @Produce json
// @Param profile body d.ProfileForm true "profile info"
// @Success 201 {object} d.Profile
// @Failure 400 {object} errs.HTTPError
// @Router /profile [post]
func (s *service) create(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	var profileForm d.ProfileForm
	err := ctx.ShouldBindJSON(&profileForm)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	exists, err := s.db.Profile.Query().Where(profile.Call(profileForm.Call)).Exist(ctx)
	if err != nil {
		return 0, nil, err
	}
	if exists {
		return 0, nil, errs.CallDuplicateError
	}

	p, err := s.db.Profile.Create().
		SetTitle(profileForm.Title).
		SetCall(profileForm.Call).
		SetCategory(profileForm.Category).
		SetAvatar(profileForm.Avatar).
		SetAccountID(operator.Account().ID).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, d.ProfileFromEnt(p), nil
}

// create godoc
// @Summary update profile
// @Tags profile
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param profile body d.ProfileForm true "profile info"
// @Success 201 {object} d.Profile
// @Failure 400 {object} errs.HTTPError
// @Router /profile/{pid} [put]
func (s *service) update(ctx *gin.Context, id string) (int, interface{}, error) {
	var profileForm d.ProfileForm
	err := ctx.ShouldBindJSON(&profileForm)
	if err != nil {
		return 0, nil, errs.BindingError(err)
	}

	p, err := s.db.Profile.UpdateOneID(uuid.MustParse(id)).
		SetTitle(profileForm.Title).
		SetCall(profileForm.Call).
		SetCategory(profileForm.Category).
		SetAvatar(profileForm.Avatar).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.ProfileFromEnt(p), nil

}

// get godoc
// @Summary get profile
// @Tags profile
// @Produce json
// @Param pid path string true "profile id"
// @Success 200 {object} d.Profile
// @Failure 400 {object} errs.HTTPError
// @Router /profile/{pid} [get]
func (s *service) get(ctx *gin.Context, id string) (int, interface{}, error) {
	p, err := s.db.Profile.Query().Where(profile.ID(uuid.MustParse(id))).WithAccount().Only(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.ProfileFromEnt(p), nil
}

func (s *service) owned(ctx *gin.Context, operator *operator.Identity, objID string) error {
	profileData, err := s.db.Profile.Get(ctx, uuid.MustParse(objID))
	if err != nil {
		return err
	}
	if profileData.Edges.Account.ID.String() != operator.Account().ID.String() {
		return errs.NotBelongsToOperator
	}
	return nil
}
