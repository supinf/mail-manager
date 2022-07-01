// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supinf/supinf-mail/app-api/generated/swagger/v1/models"
)

// PatchAdminUsersEnabledOKCode is the HTTP code returned for type PatchAdminUsersEnabledOK
const PatchAdminUsersEnabledOKCode int = 200

/*PatchAdminUsersEnabledOK OK

swagger:response patchAdminUsersEnabledOK
*/
type PatchAdminUsersEnabledOK struct {
}

// NewPatchAdminUsersEnabledOK creates PatchAdminUsersEnabledOK with default headers values
func NewPatchAdminUsersEnabledOK() *PatchAdminUsersEnabledOK {

	return &PatchAdminUsersEnabledOK{}
}

// WriteResponse to the client
func (o *PatchAdminUsersEnabledOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*PatchAdminUsersEnabledDefault unexpected error

swagger:response patchAdminUsersEnabledDefault
*/
type PatchAdminUsersEnabledDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchAdminUsersEnabledDefault creates PatchAdminUsersEnabledDefault with default headers values
func NewPatchAdminUsersEnabledDefault(code int) *PatchAdminUsersEnabledDefault {
	if code <= 0 {
		code = 500
	}

	return &PatchAdminUsersEnabledDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the patch admin users enabled default response
func (o *PatchAdminUsersEnabledDefault) WithStatusCode(code int) *PatchAdminUsersEnabledDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the patch admin users enabled default response
func (o *PatchAdminUsersEnabledDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the patch admin users enabled default response
func (o *PatchAdminUsersEnabledDefault) WithPayload(payload *models.Error) *PatchAdminUsersEnabledDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch admin users enabled default response
func (o *PatchAdminUsersEnabledDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAdminUsersEnabledDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}