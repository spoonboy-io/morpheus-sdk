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

// ZoneNetworkOptionsNetworks struct for ZoneNetworkOptionsNetworks
type ZoneNetworkOptionsNetworks struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DhcpServer *bool `json:"dhcpServer,omitempty"`
	AllowStaticOverride *bool `json:"allowStaticOverride,omitempty"`
	Pool NullableString `json:"pool,omitempty"`
}

// NewZoneNetworkOptionsNetworks instantiates a new ZoneNetworkOptionsNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneNetworkOptionsNetworks() *ZoneNetworkOptionsNetworks {
	this := ZoneNetworkOptionsNetworks{}
	return &this
}

// NewZoneNetworkOptionsNetworksWithDefaults instantiates a new ZoneNetworkOptionsNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneNetworkOptionsNetworksWithDefaults() *ZoneNetworkOptionsNetworks {
	this := ZoneNetworkOptionsNetworks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneNetworkOptionsNetworks) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptionsNetworks) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneNetworkOptionsNetworks) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ZoneNetworkOptionsNetworks) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneNetworkOptionsNetworks) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptionsNetworks) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneNetworkOptionsNetworks) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneNetworkOptionsNetworks) SetName(v string) {
	o.Name = &v
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *ZoneNetworkOptionsNetworks) GetDhcpServer() bool {
	if o == nil || o.DhcpServer == nil {
		var ret bool
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptionsNetworks) GetDhcpServerOk() (*bool, bool) {
	if o == nil || o.DhcpServer == nil {
		return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *ZoneNetworkOptionsNetworks) HasDhcpServer() bool {
	if o != nil && o.DhcpServer != nil {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given bool and assigns it to the DhcpServer field.
func (o *ZoneNetworkOptionsNetworks) SetDhcpServer(v bool) {
	o.DhcpServer = &v
}

// GetAllowStaticOverride returns the AllowStaticOverride field value if set, zero value otherwise.
func (o *ZoneNetworkOptionsNetworks) GetAllowStaticOverride() bool {
	if o == nil || o.AllowStaticOverride == nil {
		var ret bool
		return ret
	}
	return *o.AllowStaticOverride
}

// GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneNetworkOptionsNetworks) GetAllowStaticOverrideOk() (*bool, bool) {
	if o == nil || o.AllowStaticOverride == nil {
		return nil, false
	}
	return o.AllowStaticOverride, true
}

// HasAllowStaticOverride returns a boolean if a field has been set.
func (o *ZoneNetworkOptionsNetworks) HasAllowStaticOverride() bool {
	if o != nil && o.AllowStaticOverride != nil {
		return true
	}

	return false
}

// SetAllowStaticOverride gets a reference to the given bool and assigns it to the AllowStaticOverride field.
func (o *ZoneNetworkOptionsNetworks) SetAllowStaticOverride(v bool) {
	o.AllowStaticOverride = &v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZoneNetworkOptionsNetworks) GetPool() string {
	if o == nil || o.Pool.Get() == nil {
		var ret string
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneNetworkOptionsNetworks) GetPoolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *ZoneNetworkOptionsNetworks) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableString and assigns it to the Pool field.
func (o *ZoneNetworkOptionsNetworks) SetPool(v string) {
	o.Pool.Set(&v)
}
// SetPoolNil sets the value for Pool to be an explicit nil
func (o *ZoneNetworkOptionsNetworks) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *ZoneNetworkOptionsNetworks) UnsetPool() {
	o.Pool.Unset()
}

func (o ZoneNetworkOptionsNetworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DhcpServer != nil {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if o.AllowStaticOverride != nil {
		toSerialize["allowStaticOverride"] = o.AllowStaticOverride
	}
	if o.Pool.IsSet() {
		toSerialize["pool"] = o.Pool.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableZoneNetworkOptionsNetworks struct {
	value *ZoneNetworkOptionsNetworks
	isSet bool
}

func (v NullableZoneNetworkOptionsNetworks) Get() *ZoneNetworkOptionsNetworks {
	return v.value
}

func (v *NullableZoneNetworkOptionsNetworks) Set(val *ZoneNetworkOptionsNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneNetworkOptionsNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneNetworkOptionsNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneNetworkOptionsNetworks(val *ZoneNetworkOptionsNetworks) *NullableZoneNetworkOptionsNetworks {
	return &NullableZoneNetworkOptionsNetworks{value: val, isSet: true}
}

func (v NullableZoneNetworkOptionsNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneNetworkOptionsNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


