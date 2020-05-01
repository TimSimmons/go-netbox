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

package extras

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
)

// NewExtrasChoicesReadParams creates a new ExtrasChoicesReadParams object
// with the default values initialized.
func NewExtrasChoicesReadParams() *ExtrasChoicesReadParams {
	var ()
	return &ExtrasChoicesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasChoicesReadParamsWithTimeout creates a new ExtrasChoicesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasChoicesReadParamsWithTimeout(timeout time.Duration) *ExtrasChoicesReadParams {
	var ()
	return &ExtrasChoicesReadParams{

		timeout: timeout,
	}
}

// NewExtrasChoicesReadParamsWithContext creates a new ExtrasChoicesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasChoicesReadParamsWithContext(ctx context.Context) *ExtrasChoicesReadParams {
	var ()
	return &ExtrasChoicesReadParams{

		Context: ctx,
	}
}

// NewExtrasChoicesReadParamsWithHTTPClient creates a new ExtrasChoicesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasChoicesReadParamsWithHTTPClient(client *http.Client) *ExtrasChoicesReadParams {
	var ()
	return &ExtrasChoicesReadParams{
		HTTPClient: client,
	}
}

/*ExtrasChoicesReadParams contains all the parameters to send to the API endpoint
for the extras choices read operation typically these are written to a http.Request
*/
type ExtrasChoicesReadParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras choices read params
func (o *ExtrasChoicesReadParams) WithTimeout(timeout time.Duration) *ExtrasChoicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras choices read params
func (o *ExtrasChoicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras choices read params
func (o *ExtrasChoicesReadParams) WithContext(ctx context.Context) *ExtrasChoicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras choices read params
func (o *ExtrasChoicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras choices read params
func (o *ExtrasChoicesReadParams) WithHTTPClient(client *http.Client) *ExtrasChoicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras choices read params
func (o *ExtrasChoicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras choices read params
func (o *ExtrasChoicesReadParams) WithID(id string) *ExtrasChoicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras choices read params
func (o *ExtrasChoicesReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasChoicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
