package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func AwsConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedCredentialsFiles(
			[]string{"configs/aws/credentials"},
		),
		config.WithSharedConfigFiles(
			[]string{"configs/aws/config"},
		),
	)
	if err != nil {
		return aws.Config{}, err
	}

	return cfg, nil
}
