// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// PatchAPICustomersIDReader is a Reader for the PatchAPICustomersID structure.
type PatchAPICustomersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPICustomersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPICustomersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPICustomersIDOK creates a PatchAPICustomersIDOK with default headers values
func NewPatchAPICustomersIDOK() *PatchAPICustomersIDOK {
	return &PatchAPICustomersIDOK{}
}

/*
PatchAPICustomersIDOK describes a response with status code 200, with default header values.

Returns the updated customer
*/
type PatchAPICustomersIDOK struct {
	Payload *models.CustomerEntity
}

// IsSuccess returns true when this patch Api customers Id o k response has a 2xx status code
func (o *PatchAPICustomersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api customers Id o k response has a 3xx status code
func (o *PatchAPICustomersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api customers Id o k response has a 4xx status code
func (o *PatchAPICustomersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api customers Id o k response has a 5xx status code
func (o *PatchAPICustomersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api customers Id o k response a status code equal to that given
func (o *PatchAPICustomersIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPICustomersIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/customers/{id}][%d] patchApiCustomersIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPICustomersIDOK) String() string {
	return fmt.Sprintf("[PATCH /api/customers/{id}][%d] patchApiCustomersIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPICustomersIDOK) GetPayload() *models.CustomerEntity {
	return o.Payload
}

func (o *PatchAPICustomersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
