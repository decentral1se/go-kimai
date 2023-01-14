// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIUsers(params *GetAPIUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersOK, error)

	GetAPIUsersID(params *GetAPIUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersIDOK, error)

	GetAPIUsersMe(params *GetAPIUsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersMeOK, error)

	PatchAPIUsersID(params *PatchAPIUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPIUsersIDOK, error)

	PostAPIUsers(params *PostAPIUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPIUsersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAPIUsers returns the collection of all registered users
*/
func (a *Client) GetAPIUsers(params *GetAPIUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIUsers",
		Method:             "GET",
		PathPattern:        "/api/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIUsersReader{formats: a.formats},
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
	success, ok := result.(*GetAPIUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIUsersID returns one user entity
*/
func (a *Client) GetAPIUsersID(params *GetAPIUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIUsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIUsersID",
		Method:             "GET",
		PathPattern:        "/api/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIUsersIDReader{formats: a.formats},
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
	success, ok := result.(*GetAPIUsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIUsersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPIUsersMe returns the current user entity
*/
func (a *Client) GetAPIUsersMe(params *GetAPIUsersMeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIUsersMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIUsersMeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPIUsersMe",
		Method:             "GET",
		PathPattern:        "/api/users/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIUsersMeReader{formats: a.formats},
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
	success, ok := result.(*GetAPIUsersMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIUsersMe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchAPIUsersID updates an existing user

Update an existing user, you can pass all or just a subset of all attributes (passing roles will replace all existing ones)
*/
func (a *Client) PatchAPIUsersID(params *PatchAPIUsersIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchAPIUsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPIUsersIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchAPIUsersID",
		Method:             "PATCH",
		PathPattern:        "/api/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPIUsersIDReader{formats: a.formats},
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
	success, ok := result.(*PatchAPIUsersIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPIUsersID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPIUsers creates a new user

Creates a new user and returns it afterwards
*/
func (a *Client) PostAPIUsers(params *PostAPIUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPIUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPIUsers",
		Method:             "POST",
		PathPattern:        "/api/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPIUsersReader{formats: a.formats},
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
	success, ok := result.(*PostAPIUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPIUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
