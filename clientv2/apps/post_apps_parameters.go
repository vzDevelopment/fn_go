// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	modelsv2 "github.com/fnproject/fn_go/modelsv2"
)

// NewPostAppsParams creates a new PostAppsParams object
// with the default values initialized.
func NewPostAppsParams() *PostAppsParams {
	var ()
	return &PostAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAppsParamsWithTimeout creates a new PostAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAppsParamsWithTimeout(timeout time.Duration) *PostAppsParams {
	var ()
	return &PostAppsParams{

		timeout: timeout,
	}
}

// NewPostAppsParamsWithContext creates a new PostAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAppsParamsWithContext(ctx context.Context) *PostAppsParams {
	var ()
	return &PostAppsParams{

		Context: ctx,
	}
}

// NewPostAppsParamsWithHTTPClient creates a new PostAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAppsParamsWithHTTPClient(client *http.Client) *PostAppsParams {
	var ()
	return &PostAppsParams{
		HTTPClient: client,
	}
}

/*PostAppsParams contains all the parameters to send to the API endpoint
for the post apps operation typically these are written to a http.Request
*/
type PostAppsParams struct {

	/*AppBody
	  App to modify.

	*/
	AppBody *modelsv2.App

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post apps params
func (o *PostAppsParams) WithTimeout(timeout time.Duration) *PostAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post apps params
func (o *PostAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post apps params
func (o *PostAppsParams) WithContext(ctx context.Context) *PostAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post apps params
func (o *PostAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post apps params
func (o *PostAppsParams) WithHTTPClient(client *http.Client) *PostAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post apps params
func (o *PostAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppBody adds the appBody to the post apps params
func (o *PostAppsParams) WithAppBody(appBody *modelsv2.App) *PostAppsParams {
	o.SetAppBody(appBody)
	return o
}

// SetAppBody adds the appBody to the post apps params
func (o *PostAppsParams) SetAppBody(appBody *modelsv2.App) {
	o.AppBody = appBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppBody != nil {
		if err := r.SetBodyParam(o.AppBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
