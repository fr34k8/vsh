// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// DeletePluginReader is a Reader for the DeletePlugin structure.
type DeletePluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeletePluginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePluginOK creates a DeletePluginOK with default headers values
func NewDeletePluginOK() *DeletePluginOK {
	return &DeletePluginOK{}
}

/*
DeletePluginOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeletePluginOK struct {
	Payload models.HashicorpCloudVault20201125DeletePluginResponse
}

// IsSuccess returns true when this delete plugin o k response has a 2xx status code
func (o *DeletePluginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete plugin o k response has a 3xx status code
func (o *DeletePluginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete plugin o k response has a 4xx status code
func (o *DeletePluginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete plugin o k response has a 5xx status code
func (o *DeletePluginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete plugin o k response a status code equal to that given
func (o *DeletePluginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete plugin o k response
func (o *DeletePluginOK) Code() int {
	return 200
}

func (o *DeletePluginOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/plugin/delete][%d] deletePluginOK  %+v", 200, o.Payload)
}

func (o *DeletePluginOK) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/plugin/delete][%d] deletePluginOK  %+v", 200, o.Payload)
}

func (o *DeletePluginOK) GetPayload() models.HashicorpCloudVault20201125DeletePluginResponse {
	return o.Payload
}

func (o *DeletePluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginDefault creates a DeletePluginDefault with default headers values
func NewDeletePluginDefault(code int) *DeletePluginDefault {
	return &DeletePluginDefault{
		_statusCode: code,
	}
}

/*
DeletePluginDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeletePluginDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// IsSuccess returns true when this delete plugin default response has a 2xx status code
func (o *DeletePluginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete plugin default response has a 3xx status code
func (o *DeletePluginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete plugin default response has a 4xx status code
func (o *DeletePluginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete plugin default response has a 5xx status code
func (o *DeletePluginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete plugin default response a status code equal to that given
func (o *DeletePluginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete plugin default response
func (o *DeletePluginDefault) Code() int {
	return o._statusCode
}

func (o *DeletePluginDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/plugin/delete][%d] DeletePlugin default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePluginDefault) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/plugin/delete][%d] DeletePlugin default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePluginDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeletePluginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
