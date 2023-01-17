// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIActivitiesReader is a Reader for the GetAPIActivities structure.
type GetAPIActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIActivitiesOK creates a GetAPIActivitiesOK with default headers values
func NewGetAPIActivitiesOK() *GetAPIActivitiesOK {
	return &GetAPIActivitiesOK{}
}

/*
GetAPIActivitiesOK describes a response with status code 200, with default header values.

Returns a collection of activity entities
*/
type GetAPIActivitiesOK struct {
	Payload []*models.ActivityCollection
}

// IsSuccess returns true when this get Api activities o k response has a 2xx status code
func (o *GetAPIActivitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api activities o k response has a 3xx status code
func (o *GetAPIActivitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api activities o k response has a 4xx status code
func (o *GetAPIActivitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api activities o k response has a 5xx status code
func (o *GetAPIActivitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api activities o k response a status code equal to that given
func (o *GetAPIActivitiesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /api/activities][%d] getApiActivitiesOK  %+v", 200, o.Payload)
}

func (o *GetAPIActivitiesOK) String() string {
	return fmt.Sprintf("[GET /api/activities][%d] getApiActivitiesOK  %+v", 200, o.Payload)
}

func (o *GetAPIActivitiesOK) GetPayload() []*models.ActivityCollection {
	return o.Payload
}

func (o *GetAPIActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
