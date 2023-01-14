// Code generated by go-swagger; DO NOT EDIT.

package expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// PatchAPIExpensesIDReader is a Reader for the PatchAPIExpensesID structure.
type PatchAPIExpensesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIExpensesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIExpensesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIExpensesIDOK creates a PatchAPIExpensesIDOK with default headers values
func NewPatchAPIExpensesIDOK() *PatchAPIExpensesIDOK {
	return &PatchAPIExpensesIDOK{}
}

/*
PatchAPIExpensesIDOK describes a response with status code 200, with default header values.

Returns the updated expense
*/
type PatchAPIExpensesIDOK struct {
	Payload *models.ExpenseEntity
}

// IsSuccess returns true when this patch Api expenses Id o k response has a 2xx status code
func (o *PatchAPIExpensesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api expenses Id o k response has a 3xx status code
func (o *PatchAPIExpensesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api expenses Id o k response has a 4xx status code
func (o *PatchAPIExpensesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api expenses Id o k response has a 5xx status code
func (o *PatchAPIExpensesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api expenses Id o k response a status code equal to that given
func (o *PatchAPIExpensesIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPIExpensesIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/expenses/{id}][%d] patchApiExpensesIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPIExpensesIDOK) String() string {
	return fmt.Sprintf("[PATCH /api/expenses/{id}][%d] patchApiExpensesIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPIExpensesIDOK) GetPayload() *models.ExpenseEntity {
	return o.Payload
}

func (o *PatchAPIExpensesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpenseEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
