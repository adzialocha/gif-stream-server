package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	Client *SDK.S3
	BucketName string
	Region string
}

func New(accessKeyId string, secretAccessKey string, region string, bucketName string) *S3 {
	// Prepare S3 configuration.
	creds := credentials.NewStaticCredentials(accessKeyId, secretAccessKey, "")
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	// Return new API instance.
	return &S3{
		Client: SDK.New(session.New(), cfg),
		BucketName: bucketName,
		Region: region,
	}
}

func (s3 *S3) KeyToUrl(key string) string {
	return "https://s3." + s3.Region + ".amazonaws.com/" + s3.BucketName + "/" + key
}
