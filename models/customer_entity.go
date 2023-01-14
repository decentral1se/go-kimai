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

// CustomerEntity customer entity
//
// swagger:model CustomerEntity
type CustomerEntity struct {

	// address
	Address string `json:"address,omitempty"`

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

	// comment
	Comment string `json:"comment,omitempty"`

	// company
	// Max Length: 255
	Company string `json:"company,omitempty"`

	// contact
	// Max Length: 255
	Contact string `json:"contact,omitempty"`

	// country
	// Required: true
	// Max Length: 2
	Country *string `json:"country"`

	// currency
	// Required: true
	// Max Length: 3
	Currency *string `json:"currency"`

	// Customers contact email
	//
	// Limited via RFC to 254 chars
	// Max Length: 254
	Email string `json:"email,omitempty"`

	// fax
	// Max Length: 255
	Fax string `json:"fax,omitempty"`

	// homepage
	// Max Length: 255
	Homepage string `json:"homepage,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Meta fields
	//
	// All visible meta (custom) fields registered with this customer
	MetaFields []*CustomerMeta `json:"metaFields"`

	// mobile
	// Max Length: 255
	Mobile string `json:"mobile,omitempty"`

	// name
	// Required: true
	// Max Length: 150
	// Min Length: 2
	Name *string `json:"name"`

	// number
	// Max Length: 50
	Number string `json:"number,omitempty"`

	// phone
	// Max Length: 255
	Phone string `json:"phone,omitempty"`

	// Teams
	//
	// If no team is assigned, everyone can access the customer
	Teams []*Team `json:"teams"`

	// The time budget in seconds, will be zero if unconfigured.
	// Required: true
	// Maximum: 2.1456e+09
	// Minimum: 0
	TimeBudget *int64 `json:"timeBudget"`

	// Timezone of begin and end
	//
	// Length was determined by a MySQL column via "use mysql;describe time_zone_name;"
	// Required: true
	// Max Length: 64
	Timezone *string `json:"timezone"`

	// vat Id
	// Max Length: 50
	VatID string `json:"vatId,omitempty"`

	// visible
	// Required: true
	Visible *bool `json:"visible"`
}

// Validate validates this customer entity
func (m *CustomerEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBudget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomepage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeBudget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatID(formats); err != nil {
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

func (m *CustomerEntity) validateBillable(formats strfmt.Registry) error {

	if err := validate.Required("billable", "body", m.Billable); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateBudget(formats strfmt.Registry) error {

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

func (m *CustomerEntity) validateCompany(formats strfmt.Registry) error {
	if swag.IsZero(m.Company) { // not required
		return nil
	}

	if err := validate.MaxLength("company", "body", m.Company, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateContact(formats strfmt.Registry) error {
	if swag.IsZero(m.Contact) { // not required
		return nil
	}

	if err := validate.MaxLength("contact", "body", m.Contact, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", *m.Country, 2); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	if err := validate.MaxLength("currency", "body", *m.Currency, 3); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.MaxLength("email", "body", m.Email, 254); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateFax(formats strfmt.Registry) error {
	if swag.IsZero(m.Fax) { // not required
		return nil
	}

	if err := validate.MaxLength("fax", "body", m.Fax, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateHomepage(formats strfmt.Registry) error {
	if swag.IsZero(m.Homepage) { // not required
		return nil
	}

	if err := validate.MaxLength("homepage", "body", m.Homepage, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateMetaFields(formats strfmt.Registry) error {
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

func (m *CustomerEntity) validateMobile(formats strfmt.Registry) error {
	if swag.IsZero(m.Mobile) { // not required
		return nil
	}

	if err := validate.MaxLength("mobile", "body", m.Mobile, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateName(formats strfmt.Registry) error {

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

func (m *CustomerEntity) validateNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MaxLength("number", "body", m.Number, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validatePhone(formats strfmt.Registry) error {
	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MaxLength("phone", "body", m.Phone, 255); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateTeams(formats strfmt.Registry) error {
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

func (m *CustomerEntity) validateTimeBudget(formats strfmt.Registry) error {

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

func (m *CustomerEntity) validateTimezone(formats strfmt.Registry) error {

	if err := validate.Required("timezone", "body", m.Timezone); err != nil {
		return err
	}

	if err := validate.MaxLength("timezone", "body", *m.Timezone, 64); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateVatID(formats strfmt.Registry) error {
	if swag.IsZero(m.VatID) { // not required
		return nil
	}

	if err := validate.MaxLength("vatId", "body", m.VatID, 50); err != nil {
		return err
	}

	return nil
}

func (m *CustomerEntity) validateVisible(formats strfmt.Registry) error {

	if err := validate.Required("visible", "body", m.Visible); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this customer entity based on the context it is used
func (m *CustomerEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *CustomerEntity) contextValidateMetaFields(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CustomerEntity) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CustomerEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerEntity) UnmarshalBinary(b []byte) error {
	var res CustomerEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
