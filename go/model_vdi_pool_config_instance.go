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

// VdiPoolConfigInstance struct for VdiPoolConfigInstance
type VdiPoolConfigInstance struct {
	UserGroup *ZoneVcenterConfigNetworkServer `json:"userGroup,omitempty"`
	NetworkDomain *ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"networkDomain,omitempty"`
}

// NewVdiPoolConfigInstance instantiates a new VdiPoolConfigInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVdiPoolConfigInstance() *VdiPoolConfigInstance {
	this := VdiPoolConfigInstance{}
	return &this
}

// NewVdiPoolConfigInstanceWithDefaults instantiates a new VdiPoolConfigInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVdiPoolConfigInstanceWithDefaults() *VdiPoolConfigInstance {
	this := VdiPoolConfigInstance{}
	return &this
}

// GetUserGroup returns the UserGroup field value if set, zero value otherwise.
func (o *VdiPoolConfigInstance) GetUserGroup() ZoneVcenterConfigNetworkServer {
	if o == nil || o.UserGroup == nil {
		var ret ZoneVcenterConfigNetworkServer
		return ret
	}
	return *o.UserGroup
}

// GetUserGroupOk returns a tuple with the UserGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdiPoolConfigInstance) GetUserGroupOk() (*ZoneVcenterConfigNetworkServer, bool) {
	if o == nil || o.UserGroup == nil {
		return nil, false
	}
	return o.UserGroup, true
}

// HasUserGroup returns a boolean if a field has been set.
func (o *VdiPoolConfigInstance) HasUserGroup() bool {
	if o != nil && o.UserGroup != nil {
		return true
	}

	return false
}

// SetUserGroup gets a reference to the given ZoneVcenterConfigNetworkServer and assigns it to the UserGroup field.
func (o *VdiPoolConfigInstance) SetUserGroup(v ZoneVcenterConfigNetworkServer) {
	o.UserGroup = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise.
func (o *VdiPoolConfigInstance) GetNetworkDomain() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.NetworkDomain == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.NetworkDomain
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VdiPoolConfigInstance) GetNetworkDomainOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.NetworkDomain == nil {
		return nil, false
	}
	return o.NetworkDomain, true
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *VdiPoolConfigInstance) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain != nil {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the NetworkDomain field.
func (o *VdiPoolConfigInstance) SetNetworkDomain(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.NetworkDomain = &v
}

func (o VdiPoolConfigInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserGroup != nil {
		toSerialize["userGroup"] = o.UserGroup
	}
	if o.NetworkDomain != nil {
		toSerialize["networkDomain"] = o.NetworkDomain
	}
	return json.Marshal(toSerialize)
}

type NullableVdiPoolConfigInstance struct {
	value *VdiPoolConfigInstance
	isSet bool
}

func (v NullableVdiPoolConfigInstance) Get() *VdiPoolConfigInstance {
	return v.value
}

func (v *NullableVdiPoolConfigInstance) Set(val *VdiPoolConfigInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableVdiPoolConfigInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableVdiPoolConfigInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVdiPoolConfigInstance(val *VdiPoolConfigInstance) *NullableVdiPoolConfigInstance {
	return &NullableVdiPoolConfigInstance{value: val, isSet: true}
}

func (v NullableVdiPoolConfigInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVdiPoolConfigInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


