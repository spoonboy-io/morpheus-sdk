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

// IntegrationGitRepo struct for IntegrationGitRepo
type IntegrationGitRepo struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Type *string `json:"type,omitempty"`
	IntegrationType *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"integrationType,omitempty"`
	Url *string `json:"url,omitempty"`
	ServiceKey *InlineResponse20040AppDeployInstance `json:"serviceKey,omitempty"`
	IsPlugin *bool `json:"isPlugin,omitempty"`
	Config *IntegrationGitRepoConfig `json:"config,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusDate NullableTime `json:"statusDate,omitempty"`
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	LastSync NullableString `json:"lastSync,omitempty"`
	LastSyncDuration NullableString `json:"lastSyncDuration,omitempty"`
	Credential *Creds `json:"credential,omitempty"`
}

// NewIntegrationGitRepo instantiates a new IntegrationGitRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationGitRepo() *IntegrationGitRepo {
	this := IntegrationGitRepo{}
	return &this
}

// NewIntegrationGitRepoWithDefaults instantiates a new IntegrationGitRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationGitRepoWithDefaults() *IntegrationGitRepo {
	this := IntegrationGitRepo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IntegrationGitRepo) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationGitRepo) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationGitRepo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IntegrationGitRepo) SetType(v string) {
	o.Type = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetIntegrationType() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.IntegrationType == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetIntegrationTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.IntegrationType == nil {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasIntegrationType() bool {
	if o != nil && o.IntegrationType != nil {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the IntegrationType field.
func (o *IntegrationGitRepo) SetIntegrationType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.IntegrationType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IntegrationGitRepo) SetUrl(v string) {
	o.Url = &v
}

// GetServiceKey returns the ServiceKey field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetServiceKey() InlineResponse20040AppDeployInstance {
	if o == nil || o.ServiceKey == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.ServiceKey
}

// GetServiceKeyOk returns a tuple with the ServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetServiceKeyOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.ServiceKey == nil {
		return nil, false
	}
	return o.ServiceKey, true
}

// HasServiceKey returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasServiceKey() bool {
	if o != nil && o.ServiceKey != nil {
		return true
	}

	return false
}

// SetServiceKey gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the ServiceKey field.
func (o *IntegrationGitRepo) SetServiceKey(v InlineResponse20040AppDeployInstance) {
	o.ServiceKey = &v
}

// GetIsPlugin returns the IsPlugin field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetIsPlugin() bool {
	if o == nil || o.IsPlugin == nil {
		var ret bool
		return ret
	}
	return *o.IsPlugin
}

// GetIsPluginOk returns a tuple with the IsPlugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetIsPluginOk() (*bool, bool) {
	if o == nil || o.IsPlugin == nil {
		return nil, false
	}
	return o.IsPlugin, true
}

// HasIsPlugin returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasIsPlugin() bool {
	if o != nil && o.IsPlugin != nil {
		return true
	}

	return false
}

// SetIsPlugin gets a reference to the given bool and assigns it to the IsPlugin field.
func (o *IntegrationGitRepo) SetIsPlugin(v bool) {
	o.IsPlugin = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetConfig() IntegrationGitRepoConfig {
	if o == nil || o.Config == nil {
		var ret IntegrationGitRepoConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetConfigOk() (*IntegrationGitRepoConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given IntegrationGitRepoConfig and assigns it to the Config field.
func (o *IntegrationGitRepo) SetConfig(v IntegrationGitRepoConfig) {
	o.Config = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IntegrationGitRepo) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationGitRepo) GetStatusDate() time.Time {
	if o == nil || o.StatusDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusDate.Get()
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationGitRepo) GetStatusDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDate.Get(), o.StatusDate.IsSet()
}

// HasStatusDate returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasStatusDate() bool {
	if o != nil && o.StatusDate.IsSet() {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given NullableTime and assigns it to the StatusDate field.
func (o *IntegrationGitRepo) SetStatusDate(v time.Time) {
	o.StatusDate.Set(&v)
}
// SetStatusDateNil sets the value for StatusDate to be an explicit nil
func (o *IntegrationGitRepo) SetStatusDateNil() {
	o.StatusDate.Set(nil)
}

// UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
func (o *IntegrationGitRepo) UnsetStatusDate() {
	o.StatusDate.Unset()
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationGitRepo) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationGitRepo) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *IntegrationGitRepo) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *IntegrationGitRepo) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *IntegrationGitRepo) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetLastSync returns the LastSync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationGitRepo) GetLastSync() string {
	if o == nil || o.LastSync.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSync.Get()
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationGitRepo) GetLastSyncOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSync.Get(), o.LastSync.IsSet()
}

// HasLastSync returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasLastSync() bool {
	if o != nil && o.LastSync.IsSet() {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given NullableString and assigns it to the LastSync field.
func (o *IntegrationGitRepo) SetLastSync(v string) {
	o.LastSync.Set(&v)
}
// SetLastSyncNil sets the value for LastSync to be an explicit nil
func (o *IntegrationGitRepo) SetLastSyncNil() {
	o.LastSync.Set(nil)
}

// UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
func (o *IntegrationGitRepo) UnsetLastSync() {
	o.LastSync.Unset()
}

// GetLastSyncDuration returns the LastSyncDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationGitRepo) GetLastSyncDuration() string {
	if o == nil || o.LastSyncDuration.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastSyncDuration.Get()
}

// GetLastSyncDurationOk returns a tuple with the LastSyncDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationGitRepo) GetLastSyncDurationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSyncDuration.Get(), o.LastSyncDuration.IsSet()
}

// HasLastSyncDuration returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasLastSyncDuration() bool {
	if o != nil && o.LastSyncDuration.IsSet() {
		return true
	}

	return false
}

// SetLastSyncDuration gets a reference to the given NullableString and assigns it to the LastSyncDuration field.
func (o *IntegrationGitRepo) SetLastSyncDuration(v string) {
	o.LastSyncDuration.Set(&v)
}
// SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil
func (o *IntegrationGitRepo) SetLastSyncDurationNil() {
	o.LastSyncDuration.Set(nil)
}

// UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
func (o *IntegrationGitRepo) UnsetLastSyncDuration() {
	o.LastSyncDuration.Unset()
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *IntegrationGitRepo) GetCredential() Creds {
	if o == nil || o.Credential == nil {
		var ret Creds
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationGitRepo) GetCredentialOk() (*Creds, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *IntegrationGitRepo) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given Creds and assigns it to the Credential field.
func (o *IntegrationGitRepo) SetCredential(v Creds) {
	o.Credential = &v
}

func (o IntegrationGitRepo) MarshalJSON() ([]byte, error) {
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
	if o.ServiceKey != nil {
		toSerialize["serviceKey"] = o.ServiceKey
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

type NullableIntegrationGitRepo struct {
	value *IntegrationGitRepo
	isSet bool
}

func (v NullableIntegrationGitRepo) Get() *IntegrationGitRepo {
	return v.value
}

func (v *NullableIntegrationGitRepo) Set(val *IntegrationGitRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationGitRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationGitRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationGitRepo(val *IntegrationGitRepo) *NullableIntegrationGitRepo {
	return &NullableIntegrationGitRepo{value: val, isSet: true}
}

func (v NullableIntegrationGitRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationGitRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


