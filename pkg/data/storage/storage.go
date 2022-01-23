package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"path/filepath"
)

type FilePart struct {
	ETag       string
	PartNumber int
}

type Storage struct {
	client     *s3.S3
	bucketName string
}

func New() *Storage {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials("minioadmin", "minioadmin", ""),
		Endpoint:         aws.String("http://localhost:9000"),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		panic(err)
	}

	return &Storage{
		client:     s3.New(newSession),
		bucketName: "develop",
	}

}

func (s *Storage) CreateMultiPart(path string, id uuid.UUID, filename string, mime string) (*s3.CreateMultipartUploadOutput, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(filepath.Join(path, id.String(), filename)),
		ContentType: aws.String(mime),
	}

	return s.client.CreateMultipartUpload(input)
}

func (s *Storage) PutMultiPart(stream *bytes.Reader, key string, uploadID string, partNumber int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          stream,
		Bucket:        aws.String(s.bucketName),
		Key:           aws.String(key),
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      aws.String(uploadID),
		ContentLength: aws.Int64(stream.Size()),
	}
	maxRetries := 3
	for tryNum <= maxRetries {
		uploadResult, err := s.client.UploadPart(partInput)
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

func (s *Storage) CompleteMultiPart(key string, uploadID string, fileParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	//var completedParts = make([]*s3.CompletedPart, len(fileParts))
	//for i, part := range fileParts {
	//	completedParts[i] = &s3.CompletedPart{
	//		ETag:       aws.String(part.ETag),
	//		PartNumber: aws.Int64(int64(part.PartNumber)),
	//	}
	//}

	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(s.bucketName),
		Key:      aws.String(key),
		UploadId: aws.String(uploadID),
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: fileParts,
		},
	}
	return s.client.CompleteMultipartUpload(completeInput)
}

func (s *Storage) AbortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	fmt.Println("Aborting multipart upload for UploadId#" + *resp.UploadId)
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}
