// Code generated by go-swagger; DO NOT EDIT.

package suppression

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
)

// PostSuppressionsCreatedCode is the HTTP code returned for type PostSuppressionsCreated
const PostSuppressionsCreatedCode int = 201

/*PostSuppressionsCreated OK

swagger:response postSuppressionsCreated
*/
type PostSuppressionsCreated struct {
}

// NewPostSuppressionsCreated creates PostSuppressionsCreated with default headers values
func NewPostSuppressionsCreated() *PostSuppressionsCreated {

	return &PostSuppressionsCreated{}
}

// WriteResponse to the client
func (o *PostSuppressionsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostSuppressionsDefault unexpected error

swagger:response postSuppressionsDefault
*/
type PostSuppressionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostSuppressionsDefault creates PostSuppressionsDefault with default headers values
func NewPostSuppressionsDefault(code int) *PostSuppressionsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostSuppressionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post suppressions default response
func (o *PostSuppressionsDefault) WithStatusCode(code int) *PostSuppressionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post suppressions default response
func (o *PostSuppressionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post suppressions default response
func (o *PostSuppressionsDefault) WithPayload(payload *models.Error) *PostSuppressionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post suppressions default response
func (o *PostSuppressionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostSuppressionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
