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

// RoleSites struct for RoleSites
type RoleSites struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Access *string `json:"access,omitempty"`
}

// NewRoleSites instantiates a new RoleSites object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleSites() *RoleSites {
	this := RoleSites{}
	return &this
}

// NewRoleSitesWithDefaults instantiates a new RoleSites object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleSitesWithDefaults() *RoleSites {
	this := RoleSites{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleSites) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSites) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleSites) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RoleSites) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleSites) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSites) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleSites) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleSites) SetName(v string) {
	o.Name = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *RoleSites) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSites) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *RoleSites) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *RoleSites) SetAccess(v string) {
	o.Access = &v
}

func (o RoleSites) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRoleSites struct {
	value *RoleSites
	isSet bool
}

func (v NullableRoleSites) Get() *RoleSites {
	return v.value
}

func (v *NullableRoleSites) Set(val *RoleSites) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleSites) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleSites) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleSites(val *RoleSites) *NullableRoleSites {
	return &NullableRoleSites{value: val, isSet: true}
}

func (v NullableRoleSites) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleSites) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

