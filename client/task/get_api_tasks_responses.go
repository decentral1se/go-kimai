// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPITasksReader is a Reader for the GetAPITasks structure.
type GetAPITasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPITasksOK creates a GetAPITasksOK with default headers values
func NewGetAPITasksOK() *GetAPITasksOK {
	return &GetAPITasksOK{}
}

/*
GetAPITasksOK describes a response with status code 200, with default header values.

Returns a collection of task entities
*/
type GetAPITasksOK struct {
	Payload []*models.Task
}

// IsSuccess returns true when this get Api tasks o k response has a 2xx status code
func (o *GetAPITasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api tasks o k response has a 3xx status code
func (o *GetAPITasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api tasks o k response has a 4xx status code
func (o *GetAPITasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api tasks o k response has a 5xx status code
func (o *GetAPITasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api tasks o k response a status code equal to that given
func (o *GetAPITasksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPITasksOK) Error() string {
	return fmt.Sprintf("[GET /api/tasks][%d] getApiTasksOK  %+v", 200, o.Payload)
}

func (o *GetAPITasksOK) String() string {
	return fmt.Sprintf("[GET /api/tasks][%d] getApiTasksOK  %+v", 200, o.Payload)
}

func (o *GetAPITasksOK) GetPayload() []*models.Task {
	return o.Payload
}

func (o *GetAPITasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
