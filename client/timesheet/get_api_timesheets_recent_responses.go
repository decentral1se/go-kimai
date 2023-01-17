// Code generated by go-swagger; DO NOT EDIT.

package timesheet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPITimesheetsRecentReader is a Reader for the GetAPITimesheetsRecent structure.
type GetAPITimesheetsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITimesheetsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITimesheetsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPITimesheetsRecentOK creates a GetAPITimesheetsRecentOK with default headers values
func NewGetAPITimesheetsRecentOK() *GetAPITimesheetsRecentOK {
	return &GetAPITimesheetsRecentOK{}
}

/*
GetAPITimesheetsRecentOK describes a response with status code 200, with default header values.

Returns the collection of recent user activities (always the latest entry of a unique working set grouped by customer, project and activity)
*/
type GetAPITimesheetsRecentOK struct {
	Payload []*models.TimesheetCollectionExpanded
}

// IsSuccess returns true when this get Api timesheets recent o k response has a 2xx status code
func (o *GetAPITimesheetsRecentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api timesheets recent o k response has a 3xx status code
func (o *GetAPITimesheetsRecentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api timesheets recent o k response has a 4xx status code
func (o *GetAPITimesheetsRecentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api timesheets recent o k response has a 5xx status code
func (o *GetAPITimesheetsRecentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api timesheets recent o k response a status code equal to that given
func (o *GetAPITimesheetsRecentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPITimesheetsRecentOK) Error() string {
	return fmt.Sprintf("[GET /api/timesheets/recent][%d] getApiTimesheetsRecentOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsRecentOK) String() string {
	return fmt.Sprintf("[GET /api/timesheets/recent][%d] getApiTimesheetsRecentOK  %+v", 200, o.Payload)
}

func (o *GetAPITimesheetsRecentOK) GetPayload() []*models.TimesheetCollectionExpanded {
	return o.Payload
}

func (o *GetAPITimesheetsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
