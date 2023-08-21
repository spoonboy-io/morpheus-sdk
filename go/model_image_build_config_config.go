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

// ImageBuildConfigConfig struct for ImageBuildConfigConfig
type ImageBuildConfigConfig struct {
	Template *int64 `json:"template,omitempty"`
	VmwareFolderId *string `json:"vmwareFolderId,omitempty"`
	ResourcePoolId *int64 `json:"resourcePoolId,omitempty"`
	NestedVirtualization *string `json:"nestedVirtualization,omitempty"`
}

// NewImageBuildConfigConfig instantiates a new ImageBuildConfigConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBuildConfigConfig() *ImageBuildConfigConfig {
	this := ImageBuildConfigConfig{}
	return &this
}

// NewImageBuildConfigConfigWithDefaults instantiates a new ImageBuildConfigConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBuildConfigConfigWithDefaults() *ImageBuildConfigConfig {
	this := ImageBuildConfigConfig{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ImageBuildConfigConfig) GetTemplate() int64 {
	if o == nil || o.Template == nil {
		var ret int64
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildConfigConfig) GetTemplateOk() (*int64, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ImageBuildConfigConfig) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given int64 and assigns it to the Template field.
func (o *ImageBuildConfigConfig) SetTemplate(v int64) {
	o.Template = &v
}

// GetVmwareFolderId returns the VmwareFolderId field value if set, zero value otherwise.
func (o *ImageBuildConfigConfig) GetVmwareFolderId() string {
	if o == nil || o.VmwareFolderId == nil {
		var ret string
		return ret
	}
	return *o.VmwareFolderId
}

// GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildConfigConfig) GetVmwareFolderIdOk() (*string, bool) {
	if o == nil || o.VmwareFolderId == nil {
		return nil, false
	}
	return o.VmwareFolderId, true
}

// HasVmwareFolderId returns a boolean if a field has been set.
func (o *ImageBuildConfigConfig) HasVmwareFolderId() bool {
	if o != nil && o.VmwareFolderId != nil {
		return true
	}

	return false
}

// SetVmwareFolderId gets a reference to the given string and assigns it to the VmwareFolderId field.
func (o *ImageBuildConfigConfig) SetVmwareFolderId(v string) {
	o.VmwareFolderId = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *ImageBuildConfigConfig) GetResourcePoolId() int64 {
	if o == nil || o.ResourcePoolId == nil {
		var ret int64
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildConfigConfig) GetResourcePoolIdOk() (*int64, bool) {
	if o == nil || o.ResourcePoolId == nil {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *ImageBuildConfigConfig) HasResourcePoolId() bool {
	if o != nil && o.ResourcePoolId != nil {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given int64 and assigns it to the ResourcePoolId field.
func (o *ImageBuildConfigConfig) SetResourcePoolId(v int64) {
	o.ResourcePoolId = &v
}

// GetNestedVirtualization returns the NestedVirtualization field value if set, zero value otherwise.
func (o *ImageBuildConfigConfig) GetNestedVirtualization() string {
	if o == nil || o.NestedVirtualization == nil {
		var ret string
		return ret
	}
	return *o.NestedVirtualization
}

// GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildConfigConfig) GetNestedVirtualizationOk() (*string, bool) {
	if o == nil || o.NestedVirtualization == nil {
		return nil, false
	}
	return o.NestedVirtualization, true
}

// HasNestedVirtualization returns a boolean if a field has been set.
func (o *ImageBuildConfigConfig) HasNestedVirtualization() bool {
	if o != nil && o.NestedVirtualization != nil {
		return true
	}

	return false
}

// SetNestedVirtualization gets a reference to the given string and assigns it to the NestedVirtualization field.
func (o *ImageBuildConfigConfig) SetNestedVirtualization(v string) {
	o.NestedVirtualization = &v
}

func (o ImageBuildConfigConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.VmwareFolderId != nil {
		toSerialize["vmwareFolderId"] = o.VmwareFolderId
	}
	if o.ResourcePoolId != nil {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if o.NestedVirtualization != nil {
		toSerialize["nestedVirtualization"] = o.NestedVirtualization
	}
	return json.Marshal(toSerialize)
}

type NullableImageBuildConfigConfig struct {
	value *ImageBuildConfigConfig
	isSet bool
}

func (v NullableImageBuildConfigConfig) Get() *ImageBuildConfigConfig {
	return v.value
}

func (v *NullableImageBuildConfigConfig) Set(val *ImageBuildConfigConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildConfigConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildConfigConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildConfigConfig(val *ImageBuildConfigConfig) *NullableImageBuildConfigConfig {
	return &NullableImageBuildConfigConfig{value: val, isSet: true}
}

func (v NullableImageBuildConfigConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildConfigConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


