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

// NewDeleteAPITeamsIDProjectsProjectIDParams creates a new DeleteAPITeamsIDProjectsProjectIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPITeamsIDProjectsProjectIDParams() *DeleteAPITeamsIDProjectsProjectIDParams {
	return &DeleteAPITeamsIDProjectsProjectIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPITeamsIDProjectsProjectIDParamsWithTimeout creates a new DeleteAPITeamsIDProjectsProjectIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAPITeamsIDProjectsProjectIDParamsWithTimeout(timeout time.Duration) *DeleteAPITeamsIDProjectsProjectIDParams {
	return &DeleteAPITeamsIDProjectsProjectIDParams{
		timeout: timeout,
	}
}

// NewDeleteAPITeamsIDProjectsProjectIDParamsWithContext creates a new DeleteAPITeamsIDProjectsProjectIDParams object
// with the ability to set a context for a request.
func NewDeleteAPITeamsIDProjectsProjectIDParamsWithContext(ctx context.Context) *DeleteAPITeamsIDProjectsProjectIDParams {
	return &DeleteAPITeamsIDProjectsProjectIDParams{
		Context: ctx,
	}
}

// NewDeleteAPITeamsIDProjectsProjectIDParamsWithHTTPClient creates a new DeleteAPITeamsIDProjectsProjectIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPITeamsIDProjectsProjectIDParamsWithHTTPClient(client *http.Client) *DeleteAPITeamsIDProjectsProjectIDParams {
	return &DeleteAPITeamsIDProjectsProjectIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAPITeamsIDProjectsProjectIDParams contains all the parameters to send to the API endpoint

	for the delete API teams ID projects project ID operation.

	Typically these are written to a http.Request.
*/
type DeleteAPITeamsIDProjectsProjectIDParams struct {

	/* ID.

	   The team whose permission will be revoked
	*/
	ID int64

	/* ProjectID.

	   The project to remove (Project ID)
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API teams ID projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithDefaults() *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API teams ID projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithTimeout(timeout time.Duration) *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithContext(ctx context.Context) *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithHTTPClient(client *http.Client) *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithID(id int64) *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetID(id int64) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WithProjectID(projectID int64) *DeleteAPITeamsIDProjectsProjectIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete API teams ID projects project ID params
func (o *DeleteAPITeamsIDProjectsProjectIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPITeamsIDProjectsProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}