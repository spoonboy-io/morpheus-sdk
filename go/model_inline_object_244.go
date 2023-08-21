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

// InlineObject244 struct for InlineObject244
type InlineObject244 struct {
	Subnet *ApiSubnetsSubnet `json:"subnet,omitempty"`
	ResourcePermission *ApiSubnetsResourcePermission `json:"resourcePermission,omitempty"`
}

// NewInlineObject244 instantiates a new InlineObject244 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject244() *InlineObject244 {
	this := InlineObject244{}
	return &this
}

// NewInlineObject244WithDefaults instantiates a new InlineObject244 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject244WithDefaults() *InlineObject244 {
	this := InlineObject244{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject244) GetSubnet() ApiSubnetsSubnet {
	if o == nil || o.Subnet == nil {
		var ret ApiSubnetsSubnet
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject244) GetSubnetOk() (*ApiSubnetsSubnet, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject244) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given ApiSubnetsSubnet and assigns it to the Subnet field.
func (o *InlineObject244) SetSubnet(v ApiSubnetsSubnet) {
	o.Subnet = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *InlineObject244) GetResourcePermission() ApiSubnetsResourcePermission {
	if o == nil || o.ResourcePermission == nil {
		var ret ApiSubnetsResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject244) GetResourcePermissionOk() (*ApiSubnetsResourcePermission, bool) {
	if o == nil || o.ResourcePermission == nil {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *InlineObject244) HasResourcePermission() bool {
	if o != nil && o.ResourcePermission != nil {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given ApiSubnetsResourcePermission and assigns it to the ResourcePermission field.
func (o *InlineObject244) SetResourcePermission(v ApiSubnetsResourcePermission) {
	o.ResourcePermission = &v
}

func (o InlineObject244) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.ResourcePermission != nil {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject244 struct {
	value *InlineObject244
	isSet bool
}

func (v NullableInlineObject244) Get() *InlineObject244 {
	return v.value
}

func (v *NullableInlineObject244) Set(val *InlineObject244) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject244) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject244) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject244(val *InlineObject244) *NullableInlineObject244 {
	return &NullableInlineObject244{value: val, isSet: true}
}

func (v NullableInlineObject244) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject244) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


