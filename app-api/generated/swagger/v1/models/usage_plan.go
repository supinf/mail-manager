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

// UsagePlan usage plan
//
// swagger:model UsagePlan
type UsagePlan struct {

	// API ステージリスト
	APIStage []*APIStage `json:"apiStage"`

	// プラン名
	// Required: true
	Name *string `json:"name"`

	// クォーター（特定期間内の最大リクエスト可能数）
	Quota *Quota `json:"quota,omitempty"`

	// スロットル（リクエストレート）
	Throttle *Throttle `json:"throttle,omitempty"`
}

// Validate validates this usage plan
func (m *UsagePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsagePlan) validateAPIStage(formats strfmt.Registry) error {

	if swag.IsZero(m.APIStage) { // not required
		return nil
	}

	for i := 0; i < len(m.APIStage); i++ {
		if swag.IsZero(m.APIStage[i]) { // not required
			continue
		}

		if m.APIStage[i] != nil {
			if err := m.APIStage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apiStage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UsagePlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UsagePlan) validateQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.Quota) { // not required
		return nil
	}

	if m.Quota != nil {
		if err := m.Quota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quota")
			}
			return err
		}
	}

	return nil
}

func (m *UsagePlan) validateThrottle(formats strfmt.Registry) error {

	if swag.IsZero(m.Throttle) { // not required
		return nil
	}

	if m.Throttle != nil {
		if err := m.Throttle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throttle")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsagePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsagePlan) UnmarshalBinary(b []byte) error {
	var res UsagePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}