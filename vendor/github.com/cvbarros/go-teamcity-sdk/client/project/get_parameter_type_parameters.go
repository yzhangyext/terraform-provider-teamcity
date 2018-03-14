// Code generated by go-swagger; DO NOT EDIT.

package project

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
)

// NewGetParameterTypeParams creates a new GetParameterTypeParams object
// with the default values initialized.
func NewGetParameterTypeParams() *GetParameterTypeParams {
	var ()
	return &GetParameterTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParameterTypeParamsWithTimeout creates a new GetParameterTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParameterTypeParamsWithTimeout(timeout time.Duration) *GetParameterTypeParams {
	var ()
	return &GetParameterTypeParams{

		timeout: timeout,
	}
}

// NewGetParameterTypeParamsWithContext creates a new GetParameterTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParameterTypeParamsWithContext(ctx context.Context) *GetParameterTypeParams {
	var ()
	return &GetParameterTypeParams{

		Context: ctx,
	}
}

// NewGetParameterTypeParamsWithHTTPClient creates a new GetParameterTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParameterTypeParamsWithHTTPClient(client *http.Client) *GetParameterTypeParams {
	var ()
	return &GetParameterTypeParams{
		HTTPClient: client,
	}
}

/*GetParameterTypeParams contains all the parameters to send to the API endpoint
for the get parameter type operation typically these are written to a http.Request
*/
type GetParameterTypeParams struct {

	/*Name*/
	Name string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get parameter type params
func (o *GetParameterTypeParams) WithTimeout(timeout time.Duration) *GetParameterTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameter type params
func (o *GetParameterTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameter type params
func (o *GetParameterTypeParams) WithContext(ctx context.Context) *GetParameterTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameter type params
func (o *GetParameterTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameter type params
func (o *GetParameterTypeParams) WithHTTPClient(client *http.Client) *GetParameterTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameter type params
func (o *GetParameterTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get parameter type params
func (o *GetParameterTypeParams) WithName(name string) *GetParameterTypeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get parameter type params
func (o *GetParameterTypeParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the get parameter type params
func (o *GetParameterTypeParams) WithProjectLocator(projectLocator string) *GetParameterTypeParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get parameter type params
func (o *GetParameterTypeParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetParameterTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
