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

// ActivityMeta activity meta
//
// swagger:model ActivityMeta
type ActivityMeta struct {

	// Name of the meta (custom) field
	// Required: true
	// Max Length: 50
	// Min Length: 2
	Name *string `json:"name"`

	// Value of the meta (custom) field
	// Max Length: 65535
	Value string `json:"value,omitempty"`
}

// Validate validates this activity meta
func (m *ActivityMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityMeta) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *ActivityMeta) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.MaxLength("value", "body", m.Value, 65535); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this activity meta based on context it is used
func (m *ActivityMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivityMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityMeta) UnmarshalBinary(b []byte) error {
	var res ActivityMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
