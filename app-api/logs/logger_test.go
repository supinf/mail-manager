package logs

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/supinf/supinf-mail/app-api/config"
)

type LogsTestSuite struct {
	suite.Suite
	buf *bytes.Buffer
}

func (suite *LogsTestSuite) SetupTest() {
	suite.buf = &bytes.Buffer{}
	logrus.SetOutput(suite.buf)
}

func (suite *LogsTestSuite) Out() string {
	return suite.buf.String()
}

func TestLogs(t *testing.T) {
	suite.Run(t, new(LogsTestSuite))
}

// Debug が期待した文字列を出力する
func (suite *LogsTestSuite) TestDebug1() {
	os.Setenv("LOG_LEVEL", "debug")
	config.Set()
	expect := "foo"

	Debug(expect, nil, nil)

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"\"")
}

// Map を渡しても Debug が期待した文字列を出力する
func (suite *LogsTestSuite) TestDebug2() {
	os.Setenv("LOG_LEVEL", "debug")
	config.Set()
	expect := "foo"

	Debug(expect, nil, &Map{"detail": expect})

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"\"")
	assert.Contains(suite.T(), suite.Out(), "\"detail\":\""+expect)
}

// Error を渡しても Debug が期待した文字列を出力する
func (suite *LogsTestSuite) TestDebug3() {
	os.Setenv("LOG_LEVEL", "debug")
	config.Set()
	expect := "foo"

	Debug(expect, errors.New(expect), nil)

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"("+expect+")\"")
}

// Info が期待した文字列を出力する
func (suite *LogsTestSuite) TestInfo() {
	os.Setenv("LOG_LEVEL", "info")
	config.Set()
	expect := "foo"

	Info(expect, nil, &Map{"detail": expect})

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"\"")
	assert.Contains(suite.T(), suite.Out(), "\"detail\":\""+expect)
}

// Warn が期待した文字列を出力する
func (suite *LogsTestSuite) TestWarn() {
	os.Setenv("LOG_LEVEL", "warn")
	config.Set()
	expect := "foo"

	Warn(expect, nil, &Map{"detail": expect})

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"\"")
	assert.Contains(suite.T(), suite.Out(), "\"detail\":\""+expect)
}

// Error が期待した文字列を出力する
func (suite *LogsTestSuite) TestError() {
	os.Setenv("LOG_LEVEL", "error")
	config.Set()
	expect := "foo"

	Error(expect, nil, &Map{"detail": expect})

	assert.Contains(suite.T(), suite.Out(), "\"message\":\""+expect+"\"")
	assert.Contains(suite.T(), suite.Out(), "\"detail\":\""+expect)
}

// ログレベルが fatal の場合はなにも出力されない
func (suite *LogsTestSuite) TestFatal() {
	os.Setenv("LOG_LEVEL", "fatal")
	config.Set()
	expect := "foo"

	Debug(expect, nil, &Map{"detail": expect})

	assert.Empty(suite.T(), suite.Out())
}

// StackTrace がこの関数そのものを出力する
func (suite *LogsTestSuite) TestStackTrace1() {
	os.Setenv("LOG_LEVEL", "debug")
	config.Set()

	StackTrace()

	assert.Contains(suite.T(), suite.Out(), "TestStackTrace")
}

// ログレベルが debug 出ない場合、StackTrace は何も出力しない
func (suite *LogsTestSuite) TestStackTrace2() {
	os.Setenv("LOG_LEVEL", "info")
	config.Set()

	StackTrace()

	assert.Empty(suite.T(), suite.Out())
}

// SENTRY_DSN が設定されていなければ、フックは作成されない
func (suite *LogsTestSuite) TestSetHooks1() {
	os.Unsetenv("SENTRY_DSN")
	config.Set()

	SetHooks()

	assert.Empty(suite.T(), logrus.StandardLogger().Hooks)
}

// SENTRY_DSN が設定されていても、不正な DSN の場合 logrus にフックは作成されない
func (suite *LogsTestSuite) TestSetHooks2() {
	os.Setenv("SENTRY_DSN", "https://sentry.io/123123")
	config.Set()

	SetHooks()

	assert.Empty(suite.T(), logrus.StandardLogger().Hooks)
}

// カラー整形された文字列が返る
func (suite *LogsTestSuite) TestColor() {
	expect := Red
	actual := Color(expect, "foo")
	assert.Contains(suite.T(), actual, fmt.Sprintf("%d", expect))
}
