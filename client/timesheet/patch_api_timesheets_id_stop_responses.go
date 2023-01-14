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

// PatchAPITimesheetsIDStopReader is a Reader for the PatchAPITimesheetsIDStop structure.
type PatchAPITimesheetsIDStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITimesheetsIDStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITimesheetsIDStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITimesheetsIDStopOK creates a PatchAPITimesheetsIDStopOK with default headers values
func NewPatchAPITimesheetsIDStopOK() *PatchAPITimesheetsIDStopOK {
	return &PatchAPITimesheetsIDStopOK{}
}

/*
PatchAPITimesheetsIDStopOK describes a response with status code 200, with default header values.

Stops an active timesheet record and returns it afterwards.
*/
type PatchAPITimesheetsIDStopOK struct {
	Payload *models.TimesheetEntity
}

// IsSuccess returns true when this patch Api timesheets Id stop o k response has a 2xx status code
func (o *PatchAPITimesheetsIDStopOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api timesheets Id stop o k response has a 3xx status code
func (o *PatchAPITimesheetsIDStopOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api timesheets Id stop o k response has a 4xx status code
func (o *PatchAPITimesheetsIDStopOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api timesheets Id stop o k response has a 5xx status code
func (o *PatchAPITimesheetsIDStopOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api timesheets Id stop o k response a status code equal to that given
func (o *PatchAPITimesheetsIDStopOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITimesheetsIDStopOK) Error() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/stop][%d] patchApiTimesheetsIdStopOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDStopOK) String() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/stop][%d] patchApiTimesheetsIdStopOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDStopOK) GetPayload() *models.TimesheetEntity {
	return o.Payload
}

func (o *PatchAPITimesheetsIDStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimesheetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}