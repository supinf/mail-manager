package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	v1 "github.com/supinf/supinf-mail/app-api/api/v1"
	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/suppression"
)

func listSuppression(params suppression.GetSuppressionsParams, _ *auth.Session) middleware.Responder {
	resp, err := v1.ListSuppression(
		params.From,
		params.To,
		params.Reasons,
		params.Limit,
		params.NextToken,
	)
	if err != nil {
		return suppression.NewGetSuppressionsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return suppression.NewGetSuppressionsOK().WithPayload(resp)
}

func getSuppression(params suppression.GetSuppressionsMailParams, _ *auth.Session) middleware.Responder {
	resp, err := v1.GetSuppression(params.Mail)
	if err != nil {
		return suppression.NewGetSuppressionsMailDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return suppression.NewGetSuppressionsMailOK().WithPayload(resp)
}

func postSuppression(params suppression.PostSuppressionsParams, _ *auth.Session) middleware.Responder {
	if err := v1.PostSuppression(params.Body); err != nil {
		return suppression.NewPostSuppressionsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return suppression.NewPostSuppressionsCreated()
}

func deleteSuppression(params suppression.DeleteSuppressionsParams, _ *auth.Session) middleware.Responder {
	if err := v1.DeleteSuppression(params.Body); err != nil {
		return suppression.NewDeleteSuppressionsDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return suppression.NewDeleteSuppressionsOK()
}
