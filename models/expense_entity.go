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

// ExpenseEntity expense entity
//
// swagger:model ExpenseEntity
type ExpenseEntity struct {

	// activity
	Activity int64 `json:"activity,omitempty"`

	// begin
	// Format: date-time
	Begin strfmt.DateTime `json:"begin,omitempty"`

	// category
	// Required: true
	Category *ExpenseCategory `json:"category"`

	// cost
	// Required: true
	Cost *float32 `json:"cost"`

	// description
	Description string `json:"description,omitempty"`

	// exported
	// Required: true
	Exported *bool `json:"exported"`

	// id
	ID int64 `json:"id,omitempty"`

	// meta fields
	MetaFields []*ExpenseMeta `json:"metaFields"`

	// multiplier
	// Minimum: 0
	Multiplier *float32 `json:"multiplier,omitempty"`

	// project
	Project int64 `json:"project,omitempty"`

	// refundable
	// Required: true
	Refundable *bool `json:"refundable"`

	// total
	Total float32 `json:"total,omitempty"`

	// user
	User int64 `json:"user,omitempty"`
}

// Validate validates this expense entity
func (m *ExpenseEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMultiplier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefundable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpenseEntity) validateBegin(formats strfmt.Registry) error {
	if swag.IsZero(m.Begin) { // not required
		return nil
	}

	if err := validate.FormatOf("begin", "body", "date-time", m.Begin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExpenseEntity) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *ExpenseEntity) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	return nil
}

func (m *ExpenseEntity) validateExported(formats strfmt.Registry) error {

	if err := validate.Required("exported", "body", m.Exported); err != nil {
		return err
	}

	return nil
}

func (m *ExpenseEntity) validateMetaFields(formats strfmt.Registry) error {
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

func (m *ExpenseEntity) validateMultiplier(formats strfmt.Registry) error {
	if swag.IsZero(m.Multiplier) { // not required
		return nil
	}

	if err := validate.Minimum("multiplier", "body", float64(*m.Multiplier), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ExpenseEntity) validateRefundable(formats strfmt.Registry) error {

	if err := validate.Required("refundable", "body", m.Refundable); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this expense entity based on the context it is used
func (m *ExpenseEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetaFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpenseEntity) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {
		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *ExpenseEntity) contextValidateMetaFields(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ExpenseEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpenseEntity) UnmarshalBinary(b []byte) error {
	var res ExpenseEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
