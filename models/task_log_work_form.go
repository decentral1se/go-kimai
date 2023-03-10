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

// TaskLogWorkForm task log work form
//
// swagger:model TaskLogWorkForm
type TaskLogWorkForm struct {

	// begin
	// Example: 2023-01-14T05:12:43
	// Format: date-time
	Begin strfmt.DateTime `json:"begin,omitempty"`

	// end
	// Example: 2023-01-14T05:12:43
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`
}

// Validate validates this task log work form
func (m *TaskLogWorkForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskLogWorkForm) validateBegin(formats strfmt.Registry) error {
	if swag.IsZero(m.Begin) { // not required
		return nil
	}

	if err := validate.FormatOf("begin", "body", "date-time", m.Begin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaskLogWorkForm) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task log work form based on context it is used
func (m *TaskLogWorkForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskLogWorkForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskLogWorkForm) UnmarshalBinary(b []byte) error {
	var res TaskLogWorkForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
