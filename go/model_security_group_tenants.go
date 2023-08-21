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

// SecurityGroupTenants struct for SecurityGroupTenants
type SecurityGroupTenants struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	CanManage *bool `json:"canManage,omitempty"`
}

// NewSecurityGroupTenants instantiates a new SecurityGroupTenants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupTenants() *SecurityGroupTenants {
	this := SecurityGroupTenants{}
	return &this
}

// NewSecurityGroupTenantsWithDefaults instantiates a new SecurityGroupTenants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupTenantsWithDefaults() *SecurityGroupTenants {
	this := SecurityGroupTenants{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityGroupTenants) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupTenants) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityGroupTenants) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecurityGroupTenants) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityGroupTenants) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupTenants) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityGroupTenants) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityGroupTenants) SetName(v string) {
	o.Name = &v
}

// GetCanManage returns the CanManage field value if set, zero value otherwise.
func (o *SecurityGroupTenants) GetCanManage() bool {
	if o == nil || o.CanManage == nil {
		var ret bool
		return ret
	}
	return *o.CanManage
}

// GetCanManageOk returns a tuple with the CanManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupTenants) GetCanManageOk() (*bool, bool) {
	if o == nil || o.CanManage == nil {
		return nil, false
	}
	return o.CanManage, true
}

// HasCanManage returns a boolean if a field has been set.
func (o *SecurityGroupTenants) HasCanManage() bool {
	if o != nil && o.CanManage != nil {
		return true
	}

	return false
}

// SetCanManage gets a reference to the given bool and assigns it to the CanManage field.
func (o *SecurityGroupTenants) SetCanManage(v bool) {
	o.CanManage = &v
}

func (o SecurityGroupTenants) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CanManage != nil {
		toSerialize["canManage"] = o.CanManage
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityGroupTenants struct {
	value *SecurityGroupTenants
	isSet bool
}

func (v NullableSecurityGroupTenants) Get() *SecurityGroupTenants {
	return v.value
}

func (v *NullableSecurityGroupTenants) Set(val *SecurityGroupTenants) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupTenants) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupTenants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupTenants(val *SecurityGroupTenants) *NullableSecurityGroupTenants {
	return &NullableSecurityGroupTenants{value: val, isSet: true}
}

func (v NullableSecurityGroupTenants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupTenants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


