// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ServeRootReader is a Reader for the ServeRoot structure.
type ServeRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeRootOK creates a ServeRootOK with default headers values
func NewServeRootOK() *ServeRootOK {
	return &ServeRootOK{}
}

/*ServeRootOK handles this case with default header values.

successful operation
*/
type ServeRootOK struct {
	Payload *models.VcsRoot
}

func (o *ServeRootOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-roots/{vcsRootLocator}][%d] serveRootOK  %+v", 200, o.Payload)
}

func (o *ServeRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRoot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
