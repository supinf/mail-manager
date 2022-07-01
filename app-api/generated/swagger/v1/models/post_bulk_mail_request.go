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

// PostBulkMailRequest post bulk mail request
//
// swagger:model PostBulkMailRequest
type PostBulkMailRequest struct {

	// メール内容
	// Required: true
	Content *MailContent `json:"content"`

	// 送信データ
	// Required: true
	Entries []*PostBulkMailRequestEntriesItems0 `json:"entries"`

	// 送信元メールアドレス API Key 発行時に Mail(From) を指定していた場合、このパラメータにはそれと同一のメールアドレスを指定するか、未指定にしてください。 また、API Key 発行時に Mail(From) を指定していなかった場合、このパラメータは必須になります。
	From *MailAddress `json:"from,omitempty"`
}

// Validate validates this post bulk mail request
func (m *PostBulkMailRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostBulkMailRequest) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *PostBulkMailRequest) validateEntries(formats strfmt.Registry) error {

	if err := validate.Required("entries", "body", m.Entries); err != nil {
		return err
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostBulkMailRequest) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {
		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostBulkMailRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostBulkMailRequest) UnmarshalBinary(b []byte) error {
	var res PostBulkMailRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostBulkMailRequestEntriesItems0 post bulk mail request entries items0
//
// swagger:model PostBulkMailRequestEntriesItems0
type PostBulkMailRequestEntriesItems0 struct {

	// 送信先メールアドレス
	// Required: true
	Destination *MailDestination `json:"destination"`

	// テンプレートへのマッピングデータ
	Map *Map `json:"map,omitempty"`
}

// Validate validates this post bulk mail request entries items0
func (m *PostBulkMailRequestEntriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostBulkMailRequestEntriesItems0) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

func (m *PostBulkMailRequestEntriesItems0) validateMap(formats strfmt.Registry) error {

	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if m.Map != nil {
		if err := m.Map.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostBulkMailRequestEntriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostBulkMailRequestEntriesItems0) UnmarshalBinary(b []byte) error {
	var res PostBulkMailRequestEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}