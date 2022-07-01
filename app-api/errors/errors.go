package errors

import (
	"net/http"

	"github.com/go-openapi/swag"
	apiModels "github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/i18n"
)

// ---------------------------------------------------------------------
//  ユーザ、またはクライアントアプリに返すエラー
// ---------------------------------------------------------------------

// Error エラー発生に関する情報
type Error struct {
	StatusCode          int
	ErrorCodeForClient  string
	ErrorTitleForClient string
	ErrorMsgForClient   string
}

func newError(httpStatusCode int, errCode, errTitle, errMsg string) *Error {
	return &Error{
		StatusCode:          httpStatusCode,
		ErrorCodeForClient:  errCode,
		ErrorTitleForClient: i18n.Message(errTitle),
		ErrorMsgForClient:   i18n.Message(errMsg),
	}
}

// ToAPIs API としてクライアントサイドへ返すエラー情報を生成します
func (e *Error) ToAPIs() *apiModels.Error {
	return &apiModels.Error{
		Code:    swag.String(e.ErrorCodeForClient),
		Title:   swag.String(e.ErrorTitleForClient),
		Message: swag.String(e.ErrorMsgForClient),
	}
}

// CodeToError ステータスコード（数値）をエラーに変換します
func CodeToError(code int) *Error {
	switch code {
	case http.StatusServiceUnavailable:
		return ServiceUnavailable
	case http.StatusInternalServerError:
		return InternalServerError
	case http.StatusUnauthorized:
		return Unauthorized
	case http.StatusForbidden:
		return Forbidden
	case http.StatusBadRequest:
		return InvalidParameters
	case http.StatusNotFound:
		return NotFound
	default:
		return InternalServerError
	}
}
