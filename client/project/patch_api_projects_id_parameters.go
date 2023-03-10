// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewPatchAPIProjectsIDParams creates a new PatchAPIProjectsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIProjectsIDParams() *PatchAPIProjectsIDParams {
	return &PatchAPIProjectsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIProjectsIDParamsWithTimeout creates a new PatchAPIProjectsIDParams object
// with the ability to set a timeout on a request.
func NewPatchAPIProjectsIDParamsWithTimeout(timeout time.Duration) *PatchAPIProjectsIDParams {
	return &PatchAPIProjectsIDParams{
		timeout: timeout,
	}
}

// NewPatchAPIProjectsIDParamsWithContext creates a new PatchAPIProjectsIDParams object
// with the ability to set a context for a request.
func NewPatchAPIProjectsIDParamsWithContext(ctx context.Context) *PatchAPIProjectsIDParams {
	return &PatchAPIProjectsIDParams{
		Context: ctx,
	}
}

// NewPatchAPIProjectsIDParamsWithHTTPClient creates a new PatchAPIProjectsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIProjectsIDParamsWithHTTPClient(client *http.Client) *PatchAPIProjectsIDParams {
	return &PatchAPIProjectsIDParams{
		HTTPClient: client,
	}
}

/*
PatchAPIProjectsIDParams contains all the parameters to send to the API endpoint

	for the patch API projects ID operation.

	Typically these are written to a http.Request.
*/
type PatchAPIProjectsIDParams struct {

	// Body.
	Body *models.ProjectEditForm

	/* ID.

	   Project ID to update
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIProjectsIDParams) WithDefaults() *PatchAPIProjectsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIProjectsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) WithTimeout(timeout time.Duration) *PatchAPIProjectsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) WithContext(ctx context.Context) *PatchAPIProjectsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) WithHTTPClient(client *http.Client) *PatchAPIProjectsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) WithBody(body *models.ProjectEditForm) *PatchAPIProjectsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) SetBody(body *models.ProjectEditForm) {
	o.Body = body
}

// WithID adds the id to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) WithID(id int64) *PatchAPIProjectsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API projects ID params
func (o *PatchAPIProjectsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIProjectsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
