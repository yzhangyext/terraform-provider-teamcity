// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBuildChainOptimizationLogReader is a Reader for the GetBuildChainOptimizationLog structure.
type GetBuildChainOptimizationLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildChainOptimizationLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBuildChainOptimizationLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBuildChainOptimizationLogOK creates a GetBuildChainOptimizationLogOK with default headers values
func NewGetBuildChainOptimizationLogOK() *GetBuildChainOptimizationLogOK {
	return &GetBuildChainOptimizationLogOK{}
}

/*GetBuildChainOptimizationLogOK handles this case with default header values.

successful operation
*/
type GetBuildChainOptimizationLogOK struct {
	Payload string
}

func (o *GetBuildChainOptimizationLogOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/buildChainOptimizationLog/{buildLocator}][%d] getBuildChainOptimizationLogOK  %+v", 200, o.Payload)
}

func (o *GetBuildChainOptimizationLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
