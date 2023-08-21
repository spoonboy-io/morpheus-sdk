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

// InlineResponse200127 struct for InlineResponse200127
type InlineResponse200127 struct {
	Licenses *[]ProvisioningLicenseReservations `json:"licenses,omitempty"`
}

// NewInlineResponse200127 instantiates a new InlineResponse200127 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200127() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// NewInlineResponse200127WithDefaults instantiates a new InlineResponse200127 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200127WithDefaults() *InlineResponse200127 {
	this := InlineResponse200127{}
	return &this
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *InlineResponse200127) GetLicenses() []ProvisioningLicenseReservations {
	if o == nil || o.Licenses == nil {
		var ret []ProvisioningLicenseReservations
		return ret
	}
	return *o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200127) GetLicensesOk() (*[]ProvisioningLicenseReservations, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *InlineResponse200127) HasLicenses() bool {
	if o != nil && o.Licenses != nil {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []ProvisioningLicenseReservations and assigns it to the Licenses field.
func (o *InlineResponse200127) SetLicenses(v []ProvisioningLicenseReservations) {
	o.Licenses = &v
}

func (o InlineResponse200127) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200127 struct {
	value *InlineResponse200127
	isSet bool
}

func (v NullableInlineResponse200127) Get() *InlineResponse200127 {
	return v.value
}

func (v *NullableInlineResponse200127) Set(val *InlineResponse200127) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200127) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200127) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200127(val *InlineResponse200127) *NullableInlineResponse200127 {
	return &NullableInlineResponse200127{value: val, isSet: true}
}

func (v NullableInlineResponse200127) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200127) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

