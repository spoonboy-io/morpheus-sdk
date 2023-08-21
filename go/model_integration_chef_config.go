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

// IntegrationChefConfig struct for IntegrationChefConfig
type IntegrationChefConfig struct {
	Databags *[]IntegrationChefConfigDatabags `json:"databags,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Org *string `json:"org,omitempty"`
	ChefUser *string `json:"chefUser,omitempty"`
	UserKey *string `json:"userKey,omitempty"`
	OrgKey *string `json:"orgKey,omitempty"`
	Version *string `json:"version,omitempty"`
	ChefUseFqdn *bool `json:"chefUseFqdn,omitempty"`
	WindowsVersion *string `json:"windowsVersion,omitempty"`
	WindowsInstallUrl *string `json:"windowsInstallUrl,omitempty"`
	UserKeyHash *string `json:"userKeyHash,omitempty"`
	OrgKeyHash *string `json:"orgKeyHash,omitempty"`
}

// NewIntegrationChefConfig instantiates a new IntegrationChefConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationChefConfig() *IntegrationChefConfig {
	this := IntegrationChefConfig{}
	return &this
}

// NewIntegrationChefConfigWithDefaults instantiates a new IntegrationChefConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationChefConfigWithDefaults() *IntegrationChefConfig {
	this := IntegrationChefConfig{}
	return &this
}

// GetDatabags returns the Databags field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetDatabags() []IntegrationChefConfigDatabags {
	if o == nil || o.Databags == nil {
		var ret []IntegrationChefConfigDatabags
		return ret
	}
	return *o.Databags
}

// GetDatabagsOk returns a tuple with the Databags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetDatabagsOk() (*[]IntegrationChefConfigDatabags, bool) {
	if o == nil || o.Databags == nil {
		return nil, false
	}
	return o.Databags, true
}

// HasDatabags returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasDatabags() bool {
	if o != nil && o.Databags != nil {
		return true
	}

	return false
}

// SetDatabags gets a reference to the given []IntegrationChefConfigDatabags and assigns it to the Databags field.
func (o *IntegrationChefConfig) SetDatabags(v []IntegrationChefConfigDatabags) {
	o.Databags = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *IntegrationChefConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetOrg() string {
	if o == nil || o.Org == nil {
		var ret string
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetOrgOk() (*string, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given string and assigns it to the Org field.
func (o *IntegrationChefConfig) SetOrg(v string) {
	o.Org = &v
}

// GetChefUser returns the ChefUser field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetChefUser() string {
	if o == nil || o.ChefUser == nil {
		var ret string
		return ret
	}
	return *o.ChefUser
}

// GetChefUserOk returns a tuple with the ChefUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetChefUserOk() (*string, bool) {
	if o == nil || o.ChefUser == nil {
		return nil, false
	}
	return o.ChefUser, true
}

// HasChefUser returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasChefUser() bool {
	if o != nil && o.ChefUser != nil {
		return true
	}

	return false
}

// SetChefUser gets a reference to the given string and assigns it to the ChefUser field.
func (o *IntegrationChefConfig) SetChefUser(v string) {
	o.ChefUser = &v
}

// GetUserKey returns the UserKey field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetUserKey() string {
	if o == nil || o.UserKey == nil {
		var ret string
		return ret
	}
	return *o.UserKey
}

// GetUserKeyOk returns a tuple with the UserKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetUserKeyOk() (*string, bool) {
	if o == nil || o.UserKey == nil {
		return nil, false
	}
	return o.UserKey, true
}

// HasUserKey returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasUserKey() bool {
	if o != nil && o.UserKey != nil {
		return true
	}

	return false
}

// SetUserKey gets a reference to the given string and assigns it to the UserKey field.
func (o *IntegrationChefConfig) SetUserKey(v string) {
	o.UserKey = &v
}

// GetOrgKey returns the OrgKey field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetOrgKey() string {
	if o == nil || o.OrgKey == nil {
		var ret string
		return ret
	}
	return *o.OrgKey
}

// GetOrgKeyOk returns a tuple with the OrgKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetOrgKeyOk() (*string, bool) {
	if o == nil || o.OrgKey == nil {
		return nil, false
	}
	return o.OrgKey, true
}

// HasOrgKey returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasOrgKey() bool {
	if o != nil && o.OrgKey != nil {
		return true
	}

	return false
}

// SetOrgKey gets a reference to the given string and assigns it to the OrgKey field.
func (o *IntegrationChefConfig) SetOrgKey(v string) {
	o.OrgKey = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IntegrationChefConfig) SetVersion(v string) {
	o.Version = &v
}

// GetChefUseFqdn returns the ChefUseFqdn field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetChefUseFqdn() bool {
	if o == nil || o.ChefUseFqdn == nil {
		var ret bool
		return ret
	}
	return *o.ChefUseFqdn
}

// GetChefUseFqdnOk returns a tuple with the ChefUseFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetChefUseFqdnOk() (*bool, bool) {
	if o == nil || o.ChefUseFqdn == nil {
		return nil, false
	}
	return o.ChefUseFqdn, true
}

// HasChefUseFqdn returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasChefUseFqdn() bool {
	if o != nil && o.ChefUseFqdn != nil {
		return true
	}

	return false
}

// SetChefUseFqdn gets a reference to the given bool and assigns it to the ChefUseFqdn field.
func (o *IntegrationChefConfig) SetChefUseFqdn(v bool) {
	o.ChefUseFqdn = &v
}

// GetWindowsVersion returns the WindowsVersion field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetWindowsVersion() string {
	if o == nil || o.WindowsVersion == nil {
		var ret string
		return ret
	}
	return *o.WindowsVersion
}

// GetWindowsVersionOk returns a tuple with the WindowsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetWindowsVersionOk() (*string, bool) {
	if o == nil || o.WindowsVersion == nil {
		return nil, false
	}
	return o.WindowsVersion, true
}

// HasWindowsVersion returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasWindowsVersion() bool {
	if o != nil && o.WindowsVersion != nil {
		return true
	}

	return false
}

// SetWindowsVersion gets a reference to the given string and assigns it to the WindowsVersion field.
func (o *IntegrationChefConfig) SetWindowsVersion(v string) {
	o.WindowsVersion = &v
}

// GetWindowsInstallUrl returns the WindowsInstallUrl field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetWindowsInstallUrl() string {
	if o == nil || o.WindowsInstallUrl == nil {
		var ret string
		return ret
	}
	return *o.WindowsInstallUrl
}

// GetWindowsInstallUrlOk returns a tuple with the WindowsInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetWindowsInstallUrlOk() (*string, bool) {
	if o == nil || o.WindowsInstallUrl == nil {
		return nil, false
	}
	return o.WindowsInstallUrl, true
}

// HasWindowsInstallUrl returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasWindowsInstallUrl() bool {
	if o != nil && o.WindowsInstallUrl != nil {
		return true
	}

	return false
}

// SetWindowsInstallUrl gets a reference to the given string and assigns it to the WindowsInstallUrl field.
func (o *IntegrationChefConfig) SetWindowsInstallUrl(v string) {
	o.WindowsInstallUrl = &v
}

// GetUserKeyHash returns the UserKeyHash field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetUserKeyHash() string {
	if o == nil || o.UserKeyHash == nil {
		var ret string
		return ret
	}
	return *o.UserKeyHash
}

// GetUserKeyHashOk returns a tuple with the UserKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetUserKeyHashOk() (*string, bool) {
	if o == nil || o.UserKeyHash == nil {
		return nil, false
	}
	return o.UserKeyHash, true
}

// HasUserKeyHash returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasUserKeyHash() bool {
	if o != nil && o.UserKeyHash != nil {
		return true
	}

	return false
}

// SetUserKeyHash gets a reference to the given string and assigns it to the UserKeyHash field.
func (o *IntegrationChefConfig) SetUserKeyHash(v string) {
	o.UserKeyHash = &v
}

// GetOrgKeyHash returns the OrgKeyHash field value if set, zero value otherwise.
func (o *IntegrationChefConfig) GetOrgKeyHash() string {
	if o == nil || o.OrgKeyHash == nil {
		var ret string
		return ret
	}
	return *o.OrgKeyHash
}

// GetOrgKeyHashOk returns a tuple with the OrgKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationChefConfig) GetOrgKeyHashOk() (*string, bool) {
	if o == nil || o.OrgKeyHash == nil {
		return nil, false
	}
	return o.OrgKeyHash, true
}

// HasOrgKeyHash returns a boolean if a field has been set.
func (o *IntegrationChefConfig) HasOrgKeyHash() bool {
	if o != nil && o.OrgKeyHash != nil {
		return true
	}

	return false
}

// SetOrgKeyHash gets a reference to the given string and assigns it to the OrgKeyHash field.
func (o *IntegrationChefConfig) SetOrgKeyHash(v string) {
	o.OrgKeyHash = &v
}

func (o IntegrationChefConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Databags != nil {
		toSerialize["databags"] = o.Databags
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.ChefUser != nil {
		toSerialize["chefUser"] = o.ChefUser
	}
	if o.UserKey != nil {
		toSerialize["userKey"] = o.UserKey
	}
	if o.OrgKey != nil {
		toSerialize["orgKey"] = o.OrgKey
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ChefUseFqdn != nil {
		toSerialize["chefUseFqdn"] = o.ChefUseFqdn
	}
	if o.WindowsVersion != nil {
		toSerialize["windowsVersion"] = o.WindowsVersion
	}
	if o.WindowsInstallUrl != nil {
		toSerialize["windowsInstallUrl"] = o.WindowsInstallUrl
	}
	if o.UserKeyHash != nil {
		toSerialize["userKeyHash"] = o.UserKeyHash
	}
	if o.OrgKeyHash != nil {
		toSerialize["orgKeyHash"] = o.OrgKeyHash
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationChefConfig struct {
	value *IntegrationChefConfig
	isSet bool
}

func (v NullableIntegrationChefConfig) Get() *IntegrationChefConfig {
	return v.value
}

func (v *NullableIntegrationChefConfig) Set(val *IntegrationChefConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationChefConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationChefConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationChefConfig(val *IntegrationChefConfig) *NullableIntegrationChefConfig {
	return &NullableIntegrationChefConfig{value: val, isSet: true}
}

func (v NullableIntegrationChefConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationChefConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

