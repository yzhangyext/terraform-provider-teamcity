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

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// NewAddRoleParams creates a new AddRoleParams object
// with the default values initialized.
func NewAddRoleParams() *AddRoleParams {
	var ()
	return &AddRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRoleParamsWithTimeout creates a new AddRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRoleParamsWithTimeout(timeout time.Duration) *AddRoleParams {
	var ()
	return &AddRoleParams{

		timeout: timeout,
	}
}

// NewAddRoleParamsWithContext creates a new AddRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRoleParamsWithContext(ctx context.Context) *AddRoleParams {
	var ()
	return &AddRoleParams{

		Context: ctx,
	}
}

// NewAddRoleParamsWithHTTPClient creates a new AddRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRoleParamsWithHTTPClient(client *http.Client) *AddRoleParams {
	var ()
	return &AddRoleParams{
		HTTPClient: client,
	}
}

/*AddRoleParams contains all the parameters to send to the API endpoint
for the add role operation typically these are written to a http.Request
*/
type AddRoleParams struct {

	/*Body*/
	Body *models.Role
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add role params
func (o *AddRoleParams) WithTimeout(timeout time.Duration) *AddRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add role params
func (o *AddRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add role params
func (o *AddRoleParams) WithContext(ctx context.Context) *AddRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add role params
func (o *AddRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add role params
func (o *AddRoleParams) WithHTTPClient(client *http.Client) *AddRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add role params
func (o *AddRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add role params
func (o *AddRoleParams) WithBody(body *models.Role) *AddRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add role params
func (o *AddRoleParams) SetBody(body *models.Role) {
	o.Body = body
}

// WithUserLocator adds the userLocator to the add role params
func (o *AddRoleParams) WithUserLocator(userLocator string) *AddRoleParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the add role params
func (o *AddRoleParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *AddRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
