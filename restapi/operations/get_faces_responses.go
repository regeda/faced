// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/regeda/faced/models"
)

// GetFacesOKCode is the HTTP code returned for type GetFacesOK
const GetFacesOKCode int = 200

/*GetFacesOK list of detected faces

swagger:response getFacesOK
*/
type GetFacesOK struct {

	/*
	  In: Body
	*/
	Payload *GetFacesOKBody `json:"body,omitempty"`
}

// NewGetFacesOK creates GetFacesOK with default headers values
func NewGetFacesOK() *GetFacesOK {

	return &GetFacesOK{}
}

// WithPayload adds the payload to the get faces o k response
func (o *GetFacesOK) WithPayload(payload *GetFacesOKBody) *GetFacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get faces o k response
func (o *GetFacesOK) SetPayload(payload *GetFacesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFacesDefault generic error response

swagger:response getFacesDefault
*/
type GetFacesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFacesDefault creates GetFacesDefault with default headers values
func NewGetFacesDefault(code int) *GetFacesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFacesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get faces default response
func (o *GetFacesDefault) WithStatusCode(code int) *GetFacesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get faces default response
func (o *GetFacesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get faces default response
func (o *GetFacesDefault) WithPayload(payload *models.Error) *GetFacesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get faces default response
func (o *GetFacesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFacesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}