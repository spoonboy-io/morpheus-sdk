/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AppBlueprint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppBlueprint{}

// AppBlueprint struct for AppBlueprint
type AppBlueprint struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewAppBlueprint instantiates a new AppBlueprint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppBlueprint() *AppBlueprint {
	this := AppBlueprint{}
	return &this
}

// NewAppBlueprintWithDefaults instantiates a new AppBlueprint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppBlueprintWithDefaults() *AppBlueprint {
	this := AppBlueprint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppBlueprint) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBlueprint) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppBlueprint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AppBlueprint) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppBlueprint) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBlueprint) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppBlueprint) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppBlueprint) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppBlueprint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppBlueprint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppBlueprint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppBlueprint) SetType(v string) {
	o.Type = &v
}

func (o AppBlueprint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppBlueprint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAppBlueprint struct {
	value *AppBlueprint
	isSet bool
}

func (v NullableAppBlueprint) Get() *AppBlueprint {
	return v.value
}

func (v *NullableAppBlueprint) Set(val *AppBlueprint) {
	v.value = val
	v.isSet = true
}

func (v NullableAppBlueprint) IsSet() bool {
	return v.isSet
}

func (v *NullableAppBlueprint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppBlueprint(val *AppBlueprint) *NullableAppBlueprint {
	return &NullableAppBlueprint{value: val, isSet: true}
}

func (v NullableAppBlueprint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppBlueprint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


