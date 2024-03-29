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
	"time"
)

// IdentitySourcesLDAPConfig struct for IdentitySourcesLDAPConfig
type IdentitySourcesLDAPConfig struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Code *string `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
	Active *bool `json:"active,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	AutoSyncOnLogin *bool `json:"autoSyncOnLogin,omitempty"`
	ExternalLogin *bool `json:"externalLogin,omitempty"`
	AllowCustomMappings *bool `json:"allowCustomMappings,omitempty"`
	ManualRoleAssignment *bool `json:"manualRoleAssignment,omitempty"`
	Account *InlineResponse20040AppDeployInstance `json:"account,omitempty"`
	DefaultAccountRole *IdentitySourcesLDAPConfigDefaultAccountRole `json:"defaultAccountRole,omitempty"`
	Config *IdentitySourcesLDAPConfigConfig `json:"config,omitempty"`
	RoleMappings *[]IdentitySourcesLDAPConfigRoleMappings `json:"roleMappings,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	LoginURL *string `json:"loginURL,omitempty"`
	ProviderSettings *map[string]interface{} `json:"providerSettings,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewIdentitySourcesLDAPConfig instantiates a new IdentitySourcesLDAPConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourcesLDAPConfig() *IdentitySourcesLDAPConfig {
	this := IdentitySourcesLDAPConfig{}
	return &this
}

// NewIdentitySourcesLDAPConfigWithDefaults instantiates a new IdentitySourcesLDAPConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourcesLDAPConfigWithDefaults() *IdentitySourcesLDAPConfig {
	this := IdentitySourcesLDAPConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentitySourcesLDAPConfig) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentitySourcesLDAPConfig) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentitySourcesLDAPConfig) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IdentitySourcesLDAPConfig) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentitySourcesLDAPConfig) SetType(v string) {
	o.Type = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *IdentitySourcesLDAPConfig) SetActive(v bool) {
	o.Active = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *IdentitySourcesLDAPConfig) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetAutoSyncOnLogin returns the AutoSyncOnLogin field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetAutoSyncOnLogin() bool {
	if o == nil || o.AutoSyncOnLogin == nil {
		var ret bool
		return ret
	}
	return *o.AutoSyncOnLogin
}

// GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetAutoSyncOnLoginOk() (*bool, bool) {
	if o == nil || o.AutoSyncOnLogin == nil {
		return nil, false
	}
	return o.AutoSyncOnLogin, true
}

// HasAutoSyncOnLogin returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasAutoSyncOnLogin() bool {
	if o != nil && o.AutoSyncOnLogin != nil {
		return true
	}

	return false
}

// SetAutoSyncOnLogin gets a reference to the given bool and assigns it to the AutoSyncOnLogin field.
func (o *IdentitySourcesLDAPConfig) SetAutoSyncOnLogin(v bool) {
	o.AutoSyncOnLogin = &v
}

// GetExternalLogin returns the ExternalLogin field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetExternalLogin() bool {
	if o == nil || o.ExternalLogin == nil {
		var ret bool
		return ret
	}
	return *o.ExternalLogin
}

// GetExternalLoginOk returns a tuple with the ExternalLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetExternalLoginOk() (*bool, bool) {
	if o == nil || o.ExternalLogin == nil {
		return nil, false
	}
	return o.ExternalLogin, true
}

// HasExternalLogin returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasExternalLogin() bool {
	if o != nil && o.ExternalLogin != nil {
		return true
	}

	return false
}

// SetExternalLogin gets a reference to the given bool and assigns it to the ExternalLogin field.
func (o *IdentitySourcesLDAPConfig) SetExternalLogin(v bool) {
	o.ExternalLogin = &v
}

// GetAllowCustomMappings returns the AllowCustomMappings field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetAllowCustomMappings() bool {
	if o == nil || o.AllowCustomMappings == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomMappings
}

// GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetAllowCustomMappingsOk() (*bool, bool) {
	if o == nil || o.AllowCustomMappings == nil {
		return nil, false
	}
	return o.AllowCustomMappings, true
}

// HasAllowCustomMappings returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasAllowCustomMappings() bool {
	if o != nil && o.AllowCustomMappings != nil {
		return true
	}

	return false
}

// SetAllowCustomMappings gets a reference to the given bool and assigns it to the AllowCustomMappings field.
func (o *IdentitySourcesLDAPConfig) SetAllowCustomMappings(v bool) {
	o.AllowCustomMappings = &v
}

// GetManualRoleAssignment returns the ManualRoleAssignment field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetManualRoleAssignment() bool {
	if o == nil || o.ManualRoleAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ManualRoleAssignment
}

// GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetManualRoleAssignmentOk() (*bool, bool) {
	if o == nil || o.ManualRoleAssignment == nil {
		return nil, false
	}
	return o.ManualRoleAssignment, true
}

// HasManualRoleAssignment returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasManualRoleAssignment() bool {
	if o != nil && o.ManualRoleAssignment != nil {
		return true
	}

	return false
}

// SetManualRoleAssignment gets a reference to the given bool and assigns it to the ManualRoleAssignment field.
func (o *IdentitySourcesLDAPConfig) SetManualRoleAssignment(v bool) {
	o.ManualRoleAssignment = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *IdentitySourcesLDAPConfig) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

// GetDefaultAccountRole returns the DefaultAccountRole field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetDefaultAccountRole() IdentitySourcesLDAPConfigDefaultAccountRole {
	if o == nil || o.DefaultAccountRole == nil {
		var ret IdentitySourcesLDAPConfigDefaultAccountRole
		return ret
	}
	return *o.DefaultAccountRole
}

// GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetDefaultAccountRoleOk() (*IdentitySourcesLDAPConfigDefaultAccountRole, bool) {
	if o == nil || o.DefaultAccountRole == nil {
		return nil, false
	}
	return o.DefaultAccountRole, true
}

// HasDefaultAccountRole returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasDefaultAccountRole() bool {
	if o != nil && o.DefaultAccountRole != nil {
		return true
	}

	return false
}

// SetDefaultAccountRole gets a reference to the given IdentitySourcesLDAPConfigDefaultAccountRole and assigns it to the DefaultAccountRole field.
func (o *IdentitySourcesLDAPConfig) SetDefaultAccountRole(v IdentitySourcesLDAPConfigDefaultAccountRole) {
	o.DefaultAccountRole = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetConfig() IdentitySourcesLDAPConfigConfig {
	if o == nil || o.Config == nil {
		var ret IdentitySourcesLDAPConfigConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetConfigOk() (*IdentitySourcesLDAPConfigConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentitySourcesLDAPConfigConfig and assigns it to the Config field.
func (o *IdentitySourcesLDAPConfig) SetConfig(v IdentitySourcesLDAPConfigConfig) {
	o.Config = &v
}

// GetRoleMappings returns the RoleMappings field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetRoleMappings() []IdentitySourcesLDAPConfigRoleMappings {
	if o == nil || o.RoleMappings == nil {
		var ret []IdentitySourcesLDAPConfigRoleMappings
		return ret
	}
	return *o.RoleMappings
}

// GetRoleMappingsOk returns a tuple with the RoleMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetRoleMappingsOk() (*[]IdentitySourcesLDAPConfigRoleMappings, bool) {
	if o == nil || o.RoleMappings == nil {
		return nil, false
	}
	return o.RoleMappings, true
}

// HasRoleMappings returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasRoleMappings() bool {
	if o != nil && o.RoleMappings != nil {
		return true
	}

	return false
}

// SetRoleMappings gets a reference to the given []IdentitySourcesLDAPConfigRoleMappings and assigns it to the RoleMappings field.
func (o *IdentitySourcesLDAPConfig) SetRoleMappings(v []IdentitySourcesLDAPConfigRoleMappings) {
	o.RoleMappings = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *IdentitySourcesLDAPConfig) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetLoginURL returns the LoginURL field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetLoginURL() string {
	if o == nil || o.LoginURL == nil {
		var ret string
		return ret
	}
	return *o.LoginURL
}

// GetLoginURLOk returns a tuple with the LoginURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetLoginURLOk() (*string, bool) {
	if o == nil || o.LoginURL == nil {
		return nil, false
	}
	return o.LoginURL, true
}

// HasLoginURL returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasLoginURL() bool {
	if o != nil && o.LoginURL != nil {
		return true
	}

	return false
}

// SetLoginURL gets a reference to the given string and assigns it to the LoginURL field.
func (o *IdentitySourcesLDAPConfig) SetLoginURL(v string) {
	o.LoginURL = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetProviderSettings() map[string]interface{} {
	if o == nil || o.ProviderSettings == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetProviderSettingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.ProviderSettings == nil {
		return nil, false
	}
	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasProviderSettings() bool {
	if o != nil && o.ProviderSettings != nil {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given map[string]interface{} and assigns it to the ProviderSettings field.
func (o *IdentitySourcesLDAPConfig) SetProviderSettings(v map[string]interface{}) {
	o.ProviderSettings = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *IdentitySourcesLDAPConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IdentitySourcesLDAPConfig) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesLDAPConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IdentitySourcesLDAPConfig) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IdentitySourcesLDAPConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o IdentitySourcesLDAPConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.AutoSyncOnLogin != nil {
		toSerialize["autoSyncOnLogin"] = o.AutoSyncOnLogin
	}
	if o.ExternalLogin != nil {
		toSerialize["externalLogin"] = o.ExternalLogin
	}
	if o.AllowCustomMappings != nil {
		toSerialize["allowCustomMappings"] = o.AllowCustomMappings
	}
	if o.ManualRoleAssignment != nil {
		toSerialize["manualRoleAssignment"] = o.ManualRoleAssignment
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.DefaultAccountRole != nil {
		toSerialize["defaultAccountRole"] = o.DefaultAccountRole
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.RoleMappings != nil {
		toSerialize["roleMappings"] = o.RoleMappings
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.LoginURL != nil {
		toSerialize["loginURL"] = o.LoginURL
	}
	if o.ProviderSettings != nil {
		toSerialize["providerSettings"] = o.ProviderSettings
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableIdentitySourcesLDAPConfig struct {
	value *IdentitySourcesLDAPConfig
	isSet bool
}

func (v NullableIdentitySourcesLDAPConfig) Get() *IdentitySourcesLDAPConfig {
	return v.value
}

func (v *NullableIdentitySourcesLDAPConfig) Set(val *IdentitySourcesLDAPConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourcesLDAPConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourcesLDAPConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourcesLDAPConfig(val *IdentitySourcesLDAPConfig) *NullableIdentitySourcesLDAPConfig {
	return &NullableIdentitySourcesLDAPConfig{value: val, isSet: true}
}

func (v NullableIdentitySourcesLDAPConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourcesLDAPConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


