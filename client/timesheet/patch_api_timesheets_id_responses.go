// Code generated by go-swagger; DO NOT EDIT.

package timesheet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// PatchAPITimesheetsIDReader is a Reader for the PatchAPITimesheetsID structure.
type PatchAPITimesheetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITimesheetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITimesheetsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITimesheetsIDOK creates a PatchAPITimesheetsIDOK with default headers values
func NewPatchAPITimesheetsIDOK() *PatchAPITimesheetsIDOK {
	return &PatchAPITimesheetsIDOK{}
}

/*
PatchAPITimesheetsIDOK describes a response with status code 200, with default header values.

Returns the updated timesheet
*/
type PatchAPITimesheetsIDOK struct {
	Payload *models.TimesheetEntity
}

// IsSuccess returns true when this patch Api timesheets Id o k response has a 2xx status code
func (o *PatchAPITimesheetsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api timesheets Id o k response has a 3xx status code
func (o *PatchAPITimesheetsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api timesheets Id o k response has a 4xx status code
func (o *PatchAPITimesheetsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api timesheets Id o k response has a 5xx status code
func (o *PatchAPITimesheetsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api timesheets Id o k response a status code equal to that given
func (o *PatchAPITimesheetsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITimesheetsIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}][%d] patchApiTimesheetsIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDOK) String() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}][%d] patchApiTimesheetsIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDOK) GetPayload() *models.TimesheetEntity {
	return o.Payload
}

func (o *PatchAPITimesheetsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimesheetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
