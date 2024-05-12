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

// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC hashicorp cloud vault 20201125 major version upgrade config maintenance window time window u t c
//
// swagger:model hashicorp.cloud.vault_20201125.MajorVersionUpgradeConfig.MaintenanceWindow.TimeWindowUTC
type HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC string

func NewHashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC(value HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) *HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC.
func (m HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) Pointer() *HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC {
	return &m
}

const (

	// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCTIMEWINDOWUTCINVALID captures enum value "TIME_WINDOW_UTC_INVALID"
	HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCTIMEWINDOWUTCINVALID HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC = "TIME_WINDOW_UTC_INVALID"

	// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW12AM4AM captures enum value "WINDOW_12AM_4AM"
	HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW12AM4AM HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC = "WINDOW_12AM_4AM"

	// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW6AM10AM captures enum value "WINDOW_6AM_10AM"
	HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW6AM10AM HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC = "WINDOW_6AM_10AM"

	// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW12PM4PM captures enum value "WINDOW_12PM_4PM"
	HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW12PM4PM HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC = "WINDOW_12PM_4PM"

	// HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW6PM10PM captures enum value "WINDOW_6PM_10PM"
	HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCWINDOW6PM10PM HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC = "WINDOW_6PM_10PM"
)

// for schema
var hashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum []interface{}

func init() {
	var res []HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC
	if err := json.Unmarshal([]byte(`["TIME_WINDOW_UTC_INVALID","WINDOW_12AM_4AM","WINDOW_6AM_10AM","WINDOW_12PM_4PM","WINDOW_6PM_10PM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum = append(hashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum, v)
	}
}

func (m HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) validateHashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum(path, location string, value HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault 20201125 major version upgrade config maintenance window time window u t c
func (m HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTCEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 major version upgrade config maintenance window time window u t c based on context it is used
func (m HashicorpCloudVault20201125MajorVersionUpgradeConfigMaintenanceWindowTimeWindowUTC) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}