// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// GetAPIProjectsIDReader is a Reader for the GetAPIProjectsID structure.
type GetAPIProjectsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIProjectsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIProjectsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIProjectsIDOK creates a GetAPIProjectsIDOK with default headers values
func NewGetAPIProjectsIDOK() *GetAPIProjectsIDOK {
	return &GetAPIProjectsIDOK{}
}

/*
GetAPIProjectsIDOK describes a response with status code 200, with default header values.

Returns one project entity
*/
type GetAPIProjectsIDOK struct {
	Payload *models.ProjectEntity
}

// IsSuccess returns true when this get Api projects Id o k response has a 2xx status code
func (o *GetAPIProjectsIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api projects Id o k response has a 3xx status code
func (o *GetAPIProjectsIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api projects Id o k response has a 4xx status code
func (o *GetAPIProjectsIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api projects Id o k response has a 5xx status code
func (o *GetAPIProjectsIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api projects Id o k response a status code equal to that given
func (o *GetAPIProjectsIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIProjectsIDOK) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}][%d] getApiProjectsIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIProjectsIDOK) String() string {
	return fmt.Sprintf("[GET /api/projects/{id}][%d] getApiProjectsIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIProjectsIDOK) GetPayload() *models.ProjectEntity {
	return o.Payload
}

func (o *GetAPIProjectsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
