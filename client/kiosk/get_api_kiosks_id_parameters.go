// Code generated by go-swagger; DO NOT EDIT.

package kiosk

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

// NewGetAPIKiosksIDParams creates a new GetAPIKiosksIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPIKiosksIDParams() *GetAPIKiosksIDParams {
	return &GetAPIKiosksIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIKiosksIDParamsWithTimeout creates a new GetAPIKiosksIDParams object
// with the ability to set a timeout on a request.
func NewGetAPIKiosksIDParamsWithTimeout(timeout time.Duration) *GetAPIKiosksIDParams {
	return &GetAPIKiosksIDParams{
		timeout: timeout,
	}
}

// NewGetAPIKiosksIDParamsWithContext creates a new GetAPIKiosksIDParams object
// with the ability to set a context for a request.
func NewGetAPIKiosksIDParamsWithContext(ctx context.Context) *GetAPIKiosksIDParams {
	return &GetAPIKiosksIDParams{
		Context: ctx,
	}
}

// NewGetAPIKiosksIDParamsWithHTTPClient creates a new GetAPIKiosksIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPIKiosksIDParamsWithHTTPClient(client *http.Client) *GetAPIKiosksIDParams {
	return &GetAPIKiosksIDParams{
		HTTPClient: client,
	}
}

/*
GetAPIKiosksIDParams contains all the parameters to send to the API endpoint

	for the get API kiosks ID operation.

	Typically these are written to a http.Request.
*/
type GetAPIKiosksIDParams struct {

	/* ID.

	   User ID to fetch
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API kiosks ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKiosksIDParams) WithDefaults() *GetAPIKiosksIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API kiosks ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPIKiosksIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) WithTimeout(timeout time.Duration) *GetAPIKiosksIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) WithContext(ctx context.Context) *GetAPIKiosksIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) WithHTTPClient(client *http.Client) *GetAPIKiosksIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) WithID(id int64) *GetAPIKiosksIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get API kiosks ID params
func (o *GetAPIKiosksIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIKiosksIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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