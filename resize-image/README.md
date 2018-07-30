## A typical function -resize the image files (thumbnails)

on s3, assume we have two buckets with following names:
testbucket

testbucket-thumbnail

when we upload an image to s3 bucket 'testbucket', event is triggered to go-lambda function,
and it will resize the image and write the new thumbnail into another bucket 'testbucket-thumbnail'.

To make it sample code simple, we only support 'jpeg' format in the code.

## Build linux application
As aws support to upload go binary application directly, but the application has to be 'linux' compiled.
So when we build applicaiton, we should build like following,

```
GOOS=linux go build -o your_lambda_function_name main.go
zip deployment.zip your_lambda_function_name 
```

