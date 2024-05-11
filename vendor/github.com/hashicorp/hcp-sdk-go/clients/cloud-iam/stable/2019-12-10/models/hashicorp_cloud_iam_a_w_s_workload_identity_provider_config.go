// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamAWSWorkloadIdentityProviderConfig AWSWorkloadIdentityProviderConfig configures an AWS Workload Identity
// Provider.
//
// swagger:model hashicorp.cloud.iam.AWSWorkloadIdentityProviderConfig
type HashicorpCloudIamAWSWorkloadIdentityProviderConfig struct {

	// account_id is the AWS account ID allowed to exchange identities.
	AccountID string `json:"account_id,omitempty"`
}

// Validate validates this hashicorp cloud iam a w s workload identity provider config
func (m *HashicorpCloudIamAWSWorkloadIdentityProviderConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam a w s workload identity provider config based on context it is used
func (m *HashicorpCloudIamAWSWorkloadIdentityProviderConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamAWSWorkloadIdentityProviderConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamAWSWorkloadIdentityProviderConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamAWSWorkloadIdentityProviderConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
