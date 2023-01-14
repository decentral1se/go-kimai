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

// GetAPITimesheetsReader is a Reader for the GetAPITimesheets structure.
type GetAPITimesheetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITimesheetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITimesheetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPITimesheetsOK creates a GetAPITimesheetsOK with default headers values
func NewGetAPITimesheetsOK() *GetAPITimesheetsOK {
	return &GetAPITimesheetsOK{}
}

/*
GetAPITimesheetsOK describes a response with status code 200, with default header values.

Returns a collection of timesheets records. Be aware that the datetime fields are given in the users local time including the timezone offset via ISO 8601.
*/
type GetAPITimesheetsOK struct {
	Payload []*models.TimesheetCollection
}

// IsSuccess returns true when this get Api timesheets o k response has a 2xx status code
func (o *GetAPITimesheetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api timesheets o k response has a 3xx status code
func (o *GetAPITimesheetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api timesheets o k response has a 4xx status code
func (o *GetAPITimesheetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api timesheets o k response has a 5xx status code
func (o *GetAPITimesheetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api timesheets o k response a status code equal to that given
func (o *GetAPITimesheetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPITimesheetsOK) Error() string {
	return fmt.Sprintf("[GET /api/timesheets][%d] getApiTimesheetsOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsOK) String() string {
	return fmt.Sprintf("[GET /api/timesheets][%d] getApiTimesheetsOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsOK) GetPayload() []*models.TimesheetCollection {
	return o.Payload
}

func (o *GetAPITimesheetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
