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

// InlineResponse200108NetworkFloatingIpCloud Cloud
type InlineResponse200108NetworkFloatingIpCloud struct {
	// Cloud ID
	Id *int64 `json:"id,omitempty"`
	// Cloud Name
	Name *string `json:"name,omitempty"`
	// Cloud Type Code
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse200108NetworkFloatingIpCloud instantiates a new InlineResponse200108NetworkFloatingIpCloud object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200108NetworkFloatingIpCloud() *InlineResponse200108NetworkFloatingIpCloud {
	this := InlineResponse200108NetworkFloatingIpCloud{}
	return &this
}

// NewInlineResponse200108NetworkFloatingIpCloudWithDefaults instantiates a new InlineResponse200108NetworkFloatingIpCloud object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200108NetworkFloatingIpCloudWithDefaults() *InlineResponse200108NetworkFloatingIpCloud {
	this := InlineResponse200108NetworkFloatingIpCloud{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse200108NetworkFloatingIpCloud) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200108NetworkFloatingIpCloud) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200108NetworkFloatingIpCloud) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200108NetworkFloatingIpCloud) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse200108NetworkFloatingIpCloud) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200108NetworkFloatingIpCloud struct {
	value *InlineResponse200108NetworkFloatingIpCloud
	isSet bool
}

func (v NullableInlineResponse200108NetworkFloatingIpCloud) Get() *InlineResponse200108NetworkFloatingIpCloud {
	return v.value
}

func (v *NullableInlineResponse200108NetworkFloatingIpCloud) Set(val *InlineResponse200108NetworkFloatingIpCloud) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200108NetworkFloatingIpCloud) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200108NetworkFloatingIpCloud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200108NetworkFloatingIpCloud(val *InlineResponse200108NetworkFloatingIpCloud) *NullableInlineResponse200108NetworkFloatingIpCloud {
	return &NullableInlineResponse200108NetworkFloatingIpCloud{value: val, isSet: true}
}

func (v NullableInlineResponse200108NetworkFloatingIpCloud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200108NetworkFloatingIpCloud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


