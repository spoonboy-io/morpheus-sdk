/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// License Object for logging in to the [Morpheus Hub](https://morpheushub.com) with existing credentials. This is only required for `hubmode=login`.
type License struct {
	License *LicenseLicense `json:"license,omitempty"`
	CurrentUsage *LicenseCurrentUsage `json:"currentUsage,omitempty"`
}

// NewLicense instantiates a new License object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicense() *License {
	this := License{}
	return &this
}

// NewLicenseWithDefaults instantiates a new License object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseWithDefaults() *License {
	this := License{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *License) GetLicense() LicenseLicense {
	if o == nil || o.License == nil {
		var ret LicenseLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetLicenseOk() (*LicenseLicense, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *License) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given LicenseLicense and assigns it to the License field.
func (o *License) SetLicense(v LicenseLicense) {
	o.License = &v
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *License) GetCurrentUsage() LicenseCurrentUsage {
	if o == nil || o.CurrentUsage == nil {
		var ret LicenseCurrentUsage
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *License) GetCurrentUsageOk() (*LicenseCurrentUsage, bool) {
	if o == nil || o.CurrentUsage == nil {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *License) HasCurrentUsage() bool {
	if o != nil && o.CurrentUsage != nil {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given LicenseCurrentUsage and assigns it to the CurrentUsage field.
func (o *License) SetCurrentUsage(v LicenseCurrentUsage) {
	o.CurrentUsage = &v
}

func (o License) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.CurrentUsage != nil {
		toSerialize["currentUsage"] = o.CurrentUsage
	}
	return json.Marshal(toSerialize)
}

type NullableLicense struct {
	value *License
	isSet bool
}

func (v NullableLicense) Get() *License {
	return v.value
}

func (v *NullableLicense) Set(val *License) {
	v.value = val
	v.isSet = true
}

func (v NullableLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicense(val *License) *NullableLicense {
	return &NullableLicense{value: val, isSet: true}
}

func (v NullableLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


