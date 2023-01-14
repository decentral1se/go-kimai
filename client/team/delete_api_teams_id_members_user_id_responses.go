// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// DeleteAPITeamsIDMembersUserIDReader is a Reader for the DeleteAPITeamsIDMembersUserID structure.
type DeleteAPITeamsIDMembersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPITeamsIDMembersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAPITeamsIDMembersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPITeamsIDMembersUserIDOK creates a DeleteAPITeamsIDMembersUserIDOK with default headers values
func NewDeleteAPITeamsIDMembersUserIDOK() *DeleteAPITeamsIDMembersUserIDOK {
	return &DeleteAPITeamsIDMembersUserIDOK{}
}

/*
DeleteAPITeamsIDMembersUserIDOK describes a response with status code 200, with default header values.

Removes a user from the team. The teamlead cannot be removed.
*/
type DeleteAPITeamsIDMembersUserIDOK struct {
	Payload *models.Team
}

// IsSuccess returns true when this delete Api teams Id members user Id o k response has a 2xx status code
func (o *DeleteAPITeamsIDMembersUserIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api teams Id members user Id o k response has a 3xx status code
func (o *DeleteAPITeamsIDMembersUserIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api teams Id members user Id o k response has a 4xx status code
func (o *DeleteAPITeamsIDMembersUserIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api teams Id members user Id o k response has a 5xx status code
func (o *DeleteAPITeamsIDMembersUserIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api teams Id members user Id o k response a status code equal to that given
func (o *DeleteAPITeamsIDMembersUserIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteAPITeamsIDMembersUserIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/teams/{id}/members/{userId}][%d] deleteApiTeamsIdMembersUserIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAPITeamsIDMembersUserIDOK) String() string {
	return fmt.Sprintf("[DELETE /api/teams/{id}/members/{userId}][%d] deleteApiTeamsIdMembersUserIdOK  %+v", 200, o.Payload)
}

func (o *DeleteAPITeamsIDMembersUserIDOK) GetPayload() *models.Team {
	return o.Payload
}

func (o *DeleteAPITeamsIDMembersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
