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

// NewRemoveDefaultTemplateParams creates a new RemoveDefaultTemplateParams object
// with the default values initialized.
func NewRemoveDefaultTemplateParams() *RemoveDefaultTemplateParams {
	var ()
	return &RemoveDefaultTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveDefaultTemplateParamsWithTimeout creates a new RemoveDefaultTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveDefaultTemplateParamsWithTimeout(timeout time.Duration) *RemoveDefaultTemplateParams {
	var ()
	return &RemoveDefaultTemplateParams{

		timeout: timeout,
	}
}

// NewRemoveDefaultTemplateParamsWithContext creates a new RemoveDefaultTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveDefaultTemplateParamsWithContext(ctx context.Context) *RemoveDefaultTemplateParams {
	var ()
	return &RemoveDefaultTemplateParams{

		Context: ctx,
	}
}

// NewRemoveDefaultTemplateParamsWithHTTPClient creates a new RemoveDefaultTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveDefaultTemplateParamsWithHTTPClient(client *http.Client) *RemoveDefaultTemplateParams {
	var ()
	return &RemoveDefaultTemplateParams{
		HTTPClient: client,
	}
}

/*RemoveDefaultTemplateParams contains all the parameters to send to the API endpoint
for the remove default template operation typically these are written to a http.Request
*/
type RemoveDefaultTemplateParams struct {

	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove default template params
func (o *RemoveDefaultTemplateParams) WithTimeout(timeout time.Duration) *RemoveDefaultTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove default template params
func (o *RemoveDefaultTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove default template params
func (o *RemoveDefaultTemplateParams) WithContext(ctx context.Context) *RemoveDefaultTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove default template params
func (o *RemoveDefaultTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove default template params
func (o *RemoveDefaultTemplateParams) WithHTTPClient(client *http.Client) *RemoveDefaultTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove default template params
func (o *RemoveDefaultTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the remove default template params
func (o *RemoveDefaultTemplateParams) WithFields(fields *string) *RemoveDefaultTemplateParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the remove default template params
func (o *RemoveDefaultTemplateParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the remove default template params
func (o *RemoveDefaultTemplateParams) WithProjectLocator(projectLocator string) *RemoveDefaultTemplateParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the remove default template params
func (o *RemoveDefaultTemplateParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveDefaultTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

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
