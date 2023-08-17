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

// checks if the GetAppSecurityGroups200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppSecurityGroups200Response{}

// GetAppSecurityGroups200Response struct for GetAppSecurityGroups200Response
type GetAppSecurityGroups200Response struct {
	Success *bool `json:"success,omitempty"`
	SecurityGroups []AppSecurityGroups `json:"securityGroups,omitempty"`
	FirewallEnabled *bool `json:"firewallEnabled,omitempty"`
}

// NewGetAppSecurityGroups200Response instantiates a new GetAppSecurityGroups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppSecurityGroups200Response() *GetAppSecurityGroups200Response {
	this := GetAppSecurityGroups200Response{}
	return &this
}

// NewGetAppSecurityGroups200ResponseWithDefaults instantiates a new GetAppSecurityGroups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppSecurityGroups200ResponseWithDefaults() *GetAppSecurityGroups200Response {
	this := GetAppSecurityGroups200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetAppSecurityGroups200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppSecurityGroups200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetAppSecurityGroups200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetAppSecurityGroups200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *GetAppSecurityGroups200Response) GetSecurityGroups() []AppSecurityGroups {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret []AppSecurityGroups
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppSecurityGroups200Response) GetSecurityGroupsOk() ([]AppSecurityGroups, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *GetAppSecurityGroups200Response) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []AppSecurityGroups and assigns it to the SecurityGroups field.
func (o *GetAppSecurityGroups200Response) SetSecurityGroups(v []AppSecurityGroups) {
	o.SecurityGroups = v
}

// GetFirewallEnabled returns the FirewallEnabled field value if set, zero value otherwise.
func (o *GetAppSecurityGroups200Response) GetFirewallEnabled() bool {
	if o == nil || IsNil(o.FirewallEnabled) {
		var ret bool
		return ret
	}
	return *o.FirewallEnabled
}

// GetFirewallEnabledOk returns a tuple with the FirewallEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppSecurityGroups200Response) GetFirewallEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FirewallEnabled) {
		return nil, false
	}
	return o.FirewallEnabled, true
}

// HasFirewallEnabled returns a boolean if a field has been set.
func (o *GetAppSecurityGroups200Response) HasFirewallEnabled() bool {
	if o != nil && !IsNil(o.FirewallEnabled) {
		return true
	}

	return false
}

// SetFirewallEnabled gets a reference to the given bool and assigns it to the FirewallEnabled field.
func (o *GetAppSecurityGroups200Response) SetFirewallEnabled(v bool) {
	o.FirewallEnabled = &v
}

func (o GetAppSecurityGroups200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAppSecurityGroups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	if !IsNil(o.FirewallEnabled) {
		toSerialize["firewallEnabled"] = o.FirewallEnabled
	}
	return toSerialize, nil
}

type NullableGetAppSecurityGroups200Response struct {
	value *GetAppSecurityGroups200Response
	isSet bool
}

func (v NullableGetAppSecurityGroups200Response) Get() *GetAppSecurityGroups200Response {
	return v.value
}

func (v *NullableGetAppSecurityGroups200Response) Set(val *GetAppSecurityGroups200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppSecurityGroups200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppSecurityGroups200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppSecurityGroups200Response(val *GetAppSecurityGroups200Response) *NullableGetAppSecurityGroups200Response {
	return &NullableGetAppSecurityGroups200Response{value: val, isSet: true}
}

func (v NullableGetAppSecurityGroups200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppSecurityGroups200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

