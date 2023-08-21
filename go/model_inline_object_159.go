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

// InlineObject159 struct for InlineObject159
type InlineObject159 struct {
	NetworkRoute *NetworkRouterRouteCreate `json:"networkRoute,omitempty"`
}

// NewInlineObject159 instantiates a new InlineObject159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject159() *InlineObject159 {
	this := InlineObject159{}
	return &this
}

// NewInlineObject159WithDefaults instantiates a new InlineObject159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject159WithDefaults() *InlineObject159 {
	this := InlineObject159{}
	return &this
}

// GetNetworkRoute returns the NetworkRoute field value if set, zero value otherwise.
func (o *InlineObject159) GetNetworkRoute() NetworkRouterRouteCreate {
	if o == nil || o.NetworkRoute == nil {
		var ret NetworkRouterRouteCreate
		return ret
	}
	return *o.NetworkRoute
}

// GetNetworkRouteOk returns a tuple with the NetworkRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetNetworkRouteOk() (*NetworkRouterRouteCreate, bool) {
	if o == nil || o.NetworkRoute == nil {
		return nil, false
	}
	return o.NetworkRoute, true
}

// HasNetworkRoute returns a boolean if a field has been set.
func (o *InlineObject159) HasNetworkRoute() bool {
	if o != nil && o.NetworkRoute != nil {
		return true
	}

	return false
}

// SetNetworkRoute gets a reference to the given NetworkRouterRouteCreate and assigns it to the NetworkRoute field.
func (o *InlineObject159) SetNetworkRoute(v NetworkRouterRouteCreate) {
	o.NetworkRoute = &v
}

func (o InlineObject159) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkRoute != nil {
		toSerialize["networkRoute"] = o.NetworkRoute
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject159 struct {
	value *InlineObject159
	isSet bool
}

func (v NullableInlineObject159) Get() *InlineObject159 {
	return v.value
}

func (v *NullableInlineObject159) Set(val *InlineObject159) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject159) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject159(val *InlineObject159) *NullableInlineObject159 {
	return &NullableInlineObject159{value: val, isSet: true}
}

func (v NullableInlineObject159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

