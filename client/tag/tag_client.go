// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tag API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPITagsID(params *DeleteAPITagsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPITagsIDNoContent, error)

	GetAPITags(params *GetAPITagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITagsOK, error)

	PostAPITags(params *PostAPITagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPITagsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteAPITagsID deletes a tag
*/
func (a *Client) DeleteAPITagsID(params *DeleteAPITagsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPITagsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPITagsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteAPITagsID",
		Method:             "DELETE",
		PathPattern:        "/api/tags/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPITagsIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteAPITagsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPITagsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAPITags fetches all existing tags
*/
func (a *Client) GetAPITags(params *GetAPITagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPITagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPITagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAPITags",
		Method:             "GET",
		PathPattern:        "/api/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPITagsReader{formats: a.formats},
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
	success, ok := result.(*GetAPITagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPITags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAPITags creates a new tag

Creates a new tag and returns it afterwards
*/
func (a *Client) PostAPITags(params *PostAPITagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAPITagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPITagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAPITags",
		Method:             "POST",
		PathPattern:        "/api/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPITagsReader{formats: a.formats},
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
	success, ok := result.(*PostAPITagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPITags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}