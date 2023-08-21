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

// BlueprintARMCreate struct for BlueprintARMCreate
type BlueprintARMCreate struct {
	// A name for the blueprint
	Name string `json:"name"`
	// Path to display image. Defaults to an internal Morpheus image.
	Image *string `json:"image,omitempty"`
	// Blueprint Type
	Type string `json:"type"`
	// Array of label strings, can be used for filtering.
	Labels []string `json:"labels,omitempty"`
	Arm BlueprintARMCreateArm `json:"arm"`
}

// NewBlueprintARMCreate instantiates a new BlueprintARMCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintARMCreate(name string, type_ string, arm BlueprintARMCreateArm, ) *BlueprintARMCreate {
	this := BlueprintARMCreate{}
	this.Name = name
	this.Type = type_
	this.Arm = arm
	return &this
}

// NewBlueprintARMCreateWithDefaults instantiates a new BlueprintARMCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintARMCreateWithDefaults() *BlueprintARMCreate {
	this := BlueprintARMCreate{}
	return &this
}

// GetName returns the Name field value
func (o *BlueprintARMCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlueprintARMCreate) SetName(v string) {
	o.Name = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BlueprintARMCreate) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreate) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BlueprintARMCreate) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BlueprintARMCreate) SetImage(v string) {
	o.Image = &v
}

// GetType returns the Type field value
func (o *BlueprintARMCreate) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreate) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BlueprintARMCreate) SetType(v string) {
	o.Type = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintARMCreate) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintARMCreate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BlueprintARMCreate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BlueprintARMCreate) SetLabels(v []string) {
	o.Labels = v
}

// GetArm returns the Arm field value
func (o *BlueprintARMCreate) GetArm() BlueprintARMCreateArm {
	if o == nil  {
		var ret BlueprintARMCreateArm
		return ret
	}

	return o.Arm
}

// GetArmOk returns a tuple with the Arm field value
// and a boolean to check if the value has been set.
func (o *BlueprintARMCreate) GetArmOk() (*BlueprintARMCreateArm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Arm, true
}

// SetArm sets field value
func (o *BlueprintARMCreate) SetArm(v BlueprintARMCreateArm) {
	o.Arm = v
}

func (o BlueprintARMCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["arm"] = o.Arm
	}
	return json.Marshal(toSerialize)
}

type NullableBlueprintARMCreate struct {
	value *BlueprintARMCreate
	isSet bool
}

func (v NullableBlueprintARMCreate) Get() *BlueprintARMCreate {
	return v.value
}

func (v *NullableBlueprintARMCreate) Set(val *BlueprintARMCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintARMCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintARMCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintARMCreate(val *BlueprintARMCreate) *NullableBlueprintARMCreate {
	return &NullableBlueprintARMCreate{value: val, isSet: true}
}

func (v NullableBlueprintARMCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintARMCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


