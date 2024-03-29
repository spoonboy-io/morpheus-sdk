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

// ImageBuildsConfigConfig struct for ImageBuildsConfigConfig
type ImageBuildsConfigConfig struct {
	Template *int64 `json:"template,omitempty"`
	ResourcePoolId *int64 `json:"resourcePoolId,omitempty"`
	Expose *int64 `json:"expose,omitempty"`
}

// NewImageBuildsConfigConfig instantiates a new ImageBuildsConfigConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBuildsConfigConfig() *ImageBuildsConfigConfig {
	this := ImageBuildsConfigConfig{}
	return &this
}

// NewImageBuildsConfigConfigWithDefaults instantiates a new ImageBuildsConfigConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBuildsConfigConfigWithDefaults() *ImageBuildsConfigConfig {
	this := ImageBuildsConfigConfig{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ImageBuildsConfigConfig) GetTemplate() int64 {
	if o == nil || o.Template == nil {
		var ret int64
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildsConfigConfig) GetTemplateOk() (*int64, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ImageBuildsConfigConfig) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given int64 and assigns it to the Template field.
func (o *ImageBuildsConfigConfig) SetTemplate(v int64) {
	o.Template = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *ImageBuildsConfigConfig) GetResourcePoolId() int64 {
	if o == nil || o.ResourcePoolId == nil {
		var ret int64
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildsConfigConfig) GetResourcePoolIdOk() (*int64, bool) {
	if o == nil || o.ResourcePoolId == nil {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *ImageBuildsConfigConfig) HasResourcePoolId() bool {
	if o != nil && o.ResourcePoolId != nil {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given int64 and assigns it to the ResourcePoolId field.
func (o *ImageBuildsConfigConfig) SetResourcePoolId(v int64) {
	o.ResourcePoolId = &v
}

// GetExpose returns the Expose field value if set, zero value otherwise.
func (o *ImageBuildsConfigConfig) GetExpose() int64 {
	if o == nil || o.Expose == nil {
		var ret int64
		return ret
	}
	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildsConfigConfig) GetExposeOk() (*int64, bool) {
	if o == nil || o.Expose == nil {
		return nil, false
	}
	return o.Expose, true
}

// HasExpose returns a boolean if a field has been set.
func (o *ImageBuildsConfigConfig) HasExpose() bool {
	if o != nil && o.Expose != nil {
		return true
	}

	return false
}

// SetExpose gets a reference to the given int64 and assigns it to the Expose field.
func (o *ImageBuildsConfigConfig) SetExpose(v int64) {
	o.Expose = &v
}

func (o ImageBuildsConfigConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.ResourcePoolId != nil {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if o.Expose != nil {
		toSerialize["expose"] = o.Expose
	}
	return json.Marshal(toSerialize)
}

type NullableImageBuildsConfigConfig struct {
	value *ImageBuildsConfigConfig
	isSet bool
}

func (v NullableImageBuildsConfigConfig) Get() *ImageBuildsConfigConfig {
	return v.value
}

func (v *NullableImageBuildsConfigConfig) Set(val *ImageBuildsConfigConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildsConfigConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildsConfigConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildsConfigConfig(val *ImageBuildsConfigConfig) *NullableImageBuildsConfigConfig {
	return &NullableImageBuildsConfigConfig{value: val, isSet: true}
}

func (v NullableImageBuildsConfigConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildsConfigConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


