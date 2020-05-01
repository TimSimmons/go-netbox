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
	"github.com/go-openapi/swag"
)

// NewExtrasExportTemplatesListParams creates a new ExtrasExportTemplatesListParams object
// with the default values initialized.
func NewExtrasExportTemplatesListParams() *ExtrasExportTemplatesListParams {
	var ()
	return &ExtrasExportTemplatesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesListParamsWithTimeout creates a new ExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasExportTemplatesListParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	var ()
	return &ExtrasExportTemplatesListParams{

		timeout: timeout,
	}
}

// NewExtrasExportTemplatesListParamsWithContext creates a new ExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasExportTemplatesListParamsWithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	var ()
	return &ExtrasExportTemplatesListParams{

		Context: ctx,
	}
}

// NewExtrasExportTemplatesListParamsWithHTTPClient creates a new ExtrasExportTemplatesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasExportTemplatesListParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	var ()
	return &ExtrasExportTemplatesListParams{
		HTTPClient: client,
	}
}

/*ExtrasExportTemplatesListParams contains all the parameters to send to the API endpoint
for the extras export templates list operation typically these are written to a http.Request
*/
type ExtrasExportTemplatesListParams struct {

	/*ContentType*/
	ContentType *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*TemplateLanguage*/
	TemplateLanguage *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContext(ctx context.Context) *ExtrasExportTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithContentType(contentType *string) *ExtrasExportTemplatesListParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithLimit(limit *int64) *ExtrasExportTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithName(name *string) *ExtrasExportTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithOffset(offset *int64) *ExtrasExportTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithTemplateLanguage adds the templateLanguage to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) WithTemplateLanguage(templateLanguage *string) *ExtrasExportTemplatesListParams {
	o.SetTemplateLanguage(templateLanguage)
	return o
}

// SetTemplateLanguage adds the templateLanguage to the extras export templates list params
func (o *ExtrasExportTemplatesListParams) SetTemplateLanguage(templateLanguage *string) {
	o.TemplateLanguage = templateLanguage
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string
		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {
			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.TemplateLanguage != nil {

		// query param template_language
		var qrTemplateLanguage string
		if o.TemplateLanguage != nil {
			qrTemplateLanguage = *o.TemplateLanguage
		}
		qTemplateLanguage := qrTemplateLanguage
		if qTemplateLanguage != "" {
			if err := r.SetQueryParam("template_language", qTemplateLanguage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
