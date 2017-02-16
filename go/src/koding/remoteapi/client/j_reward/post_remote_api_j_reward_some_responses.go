package j_reward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJRewardSomeReader is a Reader for the PostRemoteAPIJRewardSome structure.
type PostRemoteAPIJRewardSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJRewardSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJRewardSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJRewardSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJRewardSomeOK creates a PostRemoteAPIJRewardSomeOK with default headers values
func NewPostRemoteAPIJRewardSomeOK() *PostRemoteAPIJRewardSomeOK {
	return &PostRemoteAPIJRewardSomeOK{}
}

/*PostRemoteAPIJRewardSomeOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJRewardSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJRewardSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.some][%d] postRemoteApiJRewardSomeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJRewardSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJRewardSomeUnauthorized creates a PostRemoteAPIJRewardSomeUnauthorized with default headers values
func NewPostRemoteAPIJRewardSomeUnauthorized() *PostRemoteAPIJRewardSomeUnauthorized {
	return &PostRemoteAPIJRewardSomeUnauthorized{}
}

/*PostRemoteAPIJRewardSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJRewardSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJRewardSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JReward.some][%d] postRemoteApiJRewardSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJRewardSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
