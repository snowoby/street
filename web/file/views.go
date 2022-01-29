package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"street/ent"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

type Meta struct {
	Filename string `json:"filename"`
	Mime     string `json:"mime" binding:"required"`
	Size     int    `json:"size" binding:"required"`
	Category string `json:"category" binding:"required"`
}

type Part struct {
	Part int `uri:"part" binding:"required"`
}

type CreateResponse struct {
	UploadCredential CreateOutput `json:"uploadCredential"`
	File             ent.File     `json:"file"`
	Path             string       `json:"path"`
}

type CreateOutput struct {
	Key      string `json:"key" binding:"required"`
	UploadId string `json:"upload_id" binding:"required"`
}

func create(ctx *gin.Context, store *data.Store, visitor *controller.Identity) (int, interface{}, error) {
	var meta Meta
	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		return 0, nil, err
	}

	file, err := store.File().Create(ctx, meta.Filename, meta.Category, meta.Mime, meta.Size, visitor.Profile().ID)
	if err != nil {
		return 0, nil, err
	}

	createOutput, err := store.Storage.CreateMultiPart(meta.Category, file.ID, file.Filename, meta.Mime)
	if err != nil {
		return 0, nil, err
	}
	output := CreateOutput{
		Key:      *createOutput.Key,
		UploadId: *createOutput.UploadId,
	}
	err = store.FileRedis.Create(ctx, file.ID.String(), output)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, file, nil
}

func upload(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	//type FilePartUpload struct {
	//	Key        string `json:"key" binding:"required"`
	//	UploadID   string `json:"upload_id" binding:"required"`
	//	PartNumber int    `json:"part_number" binding:"required"`
	//	Content    string `json:"content" binding:"required"`
	//}

	type FilePartUpload struct {
		UploadID string `uri:"id" binding:"required"`
		PartID   int    `uri:"part_id" binding:"required"`
	}

	var ids FilePartUpload
	err := ctx.ShouldBindUri(&ids)
	if err != nil {
		return 0, nil, err
	}

	createObj, err := store.FileRedis.Get(ctx, ids.UploadID)
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
	part, err := store.Storage.PutMultiPart(reader, createOutput.Key, ids.UploadID, ids.PartID)
	if err != nil {
		return 0, nil, err
	}

	err = store.FileRedis.Part(ctx, ids.UploadID, ids.PartID, part)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusOK, part, nil

}

type Finish struct {
	Parts        []*s3.CompletedPart `json:"parts" binding:"required"`
	CreateOutput CreateOutput        `json:"createOutput" binding:"required"`
}

func done(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	//TODO
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	var finish Finish
	err := ctx.ShouldBindJSON(&finish)
	if err != nil {
		return 0, nil, err
	}

	completeOutput, err := store.Storage.CompleteMultiPart(finish.CreateOutput.Key, finish.CreateOutput.UploadId, finish.Parts)
	fmt.Println(completeOutput)
	if err != nil {
		return 0, nil, err
	}

	file, err := store.File().UpdateStatus(ctx, id, "uploaded")
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, file, nil
}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ok, err := store.File().IsOwner(ctx, operator.Profile().ID, objectID)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NotBelongsToOperator
	}
	return nil
}
