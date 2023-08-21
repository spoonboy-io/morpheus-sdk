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

// InstanceTypeLayoutPermissionsResourcePermissions struct for InstanceTypeLayoutPermissionsResourcePermissions
type InstanceTypeLayoutPermissionsResourcePermissions struct {
	DefaultStore *bool `json:"defaultStore,omitempty"`
	AllPlans *bool `json:"allPlans,omitempty"`
	DefaultTarget *bool `json:"defaultTarget,omitempty"`
	CanManage *bool `json:"canManage,omitempty"`
	All *bool `json:"all,omitempty"`
	Account NullableCheckContainer `json:"account,omitempty"`
	Sites []map[string]interface{} `json:"sites,omitempty"`
	Plans []map[string]interface{} `json:"plans,omitempty"`
}

// NewInstanceTypeLayoutPermissionsResourcePermissions instantiates a new InstanceTypeLayoutPermissionsResourcePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeLayoutPermissionsResourcePermissions() *InstanceTypeLayoutPermissionsResourcePermissions {
	this := InstanceTypeLayoutPermissionsResourcePermissions{}
	return &this
}

// NewInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults instantiates a new InstanceTypeLayoutPermissionsResourcePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults() *InstanceTypeLayoutPermissionsResourcePermissions {
	this := InstanceTypeLayoutPermissionsResourcePermissions{}
	return &this
}

// GetDefaultStore returns the DefaultStore field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultStore() bool {
	if o == nil || o.DefaultStore == nil {
		var ret bool
		return ret
	}
	return *o.DefaultStore
}

// GetDefaultStoreOk returns a tuple with the DefaultStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool) {
	if o == nil || o.DefaultStore == nil {
		return nil, false
	}
	return o.DefaultStore, true
}

// HasDefaultStore returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasDefaultStore() bool {
	if o != nil && o.DefaultStore != nil {
		return true
	}

	return false
}

// SetDefaultStore gets a reference to the given bool and assigns it to the DefaultStore field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetDefaultStore(v bool) {
	o.DefaultStore = &v
}

// GetAllPlans returns the AllPlans field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllPlans() bool {
	if o == nil || o.AllPlans == nil {
		var ret bool
		return ret
	}
	return *o.AllPlans
}

// GetAllPlansOk returns a tuple with the AllPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool) {
	if o == nil || o.AllPlans == nil {
		return nil, false
	}
	return o.AllPlans, true
}

// HasAllPlans returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAllPlans() bool {
	if o != nil && o.AllPlans != nil {
		return true
	}

	return false
}

// SetAllPlans gets a reference to the given bool and assigns it to the AllPlans field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAllPlans(v bool) {
	o.AllPlans = &v
}

// GetDefaultTarget returns the DefaultTarget field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultTarget() bool {
	if o == nil || o.DefaultTarget == nil {
		var ret bool
		return ret
	}
	return *o.DefaultTarget
}

// GetDefaultTargetOk returns a tuple with the DefaultTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool) {
	if o == nil || o.DefaultTarget == nil {
		return nil, false
	}
	return o.DefaultTarget, true
}

// HasDefaultTarget returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasDefaultTarget() bool {
	if o != nil && o.DefaultTarget != nil {
		return true
	}

	return false
}

// SetDefaultTarget gets a reference to the given bool and assigns it to the DefaultTarget field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetDefaultTarget(v bool) {
	o.DefaultTarget = &v
}

// GetCanManage returns the CanManage field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetCanManage() bool {
	if o == nil || o.CanManage == nil {
		var ret bool
		return ret
	}
	return *o.CanManage
}

// GetCanManageOk returns a tuple with the CanManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetCanManageOk() (*bool, bool) {
	if o == nil || o.CanManage == nil {
		return nil, false
	}
	return o.CanManage, true
}

// HasCanManage returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasCanManage() bool {
	if o != nil && o.CanManage != nil {
		return true
	}

	return false
}

// SetCanManage gets a reference to the given bool and assigns it to the CanManage field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetCanManage(v bool) {
	o.CanManage = &v
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAll() bool {
	if o == nil || o.All == nil {
		var ret bool
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAllOk() (*bool, bool) {
	if o == nil || o.All == nil {
		return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAll() bool {
	if o != nil && o.All != nil {
		return true
	}

	return false
}

// SetAll gets a reference to the given bool and assigns it to the All field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAll(v bool) {
	o.All = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAccount() CheckContainer {
	if o == nil || o.Account.Get() == nil {
		var ret CheckContainer
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetAccountOk() (*CheckContainer, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableCheckContainer and assigns it to the Account field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAccount(v CheckContainer) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *InstanceTypeLayoutPermissionsResourcePermissions) UnsetAccount() {
	o.Account.Unset()
}

// GetSites returns the Sites field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetSites() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.Sites
}

// GetSitesOk returns a tuple with the Sites field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetSitesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Sites == nil {
		return nil, false
	}
	return &o.Sites, true
}

// HasSites returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasSites() bool {
	if o != nil && o.Sites != nil {
		return true
	}

	return false
}

// SetSites gets a reference to the given []map[string]interface{} and assigns it to the Sites field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetSites(v []map[string]interface{}) {
	o.Sites = v
}

// GetPlans returns the Plans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetPlans() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeLayoutPermissionsResourcePermissions) GetPlansOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return &o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []map[string]interface{} and assigns it to the Plans field.
func (o *InstanceTypeLayoutPermissionsResourcePermissions) SetPlans(v []map[string]interface{}) {
	o.Plans = v
}

func (o InstanceTypeLayoutPermissionsResourcePermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultStore != nil {
		toSerialize["defaultStore"] = o.DefaultStore
	}
	if o.AllPlans != nil {
		toSerialize["allPlans"] = o.AllPlans
	}
	if o.DefaultTarget != nil {
		toSerialize["defaultTarget"] = o.DefaultTarget
	}
	if o.CanManage != nil {
		toSerialize["canManage"] = o.CanManage
	}
	if o.All != nil {
		toSerialize["all"] = o.All
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Sites != nil {
		toSerialize["sites"] = o.Sites
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTypeLayoutPermissionsResourcePermissions struct {
	value *InstanceTypeLayoutPermissionsResourcePermissions
	isSet bool
}

func (v NullableInstanceTypeLayoutPermissionsResourcePermissions) Get() *InstanceTypeLayoutPermissionsResourcePermissions {
	return v.value
}

func (v *NullableInstanceTypeLayoutPermissionsResourcePermissions) Set(val *InstanceTypeLayoutPermissionsResourcePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeLayoutPermissionsResourcePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeLayoutPermissionsResourcePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeLayoutPermissionsResourcePermissions(val *InstanceTypeLayoutPermissionsResourcePermissions) *NullableInstanceTypeLayoutPermissionsResourcePermissions {
	return &NullableInstanceTypeLayoutPermissionsResourcePermissions{value: val, isSet: true}
}

func (v NullableInstanceTypeLayoutPermissionsResourcePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeLayoutPermissionsResourcePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


