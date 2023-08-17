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

// checks if the AddApps200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddApps200Response{}

// AddApps200Response struct for AddApps200Response
type AddApps200Response struct {
	App *AppCreateResponse `json:"app,omitempty"`
}

// NewAddApps200Response instantiates a new AddApps200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddApps200Response() *AddApps200Response {
	this := AddApps200Response{}
	return &this
}

// NewAddApps200ResponseWithDefaults instantiates a new AddApps200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddApps200ResponseWithDefaults() *AddApps200Response {
	this := AddApps200Response{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AddApps200Response) GetApp() AppCreateResponse {
	if o == nil || IsNil(o.App) {
		var ret AppCreateResponse
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddApps200Response) GetAppOk() (*AppCreateResponse, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AddApps200Response) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppCreateResponse and assigns it to the App field.
func (o *AddApps200Response) SetApp(v AppCreateResponse) {
	o.App = &v
}

func (o AddApps200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddApps200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	return toSerialize, nil
}

type NullableAddApps200Response struct {
	value *AddApps200Response
	isSet bool
}

func (v NullableAddApps200Response) Get() *AddApps200Response {
	return v.value
}

func (v *NullableAddApps200Response) Set(val *AddApps200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddApps200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddApps200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddApps200Response(val *AddApps200Response) *NullableAddApps200Response {
	return &NullableAddApps200Response{value: val, isSet: true}
}

func (v NullableAddApps200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddApps200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

