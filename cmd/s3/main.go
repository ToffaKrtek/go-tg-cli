package main

import (
	"flag"
	"fmt"

	"github.com/ToffaKrtek/go-tg-cli/environment"
	"github.com/ToffaKrtek/go-tg-cli/s3"
)

var envPath = ""

func main() {
	environment.Parse(envPath)
	flagBucket := flag.String("bucket", environment.Get("S3_BUCKET"), "Bucket name")
	flagUrl := flag.String("url", environment.Get("S3_URL"), "S3 url")
	flagAccessKey := flag.String("access", environment.Get("S3_ACCESS_KEY"), "S3 access key")
	flagSecretKey := flag.String("secret", environment.Get("S3_SECRET_KEY"), "S3 secret key")
	flagFilePath := flag.String("file", "", "File path to upload")
	flagObjectName := flag.String("object", "", "Object s3 name")
	flag.Parse()

	if len(*flagObjectName) == 0 {
		flagObjectName = flagFilePath
	}

	file := s3.NewFile(*flagFilePath,
		s3.Url(*flagUrl),
		s3.AccessKeyID(*flagAccessKey),
		s3.SecretAccessKey(*flagSecretKey),
		s3.Bucket(*flagBucket),
		s3.ObjectName(*flagObjectName),
	)
	url, err := file.Upload()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File uploaded successfully to: ", url)
}
