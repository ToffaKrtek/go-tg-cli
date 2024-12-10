package s3

import "github.com/ToffaKrtek/go-tg-cli/environment"

type S3File struct {
	Url             string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	FilePath        string
	ObjectName      string
}

func NewFile(filePath string, opts ...s3Function) *S3File {
	file := &S3File{
		Url:             environment.Get("S3_URL"),
		AccessKeyID:     environment.Get("S3_ACCESS_KEY"),
		SecretAccessKey: environment.Get("S3_SECRET_KEY"),
		Bucket:          environment.Get("S3_BUCKET"),
		FilePath:        filePath,
	}
	for _, opt := range opts {
		opt(file)
	}
	return file
}

type s3Function func(*S3File)

func Url(url string) s3Function {
	return func(file *S3File) {
		file.Url = url
	}
}

func AccessKeyID(accessKeyID string) s3Function {
	return func(file *S3File) {
		file.AccessKeyID = accessKeyID
	}
}

func SecretAccessKey(secretAccessKey string) s3Function {
	return func(file *S3File) {
		file.SecretAccessKey = secretAccessKey
	}
}

func Bucket(bucket string) s3Function {
	return func(file *S3File) {
		file.Bucket = bucket
	}
}

func ObjectName(objectName string) s3Function {
	return func(file *S3File) {
		file.ObjectName = objectName
	}
}
