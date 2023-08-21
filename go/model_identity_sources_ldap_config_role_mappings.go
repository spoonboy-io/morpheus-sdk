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

// IdentitySourcesLDAPConfigRoleMappings struct for IdentitySourcesLDAPConfigRoleMappings
type IdentitySourcesLDAPConfigRoleMappings struct {
	SourceRoleName NullableString `json:"sourceRoleName,omitempty"`
	SourceRoleFqn *string `json:"sourceRoleFqn,omitempty"`
	MappedRole *IdentitySourcesLDAPConfigDefaultAccountRole `json:"mappedRole,omitempty"`
}

// NewIdentitySourcesLDAPConfigRoleMappings instantiates a new IdentitySourcesLDAPConfigRoleMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourcesLDAPConfigRoleMappings() *IdentitySourcesLDAPConfigRoleMappings {
	this := IdentitySourcesLDAPConfigRoleMappings{}
	return &this
}

// NewIdentitySourcesLDAPConfigRoleMappingsWithDefaults instantiates a new IdentitySourcesLDAPConfigRoleMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourcesLDAPConfigRoleMappingsWithDefaults() *IdentitySourcesLDAPConfigRoleMappings {
	this := IdentitySourcesLDAPConfigRoleMappings{}
	return &this
}

// GetSourceRoleName returns the SourceRoleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleName() string {
	if o == nil || o.SourceRoleName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceRoleName.Get()
}

// GetSourceRoleNameOk returns a tuple with the SourceRoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceRoleName.Get(), o.SourceRoleName.IsSet()
}

// HasSourceRoleName returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfigRoleMappings) HasSourceRoleName() bool {
	if o != nil && o.SourceRoleName.IsSet() {
		return true
	}

	return false
}

// SetSourceRoleName gets a reference to the given NullableString and assigns it to the SourceRoleName field.
func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleName(v string) {
	o.SourceRoleName.Set(&v)
}
// SetSourceRoleNameNil sets the value for SourceRoleName to be an explicit nil
func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleNameNil() {
	o.SourceRoleName.Set(nil)
}

// UnsetSourceRoleName ensures that no value is present for SourceRoleName, not even an explicit nil
func (o *IdentitySourcesLDAPConfigRoleMappings) UnsetSourceRoleName() {
	o.SourceRoleName.Unset()
}

// GetSourceRoleFqn returns the SourceRoleFqn field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleFqn() string {
	if o == nil || o.SourceRoleFqn == nil {
		var ret string
		return ret
	}
	return *o.SourceRoleFqn
}

// GetSourceRoleFqnOk returns a tuple with the SourceRoleFqn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfigRoleMappings) GetSourceRoleFqnOk() (*string, bool) {
	if o == nil || o.SourceRoleFqn == nil {
		return nil, false
	}
	return o.SourceRoleFqn, true
}

// HasSourceRoleFqn returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfigRoleMappings) HasSourceRoleFqn() bool {
	if o != nil && o.SourceRoleFqn != nil {
		return true
	}

	return false
}

// SetSourceRoleFqn gets a reference to the given string and assigns it to the SourceRoleFqn field.
func (o *IdentitySourcesLDAPConfigRoleMappings) SetSourceRoleFqn(v string) {
	o.SourceRoleFqn = &v
}

// GetMappedRole returns the MappedRole field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfigRoleMappings) GetMappedRole() IdentitySourcesLDAPConfigDefaultAccountRole {
	if o == nil || o.MappedRole == nil {
		var ret IdentitySourcesLDAPConfigDefaultAccountRole
		return ret
	}
	return *o.MappedRole
}

// GetMappedRoleOk returns a tuple with the MappedRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfigRoleMappings) GetMappedRoleOk() (*IdentitySourcesLDAPConfigDefaultAccountRole, bool) {
	if o == nil || o.MappedRole == nil {
		return nil, false
	}
	return o.MappedRole, true
}

// HasMappedRole returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfigRoleMappings) HasMappedRole() bool {
	if o != nil && o.MappedRole != nil {
		return true
	}

	return false
}

// SetMappedRole gets a reference to the given IdentitySourcesLDAPConfigDefaultAccountRole and assigns it to the MappedRole field.
func (o *IdentitySourcesLDAPConfigRoleMappings) SetMappedRole(v IdentitySourcesLDAPConfigDefaultAccountRole) {
	o.MappedRole = &v
}

func (o IdentitySourcesLDAPConfigRoleMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceRoleName.IsSet() {
		toSerialize["sourceRoleName"] = o.SourceRoleName.Get()
	}
	if o.SourceRoleFqn != nil {
		toSerialize["sourceRoleFqn"] = o.SourceRoleFqn
	}
	if o.MappedRole != nil {
		toSerialize["mappedRole"] = o.MappedRole
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySourcesLDAPConfigRoleMappings struct {
	value *IdentitySourcesLDAPConfigRoleMappings
	isSet bool
}

func (v NullableIdentitySourcesLDAPConfigRoleMappings) Get() *IdentitySourcesLDAPConfigRoleMappings {
	return v.value
}

func (v *NullableIdentitySourcesLDAPConfigRoleMappings) Set(val *IdentitySourcesLDAPConfigRoleMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourcesLDAPConfigRoleMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourcesLDAPConfigRoleMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourcesLDAPConfigRoleMappings(val *IdentitySourcesLDAPConfigRoleMappings) *NullableIdentitySourcesLDAPConfigRoleMappings {
	return &NullableIdentitySourcesLDAPConfigRoleMappings{value: val, isSet: true}
}

func (v NullableIdentitySourcesLDAPConfigRoleMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourcesLDAPConfigRoleMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


