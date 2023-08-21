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

// IntegrationCherwellConfig struct for IntegrationCherwellConfig
type IntegrationCherwellConfig struct {
	CherwellCustomCmdbMapping *string `json:"cherwellCustomCmdbMapping,omitempty"`
	CherwellClientKey *string `json:"cherwellClientKey,omitempty"`
	CherwellCreatedBy *string `json:"cherwellCreatedBy,omitempty"`
	CherwellStartDate *string `json:"cherwellStartDate,omitempty"`
	CherwellEndDate *string `json:"cherwellEndDate,omitempty"`
	CherwellIgnoreSSLErrors NullableString `json:"cherwellIgnoreSSLErrors,omitempty"`
	CherwellBusinessObject *string `json:"cherwellBusinessObject,omitempty"`
}

// NewIntegrationCherwellConfig instantiates a new IntegrationCherwellConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationCherwellConfig() *IntegrationCherwellConfig {
	this := IntegrationCherwellConfig{}
	return &this
}

// NewIntegrationCherwellConfigWithDefaults instantiates a new IntegrationCherwellConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationCherwellConfigWithDefaults() *IntegrationCherwellConfig {
	this := IntegrationCherwellConfig{}
	return &this
}

// GetCherwellCustomCmdbMapping returns the CherwellCustomCmdbMapping field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellCustomCmdbMapping() string {
	if o == nil || o.CherwellCustomCmdbMapping == nil {
		var ret string
		return ret
	}
	return *o.CherwellCustomCmdbMapping
}

// GetCherwellCustomCmdbMappingOk returns a tuple with the CherwellCustomCmdbMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellCustomCmdbMappingOk() (*string, bool) {
	if o == nil || o.CherwellCustomCmdbMapping == nil {
		return nil, false
	}
	return o.CherwellCustomCmdbMapping, true
}

// HasCherwellCustomCmdbMapping returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellCustomCmdbMapping() bool {
	if o != nil && o.CherwellCustomCmdbMapping != nil {
		return true
	}

	return false
}

// SetCherwellCustomCmdbMapping gets a reference to the given string and assigns it to the CherwellCustomCmdbMapping field.
func (o *IntegrationCherwellConfig) SetCherwellCustomCmdbMapping(v string) {
	o.CherwellCustomCmdbMapping = &v
}

// GetCherwellClientKey returns the CherwellClientKey field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellClientKey() string {
	if o == nil || o.CherwellClientKey == nil {
		var ret string
		return ret
	}
	return *o.CherwellClientKey
}

// GetCherwellClientKeyOk returns a tuple with the CherwellClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellClientKeyOk() (*string, bool) {
	if o == nil || o.CherwellClientKey == nil {
		return nil, false
	}
	return o.CherwellClientKey, true
}

// HasCherwellClientKey returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellClientKey() bool {
	if o != nil && o.CherwellClientKey != nil {
		return true
	}

	return false
}

// SetCherwellClientKey gets a reference to the given string and assigns it to the CherwellClientKey field.
func (o *IntegrationCherwellConfig) SetCherwellClientKey(v string) {
	o.CherwellClientKey = &v
}

// GetCherwellCreatedBy returns the CherwellCreatedBy field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellCreatedBy() string {
	if o == nil || o.CherwellCreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CherwellCreatedBy
}

// GetCherwellCreatedByOk returns a tuple with the CherwellCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellCreatedByOk() (*string, bool) {
	if o == nil || o.CherwellCreatedBy == nil {
		return nil, false
	}
	return o.CherwellCreatedBy, true
}

// HasCherwellCreatedBy returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellCreatedBy() bool {
	if o != nil && o.CherwellCreatedBy != nil {
		return true
	}

	return false
}

// SetCherwellCreatedBy gets a reference to the given string and assigns it to the CherwellCreatedBy field.
func (o *IntegrationCherwellConfig) SetCherwellCreatedBy(v string) {
	o.CherwellCreatedBy = &v
}

// GetCherwellStartDate returns the CherwellStartDate field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellStartDate() string {
	if o == nil || o.CherwellStartDate == nil {
		var ret string
		return ret
	}
	return *o.CherwellStartDate
}

// GetCherwellStartDateOk returns a tuple with the CherwellStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellStartDateOk() (*string, bool) {
	if o == nil || o.CherwellStartDate == nil {
		return nil, false
	}
	return o.CherwellStartDate, true
}

// HasCherwellStartDate returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellStartDate() bool {
	if o != nil && o.CherwellStartDate != nil {
		return true
	}

	return false
}

// SetCherwellStartDate gets a reference to the given string and assigns it to the CherwellStartDate field.
func (o *IntegrationCherwellConfig) SetCherwellStartDate(v string) {
	o.CherwellStartDate = &v
}

// GetCherwellEndDate returns the CherwellEndDate field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellEndDate() string {
	if o == nil || o.CherwellEndDate == nil {
		var ret string
		return ret
	}
	return *o.CherwellEndDate
}

// GetCherwellEndDateOk returns a tuple with the CherwellEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellEndDateOk() (*string, bool) {
	if o == nil || o.CherwellEndDate == nil {
		return nil, false
	}
	return o.CherwellEndDate, true
}

// HasCherwellEndDate returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellEndDate() bool {
	if o != nil && o.CherwellEndDate != nil {
		return true
	}

	return false
}

// SetCherwellEndDate gets a reference to the given string and assigns it to the CherwellEndDate field.
func (o *IntegrationCherwellConfig) SetCherwellEndDate(v string) {
	o.CherwellEndDate = &v
}

// GetCherwellIgnoreSSLErrors returns the CherwellIgnoreSSLErrors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationCherwellConfig) GetCherwellIgnoreSSLErrors() string {
	if o == nil || o.CherwellIgnoreSSLErrors.Get() == nil {
		var ret string
		return ret
	}
	return *o.CherwellIgnoreSSLErrors.Get()
}

// GetCherwellIgnoreSSLErrorsOk returns a tuple with the CherwellIgnoreSSLErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationCherwellConfig) GetCherwellIgnoreSSLErrorsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CherwellIgnoreSSLErrors.Get(), o.CherwellIgnoreSSLErrors.IsSet()
}

// HasCherwellIgnoreSSLErrors returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellIgnoreSSLErrors() bool {
	if o != nil && o.CherwellIgnoreSSLErrors.IsSet() {
		return true
	}

	return false
}

// SetCherwellIgnoreSSLErrors gets a reference to the given NullableString and assigns it to the CherwellIgnoreSSLErrors field.
func (o *IntegrationCherwellConfig) SetCherwellIgnoreSSLErrors(v string) {
	o.CherwellIgnoreSSLErrors.Set(&v)
}
// SetCherwellIgnoreSSLErrorsNil sets the value for CherwellIgnoreSSLErrors to be an explicit nil
func (o *IntegrationCherwellConfig) SetCherwellIgnoreSSLErrorsNil() {
	o.CherwellIgnoreSSLErrors.Set(nil)
}

// UnsetCherwellIgnoreSSLErrors ensures that no value is present for CherwellIgnoreSSLErrors, not even an explicit nil
func (o *IntegrationCherwellConfig) UnsetCherwellIgnoreSSLErrors() {
	o.CherwellIgnoreSSLErrors.Unset()
}

// GetCherwellBusinessObject returns the CherwellBusinessObject field value if set, zero value otherwise.
func (o *IntegrationCherwellConfig) GetCherwellBusinessObject() string {
	if o == nil || o.CherwellBusinessObject == nil {
		var ret string
		return ret
	}
	return *o.CherwellBusinessObject
}

// GetCherwellBusinessObjectOk returns a tuple with the CherwellBusinessObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCherwellConfig) GetCherwellBusinessObjectOk() (*string, bool) {
	if o == nil || o.CherwellBusinessObject == nil {
		return nil, false
	}
	return o.CherwellBusinessObject, true
}

// HasCherwellBusinessObject returns a boolean if a field has been set.
func (o *IntegrationCherwellConfig) HasCherwellBusinessObject() bool {
	if o != nil && o.CherwellBusinessObject != nil {
		return true
	}

	return false
}

// SetCherwellBusinessObject gets a reference to the given string and assigns it to the CherwellBusinessObject field.
func (o *IntegrationCherwellConfig) SetCherwellBusinessObject(v string) {
	o.CherwellBusinessObject = &v
}

func (o IntegrationCherwellConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CherwellCustomCmdbMapping != nil {
		toSerialize["cherwellCustomCmdbMapping"] = o.CherwellCustomCmdbMapping
	}
	if o.CherwellClientKey != nil {
		toSerialize["cherwellClientKey"] = o.CherwellClientKey
	}
	if o.CherwellCreatedBy != nil {
		toSerialize["cherwellCreatedBy"] = o.CherwellCreatedBy
	}
	if o.CherwellStartDate != nil {
		toSerialize["cherwellStartDate"] = o.CherwellStartDate
	}
	if o.CherwellEndDate != nil {
		toSerialize["cherwellEndDate"] = o.CherwellEndDate
	}
	if o.CherwellIgnoreSSLErrors.IsSet() {
		toSerialize["cherwellIgnoreSSLErrors"] = o.CherwellIgnoreSSLErrors.Get()
	}
	if o.CherwellBusinessObject != nil {
		toSerialize["cherwellBusinessObject"] = o.CherwellBusinessObject
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationCherwellConfig struct {
	value *IntegrationCherwellConfig
	isSet bool
}

func (v NullableIntegrationCherwellConfig) Get() *IntegrationCherwellConfig {
	return v.value
}

func (v *NullableIntegrationCherwellConfig) Set(val *IntegrationCherwellConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationCherwellConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationCherwellConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationCherwellConfig(val *IntegrationCherwellConfig) *NullableIntegrationCherwellConfig {
	return &NullableIntegrationCherwellConfig{value: val, isSet: true}
}

func (v NullableIntegrationCherwellConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationCherwellConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


