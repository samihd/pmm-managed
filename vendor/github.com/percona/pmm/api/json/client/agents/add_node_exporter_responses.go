// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// AddNodeExporterReader is a Reader for the AddNodeExporter structure.
type AddNodeExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodeExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddNodeExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddNodeExporterOK creates a AddNodeExporterOK with default headers values
func NewAddNodeExporterOK() *AddNodeExporterOK {
	return &AddNodeExporterOK{}
}

/*AddNodeExporterOK handles this case with default header values.

(empty)
*/
type AddNodeExporterOK struct {
	Payload *models.InventoryAddNodeExporterResponse
}

func (o *AddNodeExporterOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/AddNodeExporter][%d] addNodeExporterOK  %+v", 200, o.Payload)
}

func (o *AddNodeExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddNodeExporterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}