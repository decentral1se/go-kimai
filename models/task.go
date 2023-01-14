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

// Task task
//
// swagger:model Task
type Task struct {

	// active timesheets
	ActiveTimesheets []*TimesheetEntityExpanded `json:"activeTimesheets"`

	// activity
	// Required: true
	Activity *ActivityExpanded `json:"activity"`

	// description
	Description string `json:"description,omitempty"`

	// end
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// Estimated duration for the task in seconds or null if no estimation was added
	// Minimum: 0
	Estimation *int64 `json:"estimation,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// project
	// Required: true
	Project *ProjectExpanded `json:"project"`

	// status
	// Required: true
	// Max Length: 20
	Status *string `json:"status"`

	// tags
	Tags []string `json:"tags"`

	// team
	Team *Team `json:"team,omitempty"`

	// title
	// Required: true
	// Max Length: 100
	// Min Length: 2
	Title *string `json:"title"`

	// todo
	Todo string `json:"todo,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveTimesheets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
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

func (m *Task) validateActiveTimesheets(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveTimesheets) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveTimesheets); i++ {
		if swag.IsZero(m.ActiveTimesheets[i]) { // not required
			continue
		}

		if m.ActiveTimesheets[i] != nil {
			if err := m.ActiveTimesheets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activeTimesheets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activeTimesheets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) validateActivity(formats strfmt.Registry) error {

	if err := validate.Required("activity", "body", m.Activity); err != nil {
		return err
	}

	if m.Activity != nil {
		if err := m.Activity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateEstimation(formats strfmt.Registry) error {
	if swag.IsZero(m.Estimation) { // not required
		return nil
	}

	if err := validate.MinimumInt("estimation", "body", *m.Estimation, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
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

func (m *Task) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.MaxLength("status", "body", *m.Status, 20); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if m.Team != nil {
		if err := m.Team.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *Task) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 100); err != nil {
		return err
	}

	return nil
}

func (m *Task) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this task based on the context it is used
func (m *Task) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveTimesheets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) contextValidateActiveTimesheets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActiveTimesheets); i++ {

		if m.ActiveTimesheets[i] != nil {
			if err := m.ActiveTimesheets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activeTimesheets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activeTimesheets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Task) contextValidateActivity(ctx context.Context, formats strfmt.Registry) error {

	if m.Activity != nil {
		if err := m.Activity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("activity")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Task) contextValidateTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.Team != nil {
		if err := m.Team.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("team")
			}
			return err
		}
	}

	return nil
}

func (m *Task) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Task) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Task) UnmarshalBinary(b []byte) error {
	var res Task
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
