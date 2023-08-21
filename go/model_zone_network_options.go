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

// ZoneNetworkOptions struct for ZoneNetworkOptions
type ZoneNetworkOptions struct {
	Networks *[]ZoneNetworkOptionsNetworks `json:"networks,omitempty"`
	NetworkGroups *[]map[string]interface{} `json:"networkGroups,omitempty"`
	NetworkTypes *[]ZoneNetworkOptionsNetworkTypes `json:"networkTypes,omitempty"`
	NetworkSubnets *[]map[string]interface{} `json:"networkSubnets,omitempty"`
	HasNetworks *bool `json:"hasNetworks,omitempty"`
	MaxNetworks *int64 `json:"maxNetworks,omitempty"`
	EnableNetworkTypeSelection *string `json:"enableNetworkTypeSelection,omitempty"`
	SupportsNetworkSelection *bool `json:"supportsNetworkSelection,omitempty"`
}

// NewZoneNetworkOptions instantiates a new ZoneNetworkOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneNetworkOptions() *ZoneNetworkOptions {
	this := ZoneNetworkOptions{}
	return &this
}

// NewZoneNetworkOptionsWithDefaults instantiates a new ZoneNetworkOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneNetworkOptionsWithDefaults() *ZoneNetworkOptions {
	this := ZoneNetworkOptions{}
	return &this
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetNetworks() []ZoneNetworkOptionsNetworks {
	if o == nil || o.Networks == nil {
		var ret []ZoneNetworkOptionsNetworks
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetNetworksOk() (*[]ZoneNetworkOptionsNetworks, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []ZoneNetworkOptionsNetworks and assigns it to the Networks field.
func (o *ZoneNetworkOptions) SetNetworks(v []ZoneNetworkOptionsNetworks) {
	o.Networks = &v
}

// GetNetworkGroups returns the NetworkGroups field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetNetworkGroups() []map[string]interface{} {
	if o == nil || o.NetworkGroups == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.NetworkGroups
}

// GetNetworkGroupsOk returns a tuple with the NetworkGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetNetworkGroupsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.NetworkGroups == nil {
		return nil, false
	}
	return o.NetworkGroups, true
}

// HasNetworkGroups returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasNetworkGroups() bool {
	if o != nil && o.NetworkGroups != nil {
		return true
	}

	return false
}

// SetNetworkGroups gets a reference to the given []map[string]interface{} and assigns it to the NetworkGroups field.
func (o *ZoneNetworkOptions) SetNetworkGroups(v []map[string]interface{}) {
	o.NetworkGroups = &v
}

// GetNetworkTypes returns the NetworkTypes field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetNetworkTypes() []ZoneNetworkOptionsNetworkTypes {
	if o == nil || o.NetworkTypes == nil {
		var ret []ZoneNetworkOptionsNetworkTypes
		return ret
	}
	return *o.NetworkTypes
}

// GetNetworkTypesOk returns a tuple with the NetworkTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetNetworkTypesOk() (*[]ZoneNetworkOptionsNetworkTypes, bool) {
	if o == nil || o.NetworkTypes == nil {
		return nil, false
	}
	return o.NetworkTypes, true
}

// HasNetworkTypes returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasNetworkTypes() bool {
	if o != nil && o.NetworkTypes != nil {
		return true
	}

	return false
}

// SetNetworkTypes gets a reference to the given []ZoneNetworkOptionsNetworkTypes and assigns it to the NetworkTypes field.
func (o *ZoneNetworkOptions) SetNetworkTypes(v []ZoneNetworkOptionsNetworkTypes) {
	o.NetworkTypes = &v
}

// GetNetworkSubnets returns the NetworkSubnets field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetNetworkSubnets() []map[string]interface{} {
	if o == nil || o.NetworkSubnets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.NetworkSubnets
}

// GetNetworkSubnetsOk returns a tuple with the NetworkSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetNetworkSubnetsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.NetworkSubnets == nil {
		return nil, false
	}
	return o.NetworkSubnets, true
}

// HasNetworkSubnets returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasNetworkSubnets() bool {
	if o != nil && o.NetworkSubnets != nil {
		return true
	}

	return false
}

// SetNetworkSubnets gets a reference to the given []map[string]interface{} and assigns it to the NetworkSubnets field.
func (o *ZoneNetworkOptions) SetNetworkSubnets(v []map[string]interface{}) {
	o.NetworkSubnets = &v
}

// GetHasNetworks returns the HasNetworks field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetHasNetworks() bool {
	if o == nil || o.HasNetworks == nil {
		var ret bool
		return ret
	}
	return *o.HasNetworks
}

// GetHasNetworksOk returns a tuple with the HasNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetHasNetworksOk() (*bool, bool) {
	if o == nil || o.HasNetworks == nil {
		return nil, false
	}
	return o.HasNetworks, true
}

// HasHasNetworks returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasHasNetworks() bool {
	if o != nil && o.HasNetworks != nil {
		return true
	}

	return false
}

// SetHasNetworks gets a reference to the given bool and assigns it to the HasNetworks field.
func (o *ZoneNetworkOptions) SetHasNetworks(v bool) {
	o.HasNetworks = &v
}

// GetMaxNetworks returns the MaxNetworks field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetMaxNetworks() int64 {
	if o == nil || o.MaxNetworks == nil {
		var ret int64
		return ret
	}
	return *o.MaxNetworks
}

// GetMaxNetworksOk returns a tuple with the MaxNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetMaxNetworksOk() (*int64, bool) {
	if o == nil || o.MaxNetworks == nil {
		return nil, false
	}
	return o.MaxNetworks, true
}

// HasMaxNetworks returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasMaxNetworks() bool {
	if o != nil && o.MaxNetworks != nil {
		return true
	}

	return false
}

// SetMaxNetworks gets a reference to the given int64 and assigns it to the MaxNetworks field.
func (o *ZoneNetworkOptions) SetMaxNetworks(v int64) {
	o.MaxNetworks = &v
}

// GetEnableNetworkTypeSelection returns the EnableNetworkTypeSelection field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetEnableNetworkTypeSelection() string {
	if o == nil || o.EnableNetworkTypeSelection == nil {
		var ret string
		return ret
	}
	return *o.EnableNetworkTypeSelection
}

// GetEnableNetworkTypeSelectionOk returns a tuple with the EnableNetworkTypeSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetEnableNetworkTypeSelectionOk() (*string, bool) {
	if o == nil || o.EnableNetworkTypeSelection == nil {
		return nil, false
	}
	return o.EnableNetworkTypeSelection, true
}

// HasEnableNetworkTypeSelection returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasEnableNetworkTypeSelection() bool {
	if o != nil && o.EnableNetworkTypeSelection != nil {
		return true
	}

	return false
}

// SetEnableNetworkTypeSelection gets a reference to the given string and assigns it to the EnableNetworkTypeSelection field.
func (o *ZoneNetworkOptions) SetEnableNetworkTypeSelection(v string) {
	o.EnableNetworkTypeSelection = &v
}

// GetSupportsNetworkSelection returns the SupportsNetworkSelection field value if set, zero value otherwise.
func (o *ZoneNetworkOptions) GetSupportsNetworkSelection() bool {
	if o == nil || o.SupportsNetworkSelection == nil {
		var ret bool
		return ret
	}
	return *o.SupportsNetworkSelection
}

// GetSupportsNetworkSelectionOk returns a tuple with the SupportsNetworkSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptions) GetSupportsNetworkSelectionOk() (*bool, bool) {
	if o == nil || o.SupportsNetworkSelection == nil {
		return nil, false
	}
	return o.SupportsNetworkSelection, true
}

// HasSupportsNetworkSelection returns a boolean if a field has been set.
func (o *ZoneNetworkOptions) HasSupportsNetworkSelection() bool {
	if o != nil && o.SupportsNetworkSelection != nil {
		return true
	}

	return false
}

// SetSupportsNetworkSelection gets a reference to the given bool and assigns it to the SupportsNetworkSelection field.
func (o *ZoneNetworkOptions) SetSupportsNetworkSelection(v bool) {
	o.SupportsNetworkSelection = &v
}

func (o ZoneNetworkOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
	}
	if o.NetworkGroups != nil {
		toSerialize["networkGroups"] = o.NetworkGroups
	}
	if o.NetworkTypes != nil {
		toSerialize["networkTypes"] = o.NetworkTypes
	}
	if o.NetworkSubnets != nil {
		toSerialize["networkSubnets"] = o.NetworkSubnets
	}
	if o.HasNetworks != nil {
		toSerialize["hasNetworks"] = o.HasNetworks
	}
	if o.MaxNetworks != nil {
		toSerialize["maxNetworks"] = o.MaxNetworks
	}
	if o.EnableNetworkTypeSelection != nil {
		toSerialize["enableNetworkTypeSelection"] = o.EnableNetworkTypeSelection
	}
	if o.SupportsNetworkSelection != nil {
		toSerialize["supportsNetworkSelection"] = o.SupportsNetworkSelection
	}
	return json.Marshal(toSerialize)
}

type NullableZoneNetworkOptions struct {
	value *ZoneNetworkOptions
	isSet bool
}

func (v NullableZoneNetworkOptions) Get() *ZoneNetworkOptions {
	return v.value
}

func (v *NullableZoneNetworkOptions) Set(val *ZoneNetworkOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneNetworkOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneNetworkOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneNetworkOptions(val *ZoneNetworkOptions) *NullableZoneNetworkOptions {
	return &NullableZoneNetworkOptions{value: val, isSet: true}
}

func (v NullableZoneNetworkOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneNetworkOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

