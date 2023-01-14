// Code generated by go-swagger; DO NOT EDIT.

package customer

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

	"decentral1se/go-kimai/models"
)

// NewPatchAPICustomersIDParams creates a new PatchAPICustomersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPICustomersIDParams() *PatchAPICustomersIDParams {
	return &PatchAPICustomersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPICustomersIDParamsWithTimeout creates a new PatchAPICustomersIDParams object
// with the ability to set a timeout on a request.
func NewPatchAPICustomersIDParamsWithTimeout(timeout time.Duration) *PatchAPICustomersIDParams {
	return &PatchAPICustomersIDParams{
		timeout: timeout,
	}
}

// NewPatchAPICustomersIDParamsWithContext creates a new PatchAPICustomersIDParams object
// with the ability to set a context for a request.
func NewPatchAPICustomersIDParamsWithContext(ctx context.Context) *PatchAPICustomersIDParams {
	return &PatchAPICustomersIDParams{
		Context: ctx,
	}
}

// NewPatchAPICustomersIDParamsWithHTTPClient creates a new PatchAPICustomersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPICustomersIDParamsWithHTTPClient(client *http.Client) *PatchAPICustomersIDParams {
	return &PatchAPICustomersIDParams{
		HTTPClient: client,
	}
}

/*
PatchAPICustomersIDParams contains all the parameters to send to the API endpoint

	for the patch API customers ID operation.

	Typically these are written to a http.Request.
*/
type PatchAPICustomersIDParams struct {

	// Body.
	Body *models.CustomerEditForm

	/* ID.

	   Customer ID to update
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API customers ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPICustomersIDParams) WithDefaults() *PatchAPICustomersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API customers ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPICustomersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API customers ID params
func (o *PatchAPICustomersIDParams) WithTimeout(timeout time.Duration) *PatchAPICustomersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API customers ID params
func (o *PatchAPICustomersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API customers ID params
func (o *PatchAPICustomersIDParams) WithContext(ctx context.Context) *PatchAPICustomersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API customers ID params
func (o *PatchAPICustomersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API customers ID params
func (o *PatchAPICustomersIDParams) WithHTTPClient(client *http.Client) *PatchAPICustomersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API customers ID params
func (o *PatchAPICustomersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API customers ID params
func (o *PatchAPICustomersIDParams) WithBody(body *models.CustomerEditForm) *PatchAPICustomersIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API customers ID params
func (o *PatchAPICustomersIDParams) SetBody(body *models.CustomerEditForm) {
	o.Body = body
}

// WithID adds the id to the patch API customers ID params
func (o *PatchAPICustomersIDParams) WithID(id int64) *PatchAPICustomersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API customers ID params
func (o *PatchAPICustomersIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPICustomersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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