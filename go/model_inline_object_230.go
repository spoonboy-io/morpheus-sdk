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

// InlineObject230 struct for InlineObject230
type InlineObject230 struct {
	Certificate *ApiCertificatesCertificate `json:"certificate,omitempty"`
}

// NewInlineObject230 instantiates a new InlineObject230 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject230() *InlineObject230 {
	this := InlineObject230{}
	return &this
}

// NewInlineObject230WithDefaults instantiates a new InlineObject230 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject230WithDefaults() *InlineObject230 {
	this := InlineObject230{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *InlineObject230) GetCertificate() ApiCertificatesCertificate {
	if o == nil || o.Certificate == nil {
		var ret ApiCertificatesCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject230) GetCertificateOk() (*ApiCertificatesCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *InlineObject230) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given ApiCertificatesCertificate and assigns it to the Certificate field.
func (o *InlineObject230) SetCertificate(v ApiCertificatesCertificate) {
	o.Certificate = &v
}

func (o InlineObject230) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject230 struct {
	value *InlineObject230
	isSet bool
}

func (v NullableInlineObject230) Get() *InlineObject230 {
	return v.value
}

func (v *NullableInlineObject230) Set(val *InlineObject230) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject230) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject230) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject230(val *InlineObject230) *NullableInlineObject230 {
	return &NullableInlineObject230{value: val, isSet: true}
}

func (v NullableInlineObject230) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject230) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


