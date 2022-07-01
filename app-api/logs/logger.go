// Package logs defines how the logger works
package logs

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	sentry "github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
	"github.com/supinf/supinf-mail/app-api/config"
)

// Map マップエイリアス
type Map map[string]interface{}

// ---------------------------------------------------------------------
//  API 管理者・運用者に向けたアプリケーションログ
// ---------------------------------------------------------------------

// Debug アプリケーションログを残します
func Debug(msg string, err error, details *Map) {
	entry, _ := fields(details)
	entry.Debug(message(msg, err))
}

// Info アプリケーションログを残します
func Info(msg string, err error, details *Map) {
	entry, _ := fields(details)
	entry.Info(message(msg, err))
}

// Warn アプリケーションログを残します
func Warn(msg string, err error, details *Map) {
	entry, _ := fields(details)
	entry.Warn(message(msg, err))
}

// Error アプリケーションログを残し、Issue 管理サイトへの報告も行います
func Error(msg string, err error, details *Map) {
	entry, _ := fields(details)
	entry.Error(message(msg, err))
}

// Fatal アプリケーションログを残し、Issue 管理サイトへの報告も行い、アプリケーションを停止します
func Fatal(msg string, err error, details *Map) {
	entry, _ := fields(details)
	entry.Fatal(message(msg, err))
}

// StackTrace スタックトレースの簡易表示
func StackTrace() {
	if config.LogLevel != "debug" {
		return
	}
	for depth := 0; ; depth++ {
		pc, src, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		Warn(fmt.Sprintf(" -> %d: %s: %s(%d)\n", depth, runtime.FuncForPC(pc).Name(), src, line), nil, nil)
	}
}

// SetHooks フックの設定
func SetHooks() {
	if config.SentryDataSource != "" {
		if hook, err := sentry.NewWithTagsSentryHook(config.SentryDataSource, map[string]string{
			"time-zone": fmt.Sprint(config.TimeZone),
		}, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}); err == nil {
			hook.StacktraceConfiguration.Enable = true
			hook.Timeout = time.Duration(0)
			hook.SetRelease(config.AppVersion)
			hook.SetEnvironment(config.AppStage)
			logrus.StandardLogger().Hooks.Add(hook)
			Debug("Sentry configured", nil, nil)
		}
	}
}

type colorType int

// 指定可能な色
const (
	Red    colorType = 31
	Green  colorType = 32
	Yellow colorType = 33
	Blue   colorType = 34
	Gray   colorType = 37
)

// Color ターミナルでのカラー出力
func Color(color colorType, value string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, value)
}

func message(msg string, err error) string {
	if err != nil {
		msg += "(" + err.Error() + ")"
	}
	return msg
}

func fields(details *Map) (*logrus.Entry, logrus.Fields) {
	fields := logrus.Fields{}
	if details != nil {
		for key, value := range *details {
			fields[key] = value
		}
	}
	logger := logrus.StandardLogger()
	switch strings.ToLower(config.LogLevel) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	}
	if len(fields) == 0 {
		logger.Formatter = formatter{}
		return logrus.NewEntry(logger), fields
	}
	logrus.SetFormatter(formatter{})
	return logrus.WithFields(fields), fields
}

// KeyVal key value
type KeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type formatter struct{}

func (f formatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			data[strings.ToLower(k)] = v.Error()
		default:
			data[strings.ToLower(k)] = v
		}
	}
	if _, found := data["level"]; !found {
		data["level"] = entry.Level.String()
	}
	if entry.Message != "" {
		data["message"] = entry.Message
	}
	data["time"] = entry.Time

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

// ToJSONString modelをJSON文字列に変換
func ToJSONString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}
