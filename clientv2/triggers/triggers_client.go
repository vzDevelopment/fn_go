// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new triggers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for triggers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateTrigger creates a trigger

Creates a Trigger.
*/
func (a *Client) CreateTrigger(params *CreateTriggerParams) (*CreateTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateTrigger",
		Method:             "POST",
		PathPattern:        "/triggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateTriggerOK), nil

}

/*
DeleteTrigger deletes the trigger

Deletes the Trigger.
*/
func (a *Client) DeleteTrigger(params *DeleteTriggerParams) (*DeleteTriggerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteTrigger",
		Method:             "DELETE",
		PathPattern:        "/triggers/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTriggerNoContent), nil

}

/*
GetTrigger gets trigger by name

Gets a Trigger by Name.
*/
func (a *Client) GetTrigger(params *GetTriggerParams) (*GetTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTrigger",
		Method:             "GET",
		PathPattern:        "/triggers/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTriggerOK), nil

}

/*
ListTriggers lists triggers associated with app

This will list all Triggers for a particular Application, returned in name alphabetical order.
*/
func (a *Client) ListTriggers(params *ListTriggersParams) (*ListTriggersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTriggersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTriggers",
		Method:             "GET",
		PathPattern:        "/triggers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTriggersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTriggersOK), nil

}

/*
UpdateTrigger updates a trigger

Updates a Trigger.
*/
func (a *Client) UpdateTrigger(params *UpdateTriggerParams) (*UpdateTriggerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTriggerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateTrigger",
		Method:             "PUT",
		PathPattern:        "/triggers/{triggerID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateTriggerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTriggerOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
