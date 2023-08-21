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

// AppPrepareApplyDataGroup struct for AppPrepareApplyDataGroup
type AppPrepareApplyDataGroup struct {
	Name *string `json:"name,omitempty"`
	Id *int64 `json:"id,omitempty"`
}

// NewAppPrepareApplyDataGroup instantiates a new AppPrepareApplyDataGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPrepareApplyDataGroup() *AppPrepareApplyDataGroup {
	this := AppPrepareApplyDataGroup{}
	return &this
}

// NewAppPrepareApplyDataGroupWithDefaults instantiates a new AppPrepareApplyDataGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPrepareApplyDataGroupWithDefaults() *AppPrepareApplyDataGroup {
	this := AppPrepareApplyDataGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppPrepareApplyDataGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyDataGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppPrepareApplyDataGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppPrepareApplyDataGroup) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppPrepareApplyDataGroup) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyDataGroup) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppPrepareApplyDataGroup) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AppPrepareApplyDataGroup) SetId(v int64) {
	o.Id = &v
}

func (o AppPrepareApplyDataGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableAppPrepareApplyDataGroup struct {
	value *AppPrepareApplyDataGroup
	isSet bool
}

func (v NullableAppPrepareApplyDataGroup) Get() *AppPrepareApplyDataGroup {
	return v.value
}

func (v *NullableAppPrepareApplyDataGroup) Set(val *AppPrepareApplyDataGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPrepareApplyDataGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPrepareApplyDataGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPrepareApplyDataGroup(val *AppPrepareApplyDataGroup) *NullableAppPrepareApplyDataGroup {
	return &NullableAppPrepareApplyDataGroup{value: val, isSet: true}
}

func (v NullableAppPrepareApplyDataGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPrepareApplyDataGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


