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

// HashicorpCloudResourcemanagerOrganizationGetRoleResponse OrganizationGetRoleResponse see OrganizationService.GetRole
//
// swagger:model hashicorp.cloud.resourcemanager.OrganizationGetRoleResponse
type HashicorpCloudResourcemanagerOrganizationGetRoleResponse struct {

	// Role is the requested role.
	Role *HashicorpCloudResourcemanagerRole `json:"role,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager organization get role response
func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud resourcemanager organization get role response based on the context it is used
func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {

		if swag.IsZero(m.Role) { // not required
			return nil
		}

		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationGetRoleResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerOrganizationGetRoleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
