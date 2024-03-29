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

// SubnetType struct for SubnetType
type SubnetType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Creatable *bool `json:"creatable,omitempty"`
	Deletable *bool `json:"deletable,omitempty"`
	DhcpServerEditable *bool `json:"dhcpServerEditable,omitempty"`
	CanAssignPool *bool `json:"canAssignPool,omitempty"`
	OptionTypes *[]OptionType `json:"optionTypes,omitempty"`
}

// NewSubnetType instantiates a new SubnetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetType() *SubnetType {
	this := SubnetType{}
	return &this
}

// NewSubnetTypeWithDefaults instantiates a new SubnetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetTypeWithDefaults() *SubnetType {
	this := SubnetType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubnetType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubnetType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SubnetType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SubnetType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SubnetType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SubnetType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubnetType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubnetType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubnetType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubnetType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubnetType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubnetType) SetDescription(v string) {
	o.Description = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *SubnetType) GetCreatable() bool {
	if o == nil || o.Creatable == nil {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetCreatableOk() (*bool, bool) {
	if o == nil || o.Creatable == nil {
		return nil, false
	}
	return o.Creatable, true
}

// HasCreatable returns a boolean if a field has been set.
func (o *SubnetType) HasCreatable() bool {
	if o != nil && o.Creatable != nil {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *SubnetType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *SubnetType) GetDeletable() bool {
	if o == nil || o.Deletable == nil {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetDeletableOk() (*bool, bool) {
	if o == nil || o.Deletable == nil {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *SubnetType) HasDeletable() bool {
	if o != nil && o.Deletable != nil {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *SubnetType) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDhcpServerEditable returns the DhcpServerEditable field value if set, zero value otherwise.
func (o *SubnetType) GetDhcpServerEditable() bool {
	if o == nil || o.DhcpServerEditable == nil {
		var ret bool
		return ret
	}
	return *o.DhcpServerEditable
}

// GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetDhcpServerEditableOk() (*bool, bool) {
	if o == nil || o.DhcpServerEditable == nil {
		return nil, false
	}
	return o.DhcpServerEditable, true
}

// HasDhcpServerEditable returns a boolean if a field has been set.
func (o *SubnetType) HasDhcpServerEditable() bool {
	if o != nil && o.DhcpServerEditable != nil {
		return true
	}

	return false
}

// SetDhcpServerEditable gets a reference to the given bool and assigns it to the DhcpServerEditable field.
func (o *SubnetType) SetDhcpServerEditable(v bool) {
	o.DhcpServerEditable = &v
}

// GetCanAssignPool returns the CanAssignPool field value if set, zero value otherwise.
func (o *SubnetType) GetCanAssignPool() bool {
	if o == nil || o.CanAssignPool == nil {
		var ret bool
		return ret
	}
	return *o.CanAssignPool
}

// GetCanAssignPoolOk returns a tuple with the CanAssignPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetCanAssignPoolOk() (*bool, bool) {
	if o == nil || o.CanAssignPool == nil {
		return nil, false
	}
	return o.CanAssignPool, true
}

// HasCanAssignPool returns a boolean if a field has been set.
func (o *SubnetType) HasCanAssignPool() bool {
	if o != nil && o.CanAssignPool != nil {
		return true
	}

	return false
}

// SetCanAssignPool gets a reference to the given bool and assigns it to the CanAssignPool field.
func (o *SubnetType) SetCanAssignPool(v bool) {
	o.CanAssignPool = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *SubnetType) GetOptionTypes() []OptionType {
	if o == nil || o.OptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubnetType) GetOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *SubnetType) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []OptionType and assigns it to the OptionTypes field.
func (o *SubnetType) SetOptionTypes(v []OptionType) {
	o.OptionTypes = &v
}

func (o SubnetType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Creatable != nil {
		toSerialize["creatable"] = o.Creatable
	}
	if o.Deletable != nil {
		toSerialize["deletable"] = o.Deletable
	}
	if o.DhcpServerEditable != nil {
		toSerialize["dhcpServerEditable"] = o.DhcpServerEditable
	}
	if o.CanAssignPool != nil {
		toSerialize["canAssignPool"] = o.CanAssignPool
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableSubnetType struct {
	value *SubnetType
	isSet bool
}

func (v NullableSubnetType) Get() *SubnetType {
	return v.value
}

func (v *NullableSubnetType) Set(val *SubnetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetType(val *SubnetType) *NullableSubnetType {
	return &NullableSubnetType{value: val, isSet: true}
}

func (v NullableSubnetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


