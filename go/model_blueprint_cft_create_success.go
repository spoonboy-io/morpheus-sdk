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

// BlueprintCFTCreateSuccess struct for BlueprintCFTCreateSuccess
type BlueprintCFTCreateSuccess struct {
	// A name for the blueprint
	Name *string `json:"name,omitempty"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type *string `json:"type,omitempty"`
	CloudFormation *BlueprintCFTCreateSuccessCloudFormation `json:"cloudFormation,omitempty"`
	// Private or Public Access
	Visibility *string `json:"visibility,omitempty"`
	// Resource Permission Block
	ResourcePermission *map[string]interface{} `json:"resourcePermission,omitempty"`
	// Owner
	Owner *map[string]interface{} `json:"owner,omitempty"`
	// Tenant
	Tenant *map[string]interface{} `json:"tenant,omitempty"`
}

// NewBlueprintCFTCreateSuccess instantiates a new BlueprintCFTCreateSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintCFTCreateSuccess() *BlueprintCFTCreateSuccess {
	this := BlueprintCFTCreateSuccess{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewBlueprintCFTCreateSuccessWithDefaults instantiates a new BlueprintCFTCreateSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintCFTCreateSuccessWithDefaults() *BlueprintCFTCreateSuccess {
	this := BlueprintCFTCreateSuccess{}
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlueprintCFTCreateSuccess) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BlueprintCFTCreateSuccess) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlueprintCFTCreateSuccess) SetType(v string) {
	o.Type = &v
}

// GetCloudFormation returns the CloudFormation field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetCloudFormation() BlueprintCFTCreateSuccessCloudFormation {
	if o == nil || o.CloudFormation == nil {
		var ret BlueprintCFTCreateSuccessCloudFormation
		return ret
	}
	return *o.CloudFormation
}

// GetCloudFormationOk returns a tuple with the CloudFormation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetCloudFormationOk() (*BlueprintCFTCreateSuccessCloudFormation, bool) {
	if o == nil || o.CloudFormation == nil {
		return nil, false
	}
	return o.CloudFormation, true
}

// HasCloudFormation returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasCloudFormation() bool {
	if o != nil && o.CloudFormation != nil {
		return true
	}

	return false
}

// SetCloudFormation gets a reference to the given BlueprintCFTCreateSuccessCloudFormation and assigns it to the CloudFormation field.
func (o *BlueprintCFTCreateSuccess) SetCloudFormation(v BlueprintCFTCreateSuccessCloudFormation) {
	o.CloudFormation = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *BlueprintCFTCreateSuccess) SetVisibility(v string) {
	o.Visibility = &v
}

// GetResourcePermission returns the ResourcePermission field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetResourcePermission() map[string]interface{} {
	if o == nil || o.ResourcePermission == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ResourcePermission
}

// GetResourcePermissionOk returns a tuple with the ResourcePermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResourcePermission == nil {
		return nil, false
	}
	return o.ResourcePermission, true
}

// HasResourcePermission returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasResourcePermission() bool {
	if o != nil && o.ResourcePermission != nil {
		return true
	}

	return false
}

// SetResourcePermission gets a reference to the given map[string]interface{} and assigns it to the ResourcePermission field.
func (o *BlueprintCFTCreateSuccess) SetResourcePermission(v map[string]interface{}) {
	o.ResourcePermission = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetOwner() map[string]interface{} {
	if o == nil || o.Owner == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given map[string]interface{} and assigns it to the Owner field.
func (o *BlueprintCFTCreateSuccess) SetOwner(v map[string]interface{}) {
	o.Owner = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *BlueprintCFTCreateSuccess) GetTenant() map[string]interface{} {
	if o == nil || o.Tenant == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintCFTCreateSuccess) GetTenantOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *BlueprintCFTCreateSuccess) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *BlueprintCFTCreateSuccess) SetTenant(v map[string]interface{}) {
	o.Tenant = &v
}

func (o BlueprintCFTCreateSuccess) MarshalJSON() ([]byte, error) {
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
	if o.CloudFormation != nil {
		toSerialize["cloudFormation"] = o.CloudFormation
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

type NullableBlueprintCFTCreateSuccess struct {
	value *BlueprintCFTCreateSuccess
	isSet bool
}

func (v NullableBlueprintCFTCreateSuccess) Get() *BlueprintCFTCreateSuccess {
	return v.value
}

func (v *NullableBlueprintCFTCreateSuccess) Set(val *BlueprintCFTCreateSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintCFTCreateSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintCFTCreateSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintCFTCreateSuccess(val *BlueprintCFTCreateSuccess) *NullableBlueprintCFTCreateSuccess {
	return &NullableBlueprintCFTCreateSuccess{value: val, isSet: true}
}

func (v NullableBlueprintCFTCreateSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintCFTCreateSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


