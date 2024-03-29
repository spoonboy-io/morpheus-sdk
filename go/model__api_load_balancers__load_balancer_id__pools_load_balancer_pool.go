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

// ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool struct for ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool
type ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Balance Algorithm
	VipBalance *string `json:"vipBalance,omitempty"`
	// Min Active Members
	MinActive *int64 `json:"minActive,omitempty"`
	// Configuration object with parameters that vary by type.
	Config *map[string]interface{} `json:"config,omitempty"`
}

// NewApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool instantiates a new ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool() *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
	this := ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool{}
	return &this
}

// NewApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPoolWithDefaults instantiates a new ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPoolWithDefaults() *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
	this := ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) SetDescription(v string) {
	o.Description = &v
}

// GetVipBalance returns the VipBalance field value if set, zero value otherwise.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetVipBalance() string {
	if o == nil || o.VipBalance == nil {
		var ret string
		return ret
	}
	return *o.VipBalance
}

// GetVipBalanceOk returns a tuple with the VipBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetVipBalanceOk() (*string, bool) {
	if o == nil || o.VipBalance == nil {
		return nil, false
	}
	return o.VipBalance, true
}

// HasVipBalance returns a boolean if a field has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) HasVipBalance() bool {
	if o != nil && o.VipBalance != nil {
		return true
	}

	return false
}

// SetVipBalance gets a reference to the given string and assigns it to the VipBalance field.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) SetVipBalance(v string) {
	o.VipBalance = &v
}

// GetMinActive returns the MinActive field value if set, zero value otherwise.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetMinActive() int64 {
	if o == nil || o.MinActive == nil {
		var ret int64
		return ret
	}
	return *o.MinActive
}

// GetMinActiveOk returns a tuple with the MinActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetMinActiveOk() (*int64, bool) {
	if o == nil || o.MinActive == nil {
		return nil, false
	}
	return o.MinActive, true
}

// HasMinActive returns a boolean if a field has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) HasMinActive() bool {
	if o != nil && o.MinActive != nil {
		return true
	}

	return false
}

// SetMinActive gets a reference to the given int64 and assigns it to the MinActive field.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) SetMinActive(v int64) {
	o.MinActive = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

func (o ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.VipBalance != nil {
		toSerialize["vipBalance"] = o.VipBalance
	}
	if o.MinActive != nil {
		toSerialize["minActive"] = o.MinActive
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool struct {
	value *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool
	isSet bool
}

func (v NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) Get() *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
	return v.value
}

func (v *NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) Set(val *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool(val *ApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) *NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool {
	return &NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool{value: val, isSet: true}
}

func (v NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLoadBalancersLoadBalancerIdPoolsLoadBalancerPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


