// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewProjectServiceSetIamPolicyParams creates a new ProjectServiceSetIamPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectServiceSetIamPolicyParams() *ProjectServiceSetIamPolicyParams {
	return &ProjectServiceSetIamPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectServiceSetIamPolicyParamsWithTimeout creates a new ProjectServiceSetIamPolicyParams object
// with the ability to set a timeout on a request.
func NewProjectServiceSetIamPolicyParamsWithTimeout(timeout time.Duration) *ProjectServiceSetIamPolicyParams {
	return &ProjectServiceSetIamPolicyParams{
		timeout: timeout,
	}
}

// NewProjectServiceSetIamPolicyParamsWithContext creates a new ProjectServiceSetIamPolicyParams object
// with the ability to set a context for a request.
func NewProjectServiceSetIamPolicyParamsWithContext(ctx context.Context) *ProjectServiceSetIamPolicyParams {
	return &ProjectServiceSetIamPolicyParams{
		Context: ctx,
	}
}

// NewProjectServiceSetIamPolicyParamsWithHTTPClient creates a new ProjectServiceSetIamPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectServiceSetIamPolicyParamsWithHTTPClient(client *http.Client) *ProjectServiceSetIamPolicyParams {
	return &ProjectServiceSetIamPolicyParams{
		HTTPClient: client,
	}
}

/*
ProjectServiceSetIamPolicyParams contains all the parameters to send to the API endpoint

	for the project service set iam policy operation.

	Typically these are written to a http.Request.
*/
type ProjectServiceSetIamPolicyParams struct {

	// Body.
	Body ProjectServiceSetIamPolicyBody

	/* ID.

	   ID is the identifier of the project.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project service set iam policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceSetIamPolicyParams) WithDefaults() *ProjectServiceSetIamPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project service set iam policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectServiceSetIamPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) WithTimeout(timeout time.Duration) *ProjectServiceSetIamPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) WithContext(ctx context.Context) *ProjectServiceSetIamPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) WithHTTPClient(client *http.Client) *ProjectServiceSetIamPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) WithBody(body ProjectServiceSetIamPolicyBody) *ProjectServiceSetIamPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) SetBody(body ProjectServiceSetIamPolicyBody) {
	o.Body = body
}

// WithID adds the id to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) WithID(id string) *ProjectServiceSetIamPolicyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project service set iam policy params
func (o *ProjectServiceSetIamPolicyParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectServiceSetIamPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
