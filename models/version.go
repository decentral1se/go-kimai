// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Version version
//
// swagger:model Version
type Version struct {

	// Candidate: either "prod" or "dev"
	Candidate string `json:"candidate,omitempty"`

	// A full copyright notice
	Copyright string `json:"copyright,omitempty"`

	// The version name
	Name string `json:"name,omitempty"`

	// Full version including status, eg: "1.9-prod"
	Semver string `json:"semver,omitempty"`

	// Kimai Version, eg. "1.14"
	Version string `json:"version,omitempty"`

	// Kimai Version as integer, eg. 11400
	//
	// Follows the same logic as PHP_VERSION_ID, see https://www.php.net/manual/de/function.phpversion.php
	VersionID int64 `json:"versionId,omitempty"`
}

// Validate validates this version
func (m *Version) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version based on context it is used
func (m *Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Version) UnmarshalBinary(b []byte) error {
	var res Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}