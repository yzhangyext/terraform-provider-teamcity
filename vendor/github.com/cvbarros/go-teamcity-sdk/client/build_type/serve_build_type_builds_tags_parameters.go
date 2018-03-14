// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewServeBuildTypeBuildsTagsParams creates a new ServeBuildTypeBuildsTagsParams object
// with the default values initialized.
func NewServeBuildTypeBuildsTagsParams() *ServeBuildTypeBuildsTagsParams {
	var ()
	return &ServeBuildTypeBuildsTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildTypeBuildsTagsParamsWithTimeout creates a new ServeBuildTypeBuildsTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildTypeBuildsTagsParamsWithTimeout(timeout time.Duration) *ServeBuildTypeBuildsTagsParams {
	var ()
	return &ServeBuildTypeBuildsTagsParams{

		timeout: timeout,
	}
}

// NewServeBuildTypeBuildsTagsParamsWithContext creates a new ServeBuildTypeBuildsTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildTypeBuildsTagsParamsWithContext(ctx context.Context) *ServeBuildTypeBuildsTagsParams {
	var ()
	return &ServeBuildTypeBuildsTagsParams{

		Context: ctx,
	}
}

// NewServeBuildTypeBuildsTagsParamsWithHTTPClient creates a new ServeBuildTypeBuildsTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildTypeBuildsTagsParamsWithHTTPClient(client *http.Client) *ServeBuildTypeBuildsTagsParams {
	var ()
	return &ServeBuildTypeBuildsTagsParams{
		HTTPClient: client,
	}
}

/*ServeBuildTypeBuildsTagsParams contains all the parameters to send to the API endpoint
for the serve build type builds tags operation typically these are written to a http.Request
*/
type ServeBuildTypeBuildsTagsParams struct {

	/*BtLocator*/
	BtLocator string
	/*Field*/
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) WithTimeout(timeout time.Duration) *ServeBuildTypeBuildsTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) WithContext(ctx context.Context) *ServeBuildTypeBuildsTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) WithHTTPClient(client *http.Client) *ServeBuildTypeBuildsTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) WithBtLocator(btLocator string) *ServeBuildTypeBuildsTagsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithField adds the field to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) WithField(field string) *ServeBuildTypeBuildsTagsParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve build type builds tags params
func (o *ServeBuildTypeBuildsTagsParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildTypeBuildsTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}