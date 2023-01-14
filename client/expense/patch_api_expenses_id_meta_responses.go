// Code generated by go-swagger; DO NOT EDIT.

package expense

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

	"decentral1se/go-kimai/models"
)

// PatchAPIExpensesIDMetaReader is a Reader for the PatchAPIExpensesIDMeta structure.
type PatchAPIExpensesIDMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIExpensesIDMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIExpensesIDMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAPIExpensesIDMetaOK creates a PatchAPIExpensesIDMetaOK with default headers values
func NewPatchAPIExpensesIDMetaOK() *PatchAPIExpensesIDMetaOK {
	return &PatchAPIExpensesIDMetaOK{}
}

/*
PatchAPIExpensesIDMetaOK describes a response with status code 200, with default header values.

Sets the value of an existing/configured meta-field. You cannot create unknown meta-fields, if the given name is not a configured meta-field, this will return an exception.
*/
type PatchAPIExpensesIDMetaOK struct {
	Payload *models.ExpenseEntity
}

// IsSuccess returns true when this patch Api expenses Id meta o k response has a 2xx status code
func (o *PatchAPIExpensesIDMetaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api expenses Id meta o k response has a 3xx status code
func (o *PatchAPIExpensesIDMetaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api expenses Id meta o k response has a 4xx status code
func (o *PatchAPIExpensesIDMetaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api expenses Id meta o k response has a 5xx status code
func (o *PatchAPIExpensesIDMetaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api expenses Id meta o k response a status code equal to that given
func (o *PatchAPIExpensesIDMetaOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchAPIExpensesIDMetaOK) Error() string {
	return fmt.Sprintf("[PATCH /api/expenses/{id}/meta][%d] patchApiExpensesIdMetaOK  %+v", 200, o.Payload)
}

func (o *PatchAPIExpensesIDMetaOK) String() string {
	return fmt.Sprintf("[PATCH /api/expenses/{id}/meta][%d] patchApiExpensesIdMetaOK  %+v", 200, o.Payload)
}

func (o *PatchAPIExpensesIDMetaOK) GetPayload() *models.ExpenseEntity {
	return o.Payload
}

func (o *PatchAPIExpensesIDMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpenseEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PatchAPIExpensesIDMetaBody patch API expenses ID meta body
swagger:model PatchAPIExpensesIDMetaBody
*/
type PatchAPIExpensesIDMetaBody struct {

	// The meta-field name
	// Required: true
	Name *string `json:"name"`

	// The meta-field value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this patch API expenses ID meta body
func (o *PatchAPIExpensesIDMetaBody) Validate(formats strfmt.Registry) error {
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

func (o *PatchAPIExpensesIDMetaBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PatchAPIExpensesIDMetaBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch API expenses ID meta body based on context it is used
func (o *PatchAPIExpensesIDMetaBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchAPIExpensesIDMetaBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchAPIExpensesIDMetaBody) UnmarshalBinary(b []byte) error {
	var res PatchAPIExpensesIDMetaBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}