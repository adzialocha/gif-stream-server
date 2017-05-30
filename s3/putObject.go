package s3

import (
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
