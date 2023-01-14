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

// UserEntity user entity
//
// swagger:model UserEntity
type UserEntity struct {

	// account number
	// Max Length: 30
	AccountNumber string `json:"accountNumber,omitempty"`

	// The user alias will be displayed in the frontend instead of the username
	// Max Length: 60
	Alias string `json:"alias,omitempty"`

	// URL to the user avatar, will be auto-generated if empty
	// Max Length: 255
	Avatar string `json:"avatar,omitempty"`

	// The assigned color in HTML hex format, eg. #dd1d00
	Color string `json:"color,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Internal ID
	ID int64 `json:"id,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// memberships
	// Required: true
	Memberships *TeamMembership `json:"memberships"`

	// Read-only list of of all visible user preferences.
	Preferences []*UserPreference `json:"preferences"`

	// List of all role names
	Roles []string `json:"roles"`

	// List of all teams, this user is part of
	Teams []*Team `json:"teams"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// An additional title for the user, like the Job position or Department
	// Max Length: 50
	Title string `json:"title,omitempty"`

	// username
	// Required: true
	// Max Length: 60
	// Min Length: 2
	Username *string `json:"username"`
}

// Validate validates this user entity
func (m *UserEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserEntity) validateAccountNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("accountNumber", "body", m.AccountNumber, 30); err != nil {
		return err
	}

	return nil
}

func (m *UserEntity) validateAlias(formats strfmt.Registry) error {
	if swag.IsZero(m.Alias) { // not required
		return nil
	}

	if err := validate.MaxLength("alias", "body", m.Alias, 60); err != nil {
		return err
	}

	return nil
}

func (m *UserEntity) validateAvatar(formats strfmt.Registry) error {
	if swag.IsZero(m.Avatar) { // not required
		return nil
	}

	if err := validate.MaxLength("avatar", "body", m.Avatar, 255); err != nil {
		return err
	}

	return nil
}

func (m *UserEntity) validateMemberships(formats strfmt.Registry) error {

	if err := validate.Required("memberships", "body", m.Memberships); err != nil {
		return err
	}

	if m.Memberships != nil {
		if err := m.Memberships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memberships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memberships")
			}
			return err
		}
	}

	return nil
}

func (m *UserEntity) validatePreferences(formats strfmt.Registry) error {
	if swag.IsZero(m.Preferences) { // not required
		return nil
	}

	for i := 0; i < len(m.Preferences); i++ {
		if swag.IsZero(m.Preferences[i]) { // not required
			continue
		}

		if m.Preferences[i] != nil {
			if err := m.Preferences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserEntity) validateTeams(formats strfmt.Registry) error {
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

func (m *UserEntity) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MaxLength("title", "body", m.Title, 50); err != nil {
		return err
	}

	return nil
}

func (m *UserEntity) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 60); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user entity based on the context it is used
func (m *UserEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMemberships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferences(ctx, formats); err != nil {
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

func (m *UserEntity) contextValidateMemberships(ctx context.Context, formats strfmt.Registry) error {

	if m.Memberships != nil {
		if err := m.Memberships.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memberships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memberships")
			}
			return err
		}
	}

	return nil
}

func (m *UserEntity) contextValidatePreferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Preferences); i++ {

		if m.Preferences[i] != nil {
			if err := m.Preferences[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferences" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("preferences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserEntity) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UserEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEntity) UnmarshalBinary(b []byte) error {
	var res UserEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}