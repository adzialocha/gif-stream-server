package s3

import (
	"bytes"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

// Get a signed PUT request to upload files to S3.
func (s3 *S3) SignedPutObjectRequestURL(objectKey string, expiryTime time.Duration) (string) {
	req, _ := s3.Client.PutObjectRequest(&SDK.PutObjectInput{
		Bucket: aws.String(s3.BucketName),
		Key: aws.String(objectKey),
	})

	url, _ := req.Presign(expiryTime)

	return url
}

// Get a signed PUT request to upload files to S3.
func (s3 *S3) PutObject(objectKey string, buffer []byte) (*SDK.PutObjectOutput) {
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	res, _ := s3.Client.PutObject(&SDK.PutObjectInput{
		ACL: aws.String("public-read"),
		Body: fileBytes,
		Bucket: aws.String(s3.BucketName),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType: aws.String(fileType),
		Key: aws.String(objectKey),
	})

	return res
}
