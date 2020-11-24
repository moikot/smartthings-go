// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rules API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rules API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRule(params *CreateRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRuleOK, error)

	DeleteRule(params *DeleteRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRuleOK, error)

	ExecuteRule(params *ExecuteRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ExecuteRuleOK, error)

	GetRule(params *GetRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuleOK, error)

	ListRules(params *ListRulesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRulesOK, error)

	UpdateRule(params *UpdateRuleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRule creates a rule

  Create a rule for the location and token principal
*/
func (a *Client) CreateRule(params *CreateRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRule",
		Method:             "POST",
		PathPattern:        "/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteRule deletes a rule

  Delete a rule
*/
func (a *Client) DeleteRule(params *DeleteRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRule",
		Method:             "DELETE",
		PathPattern:        "/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ExecuteRule executes a rule

  Trigger Rule execution given a rule ID
*/
func (a *Client) ExecuteRule(params *ExecuteRuleParams, authInfo runtime.ClientAuthInfoWriter) (*ExecuteRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "executeRule",
		Method:             "POST",
		PathPattern:        "/rules/execute/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecuteRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExecuteRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExecuteRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetRule gets a rule

  Get a rule
*/
func (a *Client) GetRule(params *GetRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRule",
		Method:             "GET",
		PathPattern:        "/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListRules rules list

  List of rules for the location for the given token principal
*/
func (a *Client) ListRules(params *ListRulesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRules",
		Method:             "GET",
		PathPattern:        "/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateRule updates a rule

  Update a rule
*/
func (a *Client) UpdateRule(params *UpdateRuleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRule",
		Method:             "PUT",
		PathPattern:        "/rules/{ruleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
