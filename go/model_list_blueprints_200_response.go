/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListBlueprints200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBlueprints200Response{}

// ListBlueprints200Response struct for ListBlueprints200Response
type ListBlueprints200Response struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Blueprints []Blueprint `json:"blueprints,omitempty"`
}

// NewListBlueprints200Response instantiates a new ListBlueprints200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBlueprints200Response() *ListBlueprints200Response {
	this := ListBlueprints200Response{}
	return &this
}

// NewListBlueprints200ResponseWithDefaults instantiates a new ListBlueprints200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBlueprints200ResponseWithDefaults() *ListBlueprints200Response {
	this := ListBlueprints200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListBlueprints200Response) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBlueprints200Response) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListBlueprints200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ListBlueprints200Response) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetBlueprints returns the Blueprints field value if set, zero value otherwise.
func (o *ListBlueprints200Response) GetBlueprints() []Blueprint {
	if o == nil || IsNil(o.Blueprints) {
		var ret []Blueprint
		return ret
	}
	return o.Blueprints
}

// GetBlueprintsOk returns a tuple with the Blueprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBlueprints200Response) GetBlueprintsOk() ([]Blueprint, bool) {
	if o == nil || IsNil(o.Blueprints) {
		return nil, false
	}
	return o.Blueprints, true
}

// HasBlueprints returns a boolean if a field has been set.
func (o *ListBlueprints200Response) HasBlueprints() bool {
	if o != nil && !IsNil(o.Blueprints) {
		return true
	}

	return false
}

// SetBlueprints gets a reference to the given []Blueprint and assigns it to the Blueprints field.
func (o *ListBlueprints200Response) SetBlueprints(v []Blueprint) {
	o.Blueprints = v
}

func (o ListBlueprints200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBlueprints200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Blueprints) {
		toSerialize["blueprints"] = o.Blueprints
	}
	return toSerialize, nil
}

type NullableListBlueprints200Response struct {
	value *ListBlueprints200Response
	isSet bool
}

func (v NullableListBlueprints200Response) Get() *ListBlueprints200Response {
	return v.value
}

func (v *NullableListBlueprints200Response) Set(val *ListBlueprints200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListBlueprints200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListBlueprints200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBlueprints200Response(val *ListBlueprints200Response) *NullableListBlueprints200Response {
	return &NullableListBlueprints200Response{value: val, isSet: true}
}

func (v NullableListBlueprints200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBlueprints200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

