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

// PostAPIProjectsIDRatesReader is a Reader for the PostAPIProjectsIDRates structure.
type PostAPIProjectsIDRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIProjectsIDRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIProjectsIDRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIProjectsIDRatesOK creates a PostAPIProjectsIDRatesOK with default headers values
func NewPostAPIProjectsIDRatesOK() *PostAPIProjectsIDRatesOK {
	return &PostAPIProjectsIDRatesOK{}
}

/*
PostAPIProjectsIDRatesOK describes a response with status code 200, with default header values.

Returns the new created rate
*/
type PostAPIProjectsIDRatesOK struct {
	Payload *models.ProjectRate
}

// IsSuccess returns true when this post Api projects Id rates o k response has a 2xx status code
func (o *PostAPIProjectsIDRatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api projects Id rates o k response has a 3xx status code
func (o *PostAPIProjectsIDRatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api projects Id rates o k response has a 4xx status code
func (o *PostAPIProjectsIDRatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api projects Id rates o k response has a 5xx status code
func (o *PostAPIProjectsIDRatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api projects Id rates o k response a status code equal to that given
func (o *PostAPIProjectsIDRatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAPIProjectsIDRatesOK) Error() string {
	return fmt.Sprintf("[POST /api/projects/{id}/rates][%d] postApiProjectsIdRatesOK  %+v", 200, o.Payload)
}

func (o *PostAPIProjectsIDRatesOK) String() string {
	return fmt.Sprintf("[POST /api/projects/{id}/rates][%d] postApiProjectsIdRatesOK  %+v", 200, o.Payload)
}

func (o *PostAPIProjectsIDRatesOK) GetPayload() *models.ProjectRate {
	return o.Payload
}

func (o *PostAPIProjectsIDRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectRate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}