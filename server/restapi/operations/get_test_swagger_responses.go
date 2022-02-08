// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"wwwin-github.cisco.com/eti/swagger-authentication-test/server/models"
)

// GetTestOKCode is the HTTP code returned for type GetTestOK
const GetTestOKCode int = 200

/*GetTestOK OK

swagger:response getTestOK
*/
type GetTestOK struct {

	/*
	  In: Body
	*/
	Payload *models.OK `json:"body,omitempty"`
}

// NewGetTestOK creates GetTestOK with default headers values
func NewGetTestOK() *GetTestOK {

	return &GetTestOK{}
}

// WithPayload adds the payload to the get test o k response
func (o *GetTestOK) WithPayload(payload *models.OK) *GetTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get test o k response
func (o *GetTestOK) SetPayload(payload *models.OK) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTestUnauthorizedCode is the HTTP code returned for type GetTestUnauthorized
const GetTestUnauthorizedCode int = 401

/*GetTestUnauthorized Not authenticated

swagger:response getTestUnauthorized
*/
type GetTestUnauthorized struct {
}

// NewGetTestUnauthorized creates GetTestUnauthorized with default headers values
func NewGetTestUnauthorized() *GetTestUnauthorized {

	return &GetTestUnauthorized{}
}

// WriteResponse to the client
func (o *GetTestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetTestForbiddenCode is the HTTP code returned for type GetTestForbidden
const GetTestForbiddenCode int = 403

/*GetTestForbidden Access token does not have the required scope

swagger:response getTestForbidden
*/
type GetTestForbidden struct {
}

// NewGetTestForbidden creates GetTestForbidden with default headers values
func NewGetTestForbidden() *GetTestForbidden {

	return &GetTestForbidden{}
}

// WriteResponse to the client
func (o *GetTestForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
