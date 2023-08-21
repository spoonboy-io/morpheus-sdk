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

// InlineObject218 struct for InlineObject218
type InlineObject218 struct {
	SecurityPackage ApiSecurityPackagesSecurityPackage `json:"securityPackage"`
}

// NewInlineObject218 instantiates a new InlineObject218 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject218(securityPackage ApiSecurityPackagesSecurityPackage, ) *InlineObject218 {
	this := InlineObject218{}
	this.SecurityPackage = securityPackage
	return &this
}

// NewInlineObject218WithDefaults instantiates a new InlineObject218 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject218WithDefaults() *InlineObject218 {
	this := InlineObject218{}
	return &this
}

// GetSecurityPackage returns the SecurityPackage field value
func (o *InlineObject218) GetSecurityPackage() ApiSecurityPackagesSecurityPackage {
	if o == nil  {
		var ret ApiSecurityPackagesSecurityPackage
		return ret
	}

	return o.SecurityPackage
}

// GetSecurityPackageOk returns a tuple with the SecurityPackage field value
// and a boolean to check if the value has been set.
func (o *InlineObject218) GetSecurityPackageOk() (*ApiSecurityPackagesSecurityPackage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityPackage, true
}

// SetSecurityPackage sets field value
func (o *InlineObject218) SetSecurityPackage(v ApiSecurityPackagesSecurityPackage) {
	o.SecurityPackage = v
}

func (o InlineObject218) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["securityPackage"] = o.SecurityPackage
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject218 struct {
	value *InlineObject218
	isSet bool
}

func (v NullableInlineObject218) Get() *InlineObject218 {
	return v.value
}

func (v *NullableInlineObject218) Set(val *InlineObject218) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject218) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject218) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject218(val *InlineObject218) *NullableInlineObject218 {
	return &NullableInlineObject218{value: val, isSet: true}
}

func (v NullableInlineObject218) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject218) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

