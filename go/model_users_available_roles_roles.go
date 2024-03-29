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

// UsersAvailableRolesRoles struct for UsersAvailableRolesRoles
type UsersAvailableRolesRoles struct {
	Id *int64 `json:"id,omitempty"`
	Authority *string `json:"authority,omitempty"`
	Description NullableString `json:"description,omitempty"`
	RoleType *string `json:"roleType,omitempty"`
	Owner *InlineResponse20040AppDeployInstance `json:"owner,omitempty"`
}

// NewUsersAvailableRolesRoles instantiates a new UsersAvailableRolesRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersAvailableRolesRoles() *UsersAvailableRolesRoles {
	this := UsersAvailableRolesRoles{}
	return &this
}

// NewUsersAvailableRolesRolesWithDefaults instantiates a new UsersAvailableRolesRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersAvailableRolesRolesWithDefaults() *UsersAvailableRolesRoles {
	this := UsersAvailableRolesRoles{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsersAvailableRolesRoles) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersAvailableRolesRoles) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsersAvailableRolesRoles) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UsersAvailableRolesRoles) SetId(v int64) {
	o.Id = &v
}

// GetAuthority returns the Authority field value if set, zero value otherwise.
func (o *UsersAvailableRolesRoles) GetAuthority() string {
	if o == nil || o.Authority == nil {
		var ret string
		return ret
	}
	return *o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersAvailableRolesRoles) GetAuthorityOk() (*string, bool) {
	if o == nil || o.Authority == nil {
		return nil, false
	}
	return o.Authority, true
}

// HasAuthority returns a boolean if a field has been set.
func (o *UsersAvailableRolesRoles) HasAuthority() bool {
	if o != nil && o.Authority != nil {
		return true
	}

	return false
}

// SetAuthority gets a reference to the given string and assigns it to the Authority field.
func (o *UsersAvailableRolesRoles) SetAuthority(v string) {
	o.Authority = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersAvailableRolesRoles) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersAvailableRolesRoles) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UsersAvailableRolesRoles) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UsersAvailableRolesRoles) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UsersAvailableRolesRoles) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UsersAvailableRolesRoles) UnsetDescription() {
	o.Description.Unset()
}

// GetRoleType returns the RoleType field value if set, zero value otherwise.
func (o *UsersAvailableRolesRoles) GetRoleType() string {
	if o == nil || o.RoleType == nil {
		var ret string
		return ret
	}
	return *o.RoleType
}

// GetRoleTypeOk returns a tuple with the RoleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersAvailableRolesRoles) GetRoleTypeOk() (*string, bool) {
	if o == nil || o.RoleType == nil {
		return nil, false
	}
	return o.RoleType, true
}

// HasRoleType returns a boolean if a field has been set.
func (o *UsersAvailableRolesRoles) HasRoleType() bool {
	if o != nil && o.RoleType != nil {
		return true
	}

	return false
}

// SetRoleType gets a reference to the given string and assigns it to the RoleType field.
func (o *UsersAvailableRolesRoles) SetRoleType(v string) {
	o.RoleType = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UsersAvailableRolesRoles) GetOwner() InlineResponse20040AppDeployInstance {
	if o == nil || o.Owner == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersAvailableRolesRoles) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UsersAvailableRolesRoles) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Owner field.
func (o *UsersAvailableRolesRoles) SetOwner(v InlineResponse20040AppDeployInstance) {
	o.Owner = &v
}

func (o UsersAvailableRolesRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Authority != nil {
		toSerialize["authority"] = o.Authority
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RoleType != nil {
		toSerialize["roleType"] = o.RoleType
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableUsersAvailableRolesRoles struct {
	value *UsersAvailableRolesRoles
	isSet bool
}

func (v NullableUsersAvailableRolesRoles) Get() *UsersAvailableRolesRoles {
	return v.value
}

func (v *NullableUsersAvailableRolesRoles) Set(val *UsersAvailableRolesRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersAvailableRolesRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersAvailableRolesRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersAvailableRolesRoles(val *UsersAvailableRolesRoles) *NullableUsersAvailableRolesRoles {
	return &NullableUsersAvailableRolesRoles{value: val, isSet: true}
}

func (v NullableUsersAvailableRolesRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersAvailableRolesRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


