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

// ApiAccountsAccountIdGroupsIdGroup struct for ApiAccountsAccountIdGroupsIdGroup
type ApiAccountsAccountIdGroupsIdGroup struct {
	// A unique name scoped to the subtenant for the group
	Name *string `json:"name,omitempty"`
	// Optional description field if you want to put more info there
	Description *string `json:"description,omitempty"`
	// Optional code for use with policies
	Code *string `json:"code,omitempty"`
	// location
	Location *string `json:"location,omitempty"`
}

// NewApiAccountsAccountIdGroupsIdGroup instantiates a new ApiAccountsAccountIdGroupsIdGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccountsAccountIdGroupsIdGroup() *ApiAccountsAccountIdGroupsIdGroup {
	this := ApiAccountsAccountIdGroupsIdGroup{}
	return &this
}

// NewApiAccountsAccountIdGroupsIdGroupWithDefaults instantiates a new ApiAccountsAccountIdGroupsIdGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccountsAccountIdGroupsIdGroupWithDefaults() *ApiAccountsAccountIdGroupsIdGroup {
	this := ApiAccountsAccountIdGroupsIdGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAccountsAccountIdGroupsIdGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiAccountsAccountIdGroupsIdGroup) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApiAccountsAccountIdGroupsIdGroup) SetCode(v string) {
	o.Code = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApiAccountsAccountIdGroupsIdGroup) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ApiAccountsAccountIdGroupsIdGroup) SetLocation(v string) {
	o.Location = &v
}

func (o ApiAccountsAccountIdGroupsIdGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
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

type NullableApiAccountsAccountIdGroupsIdGroup struct {
	value *ApiAccountsAccountIdGroupsIdGroup
	isSet bool
}

func (v NullableApiAccountsAccountIdGroupsIdGroup) Get() *ApiAccountsAccountIdGroupsIdGroup {
	return v.value
}

func (v *NullableApiAccountsAccountIdGroupsIdGroup) Set(val *ApiAccountsAccountIdGroupsIdGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccountsAccountIdGroupsIdGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccountsAccountIdGroupsIdGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccountsAccountIdGroupsIdGroup(val *ApiAccountsAccountIdGroupsIdGroup) *NullableApiAccountsAccountIdGroupsIdGroup {
	return &NullableApiAccountsAccountIdGroupsIdGroup{value: val, isSet: true}
}

func (v NullableApiAccountsAccountIdGroupsIdGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccountsAccountIdGroupsIdGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

