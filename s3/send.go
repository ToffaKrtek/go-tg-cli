package s3

import (
	"context"
	"time"

	minio "github.com/minio/minio-go/v7"
	credentials "github.com/minio/minio-go/v7/pkg/credentials"
)

func (f S3File) Upload() (string, error) {
	ctx := context.Background()
	client, err := f.getClient()
	if err != nil {
		return "", err
	}
	_, err = client.FPutObject(
		ctx,
		f.Bucket,
		f.ObjectName,
		f.FilePath,
		minio.PutObjectOptions{ContentType: "application/octet-stream"},
	)
	presignedURL, err := client.PresignedGetObject(ctx, f.Bucket, f.ObjectName, 24*time.Hour, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), err
}

// func (f S3File) Download() error {
// 	return nil
// }

func (f S3File) getClient() (*minio.Client, error) {
	return minio.New(f.Url, &minio.Options{
		Creds:  credentials.NewStaticV4(f.AccessKeyID, f.SecretAccessKey, ""),
		Secure: true,
	})
}
