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

// ControllerType struct for ControllerType
type ControllerType struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayOrder *int64 `json:"displayOrder,omitempty"`
	Category *string `json:"category,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Creatable *bool `json:"creatable,omitempty"`
	MaxDevices *int64 `json:"maxDevices,omitempty"`
}

// NewControllerType instantiates a new ControllerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllerType() *ControllerType {
	this := ControllerType{}
	return &this
}

// NewControllerTypeWithDefaults instantiates a new ControllerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllerTypeWithDefaults() *ControllerType {
	this := ControllerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ControllerType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ControllerType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ControllerType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ControllerType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ControllerType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ControllerType) SetName(v string) {
	o.Name = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ControllerType) GetDisplayOrder() int64 {
	if o == nil || o.DisplayOrder == nil {
		var ret int64
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetDisplayOrderOk() (*int64, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *ControllerType) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int64 and assigns it to the DisplayOrder field.
func (o *ControllerType) SetDisplayOrder(v int64) {
	o.DisplayOrder = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ControllerType) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ControllerType) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ControllerType) SetCategory(v string) {
	o.Category = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ControllerType) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ControllerType) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ControllerType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *ControllerType) GetCreatable() bool {
	if o == nil || o.Creatable == nil {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetCreatableOk() (*bool, bool) {
	if o == nil || o.Creatable == nil {
		return nil, false
	}
	return o.Creatable, true
}

// HasCreatable returns a boolean if a field has been set.
func (o *ControllerType) HasCreatable() bool {
	if o != nil && o.Creatable != nil {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *ControllerType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *ControllerType) GetMaxDevices() int64 {
	if o == nil || o.MaxDevices == nil {
		var ret int64
		return ret
	}
	return *o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllerType) GetMaxDevicesOk() (*int64, bool) {
	if o == nil || o.MaxDevices == nil {
		return nil, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *ControllerType) HasMaxDevices() bool {
	if o != nil && o.MaxDevices != nil {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given int64 and assigns it to the MaxDevices field.
func (o *ControllerType) SetMaxDevices(v int64) {
	o.MaxDevices = &v
}

func (o ControllerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Creatable != nil {
		toSerialize["creatable"] = o.Creatable
	}
	if o.MaxDevices != nil {
		toSerialize["maxDevices"] = o.MaxDevices
	}
	return json.Marshal(toSerialize)
}

type NullableControllerType struct {
	value *ControllerType
	isSet bool
}

func (v NullableControllerType) Get() *ControllerType {
	return v.value
}

func (v *NullableControllerType) Set(val *ControllerType) {
	v.value = val
	v.isSet = true
}

func (v NullableControllerType) IsSet() bool {
	return v.isSet
}

func (v *NullableControllerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllerType(val *ControllerType) *NullableControllerType {
	return &NullableControllerType{value: val, isSet: true}
}

func (v NullableControllerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


