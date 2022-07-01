package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/supinf/supinf-mail/app-api/api/v1"
	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/history"
)

func listHistory(params history.GetMailsHistoriesParams, sess *auth.Session) middleware.Responder {
	// Mail(From) の検証
	from, err := validateMailFrom(sess.UserMail, params.From)
	if err != nil {
		return history.NewGetMailsHistoriesDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}

	resp, err := v1.ListHistory(sess.UserName, from, params.To, params.SendAtFrom, params.SendAtTo)
	if err != nil {
		return history.NewGetMailsHistoriesDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return history.NewGetMailsHistoriesOK().WithPayload(resp)
}
