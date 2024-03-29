// Code generated by go-swagger; DO NOT EDIT.

package issues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new issues API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for issues API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IssueAdd(params *IssueAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueAddCreated, error)

	IssueAddFile(params *IssueAddFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueAddFileCreated, error)

	IssueDelete(params *IssueDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueDeleteNoContent, error)

	IssueFileDelete(params *IssueFileDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueFileDeleteNoContent, error)

	IssueFileGet(params *IssueFileGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueFileGetOK, error)

	IssueGet(params *IssueGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueGetOK, error)

	IssueNoteAdd(params *IssueNoteAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteAddCreated, error)

	IssueNoteDelete(params *IssueNoteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteDeleteNoContent, error)

	IssueNoteGet(params *IssueNoteGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
IssueAdd creates an issue
*/
func (a *Client) IssueAdd(params *IssueAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueAddCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueAdd",
		Method:             "POST",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueAddReader{formats: a.formats},
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
	success, ok := result.(*IssueAddCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueAddFile attaches a file to an issue
*/
func (a *Client) IssueAddFile(params *IssueAddFileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueAddFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueAddFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueAddFile",
		Method:             "POST",
		PathPattern:        "/issues/:id/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueAddFileReader{formats: a.formats},
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
	success, ok := result.(*IssueAddFileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueAddFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueDelete deletes an issue
*/
func (a *Client) IssueDelete(params *IssueDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueDelete",
		Method:             "DELETE",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueDeleteReader{formats: a.formats},
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
	success, ok := result.(*IssueDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueFileDelete deletes a file
*/
func (a *Client) IssueFileDelete(params *IssueFileDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueFileDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueFileDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueFileDelete",
		Method:             "DELETE",
		PathPattern:        "/issues/:id/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueFileDeleteReader{formats: a.formats},
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
	success, ok := result.(*IssueFileDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueFileDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueFileGet gets file details
*/
func (a *Client) IssueFileGet(params *IssueFileGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueFileGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueFileGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueFileGet",
		Method:             "GET",
		PathPattern:        "/issues/:id/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueFileGetReader{formats: a.formats},
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
	success, ok := result.(*IssueFileGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueFileGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueGet gets issue details
*/
func (a *Client) IssueGet(params *IssueGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueGet",
		Method:             "GET",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueGetReader{formats: a.formats},
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
	success, ok := result.(*IssueGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueNoteAdd creates a note
*/
func (a *Client) IssueNoteAdd(params *IssueNoteAddParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteAddCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueNoteAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueNoteAdd",
		Method:             "POST",
		PathPattern:        "/issues/:id/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueNoteAddReader{formats: a.formats},
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
	success, ok := result.(*IssueNoteAddCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueNoteAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueNoteDelete deletes a note
*/
func (a *Client) IssueNoteDelete(params *IssueNoteDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueNoteDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueNoteDelete",
		Method:             "DELETE",
		PathPattern:        "/issues/:id/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueNoteDeleteReader{formats: a.formats},
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
	success, ok := result.(*IssueNoteDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueNoteDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueNoteGet gets note details
*/
func (a *Client) IssueNoteGet(params *IssueNoteGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IssueNoteGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueNoteGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "issueNoteGet",
		Method:             "GET",
		PathPattern:        "/issues/:id/notes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueNoteGetReader{formats: a.formats},
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
	success, ok := result.(*IssueNoteGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueNoteGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
