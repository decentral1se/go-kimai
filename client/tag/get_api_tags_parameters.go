// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewGetAPITagsParams creates a new GetAPITagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPITagsParams() *GetAPITagsParams {
	return &GetAPITagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPITagsParamsWithTimeout creates a new GetAPITagsParams object
// with the ability to set a timeout on a request.
func NewGetAPITagsParamsWithTimeout(timeout time.Duration) *GetAPITagsParams {
	return &GetAPITagsParams{
		timeout: timeout,
	}
}

// NewGetAPITagsParamsWithContext creates a new GetAPITagsParams object
// with the ability to set a context for a request.
func NewGetAPITagsParamsWithContext(ctx context.Context) *GetAPITagsParams {
	return &GetAPITagsParams{
		Context: ctx,
	}
}

// NewGetAPITagsParamsWithHTTPClient creates a new GetAPITagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPITagsParamsWithHTTPClient(client *http.Client) *GetAPITagsParams {
	return &GetAPITagsParams{
		HTTPClient: client,
	}
}

/*
GetAPITagsParams contains all the parameters to send to the API endpoint

	for the get API tags operation.

	Typically these are written to a http.Request.
*/
type GetAPITagsParams struct {

	/* Name.

	   Search term to filter tag list
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITagsParams) WithDefaults() *GetAPITagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPITagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API tags params
func (o *GetAPITagsParams) WithTimeout(timeout time.Duration) *GetAPITagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API tags params
func (o *GetAPITagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API tags params
func (o *GetAPITagsParams) WithContext(ctx context.Context) *GetAPITagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API tags params
func (o *GetAPITagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API tags params
func (o *GetAPITagsParams) WithHTTPClient(client *http.Client) *GetAPITagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API tags params
func (o *GetAPITagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get API tags params
func (o *GetAPITagsParams) WithName(name string) *GetAPITagsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API tags params
func (o *GetAPITagsParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPITagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName

	if err := r.SetQueryParam("name", qName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
