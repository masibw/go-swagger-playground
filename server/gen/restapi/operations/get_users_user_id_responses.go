// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"go-swagger-playground/server/gen/models"
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetUsersUserIDOKCode is the HTTP code returned for type GetUsersUserIDOK
const GetUsersUserIDOKCode int = 200

/*GetUsersUserIDOK User Found

swagger:response getUsersUserIdOK
*/
type GetUsersUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUsersUserIDOK creates GetUsersUserIDOK with default headers values
func NewGetUsersUserIDOK() *GetUsersUserIDOK {

	return &GetUsersUserIDOK{}
}

// WithPayload adds the payload to the get users user Id o k response
func (o *GetUsersUserIDOK) WithPayload(payload *models.User) *GetUsersUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users user Id o k response
func (o *GetUsersUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersUserIDNotFoundCode is the HTTP code returned for type GetUsersUserIDNotFound
const GetUsersUserIDNotFoundCode int = 404

/*GetUsersUserIDNotFound User Not Found

swagger:response getUsersUserIdNotFound
*/
type GetUsersUserIDNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetUsersUserIDNotFound creates GetUsersUserIDNotFound with default headers values
func NewGetUsersUserIDNotFound() *GetUsersUserIDNotFound {

	return &GetUsersUserIDNotFound{}
}

// WithPayload adds the payload to the get users user Id not found response
func (o *GetUsersUserIDNotFound) WithPayload(payload interface{}) *GetUsersUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users user Id not found response
func (o *GetUsersUserIDNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
