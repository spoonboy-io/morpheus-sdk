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

// ReportConfig struct for ReportConfig
type ReportConfig struct {
	ReportType NullableString `json:"reportType,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	CloudId *string `json:"cloudId,omitempty"`
}

// NewReportConfig instantiates a new ReportConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportConfig() *ReportConfig {
	this := ReportConfig{}
	return &this
}

// NewReportConfigWithDefaults instantiates a new ReportConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportConfigWithDefaults() *ReportConfig {
	this := ReportConfig{}
	return &this
}

// GetReportType returns the ReportType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportConfig) GetReportType() string {
	if o == nil || o.ReportType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReportType.Get()
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportConfig) GetReportTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReportType.Get(), o.ReportType.IsSet()
}

// HasReportType returns a boolean if a field has been set.
func (o *ReportConfig) HasReportType() bool {
	if o != nil && o.ReportType.IsSet() {
		return true
	}

	return false
}

// SetReportType gets a reference to the given NullableString and assigns it to the ReportType field.
func (o *ReportConfig) SetReportType(v string) {
	o.ReportType.Set(&v)
}
// SetReportTypeNil sets the value for ReportType to be an explicit nil
func (o *ReportConfig) SetReportTypeNil() {
	o.ReportType.Set(nil)
}

// UnsetReportType ensures that no value is present for ReportType, not even an explicit nil
func (o *ReportConfig) UnsetReportType() {
	o.ReportType.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ReportConfig) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportConfig) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ReportConfig) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ReportConfig) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ReportConfig) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportConfig) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ReportConfig) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ReportConfig) SetEndDate(v string) {
	o.EndDate = &v
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *ReportConfig) GetCloudId() string {
	if o == nil || o.CloudId == nil {
		var ret string
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportConfig) GetCloudIdOk() (*string, bool) {
	if o == nil || o.CloudId == nil {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *ReportConfig) HasCloudId() bool {
	if o != nil && o.CloudId != nil {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given string and assigns it to the CloudId field.
func (o *ReportConfig) SetCloudId(v string) {
	o.CloudId = &v
}

func (o ReportConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportType.IsSet() {
		toSerialize["reportType"] = o.ReportType.Get()
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.CloudId != nil {
		toSerialize["cloudId"] = o.CloudId
	}
	return json.Marshal(toSerialize)
}

type NullableReportConfig struct {
	value *ReportConfig
	isSet bool
}

func (v NullableReportConfig) Get() *ReportConfig {
	return v.value
}

func (v *NullableReportConfig) Set(val *ReportConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReportConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReportConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportConfig(val *ReportConfig) *NullableReportConfig {
	return &NullableReportConfig{value: val, isSet: true}
}

func (v NullableReportConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


