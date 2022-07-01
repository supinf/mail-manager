package apigw

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/supinf/supinf-mail/app-api/aws"
)

// APIGateway Amazon API Gateway
type APIGateway struct {
	client *apigateway.APIGateway
}

// New 新しい APIGateway Client を生成します
func New() (*APIGateway, error) {
	client, err := apiGatewayClient()
	if err != nil {
		return nil, err
	}
	return &APIGateway{
		client: client,
	}, nil
}

// Client Client を返します
func (ag *APIGateway) Client() *apigateway.APIGateway {
	return ag.client
}

func apiGatewayClient() (*apigateway.APIGateway, error) {
	sess, _, err := aws.Configure(nil)
	if err != nil {
		return nil, err
	}
	client := apigateway.New(sess)
	if client == nil {
		return nil, errors.New("unable create Amazon API Gateway")
	}
	return client, nil
}
