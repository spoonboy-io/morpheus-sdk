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

// ClusterUpdate struct for ClusterUpdate
type ClusterUpdate struct {
	// Cluster name
	Name *string `json:"name,omitempty"`
	// Cluster description
	Description *string `json:"description,omitempty"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Cluster enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Cluster API Url
	ServiceUrl *string `json:"serviceUrl,omitempty"`
	// Cluster API token
	ServiceToken *string `json:"serviceToken,omitempty"`
	// Queue cluster refresh
	Refresh *bool `json:"refresh,omitempty"`
	// Cluster managed
	Managed *bool `json:"managed,omitempty"`
}

// NewClusterUpdate instantiates a new ClusterUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterUpdate() *ClusterUpdate {
	this := ClusterUpdate{}
	return &this
}

// NewClusterUpdateWithDefaults instantiates a new ClusterUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterUpdateWithDefaults() *ClusterUpdate {
	this := ClusterUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterUpdate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterUpdate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ClusterUpdate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ClusterUpdate) SetLabels(v []string) {
	o.Labels = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClusterUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise.
func (o *ClusterUpdate) GetServiceUrl() string {
	if o == nil || o.ServiceUrl == nil {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetServiceUrlOk() (*string, bool) {
	if o == nil || o.ServiceUrl == nil {
		return nil, false
	}
	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *ClusterUpdate) HasServiceUrl() bool {
	if o != nil && o.ServiceUrl != nil {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *ClusterUpdate) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *ClusterUpdate) GetServiceToken() string {
	if o == nil || o.ServiceToken == nil {
		var ret string
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetServiceTokenOk() (*string, bool) {
	if o == nil || o.ServiceToken == nil {
		return nil, false
	}
	return o.ServiceToken, true
}

// HasServiceToken returns a boolean if a field has been set.
func (o *ClusterUpdate) HasServiceToken() bool {
	if o != nil && o.ServiceToken != nil {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given string and assigns it to the ServiceToken field.
func (o *ClusterUpdate) SetServiceToken(v string) {
	o.ServiceToken = &v
}

// GetRefresh returns the Refresh field value if set, zero value otherwise.
func (o *ClusterUpdate) GetRefresh() bool {
	if o == nil || o.Refresh == nil {
		var ret bool
		return ret
	}
	return *o.Refresh
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetRefreshOk() (*bool, bool) {
	if o == nil || o.Refresh == nil {
		return nil, false
	}
	return o.Refresh, true
}

// HasRefresh returns a boolean if a field has been set.
func (o *ClusterUpdate) HasRefresh() bool {
	if o != nil && o.Refresh != nil {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given bool and assigns it to the Refresh field.
func (o *ClusterUpdate) SetRefresh(v bool) {
	o.Refresh = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ClusterUpdate) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ClusterUpdate) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ClusterUpdate) SetManaged(v bool) {
	o.Managed = &v
}

func (o ClusterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ServiceUrl != nil {
		toSerialize["serviceUrl"] = o.ServiceUrl
	}
	if o.ServiceToken != nil {
		toSerialize["serviceToken"] = o.ServiceToken
	}
	if o.Refresh != nil {
		toSerialize["refresh"] = o.Refresh
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	return json.Marshal(toSerialize)
}

type NullableClusterUpdate struct {
	value *ClusterUpdate
	isSet bool
}

func (v NullableClusterUpdate) Get() *ClusterUpdate {
	return v.value
}

func (v *NullableClusterUpdate) Set(val *ClusterUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterUpdate(val *ClusterUpdate) *NullableClusterUpdate {
	return &NullableClusterUpdate{value: val, isSet: true}
}

func (v NullableClusterUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


