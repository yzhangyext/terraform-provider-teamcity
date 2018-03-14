// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

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

// NewCreatePoolParams creates a new CreatePoolParams object
// with the default values initialized.
func NewCreatePoolParams() *CreatePoolParams {
	var ()
	return &CreatePoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePoolParamsWithTimeout creates a new CreatePoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePoolParamsWithTimeout(timeout time.Duration) *CreatePoolParams {
	var ()
	return &CreatePoolParams{

		timeout: timeout,
	}
}

// NewCreatePoolParamsWithContext creates a new CreatePoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePoolParamsWithContext(ctx context.Context) *CreatePoolParams {
	var ()
	return &CreatePoolParams{

		Context: ctx,
	}
}

// NewCreatePoolParamsWithHTTPClient creates a new CreatePoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePoolParamsWithHTTPClient(client *http.Client) *CreatePoolParams {
	var ()
	return &CreatePoolParams{
		HTTPClient: client,
	}
}

/*CreatePoolParams contains all the parameters to send to the API endpoint
for the create pool operation typically these are written to a http.Request
*/
type CreatePoolParams struct {

	/*Body*/
	Body *models.AgentPool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create pool params
func (o *CreatePoolParams) WithTimeout(timeout time.Duration) *CreatePoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create pool params
func (o *CreatePoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create pool params
func (o *CreatePoolParams) WithContext(ctx context.Context) *CreatePoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create pool params
func (o *CreatePoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create pool params
func (o *CreatePoolParams) WithHTTPClient(client *http.Client) *CreatePoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create pool params
func (o *CreatePoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create pool params
func (o *CreatePoolParams) WithBody(body *models.AgentPool) *CreatePoolParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create pool params
func (o *CreatePoolParams) SetBody(body *models.AgentPool) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
