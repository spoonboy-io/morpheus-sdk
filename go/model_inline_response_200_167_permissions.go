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

// InlineResponse200167Permissions struct for InlineResponse200167Permissions
type InlineResponse200167Permissions struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Access *string `json:"access,omitempty"`
}

// NewInlineResponse200167Permissions instantiates a new InlineResponse200167Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200167Permissions() *InlineResponse200167Permissions {
	this := InlineResponse200167Permissions{}
	return &this
}

// NewInlineResponse200167PermissionsWithDefaults instantiates a new InlineResponse200167Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200167PermissionsWithDefaults() *InlineResponse200167Permissions {
	this := InlineResponse200167Permissions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200167Permissions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167Permissions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200167Permissions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200167Permissions) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse200167Permissions) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167Permissions) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse200167Permissions) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse200167Permissions) SetCode(v string) {
	o.Code = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200167Permissions) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200167Permissions) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200167Permissions) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200167Permissions) SetAccess(v string) {
	o.Access = &v
}

func (o InlineResponse200167Permissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200167Permissions struct {
	value *InlineResponse200167Permissions
	isSet bool
}

func (v NullableInlineResponse200167Permissions) Get() *InlineResponse200167Permissions {
	return v.value
}

func (v *NullableInlineResponse200167Permissions) Set(val *InlineResponse200167Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200167Permissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200167Permissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200167Permissions(val *InlineResponse200167Permissions) *NullableInlineResponse200167Permissions {
	return &NullableInlineResponse200167Permissions{value: val, isSet: true}
}

func (v NullableInlineResponse200167Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200167Permissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


