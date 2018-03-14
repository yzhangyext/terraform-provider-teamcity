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

// NewGetIPAddressParams creates a new GetIPAddressParams object
// with the default values initialized.
func NewGetIPAddressParams() *GetIPAddressParams {
	var ()
	return &GetIPAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIPAddressParamsWithTimeout creates a new GetIPAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIPAddressParamsWithTimeout(timeout time.Duration) *GetIPAddressParams {
	var ()
	return &GetIPAddressParams{

		timeout: timeout,
	}
}

// NewGetIPAddressParamsWithContext creates a new GetIPAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIPAddressParamsWithContext(ctx context.Context) *GetIPAddressParams {
	var ()
	return &GetIPAddressParams{

		Context: ctx,
	}
}

// NewGetIPAddressParamsWithHTTPClient creates a new GetIPAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIPAddressParamsWithHTTPClient(client *http.Client) *GetIPAddressParams {
	var ()
	return &GetIPAddressParams{
		HTTPClient: client,
	}
}

/*GetIPAddressParams contains all the parameters to send to the API endpoint
for the get Ip address operation typically these are written to a http.Request
*/
type GetIPAddressParams struct {

	/*Host*/
	Host string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Ip address params
func (o *GetIPAddressParams) WithTimeout(timeout time.Duration) *GetIPAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Ip address params
func (o *GetIPAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Ip address params
func (o *GetIPAddressParams) WithContext(ctx context.Context) *GetIPAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Ip address params
func (o *GetIPAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Ip address params
func (o *GetIPAddressParams) WithHTTPClient(client *http.Client) *GetIPAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Ip address params
func (o *GetIPAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHost adds the host to the get Ip address params
func (o *GetIPAddressParams) WithHost(host string) *GetIPAddressParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the get Ip address params
func (o *GetIPAddressParams) SetHost(host string) {
	o.Host = host
}

// WriteToRequest writes these params to a swagger request
func (o *GetIPAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param host
	if err := r.SetPathParam("host", o.Host); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
