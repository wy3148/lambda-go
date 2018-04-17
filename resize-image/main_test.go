package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func TestHandler(t *testing.T) {

	c := context.Background()
	e := events.S3Event{}
	e.Records = make([]events.S3EventRecord, 1)
	e.Records[0].S3.Bucket.Name = "lambda-deployment-3148"
	e.Records[0].S3.Object.Key = "x112345.jpg"
	Handler(c, e)
}
