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

// BlueprintTerraformCreateSuccess struct for BlueprintTerraformCreateSuccess
type BlueprintTerraformCreateSuccess struct {
	// A name for the blueprint
	Name *string `json:"name,omitempty"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type *string `json:"type,omitempty"`
	Terraform *BlueprintTerraformCreateTerraform `json:"terraform,omitempty"`
	Config *BlueprintTerraformCreateConfig `json:"config,omitempty"`
	// Private or Public Access
	Visibility *string `json:"visibility,omitempty"`
	// Resource Permission Block
	ResourcePermission *map[string]interface{} `json:"resourcePermission,omitempty"`
	// Owner
	Owner *map[string]interface{} `json:"owner,omitempty"`
	// Tenant
	Tenant *map[string]interface{} `json:"tenant,omitempty"`
}

// NewBlueprintTerraformCreateSuccess instantiates a new BlueprintTerraformCreateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintTerraformCreateSuccess() *BlueprintTerraformCreateSuccess {
	this := BlueprintTerraformCreateSuccess{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewBlueprintTerraformCreateSuccessWithDefaults instantiates a new BlueprintTerraformCreateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintTerraformCreateSuccessWithDefaults() *BlueprintTerraformCreateSuccess {
	this := BlueprintTerraformCreateSuccess{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlueprintTerraformCreateSuccess) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BlueprintTerraformCreateSuccess) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlueprintTerraformCreateSuccess) SetType(v string) {
	o.Type = &v
}

// GetTerraform returns the Terraform field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetTerraform() BlueprintTerraformCreateTerraform {
	if o == nil || o.Terraform == nil {
		var ret BlueprintTerraformCreateTerraform
		return ret
	}
	return *o.Terraform
}

// GetTerraformOk returns a tuple with the Terraform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetTerraformOk() (*BlueprintTerraformCreateTerraform, bool) {
	if o == nil || o.Terraform == nil {
		return nil, false
	}
	return o.Terraform, true
}

// HasTerraform returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasTerraform() bool {
	if o != nil && o.Terraform != nil {
		return true
	}

	return false
}

// SetTerraform gets a reference to the given BlueprintTerraformCreateTerraform and assigns it to the Terraform field.
func (o *BlueprintTerraformCreateSuccess) SetTerraform(v BlueprintTerraformCreateTerraform) {
	o.Terraform = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetConfig() BlueprintTerraformCreateConfig {
	if o == nil || o.Config == nil {
		var ret BlueprintTerraformCreateConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetConfigOk() (*BlueprintTerraformCreateConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given BlueprintTerraformCreateConfig and assigns it to the Config field.
func (o *BlueprintTerraformCreateSuccess) SetConfig(v BlueprintTerraformCreateConfig) {
	o.Config = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BlueprintTerraformCreateSuccess) SetVisibility(v string) {
	o.Visibility = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetResourcePermission() map[string]interface{} {
	if o == nil || o.ResourcePermission == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResourcePermission == nil {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasResourcePermission() bool {
	if o != nil && o.ResourcePermission != nil {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given map[string]interface{} and assigns it to the ResourcePermission field.
func (o *BlueprintTerraformCreateSuccess) SetResourcePermission(v map[string]interface{}) {
	o.ResourcePermission = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetOwner() map[string]interface{} {
	if o == nil || o.Owner == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given map[string]interface{} and assigns it to the Owner field.
func (o *BlueprintTerraformCreateSuccess) SetOwner(v map[string]interface{}) {
	o.Owner = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *BlueprintTerraformCreateSuccess) GetTenant() map[string]interface{} {
	if o == nil || o.Tenant == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintTerraformCreateSuccess) GetTenantOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *BlueprintTerraformCreateSuccess) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *BlueprintTerraformCreateSuccess) SetTenant(v map[string]interface{}) {
	o.Tenant = &v
}

func (o BlueprintTerraformCreateSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Terraform != nil {
		toSerialize["terraform"] = o.Terraform
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ResourcePermission != nil {
		toSerialize["resourcePermission"] = o.ResourcePermission
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintTerraformCreateSuccess struct {
	value *BlueprintTerraformCreateSuccess
	isSet bool
}

func (v NullableBlueprintTerraformCreateSuccess) Get() *BlueprintTerraformCreateSuccess {
	return v.value
}

func (v *NullableBlueprintTerraformCreateSuccess) Set(val *BlueprintTerraformCreateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintTerraformCreateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintTerraformCreateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintTerraformCreateSuccess(val *BlueprintTerraformCreateSuccess) *NullableBlueprintTerraformCreateSuccess {
	return &NullableBlueprintTerraformCreateSuccess{value: val, isSet: true}
}

func (v NullableBlueprintTerraformCreateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintTerraformCreateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

