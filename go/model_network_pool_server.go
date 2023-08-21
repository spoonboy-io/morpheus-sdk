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

// NetworkPoolServer struct for NetworkPoolServer
type NetworkPoolServer struct {
	// Network Pool Server ID
	Id *int64 `json:"id,omitempty"`
	Type *NetworkPoolServerType `json:"type,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Service URL
	ServiceUrl NullableString `json:"serviceUrl,omitempty"`
	// Service Host
	ServiceHost NullableString `json:"serviceHost,omitempty"`
	// Service Port
	ServicePort NullableInt32 `json:"servicePort,omitempty"`
	// Service Mode
	ServiceMode NullableString `json:"serviceMode,omitempty"`
	// Service Username
	ServiceUsername NullableString `json:"serviceUsername,omitempty"`
	// Service Password
	ServicePassword NullableString `json:"servicePassword,omitempty"`
	ServicePasswordHash *string `json:"servicePasswordHash,omitempty"`
	// Throttle Rate
	ServiceThrottleRate NullableInt64 `json:"serviceThrottleRate,omitempty"`
	// Disable SSL SNI Verification
	IgnoreSsl NullableBool `json:"ignoreSsl,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
	// Status Message
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	StatusDate NullableTime `json:"statusDate,omitempty"`
	// Config object varies with pool server type.
	Config *map[string]interface{} `json:"config,omitempty"`
	// Network Filter
	NetworkFilter NullableString `json:"networkFilter,omitempty"`
	// Zone Filter
	ZoneFilter NullableString `json:"zoneFilter,omitempty"`
	// Tenant Match
	TenantMatch NullableString `json:"tenantMatch,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Account *NetworkPoolServerAccount `json:"account,omitempty"`
	Integration *NetworkPoolServerIntegration `json:"integration,omitempty"`
	Pools *[]InlineResponse20040AppDeployInstance `json:"pools,omitempty"`
	Credential *Creds2 `json:"credential,omitempty"`
}

// NewNetworkPoolServer instantiates a new NetworkPoolServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolServer() *NetworkPoolServer {
	this := NetworkPoolServer{}
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = *NewNullableInt64(&serviceThrottleRate)
	var ignoreSsl bool = true
	this.IgnoreSsl = *NewNullableBool(&ignoreSsl)
	return &this
}

// NewNetworkPoolServerWithDefaults instantiates a new NetworkPoolServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolServerWithDefaults() *NetworkPoolServer {
	this := NetworkPoolServer{}
	var serviceThrottleRate int64 = 0
	this.ServiceThrottleRate = *NewNullableInt64(&serviceThrottleRate)
	var ignoreSsl bool = true
	this.IgnoreSsl = *NewNullableBool(&ignoreSsl)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NetworkPoolServer) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetType() NetworkPoolServerType {
	if o == nil || o.Type == nil {
		var ret NetworkPoolServerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetTypeOk() (*NetworkPoolServerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given NetworkPoolServerType and assigns it to the Type field.
func (o *NetworkPoolServer) SetType(v NetworkPoolServerType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkPoolServer) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkPoolServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServiceUrl() string {
	if o == nil || o.ServiceUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUrl.Get()
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServiceUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUrl.Get(), o.ServiceUrl.IsSet()
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServiceUrl() bool {
	if o != nil && o.ServiceUrl.IsSet() {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given NullableString and assigns it to the ServiceUrl field.
func (o *NetworkPoolServer) SetServiceUrl(v string) {
	o.ServiceUrl.Set(&v)
}
// SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil
func (o *NetworkPoolServer) SetServiceUrlNil() {
	o.ServiceUrl.Set(nil)
}

// UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
func (o *NetworkPoolServer) UnsetServiceUrl() {
	o.ServiceUrl.Unset()
}

// GetServiceHost returns the ServiceHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServiceHost() string {
	if o == nil || o.ServiceHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceHost.Get()
}

// GetServiceHostOk returns a tuple with the ServiceHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServiceHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceHost.Get(), o.ServiceHost.IsSet()
}

// HasServiceHost returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServiceHost() bool {
	if o != nil && o.ServiceHost.IsSet() {
		return true
	}

	return false
}

// SetServiceHost gets a reference to the given NullableString and assigns it to the ServiceHost field.
func (o *NetworkPoolServer) SetServiceHost(v string) {
	o.ServiceHost.Set(&v)
}
// SetServiceHostNil sets the value for ServiceHost to be an explicit nil
func (o *NetworkPoolServer) SetServiceHostNil() {
	o.ServiceHost.Set(nil)
}

// UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
func (o *NetworkPoolServer) UnsetServiceHost() {
	o.ServiceHost.Unset()
}

// GetServicePort returns the ServicePort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServicePort() int32 {
	if o == nil || o.ServicePort.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ServicePort.Get()
}

// GetServicePortOk returns a tuple with the ServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServicePortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePort.Get(), o.ServicePort.IsSet()
}

// HasServicePort returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServicePort() bool {
	if o != nil && o.ServicePort.IsSet() {
		return true
	}

	return false
}

// SetServicePort gets a reference to the given NullableInt32 and assigns it to the ServicePort field.
func (o *NetworkPoolServer) SetServicePort(v int32) {
	o.ServicePort.Set(&v)
}
// SetServicePortNil sets the value for ServicePort to be an explicit nil
func (o *NetworkPoolServer) SetServicePortNil() {
	o.ServicePort.Set(nil)
}

// UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
func (o *NetworkPoolServer) UnsetServicePort() {
	o.ServicePort.Unset()
}

// GetServiceMode returns the ServiceMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServiceMode() string {
	if o == nil || o.ServiceMode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceMode.Get()
}

// GetServiceModeOk returns a tuple with the ServiceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServiceModeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceMode.Get(), o.ServiceMode.IsSet()
}

// HasServiceMode returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServiceMode() bool {
	if o != nil && o.ServiceMode.IsSet() {
		return true
	}

	return false
}

// SetServiceMode gets a reference to the given NullableString and assigns it to the ServiceMode field.
func (o *NetworkPoolServer) SetServiceMode(v string) {
	o.ServiceMode.Set(&v)
}
// SetServiceModeNil sets the value for ServiceMode to be an explicit nil
func (o *NetworkPoolServer) SetServiceModeNil() {
	o.ServiceMode.Set(nil)
}

// UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
func (o *NetworkPoolServer) UnsetServiceMode() {
	o.ServiceMode.Unset()
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServiceUsername() string {
	if o == nil || o.ServiceUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUsername.Get()
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUsername.Get(), o.ServiceUsername.IsSet()
}

// HasServiceUsername returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServiceUsername() bool {
	if o != nil && o.ServiceUsername.IsSet() {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given NullableString and assigns it to the ServiceUsername field.
func (o *NetworkPoolServer) SetServiceUsername(v string) {
	o.ServiceUsername.Set(&v)
}
// SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil
func (o *NetworkPoolServer) SetServiceUsernameNil() {
	o.ServiceUsername.Set(nil)
}

// UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
func (o *NetworkPoolServer) UnsetServiceUsername() {
	o.ServiceUsername.Unset()
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServicePassword() string {
	if o == nil || o.ServicePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePassword.Get()
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServicePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePassword.Get(), o.ServicePassword.IsSet()
}

// HasServicePassword returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServicePassword() bool {
	if o != nil && o.ServicePassword.IsSet() {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given NullableString and assigns it to the ServicePassword field.
func (o *NetworkPoolServer) SetServicePassword(v string) {
	o.ServicePassword.Set(&v)
}
// SetServicePasswordNil sets the value for ServicePassword to be an explicit nil
func (o *NetworkPoolServer) SetServicePasswordNil() {
	o.ServicePassword.Set(nil)
}

// UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
func (o *NetworkPoolServer) UnsetServicePassword() {
	o.ServicePassword.Unset()
}

// GetServicePasswordHash returns the ServicePasswordHash field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetServicePasswordHash() string {
	if o == nil || o.ServicePasswordHash == nil {
		var ret string
		return ret
	}
	return *o.ServicePasswordHash
}

// GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetServicePasswordHashOk() (*string, bool) {
	if o == nil || o.ServicePasswordHash == nil {
		return nil, false
	}
	return o.ServicePasswordHash, true
}

// HasServicePasswordHash returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServicePasswordHash() bool {
	if o != nil && o.ServicePasswordHash != nil {
		return true
	}

	return false
}

// SetServicePasswordHash gets a reference to the given string and assigns it to the ServicePasswordHash field.
func (o *NetworkPoolServer) SetServicePasswordHash(v string) {
	o.ServicePasswordHash = &v
}

// GetServiceThrottleRate returns the ServiceThrottleRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetServiceThrottleRate() int64 {
	if o == nil || o.ServiceThrottleRate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ServiceThrottleRate.Get()
}

// GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceThrottleRate.Get(), o.ServiceThrottleRate.IsSet()
}

// HasServiceThrottleRate returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasServiceThrottleRate() bool {
	if o != nil && o.ServiceThrottleRate.IsSet() {
		return true
	}

	return false
}

// SetServiceThrottleRate gets a reference to the given NullableInt64 and assigns it to the ServiceThrottleRate field.
func (o *NetworkPoolServer) SetServiceThrottleRate(v int64) {
	o.ServiceThrottleRate.Set(&v)
}
// SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil
func (o *NetworkPoolServer) SetServiceThrottleRateNil() {
	o.ServiceThrottleRate.Set(nil)
}

// UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
func (o *NetworkPoolServer) UnsetServiceThrottleRate() {
	o.ServiceThrottleRate.Unset()
}

// GetIgnoreSsl returns the IgnoreSsl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetIgnoreSsl() bool {
	if o == nil || o.IgnoreSsl.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IgnoreSsl.Get()
}

// GetIgnoreSslOk returns a tuple with the IgnoreSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetIgnoreSslOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IgnoreSsl.Get(), o.IgnoreSsl.IsSet()
}

// HasIgnoreSsl returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasIgnoreSsl() bool {
	if o != nil && o.IgnoreSsl.IsSet() {
		return true
	}

	return false
}

// SetIgnoreSsl gets a reference to the given NullableBool and assigns it to the IgnoreSsl field.
func (o *NetworkPoolServer) SetIgnoreSsl(v bool) {
	o.IgnoreSsl.Set(&v)
}
// SetIgnoreSslNil sets the value for IgnoreSsl to be an explicit nil
func (o *NetworkPoolServer) SetIgnoreSslNil() {
	o.IgnoreSsl.Set(nil)
}

// UnsetIgnoreSsl ensures that no value is present for IgnoreSsl, not even an explicit nil
func (o *NetworkPoolServer) UnsetIgnoreSsl() {
	o.IgnoreSsl.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkPoolServer) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *NetworkPoolServer) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *NetworkPoolServer) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *NetworkPoolServer) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetStatusDate returns the StatusDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetStatusDate() time.Time {
	if o == nil || o.StatusDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusDate.Get()
}

// GetStatusDateOk returns a tuple with the StatusDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetStatusDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDate.Get(), o.StatusDate.IsSet()
}

// HasStatusDate returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasStatusDate() bool {
	if o != nil && o.StatusDate.IsSet() {
		return true
	}

	return false
}

// SetStatusDate gets a reference to the given NullableTime and assigns it to the StatusDate field.
func (o *NetworkPoolServer) SetStatusDate(v time.Time) {
	o.StatusDate.Set(&v)
}
// SetStatusDateNil sets the value for StatusDate to be an explicit nil
func (o *NetworkPoolServer) SetStatusDateNil() {
	o.StatusDate.Set(nil)
}

// UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
func (o *NetworkPoolServer) UnsetStatusDate() {
	o.StatusDate.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *NetworkPoolServer) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetNetworkFilter returns the NetworkFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetNetworkFilter() string {
	if o == nil || o.NetworkFilter.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkFilter.Get()
}

// GetNetworkFilterOk returns a tuple with the NetworkFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetNetworkFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkFilter.Get(), o.NetworkFilter.IsSet()
}

// HasNetworkFilter returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasNetworkFilter() bool {
	if o != nil && o.NetworkFilter.IsSet() {
		return true
	}

	return false
}

// SetNetworkFilter gets a reference to the given NullableString and assigns it to the NetworkFilter field.
func (o *NetworkPoolServer) SetNetworkFilter(v string) {
	o.NetworkFilter.Set(&v)
}
// SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil
func (o *NetworkPoolServer) SetNetworkFilterNil() {
	o.NetworkFilter.Set(nil)
}

// UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
func (o *NetworkPoolServer) UnsetNetworkFilter() {
	o.NetworkFilter.Unset()
}

// GetZoneFilter returns the ZoneFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetZoneFilter() string {
	if o == nil || o.ZoneFilter.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZoneFilter.Get()
}

// GetZoneFilterOk returns a tuple with the ZoneFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetZoneFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZoneFilter.Get(), o.ZoneFilter.IsSet()
}

// HasZoneFilter returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasZoneFilter() bool {
	if o != nil && o.ZoneFilter.IsSet() {
		return true
	}

	return false
}

// SetZoneFilter gets a reference to the given NullableString and assigns it to the ZoneFilter field.
func (o *NetworkPoolServer) SetZoneFilter(v string) {
	o.ZoneFilter.Set(&v)
}
// SetZoneFilterNil sets the value for ZoneFilter to be an explicit nil
func (o *NetworkPoolServer) SetZoneFilterNil() {
	o.ZoneFilter.Set(nil)
}

// UnsetZoneFilter ensures that no value is present for ZoneFilter, not even an explicit nil
func (o *NetworkPoolServer) UnsetZoneFilter() {
	o.ZoneFilter.Unset()
}

// GetTenantMatch returns the TenantMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkPoolServer) GetTenantMatch() string {
	if o == nil || o.TenantMatch.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantMatch.Get()
}

// GetTenantMatchOk returns a tuple with the TenantMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkPoolServer) GetTenantMatchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantMatch.Get(), o.TenantMatch.IsSet()
}

// HasTenantMatch returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasTenantMatch() bool {
	if o != nil && o.TenantMatch.IsSet() {
		return true
	}

	return false
}

// SetTenantMatch gets a reference to the given NullableString and assigns it to the TenantMatch field.
func (o *NetworkPoolServer) SetTenantMatch(v string) {
	o.TenantMatch.Set(&v)
}
// SetTenantMatchNil sets the value for TenantMatch to be an explicit nil
func (o *NetworkPoolServer) SetTenantMatchNil() {
	o.TenantMatch.Set(nil)
}

// UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
func (o *NetworkPoolServer) UnsetTenantMatch() {
	o.TenantMatch.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *NetworkPoolServer) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *NetworkPoolServer) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetAccount() NetworkPoolServerAccount {
	if o == nil || o.Account == nil {
		var ret NetworkPoolServerAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetAccountOk() (*NetworkPoolServerAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NetworkPoolServerAccount and assigns it to the Account field.
func (o *NetworkPoolServer) SetAccount(v NetworkPoolServerAccount) {
	o.Account = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetIntegration() NetworkPoolServerIntegration {
	if o == nil || o.Integration == nil {
		var ret NetworkPoolServerIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetIntegrationOk() (*NetworkPoolServerIntegration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given NetworkPoolServerIntegration and assigns it to the Integration field.
func (o *NetworkPoolServer) SetIntegration(v NetworkPoolServerIntegration) {
	o.Integration = &v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetPools() []InlineResponse20040AppDeployInstance {
	if o == nil || o.Pools == nil {
		var ret []InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetPoolsOk() (*[]InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Pools == nil {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasPools() bool {
	if o != nil && o.Pools != nil {
		return true
	}

	return false
}

// SetPools gets a reference to the given []InlineResponse20040AppDeployInstance and assigns it to the Pools field.
func (o *NetworkPoolServer) SetPools(v []InlineResponse20040AppDeployInstance) {
	o.Pools = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *NetworkPoolServer) GetCredential() Creds2 {
	if o == nil || o.Credential == nil {
		var ret Creds2
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPoolServer) GetCredentialOk() (*Creds2, bool) {
	if o == nil || o.Credential == nil {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *NetworkPoolServer) HasCredential() bool {
	if o != nil && o.Credential != nil {
		return true
	}

	return false
}

// SetCredential gets a reference to the given Creds2 and assigns it to the Credential field.
func (o *NetworkPoolServer) SetCredential(v Creds2) {
	o.Credential = &v
}

func (o NetworkPoolServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ServiceUrl.IsSet() {
		toSerialize["serviceUrl"] = o.ServiceUrl.Get()
	}
	if o.ServiceHost.IsSet() {
		toSerialize["serviceHost"] = o.ServiceHost.Get()
	}
	if o.ServicePort.IsSet() {
		toSerialize["servicePort"] = o.ServicePort.Get()
	}
	if o.ServiceMode.IsSet() {
		toSerialize["serviceMode"] = o.ServiceMode.Get()
	}
	if o.ServiceUsername.IsSet() {
		toSerialize["serviceUsername"] = o.ServiceUsername.Get()
	}
	if o.ServicePassword.IsSet() {
		toSerialize["servicePassword"] = o.ServicePassword.Get()
	}
	if o.ServicePasswordHash != nil {
		toSerialize["servicePasswordHash"] = o.ServicePasswordHash
	}
	if o.ServiceThrottleRate.IsSet() {
		toSerialize["serviceThrottleRate"] = o.ServiceThrottleRate.Get()
	}
	if o.IgnoreSsl.IsSet() {
		toSerialize["ignoreSsl"] = o.IgnoreSsl.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.StatusDate.IsSet() {
		toSerialize["statusDate"] = o.StatusDate.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.NetworkFilter.IsSet() {
		toSerialize["networkFilter"] = o.NetworkFilter.Get()
	}
	if o.ZoneFilter.IsSet() {
		toSerialize["zoneFilter"] = o.ZoneFilter.Get()
	}
	if o.TenantMatch.IsSet() {
		toSerialize["tenantMatch"] = o.TenantMatch.Get()
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.Pools != nil {
		toSerialize["pools"] = o.Pools
	}
	if o.Credential != nil {
		toSerialize["credential"] = o.Credential
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkPoolServer struct {
	value *NetworkPoolServer
	isSet bool
}

func (v NullableNetworkPoolServer) Get() *NetworkPoolServer {
	return v.value
}

func (v *NullableNetworkPoolServer) Set(val *NetworkPoolServer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolServer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolServer(val *NetworkPoolServer) *NullableNetworkPoolServer {
	return &NullableNetworkPoolServer{value: val, isSet: true}
}

func (v NullableNetworkPoolServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


