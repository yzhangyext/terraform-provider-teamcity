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

// NewDeleteLicenseKeyParams creates a new DeleteLicenseKeyParams object
// with the default values initialized.
func NewDeleteLicenseKeyParams() *DeleteLicenseKeyParams {
	var ()
	return &DeleteLicenseKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLicenseKeyParamsWithTimeout creates a new DeleteLicenseKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLicenseKeyParamsWithTimeout(timeout time.Duration) *DeleteLicenseKeyParams {
	var ()
	return &DeleteLicenseKeyParams{

		timeout: timeout,
	}
}

// NewDeleteLicenseKeyParamsWithContext creates a new DeleteLicenseKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLicenseKeyParamsWithContext(ctx context.Context) *DeleteLicenseKeyParams {
	var ()
	return &DeleteLicenseKeyParams{

		Context: ctx,
	}
}

// NewDeleteLicenseKeyParamsWithHTTPClient creates a new DeleteLicenseKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLicenseKeyParamsWithHTTPClient(client *http.Client) *DeleteLicenseKeyParams {
	var ()
	return &DeleteLicenseKeyParams{
		HTTPClient: client,
	}
}

/*DeleteLicenseKeyParams contains all the parameters to send to the API endpoint
for the delete license key operation typically these are written to a http.Request
*/
type DeleteLicenseKeyParams struct {

	/*LicenseKey*/
	LicenseKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete license key params
func (o *DeleteLicenseKeyParams) WithTimeout(timeout time.Duration) *DeleteLicenseKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete license key params
func (o *DeleteLicenseKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete license key params
func (o *DeleteLicenseKeyParams) WithContext(ctx context.Context) *DeleteLicenseKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete license key params
func (o *DeleteLicenseKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete license key params
func (o *DeleteLicenseKeyParams) WithHTTPClient(client *http.Client) *DeleteLicenseKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete license key params
func (o *DeleteLicenseKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLicenseKey adds the licenseKey to the delete license key params
func (o *DeleteLicenseKeyParams) WithLicenseKey(licenseKey string) *DeleteLicenseKeyParams {
	o.SetLicenseKey(licenseKey)
	return o
}

// SetLicenseKey adds the licenseKey to the delete license key params
func (o *DeleteLicenseKeyParams) SetLicenseKey(licenseKey string) {
	o.LicenseKey = licenseKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLicenseKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param licenseKey
	if err := r.SetPathParam("licenseKey", o.LicenseKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
