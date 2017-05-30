package s3

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

func (s3 *S3) SignedPutObjectRequestURL(objectKey string, expiryTime time.Duration) (string) {
	// Get a signed PUT request to upload files to S3.
	req, _ := s3.Client.PutObjectRequest(&SDK.PutObjectInput{
		Bucket: aws.String(s3.BucketName),
		Key: aws.String(objectKey),
	})
	url, _ := req.Presign(expiryTime)
	return url
}
