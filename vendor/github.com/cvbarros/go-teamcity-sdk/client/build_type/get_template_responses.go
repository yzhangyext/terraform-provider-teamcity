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

// GetTemplateReader is a Reader for the GetTemplate structure.
type GetTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTemplateOK creates a GetTemplateOK with default headers values
func NewGetTemplateOK() *GetTemplateOK {
	return &GetTemplateOK{}
}

/*GetTemplateOK handles this case with default header values.

successful operation
*/
type GetTemplateOK struct {
	Payload *models.BuildType
}

func (o *GetTemplateOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/templates/{templateLocator}][%d] getTemplateOK  %+v", 200, o.Payload)
}

func (o *GetTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
