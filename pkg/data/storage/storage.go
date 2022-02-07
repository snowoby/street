package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"os"
	"strings"
)

type FilePart struct {
	ETag       string
	PartNumber int
}

type Storage struct {
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	bucketName string
}

func NewDefaultConfig() *aws.Config {
	return &aws.Config{
		Credentials:      credentials.NewStaticCredentials(os.Getenv("s3_accesskey"), os.Getenv("s3_secretkey"), ""),
		Endpoint:         aws.String(os.Getenv("storage_access_endpoint")),
		Region:           aws.String(os.Getenv("s3_region")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
}

func New(s3Config *aws.Config) *Storage {

	newSession, err := session.NewSession(s3Config)

	if err != nil {
		panic(err)
	}

	client := s3.New(newSession)

	//_, err = client.CreateBucket(&s3.CreateBucketInput{
	//	Bucket: aws.String("develop"),
	//})
	if err != nil {
		panic(err)
	}
	return &Storage{
		client:     client,
		uploader:   s3manager.NewUploader(newSession),
		downloader: s3manager.NewDownloader(newSession),
		bucketName: "develop",
	}

}

func (s *Storage) Delete(path string, id uuid.UUID, fileIdentifier string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(strings.Join([]string{path, id.String(), fileIdentifier}, "/")),
	})
	return err
}

func (s *Storage) PutSingle(stream *bytes.Reader, path string, id uuid.UUID, filename string, fileIdentifier string, mime string) (*s3manager.UploadOutput, error) {
	return s.uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(s.bucketName),
		Key:                aws.String(strings.Join([]string{path, id.String(), fileIdentifier}, "/")),
		ContentType:        aws.String(mime),
		Body:               stream,
		ContentDisposition: aws.String(fmt.Sprintf("filename=%s", filename)),
	})

}

func (s *Storage) Get(path string, id string, fileIdentifier string) ([]byte, error) {
	buff := &aws.WriteAtBuffer{}
	_, err := s.downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(strings.Join([]string{path, id, fileIdentifier}, "/")),
	})

	if err != nil {
		return nil, err
	}
	data := buff.Bytes()
	return data, nil

}

func (s *Storage) CreateMultiPart(path string, id uuid.UUID, filename string, mime string) (*s3.CreateMultipartUploadOutput, error) {
	input := &s3.CreateMultipartUploadInput{
		Bucket:             aws.String(s.bucketName),
		Key:                aws.String(strings.Join([]string{path, id.String(), "original"}, "/")),
		ContentType:        aws.String(mime),
		ContentDisposition: aws.String(fmt.Sprintf("filename =%s", filename)),
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
	// TODO sort
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
