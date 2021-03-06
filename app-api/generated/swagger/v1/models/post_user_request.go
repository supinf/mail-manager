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

// PostUserRequest post user request
//
// swagger:model PostUserRequest
type PostUserRequest struct {

	// Usage Plan ID
	// Required: true
	UsagePlanID *string `json:"usagePlanID"`

	// 利用者情報 user.mail を指定した場合、メール送信 API や 履歴取得 API の利用時にそれが使われるようになります。
	// Required: true
	User *User `json:"user"`
}

// Validate validates this post user request
func (m *PostUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsagePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostUserRequest) validateUsagePlanID(formats strfmt.Registry) error {

	if err := validate.Required("usagePlanID", "body", m.UsagePlanID); err != nil {
		return err
	}

	return nil
}

func (m *PostUserRequest) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUserRequest) UnmarshalBinary(b []byte) error {
	var res PostUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
