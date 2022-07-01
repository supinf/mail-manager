package errors

import (
	"net/http"
)

// 事前定義済みエラー一覧
var (
	//COMMON系 ERROR ------------------------------------------------------------------
	ServiceUnavailable  = newError(http.StatusServiceUnavailable, "E100000", "error-title.unavailable", "error.unavailable")
	InternalServerError = newError(http.StatusInternalServerError, "E100001", "error-title.internal", "error.internal")
	Unauthorized        = newError(http.StatusUnauthorized, "E100002", "error-title.unauthorized", "error.unauthorized")
	Forbidden           = newError(http.StatusForbidden, "E100003", "error-title.forbidden", "error.forbidden")
	InvalidParameters   = newError(http.StatusBadRequest, "E100004", "error-title.invalid-parameters", "error.invalid-parameters")
	NotFound            = newError(http.StatusNotFound, "E100005", "error-title.not-found", "error.not-found")

	//アプリケーション系 ERROR ------------------------------------------------------------------
	AlreadySubscribed = newError(http.StatusBadRequest, "A100001", "error-title.already-subscribed", "error.already-subscribed")
)
