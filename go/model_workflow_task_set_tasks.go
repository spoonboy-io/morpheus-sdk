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

// WorkflowTaskSetTasks struct for WorkflowTaskSetTasks
type WorkflowTaskSetTasks struct {
	Id *int64 `json:"id,omitempty"`
	TaskPhase *string `json:"taskPhase,omitempty"`
	TaskOrder *int64 `json:"taskOrder,omitempty"`
	Task *WorkflowTask `json:"task,omitempty"`
}

// NewWorkflowTaskSetTasks instantiates a new WorkflowTaskSetTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskSetTasks() *WorkflowTaskSetTasks {
	this := WorkflowTaskSetTasks{}
	return &this
}

// NewWorkflowTaskSetTasksWithDefaults instantiates a new WorkflowTaskSetTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskSetTasksWithDefaults() *WorkflowTaskSetTasks {
	this := WorkflowTaskSetTasks{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowTaskSetTasks) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSetTasks) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowTaskSetTasks) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WorkflowTaskSetTasks) SetId(v int64) {
	o.Id = &v
}

// GetTaskPhase returns the TaskPhase field value if set, zero value otherwise.
func (o *WorkflowTaskSetTasks) GetTaskPhase() string {
	if o == nil || o.TaskPhase == nil {
		var ret string
		return ret
	}
	return *o.TaskPhase
}

// GetTaskPhaseOk returns a tuple with the TaskPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSetTasks) GetTaskPhaseOk() (*string, bool) {
	if o == nil || o.TaskPhase == nil {
		return nil, false
	}
	return o.TaskPhase, true
}

// HasTaskPhase returns a boolean if a field has been set.
func (o *WorkflowTaskSetTasks) HasTaskPhase() bool {
	if o != nil && o.TaskPhase != nil {
		return true
	}

	return false
}

// SetTaskPhase gets a reference to the given string and assigns it to the TaskPhase field.
func (o *WorkflowTaskSetTasks) SetTaskPhase(v string) {
	o.TaskPhase = &v
}

// GetTaskOrder returns the TaskOrder field value if set, zero value otherwise.
func (o *WorkflowTaskSetTasks) GetTaskOrder() int64 {
	if o == nil || o.TaskOrder == nil {
		var ret int64
		return ret
	}
	return *o.TaskOrder
}

// GetTaskOrderOk returns a tuple with the TaskOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSetTasks) GetTaskOrderOk() (*int64, bool) {
	if o == nil || o.TaskOrder == nil {
		return nil, false
	}
	return o.TaskOrder, true
}

// HasTaskOrder returns a boolean if a field has been set.
func (o *WorkflowTaskSetTasks) HasTaskOrder() bool {
	if o != nil && o.TaskOrder != nil {
		return true
	}

	return false
}

// SetTaskOrder gets a reference to the given int64 and assigns it to the TaskOrder field.
func (o *WorkflowTaskSetTasks) SetTaskOrder(v int64) {
	o.TaskOrder = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *WorkflowTaskSetTasks) GetTask() WorkflowTask {
	if o == nil || o.Task == nil {
		var ret WorkflowTask
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskSetTasks) GetTaskOk() (*WorkflowTask, bool) {
	if o == nil || o.Task == nil {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *WorkflowTaskSetTasks) HasTask() bool {
	if o != nil && o.Task != nil {
		return true
	}

	return false
}

// SetTask gets a reference to the given WorkflowTask and assigns it to the Task field.
func (o *WorkflowTaskSetTasks) SetTask(v WorkflowTask) {
	o.Task = &v
}

func (o WorkflowTaskSetTasks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TaskPhase != nil {
		toSerialize["taskPhase"] = o.TaskPhase
	}
	if o.TaskOrder != nil {
		toSerialize["taskOrder"] = o.TaskOrder
	}
	if o.Task != nil {
		toSerialize["task"] = o.Task
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowTaskSetTasks struct {
	value *WorkflowTaskSetTasks
	isSet bool
}

func (v NullableWorkflowTaskSetTasks) Get() *WorkflowTaskSetTasks {
	return v.value
}

func (v *NullableWorkflowTaskSetTasks) Set(val *WorkflowTaskSetTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskSetTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskSetTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskSetTasks(val *WorkflowTaskSetTasks) *NullableWorkflowTaskSetTasks {
	return &NullableWorkflowTaskSetTasks{value: val, isSet: true}
}

func (v NullableWorkflowTaskSetTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskSetTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

