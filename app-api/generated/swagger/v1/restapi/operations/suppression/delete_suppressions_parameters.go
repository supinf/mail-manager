// Code generated by go-swagger; DO NOT EDIT.

package suppression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
)

// NewDeleteSuppressionsParams creates a new DeleteSuppressionsParams object
// no default values defined in spec.
func NewDeleteSuppressionsParams() DeleteSuppressionsParams {

	return DeleteSuppressionsParams{}
}

// DeleteSuppressionsParams contains all the bound params for the delete suppressions operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteSuppressions
type DeleteSuppressionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*対象メールアドレス
	  Required: true
	  In: body
	*/
	Body *models.DeleteSuppressionRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteSuppressionsParams() beforehand.
func (o *DeleteSuppressionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DeleteSuppressionRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
