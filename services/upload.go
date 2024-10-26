package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadToSpaces -
func UploadToSpaces(folder, path, id string) string {
	// https://jto.nyc3.digitaloceanspaces.com
	// The session the S3 Uploader will use
	endpoint := "fra1.digitaloceanspaces.com"
	region := "fra1"
	key := "DO00JLPUN74PAYPU8W3P"                           // Access key pair. You can create access key pairs using the control panel or API.
	secret := "tl880N/cOTqML2HwaOB8rk6+Y2FrS+6Mvub/pV5tthk" //os.Getenv("SPACES_SECRET") // Secret access key defined through an environment variable.

	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    &endpoint,
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", path, err)
	}

	myBucket := "prelo"
	myString := folder + "/" + id + ".jpg"
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myString),
		Body:   f,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	errd := os.Remove(path)
	if errd != nil {
		fmt.Printf("failed to open file %q, %v", path, errd)
	}
	return aws.StringValue(&result.Location)
}

/*

  service: S3
  endpoint: https://prelo.fra1.digitaloceanspaces.com
  access_key_id: DO00JLPUN74PAYPU8W3P
  secret_access_key: tl880N/cOTqML2HwaOB8rk6+Y2FrS+6Mvub/pV5tthk
  region: fra1
  bucket: prelo.fra1.digitaloceanspaces.com

*/
