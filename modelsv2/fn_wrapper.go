// Code generated by go-swagger; DO NOT EDIT.

package modelsv2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FnWrapper fn wrapper
// swagger:model FnWrapper
type FnWrapper struct {

	// error
	Error *ErrorBody `json:"error,omitempty"`

	// fn
	// Required: true
	Fn *Fn `json:"fn"`
}

// Validate validates this fn wrapper
func (m *FnWrapper) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FnWrapper) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *FnWrapper) validateFn(formats strfmt.Registry) error {

	if err := validate.Required("fn", "body", m.Fn); err != nil {
		return err
	}

	if m.Fn != nil {
		if err := m.Fn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FnWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FnWrapper) UnmarshalBinary(b []byte) error {
	var res FnWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
