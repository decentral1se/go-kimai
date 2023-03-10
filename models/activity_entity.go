// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActivityEntity activity entity
//
// swagger:model ActivityEntity
type ActivityEntity struct {

	// billable
	// Required: true
	Billable *bool `json:"billable"`

	// The total monetary budget, will be zero if unconfigured.
	// Required: true
	// Maximum: 9e+11
	// Minimum: 0
	Budget *float32 `json:"budget"`

	// The type of budget:
	//  - null      = default / full time
	//  - month     = monthly budget
	BudgetType string `json:"budgetType,omitempty"`

	// The assigned color in HTML hex format, eg. #dd1d00
	Color string `json:"color,omitempty"`

	// Description of this activity
	Comment string `json:"comment,omitempty"`

	// Internal ID
	ID int64 `json:"id,omitempty"`

	// Meta fields
	//
	// All visible meta (custom) fields registered with this activity
	MetaFields []*ActivityMeta `json:"metaFields"`

	// Name of this activity
	// Required: true
	// Max Length: 150
	// Min Length: 2
	Name *string `json:"name"`

	// parent title
	ParentTitle string `json:"parentTitle,omitempty"`

	// project
	Project int64 `json:"project,omitempty"`

	// Teams
	//
	// If no team is assigned, everyone can access the activity
	Teams []*Team `json:"teams"`

	// The time budget in seconds, will be zero if unconfigured.
	// Required: true
	// Maximum: 2.1456e+09
	// Minimum: 0
	TimeBudget *int64 `json:"timeBudget"`

	// Whether this activity is visible and can be used for timesheets
	// Required: true
	Visible *bool `json:"visible"`
}

// Validate validates this activity entity
func (m *ActivityEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBudget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeBudget(formats); err != nil {
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

func (m *ActivityEntity) validateBillable(formats strfmt.Registry) error {

	if err := validate.Required("billable", "body", m.Billable); err != nil {
		return err
	}

	return nil
}

func (m *ActivityEntity) validateBudget(formats strfmt.Registry) error {

	if err := validate.Required("budget", "body", m.Budget); err != nil {
		return err
	}

	if err := validate.Minimum("budget", "body", float64(*m.Budget), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("budget", "body", float64(*m.Budget), 9e+11, false); err != nil {
		return err
	}

	return nil
}

func (m *ActivityEntity) validateMetaFields(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaFields) { // not required
		return nil
	}

	for i := 0; i < len(m.MetaFields); i++ {
		if swag.IsZero(m.MetaFields[i]) { // not required
			continue
		}

		if m.MetaFields[i] != nil {
			if err := m.MetaFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metaFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metaFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActivityEntity) validateName(formats strfmt.Registry) error {

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

func (m *ActivityEntity) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActivityEntity) validateTimeBudget(formats strfmt.Registry) error {

	if err := validate.Required("timeBudget", "body", m.TimeBudget); err != nil {
		return err
	}

	if err := validate.MinimumInt("timeBudget", "body", *m.TimeBudget, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeBudget", "body", *m.TimeBudget, 2.1456e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ActivityEntity) validateVisible(formats strfmt.Registry) error {

	if err := validate.Required("visible", "body", m.Visible); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this activity entity based on the context it is used
func (m *ActivityEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetaFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityEntity) contextValidateMetaFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetaFields); i++ {

		if m.MetaFields[i] != nil {
			if err := m.MetaFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metaFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metaFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActivityEntity) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {
			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityEntity) UnmarshalBinary(b []byte) error {
	var res ActivityEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
