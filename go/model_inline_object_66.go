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

// InlineObject66 The following parameters are available under the root context of the JSON body. The secret mount is capable of storing the entire JSON object as key=value pairs, or just a single string. To store a string instead, use the value query parameter instead of JSON, or pass type=string. There are a couple of special keys that the API will look for in the payload. The ttl key is a special key that if present in the payload will be parsed and used as the ttl parameter (lease duration in seconds). The value key is a special key that if present in the payload will be parsed and used as the secret data (instead of the entire payload). This is true when type=string. 
type InlineObject66 struct {
	// Time to Live. The lease duration in seconds, or a human readable format eg. 15m, 8h, 7d. The default is 0 meaning Never expires. This only is applied if the cypher does not yet exist and is created. 
	Ttl *OneOfstringlong `json:"ttl,omitempty"`
	// The secret value to be stored. This is only parsed if type is passed as `string`.
	Value *string `json:"value,omitempty"`
}

// NewInlineObject66 instantiates a new InlineObject66 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject66() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// NewInlineObject66WithDefaults instantiates a new InlineObject66 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject66WithDefaults() *InlineObject66 {
	this := InlineObject66{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *InlineObject66) GetTtl() OneOfstringlong {
	if o == nil || o.Ttl == nil {
		var ret OneOfstringlong
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetTtlOk() (*OneOfstringlong, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *InlineObject66) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given OneOfstringlong and assigns it to the Ttl field.
func (o *InlineObject66) SetTtl(v OneOfstringlong) {
	o.Ttl = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineObject66) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject66) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineObject66) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineObject66) SetValue(v string) {
	o.Value = &v
}

func (o InlineObject66) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject66 struct {
	value *InlineObject66
	isSet bool
}

func (v NullableInlineObject66) Get() *InlineObject66 {
	return v.value
}

func (v *NullableInlineObject66) Set(val *InlineObject66) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject66) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject66) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject66(val *InlineObject66) *NullableInlineObject66 {
	return &NullableInlineObject66{value: val, isSet: true}
}

func (v NullableInlineObject66) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject66) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

