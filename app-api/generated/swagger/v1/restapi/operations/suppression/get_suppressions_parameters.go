// Code generated by go-swagger; DO NOT EDIT.

package suppression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetSuppressionsParams creates a new GetSuppressionsParams object
// no default values defined in spec.
func NewGetSuppressionsParams() GetSuppressionsParams {

	return GetSuppressionsParams{}
}

// GetSuppressionsParams contains all the bound params for the get suppressions operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSuppressions
type GetSuppressionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*検索期間 (From)
	  In: query
	*/
	From *strfmt.DateTime
	/*取得件数上限
	  In: query
	*/
	Limit *int64
	/*検索開始位置のトークン
	  In: query
	*/
	NextToken *string
	/*追加された要因
	  In: query
	*/
	Reasons []string
	/*検索期間 (To)
	  In: query
	*/
	To *strfmt.DateTime
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSuppressionsParams() beforehand.
func (o *GetSuppressionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qNextToken, qhkNextToken, _ := qs.GetOK("nextToken")
	if err := o.bindNextToken(qNextToken, qhkNextToken, route.Formats); err != nil {
		res = append(res, err)
	}

	qReasons, qhkReasons, _ := qs.GetOK("reasons")
	if err := o.bindReasons(qReasons, qhkReasons, route.Formats); err != nil {
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
func (o *GetSuppressionsParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("from", "query", "strfmt.DateTime", raw)
	}
	o.From = (value.(*strfmt.DateTime))

	if err := o.validateFrom(formats); err != nil {
		return err
	}

	return nil
}

// validateFrom carries on validations for parameter From
func (o *GetSuppressionsParams) validateFrom(formats strfmt.Registry) error {

	if err := validate.FormatOf("from", "query", "date-time", o.From.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetSuppressionsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindNextToken binds and validates parameter NextToken from query.
func (o *GetSuppressionsParams) bindNextToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.NextToken = &raw

	return nil
}

// bindReasons binds and validates array parameter Reasons from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetSuppressionsParams) bindReasons(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvReasons string
	if len(rawData) > 0 {
		qvReasons = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	reasonsIC := swag.SplitByFormat(qvReasons, "")
	if len(reasonsIC) == 0 {
		return nil
	}

	var reasonsIR []string
	for i, reasonsIV := range reasonsIC {
		reasonsI := reasonsIV

		if err := validate.EnumCase(fmt.Sprintf("%s.%v", "reasons", i), "query", reasonsI, []interface{}{"BOUNCE", "COMPLAINT"}, true); err != nil {
			return err
		}

		reasonsIR = append(reasonsIR, reasonsI)
	}

	o.Reasons = reasonsIR

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetSuppressionsParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("to", "query", "strfmt.DateTime", raw)
	}
	o.To = (value.(*strfmt.DateTime))

	if err := o.validateTo(formats); err != nil {
		return err
	}

	return nil
}

// validateTo carries on validations for parameter To
func (o *GetSuppressionsParams) validateTo(formats strfmt.Registry) error {

	if err := validate.FormatOf("to", "query", "date-time", o.To.String(), formats); err != nil {
		return err
	}
	return nil
}
