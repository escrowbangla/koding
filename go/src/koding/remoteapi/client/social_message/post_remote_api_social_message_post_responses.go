package social_message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialMessagePostReader is a Reader for the PostRemoteAPISocialMessagePost structure.
type PostRemoteAPISocialMessagePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialMessagePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialMessagePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialMessagePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialMessagePostOK creates a PostRemoteAPISocialMessagePostOK with default headers values
func NewPostRemoteAPISocialMessagePostOK() *PostRemoteAPISocialMessagePostOK {
	return &PostRemoteAPISocialMessagePostOK{}
}

/*PostRemoteAPISocialMessagePostOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialMessagePostOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialMessagePostOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.post][%d] postRemoteApiSocialMessagePostOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialMessagePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialMessagePostUnauthorized creates a PostRemoteAPISocialMessagePostUnauthorized with default headers values
func NewPostRemoteAPISocialMessagePostUnauthorized() *PostRemoteAPISocialMessagePostUnauthorized {
	return &PostRemoteAPISocialMessagePostUnauthorized{}
}

/*PostRemoteAPISocialMessagePostUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialMessagePostUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialMessagePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialMessage.post][%d] postRemoteApiSocialMessagePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialMessagePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
