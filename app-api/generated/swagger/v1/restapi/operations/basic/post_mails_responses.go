// Code generated by go-swagger; DO NOT EDIT.

package basic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
)

// PostMailsCreatedCode is the HTTP code returned for type PostMailsCreated
const PostMailsCreatedCode int = 201

/*PostMailsCreated OK

swagger:response postMailsCreated
*/
type PostMailsCreated struct {
}

// NewPostMailsCreated creates PostMailsCreated with default headers values
func NewPostMailsCreated() *PostMailsCreated {

	return &PostMailsCreated{}
}

// WriteResponse to the client
func (o *PostMailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostMailsDefault unexpected error

swagger:response postMailsDefault
*/
type PostMailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostMailsDefault creates PostMailsDefault with default headers values
func NewPostMailsDefault(code int) *PostMailsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostMailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post mails default response
func (o *PostMailsDefault) WithStatusCode(code int) *PostMailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post mails default response
func (o *PostMailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post mails default response
func (o *PostMailsDefault) WithPayload(payload *models.Error) *PostMailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post mails default response
func (o *PostMailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
