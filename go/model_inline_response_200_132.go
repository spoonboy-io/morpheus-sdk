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

// InlineResponse200132 struct for InlineResponse200132
type InlineResponse200132 struct {
	ResourcePoolGroup *map[string]interface{} `json:"resourcePoolGroup,omitempty"`
}

// NewInlineResponse200132 instantiates a new InlineResponse200132 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200132() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// NewInlineResponse200132WithDefaults instantiates a new InlineResponse200132 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200132WithDefaults() *InlineResponse200132 {
	this := InlineResponse200132{}
	return &this
}

// GetResourcePoolGroup returns the ResourcePoolGroup field value if set, zero value otherwise.
func (o *InlineResponse200132) GetResourcePoolGroup() map[string]interface{} {
	if o == nil || o.ResourcePoolGroup == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResourcePoolGroup
}

// GetResourcePoolGroupOk returns a tuple with the ResourcePoolGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200132) GetResourcePoolGroupOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResourcePoolGroup == nil {
		return nil, false
	}
	return o.ResourcePoolGroup, true
}

// HasResourcePoolGroup returns a boolean if a field has been set.
func (o *InlineResponse200132) HasResourcePoolGroup() bool {
	if o != nil && o.ResourcePoolGroup != nil {
		return true
	}

	return false
}

// SetResourcePoolGroup gets a reference to the given map[string]interface{} and assigns it to the ResourcePoolGroup field.
func (o *InlineResponse200132) SetResourcePoolGroup(v map[string]interface{}) {
	o.ResourcePoolGroup = &v
}

func (o InlineResponse200132) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePoolGroup != nil {
		toSerialize["resourcePoolGroup"] = o.ResourcePoolGroup
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200132 struct {
	value *InlineResponse200132
	isSet bool
}

func (v NullableInlineResponse200132) Get() *InlineResponse200132 {
	return v.value
}

func (v *NullableInlineResponse200132) Set(val *InlineResponse200132) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200132) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200132) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200132(val *InlineResponse200132) *NullableInlineResponse200132 {
	return &NullableInlineResponse200132{value: val, isSet: true}
}

func (v NullableInlineResponse200132) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200132) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


