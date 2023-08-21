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

// ClusterApiConfig struct for ClusterApiConfig
type ClusterApiConfig struct {
	ServiceUrl NullableString `json:"serviceUrl,omitempty"`
	ServiceHost NullableString `json:"serviceHost,omitempty"`
	ServicePath NullableString `json:"servicePath,omitempty"`
	ServiceHostname NullableString `json:"serviceHostname,omitempty"`
	ServicePort *int64 `json:"servicePort,omitempty"`
	ServiceUsername NullableString `json:"serviceUsername,omitempty"`
	ServicePassword NullableString `json:"servicePassword,omitempty"`
	ServicePasswordHash NullableString `json:"servicePasswordHash,omitempty"`
	// API Token
	ServiceToken NullableString `json:"serviceToken,omitempty"`
	// Kube Config
	ServiceAccess NullableString `json:"serviceAccess,omitempty"`
	ServiceCert NullableString `json:"serviceCert,omitempty"`
	ServiceVersion NullableString `json:"serviceVersion,omitempty"`
}

// NewClusterApiConfig instantiates a new ClusterApiConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterApiConfig() *ClusterApiConfig {
	this := ClusterApiConfig{}
	return &this
}

// NewClusterApiConfigWithDefaults instantiates a new ClusterApiConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterApiConfigWithDefaults() *ClusterApiConfig {
	this := ClusterApiConfig{}
	return &this
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceUrl() string {
	if o == nil || o.ServiceUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUrl.Get()
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUrl.Get(), o.ServiceUrl.IsSet()
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceUrl() bool {
	if o != nil && o.ServiceUrl.IsSet() {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given NullableString and assigns it to the ServiceUrl field.
func (o *ClusterApiConfig) SetServiceUrl(v string) {
	o.ServiceUrl.Set(&v)
}
// SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil
func (o *ClusterApiConfig) SetServiceUrlNil() {
	o.ServiceUrl.Set(nil)
}

// UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceUrl() {
	o.ServiceUrl.Unset()
}

// GetServiceHost returns the ServiceHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceHost() string {
	if o == nil || o.ServiceHost.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceHost.Get()
}

// GetServiceHostOk returns a tuple with the ServiceHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceHost.Get(), o.ServiceHost.IsSet()
}

// HasServiceHost returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceHost() bool {
	if o != nil && o.ServiceHost.IsSet() {
		return true
	}

	return false
}

// SetServiceHost gets a reference to the given NullableString and assigns it to the ServiceHost field.
func (o *ClusterApiConfig) SetServiceHost(v string) {
	o.ServiceHost.Set(&v)
}
// SetServiceHostNil sets the value for ServiceHost to be an explicit nil
func (o *ClusterApiConfig) SetServiceHostNil() {
	o.ServiceHost.Set(nil)
}

// UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceHost() {
	o.ServiceHost.Unset()
}

// GetServicePath returns the ServicePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServicePath() string {
	if o == nil || o.ServicePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePath.Get()
}

// GetServicePathOk returns a tuple with the ServicePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServicePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePath.Get(), o.ServicePath.IsSet()
}

// HasServicePath returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServicePath() bool {
	if o != nil && o.ServicePath.IsSet() {
		return true
	}

	return false
}

// SetServicePath gets a reference to the given NullableString and assigns it to the ServicePath field.
func (o *ClusterApiConfig) SetServicePath(v string) {
	o.ServicePath.Set(&v)
}
// SetServicePathNil sets the value for ServicePath to be an explicit nil
func (o *ClusterApiConfig) SetServicePathNil() {
	o.ServicePath.Set(nil)
}

// UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
func (o *ClusterApiConfig) UnsetServicePath() {
	o.ServicePath.Unset()
}

// GetServiceHostname returns the ServiceHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceHostname() string {
	if o == nil || o.ServiceHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceHostname.Get()
}

// GetServiceHostnameOk returns a tuple with the ServiceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceHostname.Get(), o.ServiceHostname.IsSet()
}

// HasServiceHostname returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceHostname() bool {
	if o != nil && o.ServiceHostname.IsSet() {
		return true
	}

	return false
}

// SetServiceHostname gets a reference to the given NullableString and assigns it to the ServiceHostname field.
func (o *ClusterApiConfig) SetServiceHostname(v string) {
	o.ServiceHostname.Set(&v)
}
// SetServiceHostnameNil sets the value for ServiceHostname to be an explicit nil
func (o *ClusterApiConfig) SetServiceHostnameNil() {
	o.ServiceHostname.Set(nil)
}

// UnsetServiceHostname ensures that no value is present for ServiceHostname, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceHostname() {
	o.ServiceHostname.Unset()
}

// GetServicePort returns the ServicePort field value if set, zero value otherwise.
func (o *ClusterApiConfig) GetServicePort() int64 {
	if o == nil || o.ServicePort == nil {
		var ret int64
		return ret
	}
	return *o.ServicePort
}

// GetServicePortOk returns a tuple with the ServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterApiConfig) GetServicePortOk() (*int64, bool) {
	if o == nil || o.ServicePort == nil {
		return nil, false
	}
	return o.ServicePort, true
}

// HasServicePort returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServicePort() bool {
	if o != nil && o.ServicePort != nil {
		return true
	}

	return false
}

// SetServicePort gets a reference to the given int64 and assigns it to the ServicePort field.
func (o *ClusterApiConfig) SetServicePort(v int64) {
	o.ServicePort = &v
}

// GetServiceUsername returns the ServiceUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceUsername() string {
	if o == nil || o.ServiceUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceUsername.Get()
}

// GetServiceUsernameOk returns a tuple with the ServiceUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceUsername.Get(), o.ServiceUsername.IsSet()
}

// HasServiceUsername returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceUsername() bool {
	if o != nil && o.ServiceUsername.IsSet() {
		return true
	}

	return false
}

// SetServiceUsername gets a reference to the given NullableString and assigns it to the ServiceUsername field.
func (o *ClusterApiConfig) SetServiceUsername(v string) {
	o.ServiceUsername.Set(&v)
}
// SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil
func (o *ClusterApiConfig) SetServiceUsernameNil() {
	o.ServiceUsername.Set(nil)
}

// UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceUsername() {
	o.ServiceUsername.Unset()
}

// GetServicePassword returns the ServicePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServicePassword() string {
	if o == nil || o.ServicePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePassword.Get()
}

// GetServicePasswordOk returns a tuple with the ServicePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServicePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePassword.Get(), o.ServicePassword.IsSet()
}

// HasServicePassword returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServicePassword() bool {
	if o != nil && o.ServicePassword.IsSet() {
		return true
	}

	return false
}

// SetServicePassword gets a reference to the given NullableString and assigns it to the ServicePassword field.
func (o *ClusterApiConfig) SetServicePassword(v string) {
	o.ServicePassword.Set(&v)
}
// SetServicePasswordNil sets the value for ServicePassword to be an explicit nil
func (o *ClusterApiConfig) SetServicePasswordNil() {
	o.ServicePassword.Set(nil)
}

// UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
func (o *ClusterApiConfig) UnsetServicePassword() {
	o.ServicePassword.Unset()
}

// GetServicePasswordHash returns the ServicePasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServicePasswordHash() string {
	if o == nil || o.ServicePasswordHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServicePasswordHash.Get()
}

// GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServicePasswordHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServicePasswordHash.Get(), o.ServicePasswordHash.IsSet()
}

// HasServicePasswordHash returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServicePasswordHash() bool {
	if o != nil && o.ServicePasswordHash.IsSet() {
		return true
	}

	return false
}

// SetServicePasswordHash gets a reference to the given NullableString and assigns it to the ServicePasswordHash field.
func (o *ClusterApiConfig) SetServicePasswordHash(v string) {
	o.ServicePasswordHash.Set(&v)
}
// SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil
func (o *ClusterApiConfig) SetServicePasswordHashNil() {
	o.ServicePasswordHash.Set(nil)
}

// UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
func (o *ClusterApiConfig) UnsetServicePasswordHash() {
	o.ServicePasswordHash.Unset()
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceToken() string {
	if o == nil || o.ServiceToken.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceToken.Get()
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceToken.Get(), o.ServiceToken.IsSet()
}

// HasServiceToken returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceToken() bool {
	if o != nil && o.ServiceToken.IsSet() {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given NullableString and assigns it to the ServiceToken field.
func (o *ClusterApiConfig) SetServiceToken(v string) {
	o.ServiceToken.Set(&v)
}
// SetServiceTokenNil sets the value for ServiceToken to be an explicit nil
func (o *ClusterApiConfig) SetServiceTokenNil() {
	o.ServiceToken.Set(nil)
}

// UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceToken() {
	o.ServiceToken.Unset()
}

// GetServiceAccess returns the ServiceAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceAccess() string {
	if o == nil || o.ServiceAccess.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccess.Get()
}

// GetServiceAccessOk returns a tuple with the ServiceAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceAccessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceAccess.Get(), o.ServiceAccess.IsSet()
}

// HasServiceAccess returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceAccess() bool {
	if o != nil && o.ServiceAccess.IsSet() {
		return true
	}

	return false
}

// SetServiceAccess gets a reference to the given NullableString and assigns it to the ServiceAccess field.
func (o *ClusterApiConfig) SetServiceAccess(v string) {
	o.ServiceAccess.Set(&v)
}
// SetServiceAccessNil sets the value for ServiceAccess to be an explicit nil
func (o *ClusterApiConfig) SetServiceAccessNil() {
	o.ServiceAccess.Set(nil)
}

// UnsetServiceAccess ensures that no value is present for ServiceAccess, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceAccess() {
	o.ServiceAccess.Unset()
}

// GetServiceCert returns the ServiceCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceCert() string {
	if o == nil || o.ServiceCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceCert.Get()
}

// GetServiceCertOk returns a tuple with the ServiceCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceCert.Get(), o.ServiceCert.IsSet()
}

// HasServiceCert returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceCert() bool {
	if o != nil && o.ServiceCert.IsSet() {
		return true
	}

	return false
}

// SetServiceCert gets a reference to the given NullableString and assigns it to the ServiceCert field.
func (o *ClusterApiConfig) SetServiceCert(v string) {
	o.ServiceCert.Set(&v)
}
// SetServiceCertNil sets the value for ServiceCert to be an explicit nil
func (o *ClusterApiConfig) SetServiceCertNil() {
	o.ServiceCert.Set(nil)
}

// UnsetServiceCert ensures that no value is present for ServiceCert, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceCert() {
	o.ServiceCert.Unset()
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterApiConfig) GetServiceVersion() string {
	if o == nil || o.ServiceVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion.Get()
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterApiConfig) GetServiceVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServiceVersion.Get(), o.ServiceVersion.IsSet()
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *ClusterApiConfig) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion.IsSet() {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given NullableString and assigns it to the ServiceVersion field.
func (o *ClusterApiConfig) SetServiceVersion(v string) {
	o.ServiceVersion.Set(&v)
}
// SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil
func (o *ClusterApiConfig) SetServiceVersionNil() {
	o.ServiceVersion.Set(nil)
}

// UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
func (o *ClusterApiConfig) UnsetServiceVersion() {
	o.ServiceVersion.Unset()
}

func (o ClusterApiConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceUrl.IsSet() {
		toSerialize["serviceUrl"] = o.ServiceUrl.Get()
	}
	if o.ServiceHost.IsSet() {
		toSerialize["serviceHost"] = o.ServiceHost.Get()
	}
	if o.ServicePath.IsSet() {
		toSerialize["servicePath"] = o.ServicePath.Get()
	}
	if o.ServiceHostname.IsSet() {
		toSerialize["serviceHostname"] = o.ServiceHostname.Get()
	}
	if o.ServicePort != nil {
		toSerialize["servicePort"] = o.ServicePort
	}
	if o.ServiceUsername.IsSet() {
		toSerialize["serviceUsername"] = o.ServiceUsername.Get()
	}
	if o.ServicePassword.IsSet() {
		toSerialize["servicePassword"] = o.ServicePassword.Get()
	}
	if o.ServicePasswordHash.IsSet() {
		toSerialize["servicePasswordHash"] = o.ServicePasswordHash.Get()
	}
	if o.ServiceToken.IsSet() {
		toSerialize["serviceToken"] = o.ServiceToken.Get()
	}
	if o.ServiceAccess.IsSet() {
		toSerialize["serviceAccess"] = o.ServiceAccess.Get()
	}
	if o.ServiceCert.IsSet() {
		toSerialize["serviceCert"] = o.ServiceCert.Get()
	}
	if o.ServiceVersion.IsSet() {
		toSerialize["serviceVersion"] = o.ServiceVersion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableClusterApiConfig struct {
	value *ClusterApiConfig
	isSet bool
}

func (v NullableClusterApiConfig) Get() *ClusterApiConfig {
	return v.value
}

func (v *NullableClusterApiConfig) Set(val *ClusterApiConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterApiConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterApiConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterApiConfig(val *ClusterApiConfig) *NullableClusterApiConfig {
	return &NullableClusterApiConfig{value: val, isSet: true}
}

func (v NullableClusterApiConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterApiConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

