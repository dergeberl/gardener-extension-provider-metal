// Code generated by go-swagger; DO NOT EDIT.

package size

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// ListSizesReader is a Reader for the ListSizes structure.
type ListSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSizesOK creates a ListSizesOK with default headers values
func NewListSizesOK() *ListSizesOK {
	return &ListSizesOK{}
}

/*ListSizesOK handles this case with default header values.

OK
*/
type ListSizesOK struct {
	Payload []*models.V1SizeResponse
}

func (o *ListSizesOK) Error() string {
	return fmt.Sprintf("[GET /v1/size][%d] listSizesOK  %+v", 200, o.Payload)
}

func (o *ListSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSizesDefault creates a ListSizesDefault with default headers values
func NewListSizesDefault(code int) *ListSizesDefault {
	return &ListSizesDefault{
		_statusCode: code,
	}
}

/*ListSizesDefault handles this case with default header values.

Error
*/
type ListSizesDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the list sizes default response
func (o *ListSizesDefault) Code() int {
	return o._statusCode
}

func (o *ListSizesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/size][%d] listSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
