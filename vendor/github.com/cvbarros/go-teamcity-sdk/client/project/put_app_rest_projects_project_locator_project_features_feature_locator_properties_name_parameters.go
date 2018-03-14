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

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams creates a new PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams object
// with the default values initialized.
func NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams() *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	var ()
	return &PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithTimeout creates a new PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithTimeout(timeout time.Duration) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	var ()
	return &PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams{

		timeout: timeout,
	}
}

// NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithContext creates a new PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithContext(ctx context.Context) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	var ()
	return &PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams{

		Context: ctx,
	}
}

// NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithHTTPClient creates a new PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParamsWithHTTPClient(client *http.Client) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	var ()
	return &PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams{
		HTTPClient: client,
	}
}

/*PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams contains all the parameters to send to the API endpoint
for the put app rest projects project locator project features feature locator properties name operation typically these are written to a http.Request
*/
type PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams struct {

	/*Body*/
	Body *models.Property
	/*FeatureLocator*/
	FeatureLocator string
	/*Fields*/
	Fields *string
	/*Name*/
	Name string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithTimeout(timeout time.Duration) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithContext(ctx context.Context) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithHTTPClient(client *http.Client) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithBody(body *models.Property) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithFeatureLocator adds the featureLocator to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithFeatureLocator(featureLocator string) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithFields adds the fields to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithFields(fields *string) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithName(name string) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WithProjectLocator(projectLocator string) *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the put app rest projects project locator project features feature locator properties name params
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
