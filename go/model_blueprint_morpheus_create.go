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

// BlueprintMorpheusCreate struct for BlueprintMorpheusCreate
type BlueprintMorpheusCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	// Tier definitions - Create in UI to view a baseline for object
	Tiers map[string]interface{} `json:"tiers"`
}

// NewBlueprintMorpheusCreate instantiates a new BlueprintMorpheusCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintMorpheusCreate(name string, type_ string, tiers map[string]interface{}, ) *BlueprintMorpheusCreate {
	this := BlueprintMorpheusCreate{}
	this.Name = name
	this.Type = type_
	this.Tiers = tiers
	return &this
}

// NewBlueprintMorpheusCreateWithDefaults instantiates a new BlueprintMorpheusCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintMorpheusCreateWithDefaults() *BlueprintMorpheusCreate {
	this := BlueprintMorpheusCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintMorpheusCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintMorpheusCreate) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BlueprintMorpheusCreate) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintMorpheusCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintMorpheusCreate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintMorpheusCreate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BlueprintMorpheusCreate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintMorpheusCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetTiers returns the Tiers field value
func (o *BlueprintMorpheusCreate) GetTiers() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value
// and a boolean to check if the value has been set.
func (o *BlueprintMorpheusCreate) GetTiersOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tiers, true
}

// SetTiers sets field value
func (o *BlueprintMorpheusCreate) SetTiers(v map[string]interface{}) {
	o.Tiers = v
}

func (o BlueprintMorpheusCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["tiers"] = o.Tiers
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintMorpheusCreate struct {
	value *BlueprintMorpheusCreate
	isSet bool
}

func (v NullableBlueprintMorpheusCreate) Get() *BlueprintMorpheusCreate {
	return v.value
}

func (v *NullableBlueprintMorpheusCreate) Set(val *BlueprintMorpheusCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintMorpheusCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintMorpheusCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintMorpheusCreate(val *BlueprintMorpheusCreate) *NullableBlueprintMorpheusCreate {
	return &NullableBlueprintMorpheusCreate{value: val, isSet: true}
}

func (v NullableBlueprintMorpheusCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintMorpheusCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


