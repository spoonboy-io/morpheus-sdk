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

// InlineResponse20094 struct for InlineResponse20094
type InlineResponse20094 struct {
	NetworkRouters *[]InlineResponse20094NetworkRouters `json:"networkRouters,omitempty"`
}

// NewInlineResponse20094 instantiates a new InlineResponse20094 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20094() *InlineResponse20094 {
	this := InlineResponse20094{}
	return &this
}

// NewInlineResponse20094WithDefaults instantiates a new InlineResponse20094 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20094WithDefaults() *InlineResponse20094 {
	this := InlineResponse20094{}
	return &this
}

// GetNetworkRouters returns the NetworkRouters field value if set, zero value otherwise.
func (o *InlineResponse20094) GetNetworkRouters() []InlineResponse20094NetworkRouters {
	if o == nil || o.NetworkRouters == nil {
		var ret []InlineResponse20094NetworkRouters
		return ret
	}
	return *o.NetworkRouters
}

// GetNetworkRoutersOk returns a tuple with the NetworkRouters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094) GetNetworkRoutersOk() (*[]InlineResponse20094NetworkRouters, bool) {
	if o == nil || o.NetworkRouters == nil {
		return nil, false
	}
	return o.NetworkRouters, true
}

// HasNetworkRouters returns a boolean if a field has been set.
func (o *InlineResponse20094) HasNetworkRouters() bool {
	if o != nil && o.NetworkRouters != nil {
		return true
	}

	return false
}

// SetNetworkRouters gets a reference to the given []InlineResponse20094NetworkRouters and assigns it to the NetworkRouters field.
func (o *InlineResponse20094) SetNetworkRouters(v []InlineResponse20094NetworkRouters) {
	o.NetworkRouters = &v
}

func (o InlineResponse20094) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkRouters != nil {
		toSerialize["networkRouters"] = o.NetworkRouters
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20094 struct {
	value *InlineResponse20094
	isSet bool
}

func (v NullableInlineResponse20094) Get() *InlineResponse20094 {
	return v.value
}

func (v *NullableInlineResponse20094) Set(val *InlineResponse20094) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20094) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20094) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20094(val *InlineResponse20094) *NullableInlineResponse20094 {
	return &NullableInlineResponse20094{value: val, isSet: true}
}

func (v NullableInlineResponse20094) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20094) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

