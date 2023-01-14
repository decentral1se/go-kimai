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

// ActivityExpanded activity expanded
//
// swagger:model ActivityExpanded
type ActivityExpanded struct {

	// billable
	// Required: true
	Billable *bool `json:"billable"`

	// The assigned color in HTML hex format, eg. #dd1d00
	Color string `json:"color,omitempty"`

	// Description of this activity
	Comment string `json:"comment,omitempty"`

	// Internal ID
	ID int64 `json:"id,omitempty"`

	// Name of this activity
	// Required: true
	// Max Length: 150
	// Min Length: 2
	Name *string `json:"name"`

	// project
	Project *Project `json:"project,omitempty"`

	// Whether this activity is visible and can be used for timesheets
	// Required: true
	Visible *bool `json:"visible"`
}

// Validate validates this activity expanded
func (m *ActivityExpanded) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
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

func (m *ActivityExpanded) validateBillable(formats strfmt.Registry) error {

	if err := validate.Required("billable", "body", m.Billable); err != nil {
		return err
	}

	return nil
}

func (m *ActivityExpanded) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 150); err != nil {
		return err
	}

	return nil
}

func (m *ActivityExpanded) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityExpanded) validateVisible(formats strfmt.Registry) error {

	if err := validate.Required("visible", "body", m.Visible); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this activity expanded based on the context it is used
func (m *ActivityExpanded) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityExpanded) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityExpanded) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityExpanded) UnmarshalBinary(b []byte) error {
	var res ActivityExpanded
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}