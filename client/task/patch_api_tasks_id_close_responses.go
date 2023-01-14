// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// PatchAPITasksIDCloseReader is a Reader for the PatchAPITasksIDClose structure.
type PatchAPITasksIDCloseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITasksIDCloseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITasksIDCloseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITasksIDCloseOK creates a PatchAPITasksIDCloseOK with default headers values
func NewPatchAPITasksIDCloseOK() *PatchAPITasksIDCloseOK {
	return &PatchAPITasksIDCloseOK{}
}

/*
PatchAPITasksIDCloseOK describes a response with status code 200, with default header values.

Returns the closed task
*/
type PatchAPITasksIDCloseOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this patch Api tasks Id close o k response has a 2xx status code
func (o *PatchAPITasksIDCloseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api tasks Id close o k response has a 3xx status code
func (o *PatchAPITasksIDCloseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api tasks Id close o k response has a 4xx status code
func (o *PatchAPITasksIDCloseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api tasks Id close o k response has a 5xx status code
func (o *PatchAPITasksIDCloseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api tasks Id close o k response a status code equal to that given
func (o *PatchAPITasksIDCloseOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITasksIDCloseOK) Error() string {
	return fmt.Sprintf("[PATCH /api/tasks/{id}/close][%d] patchApiTasksIdCloseOK  %+v", 200, o.Payload)
}

func (o *PatchAPITasksIDCloseOK) String() string {
	return fmt.Sprintf("[PATCH /api/tasks/{id}/close][%d] patchApiTasksIdCloseOK  %+v", 200, o.Payload)
}

func (o *PatchAPITasksIDCloseOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *PatchAPITasksIDCloseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}