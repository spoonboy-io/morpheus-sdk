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

// checks if the ListClients200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClients200Response{}

// ListClients200Response struct for ListClients200Response
type ListClients200Response struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Clients []Client `json:"clients,omitempty"`
}

// NewListClients200Response instantiates a new ListClients200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClients200Response() *ListClients200Response {
	this := ListClients200Response{}
	return &this
}

// NewListClients200ResponseWithDefaults instantiates a new ListClients200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClients200ResponseWithDefaults() *ListClients200Response {
	this := ListClients200Response{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListClients200Response) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClients200Response) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListClients200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ListClients200Response) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ListClients200Response) GetClients() []Client {
	if o == nil || IsNil(o.Clients) {
		var ret []Client
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClients200Response) GetClientsOk() ([]Client, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ListClients200Response) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []Client and assigns it to the Clients field.
func (o *ListClients200Response) SetClients(v []Client) {
	o.Clients = v
}

func (o ListClients200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClients200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	return toSerialize, nil
}

type NullableListClients200Response struct {
	value *ListClients200Response
	isSet bool
}

func (v NullableListClients200Response) Get() *ListClients200Response {
	return v.value
}

func (v *NullableListClients200Response) Set(val *ListClients200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListClients200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListClients200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClients200Response(val *ListClients200Response) *NullableListClients200Response {
	return &NullableListClients200Response{value: val, isSet: true}
}

func (v NullableListClients200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClients200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

