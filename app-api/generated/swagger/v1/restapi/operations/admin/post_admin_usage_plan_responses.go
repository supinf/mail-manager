// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
)

// PostAdminUsagePlanCreatedCode is the HTTP code returned for type PostAdminUsagePlanCreated
const PostAdminUsagePlanCreatedCode int = 201

/*PostAdminUsagePlanCreated OK

swagger:response postAdminUsagePlanCreated
*/
type PostAdminUsagePlanCreated struct {

	/*Usage Plan ID
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAdminUsagePlanCreated creates PostAdminUsagePlanCreated with default headers values
func NewPostAdminUsagePlanCreated() *PostAdminUsagePlanCreated {

	return &PostAdminUsagePlanCreated{}
}

// WithPayload adds the payload to the post admin usage plan created response
func (o *PostAdminUsagePlanCreated) WithPayload(payload string) *PostAdminUsagePlanCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post admin usage plan created response
func (o *PostAdminUsagePlanCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAdminUsagePlanCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostAdminUsagePlanDefault unexpected error

swagger:response postAdminUsagePlanDefault
*/
type PostAdminUsagePlanDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostAdminUsagePlanDefault creates PostAdminUsagePlanDefault with default headers values
func NewPostAdminUsagePlanDefault(code int) *PostAdminUsagePlanDefault {
	if code <= 0 {
		code = 500
	}

	return &PostAdminUsagePlanDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post admin usage plan default response
func (o *PostAdminUsagePlanDefault) WithStatusCode(code int) *PostAdminUsagePlanDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post admin usage plan default response
func (o *PostAdminUsagePlanDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post admin usage plan default response
func (o *PostAdminUsagePlanDefault) WithPayload(payload *models.Error) *PostAdminUsagePlanDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post admin usage plan default response
func (o *PostAdminUsagePlanDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAdminUsagePlanDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}