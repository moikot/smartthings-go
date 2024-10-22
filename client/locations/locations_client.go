// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLocation(params *CreateLocationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLocationOK, error)

	DeleteLocation(params *DeleteLocationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLocationOK, error)

	GetLocation(params *GetLocationParams, authInfo runtime.ClientAuthInfoWriter) (*GetLocationOK, error)

	ListLocations(params *ListLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListLocationsOK, error)

	UpdateLocation(params *UpdateLocationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLocationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateLocation creates a location

  Create a Location for a user. We will try and create the Location geographically close to the country code provided in the request body. If we do not support Location creation in the requested country code, then the API will return a 422 error response with an error code of UnsupportedGeoRegionError.

*/
func (a *Client) CreateLocation(params *CreateLocationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLocation",
		Method:             "POST",
		PathPattern:        "/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteLocation deletes a location

  Delete a Location from a user's account.
*/
func (a *Client) DeleteLocation(params *DeleteLocationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLocation",
		Method:             "DELETE",
		PathPattern:        "/locations/{locationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLocation gets a location

  Get a specific Location from a user's account.
*/
func (a *Client) GetLocation(params *GetLocationParams, authInfo runtime.ClientAuthInfoWriter) (*GetLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLocation",
		Method:             "GET",
		PathPattern:        "/locations/{locationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListLocations lists locations

  List all Locations currently available in a user account.
*/
func (a *Client) ListLocations(params *ListLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLocationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listLocations",
		Method:             "GET",
		PathPattern:        "/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListLocationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListLocationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateLocation updates a location

  All the fields in the request body are optional. Only the specified fields will be updated.
*/
func (a *Client) UpdateLocation(params *UpdateLocationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLocation",
		Method:             "PUT",
		PathPattern:        "/locations/{locationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateLocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateLocationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
