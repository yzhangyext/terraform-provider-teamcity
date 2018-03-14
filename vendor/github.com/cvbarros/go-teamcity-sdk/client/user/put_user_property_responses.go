// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutUserPropertyReader is a Reader for the PutUserProperty structure.
type PutUserPropertyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserPropertyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutUserPropertyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutUserPropertyOK creates a PutUserPropertyOK with default headers values
func NewPutUserPropertyOK() *PutUserPropertyOK {
	return &PutUserPropertyOK{}
}

/*PutUserPropertyOK handles this case with default header values.

successful operation
*/
type PutUserPropertyOK struct {
	Payload string
}

func (o *PutUserPropertyOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/users/{userLocator}/properties/{name}][%d] putUserPropertyOK  %+v", 200, o.Payload)
}

func (o *PutUserPropertyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
