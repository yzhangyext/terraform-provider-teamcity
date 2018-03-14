// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteRoleDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteRoleDefault creates a DeleteRoleDefault with default headers values
func NewDeleteRoleDefault(code int) *DeleteRoleDefault {
	return &DeleteRoleDefault{
		_statusCode: code,
	}
}

/*DeleteRoleDefault handles this case with default header values.

successful operation
*/
type DeleteRoleDefault struct {
	_statusCode int
}

// Code gets the status code for the delete role default response
func (o *DeleteRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/users/{userLocator}/roles/{roleId}/{scope}][%d] deleteRole default ", o._statusCode)
}

func (o *DeleteRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
