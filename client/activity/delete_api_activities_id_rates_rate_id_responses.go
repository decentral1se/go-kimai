// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPIActivitiesIDRatesRateIDReader is a Reader for the DeleteAPIActivitiesIDRatesRateID structure.
type DeleteAPIActivitiesIDRatesRateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIActivitiesIDRatesRateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIActivitiesIDRatesRateIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPIActivitiesIDRatesRateIDNoContent creates a DeleteAPIActivitiesIDRatesRateIDNoContent with default headers values
func NewDeleteAPIActivitiesIDRatesRateIDNoContent() *DeleteAPIActivitiesIDRatesRateIDNoContent {
	return &DeleteAPIActivitiesIDRatesRateIDNoContent{}
}

/*
DeleteAPIActivitiesIDRatesRateIDNoContent describes a response with status code 204, with default header values.

Returns no content: 204 on successful delete
*/
type DeleteAPIActivitiesIDRatesRateIDNoContent struct {
}

// IsSuccess returns true when this delete Api activities Id rates rate Id no content response has a 2xx status code
func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api activities Id rates rate Id no content response has a 3xx status code
func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api activities Id rates rate Id no content response has a 4xx status code
func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api activities Id rates rate Id no content response has a 5xx status code
func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api activities Id rates rate Id no content response a status code equal to that given
func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/activities/{id}/rates/{rateId}][%d] deleteApiActivitiesIdRatesRateIdNoContent ", 204)
}

func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/activities/{id}/rates/{rateId}][%d] deleteApiActivitiesIdRatesRateIdNoContent ", 204)
}

func (o *DeleteAPIActivitiesIDRatesRateIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
