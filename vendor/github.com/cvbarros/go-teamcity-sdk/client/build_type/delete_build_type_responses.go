// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteBuildTypeReader is a Reader for the DeleteBuildType structure.
type DeleteBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteBuildTypeDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteBuildTypeDefault creates a DeleteBuildTypeDefault with default headers values
func NewDeleteBuildTypeDefault(code int) *DeleteBuildTypeDefault {
	return &DeleteBuildTypeDefault{
		_statusCode: code,
	}
}

/*DeleteBuildTypeDefault handles this case with default header values.

successful operation
*/
type DeleteBuildTypeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete build type default response
func (o *DeleteBuildTypeDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildTypeDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}][%d] deleteBuildType default ", o._statusCode)
}

func (o *DeleteBuildTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
