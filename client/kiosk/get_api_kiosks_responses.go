// Code generated by go-swagger; DO NOT EDIT.

package kiosk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIKiosksReader is a Reader for the GetAPIKiosks structure.
type GetAPIKiosksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIKiosksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIKiosksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIKiosksOK creates a GetAPIKiosksOK with default headers values
func NewGetAPIKiosksOK() *GetAPIKiosksOK {
	return &GetAPIKiosksOK{}
}

/*
GetAPIKiosksOK describes a response with status code 200, with default header values.

Returns a collection of UserAuthCodes
*/
type GetAPIKiosksOK struct {
	Payload []*models.UserAuthCodes
}

// IsSuccess returns true when this get Api kiosks o k response has a 2xx status code
func (o *GetAPIKiosksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api kiosks o k response has a 3xx status code
func (o *GetAPIKiosksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api kiosks o k response has a 4xx status code
func (o *GetAPIKiosksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api kiosks o k response has a 5xx status code
func (o *GetAPIKiosksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api kiosks o k response a status code equal to that given
func (o *GetAPIKiosksOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIKiosksOK) Error() string {
	return fmt.Sprintf("[GET /api/kiosks][%d] getApiKiosksOK  %+v", 200, o.Payload)
}

func (o *GetAPIKiosksOK) String() string {
	return fmt.Sprintf("[GET /api/kiosks][%d] getApiKiosksOK  %+v", 200, o.Payload)
}

func (o *GetAPIKiosksOK) GetPayload() []*models.UserAuthCodes {
	return o.Payload
}

func (o *GetAPIKiosksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
