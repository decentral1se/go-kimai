// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIUsersIDReader is a Reader for the GetAPIUsersID structure.
type GetAPIUsersIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIUsersIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIUsersIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIUsersIDOK creates a GetAPIUsersIDOK with default headers values
func NewGetAPIUsersIDOK() *GetAPIUsersIDOK {
	return &GetAPIUsersIDOK{}
}

/*
GetAPIUsersIDOK describes a response with status code 200, with default header values.

Return one user entity.
*/
type GetAPIUsersIDOK struct {
	Payload *models.UserEntity
}

// IsSuccess returns true when this get Api users Id o k response has a 2xx status code
func (o *GetAPIUsersIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api users Id o k response has a 3xx status code
func (o *GetAPIUsersIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api users Id o k response has a 4xx status code
func (o *GetAPIUsersIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api users Id o k response has a 5xx status code
func (o *GetAPIUsersIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api users Id o k response a status code equal to that given
func (o *GetAPIUsersIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIUsersIDOK) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}][%d] getApiUsersIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIUsersIDOK) String() string {
	return fmt.Sprintf("[GET /api/users/{id}][%d] getApiUsersIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIUsersIDOK) GetPayload() *models.UserEntity {
	return o.Payload
}

func (o *GetAPIUsersIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
