// Code generated by go-swagger; DO NOT EDIT.

package team

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

// NewGetAPITeamsParams creates a new GetAPITeamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITeamsParams() *GetAPITeamsParams {
	return &GetAPITeamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITeamsParamsWithTimeout creates a new GetAPITeamsParams object
// with the ability to set a timeout on a request.
func NewGetAPITeamsParamsWithTimeout(timeout time.Duration) *GetAPITeamsParams {
	return &GetAPITeamsParams{
		timeout: timeout,
	}
}

// NewGetAPITeamsParamsWithContext creates a new GetAPITeamsParams object
// with the ability to set a context for a request.
func NewGetAPITeamsParamsWithContext(ctx context.Context) *GetAPITeamsParams {
	return &GetAPITeamsParams{
		Context: ctx,
	}
}

// NewGetAPITeamsParamsWithHTTPClient creates a new GetAPITeamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITeamsParamsWithHTTPClient(client *http.Client) *GetAPITeamsParams {
	return &GetAPITeamsParams{
		HTTPClient: client,
	}
}

/*
GetAPITeamsParams contains all the parameters to send to the API endpoint

	for the get API teams operation.

	Typically these are written to a http.Request.
*/
type GetAPITeamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITeamsParams) WithDefaults() *GetAPITeamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API teams params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITeamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API teams params
func (o *GetAPITeamsParams) WithTimeout(timeout time.Duration) *GetAPITeamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API teams params
func (o *GetAPITeamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API teams params
func (o *GetAPITeamsParams) WithContext(ctx context.Context) *GetAPITeamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API teams params
func (o *GetAPITeamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API teams params
func (o *GetAPITeamsParams) WithHTTPClient(client *http.Client) *GetAPITeamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API teams params
func (o *GetAPITeamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITeamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
