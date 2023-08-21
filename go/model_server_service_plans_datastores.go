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

// ServerServicePlansDatastores struct for ServerServicePlansDatastores
type ServerServicePlansDatastores struct {
	Cluster *[]map[string]interface{} `json:"cluster,omitempty"`
	Store *[]map[string]interface{} `json:"store,omitempty"`
}

// NewServerServicePlansDatastores instantiates a new ServerServicePlansDatastores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerServicePlansDatastores() *ServerServicePlansDatastores {
	this := ServerServicePlansDatastores{}
	return &this
}

// NewServerServicePlansDatastoresWithDefaults instantiates a new ServerServicePlansDatastores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerServicePlansDatastoresWithDefaults() *ServerServicePlansDatastores {
	this := ServerServicePlansDatastores{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ServerServicePlansDatastores) GetCluster() []map[string]interface{} {
	if o == nil || o.Cluster == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServicePlansDatastores) GetClusterOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ServerServicePlansDatastores) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given []map[string]interface{} and assigns it to the Cluster field.
func (o *ServerServicePlansDatastores) SetCluster(v []map[string]interface{}) {
	o.Cluster = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *ServerServicePlansDatastores) GetStore() []map[string]interface{} {
	if o == nil || o.Store == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServicePlansDatastores) GetStoreOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Store == nil {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *ServerServicePlansDatastores) HasStore() bool {
	if o != nil && o.Store != nil {
		return true
	}

	return false
}

// SetStore gets a reference to the given []map[string]interface{} and assigns it to the Store field.
func (o *ServerServicePlansDatastores) SetStore(v []map[string]interface{}) {
	o.Store = &v
}

func (o ServerServicePlansDatastores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Store != nil {
		toSerialize["store"] = o.Store
	}
	return json.Marshal(toSerialize)
}

type NullableServerServicePlansDatastores struct {
	value *ServerServicePlansDatastores
	isSet bool
}

func (v NullableServerServicePlansDatastores) Get() *ServerServicePlansDatastores {
	return v.value
}

func (v *NullableServerServicePlansDatastores) Set(val *ServerServicePlansDatastores) {
	v.value = val
	v.isSet = true
}

func (v NullableServerServicePlansDatastores) IsSet() bool {
	return v.isSet
}

func (v *NullableServerServicePlansDatastores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerServicePlansDatastores(val *ServerServicePlansDatastores) *NullableServerServicePlansDatastores {
	return &NullableServerServicePlansDatastores{value: val, isSet: true}
}

func (v NullableServerServicePlansDatastores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerServicePlansDatastores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

