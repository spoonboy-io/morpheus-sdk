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
	"time"
)

// Workflow struct for Workflow
type Workflow struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Labels []string `json:"labels,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	Platform NullableString `json:"platform,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	AllowCustomConfig *bool `json:"allowCustomConfig,omitempty"`
	Tasks *[]int64 `json:"tasks,omitempty"`
	OptionTypes []WorkflowOptionTypes `json:"optionTypes,omitempty"`
	TaskSetTasks *[]WorkflowTaskSetTasks `json:"taskSetTasks,omitempty"`
}

// NewWorkflow instantiates a new Workflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow() *Workflow {
	this := Workflow{}
	return &this
}

// NewWorkflowWithDefaults instantiates a new Workflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWithDefaults() *Workflow {
	this := Workflow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Workflow) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Workflow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Workflow) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Workflow) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Workflow) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Workflow) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Workflow) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Workflow) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Workflow) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workflow) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workflow) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Workflow) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Workflow) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Workflow) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Workflow) UnsetDescription() {
	o.Description.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workflow) GetLabels() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workflow) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Workflow) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Workflow) SetLabels(v []string) {
	o.Labels = v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Workflow) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Workflow) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Workflow) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Workflow) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Workflow) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Workflow) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Workflow) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Workflow) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *Workflow) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workflow) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workflow) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Workflow) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *Workflow) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Workflow) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Workflow) UnsetPlatform() {
	o.Platform.Unset()
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Workflow) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Workflow) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *Workflow) SetVisibility(v string) {
	o.Visibility = &v
}

// GetAllowCustomConfig returns the AllowCustomConfig field value if set, zero value otherwise.
func (o *Workflow) GetAllowCustomConfig() bool {
	if o == nil || o.AllowCustomConfig == nil {
		var ret bool
		return ret
	}
	return *o.AllowCustomConfig
}

// GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetAllowCustomConfigOk() (*bool, bool) {
	if o == nil || o.AllowCustomConfig == nil {
		return nil, false
	}
	return o.AllowCustomConfig, true
}

// HasAllowCustomConfig returns a boolean if a field has been set.
func (o *Workflow) HasAllowCustomConfig() bool {
	if o != nil && o.AllowCustomConfig != nil {
		return true
	}

	return false
}

// SetAllowCustomConfig gets a reference to the given bool and assigns it to the AllowCustomConfig field.
func (o *Workflow) SetAllowCustomConfig(v bool) {
	o.AllowCustomConfig = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *Workflow) GetTasks() []int64 {
	if o == nil || o.Tasks == nil {
		var ret []int64
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetTasksOk() (*[]int64, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *Workflow) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []int64 and assigns it to the Tasks field.
func (o *Workflow) SetTasks(v []int64) {
	o.Tasks = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Workflow) GetOptionTypes() []WorkflowOptionTypes {
	if o == nil  {
		var ret []WorkflowOptionTypes
		return ret
	}
	return o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Workflow) GetOptionTypesOk() (*[]WorkflowOptionTypes, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return &o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *Workflow) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []WorkflowOptionTypes and assigns it to the OptionTypes field.
func (o *Workflow) SetOptionTypes(v []WorkflowOptionTypes) {
	o.OptionTypes = v
}

// GetTaskSetTasks returns the TaskSetTasks field value if set, zero value otherwise.
func (o *Workflow) GetTaskSetTasks() []WorkflowTaskSetTasks {
	if o == nil || o.TaskSetTasks == nil {
		var ret []WorkflowTaskSetTasks
		return ret
	}
	return *o.TaskSetTasks
}

// GetTaskSetTasksOk returns a tuple with the TaskSetTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetTaskSetTasksOk() (*[]WorkflowTaskSetTasks, bool) {
	if o == nil || o.TaskSetTasks == nil {
		return nil, false
	}
	return o.TaskSetTasks, true
}

// HasTaskSetTasks returns a boolean if a field has been set.
func (o *Workflow) HasTaskSetTasks() bool {
	if o != nil && o.TaskSetTasks != nil {
		return true
	}

	return false
}

// SetTaskSetTasks gets a reference to the given []WorkflowTaskSetTasks and assigns it to the TaskSetTasks field.
func (o *Workflow) SetTaskSetTasks(v []WorkflowTaskSetTasks) {
	o.TaskSetTasks = &v
}

func (o Workflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.AllowCustomConfig != nil {
		toSerialize["allowCustomConfig"] = o.AllowCustomConfig
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.TaskSetTasks != nil {
		toSerialize["taskSetTasks"] = o.TaskSetTasks
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflow struct {
	value *Workflow
	isSet bool
}

func (v NullableWorkflow) Get() *Workflow {
	return v.value
}

func (v *NullableWorkflow) Set(val *Workflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow(val *Workflow) *NullableWorkflow {
	return &NullableWorkflow{value: val, isSet: true}
}

func (v NullableWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

