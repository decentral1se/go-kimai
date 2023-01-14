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
)

// NewGetAPIProjectsIDParams creates a new GetAPIProjectsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIProjectsIDParams() *GetAPIProjectsIDParams {
	return &GetAPIProjectsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIProjectsIDParamsWithTimeout creates a new GetAPIProjectsIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIProjectsIDParamsWithTimeout(timeout time.Duration) *GetAPIProjectsIDParams {
	return &GetAPIProjectsIDParams{
		timeout: timeout,
	}
}

// NewGetAPIProjectsIDParamsWithContext creates a new GetAPIProjectsIDParams object
// with the ability to set a context for a request.
func NewGetAPIProjectsIDParamsWithContext(ctx context.Context) *GetAPIProjectsIDParams {
	return &GetAPIProjectsIDParams{
		Context: ctx,
	}
}

// NewGetAPIProjectsIDParamsWithHTTPClient creates a new GetAPIProjectsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIProjectsIDParamsWithHTTPClient(client *http.Client) *GetAPIProjectsIDParams {
	return &GetAPIProjectsIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIProjectsIDParams contains all the parameters to send to the API endpoint

	for the get API projects ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIProjectsIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIProjectsIDParams) WithDefaults() *GetAPIProjectsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API projects ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIProjectsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API projects ID params
func (o *GetAPIProjectsIDParams) WithTimeout(timeout time.Duration) *GetAPIProjectsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API projects ID params
func (o *GetAPIProjectsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API projects ID params
func (o *GetAPIProjectsIDParams) WithContext(ctx context.Context) *GetAPIProjectsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API projects ID params
func (o *GetAPIProjectsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API projects ID params
func (o *GetAPIProjectsIDParams) WithHTTPClient(client *http.Client) *GetAPIProjectsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API projects ID params
func (o *GetAPIProjectsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get API projects ID params
func (o *GetAPIProjectsIDParams) WithID(id string) *GetAPIProjectsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API projects ID params
func (o *GetAPIProjectsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIProjectsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
