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

// InlineObject203 struct for InlineObject203
type InlineObject203 struct {
	License *ProvisioningLicensesUpdate `json:"license,omitempty"`
}

// NewInlineObject203 instantiates a new InlineObject203 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject203() *InlineObject203 {
	this := InlineObject203{}
	return &this
}

// NewInlineObject203WithDefaults instantiates a new InlineObject203 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject203WithDefaults() *InlineObject203 {
	this := InlineObject203{}
	return &this
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *InlineObject203) GetLicense() ProvisioningLicensesUpdate {
	if o == nil || o.License == nil {
		var ret ProvisioningLicensesUpdate
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject203) GetLicenseOk() (*ProvisioningLicensesUpdate, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *InlineObject203) HasLicense() bool {
	if o != nil && o.License != nil {
		return true
	}

	return false
}

// SetLicense gets a reference to the given ProvisioningLicensesUpdate and assigns it to the License field.
func (o *InlineObject203) SetLicense(v ProvisioningLicensesUpdate) {
	o.License = &v
}

func (o InlineObject203) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject203 struct {
	value *InlineObject203
	isSet bool
}

func (v NullableInlineObject203) Get() *InlineObject203 {
	return v.value
}

func (v *NullableInlineObject203) Set(val *InlineObject203) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject203) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject203) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject203(val *InlineObject203) *NullableInlineObject203 {
	return &NullableInlineObject203{value: val, isSet: true}
}

func (v NullableInlineObject203) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject203) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


