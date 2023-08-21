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

// MonitoringSettings struct for MonitoringSettings
type MonitoringSettings struct {
	// Auto Create Checks
	AutoManageChecks *bool `json:"autoManageChecks,omitempty"`
	// Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval.
	AvailabilityTimeFrame NullableInt32 `json:"availabilityTimeFrame,omitempty"`
	// Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5.
	AvailabilityPrecision NullableInt32 `json:"availabilityPrecision,omitempty"`
	// Default Check Interval. The number of minutes to use as the default interval to use when creating new checks.
	DefaultCheckInterval NullableInt32 `json:"defaultCheckInterval,omitempty"`
	ServiceNow *MonitoringSettingsServiceNow `json:"serviceNow,omitempty"`
}

// NewMonitoringSettings instantiates a new MonitoringSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSettings() *MonitoringSettings {
	this := MonitoringSettings{}
	return &this
}

// NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSettingsWithDefaults() *MonitoringSettings {
	this := MonitoringSettings{}
	return &this
}

// GetAutoManageChecks returns the AutoManageChecks field value if set, zero value otherwise.
func (o *MonitoringSettings) GetAutoManageChecks() bool {
	if o == nil || o.AutoManageChecks == nil {
		var ret bool
		return ret
	}
	return *o.AutoManageChecks
}

// GetAutoManageChecksOk returns a tuple with the AutoManageChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetAutoManageChecksOk() (*bool, bool) {
	if o == nil || o.AutoManageChecks == nil {
		return nil, false
	}
	return o.AutoManageChecks, true
}

// HasAutoManageChecks returns a boolean if a field has been set.
func (o *MonitoringSettings) HasAutoManageChecks() bool {
	if o != nil && o.AutoManageChecks != nil {
		return true
	}

	return false
}

// SetAutoManageChecks gets a reference to the given bool and assigns it to the AutoManageChecks field.
func (o *MonitoringSettings) SetAutoManageChecks(v bool) {
	o.AutoManageChecks = &v
}

// GetAvailabilityTimeFrame returns the AvailabilityTimeFrame field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettings) GetAvailabilityTimeFrame() int32 {
	if o == nil || o.AvailabilityTimeFrame.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AvailabilityTimeFrame.Get()
}

// GetAvailabilityTimeFrameOk returns a tuple with the AvailabilityTimeFrame field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettings) GetAvailabilityTimeFrameOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityTimeFrame.Get(), o.AvailabilityTimeFrame.IsSet()
}

// HasAvailabilityTimeFrame returns a boolean if a field has been set.
func (o *MonitoringSettings) HasAvailabilityTimeFrame() bool {
	if o != nil && o.AvailabilityTimeFrame.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityTimeFrame gets a reference to the given NullableInt32 and assigns it to the AvailabilityTimeFrame field.
func (o *MonitoringSettings) SetAvailabilityTimeFrame(v int32) {
	o.AvailabilityTimeFrame.Set(&v)
}
// SetAvailabilityTimeFrameNil sets the value for AvailabilityTimeFrame to be an explicit nil
func (o *MonitoringSettings) SetAvailabilityTimeFrameNil() {
	o.AvailabilityTimeFrame.Set(nil)
}

// UnsetAvailabilityTimeFrame ensures that no value is present for AvailabilityTimeFrame, not even an explicit nil
func (o *MonitoringSettings) UnsetAvailabilityTimeFrame() {
	o.AvailabilityTimeFrame.Unset()
}

// GetAvailabilityPrecision returns the AvailabilityPrecision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettings) GetAvailabilityPrecision() int32 {
	if o == nil || o.AvailabilityPrecision.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AvailabilityPrecision.Get()
}

// GetAvailabilityPrecisionOk returns a tuple with the AvailabilityPrecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettings) GetAvailabilityPrecisionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailabilityPrecision.Get(), o.AvailabilityPrecision.IsSet()
}

// HasAvailabilityPrecision returns a boolean if a field has been set.
func (o *MonitoringSettings) HasAvailabilityPrecision() bool {
	if o != nil && o.AvailabilityPrecision.IsSet() {
		return true
	}

	return false
}

// SetAvailabilityPrecision gets a reference to the given NullableInt32 and assigns it to the AvailabilityPrecision field.
func (o *MonitoringSettings) SetAvailabilityPrecision(v int32) {
	o.AvailabilityPrecision.Set(&v)
}
// SetAvailabilityPrecisionNil sets the value for AvailabilityPrecision to be an explicit nil
func (o *MonitoringSettings) SetAvailabilityPrecisionNil() {
	o.AvailabilityPrecision.Set(nil)
}

// UnsetAvailabilityPrecision ensures that no value is present for AvailabilityPrecision, not even an explicit nil
func (o *MonitoringSettings) UnsetAvailabilityPrecision() {
	o.AvailabilityPrecision.Unset()
}

// GetDefaultCheckInterval returns the DefaultCheckInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettings) GetDefaultCheckInterval() int32 {
	if o == nil || o.DefaultCheckInterval.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DefaultCheckInterval.Get()
}

// GetDefaultCheckIntervalOk returns a tuple with the DefaultCheckInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettings) GetDefaultCheckIntervalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultCheckInterval.Get(), o.DefaultCheckInterval.IsSet()
}

// HasDefaultCheckInterval returns a boolean if a field has been set.
func (o *MonitoringSettings) HasDefaultCheckInterval() bool {
	if o != nil && o.DefaultCheckInterval.IsSet() {
		return true
	}

	return false
}

// SetDefaultCheckInterval gets a reference to the given NullableInt32 and assigns it to the DefaultCheckInterval field.
func (o *MonitoringSettings) SetDefaultCheckInterval(v int32) {
	o.DefaultCheckInterval.Set(&v)
}
// SetDefaultCheckIntervalNil sets the value for DefaultCheckInterval to be an explicit nil
func (o *MonitoringSettings) SetDefaultCheckIntervalNil() {
	o.DefaultCheckInterval.Set(nil)
}

// UnsetDefaultCheckInterval ensures that no value is present for DefaultCheckInterval, not even an explicit nil
func (o *MonitoringSettings) UnsetDefaultCheckInterval() {
	o.DefaultCheckInterval.Unset()
}

// GetServiceNow returns the ServiceNow field value if set, zero value otherwise.
func (o *MonitoringSettings) GetServiceNow() MonitoringSettingsServiceNow {
	if o == nil || o.ServiceNow == nil {
		var ret MonitoringSettingsServiceNow
		return ret
	}
	return *o.ServiceNow
}

// GetServiceNowOk returns a tuple with the ServiceNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettings) GetServiceNowOk() (*MonitoringSettingsServiceNow, bool) {
	if o == nil || o.ServiceNow == nil {
		return nil, false
	}
	return o.ServiceNow, true
}

// HasServiceNow returns a boolean if a field has been set.
func (o *MonitoringSettings) HasServiceNow() bool {
	if o != nil && o.ServiceNow != nil {
		return true
	}

	return false
}

// SetServiceNow gets a reference to the given MonitoringSettingsServiceNow and assigns it to the ServiceNow field.
func (o *MonitoringSettings) SetServiceNow(v MonitoringSettingsServiceNow) {
	o.ServiceNow = &v
}

func (o MonitoringSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoManageChecks != nil {
		toSerialize["autoManageChecks"] = o.AutoManageChecks
	}
	if o.AvailabilityTimeFrame.IsSet() {
		toSerialize["availabilityTimeFrame"] = o.AvailabilityTimeFrame.Get()
	}
	if o.AvailabilityPrecision.IsSet() {
		toSerialize["availabilityPrecision"] = o.AvailabilityPrecision.Get()
	}
	if o.DefaultCheckInterval.IsSet() {
		toSerialize["defaultCheckInterval"] = o.DefaultCheckInterval.Get()
	}
	if o.ServiceNow != nil {
		toSerialize["serviceNow"] = o.ServiceNow
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSettings struct {
	value *MonitoringSettings
	isSet bool
}

func (v NullableMonitoringSettings) Get() *MonitoringSettings {
	return v.value
}

func (v *NullableMonitoringSettings) Set(val *MonitoringSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSettings(val *MonitoringSettings) *NullableMonitoringSettings {
	return &NullableMonitoringSettings{value: val, isSet: true}
}

func (v NullableMonitoringSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

