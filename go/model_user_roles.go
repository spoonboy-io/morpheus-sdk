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

// UserRoles struct for UserRoles
type UserRoles struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Authority *string `json:"authority,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewUserRoles instantiates a new UserRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRoles() *UserRoles {
	this := UserRoles{}
	return &this
}

// NewUserRolesWithDefaults instantiates a new UserRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRolesWithDefaults() *UserRoles {
	this := UserRoles{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserRoles) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserRoles) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UserRoles) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserRoles) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserRoles) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserRoles) SetName(v string) {
	o.Name = &v
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *UserRoles) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetAuthorityOk() (*string, bool) {
	if o == nil || o.Authority == nil {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *UserRoles) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *UserRoles) SetAuthority(v string) {
	o.Authority = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserRoles) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRoles) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserRoles) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserRoles) SetDescription(v string) {
	o.Description = &v
}

func (o UserRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Authority != nil {
		toSerialize["authority"] = o.Authority
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUserRoles struct {
	value *UserRoles
	isSet bool
}

func (v NullableUserRoles) Get() *UserRoles {
	return v.value
}

func (v *NullableUserRoles) Set(val *UserRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRoles(val *UserRoles) *NullableUserRoles {
	return &NullableUserRoles{value: val, isSet: true}
}

func (v NullableUserRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


