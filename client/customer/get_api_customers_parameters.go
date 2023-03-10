// Code generated by go-swagger; DO NOT EDIT.

package customer

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

// NewGetAPICustomersParams creates a new GetAPICustomersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPICustomersParams() *GetAPICustomersParams {
	return &GetAPICustomersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPICustomersParamsWithTimeout creates a new GetAPICustomersParams object
// with the ability to set a timeout on a request.
func NewGetAPICustomersParamsWithTimeout(timeout time.Duration) *GetAPICustomersParams {
	return &GetAPICustomersParams{
		timeout: timeout,
	}
}

// NewGetAPICustomersParamsWithContext creates a new GetAPICustomersParams object
// with the ability to set a context for a request.
func NewGetAPICustomersParamsWithContext(ctx context.Context) *GetAPICustomersParams {
	return &GetAPICustomersParams{
		Context: ctx,
	}
}

// NewGetAPICustomersParamsWithHTTPClient creates a new GetAPICustomersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPICustomersParamsWithHTTPClient(client *http.Client) *GetAPICustomersParams {
	return &GetAPICustomersParams{
		HTTPClient: client,
	}
}

/*
GetAPICustomersParams contains all the parameters to send to the API endpoint

	for the get API customers operation.

	Typically these are written to a http.Request.
*/
type GetAPICustomersParams struct {

	/* Order.

	   The result order. Allowed values: ASC, DESC (default: ASC)
	*/
	Order string

	/* OrderBy.

	   The field by which results will be ordered. Allowed values: id, name (default: name)
	*/
	OrderBy string

	/* Term.

	   Free search term
	*/
	Term *string

	/* Visible.

	   Visibility status to filter activities (1=visible, 2=hidden, 3=both)
	*/
	Visible string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICustomersParams) WithDefaults() *GetAPICustomersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API customers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPICustomersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API customers params
func (o *GetAPICustomersParams) WithTimeout(timeout time.Duration) *GetAPICustomersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API customers params
func (o *GetAPICustomersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API customers params
func (o *GetAPICustomersParams) WithContext(ctx context.Context) *GetAPICustomersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API customers params
func (o *GetAPICustomersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API customers params
func (o *GetAPICustomersParams) WithHTTPClient(client *http.Client) *GetAPICustomersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API customers params
func (o *GetAPICustomersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrder adds the order to the get API customers params
func (o *GetAPICustomersParams) WithOrder(order string) *GetAPICustomersParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get API customers params
func (o *GetAPICustomersParams) SetOrder(order string) {
	o.Order = order
}

// WithOrderBy adds the orderBy to the get API customers params
func (o *GetAPICustomersParams) WithOrderBy(orderBy string) *GetAPICustomersParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get API customers params
func (o *GetAPICustomersParams) SetOrderBy(orderBy string) {
	o.OrderBy = orderBy
}

// WithTerm adds the term to the get API customers params
func (o *GetAPICustomersParams) WithTerm(term *string) *GetAPICustomersParams {
	o.SetTerm(term)
	return o
}

// SetTerm adds the term to the get API customers params
func (o *GetAPICustomersParams) SetTerm(term *string) {
	o.Term = term
}

// WithVisible adds the visible to the get API customers params
func (o *GetAPICustomersParams) WithVisible(visible string) *GetAPICustomersParams {
	o.SetVisible(visible)
	return o
}

// SetVisible adds the visible to the get API customers params
func (o *GetAPICustomersParams) SetVisible(visible string) {
	o.Visible = visible
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPICustomersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param order
	qrOrder := o.Order
	qOrder := qrOrder

	if err := r.SetQueryParam("order", qOrder); err != nil {
		return err
	}

	// query param orderBy
	qrOrderBy := o.OrderBy
	qOrderBy := qrOrderBy

	if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
		return err
	}

	if o.Term != nil {

		// query param term
		var qrTerm string

		if o.Term != nil {
			qrTerm = *o.Term
		}
		qTerm := qrTerm
		if qTerm != "" {

			if err := r.SetQueryParam("term", qTerm); err != nil {
				return err
			}
		}
	}

	// query param visible
	qrVisible := o.Visible
	qVisible := qrVisible

	if err := r.SetQueryParam("visible", qVisible); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
