// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewPatchAPIActivitiesIDMetaParams creates a new PatchAPIActivitiesIDMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIActivitiesIDMetaParams() *PatchAPIActivitiesIDMetaParams {
	return &PatchAPIActivitiesIDMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIActivitiesIDMetaParamsWithTimeout creates a new PatchAPIActivitiesIDMetaParams object
// with the ability to set a timeout on a request.
func NewPatchAPIActivitiesIDMetaParamsWithTimeout(timeout time.Duration) *PatchAPIActivitiesIDMetaParams {
	return &PatchAPIActivitiesIDMetaParams{
		timeout: timeout,
	}
}

// NewPatchAPIActivitiesIDMetaParamsWithContext creates a new PatchAPIActivitiesIDMetaParams object
// with the ability to set a context for a request.
func NewPatchAPIActivitiesIDMetaParamsWithContext(ctx context.Context) *PatchAPIActivitiesIDMetaParams {
	return &PatchAPIActivitiesIDMetaParams{
		Context: ctx,
	}
}

// NewPatchAPIActivitiesIDMetaParamsWithHTTPClient creates a new PatchAPIActivitiesIDMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIActivitiesIDMetaParamsWithHTTPClient(client *http.Client) *PatchAPIActivitiesIDMetaParams {
	return &PatchAPIActivitiesIDMetaParams{
		HTTPClient: client,
	}
}

/*
PatchAPIActivitiesIDMetaParams contains all the parameters to send to the API endpoint

	for the patch API activities ID meta operation.

	Typically these are written to a http.Request.
*/
type PatchAPIActivitiesIDMetaParams struct {

	// Body.
	Body PatchAPIActivitiesIDMetaBody

	/* ID.

	   Activity record ID to set the meta-field value for
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API activities ID meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIActivitiesIDMetaParams) WithDefaults() *PatchAPIActivitiesIDMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API activities ID meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIActivitiesIDMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) WithTimeout(timeout time.Duration) *PatchAPIActivitiesIDMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) WithContext(ctx context.Context) *PatchAPIActivitiesIDMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) WithHTTPClient(client *http.Client) *PatchAPIActivitiesIDMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) WithBody(body PatchAPIActivitiesIDMetaBody) *PatchAPIActivitiesIDMetaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) SetBody(body PatchAPIActivitiesIDMetaBody) {
	o.Body = body
}

// WithID adds the id to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) WithID(id int64) *PatchAPIActivitiesIDMetaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API activities ID meta params
func (o *PatchAPIActivitiesIDMetaParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIActivitiesIDMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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