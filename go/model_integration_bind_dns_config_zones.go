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

// IntegrationBindDNSConfigZones struct for IntegrationBindDNSConfigZones
type IntegrationBindDNSConfigZones struct {
	ZoneName *string `json:"zoneName,omitempty"`
	ReverseZone *string `json:"reverseZone,omitempty"`
}

// NewIntegrationBindDNSConfigZones instantiates a new IntegrationBindDNSConfigZones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationBindDNSConfigZones() *IntegrationBindDNSConfigZones {
	this := IntegrationBindDNSConfigZones{}
	return &this
}

// NewIntegrationBindDNSConfigZonesWithDefaults instantiates a new IntegrationBindDNSConfigZones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationBindDNSConfigZonesWithDefaults() *IntegrationBindDNSConfigZones {
	this := IntegrationBindDNSConfigZones{}
	return &this
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *IntegrationBindDNSConfigZones) GetZoneName() string {
	if o == nil || o.ZoneName == nil {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationBindDNSConfigZones) GetZoneNameOk() (*string, bool) {
	if o == nil || o.ZoneName == nil {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *IntegrationBindDNSConfigZones) HasZoneName() bool {
	if o != nil && o.ZoneName != nil {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *IntegrationBindDNSConfigZones) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetReverseZone returns the ReverseZone field value if set, zero value otherwise.
func (o *IntegrationBindDNSConfigZones) GetReverseZone() string {
	if o == nil || o.ReverseZone == nil {
		var ret string
		return ret
	}
	return *o.ReverseZone
}

// GetReverseZoneOk returns a tuple with the ReverseZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationBindDNSConfigZones) GetReverseZoneOk() (*string, bool) {
	if o == nil || o.ReverseZone == nil {
		return nil, false
	}
	return o.ReverseZone, true
}

// HasReverseZone returns a boolean if a field has been set.
func (o *IntegrationBindDNSConfigZones) HasReverseZone() bool {
	if o != nil && o.ReverseZone != nil {
		return true
	}

	return false
}

// SetReverseZone gets a reference to the given string and assigns it to the ReverseZone field.
func (o *IntegrationBindDNSConfigZones) SetReverseZone(v string) {
	o.ReverseZone = &v
}

func (o IntegrationBindDNSConfigZones) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneName != nil {
		toSerialize["zoneName"] = o.ZoneName
	}
	if o.ReverseZone != nil {
		toSerialize["reverseZone"] = o.ReverseZone
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationBindDNSConfigZones struct {
	value *IntegrationBindDNSConfigZones
	isSet bool
}

func (v NullableIntegrationBindDNSConfigZones) Get() *IntegrationBindDNSConfigZones {
	return v.value
}

func (v *NullableIntegrationBindDNSConfigZones) Set(val *IntegrationBindDNSConfigZones) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationBindDNSConfigZones) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationBindDNSConfigZones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationBindDNSConfigZones(val *IntegrationBindDNSConfigZones) *NullableIntegrationBindDNSConfigZones {
	return &NullableIntegrationBindDNSConfigZones{value: val, isSet: true}
}

func (v NullableIntegrationBindDNSConfigZones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationBindDNSConfigZones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


