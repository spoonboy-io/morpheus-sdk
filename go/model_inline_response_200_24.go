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

// InlineResponse20024 struct for InlineResponse20024
type InlineResponse20024 struct {
	ResourcePool *ZoneResourcePool `json:"resourcePool,omitempty"`
}

// NewInlineResponse20024 instantiates a new InlineResponse20024 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20024() *InlineResponse20024 {
	this := InlineResponse20024{}
	return &this
}

// NewInlineResponse20024WithDefaults instantiates a new InlineResponse20024 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20024WithDefaults() *InlineResponse20024 {
	this := InlineResponse20024{}
	return &this
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *InlineResponse20024) GetResourcePool() ZoneResourcePool {
	if o == nil || o.ResourcePool == nil {
		var ret ZoneResourcePool
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20024) GetResourcePoolOk() (*ZoneResourcePool, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *InlineResponse20024) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given ZoneResourcePool and assigns it to the ResourcePool field.
func (o *InlineResponse20024) SetResourcePool(v ZoneResourcePool) {
	o.ResourcePool = &v
}

func (o InlineResponse20024) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePool != nil {
		toSerialize["resourcePool"] = o.ResourcePool
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20024 struct {
	value *InlineResponse20024
	isSet bool
}

func (v NullableInlineResponse20024) Get() *InlineResponse20024 {
	return v.value
}

func (v *NullableInlineResponse20024) Set(val *InlineResponse20024) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20024) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20024) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20024(val *InlineResponse20024) *NullableInlineResponse20024 {
	return &NullableInlineResponse20024{value: val, isSet: true}
}

func (v NullableInlineResponse20024) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20024) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

