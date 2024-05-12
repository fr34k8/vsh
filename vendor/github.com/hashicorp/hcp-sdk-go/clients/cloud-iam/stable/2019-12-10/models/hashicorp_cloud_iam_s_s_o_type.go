// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudIamSSOType SSOType is the type of SSO.
//
//   - UNSET: UNSET is the default value.
//
// It should never be used at runtime for valid messages.
//   - SAML: SAML is a SAML connection.
//   - OIDC: OIDC is a OIDC connection stored in cloud-idp.
//
// swagger:model hashicorp.cloud.iam.SSOType
type HashicorpCloudIamSSOType string

func NewHashicorpCloudIamSSOType(value HashicorpCloudIamSSOType) *HashicorpCloudIamSSOType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudIamSSOType.
func (m HashicorpCloudIamSSOType) Pointer() *HashicorpCloudIamSSOType {
	return &m
}

const (

	// HashicorpCloudIamSSOTypeUNSET captures enum value "UNSET"
	HashicorpCloudIamSSOTypeUNSET HashicorpCloudIamSSOType = "UNSET"

	// HashicorpCloudIamSSOTypeSAML captures enum value "SAML"
	HashicorpCloudIamSSOTypeSAML HashicorpCloudIamSSOType = "SAML"

	// HashicorpCloudIamSSOTypeOIDC captures enum value "OIDC"
	HashicorpCloudIamSSOTypeOIDC HashicorpCloudIamSSOType = "OIDC"
)

// for schema
var hashicorpCloudIamSSOTypeEnum []interface{}

func init() {
	var res []HashicorpCloudIamSSOType
	if err := json.Unmarshal([]byte(`["UNSET","SAML","OIDC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudIamSSOTypeEnum = append(hashicorpCloudIamSSOTypeEnum, v)
	}
}

func (m HashicorpCloudIamSSOType) validateHashicorpCloudIamSSOTypeEnum(path, location string, value HashicorpCloudIamSSOType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudIamSSOTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud iam s s o type
func (m HashicorpCloudIamSSOType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudIamSSOTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud iam s s o type based on context it is used
func (m HashicorpCloudIamSSOType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}