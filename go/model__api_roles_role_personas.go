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

// ApiRolesRolePersonas struct for ApiRolesRolePersonas
type ApiRolesRolePersonas struct {
	// `code` of the persona
	Code *string `json:"code,omitempty"`
	// The new access level.
	Access string `json:"access"`
}

// NewApiRolesRolePersonas instantiates a new ApiRolesRolePersonas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRolesRolePersonas(access string, ) *ApiRolesRolePersonas {
	this := ApiRolesRolePersonas{}
	this.Access = access
	return &this
}

// NewApiRolesRolePersonasWithDefaults instantiates a new ApiRolesRolePersonas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRolesRolePersonasWithDefaults() *ApiRolesRolePersonas {
	this := ApiRolesRolePersonas{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiRolesRolePersonas) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRolesRolePersonas) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiRolesRolePersonas) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiRolesRolePersonas) SetCode(v string) {
	o.Code = &v
}

// GetAccess returns the Access field value
func (o *ApiRolesRolePersonas) GetAccess() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *ApiRolesRolePersonas) GetAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *ApiRolesRolePersonas) SetAccess(v string) {
	o.Access = v
}

func (o ApiRolesRolePersonas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableApiRolesRolePersonas struct {
	value *ApiRolesRolePersonas
	isSet bool
}

func (v NullableApiRolesRolePersonas) Get() *ApiRolesRolePersonas {
	return v.value
}

func (v *NullableApiRolesRolePersonas) Set(val *ApiRolesRolePersonas) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRolesRolePersonas) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRolesRolePersonas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRolesRolePersonas(val *ApiRolesRolePersonas) *NullableApiRolesRolePersonas {
	return &NullableApiRolesRolePersonas{value: val, isSet: true}
}

func (v NullableApiRolesRolePersonas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRolesRolePersonas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


