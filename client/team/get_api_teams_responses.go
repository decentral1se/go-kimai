// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPITeamsReader is a Reader for the GetAPITeams structure.
type GetAPITeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPITeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPITeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPITeamsOK creates a GetAPITeamsOK with default headers values
func NewGetAPITeamsOK() *GetAPITeamsOK {
	return &GetAPITeamsOK{}
}

/*
GetAPITeamsOK describes a response with status code 200, with default header values.

Returns the collection of all existing teams
*/
type GetAPITeamsOK struct {
	Payload []*models.TeamCollection
}

// IsSuccess returns true when this get Api teams o k response has a 2xx status code
func (o *GetAPITeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api teams o k response has a 3xx status code
func (o *GetAPITeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api teams o k response has a 4xx status code
func (o *GetAPITeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api teams o k response has a 5xx status code
func (o *GetAPITeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api teams o k response a status code equal to that given
func (o *GetAPITeamsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPITeamsOK) Error() string {
	return fmt.Sprintf("[GET /api/teams][%d] getApiTeamsOK  %+v", 200, o.Payload)
}

func (o *GetAPITeamsOK) String() string {
	return fmt.Sprintf("[GET /api/teams][%d] getApiTeamsOK  %+v", 200, o.Payload)
}

func (o *GetAPITeamsOK) GetPayload() []*models.TeamCollection {
	return o.Payload
}

func (o *GetAPITeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
