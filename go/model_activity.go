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

// Activity struct for Activity
type Activity struct {
	Activity *[]ActivityActivity `json:"activity,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise.
func (o *Activity) GetActivity() []ActivityActivity {
	if o == nil || o.Activity == nil {
		var ret []ActivityActivity
		return ret
	}
	return *o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActivityOk() (*[]ActivityActivity, bool) {
	if o == nil || o.Activity == nil {
		return nil, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *Activity) HasActivity() bool {
	if o != nil && o.Activity != nil {
		return true
	}

	return false
}

// SetActivity gets a reference to the given []ActivityActivity and assigns it to the Activity field.
func (o *Activity) SetActivity(v []ActivityActivity) {
	o.Activity = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activity != nil {
		toSerialize["activity"] = o.Activity
	}
	return json.Marshal(toSerialize)
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

