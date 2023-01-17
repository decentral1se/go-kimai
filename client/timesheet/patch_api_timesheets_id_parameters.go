// Code generated by go-swagger; DO NOT EDIT.

package timesheet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/decentral1se/go-kimai/models"
)

// NewPatchAPITimesheetsIDParams creates a new PatchAPITimesheetsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPITimesheetsIDParams() *PatchAPITimesheetsIDParams {
	return &PatchAPITimesheetsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPITimesheetsIDParamsWithTimeout creates a new PatchAPITimesheetsIDParams object
// with the ability to set a timeout on a request.
func NewPatchAPITimesheetsIDParamsWithTimeout(timeout time.Duration) *PatchAPITimesheetsIDParams {
	return &PatchAPITimesheetsIDParams{
		timeout: timeout,
	}
}

// NewPatchAPITimesheetsIDParamsWithContext creates a new PatchAPITimesheetsIDParams object
// with the ability to set a context for a request.
func NewPatchAPITimesheetsIDParamsWithContext(ctx context.Context) *PatchAPITimesheetsIDParams {
	return &PatchAPITimesheetsIDParams{
		Context: ctx,
	}
}

// NewPatchAPITimesheetsIDParamsWithHTTPClient creates a new PatchAPITimesheetsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPITimesheetsIDParamsWithHTTPClient(client *http.Client) *PatchAPITimesheetsIDParams {
	return &PatchAPITimesheetsIDParams{
		HTTPClient: client,
	}
}

/*
PatchAPITimesheetsIDParams contains all the parameters to send to the API endpoint

	for the patch API timesheets ID operation.

	Typically these are written to a http.Request.
*/
type PatchAPITimesheetsIDParams struct {

	// Body.
	Body *models.TimesheetEditForm

	/* ID.

	   Timesheet record ID to update
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API timesheets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPITimesheetsIDParams) WithDefaults() *PatchAPITimesheetsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API timesheets ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPITimesheetsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) WithTimeout(timeout time.Duration) *PatchAPITimesheetsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) WithContext(ctx context.Context) *PatchAPITimesheetsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) WithHTTPClient(client *http.Client) *PatchAPITimesheetsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) WithBody(body *models.TimesheetEditForm) *PatchAPITimesheetsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) SetBody(body *models.TimesheetEditForm) {
	o.Body = body
}

// WithID adds the id to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) WithID(id int64) *PatchAPITimesheetsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API timesheets ID params
func (o *PatchAPITimesheetsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPITimesheetsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
