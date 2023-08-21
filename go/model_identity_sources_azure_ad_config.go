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

// IdentitySourcesAzureADConfig struct for IdentitySourcesAzureADConfig
type IdentitySourcesAzureADConfig struct {
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
	Config *IdentitySourcesAzureADConfigConfig `json:"config,omitempty"`
	RoleMappings *[]IdentitySourcesSAMLConfigRoleMappings `json:"roleMappings,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	LoginURL *string `json:"loginURL,omitempty"`
	ProviderSettings *IdentitySourcesSAMLConfigProviderSettings `json:"providerSettings,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewIdentitySourcesAzureADConfig instantiates a new IdentitySourcesAzureADConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourcesAzureADConfig() *IdentitySourcesAzureADConfig {
	this := IdentitySourcesAzureADConfig{}
	return &this
}

// NewIdentitySourcesAzureADConfigWithDefaults instantiates a new IdentitySourcesAzureADConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourcesAzureADConfigWithDefaults() *IdentitySourcesAzureADConfig {
	this := IdentitySourcesAzureADConfig{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IdentitySourcesAzureADConfig) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentitySourcesAzureADConfig) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentitySourcesAzureADConfig) SetDescription(v string) {
	o.Description = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IdentitySourcesAzureADConfig) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentitySourcesAzureADConfig) SetType(v string) {
	o.Type = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *IdentitySourcesAzureADConfig) SetActive(v bool) {
	o.Active = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *IdentitySourcesAzureADConfig) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetAutoSyncOnLogin returns the AutoSyncOnLogin field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetAutoSyncOnLogin() bool {
	if o == nil || o.AutoSyncOnLogin == nil {
		var ret bool
		return ret
	}
	return *o.AutoSyncOnLogin
}

// GetAutoSyncOnLoginOk returns a tuple with the AutoSyncOnLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetAutoSyncOnLoginOk() (*bool, bool) {
	if o == nil || o.AutoSyncOnLogin == nil {
		return nil, false
	}
	return o.AutoSyncOnLogin, true
}

// HasAutoSyncOnLogin returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasAutoSyncOnLogin() bool {
	if o != nil && o.AutoSyncOnLogin != nil {
		return true
	}

	return false
}

// SetAutoSyncOnLogin gets a reference to the given bool and assigns it to the AutoSyncOnLogin field.
func (o *IdentitySourcesAzureADConfig) SetAutoSyncOnLogin(v bool) {
	o.AutoSyncOnLogin = &v
}

// GetExternalLogin returns the ExternalLogin field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetExternalLogin() bool {
	if o == nil || o.ExternalLogin == nil {
		var ret bool
		return ret
	}
	return *o.ExternalLogin
}

// GetExternalLoginOk returns a tuple with the ExternalLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetExternalLoginOk() (*bool, bool) {
	if o == nil || o.ExternalLogin == nil {
		return nil, false
	}
	return o.ExternalLogin, true
}

// HasExternalLogin returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasExternalLogin() bool {
	if o != nil && o.ExternalLogin != nil {
		return true
	}

	return false
}

// SetExternalLogin gets a reference to the given bool and assigns it to the ExternalLogin field.
func (o *IdentitySourcesAzureADConfig) SetExternalLogin(v bool) {
	o.ExternalLogin = &v
}

// GetAllowCustomMappings returns the AllowCustomMappings field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetAllowCustomMappings() bool {
	if o == nil || o.AllowCustomMappings == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomMappings
}

// GetAllowCustomMappingsOk returns a tuple with the AllowCustomMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetAllowCustomMappingsOk() (*bool, bool) {
	if o == nil || o.AllowCustomMappings == nil {
		return nil, false
	}
	return o.AllowCustomMappings, true
}

// HasAllowCustomMappings returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasAllowCustomMappings() bool {
	if o != nil && o.AllowCustomMappings != nil {
		return true
	}

	return false
}

// SetAllowCustomMappings gets a reference to the given bool and assigns it to the AllowCustomMappings field.
func (o *IdentitySourcesAzureADConfig) SetAllowCustomMappings(v bool) {
	o.AllowCustomMappings = &v
}

// GetManualRoleAssignment returns the ManualRoleAssignment field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetManualRoleAssignment() bool {
	if o == nil || o.ManualRoleAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ManualRoleAssignment
}

// GetManualRoleAssignmentOk returns a tuple with the ManualRoleAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetManualRoleAssignmentOk() (*bool, bool) {
	if o == nil || o.ManualRoleAssignment == nil {
		return nil, false
	}
	return o.ManualRoleAssignment, true
}

// HasManualRoleAssignment returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasManualRoleAssignment() bool {
	if o != nil && o.ManualRoleAssignment != nil {
		return true
	}

	return false
}

// SetManualRoleAssignment gets a reference to the given bool and assigns it to the ManualRoleAssignment field.
func (o *IdentitySourcesAzureADConfig) SetManualRoleAssignment(v bool) {
	o.ManualRoleAssignment = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *IdentitySourcesAzureADConfig) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

// GetDefaultAccountRole returns the DefaultAccountRole field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetDefaultAccountRole() IdentitySourcesLDAPConfigDefaultAccountRole {
	if o == nil || o.DefaultAccountRole == nil {
		var ret IdentitySourcesLDAPConfigDefaultAccountRole
		return ret
	}
	return *o.DefaultAccountRole
}

// GetDefaultAccountRoleOk returns a tuple with the DefaultAccountRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetDefaultAccountRoleOk() (*IdentitySourcesLDAPConfigDefaultAccountRole, bool) {
	if o == nil || o.DefaultAccountRole == nil {
		return nil, false
	}
	return o.DefaultAccountRole, true
}

// HasDefaultAccountRole returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasDefaultAccountRole() bool {
	if o != nil && o.DefaultAccountRole != nil {
		return true
	}

	return false
}

// SetDefaultAccountRole gets a reference to the given IdentitySourcesLDAPConfigDefaultAccountRole and assigns it to the DefaultAccountRole field.
func (o *IdentitySourcesAzureADConfig) SetDefaultAccountRole(v IdentitySourcesLDAPConfigDefaultAccountRole) {
	o.DefaultAccountRole = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetConfig() IdentitySourcesAzureADConfigConfig {
	if o == nil || o.Config == nil {
		var ret IdentitySourcesAzureADConfigConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetConfigOk() (*IdentitySourcesAzureADConfigConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IdentitySourcesAzureADConfigConfig and assigns it to the Config field.
func (o *IdentitySourcesAzureADConfig) SetConfig(v IdentitySourcesAzureADConfigConfig) {
	o.Config = &v
}

// GetRoleMappings returns the RoleMappings field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetRoleMappings() []IdentitySourcesSAMLConfigRoleMappings {
	if o == nil || o.RoleMappings == nil {
		var ret []IdentitySourcesSAMLConfigRoleMappings
		return ret
	}
	return *o.RoleMappings
}

// GetRoleMappingsOk returns a tuple with the RoleMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetRoleMappingsOk() (*[]IdentitySourcesSAMLConfigRoleMappings, bool) {
	if o == nil || o.RoleMappings == nil {
		return nil, false
	}
	return o.RoleMappings, true
}

// HasRoleMappings returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasRoleMappings() bool {
	if o != nil && o.RoleMappings != nil {
		return true
	}

	return false
}

// SetRoleMappings gets a reference to the given []IdentitySourcesSAMLConfigRoleMappings and assigns it to the RoleMappings field.
func (o *IdentitySourcesAzureADConfig) SetRoleMappings(v []IdentitySourcesSAMLConfigRoleMappings) {
	o.RoleMappings = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *IdentitySourcesAzureADConfig) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetLoginURL returns the LoginURL field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetLoginURL() string {
	if o == nil || o.LoginURL == nil {
		var ret string
		return ret
	}
	return *o.LoginURL
}

// GetLoginURLOk returns a tuple with the LoginURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetLoginURLOk() (*string, bool) {
	if o == nil || o.LoginURL == nil {
		return nil, false
	}
	return o.LoginURL, true
}

// HasLoginURL returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasLoginURL() bool {
	if o != nil && o.LoginURL != nil {
		return true
	}

	return false
}

// SetLoginURL gets a reference to the given string and assigns it to the LoginURL field.
func (o *IdentitySourcesAzureADConfig) SetLoginURL(v string) {
	o.LoginURL = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetProviderSettings() IdentitySourcesSAMLConfigProviderSettings {
	if o == nil || o.ProviderSettings == nil {
		var ret IdentitySourcesSAMLConfigProviderSettings
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetProviderSettingsOk() (*IdentitySourcesSAMLConfigProviderSettings, bool) {
	if o == nil || o.ProviderSettings == nil {
		return nil, false
	}
	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasProviderSettings() bool {
	if o != nil && o.ProviderSettings != nil {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given IdentitySourcesSAMLConfigProviderSettings and assigns it to the ProviderSettings field.
func (o *IdentitySourcesAzureADConfig) SetProviderSettings(v IdentitySourcesSAMLConfigProviderSettings) {
	o.ProviderSettings = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *IdentitySourcesAzureADConfig) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IdentitySourcesAzureADConfig) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourcesAzureADConfig) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IdentitySourcesAzureADConfig) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IdentitySourcesAzureADConfig) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o IdentitySourcesAzureADConfig) MarshalJSON() ([]byte, error) {
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

type NullableIdentitySourcesAzureADConfig struct {
	value *IdentitySourcesAzureADConfig
	isSet bool
}

func (v NullableIdentitySourcesAzureADConfig) Get() *IdentitySourcesAzureADConfig {
	return v.value
}

func (v *NullableIdentitySourcesAzureADConfig) Set(val *IdentitySourcesAzureADConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourcesAzureADConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourcesAzureADConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourcesAzureADConfig(val *IdentitySourcesAzureADConfig) *NullableIdentitySourcesAzureADConfig {
	return &NullableIdentitySourcesAzureADConfig{value: val, isSet: true}
}

func (v NullableIdentitySourcesAzureADConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourcesAzureADConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

