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

// checks if the AppPrepareApplyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPrepareApplyData{}

// AppPrepareApplyData struct for AppPrepareApplyData
type AppPrepareApplyData struct {
	Image *string `json:"image,omitempty"`
	Name *string `json:"name,omitempty"`
	AutoValidate *bool `json:"autoValidate,omitempty"`
	Terraform *AppPrepareApplyDataTerraform `json:"terraform,omitempty"`
	Type *string `json:"type,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
	BlueprintName *string `json:"blueprintName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	TemplateId *int64 `json:"templateId,omitempty"`
	BlueprintId *int64 `json:"blueprintId,omitempty"`
	Group *AppPrepareApplyDataGroup `json:"group,omitempty"`
}

// NewAppPrepareApplyData instantiates a new AppPrepareApplyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPrepareApplyData() *AppPrepareApplyData {
	this := AppPrepareApplyData{}
	return &this
}

// NewAppPrepareApplyDataWithDefaults instantiates a new AppPrepareApplyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPrepareApplyDataWithDefaults() *AppPrepareApplyData {
	this := AppPrepareApplyData{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *AppPrepareApplyData) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppPrepareApplyData) SetName(v string) {
	o.Name = &v
}

// GetAutoValidate returns the AutoValidate field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetAutoValidate() bool {
	if o == nil || IsNil(o.AutoValidate) {
		var ret bool
		return ret
	}
	return *o.AutoValidate
}

// GetAutoValidateOk returns a tuple with the AutoValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetAutoValidateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoValidate) {
		return nil, false
	}
	return o.AutoValidate, true
}

// HasAutoValidate returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasAutoValidate() bool {
	if o != nil && !IsNil(o.AutoValidate) {
		return true
	}

	return false
}

// SetAutoValidate gets a reference to the given bool and assigns it to the AutoValidate field.
func (o *AppPrepareApplyData) SetAutoValidate(v bool) {
	o.AutoValidate = &v
}

// GetTerraform returns the Terraform field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetTerraform() AppPrepareApplyDataTerraform {
	if o == nil || IsNil(o.Terraform) {
		var ret AppPrepareApplyDataTerraform
		return ret
	}
	return *o.Terraform
}

// GetTerraformOk returns a tuple with the Terraform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetTerraformOk() (*AppPrepareApplyDataTerraform, bool) {
	if o == nil || IsNil(o.Terraform) {
		return nil, false
	}
	return o.Terraform, true
}

// HasTerraform returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasTerraform() bool {
	if o != nil && !IsNil(o.Terraform) {
		return true
	}

	return false
}

// SetTerraform gets a reference to the given AppPrepareApplyDataTerraform and assigns it to the Terraform field.
func (o *AppPrepareApplyData) SetTerraform(v AppPrepareApplyDataTerraform) {
	o.Terraform = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppPrepareApplyData) SetType(v string) {
	o.Type = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *AppPrepareApplyData) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetBlueprintName returns the BlueprintName field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetBlueprintName() string {
	if o == nil || IsNil(o.BlueprintName) {
		var ret string
		return ret
	}
	return *o.BlueprintName
}

// GetBlueprintNameOk returns a tuple with the BlueprintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetBlueprintNameOk() (*string, bool) {
	if o == nil || IsNil(o.BlueprintName) {
		return nil, false
	}
	return o.BlueprintName, true
}

// HasBlueprintName returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasBlueprintName() bool {
	if o != nil && !IsNil(o.BlueprintName) {
		return true
	}

	return false
}

// SetBlueprintName gets a reference to the given string and assigns it to the BlueprintName field.
func (o *AppPrepareApplyData) SetBlueprintName(v string) {
	o.BlueprintName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppPrepareApplyData) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppPrepareApplyData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppPrepareApplyData) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppPrepareApplyData) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppPrepareApplyData) UnsetDescription() {
	o.Description.Unset()
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetTemplateId() int64 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int64
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetTemplateIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int64 and assigns it to the TemplateId field.
func (o *AppPrepareApplyData) SetTemplateId(v int64) {
	o.TemplateId = &v
}

// GetBlueprintId returns the BlueprintId field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetBlueprintId() int64 {
	if o == nil || IsNil(o.BlueprintId) {
		var ret int64
		return ret
	}
	return *o.BlueprintId
}

// GetBlueprintIdOk returns a tuple with the BlueprintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetBlueprintIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BlueprintId) {
		return nil, false
	}
	return o.BlueprintId, true
}

// HasBlueprintId returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasBlueprintId() bool {
	if o != nil && !IsNil(o.BlueprintId) {
		return true
	}

	return false
}

// SetBlueprintId gets a reference to the given int64 and assigns it to the BlueprintId field.
func (o *AppPrepareApplyData) SetBlueprintId(v int64) {
	o.BlueprintId = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AppPrepareApplyData) GetGroup() AppPrepareApplyDataGroup {
	if o == nil || IsNil(o.Group) {
		var ret AppPrepareApplyDataGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPrepareApplyData) GetGroupOk() (*AppPrepareApplyDataGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AppPrepareApplyData) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given AppPrepareApplyDataGroup and assigns it to the Group field.
func (o *AppPrepareApplyData) SetGroup(v AppPrepareApplyDataGroup) {
	o.Group = &v
}

func (o AppPrepareApplyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPrepareApplyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AutoValidate) {
		toSerialize["autoValidate"] = o.AutoValidate
	}
	if !IsNil(o.Terraform) {
		toSerialize["terraform"] = o.Terraform
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.BlueprintName) {
		toSerialize["blueprintName"] = o.BlueprintName
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.BlueprintId) {
		toSerialize["blueprintId"] = o.BlueprintId
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableAppPrepareApplyData struct {
	value *AppPrepareApplyData
	isSet bool
}

func (v NullableAppPrepareApplyData) Get() *AppPrepareApplyData {
	return v.value
}

func (v *NullableAppPrepareApplyData) Set(val *AppPrepareApplyData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPrepareApplyData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPrepareApplyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPrepareApplyData(val *AppPrepareApplyData) *NullableAppPrepareApplyData {
	return &NullableAppPrepareApplyData{value: val, isSet: true}
}

func (v NullableAppPrepareApplyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPrepareApplyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


