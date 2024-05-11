// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamUpdateWebConsolePreferencesResponse UpdateWebConsolePreferencesRequest is the request to update a user principal's web portal preferences on the IAM service.
//
// swagger:model hashicorp.cloud.iam.UpdateWebConsolePreferencesResponse
type HashicorpCloudIamUpdateWebConsolePreferencesResponse struct {

	// web_portal_preferences are the user preferences for the HCP Web Portal
	WebPortalPreferences HashicorpCloudIamWebConsolePreferences `json:"web_portal_preferences,omitempty"`
}

// Validate validates this hashicorp cloud iam update web console preferences response
func (m *HashicorpCloudIamUpdateWebConsolePreferencesResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam update web console preferences response based on context it is used
func (m *HashicorpCloudIamUpdateWebConsolePreferencesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamUpdateWebConsolePreferencesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamUpdateWebConsolePreferencesResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamUpdateWebConsolePreferencesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
