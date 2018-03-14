// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UnpinBuildReader is a Reader for the UnpinBuild structure.
type UnpinBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnpinBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewUnpinBuildDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewUnpinBuildDefault creates a UnpinBuildDefault with default headers values
func NewUnpinBuildDefault(code int) *UnpinBuildDefault {
	return &UnpinBuildDefault{
		_statusCode: code,
	}
}

/*UnpinBuildDefault handles this case with default header values.

successful operation
*/
type UnpinBuildDefault struct {
	_statusCode int
}

// Code gets the status code for the unpin build default response
func (o *UnpinBuildDefault) Code() int {
	return o._statusCode
}

func (o *UnpinBuildDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/builds/{buildLocator}/pin][%d] unpinBuild default ", o._statusCode)
}

func (o *UnpinBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
