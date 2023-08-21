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

// ApiStorageServersStorageServer struct for ApiStorageServersStorageServer
type ApiStorageServersStorageServer struct {
	// Name
	Name string `json:"name"`
	// The `Storage Type` Code or ID
	Type string `json:"type"`
	// description
	Description *string `json:"description,omitempty"`
	// The enabled flag
	Enabled *bool `json:"enabled,omitempty"`
	// Configuration object with parameters that vary by `type`
	Config map[string]interface{} `json:"config"`
	// private or public
	Visibility *string `json:"visibility,omitempty"`
	// Array of tenant account ids that are allowed access
	Tenants *[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"tenants,omitempty"`
}

// NewApiStorageServersStorageServer instantiates a new ApiStorageServersStorageServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStorageServersStorageServer(name string, type_ string, config map[string]interface{}, ) *ApiStorageServersStorageServer {
	this := ApiStorageServersStorageServer{}
	this.Name = name
	this.Type = type_
	var enabled bool = true
	this.Enabled = &enabled
	this.Config = config
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// NewApiStorageServersStorageServerWithDefaults instantiates a new ApiStorageServersStorageServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStorageServersStorageServerWithDefaults() *ApiStorageServersStorageServer {
	this := ApiStorageServersStorageServer{}
	var enabled bool = true
	this.Enabled = &enabled
	var visibility string = "private"
	this.Visibility = &visibility
	return &this
}

// GetName returns the Name field value
func (o *ApiStorageServersStorageServer) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiStorageServersStorageServer) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ApiStorageServersStorageServer) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiStorageServersStorageServer) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiStorageServersStorageServer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiStorageServersStorageServer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiStorageServersStorageServer) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiStorageServersStorageServer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiStorageServersStorageServer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiStorageServersStorageServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConfig returns the Config field value
func (o *ApiStorageServersStorageServer) GetConfig() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *ApiStorageServersStorageServer) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ApiStorageServersStorageServer) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ApiStorageServersStorageServer) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ApiStorageServersStorageServer) SetVisibility(v string) {
	o.Visibility = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ApiStorageServersStorageServer) GetTenants() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Tenants == nil {
		var ret []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStorageServersStorageServer) GetTenantsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ApiStorageServersStorageServer) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Tenants field.
func (o *ApiStorageServersStorageServer) SetTenants(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Tenants = &v
}

func (o ApiStorageServersStorageServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["config"] = o.Config
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	return json.Marshal(toSerialize)
}

type NullableApiStorageServersStorageServer struct {
	value *ApiStorageServersStorageServer
	isSet bool
}

func (v NullableApiStorageServersStorageServer) Get() *ApiStorageServersStorageServer {
	return v.value
}

func (v *NullableApiStorageServersStorageServer) Set(val *ApiStorageServersStorageServer) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStorageServersStorageServer) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStorageServersStorageServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStorageServersStorageServer(val *ApiStorageServersStorageServer) *NullableApiStorageServersStorageServer {
	return &NullableApiStorageServersStorageServer{value: val, isSet: true}
}

func (v NullableApiStorageServersStorageServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStorageServersStorageServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

