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

// GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions struct for GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions
type GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	TermOptions *map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions `json:"termOptions,omitempty"`
}

// NewGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions instantiates a new GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions() *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions {
	this := GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions{}
	return &this
}

// NewGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptionsWithDefaults instantiates a new GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptionsWithDefaults() *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions {
	this := GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) SetName(v string) {
	o.Name = &v
}

// GetTermOptions returns the TermOptions field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetTermOptions() map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions {
	if o == nil || o.TermOptions == nil {
		var ret map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions
		return ret
	}
	return *o.TermOptions
}

// GetTermOptionsOk returns a tuple with the TermOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) GetTermOptionsOk() (*map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions, bool) {
	if o == nil || o.TermOptions == nil {
		return nil, false
	}
	return o.TermOptions, true
}

// HasTermOptions returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) HasTermOptions() bool {
	if o != nil && o.TermOptions != nil {
		return true
	}

	return false
}

// SetTermOptions gets a reference to the given map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions and assigns it to the TermOptions field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) SetTermOptions(v map[string]GuidanceAzureReservationsConfigServicesAzureVmsTermOptions) {
	o.TermOptions = &v
}

func (o GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TermOptions != nil {
		toSerialize["termOptions"] = o.TermOptions
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions struct {
	value *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions
	isSet bool
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) Get() *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions {
	return v.value
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) Set(val *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions(val *GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) *NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions {
	return &NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions{value: val, isSet: true}
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


