// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPICustomersIDRatesRateID(params *DeleteAPICustomersIDRatesRateIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPICustomersIDRatesRateIDNoContent, error)

	GetAPICustomers(params *GetAPICustomersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersOK, error)

	GetAPICustomersID(params *GetAPICustomersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersIDOK, error)

	GetAPICustomersIDRates(params *GetAPICustomersIDRatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersIDRatesOK, error)

	PatchAPICustomersID(params *PatchAPICustomersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPICustomersIDOK, error)

	PatchAPICustomersIDMeta(params *PatchAPICustomersIDMetaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPICustomersIDMetaOK, error)

	PostAPICustomers(params *PostAPICustomersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPICustomersOK, error)

	PostAPICustomersIDRates(params *PostAPICustomersIDRatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPICustomersIDRatesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPICustomersIDRatesRateID deletes one rate for an customer
*/
func (a *Client) DeleteAPICustomersIDRatesRateID(params *DeleteAPICustomersIDRatesRateIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPICustomersIDRatesRateIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPICustomersIDRatesRateIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPICustomersIDRatesRateID",
		Method:             "DELETE",
		PathPattern:        "/api/customers/{id}/rates/{rateId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPICustomersIDRatesRateIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPICustomersIDRatesRateIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPICustomersIDRatesRateID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPICustomers returns a collection of customers
*/
func (a *Client) GetAPICustomers(params *GetAPICustomersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPICustomersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPICustomers",
		Method:             "GET",
		PathPattern:        "/api/customers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPICustomersReader{formats: a.formats},
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
	success, ok := result.(*GetAPICustomersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPICustomers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPICustomersID returns one customer
*/
func (a *Client) GetAPICustomersID(params *GetAPICustomersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPICustomersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPICustomersID",
		Method:             "GET",
		PathPattern:        "/api/customers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPICustomersIDReader{formats: a.formats},
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
	success, ok := result.(*GetAPICustomersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPICustomersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPICustomersIDRates returns a collection of all rates for one customer
*/
func (a *Client) GetAPICustomersIDRates(params *GetAPICustomersIDRatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPICustomersIDRatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPICustomersIDRatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPICustomersIDRates",
		Method:             "GET",
		PathPattern:        "/api/customers/{id}/rates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPICustomersIDRatesReader{formats: a.formats},
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
	success, ok := result.(*GetAPICustomersIDRatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPICustomersIDRates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPICustomersID updates an existing customer

Update an existing customer, you can pass all or just a subset of all attributes
*/
func (a *Client) PatchAPICustomersID(params *PatchAPICustomersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPICustomersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPICustomersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPICustomersID",
		Method:             "PATCH",
		PathPattern:        "/api/customers/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPICustomersIDReader{formats: a.formats},
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
	success, ok := result.(*PatchAPICustomersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPICustomersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPICustomersIDMeta sets the value of a meta field for an existing customer
*/
func (a *Client) PatchAPICustomersIDMeta(params *PatchAPICustomersIDMetaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPICustomersIDMetaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPICustomersIDMetaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPICustomersIDMeta",
		Method:             "PATCH",
		PathPattern:        "/api/customers/{id}/meta",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPICustomersIDMetaReader{formats: a.formats},
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
	success, ok := result.(*PatchAPICustomersIDMetaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPICustomersIDMeta: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPICustomers creates a new customer

Creates a new customer and returns it afterwards
*/
func (a *Client) PostAPICustomers(params *PostAPICustomersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPICustomersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPICustomersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPICustomers",
		Method:             "POST",
		PathPattern:        "/api/customers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPICustomersReader{formats: a.formats},
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
	success, ok := result.(*PostAPICustomersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPICustomers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPICustomersIDRates adds a new rate to a customer
*/
func (a *Client) PostAPICustomersIDRates(params *PostAPICustomersIDRatesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPICustomersIDRatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPICustomersIDRatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPICustomersIDRates",
		Method:             "POST",
		PathPattern:        "/api/customers/{id}/rates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPICustomersIDRatesReader{formats: a.formats},
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
	success, ok := result.(*PostAPICustomersIDRatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPICustomersIDRates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
