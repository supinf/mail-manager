package controllers

import (
	"github.com/go-openapi/strfmt"
	"github.com/supinf/supinf-mail/app-api/errors"
	"github.com/supinf/supinf-mail/app-api/logs"
	"github.com/supinf/supinf-mail/app-api/misc"
)

// validateMailFrom API Key 発行時の Mail(From Address or Domain) とパラメータの From を比較し、検証します
func validateMailFrom(apiKeyFrom string, paramFrom *strfmt.Email) (*strfmt.Email, *errors.Error) {
	// API Key 発行時にドメインのみ指定していた場合
	if misc.IsMailDomainOnly(apiKeyFrom) {
		if paramFrom == nil || !misc.IsSameMailDomain(apiKeyFrom, paramFrom.String()) {
			logs.Error("need to specify an email address with a domain associated with the API Key", nil, nil)
			return nil, errors.InvalidParameters
		}
		return paramFrom, nil
	}

	// API Key 発行時にメールアドレスを指定していた場合
	if paramFrom != nil && apiKeyFrom != paramFrom.String() {
		logs.Error("cannot send from this email address. check the email address associated with the API Key.", nil, nil)
		return nil, errors.InvalidParameters
	}
	from := strfmt.Email(apiKeyFrom)
	return &from, nil
}
