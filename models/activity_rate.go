// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActivityRate activity rate
//
// swagger:model ActivityRate
type ActivityRate struct {

	// id
	ID int64 `json:"id,omitempty"`

	// internal rate
	InternalRate float32 `json:"internalRate,omitempty"`

	// is fixed
	// Required: true
	IsFixed *bool `json:"isFixed"`

	// rate
	// Minimum: 0
	Rate *float32 `json:"rate,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this activity rate
func (m *ActivityRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsFixed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
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

func (m *ActivityRate) validateIsFixed(formats strfmt.Registry) error {

	if err := validate.Required("isFixed", "body", m.IsFixed); err != nil {
		return err
	}

	return nil
}

func (m *ActivityRate) validateRate(formats strfmt.Registry) error {
	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := validate.Minimum("rate", "body", float64(*m.Rate), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ActivityRate) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this activity rate based on the context it is used
func (m *ActivityRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityRate) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityRate) UnmarshalBinary(b []byte) error {
	var res ActivityRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
