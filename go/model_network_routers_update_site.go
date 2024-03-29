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

// NetworkRoutersUpdateSite struct for NetworkRoutersUpdateSite
type NetworkRoutersUpdateSite struct {
	Id *OneOflongstring `json:"id,omitempty"`
}

// NewNetworkRoutersUpdateSite instantiates a new NetworkRoutersUpdateSite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRoutersUpdateSite() *NetworkRoutersUpdateSite {
	this := NetworkRoutersUpdateSite{}
	return &this
}

// NewNetworkRoutersUpdateSiteWithDefaults instantiates a new NetworkRoutersUpdateSite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRoutersUpdateSiteWithDefaults() *NetworkRoutersUpdateSite {
	this := NetworkRoutersUpdateSite{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkRoutersUpdateSite) GetId() OneOflongstring {
	if o == nil || o.Id == nil {
		var ret OneOflongstring
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRoutersUpdateSite) GetIdOk() (*OneOflongstring, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkRoutersUpdateSite) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given OneOflongstring and assigns it to the Id field.
func (o *NetworkRoutersUpdateSite) SetId(v OneOflongstring) {
	o.Id = &v
}

func (o NetworkRoutersUpdateSite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRoutersUpdateSite struct {
	value *NetworkRoutersUpdateSite
	isSet bool
}

func (v NullableNetworkRoutersUpdateSite) Get() *NetworkRoutersUpdateSite {
	return v.value
}

func (v *NullableNetworkRoutersUpdateSite) Set(val *NetworkRoutersUpdateSite) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRoutersUpdateSite) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRoutersUpdateSite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRoutersUpdateSite(val *NetworkRoutersUpdateSite) *NullableNetworkRoutersUpdateSite {
	return &NullableNetworkRoutersUpdateSite{value: val, isSet: true}
}

func (v NullableNetworkRoutersUpdateSite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRoutersUpdateSite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


