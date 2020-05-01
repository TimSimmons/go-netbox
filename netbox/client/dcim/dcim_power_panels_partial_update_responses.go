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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/timsimmons/go-netbox/netbox/models"
)

// DcimPowerPanelsPartialUpdateReader is a Reader for the DcimPowerPanelsPartialUpdate structure.
type DcimPowerPanelsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPanelsPartialUpdateOK creates a DcimPowerPanelsPartialUpdateOK with default headers values
func NewDcimPowerPanelsPartialUpdateOK() *DcimPowerPanelsPartialUpdateOK {
	return &DcimPowerPanelsPartialUpdateOK{}
}

/*DcimPowerPanelsPartialUpdateOK handles this case with default header values.

DcimPowerPanelsPartialUpdateOK dcim power panels partial update o k
*/
type DcimPowerPanelsPartialUpdateOK struct {
	Payload *models.PowerPanel
}

func (o *DcimPowerPanelsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
