package base3

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"strings"
)

type prototype struct {
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	bucketName string
}

type Prototype interface {
	Get(path string, id string, fileIdentifier string) ([]byte, error)
	Delete(path string, id uuid.UUID, fileIdentifier string) error
	PutSingle(stream *bytes.Reader, path string, id uuid.UUID, filename string, fileIdentifier string, mime string) (*s3manager.UploadOutput, error)

	Client() *s3.S3
	Uploader() *s3manager.Uploader
	Downloader() *s3manager.Downloader
	BucketName() string
}

func (p *prototype) Client() *s3.S3 {
	return p.client
}

func (p *prototype) Uploader() *s3manager.Uploader {
	return p.uploader
}

func (p *prototype) Downloader() *s3manager.Downloader {
	return p.downloader
}

func (p *prototype) BucketName() string {
	return p.bucketName
}

func (p *prototype) Get(path string, id string, fileIdentifier string) ([]byte, error) {
	buff := &aws.WriteAtBuffer{}
	_, err := p.downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(strings.Join([]string{path, id, fileIdentifier}, "/")),
	})

	if err != nil {
		return nil, err
	}
	data := buff.Bytes()
	return data, nil

}

func (p *prototype) Delete(path string, id uuid.UUID, fileIdentifier string) error {
	_, err := p.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(strings.Join([]string{path, id.String(), fileIdentifier}, "/")),
	})
	return err
}

func (p *prototype) PutSingle(stream *bytes.Reader, path string, id uuid.UUID, filename string, fileIdentifier string, mime string) (*s3manager.UploadOutput, error) {
	return p.uploader.Upload(&s3manager.UploadInput{
		Bucket:             aws.String(p.bucketName),
		Key:                aws.String(strings.Join([]string{path, id.String(), fileIdentifier}, "/")),
		ContentType:        aws.String(mime),
		Body:               stream,
		ContentDisposition: aws.String(fmt.Sprintf("filename=%s", filename)),
	})

}

func New(s3Config *aws.Config) *prototype {

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
	return &prototype{
		client:     client,
		uploader:   s3manager.NewUploader(newSession),
		downloader: s3manager.NewDownloader(newSession),
		bucketName: "develop",
	}

}
