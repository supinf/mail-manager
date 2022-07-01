// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetMailsHistoriesParams creates a new GetMailsHistoriesParams object
// no default values defined in spec.
func NewGetMailsHistoriesParams() GetMailsHistoriesParams {

	return GetMailsHistoriesParams{}
}

// GetMailsHistoriesParams contains all the bound params for the get mails histories operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMailsHistories
type GetMailsHistoriesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*送信元メールアドレス API Key 発行時に Mail(From) を指定していた場合、このパラメータにはそれと同一のメールアドレスを指定するか、未指定にしてください。
	  In: query
	*/
	From *strfmt.Email
	/*送信日時 (From)
	  In: query
	*/
	SendAtFrom *strfmt.DateTime
	/*送信日時 (To)
	  In: query
	*/
	SendAtTo *strfmt.DateTime
	/*送信先メールアドレス
	  In: query
	*/
	To *strfmt.Email
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMailsHistoriesParams() beforehand.
func (o *GetMailsHistoriesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qSendAtFrom, qhkSendAtFrom, _ := qs.GetOK("sendAtFrom")
	if err := o.bindSendAtFrom(qSendAtFrom, qhkSendAtFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qSendAtTo, qhkSendAtTo, _ := qs.GetOK("sendAtTo")
	if err := o.bindSendAtTo(qSendAtTo, qhkSendAtTo, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetMailsHistoriesParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("from", "query", "strfmt.Email", raw)
	}
	o.From = (value.(*strfmt.Email))

	if err := o.validateFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateFrom carries on validations for parameter From
func (o *GetMailsHistoriesParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "email", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindSendAtFrom binds and validates parameter SendAtFrom from query.
func (o *GetMailsHistoriesParams) bindSendAtFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("sendAtFrom", "query", "strfmt.DateTime", raw)
	}
	o.SendAtFrom = (value.(*strfmt.DateTime))

	if err := o.validateSendAtFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateSendAtFrom carries on validations for parameter SendAtFrom
func (o *GetMailsHistoriesParams) validateSendAtFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("sendAtFrom", "query", "date-time", o.SendAtFrom.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindSendAtTo binds and validates parameter SendAtTo from query.
func (o *GetMailsHistoriesParams) bindSendAtTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date-time
	value, err := formats.Parse("date-time", raw)
	if err != nil {
		return errors.InvalidType("sendAtTo", "query", "strfmt.DateTime", raw)
	}
	o.SendAtTo = (value.(*strfmt.DateTime))

	if err := o.validateSendAtTo(formats); err != nil {
		return err
	}

	return nil
}

// validateSendAtTo carries on validations for parameter SendAtTo
func (o *GetMailsHistoriesParams) validateSendAtTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("sendAtTo", "query", "date-time", o.SendAtTo.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetMailsHistoriesParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: email
	value, err := formats.Parse("email", raw)
	if err != nil {
		return errors.InvalidType("to", "query", "strfmt.Email", raw)
	}
	o.To = (value.(*strfmt.Email))

	if err := o.validateTo(formats); err != nil {
		return err
	}

	return nil
}

// validateTo carries on validations for parameter To
func (o *GetMailsHistoriesParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "email", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}
