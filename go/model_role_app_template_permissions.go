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

// RoleAppTemplatePermissions struct for RoleAppTemplatePermissions
type RoleAppTemplatePermissions struct {
	Id *int64 `json:"id,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Access *string `json:"access,omitempty"`
}

// NewRoleAppTemplatePermissions instantiates a new RoleAppTemplatePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAppTemplatePermissions() *RoleAppTemplatePermissions {
	this := RoleAppTemplatePermissions{}
	return &this
}

// NewRoleAppTemplatePermissionsWithDefaults instantiates a new RoleAppTemplatePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAppTemplatePermissionsWithDefaults() *RoleAppTemplatePermissions {
	this := RoleAppTemplatePermissions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleAppTemplatePermissions) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAppTemplatePermissions) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleAppTemplatePermissions) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RoleAppTemplatePermissions) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAppTemplatePermissions) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAppTemplatePermissions) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *RoleAppTemplatePermissions) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *RoleAppTemplatePermissions) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *RoleAppTemplatePermissions) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *RoleAppTemplatePermissions) UnsetCode() {
	o.Code.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleAppTemplatePermissions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAppTemplatePermissions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleAppTemplatePermissions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleAppTemplatePermissions) SetName(v string) {
	o.Name = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *RoleAppTemplatePermissions) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAppTemplatePermissions) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *RoleAppTemplatePermissions) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *RoleAppTemplatePermissions) SetAccess(v string) {
	o.Access = &v
}

func (o RoleAppTemplatePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAppTemplatePermissions struct {
	value *RoleAppTemplatePermissions
	isSet bool
}

func (v NullableRoleAppTemplatePermissions) Get() *RoleAppTemplatePermissions {
	return v.value
}

func (v *NullableRoleAppTemplatePermissions) Set(val *RoleAppTemplatePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAppTemplatePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAppTemplatePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAppTemplatePermissions(val *RoleAppTemplatePermissions) *NullableRoleAppTemplatePermissions {
	return &NullableRoleAppTemplatePermissions{value: val, isSet: true}
}

func (v NullableRoleAppTemplatePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAppTemplatePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

