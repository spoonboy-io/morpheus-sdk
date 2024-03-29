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

// CheckType struct for CheckType
type CheckType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	DefaultInterval *int64 `json:"defaultInterval,omitempty"`
	MetricName *string `json:"metricName,omitempty"`
	InUptime *bool `json:"inUptime,omitempty"`
	CreateIncident *bool `json:"createIncident,omitempty"`
	PushOnly *bool `json:"pushOnly,omitempty"`
	TunnelSupported *bool `json:"tunnelSupported,omitempty"`
}

// NewCheckType instantiates a new CheckType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckType() *CheckType {
	this := CheckType{}
	return &this
}

// NewCheckTypeWithDefaults instantiates a new CheckType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckTypeWithDefaults() *CheckType {
	this := CheckType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CheckType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CheckType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CheckType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CheckType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CheckType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CheckType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CheckType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CheckType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CheckType) SetName(v string) {
	o.Name = &v
}

// GetDefaultInterval returns the DefaultInterval field value if set, zero value otherwise.
func (o *CheckType) GetDefaultInterval() int64 {
	if o == nil || o.DefaultInterval == nil {
		var ret int64
		return ret
	}
	return *o.DefaultInterval
}

// GetDefaultIntervalOk returns a tuple with the DefaultInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetDefaultIntervalOk() (*int64, bool) {
	if o == nil || o.DefaultInterval == nil {
		return nil, false
	}
	return o.DefaultInterval, true
}

// HasDefaultInterval returns a boolean if a field has been set.
func (o *CheckType) HasDefaultInterval() bool {
	if o != nil && o.DefaultInterval != nil {
		return true
	}

	return false
}

// SetDefaultInterval gets a reference to the given int64 and assigns it to the DefaultInterval field.
func (o *CheckType) SetDefaultInterval(v int64) {
	o.DefaultInterval = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *CheckType) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *CheckType) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *CheckType) SetMetricName(v string) {
	o.MetricName = &v
}

// GetInUptime returns the InUptime field value if set, zero value otherwise.
func (o *CheckType) GetInUptime() bool {
	if o == nil || o.InUptime == nil {
		var ret bool
		return ret
	}
	return *o.InUptime
}

// GetInUptimeOk returns a tuple with the InUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetInUptimeOk() (*bool, bool) {
	if o == nil || o.InUptime == nil {
		return nil, false
	}
	return o.InUptime, true
}

// HasInUptime returns a boolean if a field has been set.
func (o *CheckType) HasInUptime() bool {
	if o != nil && o.InUptime != nil {
		return true
	}

	return false
}

// SetInUptime gets a reference to the given bool and assigns it to the InUptime field.
func (o *CheckType) SetInUptime(v bool) {
	o.InUptime = &v
}

// GetCreateIncident returns the CreateIncident field value if set, zero value otherwise.
func (o *CheckType) GetCreateIncident() bool {
	if o == nil || o.CreateIncident == nil {
		var ret bool
		return ret
	}
	return *o.CreateIncident
}

// GetCreateIncidentOk returns a tuple with the CreateIncident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetCreateIncidentOk() (*bool, bool) {
	if o == nil || o.CreateIncident == nil {
		return nil, false
	}
	return o.CreateIncident, true
}

// HasCreateIncident returns a boolean if a field has been set.
func (o *CheckType) HasCreateIncident() bool {
	if o != nil && o.CreateIncident != nil {
		return true
	}

	return false
}

// SetCreateIncident gets a reference to the given bool and assigns it to the CreateIncident field.
func (o *CheckType) SetCreateIncident(v bool) {
	o.CreateIncident = &v
}

// GetPushOnly returns the PushOnly field value if set, zero value otherwise.
func (o *CheckType) GetPushOnly() bool {
	if o == nil || o.PushOnly == nil {
		var ret bool
		return ret
	}
	return *o.PushOnly
}

// GetPushOnlyOk returns a tuple with the PushOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetPushOnlyOk() (*bool, bool) {
	if o == nil || o.PushOnly == nil {
		return nil, false
	}
	return o.PushOnly, true
}

// HasPushOnly returns a boolean if a field has been set.
func (o *CheckType) HasPushOnly() bool {
	if o != nil && o.PushOnly != nil {
		return true
	}

	return false
}

// SetPushOnly gets a reference to the given bool and assigns it to the PushOnly field.
func (o *CheckType) SetPushOnly(v bool) {
	o.PushOnly = &v
}

// GetTunnelSupported returns the TunnelSupported field value if set, zero value otherwise.
func (o *CheckType) GetTunnelSupported() bool {
	if o == nil || o.TunnelSupported == nil {
		var ret bool
		return ret
	}
	return *o.TunnelSupported
}

// GetTunnelSupportedOk returns a tuple with the TunnelSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckType) GetTunnelSupportedOk() (*bool, bool) {
	if o == nil || o.TunnelSupported == nil {
		return nil, false
	}
	return o.TunnelSupported, true
}

// HasTunnelSupported returns a boolean if a field has been set.
func (o *CheckType) HasTunnelSupported() bool {
	if o != nil && o.TunnelSupported != nil {
		return true
	}

	return false
}

// SetTunnelSupported gets a reference to the given bool and assigns it to the TunnelSupported field.
func (o *CheckType) SetTunnelSupported(v bool) {
	o.TunnelSupported = &v
}

func (o CheckType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultInterval != nil {
		toSerialize["defaultInterval"] = o.DefaultInterval
	}
	if o.MetricName != nil {
		toSerialize["metricName"] = o.MetricName
	}
	if o.InUptime != nil {
		toSerialize["inUptime"] = o.InUptime
	}
	if o.CreateIncident != nil {
		toSerialize["createIncident"] = o.CreateIncident
	}
	if o.PushOnly != nil {
		toSerialize["pushOnly"] = o.PushOnly
	}
	if o.TunnelSupported != nil {
		toSerialize["tunnelSupported"] = o.TunnelSupported
	}
	return json.Marshal(toSerialize)
}

type NullableCheckType struct {
	value *CheckType
	isSet bool
}

func (v NullableCheckType) Get() *CheckType {
	return v.value
}

func (v *NullableCheckType) Set(val *CheckType) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckType) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckType(val *CheckType) *NullableCheckType {
	return &NullableCheckType{value: val, isSet: true}
}

func (v NullableCheckType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


