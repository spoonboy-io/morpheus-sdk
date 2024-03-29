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

// GuidanceAzureReservationsConfigServices struct for GuidanceAzureReservationsConfigServices
type GuidanceAzureReservationsConfigServices struct {
	AzureVms *GuidanceAzureReservationsConfigServicesAzureVms `json:"azureVms,omitempty"`
}

// NewGuidanceAzureReservationsConfigServices instantiates a new GuidanceAzureReservationsConfigServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceAzureReservationsConfigServices() *GuidanceAzureReservationsConfigServices {
	this := GuidanceAzureReservationsConfigServices{}
	return &this
}

// NewGuidanceAzureReservationsConfigServicesWithDefaults instantiates a new GuidanceAzureReservationsConfigServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceAzureReservationsConfigServicesWithDefaults() *GuidanceAzureReservationsConfigServices {
	this := GuidanceAzureReservationsConfigServices{}
	return &this
}

// GetAzureVms returns the AzureVms field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServices) GetAzureVms() GuidanceAzureReservationsConfigServicesAzureVms {
	if o == nil || o.AzureVms == nil {
		var ret GuidanceAzureReservationsConfigServicesAzureVms
		return ret
	}
	return *o.AzureVms
}

// GetAzureVmsOk returns a tuple with the AzureVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServices) GetAzureVmsOk() (*GuidanceAzureReservationsConfigServicesAzureVms, bool) {
	if o == nil || o.AzureVms == nil {
		return nil, false
	}
	return o.AzureVms, true
}

// HasAzureVms returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServices) HasAzureVms() bool {
	if o != nil && o.AzureVms != nil {
		return true
	}

	return false
}

// SetAzureVms gets a reference to the given GuidanceAzureReservationsConfigServicesAzureVms and assigns it to the AzureVms field.
func (o *GuidanceAzureReservationsConfigServices) SetAzureVms(v GuidanceAzureReservationsConfigServicesAzureVms) {
	o.AzureVms = &v
}

func (o GuidanceAzureReservationsConfigServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureVms != nil {
		toSerialize["azureVms"] = o.AzureVms
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceAzureReservationsConfigServices struct {
	value *GuidanceAzureReservationsConfigServices
	isSet bool
}

func (v NullableGuidanceAzureReservationsConfigServices) Get() *GuidanceAzureReservationsConfigServices {
	return v.value
}

func (v *NullableGuidanceAzureReservationsConfigServices) Set(val *GuidanceAzureReservationsConfigServices) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceAzureReservationsConfigServices) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceAzureReservationsConfigServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceAzureReservationsConfigServices(val *GuidanceAzureReservationsConfigServices) *NullableGuidanceAzureReservationsConfigServices {
	return &NullableGuidanceAzureReservationsConfigServices{value: val, isSet: true}
}

func (v NullableGuidanceAzureReservationsConfigServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceAzureReservationsConfigServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


