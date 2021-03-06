// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// GetContractPoIReader is a Reader for the GetContractPoI structure.
type GetContractPoIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractPoIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContractPoIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetContractPoIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetContractPoINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractPoIOK creates a GetContractPoIOK with default headers values
func NewGetContractPoIOK() *GetContractPoIOK {
	return &GetContractPoIOK{}
}

/*GetContractPoIOK handles this case with default header values.

Successful operation
*/
type GetContractPoIOK struct {
	Payload *models.PoI
}

func (o *GetContractPoIOK) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/poi][%d] getContractPoIOK  %+v", 200, o.Payload)
}

func (o *GetContractPoIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PoI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractPoIBadRequest creates a GetContractPoIBadRequest with default headers values
func NewGetContractPoIBadRequest() *GetContractPoIBadRequest {
	return &GetContractPoIBadRequest{}
}

/*GetContractPoIBadRequest handles this case with default header values.

Invalid contract key
*/
type GetContractPoIBadRequest struct {
	Payload *models.Error
}

func (o *GetContractPoIBadRequest) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/poi][%d] getContractPoIBadRequest  %+v", 400, o.Payload)
}

func (o *GetContractPoIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractPoINotFound creates a GetContractPoINotFound with default headers values
func NewGetContractPoINotFound() *GetContractPoINotFound {
	return &GetContractPoINotFound{}
}

/*GetContractPoINotFound handles this case with default header values.

Contract not found
*/
type GetContractPoINotFound struct {
	Payload *models.Error
}

func (o *GetContractPoINotFound) Error() string {
	return fmt.Sprintf("[GET /contracts/{pubkey}/poi][%d] getContractPoINotFound  %+v", 404, o.Payload)
}

func (o *GetContractPoINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
