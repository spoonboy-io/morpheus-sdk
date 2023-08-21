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

// ReportRows struct for ReportRows
type ReportRows struct {
	Id *int64 `json:"id,omitempty"`
	Section *string `json:"section,omitempty"`
	Data *string `json:"data,omitempty"`
	DisplayOrder NullableString `json:"displayOrder,omitempty"`
}

// NewReportRows instantiates a new ReportRows object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportRows() *ReportRows {
	this := ReportRows{}
	return &this
}

// NewReportRowsWithDefaults instantiates a new ReportRows object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportRowsWithDefaults() *ReportRows {
	this := ReportRows{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportRows) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportRows) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ReportRows) SetId(v int64) {
	o.Id = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ReportRows) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ReportRows) HasSection() bool {
	if o != nil && o.Section != nil {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *ReportRows) SetSection(v string) {
	o.Section = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReportRows) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRows) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReportRows) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ReportRows) SetData(v string) {
	o.Data = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportRows) GetDisplayOrder() string {
	if o == nil || o.DisplayOrder.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayOrder.Get()
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportRows) GetDisplayOrderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayOrder.Get(), o.DisplayOrder.IsSet()
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *ReportRows) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder.IsSet() {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given NullableString and assigns it to the DisplayOrder field.
func (o *ReportRows) SetDisplayOrder(v string) {
	o.DisplayOrder.Set(&v)
}
// SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil
func (o *ReportRows) SetDisplayOrderNil() {
	o.DisplayOrder.Set(nil)
}

// UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
func (o *ReportRows) UnsetDisplayOrder() {
	o.DisplayOrder.Unset()
}

func (o ReportRows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DisplayOrder.IsSet() {
		toSerialize["displayOrder"] = o.DisplayOrder.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReportRows struct {
	value *ReportRows
	isSet bool
}

func (v NullableReportRows) Get() *ReportRows {
	return v.value
}

func (v *NullableReportRows) Set(val *ReportRows) {
	v.value = val
	v.isSet = true
}

func (v NullableReportRows) IsSet() bool {
	return v.isSet
}

func (v *NullableReportRows) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportRows(val *ReportRows) *NullableReportRows {
	return &NullableReportRows{value: val, isSet: true}
}

func (v NullableReportRows) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportRows) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


