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

// InlineResponse20040AppDeployDeployment struct for InlineResponse20040AppDeployDeployment
type InlineResponse20040AppDeployDeployment struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DeployType *string `json:"deployType,omitempty"`
}

// NewInlineResponse20040AppDeployDeployment instantiates a new InlineResponse20040AppDeployDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20040AppDeployDeployment() *InlineResponse20040AppDeployDeployment {
	this := InlineResponse20040AppDeployDeployment{}
	return &this
}

// NewInlineResponse20040AppDeployDeploymentWithDefaults instantiates a new InlineResponse20040AppDeployDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20040AppDeployDeploymentWithDefaults() *InlineResponse20040AppDeployDeployment {
	this := InlineResponse20040AppDeployDeployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20040AppDeployDeployment) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040AppDeployDeployment) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20040AppDeployDeployment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20040AppDeployDeployment) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20040AppDeployDeployment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040AppDeployDeployment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20040AppDeployDeployment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20040AppDeployDeployment) SetName(v string) {
	o.Name = &v
}

// GetDeployType returns the DeployType field value if set, zero value otherwise.
func (o *InlineResponse20040AppDeployDeployment) GetDeployType() string {
	if o == nil || o.DeployType == nil {
		var ret string
		return ret
	}
	return *o.DeployType
}

// GetDeployTypeOk returns a tuple with the DeployType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040AppDeployDeployment) GetDeployTypeOk() (*string, bool) {
	if o == nil || o.DeployType == nil {
		return nil, false
	}
	return o.DeployType, true
}

// HasDeployType returns a boolean if a field has been set.
func (o *InlineResponse20040AppDeployDeployment) HasDeployType() bool {
	if o != nil && o.DeployType != nil {
		return true
	}

	return false
}

// SetDeployType gets a reference to the given string and assigns it to the DeployType field.
func (o *InlineResponse20040AppDeployDeployment) SetDeployType(v string) {
	o.DeployType = &v
}

func (o InlineResponse20040AppDeployDeployment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DeployType != nil {
		toSerialize["deployType"] = o.DeployType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20040AppDeployDeployment struct {
	value *InlineResponse20040AppDeployDeployment
	isSet bool
}

func (v NullableInlineResponse20040AppDeployDeployment) Get() *InlineResponse20040AppDeployDeployment {
	return v.value
}

func (v *NullableInlineResponse20040AppDeployDeployment) Set(val *InlineResponse20040AppDeployDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20040AppDeployDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20040AppDeployDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20040AppDeployDeployment(val *InlineResponse20040AppDeployDeployment) *NullableInlineResponse20040AppDeployDeployment {
	return &NullableInlineResponse20040AppDeployDeployment{value: val, isSet: true}
}

func (v NullableInlineResponse20040AppDeployDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20040AppDeployDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


