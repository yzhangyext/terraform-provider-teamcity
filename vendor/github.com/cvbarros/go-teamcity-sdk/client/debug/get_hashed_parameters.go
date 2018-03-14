// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewGetHashedParams creates a new GetHashedParams object
// with the default values initialized.
func NewGetHashedParams() *GetHashedParams {
	var ()
	return &GetHashedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHashedParamsWithTimeout creates a new GetHashedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHashedParamsWithTimeout(timeout time.Duration) *GetHashedParams {
	var ()
	return &GetHashedParams{

		timeout: timeout,
	}
}

// NewGetHashedParamsWithContext creates a new GetHashedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHashedParamsWithContext(ctx context.Context) *GetHashedParams {
	var ()
	return &GetHashedParams{

		Context: ctx,
	}
}

// NewGetHashedParamsWithHTTPClient creates a new GetHashedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHashedParamsWithHTTPClient(client *http.Client) *GetHashedParams {
	var ()
	return &GetHashedParams{
		HTTPClient: client,
	}
}

/*GetHashedParams contains all the parameters to send to the API endpoint
for the get hashed operation typically these are written to a http.Request
*/
type GetHashedParams struct {

	/*Method*/
	Method string
	/*Value*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hashed params
func (o *GetHashedParams) WithTimeout(timeout time.Duration) *GetHashedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hashed params
func (o *GetHashedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hashed params
func (o *GetHashedParams) WithContext(ctx context.Context) *GetHashedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hashed params
func (o *GetHashedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hashed params
func (o *GetHashedParams) WithHTTPClient(client *http.Client) *GetHashedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hashed params
func (o *GetHashedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMethod adds the method to the get hashed params
func (o *GetHashedParams) WithMethod(method string) *GetHashedParams {
	o.SetMethod(method)
	return o
}

// SetMethod adds the method to the get hashed params
func (o *GetHashedParams) SetMethod(method string) {
	o.Method = method
}

// WithValue adds the value to the get hashed params
func (o *GetHashedParams) WithValue(value *string) *GetHashedParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get hashed params
func (o *GetHashedParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetHashedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param method
	if err := r.SetPathParam("method", o.Method); err != nil {
		return err
	}

	if o.Value != nil {

		// query param value
		var qrValue string
		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {
			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
