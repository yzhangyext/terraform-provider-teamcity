// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewSetUserFieldParams creates a new SetUserFieldParams object
// with the default values initialized.
func NewSetUserFieldParams() *SetUserFieldParams {
	var ()
	return &SetUserFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetUserFieldParamsWithTimeout creates a new SetUserFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetUserFieldParamsWithTimeout(timeout time.Duration) *SetUserFieldParams {
	var ()
	return &SetUserFieldParams{

		timeout: timeout,
	}
}

// NewSetUserFieldParamsWithContext creates a new SetUserFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetUserFieldParamsWithContext(ctx context.Context) *SetUserFieldParams {
	var ()
	return &SetUserFieldParams{

		Context: ctx,
	}
}

// NewSetUserFieldParamsWithHTTPClient creates a new SetUserFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetUserFieldParamsWithHTTPClient(client *http.Client) *SetUserFieldParams {
	var ()
	return &SetUserFieldParams{
		HTTPClient: client,
	}
}

/*SetUserFieldParams contains all the parameters to send to the API endpoint
for the set user field operation typically these are written to a http.Request
*/
type SetUserFieldParams struct {

	/*Body*/
	Body string
	/*Field*/
	Field string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set user field params
func (o *SetUserFieldParams) WithTimeout(timeout time.Duration) *SetUserFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set user field params
func (o *SetUserFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set user field params
func (o *SetUserFieldParams) WithContext(ctx context.Context) *SetUserFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set user field params
func (o *SetUserFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set user field params
func (o *SetUserFieldParams) WithHTTPClient(client *http.Client) *SetUserFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set user field params
func (o *SetUserFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set user field params
func (o *SetUserFieldParams) WithBody(body string) *SetUserFieldParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set user field params
func (o *SetUserFieldParams) SetBody(body string) {
	o.Body = body
}

// WithField adds the field to the set user field params
func (o *SetUserFieldParams) WithField(field string) *SetUserFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the set user field params
func (o *SetUserFieldParams) SetField(field string) {
	o.Field = field
}

// WithUserLocator adds the userLocator to the set user field params
func (o *SetUserFieldParams) WithUserLocator(userLocator string) *SetUserFieldParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the set user field params
func (o *SetUserFieldParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetUserFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
