package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

// Get the contents of an object.
func (s3 *S3) DeleteObjects(keys []string) (*SDK.DeleteObjectsOutput) {
	objects := []*SDK.ObjectIdentifier{}

	for _, key := range keys {
		objects = append(objects, &SDK.ObjectIdentifier{
			Key: aws.String(key),
		})
	}

	deleteInput := &SDK.Delete{
		Objects: objects,
		Quiet: aws.Bool(true),
	}

	result, _ := s3.Client.DeleteObjects(&SDK.DeleteObjectsInput{
		Bucket: aws.String(s3.BucketName),
		Delete: deleteInput,
	})

	return result
}
