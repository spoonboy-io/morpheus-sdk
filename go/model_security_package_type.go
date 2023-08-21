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

// SecurityPackageType struct for SecurityPackageType
type SecurityPackageType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	OptionTypes *[]OptionType `json:"optionTypes,omitempty"`
}

// NewSecurityPackageType instantiates a new SecurityPackageType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityPackageType() *SecurityPackageType {
	this := SecurityPackageType{}
	return &this
}

// NewSecurityPackageTypeWithDefaults instantiates a new SecurityPackageType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityPackageTypeWithDefaults() *SecurityPackageType {
	this := SecurityPackageType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityPackageType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPackageType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityPackageType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecurityPackageType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SecurityPackageType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPackageType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SecurityPackageType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SecurityPackageType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityPackageType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPackageType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityPackageType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityPackageType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityPackageType) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityPackageType) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityPackageType) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SecurityPackageType) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SecurityPackageType) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SecurityPackageType) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SecurityPackageType) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPackageType) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityPackageType) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SecurityPackageType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *SecurityPackageType) GetOptionTypes() []OptionType {
	if o == nil || o.OptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityPackageType) GetOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *SecurityPackageType) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []OptionType and assigns it to the OptionTypes field.
func (o *SecurityPackageType) SetOptionTypes(v []OptionType) {
	o.OptionTypes = &v
}

func (o SecurityPackageType) MarshalJSON() ([]byte, error) {
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityPackageType struct {
	value *SecurityPackageType
	isSet bool
}

func (v NullableSecurityPackageType) Get() *SecurityPackageType {
	return v.value
}

func (v *NullableSecurityPackageType) Set(val *SecurityPackageType) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityPackageType) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityPackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityPackageType(val *SecurityPackageType) *NullableSecurityPackageType {
	return &NullableSecurityPackageType{value: val, isSet: true}
}

func (v NullableSecurityPackageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityPackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


