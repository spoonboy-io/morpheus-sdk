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

// BlueprintCreateSuccess struct for BlueprintCreateSuccess
type BlueprintCreateSuccess struct {
	// Blueprint ID
	Id *int64 `json:"id,omitempty"`
	// A name for the blueprint
	Name *string `json:"name,omitempty"`
	// A description for the blueprint
	Description NullableString `json:"description,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	// Category
	Category NullableString `json:"category,omitempty"`
	Config *OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess `json:"config,omitempty"`
}

// NewBlueprintCreateSuccess instantiates a new BlueprintCreateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintCreateSuccess() *BlueprintCreateSuccess {
	this := BlueprintCreateSuccess{}
	return &this
}

// NewBlueprintCreateSuccessWithDefaults instantiates a new BlueprintCreateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintCreateSuccessWithDefaults() *BlueprintCreateSuccess {
	this := BlueprintCreateSuccess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlueprintCreateSuccess) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCreateSuccess) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BlueprintCreateSuccess) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlueprintCreateSuccess) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCreateSuccess) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlueprintCreateSuccess) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintCreateSuccess) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintCreateSuccess) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *BlueprintCreateSuccess) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *BlueprintCreateSuccess) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *BlueprintCreateSuccess) UnsetDescription() {
	o.Description.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BlueprintCreateSuccess) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCreateSuccess) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintCreateSuccess) SetLabels(v []string) {
	o.Labels = &v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintCreateSuccess) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintCreateSuccess) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *BlueprintCreateSuccess) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *BlueprintCreateSuccess) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *BlueprintCreateSuccess) UnsetCategory() {
	o.Category.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *BlueprintCreateSuccess) GetConfig() OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess {
	if o == nil || o.Config == nil {
		var ret OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCreateSuccess) GetConfigOk() (*OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *BlueprintCreateSuccess) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess and assigns it to the Config field.
func (o *BlueprintCreateSuccess) SetConfig(v OneOfblueprintARMCreateSuccessblueprintCFTCreateSuccessblueprintHelmCreateSuccessblueprintKubernetesCreateSuccessblueprintMorpheusCreateSuccessblueprintTerraformCreateSuccess) {
	o.Config = &v
}

func (o BlueprintCreateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintCreateSuccess struct {
	value *BlueprintCreateSuccess
	isSet bool
}

func (v NullableBlueprintCreateSuccess) Get() *BlueprintCreateSuccess {
	return v.value
}

func (v *NullableBlueprintCreateSuccess) Set(val *BlueprintCreateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintCreateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintCreateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintCreateSuccess(val *BlueprintCreateSuccess) *NullableBlueprintCreateSuccess {
	return &NullableBlueprintCreateSuccess{value: val, isSet: true}
}

func (v NullableBlueprintCreateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintCreateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

