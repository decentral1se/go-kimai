// Code generated by go-swagger; DO NOT EDIT.

package meta_field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// GetAPIMetafieldsReader is a Reader for the GetAPIMetafields structure.
type GetAPIMetafieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIMetafieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIMetafieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIMetafieldsOK creates a GetAPIMetafieldsOK with default headers values
func NewGetAPIMetafieldsOK() *GetAPIMetafieldsOK {
	return &GetAPIMetafieldsOK{}
}

/*
GetAPIMetafieldsOK describes a response with status code 200, with default header values.

Returns a collection of meta-fields
*/
type GetAPIMetafieldsOK struct {
	Payload []*models.MetaFieldRule
}

// IsSuccess returns true when this get Api metafields o k response has a 2xx status code
func (o *GetAPIMetafieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api metafields o k response has a 3xx status code
func (o *GetAPIMetafieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api metafields o k response has a 4xx status code
func (o *GetAPIMetafieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api metafields o k response has a 5xx status code
func (o *GetAPIMetafieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api metafields o k response a status code equal to that given
func (o *GetAPIMetafieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIMetafieldsOK) Error() string {
	return fmt.Sprintf("[GET /api/metafields][%d] getApiMetafieldsOK  %+v", 200, o.Payload)
}

func (o *GetAPIMetafieldsOK) String() string {
	return fmt.Sprintf("[GET /api/metafields][%d] getApiMetafieldsOK  %+v", 200, o.Payload)
}

func (o *GetAPIMetafieldsOK) GetPayload() []*models.MetaFieldRule {
	return o.Payload
}

func (o *GetAPIMetafieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
