package errors

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

type WrappedError struct {
	orig awserr.Error
}

func Error(err error) *WrappedError {
	if err == nil {
		return nil
	}
	if e, ok := err.(awserr.Error); ok {
		return &WrappedError{
			orig: e,
		}
	}
	return nil
}

func Code(err error) *int {
	return Error(err).Code()
}

func Message(err error) *string {
	return Error(err).Message()
}

func (e *WrappedError) OrigError() awserr.Error {
	return e.orig
}

func (e *WrappedError) Code() *int {
	if e == nil || e.orig == nil {
		return nil
	}
	// TODO: エラーパターン追加
	switch e.orig.Code() {
	case "NotFoundException":
		return aws.Int(http.StatusNotFound)
	default:
		return aws.Int(http.StatusInternalServerError)
	}
}

func (e *WrappedError) Message() *string {
	if e == nil || e.orig == nil {
		return nil
	}
	return aws.String(e.orig.Message())
}
