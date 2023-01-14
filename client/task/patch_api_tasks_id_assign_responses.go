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

// PatchAPITasksIDAssignReader is a Reader for the PatchAPITasksIDAssign structure.
type PatchAPITasksIDAssignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITasksIDAssignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPITasksIDAssignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPITasksIDAssignOK creates a PatchAPITasksIDAssignOK with default headers values
func NewPatchAPITasksIDAssignOK() *PatchAPITasksIDAssignOK {
	return &PatchAPITasksIDAssignOK{}
}

/*
PatchAPITasksIDAssignOK describes a response with status code 200, with default header values.

Returns the updated task
*/
type PatchAPITasksIDAssignOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this patch Api tasks Id assign o k response has a 2xx status code
func (o *PatchAPITasksIDAssignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api tasks Id assign o k response has a 3xx status code
func (o *PatchAPITasksIDAssignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api tasks Id assign o k response has a 4xx status code
func (o *PatchAPITasksIDAssignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api tasks Id assign o k response has a 5xx status code
func (o *PatchAPITasksIDAssignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api tasks Id assign o k response a status code equal to that given
func (o *PatchAPITasksIDAssignOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPITasksIDAssignOK) Error() string {
	return fmt.Sprintf("[PATCH /api/tasks/{id}/assign][%d] patchApiTasksIdAssignOK  %+v", 200, o.Payload)
}

func (o *PatchAPITasksIDAssignOK) String() string {
	return fmt.Sprintf("[PATCH /api/tasks/{id}/assign][%d] patchApiTasksIdAssignOK  %+v", 200, o.Payload)
}

func (o *PatchAPITasksIDAssignOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *PatchAPITasksIDAssignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}