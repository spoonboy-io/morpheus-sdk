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

// ApiZonesZoneIdFoldersIdFolder struct for ApiZonesZoneIdFoldersIdFolder
type ApiZonesZoneIdFoldersIdFolder struct {
	DefaultFolder *bool `json:"defaultFolder,omitempty"`
	DefaultImage *bool `json:"defaultImage,omitempty"`
	// Activate `true` or disable `false` the folder
	Active *bool `json:"active,omitempty"`
	// Setting `private` or `public`
	Visibility *string `json:"visibility,omitempty"`
	TenantPermissions *[]ApiZonesZoneIdFoldersIdFolderTenantPermissions `json:"tenantPermissions,omitempty"`
	ResourcePermissions *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions `json:"resourcePermissions,omitempty"`
}

// NewApiZonesZoneIdFoldersIdFolder instantiates a new ApiZonesZoneIdFoldersIdFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiZonesZoneIdFoldersIdFolder() *ApiZonesZoneIdFoldersIdFolder {
	this := ApiZonesZoneIdFoldersIdFolder{}
	var defaultFolder bool = false
	this.DefaultFolder = &defaultFolder
	var defaultImage bool = false
	this.DefaultImage = &defaultImage
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewApiZonesZoneIdFoldersIdFolderWithDefaults instantiates a new ApiZonesZoneIdFoldersIdFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiZonesZoneIdFoldersIdFolderWithDefaults() *ApiZonesZoneIdFoldersIdFolder {
	this := ApiZonesZoneIdFoldersIdFolder{}
	var defaultFolder bool = false
	this.DefaultFolder = &defaultFolder
	var defaultImage bool = false
	this.DefaultImage = &defaultImage
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetDefaultFolder returns the DefaultFolder field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetDefaultFolder() bool {
	if o == nil || o.DefaultFolder == nil {
		var ret bool
		return ret
	}
	return *o.DefaultFolder
}

// GetDefaultFolderOk returns a tuple with the DefaultFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetDefaultFolderOk() (*bool, bool) {
	if o == nil || o.DefaultFolder == nil {
		return nil, false
	}
	return o.DefaultFolder, true
}

// HasDefaultFolder returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasDefaultFolder() bool {
	if o != nil && o.DefaultFolder != nil {
		return true
	}

	return false
}

// SetDefaultFolder gets a reference to the given bool and assigns it to the DefaultFolder field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetDefaultFolder(v bool) {
	o.DefaultFolder = &v
}

// GetDefaultImage returns the DefaultImage field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetDefaultImage() bool {
	if o == nil || o.DefaultImage == nil {
		var ret bool
		return ret
	}
	return *o.DefaultImage
}

// GetDefaultImageOk returns a tuple with the DefaultImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetDefaultImageOk() (*bool, bool) {
	if o == nil || o.DefaultImage == nil {
		return nil, false
	}
	return o.DefaultImage, true
}

// HasDefaultImage returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasDefaultImage() bool {
	if o != nil && o.DefaultImage != nil {
		return true
	}

	return false
}

// SetDefaultImage gets a reference to the given bool and assigns it to the DefaultImage field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetDefaultImage(v bool) {
	o.DefaultImage = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetActive(v bool) {
	o.Active = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenantPermissions returns the TenantPermissions field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetTenantPermissions() []ApiZonesZoneIdFoldersIdFolderTenantPermissions {
	if o == nil || o.TenantPermissions == nil {
		var ret []ApiZonesZoneIdFoldersIdFolderTenantPermissions
		return ret
	}
	return *o.TenantPermissions
}

// GetTenantPermissionsOk returns a tuple with the TenantPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetTenantPermissionsOk() (*[]ApiZonesZoneIdFoldersIdFolderTenantPermissions, bool) {
	if o == nil || o.TenantPermissions == nil {
		return nil, false
	}
	return o.TenantPermissions, true
}

// HasTenantPermissions returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasTenantPermissions() bool {
	if o != nil && o.TenantPermissions != nil {
		return true
	}

	return false
}

// SetTenantPermissions gets a reference to the given []ApiZonesZoneIdFoldersIdFolderTenantPermissions and assigns it to the TenantPermissions field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetTenantPermissions(v []ApiZonesZoneIdFoldersIdFolderTenantPermissions) {
	o.TenantPermissions = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *ApiZonesZoneIdFoldersIdFolder) GetResourcePermissions() ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) GetResourcePermissionsOk() (*ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *ApiZonesZoneIdFoldersIdFolder) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions and assigns it to the ResourcePermissions field.
func (o *ApiZonesZoneIdFoldersIdFolder) SetResourcePermissions(v ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) {
	o.ResourcePermissions = &v
}

func (o ApiZonesZoneIdFoldersIdFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultFolder != nil {
		toSerialize["defaultFolder"] = o.DefaultFolder
	}
	if o.DefaultImage != nil {
		toSerialize["defaultImage"] = o.DefaultImage
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.TenantPermissions != nil {
		toSerialize["tenantPermissions"] = o.TenantPermissions
	}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	return json.Marshal(toSerialize)
}

type NullableApiZonesZoneIdFoldersIdFolder struct {
	value *ApiZonesZoneIdFoldersIdFolder
	isSet bool
}

func (v NullableApiZonesZoneIdFoldersIdFolder) Get() *ApiZonesZoneIdFoldersIdFolder {
	return v.value
}

func (v *NullableApiZonesZoneIdFoldersIdFolder) Set(val *ApiZonesZoneIdFoldersIdFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableApiZonesZoneIdFoldersIdFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableApiZonesZoneIdFoldersIdFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiZonesZoneIdFoldersIdFolder(val *ApiZonesZoneIdFoldersIdFolder) *NullableApiZonesZoneIdFoldersIdFolder {
	return &NullableApiZonesZoneIdFoldersIdFolder{value: val, isSet: true}
}

func (v NullableApiZonesZoneIdFoldersIdFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiZonesZoneIdFoldersIdFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


