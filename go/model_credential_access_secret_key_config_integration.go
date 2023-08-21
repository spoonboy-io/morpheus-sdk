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

// CredentialAccessSecretKeyConfigIntegration Credential Store. ID of a Credential Integration. This can be set to store the credential in an external store. 
type CredentialAccessSecretKeyConfigIntegration struct {
	Id *OneOfstringlong `json:"id,omitempty"`
}

// NewCredentialAccessSecretKeyConfigIntegration instantiates a new CredentialAccessSecretKeyConfigIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialAccessSecretKeyConfigIntegration() *CredentialAccessSecretKeyConfigIntegration {
	this := CredentialAccessSecretKeyConfigIntegration{}
	return &this
}

// NewCredentialAccessSecretKeyConfigIntegrationWithDefaults instantiates a new CredentialAccessSecretKeyConfigIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialAccessSecretKeyConfigIntegrationWithDefaults() *CredentialAccessSecretKeyConfigIntegration {
	this := CredentialAccessSecretKeyConfigIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialAccessSecretKeyConfigIntegration) GetId() OneOfstringlong {
	if o == nil || o.Id == nil {
		var ret OneOfstringlong
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialAccessSecretKeyConfigIntegration) GetIdOk() (*OneOfstringlong, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialAccessSecretKeyConfigIntegration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given OneOfstringlong and assigns it to the Id field.
func (o *CredentialAccessSecretKeyConfigIntegration) SetId(v OneOfstringlong) {
	o.Id = &v
}

func (o CredentialAccessSecretKeyConfigIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialAccessSecretKeyConfigIntegration struct {
	value *CredentialAccessSecretKeyConfigIntegration
	isSet bool
}

func (v NullableCredentialAccessSecretKeyConfigIntegration) Get() *CredentialAccessSecretKeyConfigIntegration {
	return v.value
}

func (v *NullableCredentialAccessSecretKeyConfigIntegration) Set(val *CredentialAccessSecretKeyConfigIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialAccessSecretKeyConfigIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialAccessSecretKeyConfigIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialAccessSecretKeyConfigIntegration(val *CredentialAccessSecretKeyConfigIntegration) *NullableCredentialAccessSecretKeyConfigIntegration {
	return &NullableCredentialAccessSecretKeyConfigIntegration{value: val, isSet: true}
}

func (v NullableCredentialAccessSecretKeyConfigIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialAccessSecretKeyConfigIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


