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

// ServeBuildTypeBuildsTagsReader is a Reader for the ServeBuildTypeBuildsTags structure.
type ServeBuildTypeBuildsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildTypeBuildsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeBuildTypeBuildsTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildTypeBuildsTagsOK creates a ServeBuildTypeBuildsTagsOK with default headers values
func NewServeBuildTypeBuildsTagsOK() *ServeBuildTypeBuildsTagsOK {
	return &ServeBuildTypeBuildsTagsOK{}
}

/*ServeBuildTypeBuildsTagsOK handles this case with default header values.

successful operation
*/
type ServeBuildTypeBuildsTagsOK struct {
	Payload *models.Tags
}

func (o *ServeBuildTypeBuildsTagsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/buildTags][%d] serveBuildTypeBuildsTagsOK  %+v", 200, o.Payload)
}

func (o *ServeBuildTypeBuildsTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
