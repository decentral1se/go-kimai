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

	"decentral1se/go-kimai/models"
)

// NewPostAPITasksParams creates a new PostAPITasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPITasksParams() *PostAPITasksParams {
	return &PostAPITasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPITasksParamsWithTimeout creates a new PostAPITasksParams object
// with the ability to set a timeout on a request.
func NewPostAPITasksParamsWithTimeout(timeout time.Duration) *PostAPITasksParams {
	return &PostAPITasksParams{
		timeout: timeout,
	}
}

// NewPostAPITasksParamsWithContext creates a new PostAPITasksParams object
// with the ability to set a context for a request.
func NewPostAPITasksParamsWithContext(ctx context.Context) *PostAPITasksParams {
	return &PostAPITasksParams{
		Context: ctx,
	}
}

// NewPostAPITasksParamsWithHTTPClient creates a new PostAPITasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPITasksParamsWithHTTPClient(client *http.Client) *PostAPITasksParams {
	return &PostAPITasksParams{
		HTTPClient: client,
	}
}

/*
PostAPITasksParams contains all the parameters to send to the API endpoint

	for the post API tasks operation.

	Typically these are written to a http.Request.
*/
type PostAPITasksParams struct {

	// Body.
	Body *models.TaskEditForm

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post API tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPITasksParams) WithDefaults() *PostAPITasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPITasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API tasks params
func (o *PostAPITasksParams) WithTimeout(timeout time.Duration) *PostAPITasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API tasks params
func (o *PostAPITasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API tasks params
func (o *PostAPITasksParams) WithContext(ctx context.Context) *PostAPITasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API tasks params
func (o *PostAPITasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API tasks params
func (o *PostAPITasksParams) WithHTTPClient(client *http.Client) *PostAPITasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API tasks params
func (o *PostAPITasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post API tasks params
func (o *PostAPITasksParams) WithBody(body *models.TaskEditForm) *PostAPITasksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API tasks params
func (o *PostAPITasksParams) SetBody(body *models.TaskEditForm) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPITasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
