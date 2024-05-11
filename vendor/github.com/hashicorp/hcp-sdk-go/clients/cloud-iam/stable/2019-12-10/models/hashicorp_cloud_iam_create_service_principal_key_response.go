// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamCreateServicePrincipalKeyResponse CreateServicePrincipalKeyResponse is the response message returned after creating
// a service principal key.
//
// swagger:model hashicorp.cloud.iam.CreateServicePrincipalKeyResponse
type HashicorpCloudIamCreateServicePrincipalKeyResponse struct {

	// client_secret is the secret part of the credential set modelled by the service
	// principal key. Its counterpart is the "client ID", which is part of the key
	// message. This is the only time that this client secret transits through our
	// systems, as after the creation of the service principal key it's forgotten.
	// Consumers of this service endpoint are expected to store the client secret
	// in order to use it later.
	ClientSecret string `json:"client_secret,omitempty"`

	// key is the service principal key that has just been created.
	Key *HashicorpCloudIamServicePrincipalKey `json:"key,omitempty"`
}

// Validate validates this hashicorp cloud iam create service principal key response
func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam create service principal key response based on the context it is used
func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if m.Key != nil {

		if swag.IsZero(m.Key) { // not required
			return nil
		}

		if err := m.Key.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalKeyResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateServicePrincipalKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
