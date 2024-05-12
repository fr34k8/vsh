// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudIamListServicePrincipalsResponse ListServicePrincipalsResponse is the response message for listing the service
// principals at a given scope.
//
// swagger:model hashicorp.cloud.iam.ListServicePrincipalsResponse
type HashicorpCloudIamListServicePrincipalsResponse struct {

	// pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`

	// service_principals is the list of principals in the organization.
	ServicePrincipals []*HashicorpCloudIamServicePrincipal `json:"service_principals"`
}

// Validate validates this hashicorp cloud iam list service principals response
func (m *HashicorpCloudIamListServicePrincipalsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePrincipals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamListServicePrincipalsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudIamListServicePrincipalsResponse) validateServicePrincipals(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePrincipals) { // not required
		return nil
	}

	for i := 0; i < len(m.ServicePrincipals); i++ {
		if swag.IsZero(m.ServicePrincipals[i]) { // not required
			continue
		}

		if m.ServicePrincipals[i] != nil {
			if err := m.ServicePrincipals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_principals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_principals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam list service principals response based on the context it is used
func (m *HashicorpCloudIamListServicePrincipalsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServicePrincipals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamListServicePrincipalsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudIamListServicePrincipalsResponse) contextValidateServicePrincipals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServicePrincipals); i++ {

		if m.ServicePrincipals[i] != nil {

			if swag.IsZero(m.ServicePrincipals[i]) { // not required
				return nil
			}

			if err := m.ServicePrincipals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_principals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("service_principals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamListServicePrincipalsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamListServicePrincipalsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamListServicePrincipalsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}