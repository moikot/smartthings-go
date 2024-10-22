// Code generated by go-swagger; DO NOT EDIT.

package scenes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scenes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scenes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ExecuteScene(params *ExecuteSceneParams, authInfo runtime.ClientAuthInfoWriter) (*ExecuteSceneOK, error)

	ListScenes(params *ListScenesParams, authInfo runtime.ClientAuthInfoWriter) (*ListScenesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ExecuteScene executes scene

  Execute a Scene by id for the logged in user and given locationId
*/
func (a *Client) ExecuteScene(params *ExecuteSceneParams, authInfo runtime.ClientAuthInfoWriter) (*ExecuteSceneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteSceneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "executeScene",
		Method:             "POST",
		PathPattern:        "/scenes/{sceneId}/execute",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteSceneReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExecuteSceneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExecuteSceneDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListScenes lists scenes

  Fetch a list of Scenes for the logged in user filtered by locationIds. If no locationId is sent, return scenes for all available locations
*/
func (a *Client) ListScenes(params *ListScenesParams, authInfo runtime.ClientAuthInfoWriter) (*ListScenesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScenesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listScenes",
		Method:             "GET",
		PathPattern:        "/scenes",
		ProducesMediaTypes: []string{"application/vnd.smartthings+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListScenesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListScenesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListScenesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
