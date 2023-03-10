// Code generated by go-swagger; DO NOT EDIT.

package default_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIConfigTimesheetReader is a Reader for the GetAPIConfigTimesheet structure.
type GetAPIConfigTimesheetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIConfigTimesheetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIConfigTimesheetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIConfigTimesheetOK creates a GetAPIConfigTimesheetOK with default headers values
func NewGetAPIConfigTimesheetOK() *GetAPIConfigTimesheetOK {
	return &GetAPIConfigTimesheetOK{}
}

/*
GetAPIConfigTimesheetOK describes a response with status code 200, with default header values.

Returns the instance specific timesheet configuration
*/
type GetAPIConfigTimesheetOK struct {
	Payload *models.TimesheetConfig
}

// IsSuccess returns true when this get Api config timesheet o k response has a 2xx status code
func (o *GetAPIConfigTimesheetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api config timesheet o k response has a 3xx status code
func (o *GetAPIConfigTimesheetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api config timesheet o k response has a 4xx status code
func (o *GetAPIConfigTimesheetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api config timesheet o k response has a 5xx status code
func (o *GetAPIConfigTimesheetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api config timesheet o k response a status code equal to that given
func (o *GetAPIConfigTimesheetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIConfigTimesheetOK) Error() string {
	return fmt.Sprintf("[GET /api/config/timesheet][%d] getApiConfigTimesheetOK  %+v", 200, o.Payload)
}

func (o *GetAPIConfigTimesheetOK) String() string {
	return fmt.Sprintf("[GET /api/config/timesheet][%d] getApiConfigTimesheetOK  %+v", 200, o.Payload)
}

func (o *GetAPIConfigTimesheetOK) GetPayload() *models.TimesheetConfig {
	return o.Payload
}

func (o *GetAPIConfigTimesheetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimesheetConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
