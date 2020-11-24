// Code generated by go-swagger; DO NOT EDIT.

package presentations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new presentations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for presentations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCustomCapabilityPresentation(params *CreateCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCustomCapabilityPresentationOK, error)

	CreateDeviceConfiguration(params *CreateDeviceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceConfigurationOK, error)

	GenerateDeviceConfig(params *GenerateDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateDeviceConfigOK, error)

	GetCustomCapabilityPresentation(params *GetCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomCapabilityPresentationOK, error)

	GetDeviceConfiguration(params *GetDeviceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceConfigurationOK, error)

	GetDevicePresentation(params *GetDevicePresentationParams, authInfo runtime.ClientAuthInfoWriter) (*GetDevicePresentationOK, error)

	UpdateCustomCapabilityPresentation(params *UpdateCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCustomCapabilityPresentationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCustomCapabilityPresentation creates a capability presentation

  Create a capability presentation.<br /><br />
**Note:** This API functionality is in BETA

*/
func (a *Client) CreateCustomCapabilityPresentation(params *CreateCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCustomCapabilityPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCustomCapabilityPresentationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCustomCapabilityPresentation",
		Method:             "POST",
		PathPattern:        "/capabilities/{capabilityId}/{capabilityVersion}/presentation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateCustomCapabilityPresentationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCustomCapabilityPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCustomCapabilityPresentationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateDeviceConfiguration creates a device configuration

  Make an idempotent call to either create or get a device configuration based on the structure of the provided payload
Note: This API functionality is in BETA

*/
func (a *Client) CreateDeviceConfiguration(params *CreateDeviceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDeviceConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDeviceConfiguration",
		Method:             "POST",
		PathPattern:        "/presentation/deviceconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeviceConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDeviceConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GenerateDeviceConfig generates device config from a device profile or d t h

  Examines the profile of the device and constructs a default device configuration.
Note: This API functionality is in BETA

*/
func (a *Client) GenerateDeviceConfig(params *GenerateDeviceConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GenerateDeviceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGenerateDeviceConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "generateDeviceConfig",
		Method:             "GET",
		PathPattern:        "/presentation/types/{typeIntegrationId}/deviceconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GenerateDeviceConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GenerateDeviceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GenerateDeviceConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCustomCapabilityPresentation gets a capability presentation by id and version

  Get a capability presentation by id and version.<br /><br />
**Note:** This API functionality is in BETA

*/
func (a *Client) GetCustomCapabilityPresentation(params *GetCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*GetCustomCapabilityPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomCapabilityPresentationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomCapabilityPresentation",
		Method:             "GET",
		PathPattern:        "/capabilities/{capabilityId}/{capabilityVersion}/presentation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCustomCapabilityPresentationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomCapabilityPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCustomCapabilityPresentationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDeviceConfiguration gets a device configuration

  Get a device configuration.
Note: This API functionality is in BETA

*/
func (a *Client) GetDeviceConfiguration(params *GetDeviceConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*GetDeviceConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeviceConfiguration",
		Method:             "GET",
		PathPattern:        "/presentation/deviceconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDeviceConfigurationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDevicePresentation gets a device presentation

  Get a device presentation. If Manufacturer Name is omitted we will assume the default - `SmartThingsCommunity`.
This endpoint currently supports automatic generation of UI-metadata only for those presentations created
through the presentation APIs, and not custom UI-metadata created through the legacy workflow.
Note: This API functionality is in BETA

*/
func (a *Client) GetDevicePresentation(params *GetDevicePresentationParams, authInfo runtime.ClientAuthInfoWriter) (*GetDevicePresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDevicePresentationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDevicePresentation",
		Method:             "GET",
		PathPattern:        "/presentation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDevicePresentationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDevicePresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDevicePresentationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateCustomCapabilityPresentation updates a capability presentation

  Update a capability presentation.<br /><br />
**Note:** This API functionality is in BETA

*/
func (a *Client) UpdateCustomCapabilityPresentation(params *UpdateCustomCapabilityPresentationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCustomCapabilityPresentationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCustomCapabilityPresentationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCustomCapabilityPresentation",
		Method:             "PUT",
		PathPattern:        "/capabilities/{capabilityId}/{capabilityVersion}/presentation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateCustomCapabilityPresentationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCustomCapabilityPresentationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCustomCapabilityPresentationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}