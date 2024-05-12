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

// HashicorpCloudVault20201125InputClusterConfig hashicorp cloud vault 20201125 input cluster config
//
// swagger:model hashicorp.cloud.vault_20201125.InputClusterConfig
type HashicorpCloudVault20201125InputClusterConfig struct {

	// audit_log_export_config is the configuration settings for exporting Vault's log information
	AuditLogExportConfig *HashicorpCloudVault20201125ObservabilityConfig `json:"audit_log_export_config,omitempty"`

	// metrics_config is the configuration settings for exporting Vault's observability information
	MetricsConfig *HashicorpCloudVault20201125ObservabilityConfig `json:"metrics_config,omitempty"`

	// network_config is the network configuration for the cluster
	NetworkConfig *HashicorpCloudVault20201125InputNetworkConfig `json:"network_config,omitempty"`

	// Tier is the type of Vault cluster that should be provisioned
	Tier *HashicorpCloudVault20201125Tier `json:"tier,omitempty"`

	// vault_config is the Vault specific configuration
	VaultConfig *HashicorpCloudVault20201125VaultConfig `json:"vault_config,omitempty"`

	// vault_insights_config is the configuration for Vault Insights audit-log streaming
	VaultInsightsConfig *HashicorpCloudVault20201125InputVaultInsightsConfig `json:"vault_insights_config,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 input cluster config
func (m *HashicorpCloudVault20201125InputClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditLogExportConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultInsightsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateAuditLogExportConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogExportConfig) { // not required
		return nil
	}

	if m.AuditLogExportConfig != nil {
		if err := m.AuditLogExportConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit_log_export_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audit_log_export_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateMetricsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricsConfig) { // not required
		return nil
	}

	if m.MetricsConfig != nil {
		if err := m.MetricsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateNetworkConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkConfig) { // not required
		return nil
	}

	if m.NetworkConfig != nil {
		if err := m.NetworkConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateTier(formats strfmt.Registry) error {
	if swag.IsZero(m.Tier) { // not required
		return nil
	}

	if m.Tier != nil {
		if err := m.Tier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateVaultConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VaultConfig) { // not required
		return nil
	}

	if m.VaultConfig != nil {
		if err := m.VaultConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vault_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) validateVaultInsightsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.VaultInsightsConfig) { // not required
		return nil
	}

	if m.VaultInsightsConfig != nil {
		if err := m.VaultInsightsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_insights_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vault_insights_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 input cluster config based on the context it is used
func (m *HashicorpCloudVault20201125InputClusterConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogExportConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVaultConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVaultInsightsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateAuditLogExportConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AuditLogExportConfig != nil {

		if swag.IsZero(m.AuditLogExportConfig) { // not required
			return nil
		}

		if err := m.AuditLogExportConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audit_log_export_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audit_log_export_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateMetricsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.MetricsConfig != nil {

		if swag.IsZero(m.MetricsConfig) { // not required
			return nil
		}

		if err := m.MetricsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateNetworkConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfig != nil {

		if swag.IsZero(m.NetworkConfig) { // not required
			return nil
		}

		if err := m.NetworkConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateTier(ctx context.Context, formats strfmt.Registry) error {

	if m.Tier != nil {

		if swag.IsZero(m.Tier) { // not required
			return nil
		}

		if err := m.Tier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tier")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateVaultConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VaultConfig != nil {

		if swag.IsZero(m.VaultConfig) { // not required
			return nil
		}

		if err := m.VaultConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vault_config")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125InputClusterConfig) contextValidateVaultInsightsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.VaultInsightsConfig != nil {

		if swag.IsZero(m.VaultInsightsConfig) { // not required
			return nil
		}

		if err := m.VaultInsightsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault_insights_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vault_insights_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125InputClusterConfig) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125InputClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}