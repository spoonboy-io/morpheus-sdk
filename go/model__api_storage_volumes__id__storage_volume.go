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

// ApiStorageVolumesIdStorageVolume struct for ApiStorageVolumesIdStorageVolume
type ApiStorageVolumesIdStorageVolume struct {
	// A unique name scoped to your account for the storage volume
	Name *string `json:"name,omitempty"`
	// Storage Type Code or ID
	Type *string `json:"type,omitempty"`
	// Configuration object with parameters that vary by `type`.
	Config *map[string]interface{} `json:"config,omitempty"`
	StorageServer *ApiStorageVolumesStorageVolumeStorageServer `json:"storageServer,omitempty"`
	StorageGroup *ApiStorageVolumesStorageVolumeStorageServer `json:"storageGroup,omitempty"`
}

// NewApiStorageVolumesIdStorageVolume instantiates a new ApiStorageVolumesIdStorageVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStorageVolumesIdStorageVolume() *ApiStorageVolumesIdStorageVolume {
	this := ApiStorageVolumesIdStorageVolume{}
	return &this
}

// NewApiStorageVolumesIdStorageVolumeWithDefaults instantiates a new ApiStorageVolumesIdStorageVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStorageVolumesIdStorageVolumeWithDefaults() *ApiStorageVolumesIdStorageVolume {
	this := ApiStorageVolumesIdStorageVolume{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiStorageVolumesIdStorageVolume) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesIdStorageVolume) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiStorageVolumesIdStorageVolume) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiStorageVolumesIdStorageVolume) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiStorageVolumesIdStorageVolume) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesIdStorageVolume) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiStorageVolumesIdStorageVolume) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiStorageVolumesIdStorageVolume) SetType(v string) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ApiStorageVolumesIdStorageVolume) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesIdStorageVolume) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ApiStorageVolumesIdStorageVolume) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ApiStorageVolumesIdStorageVolume) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetStorageServer returns the StorageServer field value if set, zero value otherwise.
func (o *ApiStorageVolumesIdStorageVolume) GetStorageServer() ApiStorageVolumesStorageVolumeStorageServer {
	if o == nil || o.StorageServer == nil {
		var ret ApiStorageVolumesStorageVolumeStorageServer
		return ret
	}
	return *o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesIdStorageVolume) GetStorageServerOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool) {
	if o == nil || o.StorageServer == nil {
		return nil, false
	}
	return o.StorageServer, true
}

// HasStorageServer returns a boolean if a field has been set.
func (o *ApiStorageVolumesIdStorageVolume) HasStorageServer() bool {
	if o != nil && o.StorageServer != nil {
		return true
	}

	return false
}

// SetStorageServer gets a reference to the given ApiStorageVolumesStorageVolumeStorageServer and assigns it to the StorageServer field.
func (o *ApiStorageVolumesIdStorageVolume) SetStorageServer(v ApiStorageVolumesStorageVolumeStorageServer) {
	o.StorageServer = &v
}

// GetStorageGroup returns the StorageGroup field value if set, zero value otherwise.
func (o *ApiStorageVolumesIdStorageVolume) GetStorageGroup() ApiStorageVolumesStorageVolumeStorageServer {
	if o == nil || o.StorageGroup == nil {
		var ret ApiStorageVolumesStorageVolumeStorageServer
		return ret
	}
	return *o.StorageGroup
}

// GetStorageGroupOk returns a tuple with the StorageGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageVolumesIdStorageVolume) GetStorageGroupOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool) {
	if o == nil || o.StorageGroup == nil {
		return nil, false
	}
	return o.StorageGroup, true
}

// HasStorageGroup returns a boolean if a field has been set.
func (o *ApiStorageVolumesIdStorageVolume) HasStorageGroup() bool {
	if o != nil && o.StorageGroup != nil {
		return true
	}

	return false
}

// SetStorageGroup gets a reference to the given ApiStorageVolumesStorageVolumeStorageServer and assigns it to the StorageGroup field.
func (o *ApiStorageVolumesIdStorageVolume) SetStorageGroup(v ApiStorageVolumesStorageVolumeStorageServer) {
	o.StorageGroup = &v
}

func (o ApiStorageVolumesIdStorageVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.StorageServer != nil {
		toSerialize["storageServer"] = o.StorageServer
	}
	if o.StorageGroup != nil {
		toSerialize["storageGroup"] = o.StorageGroup
	}
	return json.Marshal(toSerialize)
}

type NullableApiStorageVolumesIdStorageVolume struct {
	value *ApiStorageVolumesIdStorageVolume
	isSet bool
}

func (v NullableApiStorageVolumesIdStorageVolume) Get() *ApiStorageVolumesIdStorageVolume {
	return v.value
}

func (v *NullableApiStorageVolumesIdStorageVolume) Set(val *ApiStorageVolumesIdStorageVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStorageVolumesIdStorageVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStorageVolumesIdStorageVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStorageVolumesIdStorageVolume(val *ApiStorageVolumesIdStorageVolume) *NullableApiStorageVolumesIdStorageVolume {
	return &NullableApiStorageVolumesIdStorageVolume{value: val, isSet: true}
}

func (v NullableApiStorageVolumesIdStorageVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStorageVolumesIdStorageVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


