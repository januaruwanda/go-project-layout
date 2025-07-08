package uploader

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type Uploader struct {
	storageAddress string
	bucketName     string
	accessKey      string
	secretKey      string
	secure         bool
	minioClient    *minio.Client
}

func NewUploaderWithConfig() (*Uploader, error) {
	storageAddress := viper.GetString("minio.storage_address")
	bucketName := viper.GetString("minio.bucket_name")
	accessKey := viper.GetString("minio.access_key")
	secretKey := viper.GetString("minio.secret_key")
	secure := viper.GetBool("minio.secure")

	minioClient, err := minio.New(storageAddress, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		return nil, err
	}

	return &Uploader{
		storageAddress: storageAddress,
		bucketName:     bucketName,
		accessKey:      accessKey,
		secretKey:      secretKey,
		secure:         secure,
		minioClient:    minioClient,
	}, nil
}

func (u *Uploader) UploadToMinio(objectName string, file multipart.File, fileSize int64, contentType string) error {
	_, err := u.minioClient.PutObject(
		context.Background(),
		u.bucketName,
		objectName,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *Uploader) GeneratePresignedURL(objectName string, expiry time.Duration) (string, error) {
	presignedURL, err := u.minioClient.PresignedGetObject(context.Background(), u.bucketName, objectName, expiry, nil)
	if err != nil {
		return "", err
	}

	protocol := "http://"
	if u.secure {
		protocol = "https://"
	}
	return protocol + presignedURL.Host + presignedURL.Path + "?" + presignedURL.RawQuery, nil
}

type UploaderHandlerFunc func(objectName string, file multipart.File, fileSize int64, contentType string) error

func CreateUploaderHandler() (UploaderHandlerFunc, error) {
	uploader, err := NewUploaderWithConfig()
	if err != nil {
		return nil, err
	}

	return func(objectName string, file multipart.File, fileSize int64, contentType string) error {
		return uploader.UploadToMinio(objectName, file, fileSize, contentType)
	}, nil
}

type PresignedURLHandlerFunc func(objectName string) (string, error)

func CreatePresignedURLHandler() (PresignedURLHandlerFunc, error) {
	uploader, err := NewUploaderWithConfig()
	if err != nil {
		return nil, err
	}

	return func(objectName string) (string, error) {
		return uploader.GeneratePresignedURL(objectName, 5*time.Minute)
	}, nil
}
