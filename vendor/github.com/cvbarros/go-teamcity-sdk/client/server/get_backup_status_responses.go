// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBackupStatusReader is a Reader for the GetBackupStatus structure.
type GetBackupStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBackupStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBackupStatusOK creates a GetBackupStatusOK with default headers values
func NewGetBackupStatusOK() *GetBackupStatusOK {
	return &GetBackupStatusOK{}
}

/*GetBackupStatusOK handles this case with default header values.

successful operation
*/
type GetBackupStatusOK struct {
	Payload string
}

func (o *GetBackupStatusOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/server/backup][%d] getBackupStatusOK  %+v", 200, o.Payload)
}

func (o *GetBackupStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
