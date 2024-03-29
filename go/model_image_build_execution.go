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

// ImageBuildExecution struct for ImageBuildExecution
type ImageBuildExecution struct {
	Id *int64 `json:"id,omitempty"`
	ImageBuild *InlineResponse20040AppDeployInstance `json:"imageBuild,omitempty"`
	BuildNumber *int64 `json:"buildNumber,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate NullableTime `json:"endDate,omitempty"`
	StatusMessage NullableString `json:"statusMessage,omitempty"`
	StatusPercent *int64 `json:"statusPercent,omitempty"`
	StatusEta NullableString `json:"statusEta,omitempty"`
	Status *string `json:"status,omitempty"`
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	CreatedBy *ArchiveBucketFileCreatedBy `json:"createdBy,omitempty"`
	TempInstance NullableString `json:"tempInstance,omitempty"`
	VirtualImages []map[string]interface{} `json:"virtualImages,omitempty"`
}

// NewImageBuildExecution instantiates a new ImageBuildExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBuildExecution() *ImageBuildExecution {
	this := ImageBuildExecution{}
	return &this
}

// NewImageBuildExecutionWithDefaults instantiates a new ImageBuildExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBuildExecutionWithDefaults() *ImageBuildExecution {
	this := ImageBuildExecution{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ImageBuildExecution) SetId(v int64) {
	o.Id = &v
}

// GetImageBuild returns the ImageBuild field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetImageBuild() InlineResponse20040AppDeployInstance {
	if o == nil || o.ImageBuild == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.ImageBuild
}

// GetImageBuildOk returns a tuple with the ImageBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetImageBuildOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.ImageBuild == nil {
		return nil, false
	}
	return o.ImageBuild, true
}

// HasImageBuild returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasImageBuild() bool {
	if o != nil && o.ImageBuild != nil {
		return true
	}

	return false
}

// SetImageBuild gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the ImageBuild field.
func (o *ImageBuildExecution) SetImageBuild(v InlineResponse20040AppDeployInstance) {
	o.ImageBuild = &v
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetBuildNumber() int64 {
	if o == nil || o.BuildNumber == nil {
		var ret int64
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetBuildNumberOk() (*int64, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasBuildNumber() bool {
	if o != nil && o.BuildNumber != nil {
		return true
	}

	return false
}

// SetBuildNumber gets a reference to the given int64 and assigns it to the BuildNumber field.
func (o *ImageBuildExecution) SetBuildNumber(v int64) {
	o.BuildNumber = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ImageBuildExecution) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetEndDate() time.Time {
	if o == nil || o.EndDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetEndDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *ImageBuildExecution) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *ImageBuildExecution) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *ImageBuildExecution) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetStatusMessage() string {
	if o == nil || o.StatusMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetStatusMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *ImageBuildExecution) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *ImageBuildExecution) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *ImageBuildExecution) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetStatusPercent returns the StatusPercent field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetStatusPercent() int64 {
	if o == nil || o.StatusPercent == nil {
		var ret int64
		return ret
	}
	return *o.StatusPercent
}

// GetStatusPercentOk returns a tuple with the StatusPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetStatusPercentOk() (*int64, bool) {
	if o == nil || o.StatusPercent == nil {
		return nil, false
	}
	return o.StatusPercent, true
}

// HasStatusPercent returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasStatusPercent() bool {
	if o != nil && o.StatusPercent != nil {
		return true
	}

	return false
}

// SetStatusPercent gets a reference to the given int64 and assigns it to the StatusPercent field.
func (o *ImageBuildExecution) SetStatusPercent(v int64) {
	o.StatusPercent = &v
}

// GetStatusEta returns the StatusEta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetStatusEta() string {
	if o == nil || o.StatusEta.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusEta.Get()
}

// GetStatusEtaOk returns a tuple with the StatusEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetStatusEtaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusEta.Get(), o.StatusEta.IsSet()
}

// HasStatusEta returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasStatusEta() bool {
	if o != nil && o.StatusEta.IsSet() {
		return true
	}

	return false
}

// SetStatusEta gets a reference to the given NullableString and assigns it to the StatusEta field.
func (o *ImageBuildExecution) SetStatusEta(v string) {
	o.StatusEta.Set(&v)
}
// SetStatusEtaNil sets the value for StatusEta to be an explicit nil
func (o *ImageBuildExecution) SetStatusEtaNil() {
	o.StatusEta.Set(nil)
}

// UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
func (o *ImageBuildExecution) UnsetStatusEta() {
	o.StatusEta.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImageBuildExecution) SetStatus(v string) {
	o.Status = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ImageBuildExecution) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ImageBuildExecution) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ImageBuildExecution) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ImageBuildExecution) GetCreatedBy() ArchiveBucketFileCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret ArchiveBucketFileCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildExecution) GetCreatedByOk() (*ArchiveBucketFileCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ArchiveBucketFileCreatedBy and assigns it to the CreatedBy field.
func (o *ImageBuildExecution) SetCreatedBy(v ArchiveBucketFileCreatedBy) {
	o.CreatedBy = &v
}

// GetTempInstance returns the TempInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetTempInstance() string {
	if o == nil || o.TempInstance.Get() == nil {
		var ret string
		return ret
	}
	return *o.TempInstance.Get()
}

// GetTempInstanceOk returns a tuple with the TempInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetTempInstanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TempInstance.Get(), o.TempInstance.IsSet()
}

// HasTempInstance returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasTempInstance() bool {
	if o != nil && o.TempInstance.IsSet() {
		return true
	}

	return false
}

// SetTempInstance gets a reference to the given NullableString and assigns it to the TempInstance field.
func (o *ImageBuildExecution) SetTempInstance(v string) {
	o.TempInstance.Set(&v)
}
// SetTempInstanceNil sets the value for TempInstance to be an explicit nil
func (o *ImageBuildExecution) SetTempInstanceNil() {
	o.TempInstance.Set(nil)
}

// UnsetTempInstance ensures that no value is present for TempInstance, not even an explicit nil
func (o *ImageBuildExecution) UnsetTempInstance() {
	o.TempInstance.Unset()
}

// GetVirtualImages returns the VirtualImages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildExecution) GetVirtualImages() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.VirtualImages
}

// GetVirtualImagesOk returns a tuple with the VirtualImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildExecution) GetVirtualImagesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.VirtualImages == nil {
		return nil, false
	}
	return &o.VirtualImages, true
}

// HasVirtualImages returns a boolean if a field has been set.
func (o *ImageBuildExecution) HasVirtualImages() bool {
	if o != nil && o.VirtualImages != nil {
		return true
	}

	return false
}

// SetVirtualImages gets a reference to the given []map[string]interface{} and assigns it to the VirtualImages field.
func (o *ImageBuildExecution) SetVirtualImages(v []map[string]interface{}) {
	o.VirtualImages = v
}

func (o ImageBuildExecution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ImageBuild != nil {
		toSerialize["imageBuild"] = o.ImageBuild
	}
	if o.BuildNumber != nil {
		toSerialize["buildNumber"] = o.BuildNumber
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if o.StatusMessage.IsSet() {
		toSerialize["statusMessage"] = o.StatusMessage.Get()
	}
	if o.StatusPercent != nil {
		toSerialize["statusPercent"] = o.StatusPercent
	}
	if o.StatusEta.IsSet() {
		toSerialize["statusEta"] = o.StatusEta.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.TempInstance.IsSet() {
		toSerialize["tempInstance"] = o.TempInstance.Get()
	}
	if o.VirtualImages != nil {
		toSerialize["virtualImages"] = o.VirtualImages
	}
	return json.Marshal(toSerialize)
}

type NullableImageBuildExecution struct {
	value *ImageBuildExecution
	isSet bool
}

func (v NullableImageBuildExecution) Get() *ImageBuildExecution {
	return v.value
}

func (v *NullableImageBuildExecution) Set(val *ImageBuildExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildExecution(val *ImageBuildExecution) *NullableImageBuildExecution {
	return &NullableImageBuildExecution{value: val, isSet: true}
}

func (v NullableImageBuildExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


