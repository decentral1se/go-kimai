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

// GetAPITimesheetsActiveReader is a Reader for the GetAPITimesheetsActive structure.
type GetAPITimesheetsActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITimesheetsActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITimesheetsActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPITimesheetsActiveOK creates a GetAPITimesheetsActiveOK with default headers values
func NewGetAPITimesheetsActiveOK() *GetAPITimesheetsActiveOK {
	return &GetAPITimesheetsActiveOK{}
}

/*
GetAPITimesheetsActiveOK describes a response with status code 200, with default header values.

Returns the collection of active timesheet records for the current user
*/
type GetAPITimesheetsActiveOK struct {
	Payload []*models.TimesheetCollectionExpanded
}

// IsSuccess returns true when this get Api timesheets active o k response has a 2xx status code
func (o *GetAPITimesheetsActiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api timesheets active o k response has a 3xx status code
func (o *GetAPITimesheetsActiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api timesheets active o k response has a 4xx status code
func (o *GetAPITimesheetsActiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api timesheets active o k response has a 5xx status code
func (o *GetAPITimesheetsActiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api timesheets active o k response a status code equal to that given
func (o *GetAPITimesheetsActiveOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPITimesheetsActiveOK) Error() string {
	return fmt.Sprintf("[GET /api/timesheets/active][%d] getApiTimesheetsActiveOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsActiveOK) String() string {
	return fmt.Sprintf("[GET /api/timesheets/active][%d] getApiTimesheetsActiveOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsActiveOK) GetPayload() []*models.TimesheetCollectionExpanded {
	return o.Payload
}

func (o *GetAPITimesheetsActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}