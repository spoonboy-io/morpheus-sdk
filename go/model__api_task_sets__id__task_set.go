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

// ApiTaskSetsIdTaskSet struct for ApiTaskSetsIdTaskSet
type ApiTaskSetsIdTaskSet struct {
	// A unique name for the workflow
	Name *string `json:"name,omitempty"`
	// A description of the workflow
	Description *string `json:"description,omitempty"`
	// An array of Category labels for filtering
	Labels *[]string `json:"labels,omitempty"`
	// Workflow type
	Type *string `json:"type,omitempty"`
	// List of option type IDs for use with operational workflow configuration.
	OptionTypes *[]int64 `json:"optionTypes,omitempty"`
	Tasks *ApiTaskSetsTaskSetTasks `json:"tasks,omitempty"`
}

// NewApiTaskSetsIdTaskSet instantiates a new ApiTaskSetsIdTaskSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTaskSetsIdTaskSet() *ApiTaskSetsIdTaskSet {
	this := ApiTaskSetsIdTaskSet{}
	var type_ string = "provision"
	this.Type = &type_
	return &this
}

// NewApiTaskSetsIdTaskSetWithDefaults instantiates a new ApiTaskSetsIdTaskSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTaskSetsIdTaskSetWithDefaults() *ApiTaskSetsIdTaskSet {
	this := ApiTaskSetsIdTaskSet{}
	var type_ string = "provision"
	this.Type = &type_
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiTaskSetsIdTaskSet) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiTaskSetsIdTaskSet) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ApiTaskSetsIdTaskSet) SetLabels(v []string) {
	o.Labels = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiTaskSetsIdTaskSet) SetType(v string) {
	o.Type = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetOptionTypes() []int64 {
	if o == nil || o.OptionTypes == nil {
		var ret []int64
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetOptionTypesOk() (*[]int64, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []int64 and assigns it to the OptionTypes field.
func (o *ApiTaskSetsIdTaskSet) SetOptionTypes(v []int64) {
	o.OptionTypes = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ApiTaskSetsIdTaskSet) GetTasks() ApiTaskSetsTaskSetTasks {
	if o == nil || o.Tasks == nil {
		var ret ApiTaskSetsTaskSetTasks
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTaskSetsIdTaskSet) GetTasksOk() (*ApiTaskSetsTaskSetTasks, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ApiTaskSetsIdTaskSet) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given ApiTaskSetsTaskSetTasks and assigns it to the Tasks field.
func (o *ApiTaskSetsIdTaskSet) SetTasks(v ApiTaskSetsTaskSetTasks) {
	o.Tasks = &v
}

func (o ApiTaskSetsIdTaskSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableApiTaskSetsIdTaskSet struct {
	value *ApiTaskSetsIdTaskSet
	isSet bool
}

func (v NullableApiTaskSetsIdTaskSet) Get() *ApiTaskSetsIdTaskSet {
	return v.value
}

func (v *NullableApiTaskSetsIdTaskSet) Set(val *ApiTaskSetsIdTaskSet) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTaskSetsIdTaskSet) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTaskSetsIdTaskSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTaskSetsIdTaskSet(val *ApiTaskSetsIdTaskSet) *NullableApiTaskSetsIdTaskSet {
	return &NullableApiTaskSetsIdTaskSet{value: val, isSet: true}
}

func (v NullableApiTaskSetsIdTaskSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTaskSetsIdTaskSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

