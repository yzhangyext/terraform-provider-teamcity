// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ReplaceRolesReader is a Reader for the ReplaceRoles structure.
type ReplaceRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceRolesOK creates a ReplaceRolesOK with default headers values
func NewReplaceRolesOK() *ReplaceRolesOK {
	return &ReplaceRolesOK{}
}

/*ReplaceRolesOK handles this case with default header values.

successful operation
*/
type ReplaceRolesOK struct {
	Payload *models.Roles
}

func (o *ReplaceRolesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/users/{userLocator}/roles][%d] replaceRolesOK  %+v", 200, o.Payload)
}

func (o *ReplaceRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Roles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
