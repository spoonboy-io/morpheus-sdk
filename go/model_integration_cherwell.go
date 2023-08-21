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

// IntegrationCherwell struct for IntegrationCherwell
type IntegrationCherwell struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Type *string `json:"type,omitempty"`
	IntegrationType *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"integrationType,omitempty"`
	Url *string `json:"url,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	PasswordHash *string `json:"passwordHash,omitempty"`
	IsPlugin *bool `json:"isPlugin,omitempty"`
	Config *IntegrationCherwellConfig `json:"config,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusDate NullableTime `json:"statusDate,omitempty"`
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	LastSync NullableString `json:"lastSync,omitempty"`
	LastSyncDuration NullableString `json:"lastSyncDuration,omitempty"`
	Credential *Creds `json:"credential,omitempty"`
}

// NewIntegrationCherwell instantiates a new IntegrationCherwell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationCherwell() *IntegrationCherwell {
	this := IntegrationCherwell{}
	return &this
}

// NewIntegrationCherwellWithDefaults instantiates a new IntegrationCherwell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationCherwellWithDefaults() *IntegrationCherwell {
	this := IntegrationCherwell{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IntegrationCherwell) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationCherwell) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationCherwell) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IntegrationCherwell) SetType(v string) {
	o.Type = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetIntegrationType() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.IntegrationType == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetIntegrationTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.IntegrationType == nil {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasIntegrationType() bool {
	if o != nil && o.IntegrationType != nil {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the IntegrationType field.
func (o *IntegrationCherwell) SetIntegrationType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.IntegrationType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IntegrationCherwell) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *IntegrationCherwell) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IntegrationCherwell) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetPasswordHash() string {
	if o == nil || o.PasswordHash == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetPasswordHashOk() (*string, bool) {
	if o == nil || o.PasswordHash == nil {
		return nil, false
	}
	return o.PasswordHash, true
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasPasswordHash() bool {
	if o != nil && o.PasswordHash != nil {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given string and assigns it to the PasswordHash field.
func (o *IntegrationCherwell) SetPasswordHash(v string) {
	o.PasswordHash = &v
}

// GetIsPlugin returns the IsPlugin field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetIsPlugin() bool {
	if o == nil || o.IsPlugin == nil {
		var ret bool
		return ret
	}
	return *o.IsPlugin
}

// GetIsPluginOk returns a tuple with the IsPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetIsPluginOk() (*bool, bool) {
	if o == nil || o.IsPlugin == nil {
		return nil, false
	}
	return o.IsPlugin, true
}

// HasIsPlugin returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasIsPlugin() bool {
	if o != nil && o.IsPlugin != nil {
		return true
	}

	return false
}

// SetIsPlugin gets a reference to the given bool and assigns it to the IsPlugin field.
func (o *IntegrationCherwell) SetIsPlugin(v bool) {
	o.IsPlugin = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetConfig() IntegrationCherwellConfig {
	if o == nil || o.Config == nil {
		var ret IntegrationCherwellConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetConfigOk() (*IntegrationCherwellConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IntegrationCherwellConfig and assigns it to the Config field.
func (o *IntegrationCherwell) SetConfig(v IntegrationCherwellConfig) {
	o.Config = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IntegrationCherwell) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationCherwell) GetStatusDate() time.Time {
	if o == nil || o.StatusDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusDate.Get()
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationCherwell) GetStatusDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDate.Get(), o.StatusDate.IsSet()
}

// HasStatusDate returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasStatusDate() bool {
	if o != nil && o.StatusDate.IsSet() {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given NullableTime and assigns it to the StatusDate field.
func (o *IntegrationCherwell) SetStatusDate(v time.Time) {
	o.StatusDate.Set(&v)
}
// SetStatusDateNil sets the value for StatusDate to be an explicit nil
func (o *IntegrationCherwell) SetStatusDateNil() {
	o.StatusDate.Set(nil)
}

// UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
func (o *IntegrationCherwell) UnsetStatusDate() {
	o.StatusDate.Unset()
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationCherwell) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationCherwell) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *IntegrationCherwell) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *IntegrationCherwell) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *IntegrationCherwell) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetLastSync returns the LastSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationCherwell) GetLastSync() string {
	if o == nil || o.LastSync.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSync.Get()
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationCherwell) GetLastSyncOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSync.Get(), o.LastSync.IsSet()
}

// HasLastSync returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasLastSync() bool {
	if o != nil && o.LastSync.IsSet() {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given NullableString and assigns it to the LastSync field.
func (o *IntegrationCherwell) SetLastSync(v string) {
	o.LastSync.Set(&v)
}
// SetLastSyncNil sets the value for LastSync to be an explicit nil
func (o *IntegrationCherwell) SetLastSyncNil() {
	o.LastSync.Set(nil)
}

// UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
func (o *IntegrationCherwell) UnsetLastSync() {
	o.LastSync.Unset()
}

// GetLastSyncDuration returns the LastSyncDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationCherwell) GetLastSyncDuration() string {
	if o == nil || o.LastSyncDuration.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSyncDuration.Get()
}

// GetLastSyncDurationOk returns a tuple with the LastSyncDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationCherwell) GetLastSyncDurationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSyncDuration.Get(), o.LastSyncDuration.IsSet()
}

// HasLastSyncDuration returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasLastSyncDuration() bool {
	if o != nil && o.LastSyncDuration.IsSet() {
		return true
	}

	return false
}

// SetLastSyncDuration gets a reference to the given NullableString and assigns it to the LastSyncDuration field.
func (o *IntegrationCherwell) SetLastSyncDuration(v string) {
	o.LastSyncDuration.Set(&v)
}
// SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil
func (o *IntegrationCherwell) SetLastSyncDurationNil() {
	o.LastSyncDuration.Set(nil)
}

// UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
func (o *IntegrationCherwell) UnsetLastSyncDuration() {
	o.LastSyncDuration.Unset()
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *IntegrationCherwell) GetCredential() Creds {
	if o == nil || o.Credential == nil {
		var ret Creds
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwell) GetCredentialOk() (*Creds, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *IntegrationCherwell) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given Creds and assigns it to the Credential field.
func (o *IntegrationCherwell) SetCredential(v Creds) {
	o.Credential = &v
}

func (o IntegrationCherwell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IntegrationType != nil {
		toSerialize["integrationType"] = o.IntegrationType
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordHash != nil {
		toSerialize["passwordHash"] = o.PasswordHash
	}
	if o.IsPlugin != nil {
		toSerialize["isPlugin"] = o.IsPlugin
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDate.IsSet() {
		toSerialize["statusDate"] = o.StatusDate.Get()
	}
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.LastSync.IsSet() {
		toSerialize["lastSync"] = o.LastSync.Get()
	}
	if o.LastSyncDuration.IsSet() {
		toSerialize["lastSyncDuration"] = o.LastSyncDuration.Get()
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationCherwell struct {
	value *IntegrationCherwell
	isSet bool
}

func (v NullableIntegrationCherwell) Get() *IntegrationCherwell {
	return v.value
}

func (v *NullableIntegrationCherwell) Set(val *IntegrationCherwell) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationCherwell) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationCherwell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationCherwell(val *IntegrationCherwell) *NullableIntegrationCherwell {
	return &NullableIntegrationCherwell{value: val, isSet: true}
}

func (v NullableIntegrationCherwell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationCherwell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


