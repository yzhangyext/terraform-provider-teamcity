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

// NewGetParametersParams creates a new GetParametersParams object
// with the default values initialized.
func NewGetParametersParams() *GetParametersParams {
	var ()
	return &GetParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParametersParamsWithTimeout creates a new GetParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParametersParamsWithTimeout(timeout time.Duration) *GetParametersParams {
	var ()
	return &GetParametersParams{

		timeout: timeout,
	}
}

// NewGetParametersParamsWithContext creates a new GetParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParametersParamsWithContext(ctx context.Context) *GetParametersParams {
	var ()
	return &GetParametersParams{

		Context: ctx,
	}
}

// NewGetParametersParamsWithHTTPClient creates a new GetParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParametersParamsWithHTTPClient(client *http.Client) *GetParametersParams {
	var ()
	return &GetParametersParams{
		HTTPClient: client,
	}
}

/*GetParametersParams contains all the parameters to send to the API endpoint
for the get parameters operation typically these are written to a http.Request
*/
type GetParametersParams struct {

	/*FeatureLocator*/
	FeatureLocator string
	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get parameters params
func (o *GetParametersParams) WithTimeout(timeout time.Duration) *GetParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameters params
func (o *GetParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameters params
func (o *GetParametersParams) WithContext(ctx context.Context) *GetParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameters params
func (o *GetParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameters params
func (o *GetParametersParams) WithHTTPClient(client *http.Client) *GetParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameters params
func (o *GetParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureLocator adds the featureLocator to the get parameters params
func (o *GetParametersParams) WithFeatureLocator(featureLocator string) *GetParametersParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the get parameters params
func (o *GetParametersParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithFields adds the fields to the get parameters params
func (o *GetParametersParams) WithFields(fields *string) *GetParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get parameters params
func (o *GetParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get parameters params
func (o *GetParametersParams) WithLocator(locator *string) *GetParametersParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get parameters params
func (o *GetParametersParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithProjectLocator adds the projectLocator to the get parameters params
func (o *GetParametersParams) WithProjectLocator(projectLocator string) *GetParametersParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get parameters params
func (o *GetParametersParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureLocator
	if err := r.SetPathParam("featureLocator", o.FeatureLocator); err != nil {
		return err
	}

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string
		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {
			if err := r.SetQueryParam("locator", qLocator); err != nil {
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
