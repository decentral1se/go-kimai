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
	"github.com/go-openapi/swag"
)

// NewDeleteAPITeamsIDParams creates a new DeleteAPITeamsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPITeamsIDParams() *DeleteAPITeamsIDParams {
	return &DeleteAPITeamsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPITeamsIDParamsWithTimeout creates a new DeleteAPITeamsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAPITeamsIDParamsWithTimeout(timeout time.Duration) *DeleteAPITeamsIDParams {
	return &DeleteAPITeamsIDParams{
		timeout: timeout,
	}
}

// NewDeleteAPITeamsIDParamsWithContext creates a new DeleteAPITeamsIDParams object
// with the ability to set a context for a request.
func NewDeleteAPITeamsIDParamsWithContext(ctx context.Context) *DeleteAPITeamsIDParams {
	return &DeleteAPITeamsIDParams{
		Context: ctx,
	}
}

// NewDeleteAPITeamsIDParamsWithHTTPClient creates a new DeleteAPITeamsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPITeamsIDParamsWithHTTPClient(client *http.Client) *DeleteAPITeamsIDParams {
	return &DeleteAPITeamsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAPITeamsIDParams contains all the parameters to send to the API endpoint

	for the delete API teams ID operation.

	Typically these are written to a http.Request.
*/
type DeleteAPITeamsIDParams struct {

	/* ID.

	   Team ID to delete
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API teams ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITeamsIDParams) WithDefaults() *DeleteAPITeamsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API teams ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITeamsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) WithTimeout(timeout time.Duration) *DeleteAPITeamsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) WithContext(ctx context.Context) *DeleteAPITeamsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) WithHTTPClient(client *http.Client) *DeleteAPITeamsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) WithID(id int64) *DeleteAPITeamsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete API teams ID params
func (o *DeleteAPITeamsIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPITeamsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
