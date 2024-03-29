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

// InlineObject64 struct for InlineObject64
type InlineObject64 struct {
	// Payload for creating a new credential
	Credential OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config `json:"credential"`
}

// NewInlineObject64 instantiates a new InlineObject64 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject64(credential OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config, ) *InlineObject64 {
	this := InlineObject64{}
	this.Credential = credential
	return &this
}

// NewInlineObject64WithDefaults instantiates a new InlineObject64 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject64WithDefaults() *InlineObject64 {
	this := InlineObject64{}
	return &this
}

// GetCredential returns the Credential field value
func (o *InlineObject64) GetCredential() OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config {
	if o == nil  {
		var ret OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config
		return ret
	}

	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value
// and a boolean to check if the value has been set.
func (o *InlineObject64) GetCredentialOk() (*OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Credential, true
}

// SetCredential sets field value
func (o *InlineObject64) SetCredential(v OneOfcredentialAccessSecretKeyConfigcredentialClientIDSecretConfigcredentialEmailPrivateKeyConfigcredentialTenantUsernameKeypairConfigcredentialUsernameAPIKeyConfigcredentialUsernameKeypairConfigcredentialUsernamePasswordConfigcredentialUsernamePasswordKeypairConfigcredentialOauth2Config) {
	o.Credential = v
}

func (o InlineObject64) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject64 struct {
	value *InlineObject64
	isSet bool
}

func (v NullableInlineObject64) Get() *InlineObject64 {
	return v.value
}

func (v *NullableInlineObject64) Set(val *InlineObject64) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject64) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject64(val *InlineObject64) *NullableInlineObject64 {
	return &NullableInlineObject64{value: val, isSet: true}
}

func (v NullableInlineObject64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


