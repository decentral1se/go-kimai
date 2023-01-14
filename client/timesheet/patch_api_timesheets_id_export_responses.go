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

// PatchAPITimesheetsIDExportReader is a Reader for the PatchAPITimesheetsIDExport structure.
type PatchAPITimesheetsIDExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITimesheetsIDExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITimesheetsIDExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITimesheetsIDExportOK creates a PatchAPITimesheetsIDExportOK with default headers values
func NewPatchAPITimesheetsIDExportOK() *PatchAPITimesheetsIDExportOK {
	return &PatchAPITimesheetsIDExportOK{}
}

/*
PatchAPITimesheetsIDExportOK describes a response with status code 200, with default header values.

Switches the exported state on the record and therefor locks / unlocks it for further updates. Needs edit_export_*_timesheet permission.
*/
type PatchAPITimesheetsIDExportOK struct {
	Payload *models.TimesheetEntity
}

// IsSuccess returns true when this patch Api timesheets Id export o k response has a 2xx status code
func (o *PatchAPITimesheetsIDExportOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api timesheets Id export o k response has a 3xx status code
func (o *PatchAPITimesheetsIDExportOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api timesheets Id export o k response has a 4xx status code
func (o *PatchAPITimesheetsIDExportOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api timesheets Id export o k response has a 5xx status code
func (o *PatchAPITimesheetsIDExportOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api timesheets Id export o k response a status code equal to that given
func (o *PatchAPITimesheetsIDExportOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITimesheetsIDExportOK) Error() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/export][%d] patchApiTimesheetsIdExportOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDExportOK) String() string {
	return fmt.Sprintf("[PATCH /api/timesheets/{id}/export][%d] patchApiTimesheetsIdExportOK  %+v", 200, o.Payload)
}

func (o *PatchAPITimesheetsIDExportOK) GetPayload() *models.TimesheetEntity {
	return o.Payload
}

func (o *PatchAPITimesheetsIDExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimesheetEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
