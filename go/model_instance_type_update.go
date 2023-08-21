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

// InstanceTypeUpdate struct for InstanceTypeUpdate
type InstanceTypeUpdate struct {
	// Instance type name
	Name *string `json:"name,omitempty"`
	// Instance type description
	Description *string `json:"description,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Instance type code
	Code *string `json:"code,omitempty"`
	// Category
	Category *string `json:"category,omitempty"`
	// Visibility
	Visibility *string `json:"visibility,omitempty"`
	// Featured
	Featured *bool `json:"featured,omitempty"`
	// Enable Settings
	HasSettings *bool `json:"hasSettings,omitempty"`
	// Enable Scaling (Horizontal)
	HasAutoScale *bool `json:"hasAutoScale,omitempty"`
	// Supports Deployments
	HasDeployment *bool `json:"hasDeployment,omitempty"`
	// Environment Prefix, can be used to make exported evars unique.
	EnvironmentPrefix *string `json:"environmentPrefix,omitempty"`
	// Array of instance type env variables.
	EnvironmentVariables *[]ClusterLayoutCreateEnvironmentVariables `json:"environmentVariables,omitempty"`
	// Array of price set objects
	PriceSets *[]InstanceTypeCreatePriceSets `json:"priceSets,omitempty"`
	// Array of instance type option type IDs
	OptionTypes *[]int64 `json:"optionTypes,omitempty"`
}

// NewInstanceTypeUpdate instantiates a new InstanceTypeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeUpdate() *InstanceTypeUpdate {
	this := InstanceTypeUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewInstanceTypeUpdateWithDefaults instantiates a new InstanceTypeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeUpdateWithDefaults() *InstanceTypeUpdate {
	this := InstanceTypeUpdate{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTypeUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceTypeUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeUpdate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeUpdate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *InstanceTypeUpdate) SetLabels(v []string) {
	o.Labels = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InstanceTypeUpdate) SetCode(v string) {
	o.Code = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InstanceTypeUpdate) SetCategory(v string) {
	o.Category = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *InstanceTypeUpdate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetFeatured() bool {
	if o == nil || o.Featured == nil {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetFeaturedOk() (*bool, bool) {
	if o == nil || o.Featured == nil {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasFeatured() bool {
	if o != nil && o.Featured != nil {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *InstanceTypeUpdate) SetFeatured(v bool) {
	o.Featured = &v
}

// GetHasSettings returns the HasSettings field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasSettings() bool {
	if o == nil || o.HasSettings == nil {
		var ret bool
		return ret
	}
	return *o.HasSettings
}

// GetHasSettingsOk returns a tuple with the HasSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasSettingsOk() (*bool, bool) {
	if o == nil || o.HasSettings == nil {
		return nil, false
	}
	return o.HasSettings, true
}

// HasHasSettings returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasHasSettings() bool {
	if o != nil && o.HasSettings != nil {
		return true
	}

	return false
}

// SetHasSettings gets a reference to the given bool and assigns it to the HasSettings field.
func (o *InstanceTypeUpdate) SetHasSettings(v bool) {
	o.HasSettings = &v
}

// GetHasAutoScale returns the HasAutoScale field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasAutoScale() bool {
	if o == nil || o.HasAutoScale == nil {
		var ret bool
		return ret
	}
	return *o.HasAutoScale
}

// GetHasAutoScaleOk returns a tuple with the HasAutoScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasAutoScaleOk() (*bool, bool) {
	if o == nil || o.HasAutoScale == nil {
		return nil, false
	}
	return o.HasAutoScale, true
}

// HasHasAutoScale returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasHasAutoScale() bool {
	if o != nil && o.HasAutoScale != nil {
		return true
	}

	return false
}

// SetHasAutoScale gets a reference to the given bool and assigns it to the HasAutoScale field.
func (o *InstanceTypeUpdate) SetHasAutoScale(v bool) {
	o.HasAutoScale = &v
}

// GetHasDeployment returns the HasDeployment field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetHasDeployment() bool {
	if o == nil || o.HasDeployment == nil {
		var ret bool
		return ret
	}
	return *o.HasDeployment
}

// GetHasDeploymentOk returns a tuple with the HasDeployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetHasDeploymentOk() (*bool, bool) {
	if o == nil || o.HasDeployment == nil {
		return nil, false
	}
	return o.HasDeployment, true
}

// HasHasDeployment returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasHasDeployment() bool {
	if o != nil && o.HasDeployment != nil {
		return true
	}

	return false
}

// SetHasDeployment gets a reference to the given bool and assigns it to the HasDeployment field.
func (o *InstanceTypeUpdate) SetHasDeployment(v bool) {
	o.HasDeployment = &v
}

// GetEnvironmentPrefix returns the EnvironmentPrefix field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetEnvironmentPrefix() string {
	if o == nil || o.EnvironmentPrefix == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentPrefix
}

// GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetEnvironmentPrefixOk() (*string, bool) {
	if o == nil || o.EnvironmentPrefix == nil {
		return nil, false
	}
	return o.EnvironmentPrefix, true
}

// HasEnvironmentPrefix returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasEnvironmentPrefix() bool {
	if o != nil && o.EnvironmentPrefix != nil {
		return true
	}

	return false
}

// SetEnvironmentPrefix gets a reference to the given string and assigns it to the EnvironmentPrefix field.
func (o *InstanceTypeUpdate) SetEnvironmentPrefix(v string) {
	o.EnvironmentPrefix = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetEnvironmentVariables() []ClusterLayoutCreateEnvironmentVariables {
	if o == nil || o.EnvironmentVariables == nil {
		var ret []ClusterLayoutCreateEnvironmentVariables
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetEnvironmentVariablesOk() (*[]ClusterLayoutCreateEnvironmentVariables, bool) {
	if o == nil || o.EnvironmentVariables == nil {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasEnvironmentVariables() bool {
	if o != nil && o.EnvironmentVariables != nil {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []ClusterLayoutCreateEnvironmentVariables and assigns it to the EnvironmentVariables field.
func (o *InstanceTypeUpdate) SetEnvironmentVariables(v []ClusterLayoutCreateEnvironmentVariables) {
	o.EnvironmentVariables = &v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetPriceSets() []InstanceTypeCreatePriceSets {
	if o == nil || o.PriceSets == nil {
		var ret []InstanceTypeCreatePriceSets
		return ret
	}
	return *o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetPriceSetsOk() (*[]InstanceTypeCreatePriceSets, bool) {
	if o == nil || o.PriceSets == nil {
		return nil, false
	}
	return o.PriceSets, true
}

// HasPriceSets returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasPriceSets() bool {
	if o != nil && o.PriceSets != nil {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []InstanceTypeCreatePriceSets and assigns it to the PriceSets field.
func (o *InstanceTypeUpdate) SetPriceSets(v []InstanceTypeCreatePriceSets) {
	o.PriceSets = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *InstanceTypeUpdate) GetOptionTypes() []int64 {
	if o == nil || o.OptionTypes == nil {
		var ret []int64
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeUpdate) GetOptionTypesOk() (*[]int64, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *InstanceTypeUpdate) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []int64 and assigns it to the OptionTypes field.
func (o *InstanceTypeUpdate) SetOptionTypes(v []int64) {
	o.OptionTypes = &v
}

func (o InstanceTypeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Featured != nil {
		toSerialize["featured"] = o.Featured
	}
	if o.HasSettings != nil {
		toSerialize["hasSettings"] = o.HasSettings
	}
	if o.HasAutoScale != nil {
		toSerialize["hasAutoScale"] = o.HasAutoScale
	}
	if o.HasDeployment != nil {
		toSerialize["hasDeployment"] = o.HasDeployment
	}
	if o.EnvironmentPrefix != nil {
		toSerialize["environmentPrefix"] = o.EnvironmentPrefix
	}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if o.PriceSets != nil {
		toSerialize["priceSets"] = o.PriceSets
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTypeUpdate struct {
	value *InstanceTypeUpdate
	isSet bool
}

func (v NullableInstanceTypeUpdate) Get() *InstanceTypeUpdate {
	return v.value
}

func (v *NullableInstanceTypeUpdate) Set(val *InstanceTypeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeUpdate(val *InstanceTypeUpdate) *NullableInstanceTypeUpdate {
	return &NullableInstanceTypeUpdate{value: val, isSet: true}
}

func (v NullableInstanceTypeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


