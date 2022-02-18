package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"street/ent"
	"street/ent/comment"
	"street/ent/episode"
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

func New(db *ent.Client, auth auth.Service, router *gin.RouterGroup) *service {
	s := &service{
		db:     db,
		auth:   auth,
		router: router,
	}
	s.registerRouters()
	return s
}

func (s *service) registerRouters() {
	s.router.GET("/:id", composer.ID(s.getByEpisodeID))
	s.router.Use(s.auth.MustLogin)
	s.router.POST("/", composer.Authed(s.create))
	s.router.PUT("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.update))
	s.router.DELETE("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.delete))

}

func (s *service) getByEpisodeID(ctx *gin.Context, id string) (int, interface{}, error) {
	comments, err := s.db.Comment.Query().Where(comment.HasEpisodeWith(episode.ID(uuid.MustParse(id)))).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, d.CommentsFromEnt(comments), nil
}

func (s *service) create(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	var commentForm d.CommentForm
	if err := ctx.ShouldBindJSON(&commentForm); err != nil {
		return 0, nil, err
	}
	if err := operator.HaveProfileX(ctx, commentForm.From); err != nil {
		return 0, nil, err
	}

	c, err := s.db.Comment.Create().
		SetContent(commentForm.Content).
		SetAuthorID(uuid.MustParse(commentForm.From)).
		SetEpisodeID(uuid.MustParse(commentForm.To)).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, d.CommentFromEnt(c), nil
}

func (s *service) update(ctx *gin.Context, id string) (int, interface{}, error) {
	var commentForm d.CommentForm
	if err := ctx.ShouldBindJSON(&commentForm); err != nil {
		return 0, nil, err
	}

	c, err := s.db.Comment.UpdateOneID(uuid.MustParse(id)).
		SetContent(commentForm.Content).
		SetAuthorID(uuid.MustParse(commentForm.From)).
		SetEpisodeID(uuid.MustParse(commentForm.To)).
		Save(ctx)

	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.CommentFromEnt(c), nil
}

func (s *service) delete(ctx *gin.Context, id string) (int, interface{}, error) {

	err := s.db.Comment.DeleteOneID(uuid.MustParse(id)).Exec(ctx)

	if err != nil {
		return 0, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (s *service) owned(ctx *gin.Context, operator *operator.Identity, objID string) error {
	id, err := s.db.Comment.Query().Where(comment.ID(uuid.MustParse(objID))).QueryAuthor().OnlyID(ctx)
	if err != nil {
		return err
	}

	err = operator.HaveProfileX(ctx, id.String())
	return err
}
