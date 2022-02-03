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
	"strings"
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

func createSingle(ctx *gin.Context, store *data.Store, visitor *controller.Identity) (int, interface{}, error) {
	var meta Meta
	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		return 0, nil, err
	}

	file, err := store.DB.File.Create(ctx, meta.Filename, meta.Category, meta.Mime, meta.Size, visitor.Profile().ID)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, file, nil
}

func putSingle(ctx *gin.Context, store *data.Store, visitor *controller.Identity) (int, interface{}, error) {
	id := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)

	file, err := store.DB.File.Get(ctx, id)
	if err != nil {
		return 0, nil, err
	}

	if file.Edges.Profile.ID != visitor.Profile().ID {
		return 0, nil, errs.NotFoundError
	}

	if file.Status != "created" {
		return 0, nil, errs.FileUploadedError
	}

	raw, _ := ioutil.ReadAll(ctx.Request.Body)
	reader := bytes.NewReader(raw)
	_, err = store.Storage.PutSingle(reader, file.Path, id, file.Filename, "original", file.Mime)
	if err != nil {
		return 0, nil, err
	}

	file, err = store.DB.File.UpdateStatus(ctx, id, "uploaded")
	if err != nil {
		return 0, nil, err
	}

	if strings.HasPrefix(file.Mime, "image/") {
		info, err := store.Task.ImageCompress(file)
		if err != nil {
			return 0, nil, err
		}
		fmt.Println(info)
	}

	return http.StatusOK, file, nil

}

func createMulti(ctx *gin.Context, store *data.Store, visitor *controller.Identity) (int, interface{}, error) {
	var meta Meta
	err := ctx.ShouldBindJSON(&meta)
	if err != nil {
		return 0, nil, err
	}

	file, err := store.DB.File.Create(ctx, meta.Filename, meta.Category, meta.Mime, meta.Size, visitor.Profile().ID)
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
	err = store.MultiPartRedis.Create(ctx, file.ID.String(), output)
	if err != nil {
		return 0, nil, err
	}

	return http.StatusCreated, file, nil
}

func upload(ctx *gin.Context, store *data.Store) (int, interface{}, error) {

	type FilePartUpload struct {
		UploadID string `uri:"id" binding:"required"`
		PartID   int    `uri:"part_id" binding:"required"`
	}

	var ids FilePartUpload
	err := ctx.ShouldBindUri(&ids)
	if err != nil {
		return 0, nil, err
	}

	createObj, err := store.MultiPartRedis.Get(ctx, ids.UploadID)
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
	part, err := store.Storage.PutMultiPart(reader, createOutput.Key, createOutput.UploadId, ids.PartID)
	if err != nil {
		return 0, nil, err
	}

	err = store.MultiPartRedis.Part(ctx, ids.UploadID, part)
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
	partStringArray, err := store.MultiPartRedis.GetParts(ctx, id.String())
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
	createObj, err := store.MultiPartRedis.Get(ctx, id.String())
	err = json.Unmarshal([]byte(createObj), &createOutput)
	if err != nil {
		return 0, nil, err
	}
	finish.CreateOutput = createOutput

	completeOutput, err := store.Storage.CompleteMultiPart(finish.CreateOutput.Key, finish.CreateOutput.UploadId, finish.Parts)
	fmt.Println(completeOutput)
	if err != nil {
		return 0, nil, err
	}

	file, err := store.DB.File.UpdateStatus(ctx, id, "uploaded")
	if err != nil {
		return 0, nil, err
	}

	err = store.MultiPartRedis.Finish(ctx, id.String())
	if err != nil {
		return 0, nil, err
	}

	if strings.HasPrefix(file.Mime, "image/") {
		info, err := store.Task.ImageCompress(file)
		if err != nil {
			return 0, nil, err
		}
		fmt.Println(info)
	}

	return http.StatusCreated, file, nil
}

func owned(ctx *gin.Context, store *data.Store, operator *controller.Identity) error {
	objectID := ctx.MustGet(value.StringObjectUUID).(uuid.UUID)
	ok, err := store.DB.File.IsOwner(ctx, operator.Profile().ID, objectID)
	if err != nil {
		return err
	}
	if !ok {
		return errs.NotBelongsToOperator
	}
	return nil
}
