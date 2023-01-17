// Code generated by go-swagger; DO NOT EDIT.

package expense

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/decentral1se/go-kimai/models"
)

// GetAPIExpensesReader is a Reader for the GetAPIExpenses structure.
type GetAPIExpensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIExpensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIExpensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIExpensesOK creates a GetAPIExpensesOK with default headers values
func NewGetAPIExpensesOK() *GetAPIExpensesOK {
	return &GetAPIExpensesOK{}
}

/*
GetAPIExpensesOK describes a response with status code 200, with default header values.

Returns a collection of expense entities
*/
type GetAPIExpensesOK struct {
	Payload []*models.ExpenseEntity
}

// IsSuccess returns true when this get Api expenses o k response has a 2xx status code
func (o *GetAPIExpensesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api expenses o k response has a 3xx status code
func (o *GetAPIExpensesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api expenses o k response has a 4xx status code
func (o *GetAPIExpensesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api expenses o k response has a 5xx status code
func (o *GetAPIExpensesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api expenses o k response a status code equal to that given
func (o *GetAPIExpensesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIExpensesOK) Error() string {
	return fmt.Sprintf("[GET /api/expenses][%d] getApiExpensesOK  %+v", 200, o.Payload)
}

func (o *GetAPIExpensesOK) String() string {
	return fmt.Sprintf("[GET /api/expenses][%d] getApiExpensesOK  %+v", 200, o.Payload)
}

func (o *GetAPIExpensesOK) GetPayload() []*models.ExpenseEntity {
	return o.Payload
}

func (o *GetAPIExpensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
