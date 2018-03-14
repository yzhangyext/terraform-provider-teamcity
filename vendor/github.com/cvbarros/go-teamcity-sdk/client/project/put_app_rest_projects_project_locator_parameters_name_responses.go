// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// PutAppRestProjectsProjectLocatorParametersNameReader is a Reader for the PutAppRestProjectsProjectLocatorParametersName structure.
type PutAppRestProjectsProjectLocatorParametersNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppRestProjectsProjectLocatorParametersNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAppRestProjectsProjectLocatorParametersNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAppRestProjectsProjectLocatorParametersNameOK creates a PutAppRestProjectsProjectLocatorParametersNameOK with default headers values
func NewPutAppRestProjectsProjectLocatorParametersNameOK() *PutAppRestProjectsProjectLocatorParametersNameOK {
	return &PutAppRestProjectsProjectLocatorParametersNameOK{}
}

/*PutAppRestProjectsProjectLocatorParametersNameOK handles this case with default header values.

successful operation
*/
type PutAppRestProjectsProjectLocatorParametersNameOK struct {
	Payload *models.Property
}

func (o *PutAppRestProjectsProjectLocatorParametersNameOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/parameters/{name}][%d] putAppRestProjectsProjectLocatorParametersNameOK  %+v", 200, o.Payload)
}

func (o *PutAppRestProjectsProjectLocatorParametersNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
