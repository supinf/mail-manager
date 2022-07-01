// Code generated by go-swagger; DO NOT EDIT.

package suppression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/supinf/supinf-mail/app-api/auth"
)

// PostSuppressionsHandlerFunc turns a function with the right signature into a post suppressions handler
type PostSuppressionsHandlerFunc func(PostSuppressionsParams, *auth.Session) middleware.Responder

// Handle executing the request and returning a response
func (fn PostSuppressionsHandlerFunc) Handle(params PostSuppressionsParams, principal *auth.Session) middleware.Responder {
	return fn(params, principal)
}

// PostSuppressionsHandler interface for that can handle valid post suppressions params
type PostSuppressionsHandler interface {
	Handle(PostSuppressionsParams, *auth.Session) middleware.Responder
}

// NewPostSuppressions creates a new http.Handler for the post suppressions operation
func NewPostSuppressions(ctx *middleware.Context, handler PostSuppressionsHandler) *PostSuppressions {
	return &PostSuppressions{Context: ctx, Handler: handler}
}

/*PostSuppressions swagger:route POST /suppressions suppression postSuppressions

サプレッションリストへ追加

アカウントレベルのサプレッションリストへ追加

*/
type PostSuppressions struct {
	Context *middleware.Context
	Handler PostSuppressionsHandler
}

func (o *PostSuppressions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostSuppressionsParams()

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
