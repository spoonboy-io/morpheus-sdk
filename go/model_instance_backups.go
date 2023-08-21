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

// InstanceBackups struct for InstanceBackups
type InstanceBackups struct {
	Instance *InstanceBackupsInstance `json:"instance,omitempty"`
	// List of backup objects
	Backups *[]map[string]interface{} `json:"backups,omitempty"`
}

// NewInstanceBackups instantiates a new InstanceBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceBackups() *InstanceBackups {
	this := InstanceBackups{}
	return &this
}

// NewInstanceBackupsWithDefaults instantiates a new InstanceBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceBackupsWithDefaults() *InstanceBackups {
	this := InstanceBackups{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *InstanceBackups) GetInstance() InstanceBackupsInstance {
	if o == nil || o.Instance == nil {
		var ret InstanceBackupsInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceBackups) GetInstanceOk() (*InstanceBackupsInstance, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *InstanceBackups) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given InstanceBackupsInstance and assigns it to the Instance field.
func (o *InstanceBackups) SetInstance(v InstanceBackupsInstance) {
	o.Instance = &v
}

// GetBackups returns the Backups field value if set, zero value otherwise.
func (o *InstanceBackups) GetBackups() []map[string]interface{} {
	if o == nil || o.Backups == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceBackups) GetBackupsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Backups == nil {
		return nil, false
	}
	return o.Backups, true
}

// HasBackups returns a boolean if a field has been set.
func (o *InstanceBackups) HasBackups() bool {
	if o != nil && o.Backups != nil {
		return true
	}

	return false
}

// SetBackups gets a reference to the given []map[string]interface{} and assigns it to the Backups field.
func (o *InstanceBackups) SetBackups(v []map[string]interface{}) {
	o.Backups = &v
}

func (o InstanceBackups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.Backups != nil {
		toSerialize["backups"] = o.Backups
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceBackups struct {
	value *InstanceBackups
	isSet bool
}

func (v NullableInstanceBackups) Get() *InstanceBackups {
	return v.value
}

func (v *NullableInstanceBackups) Set(val *InstanceBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceBackups(val *InstanceBackups) *NullableInstanceBackups {
	return &NullableInstanceBackups{value: val, isSet: true}
}

func (v NullableInstanceBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

