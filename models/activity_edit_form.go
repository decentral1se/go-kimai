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

// ActivityEditForm activity edit form
//
// swagger:model ActivityEditForm
type ActivityEditForm struct {

	// billable
	Billable bool `json:"billable,omitempty"`

	// The hexadecimal color code (default: #d2d6de)
	Color string `json:"color,omitempty"`

	// comment
	Comment string `json:"comment,omitempty"`

	// invoice text
	InvoiceText string `json:"invoiceText,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// Project ID
	Project int64 `json:"project,omitempty"`

	// visible
	Visible bool `json:"visible,omitempty"`
}

// Validate validates this activity edit form
func (m *ActivityEditForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityEditForm) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this activity edit form based on context it is used
func (m *ActivityEditForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivityEditForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityEditForm) UnmarshalBinary(b []byte) error {
	var res ActivityEditForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
