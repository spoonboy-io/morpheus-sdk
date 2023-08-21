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

// InlineResponse200136 struct for InlineResponse200136
type InlineResponse200136 struct {
	SecurityPackageType *SecurityPackageType `json:"securityPackageType,omitempty"`
}

// NewInlineResponse200136 instantiates a new InlineResponse200136 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200136() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// NewInlineResponse200136WithDefaults instantiates a new InlineResponse200136 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200136WithDefaults() *InlineResponse200136 {
	this := InlineResponse200136{}
	return &this
}

// GetSecurityPackageType returns the SecurityPackageType field value if set, zero value otherwise.
func (o *InlineResponse200136) GetSecurityPackageType() SecurityPackageType {
	if o == nil || o.SecurityPackageType == nil {
		var ret SecurityPackageType
		return ret
	}
	return *o.SecurityPackageType
}

// GetSecurityPackageTypeOk returns a tuple with the SecurityPackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200136) GetSecurityPackageTypeOk() (*SecurityPackageType, bool) {
	if o == nil || o.SecurityPackageType == nil {
		return nil, false
	}
	return o.SecurityPackageType, true
}

// HasSecurityPackageType returns a boolean if a field has been set.
func (o *InlineResponse200136) HasSecurityPackageType() bool {
	if o != nil && o.SecurityPackageType != nil {
		return true
	}

	return false
}

// SetSecurityPackageType gets a reference to the given SecurityPackageType and assigns it to the SecurityPackageType field.
func (o *InlineResponse200136) SetSecurityPackageType(v SecurityPackageType) {
	o.SecurityPackageType = &v
}

func (o InlineResponse200136) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityPackageType != nil {
		toSerialize["securityPackageType"] = o.SecurityPackageType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200136 struct {
	value *InlineResponse200136
	isSet bool
}

func (v NullableInlineResponse200136) Get() *InlineResponse200136 {
	return v.value
}

func (v *NullableInlineResponse200136) Set(val *InlineResponse200136) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200136) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200136) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200136(val *InlineResponse200136) *NullableInlineResponse200136 {
	return &NullableInlineResponse200136{value: val, isSet: true}
}

func (v NullableInlineResponse200136) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200136) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

