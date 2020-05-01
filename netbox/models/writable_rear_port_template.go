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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableRearPortTemplate writable rear port template
//
// swagger:model WritableRearPortTemplate
type WritableRearPortTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 64
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// Type
	// Required: true
	// Enum: [8p8c 110-punch bnc fc lc lc-apc lsh lsh-apc mpo mtrj sc sc-apc st]
	Type *string `json:"type"`
}

// Validate validates this writable rear port template
func (m *WritableRearPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableRearPortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *WritableRearPortTemplate) validatePositions(formats strfmt.Registry) error {

	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", int64(m.Positions), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", int64(m.Positions), 64, false); err != nil {
		return err
	}

	return nil
}

var writableRearPortTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","110-punch","bnc","fc","lc","lc-apc","lsh","lsh-apc","mpo","mtrj","sc","sc-apc","st"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableRearPortTemplateTypeTypePropEnum = append(writableRearPortTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableRearPortTemplateTypeNr8p8c captures enum value "8p8c"
	WritableRearPortTemplateTypeNr8p8c string = "8p8c"

	// WritableRearPortTemplateTypeNr110Punch captures enum value "110-punch"
	WritableRearPortTemplateTypeNr110Punch string = "110-punch"

	// WritableRearPortTemplateTypeBnc captures enum value "bnc"
	WritableRearPortTemplateTypeBnc string = "bnc"

	// WritableRearPortTemplateTypeFc captures enum value "fc"
	WritableRearPortTemplateTypeFc string = "fc"

	// WritableRearPortTemplateTypeLc captures enum value "lc"
	WritableRearPortTemplateTypeLc string = "lc"

	// WritableRearPortTemplateTypeLcApc captures enum value "lc-apc"
	WritableRearPortTemplateTypeLcApc string = "lc-apc"

	// WritableRearPortTemplateTypeLsh captures enum value "lsh"
	WritableRearPortTemplateTypeLsh string = "lsh"

	// WritableRearPortTemplateTypeLshApc captures enum value "lsh-apc"
	WritableRearPortTemplateTypeLshApc string = "lsh-apc"

	// WritableRearPortTemplateTypeMpo captures enum value "mpo"
	WritableRearPortTemplateTypeMpo string = "mpo"

	// WritableRearPortTemplateTypeMtrj captures enum value "mtrj"
	WritableRearPortTemplateTypeMtrj string = "mtrj"

	// WritableRearPortTemplateTypeSc captures enum value "sc"
	WritableRearPortTemplateTypeSc string = "sc"

	// WritableRearPortTemplateTypeScApc captures enum value "sc-apc"
	WritableRearPortTemplateTypeScApc string = "sc-apc"

	// WritableRearPortTemplateTypeSt captures enum value "st"
	WritableRearPortTemplateTypeSt string = "st"
)

// prop value enum
func (m *WritableRearPortTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableRearPortTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableRearPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableRearPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableRearPortTemplate) UnmarshalBinary(b []byte) error {
	var res WritableRearPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
