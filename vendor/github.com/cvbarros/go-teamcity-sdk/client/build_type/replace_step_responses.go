// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ReplaceStepReader is a Reader for the ReplaceStep structure.
type ReplaceStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceStepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceStepOK creates a ReplaceStepOK with default headers values
func NewReplaceStepOK() *ReplaceStepOK {
	return &ReplaceStepOK{}
}

/*ReplaceStepOK handles this case with default header values.

successful operation
*/
type ReplaceStepOK struct {
	Payload *models.Step
}

func (o *ReplaceStepOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/steps/{stepId}][%d] replaceStepOK  %+v", 200, o.Payload)
}

func (o *ReplaceStepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Step)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
