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
)

// NewGetAPITimesheetsActiveParams creates a new GetAPITimesheetsActiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITimesheetsActiveParams() *GetAPITimesheetsActiveParams {
	return &GetAPITimesheetsActiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITimesheetsActiveParamsWithTimeout creates a new GetAPITimesheetsActiveParams object
// with the ability to set a timeout on a request.
func NewGetAPITimesheetsActiveParamsWithTimeout(timeout time.Duration) *GetAPITimesheetsActiveParams {
	return &GetAPITimesheetsActiveParams{
		timeout: timeout,
	}
}

// NewGetAPITimesheetsActiveParamsWithContext creates a new GetAPITimesheetsActiveParams object
// with the ability to set a context for a request.
func NewGetAPITimesheetsActiveParamsWithContext(ctx context.Context) *GetAPITimesheetsActiveParams {
	return &GetAPITimesheetsActiveParams{
		Context: ctx,
	}
}

// NewGetAPITimesheetsActiveParamsWithHTTPClient creates a new GetAPITimesheetsActiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITimesheetsActiveParamsWithHTTPClient(client *http.Client) *GetAPITimesheetsActiveParams {
	return &GetAPITimesheetsActiveParams{
		HTTPClient: client,
	}
}

/*
GetAPITimesheetsActiveParams contains all the parameters to send to the API endpoint

	for the get API timesheets active operation.

	Typically these are written to a http.Request.
*/
type GetAPITimesheetsActiveParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API timesheets active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITimesheetsActiveParams) WithDefaults() *GetAPITimesheetsActiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API timesheets active params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITimesheetsActiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) WithTimeout(timeout time.Duration) *GetAPITimesheetsActiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) WithContext(ctx context.Context) *GetAPITimesheetsActiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) WithHTTPClient(client *http.Client) *GetAPITimesheetsActiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API timesheets active params
func (o *GetAPITimesheetsActiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITimesheetsActiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
