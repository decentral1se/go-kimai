// Code generated by go-swagger; DO NOT EDIT.

package task

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
)

// NewPatchAPITasksIDUnassignParams creates a new PatchAPITasksIDUnassignParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPITasksIDUnassignParams() *PatchAPITasksIDUnassignParams {
	return &PatchAPITasksIDUnassignParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPITasksIDUnassignParamsWithTimeout creates a new PatchAPITasksIDUnassignParams object
// with the ability to set a timeout on a request.
func NewPatchAPITasksIDUnassignParamsWithTimeout(timeout time.Duration) *PatchAPITasksIDUnassignParams {
	return &PatchAPITasksIDUnassignParams{
		timeout: timeout,
	}
}

// NewPatchAPITasksIDUnassignParamsWithContext creates a new PatchAPITasksIDUnassignParams object
// with the ability to set a context for a request.
func NewPatchAPITasksIDUnassignParamsWithContext(ctx context.Context) *PatchAPITasksIDUnassignParams {
	return &PatchAPITasksIDUnassignParams{
		Context: ctx,
	}
}

// NewPatchAPITasksIDUnassignParamsWithHTTPClient creates a new PatchAPITasksIDUnassignParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPITasksIDUnassignParamsWithHTTPClient(client *http.Client) *PatchAPITasksIDUnassignParams {
	return &PatchAPITasksIDUnassignParams{
		HTTPClient: client,
	}
}

/*
PatchAPITasksIDUnassignParams contains all the parameters to send to the API endpoint

	for the patch API tasks ID unassign operation.

	Typically these are written to a http.Request.
*/
type PatchAPITasksIDUnassignParams struct {

	/* ID.

	   Task ID to update
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API tasks ID unassign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPITasksIDUnassignParams) WithDefaults() *PatchAPITasksIDUnassignParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API tasks ID unassign params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPITasksIDUnassignParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) WithTimeout(timeout time.Duration) *PatchAPITasksIDUnassignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) WithContext(ctx context.Context) *PatchAPITasksIDUnassignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) WithHTTPClient(client *http.Client) *PatchAPITasksIDUnassignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) WithID(id int64) *PatchAPITasksIDUnassignParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API tasks ID unassign params
func (o *PatchAPITasksIDUnassignParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPITasksIDUnassignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
