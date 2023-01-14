// Code generated by go-swagger; DO NOT EDIT.

package default_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"decentral1se/go-kimai/models"
)

// GetAPIConfigI18nReader is a Reader for the GetAPIConfigI18n structure.
type GetAPIConfigI18nReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIConfigI18nReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIConfigI18nOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIConfigI18nOK creates a GetAPIConfigI18nOK with default headers values
func NewGetAPIConfigI18nOK() *GetAPIConfigI18nOK {
	return &GetAPIConfigI18nOK{}
}

/*
GetAPIConfigI18nOK describes a response with status code 200, with default header values.

Returns the locale specific configurations for this user
*/
type GetAPIConfigI18nOK struct {
	Payload *models.I18nConfig
}

// IsSuccess returns true when this get Api config i18n o k response has a 2xx status code
func (o *GetAPIConfigI18nOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api config i18n o k response has a 3xx status code
func (o *GetAPIConfigI18nOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api config i18n o k response has a 4xx status code
func (o *GetAPIConfigI18nOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api config i18n o k response has a 5xx status code
func (o *GetAPIConfigI18nOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api config i18n o k response a status code equal to that given
func (o *GetAPIConfigI18nOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIConfigI18nOK) Error() string {
	return fmt.Sprintf("[GET /api/config/i18n][%d] getApiConfigI18nOK  %+v", 200, o.Payload)
}

func (o *GetAPIConfigI18nOK) String() string {
	return fmt.Sprintf("[GET /api/config/i18n][%d] getApiConfigI18nOK  %+v", 200, o.Payload)
}

func (o *GetAPIConfigI18nOK) GetPayload() *models.I18nConfig {
	return o.Payload
}

func (o *GetAPIConfigI18nOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.I18nConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
