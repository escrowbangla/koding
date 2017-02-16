package j_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJSnapshotSomeReader is a Reader for the PostRemoteAPIJSnapshotSome structure.
type PostRemoteAPIJSnapshotSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJSnapshotSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJSnapshotSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJSnapshotSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJSnapshotSomeOK creates a PostRemoteAPIJSnapshotSomeOK with default headers values
func NewPostRemoteAPIJSnapshotSomeOK() *PostRemoteAPIJSnapshotSomeOK {
	return &PostRemoteAPIJSnapshotSomeOK{}
}

/*PostRemoteAPIJSnapshotSomeOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJSnapshotSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJSnapshotSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JSnapshot.some][%d] postRemoteApiJSnapshotSomeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJSnapshotSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJSnapshotSomeUnauthorized creates a PostRemoteAPIJSnapshotSomeUnauthorized with default headers values
func NewPostRemoteAPIJSnapshotSomeUnauthorized() *PostRemoteAPIJSnapshotSomeUnauthorized {
	return &PostRemoteAPIJSnapshotSomeUnauthorized{}
}

/*PostRemoteAPIJSnapshotSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJSnapshotSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJSnapshotSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JSnapshot.some][%d] postRemoteApiJSnapshotSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJSnapshotSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
