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

// ZoneDatastoreTenants struct for ZoneDatastoreTenants
type ZoneDatastoreTenants struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DefaultStore *bool `json:"defaultStore,omitempty"`
	DefaultTarget *bool `json:"defaultTarget,omitempty"`
}

// NewZoneDatastoreTenants instantiates a new ZoneDatastoreTenants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneDatastoreTenants() *ZoneDatastoreTenants {
	this := ZoneDatastoreTenants{}
	return &this
}

// NewZoneDatastoreTenantsWithDefaults instantiates a new ZoneDatastoreTenants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneDatastoreTenantsWithDefaults() *ZoneDatastoreTenants {
	this := ZoneDatastoreTenants{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneDatastoreTenants) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneDatastoreTenants) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneDatastoreTenants) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneDatastoreTenants) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneDatastoreTenants) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneDatastoreTenants) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneDatastoreTenants) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneDatastoreTenants) SetName(v string) {
	o.Name = &v
}

// GetDefaultStore returns the DefaultStore field value if set, zero value otherwise.
func (o *ZoneDatastoreTenants) GetDefaultStore() bool {
	if o == nil || o.DefaultStore == nil {
		var ret bool
		return ret
	}
	return *o.DefaultStore
}

// GetDefaultStoreOk returns a tuple with the DefaultStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneDatastoreTenants) GetDefaultStoreOk() (*bool, bool) {
	if o == nil || o.DefaultStore == nil {
		return nil, false
	}
	return o.DefaultStore, true
}

// HasDefaultStore returns a boolean if a field has been set.
func (o *ZoneDatastoreTenants) HasDefaultStore() bool {
	if o != nil && o.DefaultStore != nil {
		return true
	}

	return false
}

// SetDefaultStore gets a reference to the given bool and assigns it to the DefaultStore field.
func (o *ZoneDatastoreTenants) SetDefaultStore(v bool) {
	o.DefaultStore = &v
}

// GetDefaultTarget returns the DefaultTarget field value if set, zero value otherwise.
func (o *ZoneDatastoreTenants) GetDefaultTarget() bool {
	if o == nil || o.DefaultTarget == nil {
		var ret bool
		return ret
	}
	return *o.DefaultTarget
}

// GetDefaultTargetOk returns a tuple with the DefaultTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneDatastoreTenants) GetDefaultTargetOk() (*bool, bool) {
	if o == nil || o.DefaultTarget == nil {
		return nil, false
	}
	return o.DefaultTarget, true
}

// HasDefaultTarget returns a boolean if a field has been set.
func (o *ZoneDatastoreTenants) HasDefaultTarget() bool {
	if o != nil && o.DefaultTarget != nil {
		return true
	}

	return false
}

// SetDefaultTarget gets a reference to the given bool and assigns it to the DefaultTarget field.
func (o *ZoneDatastoreTenants) SetDefaultTarget(v bool) {
	o.DefaultTarget = &v
}

func (o ZoneDatastoreTenants) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultStore != nil {
		toSerialize["defaultStore"] = o.DefaultStore
	}
	if o.DefaultTarget != nil {
		toSerialize["defaultTarget"] = o.DefaultTarget
	}
	return json.Marshal(toSerialize)
}

type NullableZoneDatastoreTenants struct {
	value *ZoneDatastoreTenants
	isSet bool
}

func (v NullableZoneDatastoreTenants) Get() *ZoneDatastoreTenants {
	return v.value
}

func (v *NullableZoneDatastoreTenants) Set(val *ZoneDatastoreTenants) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneDatastoreTenants) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneDatastoreTenants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneDatastoreTenants(val *ZoneDatastoreTenants) *NullableZoneDatastoreTenants {
	return &NullableZoneDatastoreTenants{value: val, isSet: true}
}

func (v NullableZoneDatastoreTenants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneDatastoreTenants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

