// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewAddLicenseKeysParams creates a new AddLicenseKeysParams object
// with the default values initialized.
func NewAddLicenseKeysParams() *AddLicenseKeysParams {
	var ()
	return &AddLicenseKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddLicenseKeysParamsWithTimeout creates a new AddLicenseKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddLicenseKeysParamsWithTimeout(timeout time.Duration) *AddLicenseKeysParams {
	var ()
	return &AddLicenseKeysParams{

		timeout: timeout,
	}
}

// NewAddLicenseKeysParamsWithContext creates a new AddLicenseKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddLicenseKeysParamsWithContext(ctx context.Context) *AddLicenseKeysParams {
	var ()
	return &AddLicenseKeysParams{

		Context: ctx,
	}
}

// NewAddLicenseKeysParamsWithHTTPClient creates a new AddLicenseKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddLicenseKeysParamsWithHTTPClient(client *http.Client) *AddLicenseKeysParams {
	var ()
	return &AddLicenseKeysParams{
		HTTPClient: client,
	}
}

/*AddLicenseKeysParams contains all the parameters to send to the API endpoint
for the add license keys operation typically these are written to a http.Request
*/
type AddLicenseKeysParams struct {

	/*Body*/
	Body string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add license keys params
func (o *AddLicenseKeysParams) WithTimeout(timeout time.Duration) *AddLicenseKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add license keys params
func (o *AddLicenseKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add license keys params
func (o *AddLicenseKeysParams) WithContext(ctx context.Context) *AddLicenseKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add license keys params
func (o *AddLicenseKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add license keys params
func (o *AddLicenseKeysParams) WithHTTPClient(client *http.Client) *AddLicenseKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add license keys params
func (o *AddLicenseKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add license keys params
func (o *AddLicenseKeysParams) WithBody(body string) *AddLicenseKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add license keys params
func (o *AddLicenseKeysParams) SetBody(body string) {
	o.Body = body
}

// WithFields adds the fields to the add license keys params
func (o *AddLicenseKeysParams) WithFields(fields *string) *AddLicenseKeysParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add license keys params
func (o *AddLicenseKeysParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddLicenseKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
