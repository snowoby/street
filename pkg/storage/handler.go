package storage

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"street/ent"
	"street/ent/file"
	"street/errs"
	"street/pkg/auth"
	"street/pkg/composer"
	"street/pkg/d"
	"street/pkg/operator"
	"strings"
)

type service struct {
	db             *ent.Client
	auth           auth.Service
	redisService   *redisService
	router         *gin.RouterGroup
	tasker         *taskService
	storageService *storageService
}

func (s *service) registerRouters() {
	s.router.Use(s.auth.MustLogin)
	single := s.router.Group("single")
	single.POST("/", composer.Authed(s.create))
	single.PUT("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.put))

	large := s.router.Group("large")
	large.POST("/", composer.Authed(s.createMulti))
	large.PUT("/:id/:part_id", composer.AuthedIDCheck(s.owned), composer.ID(s.putMulti))
	large.POST("/:id", composer.AuthedIDCheck(s.owned), composer.ID(s.doneMulti))

}

func New(db *ent.Client,
	auth auth.Service,
	redisClient *redis.Client,
	router *gin.RouterGroup,
	asynqClient *asynq.Client,
	s3Config *aws.Config) *service {
	s := &service{
		db:             db,
		auth:           auth,
		router:         router,
		redisService:   newRedis(redisClient),
		tasker:         newTasker(asynqClient),
		storageService: newS3(s3Config),
	}

	s.registerRouters()
	return s
}

// createSingle godoc
// @Summary create single file upload
// @Tags file
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param meta body d.FileForm true "file meta"
// @Success 201 {object} d.File
// @Failure 400 {object} errs.HTTPError
// @Router /file/single/{pid} [post]
func (s *service) create(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	var meta d.FileForm
	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		return 0, nil, err
	}

	file, err := s.db.File.Create().
		SetFilename(meta.Filename).
		SetMime(meta.Mime).
		SetSize(meta.Size).
		SetAccountID(operator.Account().ID).
		SetPath(meta.Category).
		Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, d.FileFromEnt(file), nil
}

// putSingle godoc
// @Summary single file content upload
// @Tags file
// @Accept application/octet-stream
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "file id"
// @Param data body array true "file content"
// @Success 200 {object} d.File
// @Failure 400 {object} errs.HTTPError
// @Router /file/single/{id} [put]
func (s *service) put(ctx *gin.Context, id string) (int, interface{}, error) {

	file, err := s.db.File.Get(ctx, uuid.MustParse(id))
	if err != nil {
		return 0, nil, err
	}

	if file.Status != "created" {
		return 0, nil, errs.FileUploadedError
	}

	raw, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := bytes.NewReader(raw)
	_, err = s.storageService.PutSingle(reader, file.Path, uuid.MustParse(id), file.Filename, "original", file.Mime)
	if err != nil {
		return 0, nil, err
	}

	file, err = s.db.File.UpdateOneID(uuid.MustParse(id)).SetStatus("uploaded").Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	if strings.HasPrefix(file.Mime, "image/") {
		if file.Path == "avatar" {
			_, err = s.tasker.avatarCompress(file)
		} else {
			_, err = s.tasker.imageCompress(file)
		}

	}
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, d.FileFromEnt(file), nil

}

type CreateOutput struct {
	Key      string `json:"key" binding:"required"`
	UploadId string `json:"upload_id" binding:"required"`
}

// createMulti godoc
// @Summary create multipart upload
// @Tags file
// @Accept json
// @Produce json
// @Param pid path string true "profile id"
// @Param meta body d.FileForm true "file meta"
// @Success 200 {object} d.File
// @Failure 400 {object} errs.HTTPError
// @Router /file/large/{pid} [post]
func (s *service) createMulti(ctx *gin.Context, operator *operator.Identity) (int, interface{}, error) {
	var meta d.FileForm
	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		return 0, nil, err
	}

	file, err := s.db.File.Create().SetFilename(meta.Filename).SetMime(meta.Mime).SetSize(meta.Size).SetAccountID(operator.Account().ID).SetPath(meta.Category).Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	createOutput, err := s.storageService.CreateMultiPart(meta.Category, file.ID, file.Filename, meta.Mime)
	if err != nil {
		return 0, nil, err
	}

	output := CreateOutput{
		Key:      *createOutput.Key,
		UploadId: *createOutput.UploadId,
	}
	err = s.redisService.Create(ctx, file.ID.String(), output)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, d.FileFromEnt(file), nil
}

// uploadMulti godoc
// @Summary create multipart upload
// @Tags file
// @Accept application/octet-stream
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "file id"
// @Param part_id path int true "part id"
// @Param data body array true "file content"
// @Success 201
// @Failure 400 {object} errs.HTTPError
// @Router /file/large/{pid}/{id}/{part_id} [put]
func (s *service) putMulti(ctx *gin.Context, id string) (int, interface{}, error) {

	type FilePartUpload struct {
		PartID int `uri:"part_id" binding:"required"`
	}

	var ids FilePartUpload
	err := ctx.ShouldBindUri(&ids)
	if err != nil {
		return 0, nil, err
	}

	createObj, err := s.redisService.Get(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	var createOutput CreateOutput
	err = json.Unmarshal([]byte(createObj), &createOutput)
	if err != nil {
		return 0, nil, err
	}

	raw, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := bytes.NewReader(raw)
	part, err := s.storageService.PutMultiPart(reader, createOutput.Key, createOutput.UploadId, ids.PartID)
	if err != nil {
		return 0, nil, err
	}

	err = s.redisService.Part(ctx, id, part)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, nil, nil

}

type Finish struct {
	Parts        []*s3.CompletedPart `json:"parts" binding:"required"`
	CreateOutput CreateOutput        `json:"createOutput" binding:"required"`
}

// doneMulti godoc
// @Summary finish multipart upload
// @Tags file
// @Produce json
// @Param pid path string true "profile id"
// @Param id path string true "file id"
// @Success 201 {object} d.File
// @Failure 400 {object} errs.HTTPError
// @Router /file/large/{pid}/{id} [post]
func (s *service) doneMulti(ctx *gin.Context, id string) (int, interface{}, error) {
	var finish Finish
	partStringArray, err := s.redisService.GetParts(ctx, id)
	if err != nil {
		return 0, nil, err
	}
	parts := make([]*s3.CompletedPart, len(partStringArray)+1)
	if len(parts) == 0 {
		return 0, nil, errs.NoParts
	}
	max := 0
	for _, partString := range partStringArray {
		var partObj s3.CompletedPart
		err = json.Unmarshal([]byte(partString), &partObj)
		if err != nil {
			return 0, nil, err
		}
		partNumber := int(*partObj.PartNumber)
		if partNumber > max {
			max = partNumber
		}
		parts[partNumber] = &partObj
	}
	if len(parts) < max+1 {
		return 0, nil, errs.PartsExceeded
	}

	finish.Parts = parts[1 : max+1]

	var createOutput CreateOutput
	createObj, err := s.redisService.Get(ctx, id)
	err = json.Unmarshal([]byte(createObj), &createOutput)
	if err != nil {
		return 0, nil, err
	}
	finish.CreateOutput = createOutput

	_, err = s.storageService.CompleteMultiPart(finish.CreateOutput.Key, finish.CreateOutput.UploadId, finish.Parts)
	if err != nil {
		return 0, nil, err
	}

	file, err := s.db.File.UpdateOneID(uuid.MustParse(id)).SetStatus("uploaded").Save(ctx)
	if err != nil {
		return 0, nil, err
	}

	err = s.redisService.Finish(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	if strings.HasPrefix(file.Mime, "image/") {
		_, err = s.tasker.imageCompress(file)
		if err != nil {
			return 0, nil, err
		}
	}

	return http.StatusCreated, d.FileFromEnt(file), nil
}

func (s *service) owned(ctx *gin.Context, operator *operator.Identity, objID string) error {
	f, err := s.db.File.Query().WithAccount().Where(file.ID(uuid.MustParse(objID))).Only(ctx)
	if err != nil {
		return err
	}

	if f.Edges.Account.ID != operator.Account().ID {
		return errs.NotBelongsToOperator
	}
	return nil
}
