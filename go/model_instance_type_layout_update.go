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

// InstanceTypeLayoutUpdate struct for InstanceTypeLayoutUpdate
type InstanceTypeLayoutUpdate struct {
	// Layout name
	Name *string `json:"name,omitempty"`
	Labels []string `json:"labels,omitempty"`
	// Version of the layout
	InstanceVersion *string `json:"instanceVersion,omitempty"`
	// Layout description
	Description *string `json:"description,omitempty"`
	// Can be used to enable / disable the creatability of the layout.
	Creatable *bool `json:"creatable,omitempty"`
	// Provision type code
	ProvisionTypeCode *string `json:"provisionTypeCode,omitempty"`
	// Memory requirement in megabytes
	MemoryRequirement *string `json:"memoryRequirement,omitempty"`
	// Can be used to enable / disable the horizontal scaling.
	HasAutoScale *bool `json:"hasAutoScale,omitempty"`
	// Can be used to enable / disable the supports convert to managed.
	SupportsConvertToManaged *bool `json:"supportsConvertToManaged,omitempty"`
	// Array of layout node type IDs
	ContainerTypes *[]int64 `json:"containerTypes,omitempty"`
	// Array of layout option type IDs
	OptionTypes *[]int64 `json:"optionTypes,omitempty"`
	// Array of layout spec template IDs
	SpecTemplates *[]int64 `json:"specTemplates,omitempty"`
	// The environmentVariables parameter is array of env objects
	EnvironmentVariables *[]ClusterLayoutCreateEnvironmentVariables `json:"environmentVariables,omitempty"`
	// Array of price set objects
	PriceSets *[]InstanceTypeCreatePriceSets `json:"priceSets,omitempty"`
	Permissions *ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions `json:"permissions,omitempty"`
}

// NewInstanceTypeLayoutUpdate instantiates a new InstanceTypeLayoutUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceTypeLayoutUpdate() *InstanceTypeLayoutUpdate {
	this := InstanceTypeLayoutUpdate{}
	var creatable bool = true
	this.Creatable = &creatable
	var hasAutoScale bool = false
	this.HasAutoScale = &hasAutoScale
	var supportsConvertToManaged bool = false
	this.SupportsConvertToManaged = &supportsConvertToManaged
	return &this
}

// NewInstanceTypeLayoutUpdateWithDefaults instantiates a new InstanceTypeLayoutUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceTypeLayoutUpdateWithDefaults() *InstanceTypeLayoutUpdate {
	this := InstanceTypeLayoutUpdate{}
	var creatable bool = true
	this.Creatable = &creatable
	var hasAutoScale bool = false
	this.HasAutoScale = &hasAutoScale
	var supportsConvertToManaged bool = false
	this.SupportsConvertToManaged = &supportsConvertToManaged
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceTypeLayoutUpdate) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceTypeLayoutUpdate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceTypeLayoutUpdate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *InstanceTypeLayoutUpdate) SetLabels(v []string) {
	o.Labels = v
}

// GetInstanceVersion returns the InstanceVersion field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetInstanceVersion() string {
	if o == nil || o.InstanceVersion == nil {
		var ret string
		return ret
	}
	return *o.InstanceVersion
}

// GetInstanceVersionOk returns a tuple with the InstanceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetInstanceVersionOk() (*string, bool) {
	if o == nil || o.InstanceVersion == nil {
		return nil, false
	}
	return o.InstanceVersion, true
}

// HasInstanceVersion returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasInstanceVersion() bool {
	if o != nil && o.InstanceVersion != nil {
		return true
	}

	return false
}

// SetInstanceVersion gets a reference to the given string and assigns it to the InstanceVersion field.
func (o *InstanceTypeLayoutUpdate) SetInstanceVersion(v string) {
	o.InstanceVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstanceTypeLayoutUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetCreatable() bool {
	if o == nil || o.Creatable == nil {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetCreatableOk() (*bool, bool) {
	if o == nil || o.Creatable == nil {
		return nil, false
	}
	return o.Creatable, true
}

// HasCreatable returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasCreatable() bool {
	if o != nil && o.Creatable != nil {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *InstanceTypeLayoutUpdate) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetProvisionTypeCode returns the ProvisionTypeCode field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetProvisionTypeCode() string {
	if o == nil || o.ProvisionTypeCode == nil {
		var ret string
		return ret
	}
	return *o.ProvisionTypeCode
}

// GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetProvisionTypeCodeOk() (*string, bool) {
	if o == nil || o.ProvisionTypeCode == nil {
		return nil, false
	}
	return o.ProvisionTypeCode, true
}

// HasProvisionTypeCode returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasProvisionTypeCode() bool {
	if o != nil && o.ProvisionTypeCode != nil {
		return true
	}

	return false
}

// SetProvisionTypeCode gets a reference to the given string and assigns it to the ProvisionTypeCode field.
func (o *InstanceTypeLayoutUpdate) SetProvisionTypeCode(v string) {
	o.ProvisionTypeCode = &v
}

// GetMemoryRequirement returns the MemoryRequirement field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetMemoryRequirement() string {
	if o == nil || o.MemoryRequirement == nil {
		var ret string
		return ret
	}
	return *o.MemoryRequirement
}

// GetMemoryRequirementOk returns a tuple with the MemoryRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetMemoryRequirementOk() (*string, bool) {
	if o == nil || o.MemoryRequirement == nil {
		return nil, false
	}
	return o.MemoryRequirement, true
}

// HasMemoryRequirement returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasMemoryRequirement() bool {
	if o != nil && o.MemoryRequirement != nil {
		return true
	}

	return false
}

// SetMemoryRequirement gets a reference to the given string and assigns it to the MemoryRequirement field.
func (o *InstanceTypeLayoutUpdate) SetMemoryRequirement(v string) {
	o.MemoryRequirement = &v
}

// GetHasAutoScale returns the HasAutoScale field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetHasAutoScale() bool {
	if o == nil || o.HasAutoScale == nil {
		var ret bool
		return ret
	}
	return *o.HasAutoScale
}

// GetHasAutoScaleOk returns a tuple with the HasAutoScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetHasAutoScaleOk() (*bool, bool) {
	if o == nil || o.HasAutoScale == nil {
		return nil, false
	}
	return o.HasAutoScale, true
}

// HasHasAutoScale returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasHasAutoScale() bool {
	if o != nil && o.HasAutoScale != nil {
		return true
	}

	return false
}

// SetHasAutoScale gets a reference to the given bool and assigns it to the HasAutoScale field.
func (o *InstanceTypeLayoutUpdate) SetHasAutoScale(v bool) {
	o.HasAutoScale = &v
}

// GetSupportsConvertToManaged returns the SupportsConvertToManaged field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetSupportsConvertToManaged() bool {
	if o == nil || o.SupportsConvertToManaged == nil {
		var ret bool
		return ret
	}
	return *o.SupportsConvertToManaged
}

// GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetSupportsConvertToManagedOk() (*bool, bool) {
	if o == nil || o.SupportsConvertToManaged == nil {
		return nil, false
	}
	return o.SupportsConvertToManaged, true
}

// HasSupportsConvertToManaged returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasSupportsConvertToManaged() bool {
	if o != nil && o.SupportsConvertToManaged != nil {
		return true
	}

	return false
}

// SetSupportsConvertToManaged gets a reference to the given bool and assigns it to the SupportsConvertToManaged field.
func (o *InstanceTypeLayoutUpdate) SetSupportsConvertToManaged(v bool) {
	o.SupportsConvertToManaged = &v
}

// GetContainerTypes returns the ContainerTypes field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetContainerTypes() []int64 {
	if o == nil || o.ContainerTypes == nil {
		var ret []int64
		return ret
	}
	return *o.ContainerTypes
}

// GetContainerTypesOk returns a tuple with the ContainerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetContainerTypesOk() (*[]int64, bool) {
	if o == nil || o.ContainerTypes == nil {
		return nil, false
	}
	return o.ContainerTypes, true
}

// HasContainerTypes returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasContainerTypes() bool {
	if o != nil && o.ContainerTypes != nil {
		return true
	}

	return false
}

// SetContainerTypes gets a reference to the given []int64 and assigns it to the ContainerTypes field.
func (o *InstanceTypeLayoutUpdate) SetContainerTypes(v []int64) {
	o.ContainerTypes = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetOptionTypes() []int64 {
	if o == nil || o.OptionTypes == nil {
		var ret []int64
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetOptionTypesOk() (*[]int64, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []int64 and assigns it to the OptionTypes field.
func (o *InstanceTypeLayoutUpdate) SetOptionTypes(v []int64) {
	o.OptionTypes = &v
}

// GetSpecTemplates returns the SpecTemplates field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetSpecTemplates() []int64 {
	if o == nil || o.SpecTemplates == nil {
		var ret []int64
		return ret
	}
	return *o.SpecTemplates
}

// GetSpecTemplatesOk returns a tuple with the SpecTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetSpecTemplatesOk() (*[]int64, bool) {
	if o == nil || o.SpecTemplates == nil {
		return nil, false
	}
	return o.SpecTemplates, true
}

// HasSpecTemplates returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasSpecTemplates() bool {
	if o != nil && o.SpecTemplates != nil {
		return true
	}

	return false
}

// SetSpecTemplates gets a reference to the given []int64 and assigns it to the SpecTemplates field.
func (o *InstanceTypeLayoutUpdate) SetSpecTemplates(v []int64) {
	o.SpecTemplates = &v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetEnvironmentVariables() []ClusterLayoutCreateEnvironmentVariables {
	if o == nil || o.EnvironmentVariables == nil {
		var ret []ClusterLayoutCreateEnvironmentVariables
		return ret
	}
	return *o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetEnvironmentVariablesOk() (*[]ClusterLayoutCreateEnvironmentVariables, bool) {
	if o == nil || o.EnvironmentVariables == nil {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasEnvironmentVariables() bool {
	if o != nil && o.EnvironmentVariables != nil {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []ClusterLayoutCreateEnvironmentVariables and assigns it to the EnvironmentVariables field.
func (o *InstanceTypeLayoutUpdate) SetEnvironmentVariables(v []ClusterLayoutCreateEnvironmentVariables) {
	o.EnvironmentVariables = &v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetPriceSets() []InstanceTypeCreatePriceSets {
	if o == nil || o.PriceSets == nil {
		var ret []InstanceTypeCreatePriceSets
		return ret
	}
	return *o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetPriceSetsOk() (*[]InstanceTypeCreatePriceSets, bool) {
	if o == nil || o.PriceSets == nil {
		return nil, false
	}
	return o.PriceSets, true
}

// HasPriceSets returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasPriceSets() bool {
	if o != nil && o.PriceSets != nil {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []InstanceTypeCreatePriceSets and assigns it to the PriceSets field.
func (o *InstanceTypeLayoutUpdate) SetPriceSets(v []InstanceTypeCreatePriceSets) {
	o.PriceSets = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *InstanceTypeLayoutUpdate) GetPermissions() ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions {
	if o == nil || o.Permissions == nil {
		var ret ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceTypeLayoutUpdate) GetPermissionsOk() (*ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *InstanceTypeLayoutUpdate) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions and assigns it to the Permissions field.
func (o *InstanceTypeLayoutUpdate) SetPermissions(v ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions) {
	o.Permissions = &v
}

func (o InstanceTypeLayoutUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.InstanceVersion != nil {
		toSerialize["instanceVersion"] = o.InstanceVersion
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Creatable != nil {
		toSerialize["creatable"] = o.Creatable
	}
	if o.ProvisionTypeCode != nil {
		toSerialize["provisionTypeCode"] = o.ProvisionTypeCode
	}
	if o.MemoryRequirement != nil {
		toSerialize["memoryRequirement"] = o.MemoryRequirement
	}
	if o.HasAutoScale != nil {
		toSerialize["hasAutoScale"] = o.HasAutoScale
	}
	if o.SupportsConvertToManaged != nil {
		toSerialize["supportsConvertToManaged"] = o.SupportsConvertToManaged
	}
	if o.ContainerTypes != nil {
		toSerialize["containerTypes"] = o.ContainerTypes
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.SpecTemplates != nil {
		toSerialize["specTemplates"] = o.SpecTemplates
	}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if o.PriceSets != nil {
		toSerialize["priceSets"] = o.PriceSets
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceTypeLayoutUpdate struct {
	value *InstanceTypeLayoutUpdate
	isSet bool
}

func (v NullableInstanceTypeLayoutUpdate) Get() *InstanceTypeLayoutUpdate {
	return v.value
}

func (v *NullableInstanceTypeLayoutUpdate) Set(val *InstanceTypeLayoutUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceTypeLayoutUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceTypeLayoutUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceTypeLayoutUpdate(val *InstanceTypeLayoutUpdate) *NullableInstanceTypeLayoutUpdate {
	return &NullableInstanceTypeLayoutUpdate{value: val, isSet: true}
}

func (v NullableInstanceTypeLayoutUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceTypeLayoutUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

