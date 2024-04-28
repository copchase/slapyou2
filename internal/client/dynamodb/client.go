package dynamodb

import awsddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type dynamoClient struct {
	ddb *awsddb.Client
}

func NewDynamoClient() *dynamoClient {
	c := NewDynamoConfig()

	ddb := awsddb.NewFromConfig(c)

	return &dynamoClient{
		ddb: ddb,
	}
}
