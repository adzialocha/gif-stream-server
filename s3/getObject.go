package s3

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	SDK "github.com/aws/aws-sdk-go/service/s3"
)

// Get the contents of an object.
func (s3 *S3) GetObjectBytes(objectKey string) ([]byte, error) {
	results, err := s3.Client.GetObject(&SDK.GetObjectInput{
		Bucket: aws.String(s3.BucketName),
		Key: aws.String(objectKey),
	})

	if err != nil {
		return nil, err
	}

	defer results.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, results.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
