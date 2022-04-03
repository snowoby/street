package comment

import (
	"net/http"
	"street/ent"
	"street/ent/comment"
	"street/ent/episode"
	"street/pkg/auth"
	"street/pkg/composer"
	"street/pkg/d"
	"street/pkg/operator"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

const PathSplitter = "/"

func (s *service) registerRouters() {
	s.router.GET("/episode/:id", composer.ID(s.getByEpisodeID))
	s.router.GET("/family/:id", composer.ID(s.getFamilyByCommentID))
	s.router.Use(s.auth.MustLogin)
	s.router.POST("/episode/:id", composer.AuthedID(s.create))
	s.router.PUT("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.update))
	s.router.DELETE("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.delete))
}

func (s *service) getByEpisodeID(ctx *gin.Context, id string) (int, interface{}, error) {
	comments, err := s.db.Comment.Query().Where(comment.HasEpisodeWith(episode.ID(uuid.MustParse(id)))).WithAuthor().Order(ent.Desc(comment.FieldCreateTime)).All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return http.StatusOK, d.CommentsFromEnt(comments), nil
}

func (s *service) getFamilyByCommentID(ctx *gin.Context, id string) (int, interface{}, error) {
	c, err := s.db.Comment.Query().Where(comment.ID(uuid.MustParse(id))).WithAuthor().Only(ctx)
	if err != nil {
		return 0, nil, err
	}

	parents := []*ent.Comment{}
	if len(c.Path) != 0 {
		strParentsID := strings.Split(c.Path, PathSplitter)
		parentsID := make([]uuid.UUID, len(strParentsID))
		for i, strID := range strParentsID {
			parentsID[i] = uuid.MustParse(strID)
		}
		parents, err = s.db.Comment.Query().Where(comment.IDIn(parentsID...)).WithAuthor().Order(ent.Asc(comment.FieldCreateTime)).All(ctx)
		if err != nil {
			return 0, nil, err
		}
	}

	children, err := c.QueryReplied().WithAuthor().Order(ent.Asc(comment.FieldCreateTime)).All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.CommentsFromEnt(append(append(parents, c), children...)), nil
}

func (s *service) create(ctx *gin.Context, operator *operator.Identity, id string) (int, interface{}, error) {
	var commentForm d.CommentForm
	if err := ctx.ShouldBindJSON(&commentForm); err != nil {
		return 0, nil, err
	}
	if err := operator.HaveProfileX(ctx, commentForm.From); err != nil {
		return 0, nil, err
	}

	tmpC := s.db.Comment.Create().
		SetContent(commentForm.Content).
		SetAuthorID(uuid.MustParse(commentForm.From)).
		SetEpisodeID(uuid.MustParse(id))

	if commentForm.ReplyTo != nil {
		tmpC = tmpC.SetReplyToID(uuid.MustParse(*commentForm.ReplyTo))
		replyTo, err := s.db.Comment.Get(ctx, uuid.MustParse(*commentForm.ReplyTo))
		if err != nil {
			return 0, nil, err
		}
		newPath := *commentForm.ReplyTo
		if replyTo.Path != "" {
			newPath = replyTo.Path + PathSplitter + newPath
		}
		tmpC = tmpC.SetPath(newPath)
	}

	c, err := tmpC.Save(ctx)

	if err != nil {
		return 0, nil, err
	}
	p, err := c.QueryAuthor().Only(ctx)
	if err != nil {
		return 0, nil, err
	}
	c.Edges.Author = p
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
