// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIProjectsIDRatesReader is a Reader for the GetAPIProjectsIDRates structure.
type GetAPIProjectsIDRatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIProjectsIDRatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIProjectsIDRatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIProjectsIDRatesOK creates a GetAPIProjectsIDRatesOK with default headers values
func NewGetAPIProjectsIDRatesOK() *GetAPIProjectsIDRatesOK {
	return &GetAPIProjectsIDRatesOK{}
}

/*
GetAPIProjectsIDRatesOK describes a response with status code 200, with default header values.

Returns a collection of project rate entities
*/
type GetAPIProjectsIDRatesOK struct {
	Payload []*models.ProjectRate
}

// IsSuccess returns true when this get Api projects Id rates o k response has a 2xx status code
func (o *GetAPIProjectsIDRatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api projects Id rates o k response has a 3xx status code
func (o *GetAPIProjectsIDRatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api projects Id rates o k response has a 4xx status code
func (o *GetAPIProjectsIDRatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api projects Id rates o k response has a 5xx status code
func (o *GetAPIProjectsIDRatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api projects Id rates o k response a status code equal to that given
func (o *GetAPIProjectsIDRatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIProjectsIDRatesOK) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}/rates][%d] getApiProjectsIdRatesOK  %+v", 200, o.Payload)
}

func (o *GetAPIProjectsIDRatesOK) String() string {
	return fmt.Sprintf("[GET /api/projects/{id}/rates][%d] getApiProjectsIdRatesOK  %+v", 200, o.Payload)
}

func (o *GetAPIProjectsIDRatesOK) GetPayload() []*models.ProjectRate {
	return o.Payload
}

func (o *GetAPIProjectsIDRatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
