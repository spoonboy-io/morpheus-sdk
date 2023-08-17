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

// checks if the ApprovalItemApprovalItemReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApprovalItemApprovalItemReference{}

// ApprovalItemApprovalItemReference struct for ApprovalItemApprovalItemReference
type ApprovalItemApprovalItemReference struct {
	Id *int64 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
}

// NewApprovalItemApprovalItemReference instantiates a new ApprovalItemApprovalItemReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalItemApprovalItemReference() *ApprovalItemApprovalItemReference {
	this := ApprovalItemApprovalItemReference{}
	return &this
}

// NewApprovalItemApprovalItemReferenceWithDefaults instantiates a new ApprovalItemApprovalItemReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalItemApprovalItemReferenceWithDefaults() *ApprovalItemApprovalItemReference {
	this := ApprovalItemApprovalItemReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApprovalItemApprovalItemReference) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItemApprovalItemReference) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApprovalItemApprovalItemReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApprovalItemApprovalItemReference) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApprovalItemApprovalItemReference) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItemApprovalItemReference) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApprovalItemApprovalItemReference) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApprovalItemApprovalItemReference) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApprovalItemApprovalItemReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItemApprovalItemReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApprovalItemApprovalItemReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApprovalItemApprovalItemReference) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApprovalItemApprovalItemReference) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalItemApprovalItemReference) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApprovalItemApprovalItemReference) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApprovalItemApprovalItemReference) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o ApprovalItemApprovalItemReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApprovalItemApprovalItemReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	return toSerialize, nil
}

type NullableApprovalItemApprovalItemReference struct {
	value *ApprovalItemApprovalItemReference
	isSet bool
}

func (v NullableApprovalItemApprovalItemReference) Get() *ApprovalItemApprovalItemReference {
	return v.value
}

func (v *NullableApprovalItemApprovalItemReference) Set(val *ApprovalItemApprovalItemReference) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalItemApprovalItemReference) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalItemApprovalItemReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalItemApprovalItemReference(val *ApprovalItemApprovalItemReference) *NullableApprovalItemApprovalItemReference {
	return &NullableApprovalItemApprovalItemReference{value: val, isSet: true}
}

func (v NullableApprovalItemApprovalItemReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalItemApprovalItemReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

