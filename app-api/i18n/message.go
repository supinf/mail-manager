package i18n

import (
	"os"
	"path/filepath"

	"github.com/unknwon/i18n" // nolint
)

const jaJP = "ja-JP"

// APISrcPrefix ファイルパスプリフィックス
var APISrcPrefix string

var ja i18n.Locale

func init() {
	APISrcPrefix = filepath.Join(os.Getenv("GOPATH"), "src/github.com/supinf/supinf-mail/app-api")
	i18n.SetMessage(jaJP, filepath.Join(APISrcPrefix, "i18n/locale_"+jaJP+".ini"))
	ja = i18n.Locale{Lang: jaJP}
}

// Message 日本語でのメッセージを返します
func Message(key string, args ...interface{}) string {
	return ja.Tr(key, args...)
}
