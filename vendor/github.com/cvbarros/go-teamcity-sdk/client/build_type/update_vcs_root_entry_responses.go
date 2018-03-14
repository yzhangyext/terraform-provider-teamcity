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

// UpdateVcsRootEntryReader is a Reader for the UpdateVcsRootEntry structure.
type UpdateVcsRootEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVcsRootEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateVcsRootEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVcsRootEntryOK creates a UpdateVcsRootEntryOK with default headers values
func NewUpdateVcsRootEntryOK() *UpdateVcsRootEntryOK {
	return &UpdateVcsRootEntryOK{}
}

/*UpdateVcsRootEntryOK handles this case with default header values.

successful operation
*/
type UpdateVcsRootEntryOK struct {
	Payload *models.VcsRootEntry
}

func (o *UpdateVcsRootEntryOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/vcs-root-entries/{vcsRootLocator}][%d] updateVcsRootEntryOK  %+v", 200, o.Payload)
}

func (o *UpdateVcsRootEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
