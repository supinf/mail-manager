package auth

import (
	"context"
	"net/http"

	"github.com/supinf/supinf-mail/app-api/misc"
)

// Context リクエストコンテキスト
type Context struct {
	*http.Request
	RequestID string
}

type contextKeyType string

const (
	chars      = "abcdefghijklmnopqrstuvwxyz0123456789"
	idkey      = "x-api-request-id"
	contextKey = contextKeyType("session-context")
)

// NewContext 新しいコンテキストの生成
func NewContext(request *http.Request) *Context {
	if value := request.Context().Value(contextKey); value != nil {
		if ctx, ok := value.(Context); ok {
			return &ctx
		}
	}
	var requestID string
	if candidate := request.Header.Get(idkey); candidate == "" {
		requestID = misc.RandomString([]byte(chars), 16)
		request.Header.Set(idkey, requestID)
	} else {
		requestID = candidate
	}

	ctx := Context{
		RequestID: requestID,
		Request:   request,
	}
	_ = request.WithContext(
		context.WithValue(
			request.Context(),
			contextKey,
			ctx,
		),
	)
	return &ctx
}
