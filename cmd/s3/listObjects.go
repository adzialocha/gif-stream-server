package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

// Get a paginated list of objects.
func (s3 *S3) ListObjects(prefix string, pageSize int64, marker string) (*SDK.ListObjectsOutput) {
	options := &SDK.ListObjectsInput{
		Bucket: aws.String(s3.BucketName),
		Prefix: aws.String(prefix),
		Delimiter: aws.String("/"),
		MaxKeys: aws.Int64(pageSize),
	}

	// Set marker option for pagination when given.
	if (marker != "") {
		options.SetMarker(marker)
	}

	// Do request to S3 instance.
	out, _ := s3.Client.ListObjects(options)
	return out
}

// Get all objects.
func (s3 *S3) ListAllObjects(prefix string) ([]string) {
	options := &SDK.ListObjectsInput{
		Bucket: aws.String(s3.BucketName),
		Prefix: aws.String(prefix),
		Delimiter: aws.String("/"),
	}

	objects := []string{}

	// Do request to S3 instance.
	s3.Client.ListObjectsPages(
		options,
		func(p *SDK.ListObjectsOutput, last bool) (shouldContinue bool) {
			for _, obj := range p.Contents {
				objects = append(objects, aws.StringValue(obj.Key))
			}
			return true
		},
	)

	return objects
}
