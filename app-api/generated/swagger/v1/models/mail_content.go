// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MailContent mail content
//
// swagger:model MailContent
type MailContent struct {

	// 本文 (HTML)
	HTML string `json:"html,omitempty"`

	// 本文 (Plain)
	Plain string `json:"plain,omitempty"`

	// 件名
	// Required: true
	Subject *string `json:"subject"`
}

// Validate validates this mail content
func (m *MailContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailContent) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MailContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailContent) UnmarshalBinary(b []byte) error {
	var res MailContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
