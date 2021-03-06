// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTest(params *GetTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestOK, error)

	PostTest(params *PostTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTestCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTest get test API
*/
func (a *Client) GetTest(params *GetTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTest",
		Method:             "GET",
		PathPattern:        "/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTestReader{formats: a.formats},
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
	success, ok := result.(*GetTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTest post test API
*/
func (a *Client) PostTest(params *PostTestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostTestCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTest",
		Method:             "POST",
		PathPattern:        "/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostTestReader{formats: a.formats},
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
	success, ok := result.(*PostTestCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
