// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate hashicorp cloud vault 20201125 get available templates response template
//
// swagger:model hashicorp.cloud.vault_20201125.GetAvailableTemplatesResponse.Template
type HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate struct {

	// id
	ID string `json:"id,omitempty"`

	// is beta
	IsBeta bool `json:"is_beta,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get available templates response template
func (m *HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 get available templates response template based on context it is used
func (m *HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetAvailableTemplatesResponseTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}