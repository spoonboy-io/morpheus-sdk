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

// InlineResponse200137 struct for InlineResponse200137
type InlineResponse200137 struct {
	Server *Server `json:"server,omitempty"`
	Stats *map[string]interface{} `json:"stats,omitempty"`
}

// NewInlineResponse200137 instantiates a new InlineResponse200137 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200137() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// NewInlineResponse200137WithDefaults instantiates a new InlineResponse200137 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200137WithDefaults() *InlineResponse200137 {
	this := InlineResponse200137{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *InlineResponse200137) GetServer() Server {
	if o == nil || o.Server == nil {
		var ret Server
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetServerOk() (*Server, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *InlineResponse200137) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given Server and assigns it to the Server field.
func (o *InlineResponse200137) SetServer(v Server) {
	o.Server = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *InlineResponse200137) GetStats() map[string]interface{} {
	if o == nil || o.Stats == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200137) GetStatsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *InlineResponse200137) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given map[string]interface{} and assigns it to the Stats field.
func (o *InlineResponse200137) SetStats(v map[string]interface{}) {
	o.Stats = &v
}

func (o InlineResponse200137) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200137 struct {
	value *InlineResponse200137
	isSet bool
}

func (v NullableInlineResponse200137) Get() *InlineResponse200137 {
	return v.value
}

func (v *NullableInlineResponse200137) Set(val *InlineResponse200137) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200137) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200137) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200137(val *InlineResponse200137) *NullableInlineResponse200137 {
	return &NullableInlineResponse200137{value: val, isSet: true}
}

func (v NullableInlineResponse200137) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200137) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


