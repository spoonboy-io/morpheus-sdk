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

// MonitoringSettingsServiceNow Service Now Settings
type MonitoringSettingsServiceNow struct {
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`
	Integration *MonitoringSettingsServiceNowIntegration `json:"integration,omitempty"`
	// New Incident Action
	NewIncidentAction NullableString `json:"newIncidentAction,omitempty"`
	// Close Incident Action
	CloseIncidentAction NullableString `json:"closeIncidentAction,omitempty"`
	// Info Mapping
	InfoMapping NullableString `json:"infoMapping,omitempty"`
	// Warning Mapping
	WarningMapping NullableString `json:"warningMapping,omitempty"`
	// Critical Mapping
	CriticalMapping NullableString `json:"criticalMapping,omitempty"`
}

// NewMonitoringSettingsServiceNow instantiates a new MonitoringSettingsServiceNow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSettingsServiceNow() *MonitoringSettingsServiceNow {
	this := MonitoringSettingsServiceNow{}
	return &this
}

// NewMonitoringSettingsServiceNowWithDefaults instantiates a new MonitoringSettingsServiceNow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSettingsServiceNowWithDefaults() *MonitoringSettingsServiceNow {
	this := MonitoringSettingsServiceNow{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MonitoringSettingsServiceNow) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettingsServiceNow) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MonitoringSettingsServiceNow) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *MonitoringSettingsServiceNow) GetIntegration() MonitoringSettingsServiceNowIntegration {
	if o == nil || o.Integration == nil {
		var ret MonitoringSettingsServiceNowIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSettingsServiceNow) GetIntegrationOk() (*MonitoringSettingsServiceNowIntegration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given MonitoringSettingsServiceNowIntegration and assigns it to the Integration field.
func (o *MonitoringSettingsServiceNow) SetIntegration(v MonitoringSettingsServiceNowIntegration) {
	o.Integration = &v
}

// GetNewIncidentAction returns the NewIncidentAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettingsServiceNow) GetNewIncidentAction() string {
	if o == nil || o.NewIncidentAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.NewIncidentAction.Get()
}

// GetNewIncidentActionOk returns a tuple with the NewIncidentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettingsServiceNow) GetNewIncidentActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NewIncidentAction.Get(), o.NewIncidentAction.IsSet()
}

// HasNewIncidentAction returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasNewIncidentAction() bool {
	if o != nil && o.NewIncidentAction.IsSet() {
		return true
	}

	return false
}

// SetNewIncidentAction gets a reference to the given NullableString and assigns it to the NewIncidentAction field.
func (o *MonitoringSettingsServiceNow) SetNewIncidentAction(v string) {
	o.NewIncidentAction.Set(&v)
}
// SetNewIncidentActionNil sets the value for NewIncidentAction to be an explicit nil
func (o *MonitoringSettingsServiceNow) SetNewIncidentActionNil() {
	o.NewIncidentAction.Set(nil)
}

// UnsetNewIncidentAction ensures that no value is present for NewIncidentAction, not even an explicit nil
func (o *MonitoringSettingsServiceNow) UnsetNewIncidentAction() {
	o.NewIncidentAction.Unset()
}

// GetCloseIncidentAction returns the CloseIncidentAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettingsServiceNow) GetCloseIncidentAction() string {
	if o == nil || o.CloseIncidentAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.CloseIncidentAction.Get()
}

// GetCloseIncidentActionOk returns a tuple with the CloseIncidentAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettingsServiceNow) GetCloseIncidentActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CloseIncidentAction.Get(), o.CloseIncidentAction.IsSet()
}

// HasCloseIncidentAction returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasCloseIncidentAction() bool {
	if o != nil && o.CloseIncidentAction.IsSet() {
		return true
	}

	return false
}

// SetCloseIncidentAction gets a reference to the given NullableString and assigns it to the CloseIncidentAction field.
func (o *MonitoringSettingsServiceNow) SetCloseIncidentAction(v string) {
	o.CloseIncidentAction.Set(&v)
}
// SetCloseIncidentActionNil sets the value for CloseIncidentAction to be an explicit nil
func (o *MonitoringSettingsServiceNow) SetCloseIncidentActionNil() {
	o.CloseIncidentAction.Set(nil)
}

// UnsetCloseIncidentAction ensures that no value is present for CloseIncidentAction, not even an explicit nil
func (o *MonitoringSettingsServiceNow) UnsetCloseIncidentAction() {
	o.CloseIncidentAction.Unset()
}

// GetInfoMapping returns the InfoMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettingsServiceNow) GetInfoMapping() string {
	if o == nil || o.InfoMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.InfoMapping.Get()
}

// GetInfoMappingOk returns a tuple with the InfoMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettingsServiceNow) GetInfoMappingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InfoMapping.Get(), o.InfoMapping.IsSet()
}

// HasInfoMapping returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasInfoMapping() bool {
	if o != nil && o.InfoMapping.IsSet() {
		return true
	}

	return false
}

// SetInfoMapping gets a reference to the given NullableString and assigns it to the InfoMapping field.
func (o *MonitoringSettingsServiceNow) SetInfoMapping(v string) {
	o.InfoMapping.Set(&v)
}
// SetInfoMappingNil sets the value for InfoMapping to be an explicit nil
func (o *MonitoringSettingsServiceNow) SetInfoMappingNil() {
	o.InfoMapping.Set(nil)
}

// UnsetInfoMapping ensures that no value is present for InfoMapping, not even an explicit nil
func (o *MonitoringSettingsServiceNow) UnsetInfoMapping() {
	o.InfoMapping.Unset()
}

// GetWarningMapping returns the WarningMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettingsServiceNow) GetWarningMapping() string {
	if o == nil || o.WarningMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.WarningMapping.Get()
}

// GetWarningMappingOk returns a tuple with the WarningMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettingsServiceNow) GetWarningMappingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WarningMapping.Get(), o.WarningMapping.IsSet()
}

// HasWarningMapping returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasWarningMapping() bool {
	if o != nil && o.WarningMapping.IsSet() {
		return true
	}

	return false
}

// SetWarningMapping gets a reference to the given NullableString and assigns it to the WarningMapping field.
func (o *MonitoringSettingsServiceNow) SetWarningMapping(v string) {
	o.WarningMapping.Set(&v)
}
// SetWarningMappingNil sets the value for WarningMapping to be an explicit nil
func (o *MonitoringSettingsServiceNow) SetWarningMappingNil() {
	o.WarningMapping.Set(nil)
}

// UnsetWarningMapping ensures that no value is present for WarningMapping, not even an explicit nil
func (o *MonitoringSettingsServiceNow) UnsetWarningMapping() {
	o.WarningMapping.Unset()
}

// GetCriticalMapping returns the CriticalMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringSettingsServiceNow) GetCriticalMapping() string {
	if o == nil || o.CriticalMapping.Get() == nil {
		var ret string
		return ret
	}
	return *o.CriticalMapping.Get()
}

// GetCriticalMappingOk returns a tuple with the CriticalMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringSettingsServiceNow) GetCriticalMappingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CriticalMapping.Get(), o.CriticalMapping.IsSet()
}

// HasCriticalMapping returns a boolean if a field has been set.
func (o *MonitoringSettingsServiceNow) HasCriticalMapping() bool {
	if o != nil && o.CriticalMapping.IsSet() {
		return true
	}

	return false
}

// SetCriticalMapping gets a reference to the given NullableString and assigns it to the CriticalMapping field.
func (o *MonitoringSettingsServiceNow) SetCriticalMapping(v string) {
	o.CriticalMapping.Set(&v)
}
// SetCriticalMappingNil sets the value for CriticalMapping to be an explicit nil
func (o *MonitoringSettingsServiceNow) SetCriticalMappingNil() {
	o.CriticalMapping.Set(nil)
}

// UnsetCriticalMapping ensures that no value is present for CriticalMapping, not even an explicit nil
func (o *MonitoringSettingsServiceNow) UnsetCriticalMapping() {
	o.CriticalMapping.Unset()
}

func (o MonitoringSettingsServiceNow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.NewIncidentAction.IsSet() {
		toSerialize["newIncidentAction"] = o.NewIncidentAction.Get()
	}
	if o.CloseIncidentAction.IsSet() {
		toSerialize["closeIncidentAction"] = o.CloseIncidentAction.Get()
	}
	if o.InfoMapping.IsSet() {
		toSerialize["infoMapping"] = o.InfoMapping.Get()
	}
	if o.WarningMapping.IsSet() {
		toSerialize["warningMapping"] = o.WarningMapping.Get()
	}
	if o.CriticalMapping.IsSet() {
		toSerialize["criticalMapping"] = o.CriticalMapping.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMonitoringSettingsServiceNow struct {
	value *MonitoringSettingsServiceNow
	isSet bool
}

func (v NullableMonitoringSettingsServiceNow) Get() *MonitoringSettingsServiceNow {
	return v.value
}

func (v *NullableMonitoringSettingsServiceNow) Set(val *MonitoringSettingsServiceNow) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSettingsServiceNow) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSettingsServiceNow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSettingsServiceNow(val *MonitoringSettingsServiceNow) *NullableMonitoringSettingsServiceNow {
	return &NullableMonitoringSettingsServiceNow{value: val, isSet: true}
}

func (v NullableMonitoringSettingsServiceNow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSettingsServiceNow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


