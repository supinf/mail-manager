// Code generated by go-swagger; DO NOT EDIT.

package suppression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/supinf/supinf-mail/app-api/auth"
)

// GetSuppressionsHandlerFunc turns a function with the right signature into a get suppressions handler
type GetSuppressionsHandlerFunc func(GetSuppressionsParams, *auth.Session) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSuppressionsHandlerFunc) Handle(params GetSuppressionsParams, principal *auth.Session) middleware.Responder {
	return fn(params, principal)
}

// GetSuppressionsHandler interface for that can handle valid get suppressions params
type GetSuppressionsHandler interface {
	Handle(GetSuppressionsParams, *auth.Session) middleware.Responder
}

// NewGetSuppressions creates a new http.Handler for the get suppressions operation
func NewGetSuppressions(ctx *middleware.Context, handler GetSuppressionsHandler) *GetSuppressions {
	return &GetSuppressions{Context: ctx, Handler: handler}
}

/*GetSuppressions swagger:route GET /suppressions suppression getSuppressions

サプレッションリストを取得

アカウントレベルのサプレッションリストを取得

*/
type GetSuppressions struct {
	Context *middleware.Context
	Handler GetSuppressionsHandler
}

func (o *GetSuppressions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSuppressionsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *auth.Session
	if uprinc != nil {
		principal = uprinc.(*auth.Session) // this is really a auth.Session, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}