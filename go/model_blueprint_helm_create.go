/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlueprintHelmCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintHelmCreate{}

// BlueprintHelmCreate struct for BlueprintHelmCreate
type BlueprintHelmCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	Helm BlueprintHelmCreateHelm `json:"helm"`
}

// NewBlueprintHelmCreate instantiates a new BlueprintHelmCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintHelmCreate(name string, type_ string, helm BlueprintHelmCreateHelm) *BlueprintHelmCreate {
	this := BlueprintHelmCreate{}
	this.Name = name
	this.Type = type_
	this.Helm = helm
	return &this
}

// NewBlueprintHelmCreateWithDefaults instantiates a new BlueprintHelmCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintHelmCreateWithDefaults() *BlueprintHelmCreate {
	this := BlueprintHelmCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintHelmCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintHelmCreate) SetName(v string) {
	o.Name = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BlueprintHelmCreate) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreate) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BlueprintHelmCreate) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BlueprintHelmCreate) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value
func (o *BlueprintHelmCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintHelmCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintHelmCreate) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintHelmCreate) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BlueprintHelmCreate) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintHelmCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetHelm returns the Helm field value
func (o *BlueprintHelmCreate) GetHelm() BlueprintHelmCreateHelm {
	if o == nil {
		var ret BlueprintHelmCreateHelm
		return ret
	}

	return o.Helm
}

// GetHelmOk returns a tuple with the Helm field value
// and a boolean to check if the value has been set.
func (o *BlueprintHelmCreate) GetHelmOk() (*BlueprintHelmCreateHelm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Helm, true
}

// SetHelm sets field value
func (o *BlueprintHelmCreate) SetHelm(v BlueprintHelmCreateHelm) {
	o.Helm = v
}

func (o BlueprintHelmCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintHelmCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["type"] = o.Type
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["helm"] = o.Helm
	return toSerialize, nil
}

type NullableBlueprintHelmCreate struct {
	value *BlueprintHelmCreate
	isSet bool
}

func (v NullableBlueprintHelmCreate) Get() *BlueprintHelmCreate {
	return v.value
}

func (v *NullableBlueprintHelmCreate) Set(val *BlueprintHelmCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintHelmCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintHelmCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintHelmCreate(val *BlueprintHelmCreate) *NullableBlueprintHelmCreate {
	return &NullableBlueprintHelmCreate{value: val, isSet: true}
}

func (v NullableBlueprintHelmCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintHelmCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

