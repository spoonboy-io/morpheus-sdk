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

// MetaMeta struct for MetaMeta
type MetaMeta struct {
	// Offset records, the number of records to skip, can be used to paginate over results.
	Offset *int64 `json:"offset,omitempty"`
	// Max size, the maximum number of records to include in the response.
	Max *int64 `json:"max,omitempty"`
	// Number of records returned in the response
	Size *int64 `json:"size,omitempty"`
	// Total number of records found
	Total *int64 `json:"total,omitempty"`
}

// NewMetaMeta instantiates a new MetaMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaMeta() *MetaMeta {
	this := MetaMeta{}
	var offset int64 = 0
	this.Offset = &offset
	var max int64 = 25
	this.Max = &max
	var size int64 = 0
	this.Size = &size
	var total int64 = 0
	this.Total = &total
	return &this
}

// NewMetaMetaWithDefaults instantiates a new MetaMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaMetaWithDefaults() *MetaMeta {
	this := MetaMeta{}
	var offset int64 = 0
	this.Offset = &offset
	var max int64 = 25
	this.Max = &max
	var size int64 = 0
	this.Size = &size
	var total int64 = 0
	this.Total = &total
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *MetaMeta) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMeta) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *MetaMeta) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *MetaMeta) SetOffset(v int64) {
	o.Offset = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MetaMeta) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMeta) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MetaMeta) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *MetaMeta) SetMax(v int64) {
	o.Max = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MetaMeta) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMeta) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MetaMeta) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *MetaMeta) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MetaMeta) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaMeta) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MetaMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *MetaMeta) SetTotal(v int64) {
	o.Total = &v
}

func (o MetaMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableMetaMeta struct {
	value *MetaMeta
	isSet bool
}

func (v NullableMetaMeta) Get() *MetaMeta {
	return v.value
}

func (v *NullableMetaMeta) Set(val *MetaMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaMeta(val *MetaMeta) *NullableMetaMeta {
	return &NullableMetaMeta{value: val, isSet: true}
}

func (v NullableMetaMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


