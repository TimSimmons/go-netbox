// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/timsimmons/go-netbox/netbox/models"
)

// NewDcimRackGroupsUpdateParams creates a new DcimRackGroupsUpdateParams object
// with the default values initialized.
func NewDcimRackGroupsUpdateParams() *DcimRackGroupsUpdateParams {
	var ()
	return &DcimRackGroupsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsUpdateParamsWithTimeout creates a new DcimRackGroupsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsUpdateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsUpdateParams {
	var ()
	return &DcimRackGroupsUpdateParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsUpdateParamsWithContext creates a new DcimRackGroupsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsUpdateParamsWithContext(ctx context.Context) *DcimRackGroupsUpdateParams {
	var ()
	return &DcimRackGroupsUpdateParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsUpdateParamsWithHTTPClient creates a new DcimRackGroupsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsUpdateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsUpdateParams {
	var ()
	return &DcimRackGroupsUpdateParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsUpdateParams contains all the parameters to send to the API endpoint
for the dcim rack groups update operation typically these are written to a http.Request
*/
type DcimRackGroupsUpdateParams struct {

	/*Data*/
	Data *models.WritableRackGroup
	/*ID
	  A unique integer value identifying this rack group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithContext(ctx context.Context) *DcimRackGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WithID adds the id to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) WithID(id int64) *DcimRackGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack groups update params
func (o *DcimRackGroupsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
