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

// ApiKeyPairsKeyPair struct for ApiKeyPairsKeyPair
type ApiKeyPairsKeyPair struct {
	Name string `json:"name"`
	PublicKey string `json:"publicKey"`
	PrivateKey *string `json:"privateKey,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewApiKeyPairsKeyPair instantiates a new ApiKeyPairsKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyPairsKeyPair(name string, publicKey string, ) *ApiKeyPairsKeyPair {
	this := ApiKeyPairsKeyPair{}
	this.Name = name
	this.PublicKey = publicKey
	return &this
}

// NewApiKeyPairsKeyPairWithDefaults instantiates a new ApiKeyPairsKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyPairsKeyPairWithDefaults() *ApiKeyPairsKeyPair {
	this := ApiKeyPairsKeyPair{}
	return &this
}

// GetName returns the Name field value
func (o *ApiKeyPairsKeyPair) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiKeyPairsKeyPair) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiKeyPairsKeyPair) SetName(v string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *ApiKeyPairsKeyPair) GetPublicKey() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ApiKeyPairsKeyPair) GetPublicKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ApiKeyPairsKeyPair) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *ApiKeyPairsKeyPair) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyPairsKeyPair) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *ApiKeyPairsKeyPair) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *ApiKeyPairsKeyPair) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *ApiKeyPairsKeyPair) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyPairsKeyPair) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *ApiKeyPairsKeyPair) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *ApiKeyPairsKeyPair) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o ApiKeyPairsKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["publicKey"] = o.PublicKey
	}
	if o.PrivateKey != nil {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableApiKeyPairsKeyPair struct {
	value *ApiKeyPairsKeyPair
	isSet bool
}

func (v NullableApiKeyPairsKeyPair) Get() *ApiKeyPairsKeyPair {
	return v.value
}

func (v *NullableApiKeyPairsKeyPair) Set(val *ApiKeyPairsKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyPairsKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyPairsKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyPairsKeyPair(val *ApiKeyPairsKeyPair) *NullableApiKeyPairsKeyPair {
	return &NullableApiKeyPairsKeyPair{value: val, isSet: true}
}

func (v NullableApiKeyPairsKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyPairsKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

