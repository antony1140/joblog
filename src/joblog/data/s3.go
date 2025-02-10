package data

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	// "github.com/joho/godotenv"
)

type Presigner struct {
	PresignClient *s3.PresignClient
}


func InitS3() ( *s3.Client){
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		fmt.Println("InitS3() Didnt work :/")
		log.Print(err)
	}
	client := s3.NewFromConfig(cfg)
	return client
}

func InitS3PresignClient(client *s3.Client)(Presigner){
	return Presigner {
s3.NewPresignClient(client),

	}
}

func UploadS3(client *s3.Client, file io.Reader, fileKey string)(error){
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(fileKey),
		Body:   file,
	})
	if err != nil {
		fmt.Println("UploadS3() Didnt work :/")
		return err
	}
	fmt.Println("File upload success: ",  result.Location)
	return nil


}


func DownloadS3(client *s3.Client, fileName string)(error){
	downloader := manager.NewDownloader(client)
	newFile, err := os.Create("./assets/"+ fileName )
	if err != nil {
		log.Println(err)
	}
	defer newFile.Close()
	// buf := make([]byte, int(headObject.ContentLength))
	// wrap with aws.WriteAtBuffer
	// w := manager.NewWriteAtBuffer(buf)
	numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")), 
		Key:    aws.String(fileName),
	})
	fmt.Println(numBytes, "bytes downloaded")
	if err != nil {
		return err
	}

	return nil
}

func DownloadS3WithKey(client *s3.Client, fileName string, key string)(error){
	downloader := manager.NewDownloader(client)
	newFile, err := os.Create("./assets/"+ fileName )
	if err != nil {
		log.Println(err)
	}
	defer newFile.Close()
	// buf := make([]byte, int(headObject.ContentLength))
	// wrap with aws.WriteAtBuffer
	// w := manager.NewWriteAtBuffer(buf)
	numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")), 
		Key:    aws.String(key),
	})
	fmt.Println(numBytes, "bytes downloaded")
	if err != nil {
		return err
	}

	return nil
}

func GetObject(presigner Presigner,
	ctx context.Context, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
		contentType := "application/pdf"
		bucketName := os.Getenv("AWS_S3_BUCKET")
	request, err := presigner.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		ResponseContentType: &contentType,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Hour))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request, err
}




func DeleteS3(fileKey string, client *s3.Client, ctx context.Context) (error) {
	conf := s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv(os.Getenv("AWS_S3_BUCKET"))),
		Key: aws.String(fileKey),
	}	
	_, err := client.DeleteObject(ctx, &conf)

	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}



