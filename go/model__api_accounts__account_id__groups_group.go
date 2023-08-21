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

// ApiAccountsAccountIdGroupsGroup Payload for updating an existing tenant
type ApiAccountsAccountIdGroupsGroup struct {
	// A unique name scoped to the subtenant for the group
	Name string `json:"name"`
	// Optional description field if you want to put more info there
	Description *string `json:"description,omitempty"`
	// Optional code for use with policies
	Code *string `json:"code,omitempty"`
	// location
	Location *string `json:"location,omitempty"`
}

// NewApiAccountsAccountIdGroupsGroup instantiates a new ApiAccountsAccountIdGroupsGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccountsAccountIdGroupsGroup(name string, ) *ApiAccountsAccountIdGroupsGroup {
	this := ApiAccountsAccountIdGroupsGroup{}
	this.Name = name
	return &this
}

// NewApiAccountsAccountIdGroupsGroupWithDefaults instantiates a new ApiAccountsAccountIdGroupsGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccountsAccountIdGroupsGroupWithDefaults() *ApiAccountsAccountIdGroupsGroup {
	this := ApiAccountsAccountIdGroupsGroup{}
	return &this
}

// GetName returns the Name field value
func (o *ApiAccountsAccountIdGroupsGroup) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsGroup) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiAccountsAccountIdGroupsGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAccountsAccountIdGroupsGroup) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsGroup) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsGroup) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsGroup) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiAccountsAccountIdGroupsGroup) SetCode(v string) {
	o.Code = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsGroup) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsGroup) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsGroup) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ApiAccountsAccountIdGroupsGroup) SetLocation(v string) {
	o.Location = &v
}

func (o ApiAccountsAccountIdGroupsGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableApiAccountsAccountIdGroupsGroup struct {
	value *ApiAccountsAccountIdGroupsGroup
	isSet bool
}

func (v NullableApiAccountsAccountIdGroupsGroup) Get() *ApiAccountsAccountIdGroupsGroup {
	return v.value
}

func (v *NullableApiAccountsAccountIdGroupsGroup) Set(val *ApiAccountsAccountIdGroupsGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccountsAccountIdGroupsGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccountsAccountIdGroupsGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccountsAccountIdGroupsGroup(val *ApiAccountsAccountIdGroupsGroup) *NullableApiAccountsAccountIdGroupsGroup {
	return &NullableApiAccountsAccountIdGroupsGroup{value: val, isSet: true}
}

func (v NullableApiAccountsAccountIdGroupsGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccountsAccountIdGroupsGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

