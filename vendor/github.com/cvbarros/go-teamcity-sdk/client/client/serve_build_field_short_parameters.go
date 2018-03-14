// Code generated by go-swagger; DO NOT EDIT.

package client

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

// NewServeBuildFieldShortParams creates a new ServeBuildFieldShortParams object
// with the default values initialized.
func NewServeBuildFieldShortParams() *ServeBuildFieldShortParams {
	var ()
	return &ServeBuildFieldShortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildFieldShortParamsWithTimeout creates a new ServeBuildFieldShortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildFieldShortParamsWithTimeout(timeout time.Duration) *ServeBuildFieldShortParams {
	var ()
	return &ServeBuildFieldShortParams{

		timeout: timeout,
	}
}

// NewServeBuildFieldShortParamsWithContext creates a new ServeBuildFieldShortParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildFieldShortParamsWithContext(ctx context.Context) *ServeBuildFieldShortParams {
	var ()
	return &ServeBuildFieldShortParams{

		Context: ctx,
	}
}

// NewServeBuildFieldShortParamsWithHTTPClient creates a new ServeBuildFieldShortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildFieldShortParamsWithHTTPClient(client *http.Client) *ServeBuildFieldShortParams {
	var ()
	return &ServeBuildFieldShortParams{
		HTTPClient: client,
	}
}

/*ServeBuildFieldShortParams contains all the parameters to send to the API endpoint
for the serve build field short operation typically these are written to a http.Request
*/
type ServeBuildFieldShortParams struct {

	/*BtLocator*/
	BtLocator string
	/*BuildLocator*/
	BuildLocator string
	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build field short params
func (o *ServeBuildFieldShortParams) WithTimeout(timeout time.Duration) *ServeBuildFieldShortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build field short params
func (o *ServeBuildFieldShortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build field short params
func (o *ServeBuildFieldShortParams) WithContext(ctx context.Context) *ServeBuildFieldShortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build field short params
func (o *ServeBuildFieldShortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build field short params
func (o *ServeBuildFieldShortParams) WithHTTPClient(client *http.Client) *ServeBuildFieldShortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build field short params
func (o *ServeBuildFieldShortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) WithBtLocator(btLocator string) *ServeBuildFieldShortParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithBuildLocator adds the buildLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) WithBuildLocator(buildLocator string) *ServeBuildFieldShortParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithField adds the field to the serve build field short params
func (o *ServeBuildFieldShortParams) WithField(field string) *ServeBuildFieldShortParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve build field short params
func (o *ServeBuildFieldShortParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) WithProjectLocator(projectLocator string) *ServeBuildFieldShortParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the serve build field short params
func (o *ServeBuildFieldShortParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildFieldShortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
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
