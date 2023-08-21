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

// InlineResponse20038 struct for InlineResponse20038
type InlineResponse20038 struct {
	Deployment *Deployment `json:"deployment,omitempty"`
}

// NewInlineResponse20038 instantiates a new InlineResponse20038 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20038() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// NewInlineResponse20038WithDefaults instantiates a new InlineResponse20038 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20038WithDefaults() *InlineResponse20038 {
	this := InlineResponse20038{}
	return &this
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *InlineResponse20038) GetDeployment() Deployment {
	if o == nil || o.Deployment == nil {
		var ret Deployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20038) GetDeploymentOk() (*Deployment, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *InlineResponse20038) HasDeployment() bool {
	if o != nil && o.Deployment != nil {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given Deployment and assigns it to the Deployment field.
func (o *InlineResponse20038) SetDeployment(v Deployment) {
	o.Deployment = &v
}

func (o InlineResponse20038) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20038 struct {
	value *InlineResponse20038
	isSet bool
}

func (v NullableInlineResponse20038) Get() *InlineResponse20038 {
	return v.value
}

func (v *NullableInlineResponse20038) Set(val *InlineResponse20038) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20038) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20038) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20038(val *InlineResponse20038) *NullableInlineResponse20038 {
	return &NullableInlineResponse20038{value: val, isSet: true}
}

func (v NullableInlineResponse20038) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20038) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


