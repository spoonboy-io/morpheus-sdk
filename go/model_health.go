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

// Health struct for Health
type Health struct {
	Success *bool `json:"success,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
	ApplianceUrl *string `json:"applianceUrl,omitempty"`
	BuildVersion *string `json:"buildVersion,omitempty"`
	SetupNeeded *bool `json:"setupNeeded,omitempty"`
	Date *time.Time `json:"date,omitempty"`
	Cpu *HealthCpu `json:"cpu,omitempty"`
	Memory *HealthMemory `json:"memory,omitempty"`
	Threads *HealthThreads `json:"threads,omitempty"`
	Database *HealthDatabase `json:"database,omitempty"`
	Elastic *HealthElastic `json:"elastic,omitempty"`
	Rabbit *HealthRabbit `json:"rabbit,omitempty"`
}

// NewHealth instantiates a new Health object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealth() *Health {
	this := Health{}
	return &this
}

// NewHealthWithDefaults instantiates a new Health object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthWithDefaults() *Health {
	this := Health{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *Health) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *Health) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *Health) SetSuccess(v bool) {
	o.Success = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *Health) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *Health) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *Health) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetApplianceUrl returns the ApplianceUrl field value if set, zero value otherwise.
func (o *Health) GetApplianceUrl() string {
	if o == nil || o.ApplianceUrl == nil {
		var ret string
		return ret
	}
	return *o.ApplianceUrl
}

// GetApplianceUrlOk returns a tuple with the ApplianceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetApplianceUrlOk() (*string, bool) {
	if o == nil || o.ApplianceUrl == nil {
		return nil, false
	}
	return o.ApplianceUrl, true
}

// HasApplianceUrl returns a boolean if a field has been set.
func (o *Health) HasApplianceUrl() bool {
	if o != nil && o.ApplianceUrl != nil {
		return true
	}

	return false
}

// SetApplianceUrl gets a reference to the given string and assigns it to the ApplianceUrl field.
func (o *Health) SetApplianceUrl(v string) {
	o.ApplianceUrl = &v
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *Health) GetBuildVersion() string {
	if o == nil || o.BuildVersion == nil {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetBuildVersionOk() (*string, bool) {
	if o == nil || o.BuildVersion == nil {
		return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *Health) HasBuildVersion() bool {
	if o != nil && o.BuildVersion != nil {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *Health) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetSetupNeeded returns the SetupNeeded field value if set, zero value otherwise.
func (o *Health) GetSetupNeeded() bool {
	if o == nil || o.SetupNeeded == nil {
		var ret bool
		return ret
	}
	return *o.SetupNeeded
}

// GetSetupNeededOk returns a tuple with the SetupNeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetSetupNeededOk() (*bool, bool) {
	if o == nil || o.SetupNeeded == nil {
		return nil, false
	}
	return o.SetupNeeded, true
}

// HasSetupNeeded returns a boolean if a field has been set.
func (o *Health) HasSetupNeeded() bool {
	if o != nil && o.SetupNeeded != nil {
		return true
	}

	return false
}

// SetSetupNeeded gets a reference to the given bool and assigns it to the SetupNeeded field.
func (o *Health) SetSetupNeeded(v bool) {
	o.SetupNeeded = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Health) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Health) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *Health) SetDate(v time.Time) {
	o.Date = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Health) GetCpu() HealthCpu {
	if o == nil || o.Cpu == nil {
		var ret HealthCpu
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetCpuOk() (*HealthCpu, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Health) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given HealthCpu and assigns it to the Cpu field.
func (o *Health) SetCpu(v HealthCpu) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Health) GetMemory() HealthMemory {
	if o == nil || o.Memory == nil {
		var ret HealthMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetMemoryOk() (*HealthMemory, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Health) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given HealthMemory and assigns it to the Memory field.
func (o *Health) SetMemory(v HealthMemory) {
	o.Memory = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *Health) GetThreads() HealthThreads {
	if o == nil || o.Threads == nil {
		var ret HealthThreads
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetThreadsOk() (*HealthThreads, bool) {
	if o == nil || o.Threads == nil {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *Health) HasThreads() bool {
	if o != nil && o.Threads != nil {
		return true
	}

	return false
}

// SetThreads gets a reference to the given HealthThreads and assigns it to the Threads field.
func (o *Health) SetThreads(v HealthThreads) {
	o.Threads = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *Health) GetDatabase() HealthDatabase {
	if o == nil || o.Database == nil {
		var ret HealthDatabase
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetDatabaseOk() (*HealthDatabase, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *Health) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given HealthDatabase and assigns it to the Database field.
func (o *Health) SetDatabase(v HealthDatabase) {
	o.Database = &v
}

// GetElastic returns the Elastic field value if set, zero value otherwise.
func (o *Health) GetElastic() HealthElastic {
	if o == nil || o.Elastic == nil {
		var ret HealthElastic
		return ret
	}
	return *o.Elastic
}

// GetElasticOk returns a tuple with the Elastic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetElasticOk() (*HealthElastic, bool) {
	if o == nil || o.Elastic == nil {
		return nil, false
	}
	return o.Elastic, true
}

// HasElastic returns a boolean if a field has been set.
func (o *Health) HasElastic() bool {
	if o != nil && o.Elastic != nil {
		return true
	}

	return false
}

// SetElastic gets a reference to the given HealthElastic and assigns it to the Elastic field.
func (o *Health) SetElastic(v HealthElastic) {
	o.Elastic = &v
}

// GetRabbit returns the Rabbit field value if set, zero value otherwise.
func (o *Health) GetRabbit() HealthRabbit {
	if o == nil || o.Rabbit == nil {
		var ret HealthRabbit
		return ret
	}
	return *o.Rabbit
}

// GetRabbitOk returns a tuple with the Rabbit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Health) GetRabbitOk() (*HealthRabbit, bool) {
	if o == nil || o.Rabbit == nil {
		return nil, false
	}
	return o.Rabbit, true
}

// HasRabbit returns a boolean if a field has been set.
func (o *Health) HasRabbit() bool {
	if o != nil && o.Rabbit != nil {
		return true
	}

	return false
}

// SetRabbit gets a reference to the given HealthRabbit and assigns it to the Rabbit field.
func (o *Health) SetRabbit(v HealthRabbit) {
	o.Rabbit = &v
}

func (o Health) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.StatusMessage != nil {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if o.ApplianceUrl != nil {
		toSerialize["applianceUrl"] = o.ApplianceUrl
	}
	if o.BuildVersion != nil {
		toSerialize["buildVersion"] = o.BuildVersion
	}
	if o.SetupNeeded != nil {
		toSerialize["setupNeeded"] = o.SetupNeeded
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.Elastic != nil {
		toSerialize["elastic"] = o.Elastic
	}
	if o.Rabbit != nil {
		toSerialize["rabbit"] = o.Rabbit
	}
	return json.Marshal(toSerialize)
}

type NullableHealth struct {
	value *Health
	isSet bool
}

func (v NullableHealth) Get() *Health {
	return v.value
}

func (v *NullableHealth) Set(val *Health) {
	v.value = val
	v.isSet = true
}

func (v NullableHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealth(val *Health) *NullableHealth {
	return &NullableHealth{value: val, isSet: true}
}

func (v NullableHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


