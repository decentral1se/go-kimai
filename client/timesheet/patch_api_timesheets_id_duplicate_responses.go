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

// PatchAPITimesheetsIDDuplicateReader is a Reader for the PatchAPITimesheetsIDDuplicate structure.
type PatchAPITimesheetsIDDuplicateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITimesheetsIDDuplicateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITimesheetsIDDuplicateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITimesheetsIDDuplicateOK creates a PatchAPITimesheetsIDDuplicateOK with default headers values
func NewPatchAPITimesheetsIDDuplicateOK() *PatchAPITimesheetsIDDuplicateOK {
	return &PatchAPITimesheetsIDDuplicateOK{}
}

/*
PatchAPITimesheetsIDDuplicateOK describes a response with status code 200, with default header values.

Duplicates a timesheet record, resetting the export state only.
*/
type PatchAPITimesheetsIDDuplicateOK struct {
	Payload *models.TimesheetEntity
}

// IsSuccess returns true when this patch Api timesheets Id duplicate o k response has a 2xx status code
func (o *PatchAPITimesheetsIDDuplicateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api timesheets Id duplicate o k response has a 3xx status code
func (o *PatchAPITimesheetsIDDuplicateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api timesheets Id duplicate o k response has a 4xx status code
func (o *PatchAPITimesheetsIDDuplicateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api timesheets Id duplicate o k response has a 5xx status code
func (o *PatchAPITimesheetsIDDuplicateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api timesheets Id duplicate o k response a status code equal to that given
func (o *PatchAPITimesheetsIDDuplicateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITimesheetsIDDuplicateOK) Error() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/duplicate][%d] patchApiTimesheetsIdDuplicateOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDDuplicateOK) String() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/duplicate][%d] patchApiTimesheetsIdDuplicateOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDDuplicateOK) GetPayload() *models.TimesheetEntity {
	return o.Payload
}

func (o *PatchAPITimesheetsIDDuplicateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimesheetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}