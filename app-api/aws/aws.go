package aws

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/supinf/supinf-mail/app-api/logs"
)

// Configure 新しいセッションを生成します
func Configure(region *string) (*session.Session, *awssdk.Config, error) {
	level := awssdk.LogOff
	if os.Getenv("DEBUG") == "1" {
		level = awssdk.LogDebug
	}
	cfg := &awssdk.Config{
		Region:   awssdk.String(os.Getenv("AWS_DEFAULT_REGION")),
		LogLevel: &level,
	}
	if region != nil {
		cfg.Region = region
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, nil, err
	}
	if os.Getenv("DEBUG") == "1" {
		if out, err := sts.New(sess).GetCallerIdentityWithContext(
			context.Background(), &sts.GetCallerIdentityInput{}); err == nil {
			logs.Debug("get-caller-identity", nil, &logs.Map{
				"Account": awssdk.StringValue(out.Account),
				"UserId":  awssdk.StringValue(out.UserId),
				"ARN":     awssdk.StringValue(out.Arn),
			})
		}
	}
	return sess, cfg, nil
}

// IPAddress リモートからの IP アドレスを考慮して返します
func IPAddress(request *http.Request) (string, bool) {
	forwardedFor := request.Header.Get("X-Forwarded-For")
	if forwardedFor != "" {
		return strings.TrimSpace(strings.Split(forwardedFor, ",")[0]), true
	}
	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return request.RemoteAddr, false
	}
	return ip, false
}
