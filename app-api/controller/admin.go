package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	v1 "github.com/supinf/supinf-mail/app-api/api/v1"
	"github.com/supinf/supinf-mail/app-api/auth"
	"github.com/supinf/supinf-mail/app-api/errors"
	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/restapi/operations/admin"
)

func postUsagePlan(params admin.PostAdminUsagePlanParams, sess *auth.Session) middleware.Responder {
	// 管理者のみ呼び出し可能
	if !sess.IsAdmin() {
		err := errors.Unauthorized
		return admin.NewPostAdminUsagePlanDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}

	usagePlanID, err := v1.PostUsagePlan(params.Body)
	if err != nil {
		return admin.NewPostAdminUsagePlanDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return admin.NewPostAdminUsagePlanCreated().WithPayload(swag.StringValue(usagePlanID))
}

func postUser(params admin.PostAdminUsersParams, sess *auth.Session) middleware.Responder {
	// 管理者のみ呼び出し可能
	if !sess.IsAdmin() {
		err := errors.Unauthorized
		return admin.NewPostAdminUsersDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}

	apiKey, err := v1.PostUser(params.Body)
	if err != nil {
		return admin.NewPostAdminUsersDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return admin.NewPostAdminUsersCreated().WithPayload(swag.StringValue(apiKey))
}

func patchUserEnabled(params admin.PatchAdminUsersEnabledParams, sess *auth.Session) middleware.Responder {
	// 管理者のみ呼び出し可能
	if !sess.IsAdmin() {
		err := errors.Unauthorized
		return admin.NewPatchAdminUsersEnabledDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}

	if err := v1.PatchUserEnabled(params.Body); err != nil {
		return admin.NewPatchAdminUsersEnabledDefault(err.StatusCode).WithPayload(err.ToAPIs())
	}
	return admin.NewPatchAdminUsersEnabledOK()
}
