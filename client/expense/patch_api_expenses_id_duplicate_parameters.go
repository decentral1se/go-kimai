// Code generated by go-swagger; DO NOT EDIT.

package expense

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

// NewPatchAPIExpensesIDDuplicateParams creates a new PatchAPIExpensesIDDuplicateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchAPIExpensesIDDuplicateParams() *PatchAPIExpensesIDDuplicateParams {
	return &PatchAPIExpensesIDDuplicateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIExpensesIDDuplicateParamsWithTimeout creates a new PatchAPIExpensesIDDuplicateParams object
// with the ability to set a timeout on a request.
func NewPatchAPIExpensesIDDuplicateParamsWithTimeout(timeout time.Duration) *PatchAPIExpensesIDDuplicateParams {
	return &PatchAPIExpensesIDDuplicateParams{
		timeout: timeout,
	}
}

// NewPatchAPIExpensesIDDuplicateParamsWithContext creates a new PatchAPIExpensesIDDuplicateParams object
// with the ability to set a context for a request.
func NewPatchAPIExpensesIDDuplicateParamsWithContext(ctx context.Context) *PatchAPIExpensesIDDuplicateParams {
	return &PatchAPIExpensesIDDuplicateParams{
		Context: ctx,
	}
}

// NewPatchAPIExpensesIDDuplicateParamsWithHTTPClient creates a new PatchAPIExpensesIDDuplicateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchAPIExpensesIDDuplicateParamsWithHTTPClient(client *http.Client) *PatchAPIExpensesIDDuplicateParams {
	return &PatchAPIExpensesIDDuplicateParams{
		HTTPClient: client,
	}
}

/*
PatchAPIExpensesIDDuplicateParams contains all the parameters to send to the API endpoint

	for the patch API expenses ID duplicate operation.

	Typically these are written to a http.Request.
*/
type PatchAPIExpensesIDDuplicateParams struct {

	/* ID.

	   Expense record ID to duplicate
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch API expenses ID duplicate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIExpensesIDDuplicateParams) WithDefaults() *PatchAPIExpensesIDDuplicateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch API expenses ID duplicate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchAPIExpensesIDDuplicateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) WithTimeout(timeout time.Duration) *PatchAPIExpensesIDDuplicateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) WithContext(ctx context.Context) *PatchAPIExpensesIDDuplicateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) WithHTTPClient(client *http.Client) *PatchAPIExpensesIDDuplicateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) WithID(id int64) *PatchAPIExpensesIDDuplicateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch API expenses ID duplicate params
func (o *PatchAPIExpensesIDDuplicateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIExpensesIDDuplicateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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