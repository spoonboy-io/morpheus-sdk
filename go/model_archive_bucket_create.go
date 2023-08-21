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

// ArchiveBucketCreate struct for ArchiveBucketCreate
type ArchiveBucketCreate struct {
	// A name for the archive bucket. Must be globally unique.
	Name *string `json:"name,omitempty"`
	// A description for the archive bucket
	Description *string `json:"description,omitempty"`
	StorageProvider *ArchiveBucketCreateStorageProvider `json:"storageProvider,omitempty"`
	// Visibility - Set to public to allow all tenants
	Visibility *string `json:"visibility,omitempty"`
	// Public URL - Set to true to allow anonymous access
	IsPublic *bool `json:"isPublic,omitempty"`
	Accounts *ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"accounts,omitempty"`
}

// NewArchiveBucketCreate instantiates a new ArchiveBucketCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveBucketCreate() *ArchiveBucketCreate {
	this := ArchiveBucketCreate{}
	var visibility string = "private"
	this.Visibility = &visibility
	var isPublic bool = false
	this.IsPublic = &isPublic
	return &this
}

// NewArchiveBucketCreateWithDefaults instantiates a new ArchiveBucketCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveBucketCreateWithDefaults() *ArchiveBucketCreate {
	this := ArchiveBucketCreate{}
	var visibility string = "private"
	this.Visibility = &visibility
	var isPublic bool = false
	this.IsPublic = &isPublic
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ArchiveBucketCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArchiveBucketCreate) SetDescription(v string) {
	o.Description = &v
}

// GetStorageProvider returns the StorageProvider field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetStorageProvider() ArchiveBucketCreateStorageProvider {
	if o == nil || o.StorageProvider == nil {
		var ret ArchiveBucketCreateStorageProvider
		return ret
	}
	return *o.StorageProvider
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetStorageProviderOk() (*ArchiveBucketCreateStorageProvider, bool) {
	if o == nil || o.StorageProvider == nil {
		return nil, false
	}
	return o.StorageProvider, true
}

// HasStorageProvider returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasStorageProvider() bool {
	if o != nil && o.StorageProvider != nil {
		return true
	}

	return false
}

// SetStorageProvider gets a reference to the given ArchiveBucketCreateStorageProvider and assigns it to the StorageProvider field.
func (o *ArchiveBucketCreate) SetStorageProvider(v ArchiveBucketCreateStorageProvider) {
	o.StorageProvider = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ArchiveBucketCreate) SetVisibility(v string) {
	o.Visibility = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ArchiveBucketCreate) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ArchiveBucketCreate) GetAccounts() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Accounts == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveBucketCreate) GetAccountsOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ArchiveBucketCreate) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Accounts field.
func (o *ArchiveBucketCreate) SetAccounts(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Accounts = &v
}

func (o ArchiveBucketCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.StorageProvider != nil {
		toSerialize["storageProvider"] = o.StorageProvider
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableArchiveBucketCreate struct {
	value *ArchiveBucketCreate
	isSet bool
}

func (v NullableArchiveBucketCreate) Get() *ArchiveBucketCreate {
	return v.value
}

func (v *NullableArchiveBucketCreate) Set(val *ArchiveBucketCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveBucketCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveBucketCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveBucketCreate(val *ArchiveBucketCreate) *NullableArchiveBucketCreate {
	return &NullableArchiveBucketCreate{value: val, isSet: true}
}

func (v NullableArchiveBucketCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveBucketCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


