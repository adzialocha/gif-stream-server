package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

func (s3 *S3) ListObjects(prefix string, pageSize int64, marker string) (*SDK.ListObjectsOutput) {
	options := &SDK.ListObjectsInput{
		Bucket: aws.String(s3.BucketName),
		Prefix: aws.String(prefix),
		MaxKeys: aws.Int64(pageSize),
	}

	// Set marker option for pagination when given.
	if (marker != "") {
		options.Marker = aws.String(marker)
	}

	// Do request to S3 instance.
	out, _ := s3.Client.ListObjects(options)
	return out
}
