package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"street/ent"
	"street/errs"
	"street/pkg/controller"
	"street/pkg/data"
	"street/pkg/data/value"
)

type FileMeta struct {
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
	var meta FileMeta
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

	return http.StatusCreated, CreateResponse{
		UploadCredential: output,
		File:             *file,
		Path:             filepath.Join(meta.Category, file.ID.String(), file.Filename),
	}, nil
}

func upload(ctx *gin.Context, store *data.Store) (int, interface{}, error) {
	type FilePartUpload struct {
		Bucket     string `json:"bucket" binding:"required"`
		Key        string `json:"key" binding:"required"`
		UploadID   string `json:"upload_id" binding:"required"`
		PartNumber int    `json:"part_number" binding:"required"`
		Content    string `json:"content" binding:"required"`
	}

	var chunk FilePartUpload
	err := ctx.ShouldBindUri(&chunk)
	if err != nil {
		return 0, nil, err
	}

	raw, err := base64.StdEncoding.DecodeString(chunk.Content)
	if err != nil {
		return 0, nil, err
	}

	reader := bytes.NewReader(raw)
	part, err := store.Storage.PutMultiPart(reader, chunk.Key, chunk.UploadID, chunk.PartNumber)
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
