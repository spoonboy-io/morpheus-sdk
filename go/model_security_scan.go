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

// SecurityScan struct for SecurityScan
type SecurityScan struct {
	Id *int64 `json:"id,omitempty"`
	SecurityPackage *SecurityScanSecurityPackage `json:"securityPackage,omitempty"`
	Status *string `json:"status,omitempty"`
	ScanDate *time.Time `json:"scanDate,omitempty"`
	ScanDuration *int64 `json:"scanDuration,omitempty"`
	TestCount *int64 `json:"testCount,omitempty"`
	RunCount *int64 `json:"runCount,omitempty"`
	PassCount *int64 `json:"passCount,omitempty"`
	FailCount *int64 `json:"failCount,omitempty"`
	OtherCount *int64 `json:"otherCount,omitempty"`
	ScanScore *float32 `json:"scanScore,omitempty"`
	ExternalId NullableString `json:"externalId,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	UpdatedBy *string `json:"updatedBy,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Results Summary (only returned when using query parameter results=true)
	Results *map[string]interface{} `json:"results,omitempty"`
}

// NewSecurityScan instantiates a new SecurityScan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityScan() *SecurityScan {
	this := SecurityScan{}
	return &this
}

// NewSecurityScanWithDefaults instantiates a new SecurityScan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityScanWithDefaults() *SecurityScan {
	this := SecurityScan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityScan) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityScan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SecurityScan) SetId(v int64) {
	o.Id = &v
}

// GetSecurityPackage returns the SecurityPackage field value if set, zero value otherwise.
func (o *SecurityScan) GetSecurityPackage() SecurityScanSecurityPackage {
	if o == nil || o.SecurityPackage == nil {
		var ret SecurityScanSecurityPackage
		return ret
	}
	return *o.SecurityPackage
}

// GetSecurityPackageOk returns a tuple with the SecurityPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetSecurityPackageOk() (*SecurityScanSecurityPackage, bool) {
	if o == nil || o.SecurityPackage == nil {
		return nil, false
	}
	return o.SecurityPackage, true
}

// HasSecurityPackage returns a boolean if a field has been set.
func (o *SecurityScan) HasSecurityPackage() bool {
	if o != nil && o.SecurityPackage != nil {
		return true
	}

	return false
}

// SetSecurityPackage gets a reference to the given SecurityScanSecurityPackage and assigns it to the SecurityPackage field.
func (o *SecurityScan) SetSecurityPackage(v SecurityScanSecurityPackage) {
	o.SecurityPackage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityScan) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityScan) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityScan) SetStatus(v string) {
	o.Status = &v
}

// GetScanDate returns the ScanDate field value if set, zero value otherwise.
func (o *SecurityScan) GetScanDate() time.Time {
	if o == nil || o.ScanDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ScanDate
}

// GetScanDateOk returns a tuple with the ScanDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetScanDateOk() (*time.Time, bool) {
	if o == nil || o.ScanDate == nil {
		return nil, false
	}
	return o.ScanDate, true
}

// HasScanDate returns a boolean if a field has been set.
func (o *SecurityScan) HasScanDate() bool {
	if o != nil && o.ScanDate != nil {
		return true
	}

	return false
}

// SetScanDate gets a reference to the given time.Time and assigns it to the ScanDate field.
func (o *SecurityScan) SetScanDate(v time.Time) {
	o.ScanDate = &v
}

// GetScanDuration returns the ScanDuration field value if set, zero value otherwise.
func (o *SecurityScan) GetScanDuration() int64 {
	if o == nil || o.ScanDuration == nil {
		var ret int64
		return ret
	}
	return *o.ScanDuration
}

// GetScanDurationOk returns a tuple with the ScanDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetScanDurationOk() (*int64, bool) {
	if o == nil || o.ScanDuration == nil {
		return nil, false
	}
	return o.ScanDuration, true
}

// HasScanDuration returns a boolean if a field has been set.
func (o *SecurityScan) HasScanDuration() bool {
	if o != nil && o.ScanDuration != nil {
		return true
	}

	return false
}

// SetScanDuration gets a reference to the given int64 and assigns it to the ScanDuration field.
func (o *SecurityScan) SetScanDuration(v int64) {
	o.ScanDuration = &v
}

// GetTestCount returns the TestCount field value if set, zero value otherwise.
func (o *SecurityScan) GetTestCount() int64 {
	if o == nil || o.TestCount == nil {
		var ret int64
		return ret
	}
	return *o.TestCount
}

// GetTestCountOk returns a tuple with the TestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetTestCountOk() (*int64, bool) {
	if o == nil || o.TestCount == nil {
		return nil, false
	}
	return o.TestCount, true
}

// HasTestCount returns a boolean if a field has been set.
func (o *SecurityScan) HasTestCount() bool {
	if o != nil && o.TestCount != nil {
		return true
	}

	return false
}

// SetTestCount gets a reference to the given int64 and assigns it to the TestCount field.
func (o *SecurityScan) SetTestCount(v int64) {
	o.TestCount = &v
}

// GetRunCount returns the RunCount field value if set, zero value otherwise.
func (o *SecurityScan) GetRunCount() int64 {
	if o == nil || o.RunCount == nil {
		var ret int64
		return ret
	}
	return *o.RunCount
}

// GetRunCountOk returns a tuple with the RunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetRunCountOk() (*int64, bool) {
	if o == nil || o.RunCount == nil {
		return nil, false
	}
	return o.RunCount, true
}

// HasRunCount returns a boolean if a field has been set.
func (o *SecurityScan) HasRunCount() bool {
	if o != nil && o.RunCount != nil {
		return true
	}

	return false
}

// SetRunCount gets a reference to the given int64 and assigns it to the RunCount field.
func (o *SecurityScan) SetRunCount(v int64) {
	o.RunCount = &v
}

// GetPassCount returns the PassCount field value if set, zero value otherwise.
func (o *SecurityScan) GetPassCount() int64 {
	if o == nil || o.PassCount == nil {
		var ret int64
		return ret
	}
	return *o.PassCount
}

// GetPassCountOk returns a tuple with the PassCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetPassCountOk() (*int64, bool) {
	if o == nil || o.PassCount == nil {
		return nil, false
	}
	return o.PassCount, true
}

// HasPassCount returns a boolean if a field has been set.
func (o *SecurityScan) HasPassCount() bool {
	if o != nil && o.PassCount != nil {
		return true
	}

	return false
}

// SetPassCount gets a reference to the given int64 and assigns it to the PassCount field.
func (o *SecurityScan) SetPassCount(v int64) {
	o.PassCount = &v
}

// GetFailCount returns the FailCount field value if set, zero value otherwise.
func (o *SecurityScan) GetFailCount() int64 {
	if o == nil || o.FailCount == nil {
		var ret int64
		return ret
	}
	return *o.FailCount
}

// GetFailCountOk returns a tuple with the FailCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetFailCountOk() (*int64, bool) {
	if o == nil || o.FailCount == nil {
		return nil, false
	}
	return o.FailCount, true
}

// HasFailCount returns a boolean if a field has been set.
func (o *SecurityScan) HasFailCount() bool {
	if o != nil && o.FailCount != nil {
		return true
	}

	return false
}

// SetFailCount gets a reference to the given int64 and assigns it to the FailCount field.
func (o *SecurityScan) SetFailCount(v int64) {
	o.FailCount = &v
}

// GetOtherCount returns the OtherCount field value if set, zero value otherwise.
func (o *SecurityScan) GetOtherCount() int64 {
	if o == nil || o.OtherCount == nil {
		var ret int64
		return ret
	}
	return *o.OtherCount
}

// GetOtherCountOk returns a tuple with the OtherCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetOtherCountOk() (*int64, bool) {
	if o == nil || o.OtherCount == nil {
		return nil, false
	}
	return o.OtherCount, true
}

// HasOtherCount returns a boolean if a field has been set.
func (o *SecurityScan) HasOtherCount() bool {
	if o != nil && o.OtherCount != nil {
		return true
	}

	return false
}

// SetOtherCount gets a reference to the given int64 and assigns it to the OtherCount field.
func (o *SecurityScan) SetOtherCount(v int64) {
	o.OtherCount = &v
}

// GetScanScore returns the ScanScore field value if set, zero value otherwise.
func (o *SecurityScan) GetScanScore() float32 {
	if o == nil || o.ScanScore == nil {
		var ret float32
		return ret
	}
	return *o.ScanScore
}

// GetScanScoreOk returns a tuple with the ScanScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetScanScoreOk() (*float32, bool) {
	if o == nil || o.ScanScore == nil {
		return nil, false
	}
	return o.ScanScore, true
}

// HasScanScore returns a boolean if a field has been set.
func (o *SecurityScan) HasScanScore() bool {
	if o != nil && o.ScanScore != nil {
		return true
	}

	return false
}

// SetScanScore gets a reference to the given float32 and assigns it to the ScanScore field.
func (o *SecurityScan) SetScanScore(v float32) {
	o.ScanScore = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityScan) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityScan) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *SecurityScan) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *SecurityScan) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *SecurityScan) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *SecurityScan) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SecurityScan) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SecurityScan) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SecurityScan) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SecurityScan) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SecurityScan) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SecurityScan) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SecurityScan) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SecurityScan) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SecurityScan) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SecurityScan) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SecurityScan) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SecurityScan) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *SecurityScan) GetResults() map[string]interface{} {
	if o == nil || o.Results == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScan) GetResultsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *SecurityScan) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given map[string]interface{} and assigns it to the Results field.
func (o *SecurityScan) SetResults(v map[string]interface{}) {
	o.Results = &v
}

func (o SecurityScan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SecurityPackage != nil {
		toSerialize["securityPackage"] = o.SecurityPackage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ScanDate != nil {
		toSerialize["scanDate"] = o.ScanDate
	}
	if o.ScanDuration != nil {
		toSerialize["scanDuration"] = o.ScanDuration
	}
	if o.TestCount != nil {
		toSerialize["testCount"] = o.TestCount
	}
	if o.RunCount != nil {
		toSerialize["runCount"] = o.RunCount
	}
	if o.PassCount != nil {
		toSerialize["passCount"] = o.PassCount
	}
	if o.FailCount != nil {
		toSerialize["failCount"] = o.FailCount
	}
	if o.OtherCount != nil {
		toSerialize["otherCount"] = o.OtherCount
	}
	if o.ScanScore != nil {
		toSerialize["scanScore"] = o.ScanScore
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityScan struct {
	value *SecurityScan
	isSet bool
}

func (v NullableSecurityScan) Get() *SecurityScan {
	return v.value
}

func (v *NullableSecurityScan) Set(val *SecurityScan) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityScan) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityScan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityScan(val *SecurityScan) *NullableSecurityScan {
	return &NullableSecurityScan{value: val, isSet: true}
}

func (v NullableSecurityScan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityScan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


