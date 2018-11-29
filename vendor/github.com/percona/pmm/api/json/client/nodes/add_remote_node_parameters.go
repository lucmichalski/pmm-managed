// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

	models "github.com/percona/pmm/api/json/models"
)

// NewAddRemoteNodeParams creates a new AddRemoteNodeParams object
// with the default values initialized.
func NewAddRemoteNodeParams() *AddRemoteNodeParams {
	var ()
	return &AddRemoteNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRemoteNodeParamsWithTimeout creates a new AddRemoteNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRemoteNodeParamsWithTimeout(timeout time.Duration) *AddRemoteNodeParams {
	var ()
	return &AddRemoteNodeParams{

		timeout: timeout,
	}
}

// NewAddRemoteNodeParamsWithContext creates a new AddRemoteNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRemoteNodeParamsWithContext(ctx context.Context) *AddRemoteNodeParams {
	var ()
	return &AddRemoteNodeParams{

		Context: ctx,
	}
}

// NewAddRemoteNodeParamsWithHTTPClient creates a new AddRemoteNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRemoteNodeParamsWithHTTPClient(client *http.Client) *AddRemoteNodeParams {
	var ()
	return &AddRemoteNodeParams{
		HTTPClient: client,
	}
}

/*AddRemoteNodeParams contains all the parameters to send to the API endpoint
for the add remote node operation typically these are written to a http.Request
*/
type AddRemoteNodeParams struct {

	/*Body*/
	Body *models.InventoryAddRemoteNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add remote node params
func (o *AddRemoteNodeParams) WithTimeout(timeout time.Duration) *AddRemoteNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add remote node params
func (o *AddRemoteNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add remote node params
func (o *AddRemoteNodeParams) WithContext(ctx context.Context) *AddRemoteNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add remote node params
func (o *AddRemoteNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add remote node params
func (o *AddRemoteNodeParams) WithHTTPClient(client *http.Client) *AddRemoteNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add remote node params
func (o *AddRemoteNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add remote node params
func (o *AddRemoteNodeParams) WithBody(body *models.InventoryAddRemoteNodeRequest) *AddRemoteNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add remote node params
func (o *AddRemoteNodeParams) SetBody(body *models.InventoryAddRemoteNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRemoteNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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