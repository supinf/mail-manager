package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	v1 "github.com/supinf/supinf-mail/app-api/api/v1"
	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/errors"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/basic"
)

func postMail(params basic.PostMailsParams, sess *auth.Session) middleware.Responder {
	// Mail(From) の検証
	from, err := validatePostMailFrom(sess.UserMail, params.Body.From)
	if err != nil {
		return basic.NewPostMailsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	params.Body.From = from

	if err := v1.PostMail(params.Body, sess.UserName); err != nil {
		return basic.NewPostMailsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return basic.NewPostMailsCreated()
}

func postBulkMail(params basic.PostBulkMailsParams, sess *auth.Session) middleware.Responder {
	// Mail(From) の検証
	from, err := validatePostMailFrom(sess.UserMail, params.Body.From)
	if err != nil {
		return basic.NewPostBulkMailsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	params.Body.From = from

	if err := v1.PostBulkMail(params.Body, sess.UserName); err != nil {
		return basic.NewPostBulkMailsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return basic.NewPostBulkMailsCreated()
}

func validatePostMailFrom(apiKeyFrom string, paramFrom *models.MailAddress) (*models.MailAddress, *errors.Error) {
	var pFrom *strfmt.Email
	if paramFrom != nil && paramFrom.Address != nil {
		pFrom = paramFrom.Address
	}
	from, err := validateMailFrom(apiKeyFrom, pFrom)
	if err != nil {
		return nil, err
	}
	if from == nil {
		return nil, errors.InvalidParameters
	}
	return &models.MailAddress{
		Address: from,
	}, nil
}
