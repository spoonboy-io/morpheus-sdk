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

// JobExecution struct for JobExecution
type JobExecution struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Process NullableString `json:"process,omitempty"`
	Job *JobExecutionJob `json:"job,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	ResultData NullableString `json:"resultData,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
}

// NewJobExecution instantiates a new JobExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobExecution() *JobExecution {
	this := JobExecution{}
	return &this
}

// NewJobExecutionWithDefaults instantiates a new JobExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobExecutionWithDefaults() *JobExecution {
	this := JobExecution{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JobExecution) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JobExecution) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *JobExecution) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *JobExecution) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *JobExecution) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *JobExecution) SetName(v string) {
	o.Name = &v
}

// GetProcess returns the Process field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobExecution) GetProcess() string {
	if o == nil || o.Process.Get() == nil {
		var ret string
		return ret
	}
	return *o.Process.Get()
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobExecution) GetProcessOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Process.Get(), o.Process.IsSet()
}

// HasProcess returns a boolean if a field has been set.
func (o *JobExecution) HasProcess() bool {
	if o != nil && o.Process.IsSet() {
		return true
	}

	return false
}

// SetProcess gets a reference to the given NullableString and assigns it to the Process field.
func (o *JobExecution) SetProcess(v string) {
	o.Process.Set(&v)
}
// SetProcessNil sets the value for Process to be an explicit nil
func (o *JobExecution) SetProcessNil() {
	o.Process.Set(nil)
}

// UnsetProcess ensures that no value is present for Process, not even an explicit nil
func (o *JobExecution) UnsetProcess() {
	o.Process.Unset()
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *JobExecution) GetJob() JobExecutionJob {
	if o == nil || o.Job == nil {
		var ret JobExecutionJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetJobOk() (*JobExecutionJob, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *JobExecution) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given JobExecutionJob and assigns it to the Job field.
func (o *JobExecution) SetJob(v JobExecutionJob) {
	o.Job = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobExecution) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobExecution) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *JobExecution) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *JobExecution) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *JobExecution) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *JobExecution) UnsetDescription() {
	o.Description.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *JobExecution) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *JobExecution) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *JobExecution) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *JobExecution) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *JobExecution) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *JobExecution) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *JobExecution) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *JobExecution) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *JobExecution) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *JobExecution) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *JobExecution) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *JobExecution) SetDuration(v int64) {
	o.Duration = &v
}

// GetResultData returns the ResultData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobExecution) GetResultData() string {
	if o == nil || o.ResultData.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResultData.Get()
}

// GetResultDataOk returns a tuple with the ResultData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobExecution) GetResultDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResultData.Get(), o.ResultData.IsSet()
}

// HasResultData returns a boolean if a field has been set.
func (o *JobExecution) HasResultData() bool {
	if o != nil && o.ResultData.IsSet() {
		return true
	}

	return false
}

// SetResultData gets a reference to the given NullableString and assigns it to the ResultData field.
func (o *JobExecution) SetResultData(v string) {
	o.ResultData.Set(&v)
}
// SetResultDataNil sets the value for ResultData to be an explicit nil
func (o *JobExecution) SetResultDataNil() {
	o.ResultData.Set(nil)
}

// UnsetResultData ensures that no value is present for ResultData, not even an explicit nil
func (o *JobExecution) UnsetResultData() {
	o.ResultData.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JobExecution) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobExecution) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JobExecution) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JobExecution) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobExecution) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobExecution) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *JobExecution) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *JobExecution) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *JobExecution) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *JobExecution) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobExecution) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobExecution) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *JobExecution) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *JobExecution) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *JobExecution) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *JobExecution) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

func (o JobExecution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Process.IsSet() {
		toSerialize["process"] = o.Process.Get()
	}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ResultData.IsSet() {
		toSerialize["resultData"] = o.ResultData.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJobExecution struct {
	value *JobExecution
	isSet bool
}

func (v NullableJobExecution) Get() *JobExecution {
	return v.value
}

func (v *NullableJobExecution) Set(val *JobExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableJobExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableJobExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobExecution(val *JobExecution) *NullableJobExecution {
	return &NullableJobExecution{value: val, isSet: true}
}

func (v NullableJobExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


