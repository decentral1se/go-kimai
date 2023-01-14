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

// ExpenseCategory expense category
//
// swagger:model ExpenseCategory
type ExpenseCategory struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 2
	Name *string `json:"name"`

	// visible
	// Required: true
	Visible *bool `json:"visible"`
}

// Validate validates this expense category
func (m *ExpenseCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisible(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpenseCategory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *ExpenseCategory) validateVisible(formats strfmt.Registry) error {

	if err := validate.Required("visible", "body", m.Visible); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this expense category based on context it is used
func (m *ExpenseCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExpenseCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpenseCategory) UnmarshalBinary(b []byte) error {
	var res ExpenseCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
