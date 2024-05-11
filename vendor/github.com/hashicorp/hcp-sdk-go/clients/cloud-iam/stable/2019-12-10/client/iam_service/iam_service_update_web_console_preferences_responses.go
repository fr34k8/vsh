// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// IamServiceUpdateWebConsolePreferencesReader is a Reader for the IamServiceUpdateWebConsolePreferences structure.
type IamServiceUpdateWebConsolePreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IamServiceUpdateWebConsolePreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIamServiceUpdateWebConsolePreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIamServiceUpdateWebConsolePreferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIamServiceUpdateWebConsolePreferencesOK creates a IamServiceUpdateWebConsolePreferencesOK with default headers values
func NewIamServiceUpdateWebConsolePreferencesOK() *IamServiceUpdateWebConsolePreferencesOK {
	return &IamServiceUpdateWebConsolePreferencesOK{}
}

/*
IamServiceUpdateWebConsolePreferencesOK describes a response with status code 200, with default header values.

A successful response.
*/
type IamServiceUpdateWebConsolePreferencesOK struct {
	Payload *models.HashicorpCloudIamUpdateWebConsolePreferencesResponse
}

// IsSuccess returns true when this iam service update web console preferences o k response has a 2xx status code
func (o *IamServiceUpdateWebConsolePreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this iam service update web console preferences o k response has a 3xx status code
func (o *IamServiceUpdateWebConsolePreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this iam service update web console preferences o k response has a 4xx status code
func (o *IamServiceUpdateWebConsolePreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this iam service update web console preferences o k response has a 5xx status code
func (o *IamServiceUpdateWebConsolePreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this iam service update web console preferences o k response a status code equal to that given
func (o *IamServiceUpdateWebConsolePreferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the iam service update web console preferences o k response
func (o *IamServiceUpdateWebConsolePreferencesOK) Code() int {
	return 200
}

func (o *IamServiceUpdateWebConsolePreferencesOK) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/me/web-portal-preferences][%d] iamServiceUpdateWebConsolePreferencesOK  %+v", 200, o.Payload)
}

func (o *IamServiceUpdateWebConsolePreferencesOK) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/me/web-portal-preferences][%d] iamServiceUpdateWebConsolePreferencesOK  %+v", 200, o.Payload)
}

func (o *IamServiceUpdateWebConsolePreferencesOK) GetPayload() *models.HashicorpCloudIamUpdateWebConsolePreferencesResponse {
	return o.Payload
}

func (o *IamServiceUpdateWebConsolePreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamUpdateWebConsolePreferencesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIamServiceUpdateWebConsolePreferencesDefault creates a IamServiceUpdateWebConsolePreferencesDefault with default headers values
func NewIamServiceUpdateWebConsolePreferencesDefault(code int) *IamServiceUpdateWebConsolePreferencesDefault {
	return &IamServiceUpdateWebConsolePreferencesDefault{
		_statusCode: code,
	}
}

/*
IamServiceUpdateWebConsolePreferencesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IamServiceUpdateWebConsolePreferencesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// IsSuccess returns true when this iam service update web console preferences default response has a 2xx status code
func (o *IamServiceUpdateWebConsolePreferencesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this iam service update web console preferences default response has a 3xx status code
func (o *IamServiceUpdateWebConsolePreferencesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this iam service update web console preferences default response has a 4xx status code
func (o *IamServiceUpdateWebConsolePreferencesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this iam service update web console preferences default response has a 5xx status code
func (o *IamServiceUpdateWebConsolePreferencesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this iam service update web console preferences default response a status code equal to that given
func (o *IamServiceUpdateWebConsolePreferencesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the iam service update web console preferences default response
func (o *IamServiceUpdateWebConsolePreferencesDefault) Code() int {
	return o._statusCode
}

func (o *IamServiceUpdateWebConsolePreferencesDefault) Error() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/me/web-portal-preferences][%d] IamService_UpdateWebConsolePreferences default  %+v", o._statusCode, o.Payload)
}

func (o *IamServiceUpdateWebConsolePreferencesDefault) String() string {
	return fmt.Sprintf("[PUT /iam/2019-12-10/me/web-portal-preferences][%d] IamService_UpdateWebConsolePreferences default  %+v", o._statusCode, o.Payload)
}

func (o *IamServiceUpdateWebConsolePreferencesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *IamServiceUpdateWebConsolePreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
