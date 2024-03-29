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

// InlineObject206 struct for InlineObject206
type InlineObject206 struct {
	ResourcePoolGroup *ResourcePoolGroupsCreateInput `json:"resourcePoolGroup,omitempty"`
}

// NewInlineObject206 instantiates a new InlineObject206 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject206() *InlineObject206 {
	this := InlineObject206{}
	return &this
}

// NewInlineObject206WithDefaults instantiates a new InlineObject206 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject206WithDefaults() *InlineObject206 {
	this := InlineObject206{}
	return &this
}

// GetResourcePoolGroup returns the ResourcePoolGroup field value if set, zero value otherwise.
func (o *InlineObject206) GetResourcePoolGroup() ResourcePoolGroupsCreateInput {
	if o == nil || o.ResourcePoolGroup == nil {
		var ret ResourcePoolGroupsCreateInput
		return ret
	}
	return *o.ResourcePoolGroup
}

// GetResourcePoolGroupOk returns a tuple with the ResourcePoolGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject206) GetResourcePoolGroupOk() (*ResourcePoolGroupsCreateInput, bool) {
	if o == nil || o.ResourcePoolGroup == nil {
		return nil, false
	}
	return o.ResourcePoolGroup, true
}

// HasResourcePoolGroup returns a boolean if a field has been set.
func (o *InlineObject206) HasResourcePoolGroup() bool {
	if o != nil && o.ResourcePoolGroup != nil {
		return true
	}

	return false
}

// SetResourcePoolGroup gets a reference to the given ResourcePoolGroupsCreateInput and assigns it to the ResourcePoolGroup field.
func (o *InlineObject206) SetResourcePoolGroup(v ResourcePoolGroupsCreateInput) {
	o.ResourcePoolGroup = &v
}

func (o InlineObject206) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePoolGroup != nil {
		toSerialize["resourcePoolGroup"] = o.ResourcePoolGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject206 struct {
	value *InlineObject206
	isSet bool
}

func (v NullableInlineObject206) Get() *InlineObject206 {
	return v.value
}

func (v *NullableInlineObject206) Set(val *InlineObject206) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject206) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject206) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject206(val *InlineObject206) *NullableInlineObject206 {
	return &NullableInlineObject206{value: val, isSet: true}
}

func (v NullableInlineObject206) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject206) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


