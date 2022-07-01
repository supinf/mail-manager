// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MailDestination mail destination
//
// swagger:model MailDestination
type MailDestination struct {

	// 送信先メールアドレス (Bcc)
	Bcc []*MailAddress `json:"bcc"`

	// 送信先メールアドレス (Cc)
	Cc []*MailAddress `json:"cc"`

	// 送信先メールアドレス (To)
	// Required: true
	To *MailAddress `json:"to"`
}

// Validate validates this mail destination
func (m *MailDestination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailDestination) validateBcc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bcc) { // not required
		return nil
	}

	for i := 0; i < len(m.Bcc); i++ {
		if swag.IsZero(m.Bcc[i]) { // not required
			continue
		}

		if m.Bcc[i] != nil {
			if err := m.Bcc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bcc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MailDestination) validateCc(formats strfmt.Registry) error {

	if swag.IsZero(m.Cc) { // not required
		return nil
	}

	for i := 0; i < len(m.Cc); i++ {
		if swag.IsZero(m.Cc[i]) { // not required
			continue
		}

		if m.Cc[i] != nil {
			if err := m.Cc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MailDestination) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MailDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailDestination) UnmarshalBinary(b []byte) error {
	var res MailDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}