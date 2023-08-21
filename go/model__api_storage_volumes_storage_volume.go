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

// ApiStorageVolumesStorageVolume struct for ApiStorageVolumesStorageVolume
type ApiStorageVolumesStorageVolume struct {
	// A unique name scoped to your account for the storage volume
	Name string `json:"name"`
	// Storage Type Code or ID
	Type string `json:"type"`
	// Configuration object with parameters that vary by `type`.
	Config *map[string]interface{} `json:"config,omitempty"`
	StorageServer ApiStorageVolumesStorageVolumeStorageServer `json:"storageServer"`
	StorageGroup ApiStorageVolumesStorageVolumeStorageServer `json:"storageGroup"`
}

// NewApiStorageVolumesStorageVolume instantiates a new ApiStorageVolumesStorageVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStorageVolumesStorageVolume(name string, type_ string, storageServer ApiStorageVolumesStorageVolumeStorageServer, storageGroup ApiStorageVolumesStorageVolumeStorageServer, ) *ApiStorageVolumesStorageVolume {
	this := ApiStorageVolumesStorageVolume{}
	this.Name = name
	this.Type = type_
	this.StorageServer = storageServer
	this.StorageGroup = storageGroup
	return &this
}

// NewApiStorageVolumesStorageVolumeWithDefaults instantiates a new ApiStorageVolumesStorageVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStorageVolumesStorageVolumeWithDefaults() *ApiStorageVolumesStorageVolume {
	this := ApiStorageVolumesStorageVolume{}
	return &this
}

// GetName returns the Name field value
func (o *ApiStorageVolumesStorageVolume) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesStorageVolume) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiStorageVolumesStorageVolume) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ApiStorageVolumesStorageVolume) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesStorageVolume) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiStorageVolumesStorageVolume) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ApiStorageVolumesStorageVolume) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesStorageVolume) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ApiStorageVolumesStorageVolume) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ApiStorageVolumesStorageVolume) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetStorageServer returns the StorageServer field value
func (o *ApiStorageVolumesStorageVolume) GetStorageServer() ApiStorageVolumesStorageVolumeStorageServer {
	if o == nil  {
		var ret ApiStorageVolumesStorageVolumeStorageServer
		return ret
	}

	return o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesStorageVolume) GetStorageServerOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageServer, true
}

// SetStorageServer sets field value
func (o *ApiStorageVolumesStorageVolume) SetStorageServer(v ApiStorageVolumesStorageVolumeStorageServer) {
	o.StorageServer = v
}

// GetStorageGroup returns the StorageGroup field value
func (o *ApiStorageVolumesStorageVolume) GetStorageGroup() ApiStorageVolumesStorageVolumeStorageServer {
	if o == nil  {
		var ret ApiStorageVolumesStorageVolumeStorageServer
		return ret
	}

	return o.StorageGroup
}

// GetStorageGroupOk returns a tuple with the StorageGroup field value
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesStorageVolume) GetStorageGroupOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageGroup, true
}

// SetStorageGroup sets field value
func (o *ApiStorageVolumesStorageVolume) SetStorageGroup(v ApiStorageVolumesStorageVolumeStorageServer) {
	o.StorageGroup = v
}

func (o ApiStorageVolumesStorageVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["storageServer"] = o.StorageServer
	}
	if true {
		toSerialize["storageGroup"] = o.StorageGroup
	}
	return json.Marshal(toSerialize)
}

type NullableApiStorageVolumesStorageVolume struct {
	value *ApiStorageVolumesStorageVolume
	isSet bool
}

func (v NullableApiStorageVolumesStorageVolume) Get() *ApiStorageVolumesStorageVolume {
	return v.value
}

func (v *NullableApiStorageVolumesStorageVolume) Set(val *ApiStorageVolumesStorageVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStorageVolumesStorageVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStorageVolumesStorageVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStorageVolumesStorageVolume(val *ApiStorageVolumesStorageVolume) *NullableApiStorageVolumesStorageVolume {
	return &NullableApiStorageVolumesStorageVolume{value: val, isSet: true}
}

func (v NullableApiStorageVolumesStorageVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStorageVolumesStorageVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


