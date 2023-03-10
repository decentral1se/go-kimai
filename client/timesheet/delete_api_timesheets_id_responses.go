// Code generated by go-swagger; DO NOT EDIT.

package timesheet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPITimesheetsIDReader is a Reader for the DeleteAPITimesheetsID structure.
type DeleteAPITimesheetsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPITimesheetsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPITimesheetsIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPITimesheetsIDNoContent creates a DeleteAPITimesheetsIDNoContent with default headers values
func NewDeleteAPITimesheetsIDNoContent() *DeleteAPITimesheetsIDNoContent {
	return &DeleteAPITimesheetsIDNoContent{}
}

/*
DeleteAPITimesheetsIDNoContent describes a response with status code 204, with default header values.

Delete one timesheet record
*/
type DeleteAPITimesheetsIDNoContent struct {
}

// IsSuccess returns true when this delete Api timesheets Id no content response has a 2xx status code
func (o *DeleteAPITimesheetsIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api timesheets Id no content response has a 3xx status code
func (o *DeleteAPITimesheetsIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api timesheets Id no content response has a 4xx status code
func (o *DeleteAPITimesheetsIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api timesheets Id no content response has a 5xx status code
func (o *DeleteAPITimesheetsIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api timesheets Id no content response a status code equal to that given
func (o *DeleteAPITimesheetsIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteAPITimesheetsIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/timesheets/{id}][%d] deleteApiTimesheetsIdNoContent ", 204)
}

func (o *DeleteAPITimesheetsIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/timesheets/{id}][%d] deleteApiTimesheetsIdNoContent ", 204)
}

func (o *DeleteAPITimesheetsIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
