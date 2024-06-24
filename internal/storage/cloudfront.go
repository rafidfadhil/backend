package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

const CloudFrontURLBase = "https://d3lwcb2m2mg8wp.cloudfront.net/"

func InvalidateImage(imageKey string) (err error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewSharedCredentials("./configs/aws/credentials", "default"),
		Region:      aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}

	svc := cloudfront.New(sess)

	// Specify the paths that you want to invalidate.
	paths := []*string{
		aws.String("/" + imageKey),
	}

	input := &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String("E3DR5W32TWEHS4"),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(imageKey),
			Paths: &cloudfront.Paths{
				Quantity: aws.Int64(int64(len(paths))),
				Items:    paths,
			},
		},
	}

	_, err = svc.CreateInvalidation(input)
	if err != nil {
		return err
	}

	return nil
}
