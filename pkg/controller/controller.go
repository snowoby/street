package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
	"street/ent"
	"street/errs"
	"street/pkg/data"
	"street/pkg/data/value"
)

type OriginalF func(ctx *gin.Context, store *data.Store)
type F func(ctx *gin.Context, store *data.Store) (int, interface{}, error)
type Func func(ctx *gin.Context, store *data.Store, visitor *Identity) (int, interface{}, error)
type TaskFunc func(ctx context.Context, task *asynq.Task, store *data.Store) error
type TaskHandleFunc func(context.Context, *asynq.Task) error
type OwnerFunc func(ctx *gin.Context, store *data.Store, visitor *Identity) error

type controller struct {
	store *data.Store
}

type Controller interface {
	Original(f OriginalF) gin.HandlerFunc
	Bare(f F) gin.HandlerFunc
	// General will try to bind visitor identity as param.
	General(f Func) gin.HandlerFunc
	Task(f TaskFunc) TaskHandleFunc
	Store() *data.Store
	Owned(f OwnerFunc) gin.HandlerFunc
}

func New(store *data.Store) *controller {
	return &controller{store: store}
}

func resultProcess(ctx *gin.Context, code int, responseValue interface{}, err error) {
	if err != nil {
		responseError := errs.Detect(err)
		ctx.AbortWithStatusJSON(responseError.Code, struct {
			Message string `json:"message"`
		}{responseError.Message})
		return
	}

	ctx.JSON(code, responseValue)
}

func (controller *controller) Store() *data.Store {
	return controller.store
}

func extractOperator(ctx *gin.Context) *Identity {
	a, _ := ctx.Get(value.StringAccount)
	p, _ := ctx.Get(value.StringProfile)
	ps, _ := ctx.Get(value.StringAllProfiles)
	var account *ent.Account
	if a != nil {
		account = a.(*ent.Account)
	}
	var profile *ent.Profile
	if p != nil {
		profile = p.(*ent.Profile)
	}
	var allProfile []*ent.Profile
	if ps != nil {
		allProfile = ps.([]*ent.Profile)
	}

	return &Identity{
		account:    account,
		profile:    profile,
		allProfile: allProfile,
	}
}

func (controller *controller) Original(f OriginalF) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		f(ctx, controller.store)
	}
}
func (controller *controller) Bare(f F) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, responseValue, err := f(ctx, controller.store)
		resultProcess(ctx, code, responseValue, err)

	}
}
func (controller *controller) Task(f TaskFunc) TaskHandleFunc {
	return func(ctx context.Context, task *asynq.Task) error {
		return f(ctx, task, controller.store)
	}
}

func (controller *controller) General(nf Func) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := extractOperator(ctx)
		id.profile, _ = controller.store.DB.Profile.Client().Query().First(ctx)
		code, responseValue, err := nf(ctx, controller.store, id)
		resultProcess(ctx, code, responseValue, err)
	}
}

//
//func (controller *controller) GeneralUUID(f UUIDFunc) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		identity := extractOperator(ctx)
//		id, _ := extractUUID(ctx)
//
//		code, responseValue, err := f(ctx, controller.store, identity, id)
//
//		if err != nil {
//			responseError := errs.Detect(err)
//			ctx.AbortWithStatusJSON(responseError.Code, responseError.Message)
//			return
//		}
//
//		ctx.JSON(code, responseValue)
//	}
//
//}

func (controller *controller) Owned(f OwnerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := extractOperator(ctx)
		err := f(ctx, controller.store, id)
		if err != nil {
			responseError := errs.Detect(err)
			ctx.AbortWithStatusJSON(responseError.Code, responseError.Message)
			return
		}
		ctx.Next()
	}
}
