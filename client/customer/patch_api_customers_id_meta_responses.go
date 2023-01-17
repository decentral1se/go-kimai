// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/decentral1se/go-kimai/models"
)

// PatchAPICustomersIDMetaReader is a Reader for the PatchAPICustomersIDMeta structure.
type PatchAPICustomersIDMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPICustomersIDMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPICustomersIDMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPICustomersIDMetaOK creates a PatchAPICustomersIDMetaOK with default headers values
func NewPatchAPICustomersIDMetaOK() *PatchAPICustomersIDMetaOK {
	return &PatchAPICustomersIDMetaOK{}
}

/*
PatchAPICustomersIDMetaOK describes a response with status code 200, with default header values.

Sets the value of an existing/configured meta-field. You cannot create unknown meta-fields, if the given name is not a configured meta-field, this will return an exception.
*/
type PatchAPICustomersIDMetaOK struct {
	Payload *models.CustomerEntity
}

// IsSuccess returns true when this patch Api customers Id meta o k response has a 2xx status code
func (o *PatchAPICustomersIDMetaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api customers Id meta o k response has a 3xx status code
func (o *PatchAPICustomersIDMetaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api customers Id meta o k response has a 4xx status code
func (o *PatchAPICustomersIDMetaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api customers Id meta o k response has a 5xx status code
func (o *PatchAPICustomersIDMetaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api customers Id meta o k response a status code equal to that given
func (o *PatchAPICustomersIDMetaOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPICustomersIDMetaOK) Error() string {
	return fmt.Sprintf("[PATCH /api/customers/{id}/meta][%d] patchApiCustomersIdMetaOK  %+v", 200, o.Payload)
}

func (o *PatchAPICustomersIDMetaOK) String() string {
	return fmt.Sprintf("[PATCH /api/customers/{id}/meta][%d] patchApiCustomersIdMetaOK  %+v", 200, o.Payload)
}

func (o *PatchAPICustomersIDMetaOK) GetPayload() *models.CustomerEntity {
	return o.Payload
}

func (o *PatchAPICustomersIDMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PatchAPICustomersIDMetaBody patch API customers ID meta body
swagger:model PatchAPICustomersIDMetaBody
*/
type PatchAPICustomersIDMetaBody struct {

	// The meta-field name
	// Required: true
	Name *string `json:"name"`

	// The meta-field value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this patch API customers ID meta body
func (o *PatchAPICustomersIDMetaBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchAPICustomersIDMetaBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PatchAPICustomersIDMetaBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch API customers ID meta body based on context it is used
func (o *PatchAPICustomersIDMetaBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchAPICustomersIDMetaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchAPICustomersIDMetaBody) UnmarshalBinary(b []byte) error {
	var res PatchAPICustomersIDMetaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
