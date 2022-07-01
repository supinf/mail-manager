package ses

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/supinf/supinf-mail/app-api/aws"
)

// SES Amazon SES
type SES struct {
	client *sesv2.SESV2
}

// New 新しい SES Client を生成します
func New() (*SES, error) {
	client, err := sesClient()
	if err != nil {
		return nil, err
	}
	return &SES{
		client: client,
	}, nil
}

// Client Client を返します
func (s *SES) Client() *sesv2.SESV2 {
	return s.client
}

func sesClient() (*sesv2.SESV2, error) {
	sess, _, err := aws.Configure(nil)
	if err != nil {
		return nil, err
	}
	client := sesv2.New(sess)
	if client == nil {
		return nil, errors.New("unable create Amazon SES client")
	}
	return client, nil
}
