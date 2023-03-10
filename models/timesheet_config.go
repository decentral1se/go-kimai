// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimesheetConfig timesheet config
//
// swagger:model TimesheetConfig
type TimesheetConfig struct {

	// How many running timesheets a user is allowed to have at the same time
	ActiveEntriesHardLimit int64 `json:"activeEntriesHardLimit,omitempty"`

	// How many running timesheets a user is allowed before a warning is shown
	ActiveEntriesSoftLimit int64 `json:"activeEntriesSoftLimit,omitempty"`

	// Default begin datetime in PHP format
	DefaultBeginTime string `json:"defaultBeginTime,omitempty"`

	// Whether entries for future times are allowed
	IsAllowFutureTimes bool `json:"isAllowFutureTimes,omitempty"`

	// Whether overlapping entries are allowed
	IsAllowOverlapping bool `json:"isAllowOverlapping,omitempty"`

	// The time-tracking mode, see also: https://www.kimai.org/documentation/timesheet.html#tracking-modes
	TrackingMode string `json:"trackingMode,omitempty"`
}

// Validate validates this timesheet config
func (m *TimesheetConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this timesheet config based on context it is used
func (m *TimesheetConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimesheetConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimesheetConfig) UnmarshalBinary(b []byte) error {
	var res TimesheetConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
