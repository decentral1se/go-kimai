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
	"github.com/go-openapi/swag"
)

// NewDeleteAPITagsIDParams creates a new DeleteAPITagsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPITagsIDParams() *DeleteAPITagsIDParams {
	return &DeleteAPITagsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPITagsIDParamsWithTimeout creates a new DeleteAPITagsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAPITagsIDParamsWithTimeout(timeout time.Duration) *DeleteAPITagsIDParams {
	return &DeleteAPITagsIDParams{
		timeout: timeout,
	}
}

// NewDeleteAPITagsIDParamsWithContext creates a new DeleteAPITagsIDParams object
// with the ability to set a context for a request.
func NewDeleteAPITagsIDParamsWithContext(ctx context.Context) *DeleteAPITagsIDParams {
	return &DeleteAPITagsIDParams{
		Context: ctx,
	}
}

// NewDeleteAPITagsIDParamsWithHTTPClient creates a new DeleteAPITagsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPITagsIDParamsWithHTTPClient(client *http.Client) *DeleteAPITagsIDParams {
	return &DeleteAPITagsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAPITagsIDParams contains all the parameters to send to the API endpoint

	for the delete API tags ID operation.

	Typically these are written to a http.Request.
*/
type DeleteAPITagsIDParams struct {

	/* ID.

	   Tag ID to delete
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API tags ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITagsIDParams) WithDefaults() *DeleteAPITagsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API tags ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITagsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API tags ID params
func (o *DeleteAPITagsIDParams) WithTimeout(timeout time.Duration) *DeleteAPITagsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API tags ID params
func (o *DeleteAPITagsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API tags ID params
func (o *DeleteAPITagsIDParams) WithContext(ctx context.Context) *DeleteAPITagsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API tags ID params
func (o *DeleteAPITagsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API tags ID params
func (o *DeleteAPITagsIDParams) WithHTTPClient(client *http.Client) *DeleteAPITagsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API tags ID params
func (o *DeleteAPITagsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete API tags ID params
func (o *DeleteAPITagsIDParams) WithID(id int64) *DeleteAPITagsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete API tags ID params
func (o *DeleteAPITagsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPITagsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
