package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"street/pkg/base3"
	"strings"
)

type storageService struct {
	base3.Prototype
}

//func NewDefaultConfig() *aws.Config {
//	return &aws.Config{
//		Credentials:      credentials.NewStaticCredentials(os.Getenv("s3_accesskey"), os.Getenv("s3_secretkey"), ""),
//		Endpoint:         aws.String(os.Getenv("storage_access_endpoint")),
//		Region:           aws.String(os.Getenv("s3_region")),
//		DisableSSL:       aws.Bool(true),
//		S3ForcePathStyle: aws.Bool(true),
//	}
//}

func newS3(s3Config *aws.Config) *storageService {
	return &storageService{Prototype: base3.New(s3Config)}
}

func (s *storageService) CreateMultiPart(path string, id uuid.UUID, filename string, mime string) (*s3.CreateMultipartUploadOutput, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket:             aws.String(s.BucketName()),
		Key:                aws.String(strings.Join([]string{path, id.String(), "original"}, "/")),
		ContentType:        aws.String(mime),
		ContentDisposition: aws.String(fmt.Sprintf("filename=%s", filename)),
	}

	return s.Client().CreateMultipartUpload(input)
}

func (s *storageService) PutMultiPart(stream *bytes.Reader, key string, uploadID string, partNumber int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          stream,
		Bucket:        aws.String(s.BucketName()),
		Key:           aws.String(key),
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      aws.String(uploadID),
		ContentLength: aws.Int64(stream.Size()),
	}
	maxRetries := 3
	for tryNum <= maxRetries {
		uploadResult, err := s.Client().UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aErr, ok := err.(awserr.Error); ok {
					return nil, aErr
				}
				return nil, err
			}
			fmt.Printf("Retrying to upload part #%v\n", partNumber)
			tryNum++
		} else {
			fmt.Printf("Uploaded part #%v\n", partNumber)
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func (s *storageService) CompleteMultiPart(key string, uploadID string, fileParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	//var completedParts = make([]*s3.CompletedPart, len(fileParts))
	//for i, part := range fileParts {
	//	completedParts[i] = &s3.CompletedPart{
	//		ETag:       aws.String(part.ETag),
	//		PartNumber: aws.Int64(int64(part.PartNumber)),
	//	}
	//}
	// TODO sort
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(s.BucketName()),
		Key:      aws.String(key),
		UploadId: aws.String(uploadID),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: fileParts,
		},
	}
	return s.Client().CompleteMultipartUpload(completeInput)
}

func (s *storageService) AbortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	fmt.Println("Aborting multipart upload for UploadId#" + *resp.UploadId)
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}
