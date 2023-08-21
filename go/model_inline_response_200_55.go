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

// InlineResponse20055 struct for InlineResponse20055
type InlineResponse20055 struct {
	Incident *Incident `json:"incident,omitempty"`
}

// NewInlineResponse20055 instantiates a new InlineResponse20055 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20055() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// NewInlineResponse20055WithDefaults instantiates a new InlineResponse20055 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20055WithDefaults() *InlineResponse20055 {
	this := InlineResponse20055{}
	return &this
}

// GetIncident returns the Incident field value if set, zero value otherwise.
func (o *InlineResponse20055) GetIncident() Incident {
	if o == nil || o.Incident == nil {
		var ret Incident
		return ret
	}
	return *o.Incident
}

// GetIncidentOk returns a tuple with the Incident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055) GetIncidentOk() (*Incident, bool) {
	if o == nil || o.Incident == nil {
		return nil, false
	}
	return o.Incident, true
}

// HasIncident returns a boolean if a field has been set.
func (o *InlineResponse20055) HasIncident() bool {
	if o != nil && o.Incident != nil {
		return true
	}

	return false
}

// SetIncident gets a reference to the given Incident and assigns it to the Incident field.
func (o *InlineResponse20055) SetIncident(v Incident) {
	o.Incident = &v
}

func (o InlineResponse20055) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Incident != nil {
		toSerialize["incident"] = o.Incident
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20055 struct {
	value *InlineResponse20055
	isSet bool
}

func (v NullableInlineResponse20055) Get() *InlineResponse20055 {
	return v.value
}

func (v *NullableInlineResponse20055) Set(val *InlineResponse20055) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20055) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20055) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20055(val *InlineResponse20055) *NullableInlineResponse20055 {
	return &NullableInlineResponse20055{value: val, isSet: true}
}

func (v NullableInlineResponse20055) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20055) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


