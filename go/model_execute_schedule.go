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

// ExecuteSchedule struct for ExecuteSchedule
type ExecuteSchedule struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Visibility NullableString `json:"visibility,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	ScheduleType *string `json:"scheduleType,omitempty"`
	ScheduleTimezone NullableString `json:"scheduleTimezone,omitempty"`
	Cron *string `json:"cron,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewExecuteSchedule instantiates a new ExecuteSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteSchedule() *ExecuteSchedule {
	this := ExecuteSchedule{}
	return &this
}

// NewExecuteScheduleWithDefaults instantiates a new ExecuteSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteScheduleWithDefaults() *ExecuteSchedule {
	this := ExecuteSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ExecuteSchedule) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExecuteSchedule) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecuteSchedule) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecuteSchedule) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExecuteSchedule) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExecuteSchedule) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExecuteSchedule) UnsetDescription() {
	o.Description.Unset()
}

// GetVisibility returns the Visibility field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecuteSchedule) GetVisibility() string {
	if o == nil || o.Visibility.Get() == nil {
		var ret string
		return ret
	}
	return *o.Visibility.Get()
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecuteSchedule) GetVisibilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Visibility.Get(), o.Visibility.IsSet()
}

// HasVisibility returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasVisibility() bool {
	if o != nil && o.Visibility.IsSet() {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given NullableString and assigns it to the Visibility field.
func (o *ExecuteSchedule) SetVisibility(v string) {
	o.Visibility.Set(&v)
}
// SetVisibilityNil sets the value for Visibility to be an explicit nil
func (o *ExecuteSchedule) SetVisibilityNil() {
	o.Visibility.Set(nil)
}

// UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
func (o *ExecuteSchedule) UnsetVisibility() {
	o.Visibility.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExecuteSchedule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetScheduleType returns the ScheduleType field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetScheduleType() string {
	if o == nil || o.ScheduleType == nil {
		var ret string
		return ret
	}
	return *o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetScheduleTypeOk() (*string, bool) {
	if o == nil || o.ScheduleType == nil {
		return nil, false
	}
	return o.ScheduleType, true
}

// HasScheduleType returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasScheduleType() bool {
	if o != nil && o.ScheduleType != nil {
		return true
	}

	return false
}

// SetScheduleType gets a reference to the given string and assigns it to the ScheduleType field.
func (o *ExecuteSchedule) SetScheduleType(v string) {
	o.ScheduleType = &v
}

// GetScheduleTimezone returns the ScheduleTimezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecuteSchedule) GetScheduleTimezone() string {
	if o == nil || o.ScheduleTimezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.ScheduleTimezone.Get()
}

// GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecuteSchedule) GetScheduleTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ScheduleTimezone.Get(), o.ScheduleTimezone.IsSet()
}

// HasScheduleTimezone returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasScheduleTimezone() bool {
	if o != nil && o.ScheduleTimezone.IsSet() {
		return true
	}

	return false
}

// SetScheduleTimezone gets a reference to the given NullableString and assigns it to the ScheduleTimezone field.
func (o *ExecuteSchedule) SetScheduleTimezone(v string) {
	o.ScheduleTimezone.Set(&v)
}
// SetScheduleTimezoneNil sets the value for ScheduleTimezone to be an explicit nil
func (o *ExecuteSchedule) SetScheduleTimezoneNil() {
	o.ScheduleTimezone.Set(nil)
}

// UnsetScheduleTimezone ensures that no value is present for ScheduleTimezone, not even an explicit nil
func (o *ExecuteSchedule) UnsetScheduleTimezone() {
	o.ScheduleTimezone.Unset()
}

// GetCron returns the Cron field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetCron() string {
	if o == nil || o.Cron == nil {
		var ret string
		return ret
	}
	return *o.Cron
}

// GetCronOk returns a tuple with the Cron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetCronOk() (*string, bool) {
	if o == nil || o.Cron == nil {
		return nil, false
	}
	return o.Cron, true
}

// HasCron returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasCron() bool {
	if o != nil && o.Cron != nil {
		return true
	}

	return false
}

// SetCron gets a reference to the given string and assigns it to the Cron field.
func (o *ExecuteSchedule) SetCron(v string) {
	o.Cron = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ExecuteSchedule) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ExecuteSchedule) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteSchedule) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ExecuteSchedule) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ExecuteSchedule) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ExecuteSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Visibility.IsSet() {
		toSerialize["visibility"] = o.Visibility.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ScheduleType != nil {
		toSerialize["scheduleType"] = o.ScheduleType
	}
	if o.ScheduleTimezone.IsSet() {
		toSerialize["scheduleTimezone"] = o.ScheduleTimezone.Get()
	}
	if o.Cron != nil {
		toSerialize["cron"] = o.Cron
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteSchedule struct {
	value *ExecuteSchedule
	isSet bool
}

func (v NullableExecuteSchedule) Get() *ExecuteSchedule {
	return v.value
}

func (v *NullableExecuteSchedule) Set(val *ExecuteSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteSchedule(val *ExecuteSchedule) *NullableExecuteSchedule {
	return &NullableExecuteSchedule{value: val, isSet: true}
}

func (v NullableExecuteSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

