package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type dynamoClientConfig struct {
	Endpoint string
}

func NewDynamoConfig() aws.Config {
	c, _ := config.LoadDefaultConfig(context.Background())
	return c
}
