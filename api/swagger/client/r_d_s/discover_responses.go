// Code generated by go-swagger; DO NOT EDIT.

package r_d_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm-managed/api/swagger/models"
)

// DiscoverReader is a Reader for the Discover structure.
type DiscoverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDiscoverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDiscoverOK creates a DiscoverOK with default headers values
func NewDiscoverOK() *DiscoverOK {
	return &DiscoverOK{}
}

/*DiscoverOK handles this case with default header values.

(empty)
*/
type DiscoverOK struct {
	Payload *models.APIRDSDiscoverResponse
}

func (o *DiscoverOK) Error() string {
	return fmt.Sprintf("[POST /v0/rds/discover][%d] discoverOK  %+v", 200, o.Payload)
}

func (o *DiscoverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIRDSDiscoverResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
