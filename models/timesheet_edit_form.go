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

// TimesheetEditForm timesheet edit form
//
// swagger:model TimesheetEditForm
type TimesheetEditForm struct {

	// Activity ID
	// Required: true
	Activity *int64 `json:"activity"`

	// begin
	// Example: 2023-01-14T05:12:43
	// Required: true
	// Format: date-time
	Begin *strfmt.DateTime `json:"begin"`

	// billable
	Billable bool `json:"billable,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end
	// Example: 2023-01-14T05:12:43
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// exported
	Exported bool `json:"exported,omitempty"`

	// Fixed rate
	FixedRate float64 `json:"fixedRate,omitempty"`

	// Hourly rate
	HourlyRate float64 `json:"hourlyRate,omitempty"`

	// Project ID
	// Required: true
	Project *int64 `json:"project"`

	// Comma separated list of tags
	Tags string `json:"tags,omitempty"`

	// User ID
	User int64 `json:"user,omitempty"`
}

// Validate validates this timesheet edit form
func (m *TimesheetEditForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TimesheetEditForm) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", m.Activity); err != nil {
		return err
	}

	return nil
}

func (m *TimesheetEditForm) validateBegin(formats strfmt.Registry) error {

	if err := validate.Required("begin", "body", m.Begin); err != nil {
		return err
	}

	if err := validate.FormatOf("begin", "body", "date-time", m.Begin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimesheetEditForm) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimesheetEditForm) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this timesheet edit form based on context it is used
func (m *TimesheetEditForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimesheetEditForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimesheetEditForm) UnmarshalBinary(b []byte) error {
	var res TimesheetEditForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
