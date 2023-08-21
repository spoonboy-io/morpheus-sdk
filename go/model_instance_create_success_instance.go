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

// InstanceCreateSuccessInstance Key for name, site, instanceType layout, and plan.
type InstanceCreateSuccessInstance struct {
	// Name of the instance to be created.
	Name string `json:"name"`
	Site InstanceCreateInstanceSite `json:"site"`
	InstanceType InstanceCreateInstanceInstanceType `json:"instanceType"`
	Layout InstanceCreateInstanceLayout `json:"layout"`
	Plan InstanceCreateInstancePlan `json:"plan"`
}

// NewInstanceCreateSuccessInstance instantiates a new InstanceCreateSuccessInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceCreateSuccessInstance(name string, site InstanceCreateInstanceSite, instanceType InstanceCreateInstanceInstanceType, layout InstanceCreateInstanceLayout, plan InstanceCreateInstancePlan, ) *InstanceCreateSuccessInstance {
	this := InstanceCreateSuccessInstance{}
	this.Name = name
	this.Site = site
	this.InstanceType = instanceType
	this.Layout = layout
	this.Plan = plan
	return &this
}

// NewInstanceCreateSuccessInstanceWithDefaults instantiates a new InstanceCreateSuccessInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceCreateSuccessInstanceWithDefaults() *InstanceCreateSuccessInstance {
	this := InstanceCreateSuccessInstance{}
	return &this
}

// GetName returns the Name field value
func (o *InstanceCreateSuccessInstance) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateSuccessInstance) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InstanceCreateSuccessInstance) SetName(v string) {
	o.Name = v
}

// GetSite returns the Site field value
func (o *InstanceCreateSuccessInstance) GetSite() InstanceCreateInstanceSite {
	if o == nil  {
		var ret InstanceCreateInstanceSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateSuccessInstance) GetSiteOk() (*InstanceCreateInstanceSite, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *InstanceCreateSuccessInstance) SetSite(v InstanceCreateInstanceSite) {
	o.Site = v
}

// GetInstanceType returns the InstanceType field value
func (o *InstanceCreateSuccessInstance) GetInstanceType() InstanceCreateInstanceInstanceType {
	if o == nil  {
		var ret InstanceCreateInstanceInstanceType
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateSuccessInstance) GetInstanceTypeOk() (*InstanceCreateInstanceInstanceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *InstanceCreateSuccessInstance) SetInstanceType(v InstanceCreateInstanceInstanceType) {
	o.InstanceType = v
}

// GetLayout returns the Layout field value
func (o *InstanceCreateSuccessInstance) GetLayout() InstanceCreateInstanceLayout {
	if o == nil  {
		var ret InstanceCreateInstanceLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateSuccessInstance) GetLayoutOk() (*InstanceCreateInstanceLayout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *InstanceCreateSuccessInstance) SetLayout(v InstanceCreateInstanceLayout) {
	o.Layout = v
}

// GetPlan returns the Plan field value
func (o *InstanceCreateSuccessInstance) GetPlan() InstanceCreateInstancePlan {
	if o == nil  {
		var ret InstanceCreateInstancePlan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *InstanceCreateSuccessInstance) GetPlanOk() (*InstanceCreateInstancePlan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *InstanceCreateSuccessInstance) SetPlan(v InstanceCreateInstancePlan) {
	o.Plan = v
}

func (o InstanceCreateSuccessInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["site"] = o.Site
	}
	if true {
		toSerialize["instanceType"] = o.InstanceType
	}
	if true {
		toSerialize["layout"] = o.Layout
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceCreateSuccessInstance struct {
	value *InstanceCreateSuccessInstance
	isSet bool
}

func (v NullableInstanceCreateSuccessInstance) Get() *InstanceCreateSuccessInstance {
	return v.value
}

func (v *NullableInstanceCreateSuccessInstance) Set(val *InstanceCreateSuccessInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceCreateSuccessInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceCreateSuccessInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceCreateSuccessInstance(val *InstanceCreateSuccessInstance) *NullableInstanceCreateSuccessInstance {
	return &NullableInstanceCreateSuccessInstance{value: val, isSet: true}
}

func (v NullableInstanceCreateSuccessInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceCreateSuccessInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

