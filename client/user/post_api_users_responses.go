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

// PostAPIUsersReader is a Reader for the PostAPIUsers structure.
type PostAPIUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIUsersOK creates a PostAPIUsersOK with default headers values
func NewPostAPIUsersOK() *PostAPIUsersOK {
	return &PostAPIUsersOK{}
}

/*
PostAPIUsersOK describes a response with status code 200, with default header values.

Returns the new created user
*/
type PostAPIUsersOK struct {
	Payload *models.UserEntity
}

// IsSuccess returns true when this post Api users o k response has a 2xx status code
func (o *PostAPIUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api users o k response has a 3xx status code
func (o *PostAPIUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api users o k response has a 4xx status code
func (o *PostAPIUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api users o k response has a 5xx status code
func (o *PostAPIUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api users o k response a status code equal to that given
func (o *PostAPIUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAPIUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/users][%d] postApiUsersOK  %+v", 200, o.Payload)
}

func (o *PostAPIUsersOK) String() string {
	return fmt.Sprintf("[POST /api/users][%d] postApiUsersOK  %+v", 200, o.Payload)
}

func (o *PostAPIUsersOK) GetPayload() *models.UserEntity {
	return o.Payload
}

func (o *PostAPIUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
