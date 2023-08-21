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

// CredentialUser struct for CredentialUser
type CredentialUser struct {
	Id *int64 `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// NewCredentialUser instantiates a new CredentialUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialUser() *CredentialUser {
	this := CredentialUser{}
	return &this
}

// NewCredentialUserWithDefaults instantiates a new CredentialUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialUserWithDefaults() *CredentialUser {
	this := CredentialUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialUser) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialUser) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CredentialUser) SetId(v int64) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CredentialUser) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialUser) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CredentialUser) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CredentialUser) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CredentialUser) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CredentialUser) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CredentialUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o CredentialUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialUser struct {
	value *CredentialUser
	isSet bool
}

func (v NullableCredentialUser) Get() *CredentialUser {
	return v.value
}

func (v *NullableCredentialUser) Set(val *CredentialUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialUser(val *CredentialUser) *NullableCredentialUser {
	return &NullableCredentialUser{value: val, isSet: true}
}

func (v NullableCredentialUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


