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

// CheckElastic Elasticsearch check is capable of connecting to your Elasticsearch, cluster or node, verifying its health. In addition, Morpheus will also pull statistical information such as: document size, capacity, and cpu usage. 
type CheckElastic struct {
	// Unique name scoped to your account for the check
	Name *string `json:"name,omitempty"`
	// Optional description field
	Description NullableString `json:"description,omitempty"`
	CheckType *CheckElasticCheckType `json:"checkType,omitempty"`
	// Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan)
	CheckInterval *int32 `json:"checkInterval,omitempty"`
	// Used to determine if check should affect account wide availability calculations
	InUptime *bool `json:"inUptime,omitempty"`
	// Used to determine if check should be scheduled to execute
	Active *bool `json:"active,omitempty"`
	// Severity level threshold for sending notifications.
	Severity *string `json:"severity,omitempty"`
	Config *CheckElasticsearchConfig `json:"config,omitempty"`
}

// NewCheckElastic instantiates a new CheckElastic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckElastic() *CheckElastic {
	this := CheckElastic{}
	var checkInterval int32 = 300
	this.CheckInterval = &checkInterval
	var inUptime bool = true
	this.InUptime = &inUptime
	var active bool = true
	this.Active = &active
	var severity string = "critical"
	this.Severity = &severity
	return &this
}

// NewCheckElasticWithDefaults instantiates a new CheckElastic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckElasticWithDefaults() *CheckElastic {
	this := CheckElastic{}
	var checkInterval int32 = 300
	this.CheckInterval = &checkInterval
	var inUptime bool = true
	this.InUptime = &inUptime
	var active bool = true
	this.Active = &active
	var severity string = "critical"
	this.Severity = &severity
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CheckElastic) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CheckElastic) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CheckElastic) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckElastic) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckElastic) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CheckElastic) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CheckElastic) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CheckElastic) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CheckElastic) UnsetDescription() {
	o.Description.Unset()
}

// GetCheckType returns the CheckType field value if set, zero value otherwise.
func (o *CheckElastic) GetCheckType() CheckElasticCheckType {
	if o == nil || o.CheckType == nil {
		var ret CheckElasticCheckType
		return ret
	}
	return *o.CheckType
}

// GetCheckTypeOk returns a tuple with the CheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetCheckTypeOk() (*CheckElasticCheckType, bool) {
	if o == nil || o.CheckType == nil {
		return nil, false
	}
	return o.CheckType, true
}

// HasCheckType returns a boolean if a field has been set.
func (o *CheckElastic) HasCheckType() bool {
	if o != nil && o.CheckType != nil {
		return true
	}

	return false
}

// SetCheckType gets a reference to the given CheckElasticCheckType and assigns it to the CheckType field.
func (o *CheckElastic) SetCheckType(v CheckElasticCheckType) {
	o.CheckType = &v
}

// GetCheckInterval returns the CheckInterval field value if set, zero value otherwise.
func (o *CheckElastic) GetCheckInterval() int32 {
	if o == nil || o.CheckInterval == nil {
		var ret int32
		return ret
	}
	return *o.CheckInterval
}

// GetCheckIntervalOk returns a tuple with the CheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetCheckIntervalOk() (*int32, bool) {
	if o == nil || o.CheckInterval == nil {
		return nil, false
	}
	return o.CheckInterval, true
}

// HasCheckInterval returns a boolean if a field has been set.
func (o *CheckElastic) HasCheckInterval() bool {
	if o != nil && o.CheckInterval != nil {
		return true
	}

	return false
}

// SetCheckInterval gets a reference to the given int32 and assigns it to the CheckInterval field.
func (o *CheckElastic) SetCheckInterval(v int32) {
	o.CheckInterval = &v
}

// GetInUptime returns the InUptime field value if set, zero value otherwise.
func (o *CheckElastic) GetInUptime() bool {
	if o == nil || o.InUptime == nil {
		var ret bool
		return ret
	}
	return *o.InUptime
}

// GetInUptimeOk returns a tuple with the InUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetInUptimeOk() (*bool, bool) {
	if o == nil || o.InUptime == nil {
		return nil, false
	}
	return o.InUptime, true
}

// HasInUptime returns a boolean if a field has been set.
func (o *CheckElastic) HasInUptime() bool {
	if o != nil && o.InUptime != nil {
		return true
	}

	return false
}

// SetInUptime gets a reference to the given bool and assigns it to the InUptime field.
func (o *CheckElastic) SetInUptime(v bool) {
	o.InUptime = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CheckElastic) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CheckElastic) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CheckElastic) SetActive(v bool) {
	o.Active = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CheckElastic) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CheckElastic) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CheckElastic) SetSeverity(v string) {
	o.Severity = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CheckElastic) GetConfig() CheckElasticsearchConfig {
	if o == nil || o.Config == nil {
		var ret CheckElasticsearchConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckElastic) GetConfigOk() (*CheckElasticsearchConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CheckElastic) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CheckElasticsearchConfig and assigns it to the Config field.
func (o *CheckElastic) SetConfig(v CheckElasticsearchConfig) {
	o.Config = &v
}

func (o CheckElastic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CheckType != nil {
		toSerialize["checkType"] = o.CheckType
	}
	if o.CheckInterval != nil {
		toSerialize["checkInterval"] = o.CheckInterval
	}
	if o.InUptime != nil {
		toSerialize["inUptime"] = o.InUptime
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableCheckElastic struct {
	value *CheckElastic
	isSet bool
}

func (v NullableCheckElastic) Get() *CheckElastic {
	return v.value
}

func (v *NullableCheckElastic) Set(val *CheckElastic) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckElastic) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckElastic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckElastic(val *CheckElastic) *NullableCheckElastic {
	return &NullableCheckElastic{value: val, isSet: true}
}

func (v NullableCheckElastic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckElastic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


