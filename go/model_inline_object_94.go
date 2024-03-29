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

// InlineObject94 struct for InlineObject94
type InlineObject94 struct {
	// List of all security groups ids which should be applied. If no security groups should apply, pass '[]'
	SecurityGroupIds *[]int64 `json:"securityGroupIds,omitempty"`
}

// NewInlineObject94 instantiates a new InlineObject94 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject94() *InlineObject94 {
	this := InlineObject94{}
	return &this
}

// NewInlineObject94WithDefaults instantiates a new InlineObject94 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject94WithDefaults() *InlineObject94 {
	this := InlineObject94{}
	return &this
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *InlineObject94) GetSecurityGroupIds() []int64 {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []int64
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject94) GetSecurityGroupIdsOk() (*[]int64, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		return nil, false
	}
	return o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *InlineObject94) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []int64 and assigns it to the SecurityGroupIds field.
func (o *InlineObject94) SetSecurityGroupIds(v []int64) {
	o.SecurityGroupIds = &v
}

func (o InlineObject94) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecurityGroupIds != nil {
		toSerialize["securityGroupIds"] = o.SecurityGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject94 struct {
	value *InlineObject94
	isSet bool
}

func (v NullableInlineObject94) Get() *InlineObject94 {
	return v.value
}

func (v *NullableInlineObject94) Set(val *InlineObject94) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject94) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject94) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject94(val *InlineObject94) *NullableInlineObject94 {
	return &NullableInlineObject94{value: val, isSet: true}
}

func (v NullableInlineObject94) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject94) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


