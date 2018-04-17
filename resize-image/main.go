package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"image/jpeg"
	// "os"
)

func Handler(ctx context.Context, s3Event events.S3Event) {

	// sess, _ := session.NewSession(&aws.Config{
	// 	Region: aws.String("ap-southeast-2")},
	// )

	svc := s3.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-2"))

	for _, record := range s3Event.Records {
		s := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s.Bucket.Name, s.Object.Key)

		input := &s3.GetObjectInput{}
		input.Bucket = aws.String(s.Bucket.Name)
		input.Key = aws.String(s.Object.Key)

		// downloader := s3manager.NewDownloader(sess)

		// file, err := os.Create("/tmp/tt")
		// if err != nil {
		// 	fmt.Printf("Unable to open file %s", err.Error())
		// }
		// defer file.Close()

		// _, err = downloader.Download(file,
		// 	&s3.GetObjectInput{
		// 		Bucket: aws.String(s.Bucket.Name),
		// 		Key:    aws.String(s.Object.Key),
		// 	})

		// if err != nil {
		// 	fmt.Printf("Unable to download item %q, %v", s.Object.Key, err.Error())
		// 	continue
		// }

		output, err := svc.GetObject(input)

		if err != nil {
			fmt.Printf("failed to get object:%s\n", err.Error())
			continue
		}

		fmt.Println("start to get the image successfully!")

		img, err := jpeg.Decode(output.Body)

		fmt.Println("decode the image successfully!")

		if err != nil {
			fmt.Printf("failed to read object into image object:%s\n", err.Error())
			continue
		}

		m := resize.Thumbnail(200, 100, img, resize.Bilinear)

		var buffer bytes.Buffer
		err = jpeg.Encode(&buffer, m, nil)

		if err != nil {
			fmt.Printf("failed to write thumbnail into buffer:%s\n", err.Error())
			continue
		}

		r := bytes.NewReader(buffer.Bytes())
		newObj := &s3.PutObjectInput{}
		newObj.Bucket = aws.String(s.Bucket.Name + "-thumbnail")
		newKey := s.Object.Key + "-thumbnail.jpg"
		newObj.Key = aws.String(newKey)
		newObj.Body = r

		_, err = svc.PutObject(newObj)

		fmt.Println("write the image done!")

		if err != nil {
			fmt.Printf("failed to wirte new resized image into bucket:%s\n", err.Error())
		}
	}

	return
}

func main() {
	lambda.Start(Handler)
}
