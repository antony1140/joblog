package data

import (
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"log"
	"context"
	"io"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"

)

func InitS3() ( *s3.Client){
	envErr := godotenv.Load("./data/.env")

	if envErr != nil {
		log.Fatal(envErr)
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("InitS3() Didnt work :/")
		log.Fatal(err)
	}
	client := s3.NewFromConfig(cfg)
	return client
}

func UploadS3(client *s3.Client, file io.Reader){
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("jobcontracts"),
		Key:    aws.String("newfile"),
		Body:   file,
	})
	if err != nil {
		fmt.Println("UploadS3() Didnt work :/")
		log.Fatal(err)
	}
	fmt.Println("File upload success: ",  result.Location)


}
