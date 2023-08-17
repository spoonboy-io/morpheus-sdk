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

// checks if the ApplyAppStateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyAppStateRequest{}

// ApplyAppStateRequest struct for ApplyAppStateRequest
type ApplyAppStateRequest struct {
	// Template Parameter object. A map of key-value pairs that correspond to the template variables i.e. tfvars
	TemplateParameter map[string]interface{} `json:"templateParameter,omitempty"`
}

// NewApplyAppStateRequest instantiates a new ApplyAppStateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyAppStateRequest() *ApplyAppStateRequest {
	this := ApplyAppStateRequest{}
	return &this
}

// NewApplyAppStateRequestWithDefaults instantiates a new ApplyAppStateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyAppStateRequestWithDefaults() *ApplyAppStateRequest {
	this := ApplyAppStateRequest{}
	return &this
}

// GetTemplateParameter returns the TemplateParameter field value if set, zero value otherwise.
func (o *ApplyAppStateRequest) GetTemplateParameter() map[string]interface{} {
	if o == nil || IsNil(o.TemplateParameter) {
		var ret map[string]interface{}
		return ret
	}
	return o.TemplateParameter
}

// GetTemplateParameterOk returns a tuple with the TemplateParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyAppStateRequest) GetTemplateParameterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TemplateParameter) {
		return map[string]interface{}{}, false
	}
	return o.TemplateParameter, true
}

// HasTemplateParameter returns a boolean if a field has been set.
func (o *ApplyAppStateRequest) HasTemplateParameter() bool {
	if o != nil && !IsNil(o.TemplateParameter) {
		return true
	}

	return false
}

// SetTemplateParameter gets a reference to the given map[string]interface{} and assigns it to the TemplateParameter field.
func (o *ApplyAppStateRequest) SetTemplateParameter(v map[string]interface{}) {
	o.TemplateParameter = v
}

func (o ApplyAppStateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyAppStateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateParameter) {
		toSerialize["templateParameter"] = o.TemplateParameter
	}
	return toSerialize, nil
}

type NullableApplyAppStateRequest struct {
	value *ApplyAppStateRequest
	isSet bool
}

func (v NullableApplyAppStateRequest) Get() *ApplyAppStateRequest {
	return v.value
}

func (v *NullableApplyAppStateRequest) Set(val *ApplyAppStateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyAppStateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyAppStateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyAppStateRequest(val *ApplyAppStateRequest) *NullableApplyAppStateRequest {
	return &NullableApplyAppStateRequest{value: val, isSet: true}
}

func (v NullableApplyAppStateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyAppStateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


