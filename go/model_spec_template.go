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

// SpecTemplate struct for SpecTemplate
type SpecTemplate struct {
	Id *int64 `json:"id,omitempty"`
	Account *InlineResponse20040AppDeployInstance `json:"account,omitempty"`
	Name *string `json:"name,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Type *InlineResponse20094Network `json:"type,omitempty"`
	ExternalId NullableString `json:"externalId,omitempty"`
	ExternalType NullableString `json:"externalType,omitempty"`
	DeploymentId NullableString `json:"deploymentId,omitempty"`
	Status NullableString `json:"status,omitempty"`
	File *ClusterLayoutFile `json:"file,omitempty"`
	Config *map[string]interface{} `json:"config,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewSpecTemplate instantiates a new SpecTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecTemplate() *SpecTemplate {
	this := SpecTemplate{}
	return &this
}

// NewSpecTemplateWithDefaults instantiates a new SpecTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecTemplateWithDefaults() *SpecTemplate {
	this := SpecTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpecTemplate) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpecTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *SpecTemplate) SetId(v int64) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SpecTemplate) GetAccount() InlineResponse20040AppDeployInstance {
	if o == nil || o.Account == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SpecTemplate) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Account field.
func (o *SpecTemplate) SetAccount(v InlineResponse20040AppDeployInstance) {
	o.Account = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpecTemplate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpecTemplate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpecTemplate) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *SpecTemplate) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *SpecTemplate) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *SpecTemplate) SetLabels(v []string) {
	o.Labels = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *SpecTemplate) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *SpecTemplate) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *SpecTemplate) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *SpecTemplate) UnsetCode() {
	o.Code.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpecTemplate) GetType() InlineResponse20094Network {
	if o == nil || o.Type == nil {
		var ret InlineResponse20094Network
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetTypeOk() (*InlineResponse20094Network, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpecTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given InlineResponse20094Network and assigns it to the Type field.
func (o *SpecTemplate) SetType(v InlineResponse20094Network) {
	o.Type = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *SpecTemplate) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *SpecTemplate) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *SpecTemplate) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *SpecTemplate) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetExternalType returns the ExternalType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetExternalType() string {
	if o == nil || o.ExternalType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalType.Get()
}

// GetExternalTypeOk returns a tuple with the ExternalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetExternalTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalType.Get(), o.ExternalType.IsSet()
}

// HasExternalType returns a boolean if a field has been set.
func (o *SpecTemplate) HasExternalType() bool {
	if o != nil && o.ExternalType.IsSet() {
		return true
	}

	return false
}

// SetExternalType gets a reference to the given NullableString and assigns it to the ExternalType field.
func (o *SpecTemplate) SetExternalType(v string) {
	o.ExternalType.Set(&v)
}
// SetExternalTypeNil sets the value for ExternalType to be an explicit nil
func (o *SpecTemplate) SetExternalTypeNil() {
	o.ExternalType.Set(nil)
}

// UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
func (o *SpecTemplate) UnsetExternalType() {
	o.ExternalType.Unset()
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetDeploymentId() string {
	if o == nil || o.DeploymentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId.Get()
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetDeploymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeploymentId.Get(), o.DeploymentId.IsSet()
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *SpecTemplate) HasDeploymentId() bool {
	if o != nil && o.DeploymentId.IsSet() {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given NullableString and assigns it to the DeploymentId field.
func (o *SpecTemplate) SetDeploymentId(v string) {
	o.DeploymentId.Set(&v)
}
// SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil
func (o *SpecTemplate) SetDeploymentIdNil() {
	o.DeploymentId.Set(nil)
}

// UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
func (o *SpecTemplate) UnsetDeploymentId() {
	o.DeploymentId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *SpecTemplate) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *SpecTemplate) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *SpecTemplate) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *SpecTemplate) UnsetStatus() {
	o.Status.Unset()
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *SpecTemplate) GetFile() ClusterLayoutFile {
	if o == nil || o.File == nil {
		var ret ClusterLayoutFile
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetFileOk() (*ClusterLayoutFile, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *SpecTemplate) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given ClusterLayoutFile and assigns it to the File field.
func (o *SpecTemplate) SetFile(v ClusterLayoutFile) {
	o.File = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SpecTemplate) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SpecTemplate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *SpecTemplate) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SpecTemplate) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SpecTemplate) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SpecTemplate) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpecTemplate) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpecTemplate) GetUpdatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SpecTemplate) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *SpecTemplate) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *SpecTemplate) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *SpecTemplate) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SpecTemplate) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SpecTemplate) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SpecTemplate) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SpecTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SpecTemplate) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SpecTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o SpecTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.ExternalType.IsSet() {
		toSerialize["externalType"] = o.ExternalType.Get()
	}
	if o.DeploymentId.IsSet() {
		toSerialize["deploymentId"] = o.DeploymentId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updatedBy"] = o.UpdatedBy.Get()
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableSpecTemplate struct {
	value *SpecTemplate
	isSet bool
}

func (v NullableSpecTemplate) Get() *SpecTemplate {
	return v.value
}

func (v *NullableSpecTemplate) Set(val *SpecTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecTemplate(val *SpecTemplate) *NullableSpecTemplate {
	return &NullableSpecTemplate{value: val, isSet: true}
}

func (v NullableSpecTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


