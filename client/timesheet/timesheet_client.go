// Code generated by go-swagger; DO NOT EDIT.

package timesheet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new timesheet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for timesheet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPITimesheetsID(params *DeleteAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPITimesheetsIDNoContent, error)

	GetAPITimesheets(params *GetAPITimesheetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsOK, error)

	GetAPITimesheetsActive(params *GetAPITimesheetsActiveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsActiveOK, error)

	GetAPITimesheetsID(params *GetAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsIDOK, error)

	GetAPITimesheetsRecent(params *GetAPITimesheetsRecentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsRecentOK, error)

	PatchAPITimesheetsID(params *PatchAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDOK, error)

	PatchAPITimesheetsIDDuplicate(params *PatchAPITimesheetsIDDuplicateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDDuplicateOK, error)

	PatchAPITimesheetsIDExport(params *PatchAPITimesheetsIDExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDExportOK, error)

	PatchAPITimesheetsIDMeta(params *PatchAPITimesheetsIDMetaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDMetaOK, error)

	PatchAPITimesheetsIDRestart(params *PatchAPITimesheetsIDRestartParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDRestartOK, error)

	PatchAPITimesheetsIDStop(params *PatchAPITimesheetsIDStopParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDStopOK, error)

	PostAPITimesheets(params *PostAPITimesheetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPITimesheetsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPITimesheetsID deletes an existing timesheet record
*/
func (a *Client) DeleteAPITimesheetsID(params *DeleteAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPITimesheetsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPITimesheetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPITimesheetsID",
		Method:             "DELETE",
		PathPattern:        "/api/timesheets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPITimesheetsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPITimesheetsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPITimesheetsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPITimesheets returns a collection of timesheet records
*/
func (a *Client) GetAPITimesheets(params *GetAPITimesheetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITimesheetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITimesheets",
		Method:             "GET",
		PathPattern:        "/api/timesheets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITimesheetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITimesheetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITimesheets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPITimesheetsActive returns the collection of active timesheet records
*/
func (a *Client) GetAPITimesheetsActive(params *GetAPITimesheetsActiveParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITimesheetsActiveParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITimesheetsActive",
		Method:             "GET",
		PathPattern:        "/api/timesheets/active",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITimesheetsActiveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITimesheetsActiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITimesheetsActive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPITimesheetsID returns one timesheet record
*/
func (a *Client) GetAPITimesheetsID(params *GetAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITimesheetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITimesheetsID",
		Method:             "GET",
		PathPattern:        "/api/timesheets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITimesheetsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITimesheetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITimesheetsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPITimesheetsRecent returns the collection of recent user activities
*/
func (a *Client) GetAPITimesheetsRecent(params *GetAPITimesheetsRecentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITimesheetsRecentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITimesheetsRecentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITimesheetsRecent",
		Method:             "GET",
		PathPattern:        "/api/timesheets/recent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITimesheetsRecentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPITimesheetsRecentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITimesheetsRecent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsID updates an existing timesheet record

Update an existing timesheet record, you can pass all or just a subset of the attributes.
*/
func (a *Client) PatchAPITimesheetsID(params *PatchAPITimesheetsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsID",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsIDDuplicate duplicates an existing timesheet record
*/
func (a *Client) PatchAPITimesheetsIDDuplicate(params *PatchAPITimesheetsIDDuplicateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDDuplicateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDDuplicateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsIDDuplicate",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}/duplicate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDDuplicateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDDuplicateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsIDDuplicate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsIDExport switches the export state of a timesheet record to un lock it
*/
func (a *Client) PatchAPITimesheetsIDExport(params *PatchAPITimesheetsIDExportParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDExportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsIDExport",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDExportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDExportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsIDExport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsIDMeta sets the value of a meta field for an existing timesheet
*/
func (a *Client) PatchAPITimesheetsIDMeta(params *PatchAPITimesheetsIDMetaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDMetaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDMetaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsIDMeta",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}/meta",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDMetaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDMetaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsIDMeta: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsIDRestart restarts a previously stopped timesheet record for the current user
*/
func (a *Client) PatchAPITimesheetsIDRestart(params *PatchAPITimesheetsIDRestartParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDRestartOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDRestartParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsIDRestart",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDRestartReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDRestartOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsIDRestart: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPITimesheetsIDStop stops an active timesheet record
*/
func (a *Client) PatchAPITimesheetsIDStop(params *PatchAPITimesheetsIDStopParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPITimesheetsIDStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPITimesheetsIDStopParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPITimesheetsIDStop",
		Method:             "PATCH",
		PathPattern:        "/api/timesheets/{id}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPITimesheetsIDStopReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchAPITimesheetsIDStopOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPITimesheetsIDStop: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPITimesheets creates a new timesheet record

Creates a new timesheet record for the current user and returns it afterwards.
*/
func (a *Client) PostAPITimesheets(params *PostAPITimesheetsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPITimesheetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITimesheetsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPITimesheets",
		Method:             "POST",
		PathPattern:        "/api/timesheets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITimesheetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAPITimesheetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITimesheets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
