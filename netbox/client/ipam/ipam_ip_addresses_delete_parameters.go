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

package ipam

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
)

// NewIpamIPAddressesDeleteParams creates a new IpamIPAddressesDeleteParams object
// with the default values initialized.
func NewIpamIPAddressesDeleteParams() *IpamIPAddressesDeleteParams {
	var ()
	return &IpamIPAddressesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesDeleteParamsWithTimeout creates a new IpamIPAddressesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIpamIPAddressesDeleteParamsWithTimeout(timeout time.Duration) *IpamIPAddressesDeleteParams {
	var ()
	return &IpamIPAddressesDeleteParams{

		timeout: timeout,
	}
}

// NewIpamIPAddressesDeleteParamsWithContext creates a new IpamIPAddressesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIpamIPAddressesDeleteParamsWithContext(ctx context.Context) *IpamIPAddressesDeleteParams {
	var ()
	return &IpamIPAddressesDeleteParams{

		Context: ctx,
	}
}

// NewIpamIPAddressesDeleteParamsWithHTTPClient creates a new IpamIPAddressesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIpamIPAddressesDeleteParamsWithHTTPClient(client *http.Client) *IpamIPAddressesDeleteParams {
	var ()
	return &IpamIPAddressesDeleteParams{
		HTTPClient: client,
	}
}

/*IpamIPAddressesDeleteParams contains all the parameters to send to the API endpoint
for the ipam ip addresses delete operation typically these are written to a http.Request
*/
type IpamIPAddressesDeleteParams struct {

	/*ID
	  A unique integer value identifying this IP address.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithTimeout(timeout time.Duration) *IpamIPAddressesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithContext(ctx context.Context) *IpamIPAddressesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithHTTPClient(client *http.Client) *IpamIPAddressesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) WithID(id int64) *IpamIPAddressesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses delete params
func (o *IpamIPAddressesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
