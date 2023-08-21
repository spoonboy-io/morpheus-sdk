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

// LoadBalancerUpdate struct for LoadBalancerUpdate
type LoadBalancerUpdate struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Activate (true) or disable (false)
	Enabled *bool `json:"enabled,omitempty"`
	// Configuration object with parameters that vary by load balancer type.
	Config *map[string]interface{} `json:"config,omitempty"`
	// private or public
	Visibility *string `json:"visibility,omitempty"`
	// Array of tenant account ids that are allowed access
	Tenants *[]LoadBalancerCreateTenants `json:"tenants,omitempty"`
	ResourcePermission *LoadBalancerCreateResourcePermission `json:"resourcePermission,omitempty"`
}

// NewLoadBalancerUpdate instantiates a new LoadBalancerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerUpdate() *LoadBalancerUpdate {
	this := LoadBalancerUpdate{}
	var visibility string = "public"
	this.Visibility = &visibility
	return &this
}

// NewLoadBalancerUpdateWithDefaults instantiates a new LoadBalancerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerUpdateWithDefaults() *LoadBalancerUpdate {
	this := LoadBalancerUpdate{}
	var visibility string = "public"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LoadBalancerUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *LoadBalancerUpdate) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *LoadBalancerUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetTenants() []LoadBalancerCreateTenants {
	if o == nil || o.Tenants == nil {
		var ret []LoadBalancerCreateTenants
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetTenantsOk() (*[]LoadBalancerCreateTenants, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []LoadBalancerCreateTenants and assigns it to the Tenants field.
func (o *LoadBalancerUpdate) SetTenants(v []LoadBalancerCreateTenants) {
	o.Tenants = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetResourcePermission() LoadBalancerCreateResourcePermission {
	if o == nil || o.ResourcePermission == nil {
		var ret LoadBalancerCreateResourcePermission
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetResourcePermissionOk() (*LoadBalancerCreateResourcePermission, bool) {
	if o == nil || o.ResourcePermission == nil {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasResourcePermission() bool {
	if o != nil && o.ResourcePermission != nil {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given LoadBalancerCreateResourcePermission and assigns it to the ResourcePermission field.
func (o *LoadBalancerUpdate) SetResourcePermission(v LoadBalancerCreateResourcePermission) {
	o.ResourcePermission = &v
}

func (o LoadBalancerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	if o.ResourcePermission != nil {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	return json.Marshal(toSerialize)
}

type NullableLoadBalancerUpdate struct {
	value *LoadBalancerUpdate
	isSet bool
}

func (v NullableLoadBalancerUpdate) Get() *LoadBalancerUpdate {
	return v.value
}

func (v *NullableLoadBalancerUpdate) Set(val *LoadBalancerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerUpdate(val *LoadBalancerUpdate) *NullableLoadBalancerUpdate {
	return &NullableLoadBalancerUpdate{value: val, isSet: true}
}

func (v NullableLoadBalancerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

