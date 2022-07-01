package config

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env"
)

// ApplicationsName 機能名
const ApplicationsName = "SUPINF-MAIL-APIs"

// ApplicationsNameShort 機能名（略称）
const ApplicationsNameShort = "supinf-mail"

// 設定値
var (
	AppVersion       string
	AppStage         string
	LogLevel         string
	AllowCORS        bool
	AllowCORSOrigin  string
	SecuredTransport bool
	TimeZone         = time.UTC
	RunningOnAWS     bool
	AwsXRay          bool
	AwsXRayLogLevel  string
	SentryDataSource string
)

type config struct { // nolint:maligned
	Version          string `env:"APP_VERSION" envDefault:""`
	AppStage         string `env:"APP_STAGE" envDefault:"local"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"warn"`
	AllowCORS        bool   `env:"ALLOW_CORS" envDefault:"true"`
	AllowCORSOrigin  string `env:"ALLOW_CORS_ORIGIN" envDefault:"*"`
	SecuredTransport bool   `env:"SECURED_TRANSPORT" envDefault:"false"`
	TimeZone         string `env:"TIME_ZONE" envDefault:"UTC"`
	RunningOnAWS     bool   `env:"ON_AWS" envDefault:"false"`
	AwsXRayDaemon    string `env:"AWS_XRAY_DAEMON_ADDRESS" envDefault:""`
	AwsXRayLogLvl    string `env:"AWS_XRAY_LOG_LEVEL" envDefault:"info"`
	SentryDataSource string `env:"SENTRY_DSN" envDefault:""`
}

func init() {
	Set()
}

// Set sets configurations via envoronment variables
func Set() {
	cfg := config{}
	env.Parse(&cfg)

	AppVersion = cfg.Version
	AppStage = cfg.AppStage
	LogLevel = cfg.LogLevel
	AllowCORS = cfg.AllowCORS
	AllowCORSOrigin = cfg.AllowCORSOrigin
	SecuredTransport = cfg.SecuredTransport
	RunningOnAWS = cfg.RunningOnAWS
	AwsXRayLogLevel = cfg.AwsXRayLogLvl
	SentryDataSource = cfg.SentryDataSource

	if location, err := time.LoadLocation(cfg.TimeZone); err == nil {
		TimeZone = location
	}

	AwsXRay = cfg.AwsXRayDaemon != ""
	if strings.Contains(cfg.AwsXRayDaemon, "localhost") {
		cfg.AwsXRayDaemon = strings.Replace(cfg.AwsXRayDaemon, "localhost", getHostIP(), -1)
		os.Setenv("AWS_XRAY_DAEMON_ADDRESS", cfg.AwsXRayDaemon)
	}
}

func getHostIP() string {
	client := &http.Client{
		Timeout: time.Duration(1) * time.Second,
	}
	res, err := client.Get("http://169.254.170.2/v2/metadata")
	if err != nil {
		return "localhost"
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	return "localhost"
	// }
	// return string(body)
	return "127.0.0.1"
}
