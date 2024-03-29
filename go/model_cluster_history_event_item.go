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

// ClusterHistoryEventItem struct for ClusterHistoryEventItem
type ClusterHistoryEventItem struct {
	Id *int64 `json:"id,omitempty"`
	ProcessId *int64 `json:"processId,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	UniqueId *string `json:"uniqueId,omitempty"`
	ProcessType *ClusterContainersAvailableActions `json:"processType,omitempty"`
	Description NullableString `json:"description,omitempty"`
	RefType *string `json:"refType,omitempty"`
	RefId *int64 `json:"refId,omitempty"`
	SubType NullableString `json:"subType,omitempty"`
	SubId NullableString `json:"subId,omitempty"`
	ZoneId NullableString `json:"zoneId,omitempty"`
	IntegrationId NullableString `json:"integrationId,omitempty"`
	InstanceId NullableString `json:"instanceId,omitempty"`
	ContainerId NullableString `json:"containerId,omitempty"`
	ServerId *int64 `json:"serverId,omitempty"`
	ContainerName NullableString `json:"containerName,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Status *string `json:"status,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
	Percent *int64 `json:"percent,omitempty"`
	StatusEta *int64 `json:"statusEta,omitempty"`
	Message *string `json:"message,omitempty"`
	Output NullableString `json:"output,omitempty"`
	Error NullableString `json:"error,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	CreatedBy *ClusterHistoryCreatedBy `json:"createdBy,omitempty"`
	UpdatedBy *ClusterHistoryCreatedBy `json:"updatedBy,omitempty"`
}

// NewClusterHistoryEventItem instantiates a new ClusterHistoryEventItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterHistoryEventItem() *ClusterHistoryEventItem {
	this := ClusterHistoryEventItem{}
	return &this
}

// NewClusterHistoryEventItemWithDefaults instantiates a new ClusterHistoryEventItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterHistoryEventItemWithDefaults() *ClusterHistoryEventItem {
	this := ClusterHistoryEventItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterHistoryEventItem) SetId(v int64) {
	o.Id = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetProcessId() int64 {
	if o == nil || o.ProcessId == nil {
		var ret int64
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetProcessIdOk() (*int64, bool) {
	if o == nil || o.ProcessId == nil {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasProcessId() bool {
	if o != nil && o.ProcessId != nil {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given int64 and assigns it to the ProcessId field.
func (o *ClusterHistoryEventItem) SetProcessId(v int64) {
	o.ProcessId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *ClusterHistoryEventItem) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetUniqueId() string {
	if o == nil || o.UniqueId == nil {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetUniqueIdOk() (*string, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *ClusterHistoryEventItem) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetProcessType() ClusterContainersAvailableActions {
	if o == nil || o.ProcessType == nil {
		var ret ClusterContainersAvailableActions
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetProcessTypeOk() (*ClusterContainersAvailableActions, bool) {
	if o == nil || o.ProcessType == nil {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasProcessType() bool {
	if o != nil && o.ProcessType != nil {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given ClusterContainersAvailableActions and assigns it to the ProcessType field.
func (o *ClusterHistoryEventItem) SetProcessType(v ClusterContainersAvailableActions) {
	o.ProcessType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterHistoryEventItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterHistoryEventItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetDescription() {
	o.Description.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasRefType() bool {
	if o != nil && o.RefType != nil {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *ClusterHistoryEventItem) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetRefId() int64 {
	if o == nil || o.RefId == nil {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetRefIdOk() (*int64, bool) {
	if o == nil || o.RefId == nil {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasRefId() bool {
	if o != nil && o.RefId != nil {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *ClusterHistoryEventItem) SetRefId(v int64) {
	o.RefId = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *ClusterHistoryEventItem) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *ClusterHistoryEventItem) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetSubType() {
	o.SubType.Unset()
}

// GetSubId returns the SubId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetSubId() string {
	if o == nil || o.SubId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubId.Get()
}

// GetSubIdOk returns a tuple with the SubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetSubIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubId.Get(), o.SubId.IsSet()
}

// HasSubId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasSubId() bool {
	if o != nil && o.SubId.IsSet() {
		return true
	}

	return false
}

// SetSubId gets a reference to the given NullableString and assigns it to the SubId field.
func (o *ClusterHistoryEventItem) SetSubId(v string) {
	o.SubId.Set(&v)
}
// SetSubIdNil sets the value for SubId to be an explicit nil
func (o *ClusterHistoryEventItem) SetSubIdNil() {
	o.SubId.Set(nil)
}

// UnsetSubId ensures that no value is present for SubId, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetSubId() {
	o.SubId.Unset()
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetZoneId() string {
	if o == nil || o.ZoneId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZoneId.Get()
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetZoneIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZoneId.Get(), o.ZoneId.IsSet()
}

// HasZoneId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasZoneId() bool {
	if o != nil && o.ZoneId.IsSet() {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given NullableString and assigns it to the ZoneId field.
func (o *ClusterHistoryEventItem) SetZoneId(v string) {
	o.ZoneId.Set(&v)
}
// SetZoneIdNil sets the value for ZoneId to be an explicit nil
func (o *ClusterHistoryEventItem) SetZoneIdNil() {
	o.ZoneId.Set(nil)
}

// UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetZoneId() {
	o.ZoneId.Unset()
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetIntegrationId() string {
	if o == nil || o.IntegrationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId.Get()
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetIntegrationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IntegrationId.Get(), o.IntegrationId.IsSet()
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasIntegrationId() bool {
	if o != nil && o.IntegrationId.IsSet() {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given NullableString and assigns it to the IntegrationId field.
func (o *ClusterHistoryEventItem) SetIntegrationId(v string) {
	o.IntegrationId.Set(&v)
}
// SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil
func (o *ClusterHistoryEventItem) SetIntegrationIdNil() {
	o.IntegrationId.Set(nil)
}

// UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetIntegrationId() {
	o.IntegrationId.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetInstanceId() string {
	if o == nil || o.InstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetInstanceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableString and assigns it to the InstanceId field.
func (o *ClusterHistoryEventItem) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}
// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *ClusterHistoryEventItem) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetContainerId() string {
	if o == nil || o.ContainerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerId.Get()
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetContainerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerId.Get(), o.ContainerId.IsSet()
}

// HasContainerId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasContainerId() bool {
	if o != nil && o.ContainerId.IsSet() {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given NullableString and assigns it to the ContainerId field.
func (o *ClusterHistoryEventItem) SetContainerId(v string) {
	o.ContainerId.Set(&v)
}
// SetContainerIdNil sets the value for ContainerId to be an explicit nil
func (o *ClusterHistoryEventItem) SetContainerIdNil() {
	o.ContainerId.Set(nil)
}

// UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetContainerId() {
	o.ContainerId.Unset()
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetServerId() int64 {
	if o == nil || o.ServerId == nil {
		var ret int64
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetServerIdOk() (*int64, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int64 and assigns it to the ServerId field.
func (o *ClusterHistoryEventItem) SetServerId(v int64) {
	o.ServerId = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetContainerName() string {
	if o == nil || o.ContainerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerName.Get()
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerName.Get(), o.ContainerName.IsSet()
}

// HasContainerName returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasContainerName() bool {
	if o != nil && o.ContainerName.IsSet() {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given NullableString and assigns it to the ContainerName field.
func (o *ClusterHistoryEventItem) SetContainerName(v string) {
	o.ContainerName.Set(&v)
}
// SetContainerNameNil sets the value for ContainerName to be an explicit nil
func (o *ClusterHistoryEventItem) SetContainerNameNil() {
	o.ContainerName.Set(nil)
}

// UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetContainerName() {
	o.ContainerName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterHistoryEventItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterHistoryEventItem) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *ClusterHistoryEventItem) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *ClusterHistoryEventItem) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetReason() {
	o.Reason.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetPercent() int64 {
	if o == nil || o.Percent == nil {
		var ret int64
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetPercentOk() (*int64, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given int64 and assigns it to the Percent field.
func (o *ClusterHistoryEventItem) SetPercent(v int64) {
	o.Percent = &v
}

// GetStatusEta returns the StatusEta field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetStatusEta() int64 {
	if o == nil || o.StatusEta == nil {
		var ret int64
		return ret
	}
	return *o.StatusEta
}

// GetStatusEtaOk returns a tuple with the StatusEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetStatusEtaOk() (*int64, bool) {
	if o == nil || o.StatusEta == nil {
		return nil, false
	}
	return o.StatusEta, true
}

// HasStatusEta returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasStatusEta() bool {
	if o != nil && o.StatusEta != nil {
		return true
	}

	return false
}

// SetStatusEta gets a reference to the given int64 and assigns it to the StatusEta field.
func (o *ClusterHistoryEventItem) SetStatusEta(v int64) {
	o.StatusEta = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterHistoryEventItem) SetMessage(v string) {
	o.Message = &v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetOutput() string {
	if o == nil || o.Output.Get() == nil {
		var ret string
		return ret
	}
	return *o.Output.Get()
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetOutputOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Output.Get(), o.Output.IsSet()
}

// HasOutput returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasOutput() bool {
	if o != nil && o.Output.IsSet() {
		return true
	}

	return false
}

// SetOutput gets a reference to the given NullableString and assigns it to the Output field.
func (o *ClusterHistoryEventItem) SetOutput(v string) {
	o.Output.Set(&v)
}
// SetOutputNil sets the value for Output to be an explicit nil
func (o *ClusterHistoryEventItem) SetOutputNil() {
	o.Output.Set(nil)
}

// UnsetOutput ensures that no value is present for Output, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetOutput() {
	o.Output.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistoryEventItem) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistoryEventItem) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ClusterHistoryEventItem) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ClusterHistoryEventItem) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ClusterHistoryEventItem) UnsetError() {
	o.Error.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ClusterHistoryEventItem) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ClusterHistoryEventItem) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *ClusterHistoryEventItem) SetDuration(v int64) {
	o.Duration = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ClusterHistoryEventItem) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ClusterHistoryEventItem) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetCreatedBy() ClusterHistoryCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret ClusterHistoryCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetCreatedByOk() (*ClusterHistoryCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ClusterHistoryCreatedBy and assigns it to the CreatedBy field.
func (o *ClusterHistoryEventItem) SetCreatedBy(v ClusterHistoryCreatedBy) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ClusterHistoryEventItem) GetUpdatedBy() ClusterHistoryCreatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret ClusterHistoryCreatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistoryEventItem) GetUpdatedByOk() (*ClusterHistoryCreatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ClusterHistoryEventItem) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given ClusterHistoryCreatedBy and assigns it to the UpdatedBy field.
func (o *ClusterHistoryEventItem) SetUpdatedBy(v ClusterHistoryCreatedBy) {
	o.UpdatedBy = &v
}

func (o ClusterHistoryEventItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProcessId != nil {
		toSerialize["processId"] = o.ProcessId
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.UniqueId != nil {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if o.ProcessType != nil {
		toSerialize["processType"] = o.ProcessType
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RefType != nil {
		toSerialize["refType"] = o.RefType
	}
	if o.RefId != nil {
		toSerialize["refId"] = o.RefId
	}
	if o.SubType.IsSet() {
		toSerialize["subType"] = o.SubType.Get()
	}
	if o.SubId.IsSet() {
		toSerialize["subId"] = o.SubId.Get()
	}
	if o.ZoneId.IsSet() {
		toSerialize["zoneId"] = o.ZoneId.Get()
	}
	if o.IntegrationId.IsSet() {
		toSerialize["integrationId"] = o.IntegrationId.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["instanceId"] = o.InstanceId.Get()
	}
	if o.ContainerId.IsSet() {
		toSerialize["containerId"] = o.ContainerId.Get()
	}
	if o.ServerId != nil {
		toSerialize["serverId"] = o.ServerId
	}
	if o.ContainerName.IsSet() {
		toSerialize["containerName"] = o.ContainerName.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Percent != nil {
		toSerialize["percent"] = o.Percent
	}
	if o.StatusEta != nil {
		toSerialize["statusEta"] = o.StatusEta
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Output.IsSet() {
		toSerialize["output"] = o.Output.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
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
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableClusterHistoryEventItem struct {
	value *ClusterHistoryEventItem
	isSet bool
}

func (v NullableClusterHistoryEventItem) Get() *ClusterHistoryEventItem {
	return v.value
}

func (v *NullableClusterHistoryEventItem) Set(val *ClusterHistoryEventItem) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHistoryEventItem) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHistoryEventItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHistoryEventItem(val *ClusterHistoryEventItem) *NullableClusterHistoryEventItem {
	return &NullableClusterHistoryEventItem{value: val, isSet: true}
}

func (v NullableClusterHistoryEventItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHistoryEventItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


