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

// ProvisioningLicensesUpdate struct for ProvisioningLicensesUpdate
type ProvisioningLicensesUpdate struct {
	// Name
	Name *string `json:"name,omitempty"`
	// License Version
	LicenseVersion *string `json:"licenseVersion,omitempty"`
	// Copies - The number of times the key can be used.
	Copies *int64 `json:"copies,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Virtual Images - Array of Virtual Image IDs to associate with license.
	VirtualImages *[]int64 `json:"virtualImages,omitempty"`
	// Tenants - Array of tenants that are allowed to use the key.
	Tenants *[]int64 `json:"tenants,omitempty"`
}

// NewProvisioningLicensesUpdate instantiates a new ProvisioningLicensesUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningLicensesUpdate() *ProvisioningLicensesUpdate {
	this := ProvisioningLicensesUpdate{}
	var copies int64 = 1
	this.Copies = &copies
	return &this
}

// NewProvisioningLicensesUpdateWithDefaults instantiates a new ProvisioningLicensesUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningLicensesUpdateWithDefaults() *ProvisioningLicensesUpdate {
	this := ProvisioningLicensesUpdate{}
	var copies int64 = 1
	this.Copies = &copies
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProvisioningLicensesUpdate) SetName(v string) {
	o.Name = &v
}

// GetLicenseVersion returns the LicenseVersion field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetLicenseVersion() string {
	if o == nil || o.LicenseVersion == nil {
		var ret string
		return ret
	}
	return *o.LicenseVersion
}

// GetLicenseVersionOk returns a tuple with the LicenseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetLicenseVersionOk() (*string, bool) {
	if o == nil || o.LicenseVersion == nil {
		return nil, false
	}
	return o.LicenseVersion, true
}

// HasLicenseVersion returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasLicenseVersion() bool {
	if o != nil && o.LicenseVersion != nil {
		return true
	}

	return false
}

// SetLicenseVersion gets a reference to the given string and assigns it to the LicenseVersion field.
func (o *ProvisioningLicensesUpdate) SetLicenseVersion(v string) {
	o.LicenseVersion = &v
}

// GetCopies returns the Copies field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetCopies() int64 {
	if o == nil || o.Copies == nil {
		var ret int64
		return ret
	}
	return *o.Copies
}

// GetCopiesOk returns a tuple with the Copies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetCopiesOk() (*int64, bool) {
	if o == nil || o.Copies == nil {
		return nil, false
	}
	return o.Copies, true
}

// HasCopies returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasCopies() bool {
	if o != nil && o.Copies != nil {
		return true
	}

	return false
}

// SetCopies gets a reference to the given int64 and assigns it to the Copies field.
func (o *ProvisioningLicensesUpdate) SetCopies(v int64) {
	o.Copies = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProvisioningLicensesUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetVirtualImages returns the VirtualImages field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetVirtualImages() []int64 {
	if o == nil || o.VirtualImages == nil {
		var ret []int64
		return ret
	}
	return *o.VirtualImages
}

// GetVirtualImagesOk returns a tuple with the VirtualImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetVirtualImagesOk() (*[]int64, bool) {
	if o == nil || o.VirtualImages == nil {
		return nil, false
	}
	return o.VirtualImages, true
}

// HasVirtualImages returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasVirtualImages() bool {
	if o != nil && o.VirtualImages != nil {
		return true
	}

	return false
}

// SetVirtualImages gets a reference to the given []int64 and assigns it to the VirtualImages field.
func (o *ProvisioningLicensesUpdate) SetVirtualImages(v []int64) {
	o.VirtualImages = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ProvisioningLicensesUpdate) GetTenants() []int64 {
	if o == nil || o.Tenants == nil {
		var ret []int64
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningLicensesUpdate) GetTenantsOk() (*[]int64, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ProvisioningLicensesUpdate) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []int64 and assigns it to the Tenants field.
func (o *ProvisioningLicensesUpdate) SetTenants(v []int64) {
	o.Tenants = &v
}

func (o ProvisioningLicensesUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LicenseVersion != nil {
		toSerialize["licenseVersion"] = o.LicenseVersion
	}
	if o.Copies != nil {
		toSerialize["copies"] = o.Copies
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.VirtualImages != nil {
		toSerialize["virtualImages"] = o.VirtualImages
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	return json.Marshal(toSerialize)
}

type NullableProvisioningLicensesUpdate struct {
	value *ProvisioningLicensesUpdate
	isSet bool
}

func (v NullableProvisioningLicensesUpdate) Get() *ProvisioningLicensesUpdate {
	return v.value
}

func (v *NullableProvisioningLicensesUpdate) Set(val *ProvisioningLicensesUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningLicensesUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningLicensesUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningLicensesUpdate(val *ProvisioningLicensesUpdate) *NullableProvisioningLicensesUpdate {
	return &NullableProvisioningLicensesUpdate{value: val, isSet: true}
}

func (v NullableProvisioningLicensesUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningLicensesUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

