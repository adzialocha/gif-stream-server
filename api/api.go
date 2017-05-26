package api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type API struct {
	s3Client *s3.S3
	s3BucketName string
	s3Region string
}

func New(accessKeyId string, secretAccessKey string, sessionToken string, region string, bucketName string) *API {
	// Prepare S3 configuration.
	creds := credentials.NewStaticCredentials(accessKeyId, secretAccessKey, sessionToken)
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	// Return new API instance.
	return &API{
		s3Client: s3.New(session.New(), cfg),
		s3BucketName: bucketName,
		s3Region: region,
	}
}
