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

// ApiServersIdMakeManagedServerPlan struct for ApiServersIdMakeManagedServerPlan
type ApiServersIdMakeManagedServerPlan struct {
	// Service Plan to assign to the server
	Id *int64 `json:"id,omitempty"`
}

// NewApiServersIdMakeManagedServerPlan instantiates a new ApiServersIdMakeManagedServerPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiServersIdMakeManagedServerPlan() *ApiServersIdMakeManagedServerPlan {
	this := ApiServersIdMakeManagedServerPlan{}
	return &this
}

// NewApiServersIdMakeManagedServerPlanWithDefaults instantiates a new ApiServersIdMakeManagedServerPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiServersIdMakeManagedServerPlanWithDefaults() *ApiServersIdMakeManagedServerPlan {
	this := ApiServersIdMakeManagedServerPlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiServersIdMakeManagedServerPlan) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiServersIdMakeManagedServerPlan) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiServersIdMakeManagedServerPlan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApiServersIdMakeManagedServerPlan) SetId(v int64) {
	o.Id = &v
}

func (o ApiServersIdMakeManagedServerPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableApiServersIdMakeManagedServerPlan struct {
	value *ApiServersIdMakeManagedServerPlan
	isSet bool
}

func (v NullableApiServersIdMakeManagedServerPlan) Get() *ApiServersIdMakeManagedServerPlan {
	return v.value
}

func (v *NullableApiServersIdMakeManagedServerPlan) Set(val *ApiServersIdMakeManagedServerPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableApiServersIdMakeManagedServerPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableApiServersIdMakeManagedServerPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiServersIdMakeManagedServerPlan(val *ApiServersIdMakeManagedServerPlan) *NullableApiServersIdMakeManagedServerPlan {
	return &NullableApiServersIdMakeManagedServerPlan{value: val, isSet: true}
}

func (v NullableApiServersIdMakeManagedServerPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiServersIdMakeManagedServerPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

