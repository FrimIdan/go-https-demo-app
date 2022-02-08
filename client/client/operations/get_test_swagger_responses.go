// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"wwwin-github.cisco.com/eti/swagger-authentication-test/client/models"
)

// GetTestReader is a Reader for the GetTest structure.
type GetTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTestOK creates a GetTestOK with default headers values
func NewGetTestOK() *GetTestOK {
	return &GetTestOK{}
}

/* GetTestOK describes a response with status code 200, with default header values.

OK
*/
type GetTestOK struct {
	Payload *models.OK
}

func (o *GetTestOK) Error() string {
	return fmt.Sprintf("[GET /test][%d] getTestOK  %+v", 200, o.Payload)
}
func (o *GetTestOK) GetPayload() *models.OK {
	return o.Payload
}

func (o *GetTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OK)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTestUnauthorized creates a GetTestUnauthorized with default headers values
func NewGetTestUnauthorized() *GetTestUnauthorized {
	return &GetTestUnauthorized{}
}

/* GetTestUnauthorized describes a response with status code 401, with default header values.

Not authenticated
*/
type GetTestUnauthorized struct {
}

func (o *GetTestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /test][%d] getTestUnauthorized ", 401)
}

func (o *GetTestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTestForbidden creates a GetTestForbidden with default headers values
func NewGetTestForbidden() *GetTestForbidden {
	return &GetTestForbidden{}
}

/* GetTestForbidden describes a response with status code 403, with default header values.

Access token does not have the required scope
*/
type GetTestForbidden struct {
}

func (o *GetTestForbidden) Error() string {
	return fmt.Sprintf("[GET /test][%d] getTestForbidden ", 403)
}

func (o *GetTestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}