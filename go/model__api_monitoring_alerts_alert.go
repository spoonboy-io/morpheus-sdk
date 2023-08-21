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

// ApiMonitoringAlertsAlert Payload for creating a new monitoring alert
type ApiMonitoringAlertsAlert struct {
	// Unique name scoped to your account for the alert
	Name string `json:"name"`
	// Duration in minutes of the delay before sending notification(s)
	MinDuration *int32 `json:"minDuration,omitempty"`
	// Severity level threshold for sending notifications.
	MinSeverity *string `json:"minSeverity,omitempty"`
	// Set to false to disable notifications
	Active *bool `json:"active,omitempty"`
	// Trigger for all checks
	AllChecks *bool `json:"allChecks,omitempty"`
	// Trigger for all check groups
	AllGroups *bool `json:"allGroups,omitempty"`
	// Trigger for all monitor apps
	AllApps *bool `json:"allApps,omitempty"`
	Checks *[]int32 `json:"checks,omitempty"`
	Groups *[]int32 `json:"groups,omitempty"`
	Apps *[]int32 `json:"apps,omitempty"`
	Contacts *[]ApiMonitoringAlertsAlertContacts `json:"contacts,omitempty"`
}

// NewApiMonitoringAlertsAlert instantiates a new ApiMonitoringAlertsAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMonitoringAlertsAlert(name string, ) *ApiMonitoringAlertsAlert {
	this := ApiMonitoringAlertsAlert{}
	this.Name = name
	var minDuration int32 = 0
	this.MinDuration = &minDuration
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var active bool = true
	this.Active = &active
	var allChecks bool = false
	this.AllChecks = &allChecks
	var allGroups bool = false
	this.AllGroups = &allGroups
	var allApps bool = false
	this.AllApps = &allApps
	return &this
}

// NewApiMonitoringAlertsAlertWithDefaults instantiates a new ApiMonitoringAlertsAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMonitoringAlertsAlertWithDefaults() *ApiMonitoringAlertsAlert {
	this := ApiMonitoringAlertsAlert{}
	var minDuration int32 = 0
	this.MinDuration = &minDuration
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var active bool = true
	this.Active = &active
	var allChecks bool = false
	this.AllChecks = &allChecks
	var allGroups bool = false
	this.AllGroups = &allGroups
	var allApps bool = false
	this.AllApps = &allApps
	return &this
}

// GetName returns the Name field value
func (o *ApiMonitoringAlertsAlert) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiMonitoringAlertsAlert) SetName(v string) {
	o.Name = v
}

// GetMinDuration returns the MinDuration field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetMinDuration() int32 {
	if o == nil || o.MinDuration == nil {
		var ret int32
		return ret
	}
	return *o.MinDuration
}

// GetMinDurationOk returns a tuple with the MinDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetMinDurationOk() (*int32, bool) {
	if o == nil || o.MinDuration == nil {
		return nil, false
	}
	return o.MinDuration, true
}

// HasMinDuration returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasMinDuration() bool {
	if o != nil && o.MinDuration != nil {
		return true
	}

	return false
}

// SetMinDuration gets a reference to the given int32 and assigns it to the MinDuration field.
func (o *ApiMonitoringAlertsAlert) SetMinDuration(v int32) {
	o.MinDuration = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *ApiMonitoringAlertsAlert) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiMonitoringAlertsAlert) SetActive(v bool) {
	o.Active = &v
}

// GetAllChecks returns the AllChecks field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetAllChecks() bool {
	if o == nil || o.AllChecks == nil {
		var ret bool
		return ret
	}
	return *o.AllChecks
}

// GetAllChecksOk returns a tuple with the AllChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetAllChecksOk() (*bool, bool) {
	if o == nil || o.AllChecks == nil {
		return nil, false
	}
	return o.AllChecks, true
}

// HasAllChecks returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasAllChecks() bool {
	if o != nil && o.AllChecks != nil {
		return true
	}

	return false
}

// SetAllChecks gets a reference to the given bool and assigns it to the AllChecks field.
func (o *ApiMonitoringAlertsAlert) SetAllChecks(v bool) {
	o.AllChecks = &v
}

// GetAllGroups returns the AllGroups field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetAllGroups() bool {
	if o == nil || o.AllGroups == nil {
		var ret bool
		return ret
	}
	return *o.AllGroups
}

// GetAllGroupsOk returns a tuple with the AllGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetAllGroupsOk() (*bool, bool) {
	if o == nil || o.AllGroups == nil {
		return nil, false
	}
	return o.AllGroups, true
}

// HasAllGroups returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasAllGroups() bool {
	if o != nil && o.AllGroups != nil {
		return true
	}

	return false
}

// SetAllGroups gets a reference to the given bool and assigns it to the AllGroups field.
func (o *ApiMonitoringAlertsAlert) SetAllGroups(v bool) {
	o.AllGroups = &v
}

// GetAllApps returns the AllApps field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetAllApps() bool {
	if o == nil || o.AllApps == nil {
		var ret bool
		return ret
	}
	return *o.AllApps
}

// GetAllAppsOk returns a tuple with the AllApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetAllAppsOk() (*bool, bool) {
	if o == nil || o.AllApps == nil {
		return nil, false
	}
	return o.AllApps, true
}

// HasAllApps returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasAllApps() bool {
	if o != nil && o.AllApps != nil {
		return true
	}

	return false
}

// SetAllApps gets a reference to the given bool and assigns it to the AllApps field.
func (o *ApiMonitoringAlertsAlert) SetAllApps(v bool) {
	o.AllApps = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetChecks() []int32 {
	if o == nil || o.Checks == nil {
		var ret []int32
		return ret
	}
	return *o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetChecksOk() (*[]int32, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []int32 and assigns it to the Checks field.
func (o *ApiMonitoringAlertsAlert) SetChecks(v []int32) {
	o.Checks = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetGroups() []int32 {
	if o == nil || o.Groups == nil {
		var ret []int32
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetGroupsOk() (*[]int32, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *ApiMonitoringAlertsAlert) SetGroups(v []int32) {
	o.Groups = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetApps() []int32 {
	if o == nil || o.Apps == nil {
		var ret []int32
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetAppsOk() (*[]int32, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []int32 and assigns it to the Apps field.
func (o *ApiMonitoringAlertsAlert) SetApps(v []int32) {
	o.Apps = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ApiMonitoringAlertsAlert) GetContacts() []ApiMonitoringAlertsAlertContacts {
	if o == nil || o.Contacts == nil {
		var ret []ApiMonitoringAlertsAlertContacts
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMonitoringAlertsAlert) GetContactsOk() (*[]ApiMonitoringAlertsAlertContacts, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ApiMonitoringAlertsAlert) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ApiMonitoringAlertsAlertContacts and assigns it to the Contacts field.
func (o *ApiMonitoringAlertsAlert) SetContacts(v []ApiMonitoringAlertsAlertContacts) {
	o.Contacts = &v
}

func (o ApiMonitoringAlertsAlert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.MinDuration != nil {
		toSerialize["minDuration"] = o.MinDuration
	}
	if o.MinSeverity != nil {
		toSerialize["minSeverity"] = o.MinSeverity
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.AllChecks != nil {
		toSerialize["allChecks"] = o.AllChecks
	}
	if o.AllGroups != nil {
		toSerialize["allGroups"] = o.AllGroups
	}
	if o.AllApps != nil {
		toSerialize["allApps"] = o.AllApps
	}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	return json.Marshal(toSerialize)
}

type NullableApiMonitoringAlertsAlert struct {
	value *ApiMonitoringAlertsAlert
	isSet bool
}

func (v NullableApiMonitoringAlertsAlert) Get() *ApiMonitoringAlertsAlert {
	return v.value
}

func (v *NullableApiMonitoringAlertsAlert) Set(val *ApiMonitoringAlertsAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMonitoringAlertsAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMonitoringAlertsAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMonitoringAlertsAlert(val *ApiMonitoringAlertsAlert) *NullableApiMonitoringAlertsAlert {
	return &NullableApiMonitoringAlertsAlert{value: val, isSet: true}
}

func (v NullableApiMonitoringAlertsAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMonitoringAlertsAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


