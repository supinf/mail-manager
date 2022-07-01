package config

import (
	"os"
	"testing"
	"time"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	os.Setenv("APP_VERSION", "0.1.0")
	os.Setenv("TIME_ZONE", "Asia/Tokyo")

	cfg := config{}
	env.Parse(&cfg)

	assert.Equal(t, "0.1.0", cfg.Version, "Expected an error when Version is not '0.1.0'")
	timezone := time.UTC
	if location, err := time.LoadLocation(cfg.TimeZone); err == nil {
		timezone = location
	}
	assert.NotEqual(t, time.UTC, timezone, "Expected an error when timezone is 'time.UTC'")
}
