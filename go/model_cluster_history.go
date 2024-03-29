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

// ClusterHistory struct for ClusterHistory
type ClusterHistory struct {
	Id *int64 `json:"id,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	UniqueId *string `json:"uniqueId,omitempty"`
	ProcessType *ClusterContainersAvailableActions `json:"processType,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Description NullableString `json:"description,omitempty"`
	SubType NullableString `json:"subType,omitempty"`
	SubId NullableString `json:"subId,omitempty"`
	ZoneId NullableInt64 `json:"zoneId,omitempty"`
	IntegrationId NullableInt64 `json:"integrationId,omitempty"`
	AppId NullableInt64 `json:"appId,omitempty"`
	InstanceId NullableInt64 `json:"instanceId,omitempty"`
	ContainerId NullableInt64 `json:"containerId,omitempty"`
	ServerId NullableInt64 `json:"serverId,omitempty"`
	ContainerName NullableString `json:"containerName,omitempty"`
	Status *string `json:"status,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
	Percent *float32 `json:"percent,omitempty"`
	StatusEta *int64 `json:"statusEta,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Output NullableString `json:"output,omitempty"`
	Error NullableString `json:"error,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	CreatedBy *ClusterHistoryCreatedBy `json:"createdBy,omitempty"`
	UpdatedBy *ClusterHistoryCreatedBy `json:"updatedBy,omitempty"`
	Events *[]ClusterHistoryEvents `json:"events,omitempty"`
}

// NewClusterHistory instantiates a new ClusterHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterHistory() *ClusterHistory {
	this := ClusterHistory{}
	return &this
}

// NewClusterHistoryWithDefaults instantiates a new ClusterHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterHistoryWithDefaults() *ClusterHistory {
	this := ClusterHistory{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterHistory) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterHistory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterHistory) SetId(v int64) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ClusterHistory) GetAccountId() int64 {
	if o == nil || o.AccountId == nil {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetAccountIdOk() (*int64, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ClusterHistory) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *ClusterHistory) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *ClusterHistory) GetUniqueId() string {
	if o == nil || o.UniqueId == nil {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetUniqueIdOk() (*string, bool) {
	if o == nil || o.UniqueId == nil {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *ClusterHistory) HasUniqueId() bool {
	if o != nil && o.UniqueId != nil {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *ClusterHistory) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetProcessType returns the ProcessType field value if set, zero value otherwise.
func (o *ClusterHistory) GetProcessType() ClusterContainersAvailableActions {
	if o == nil || o.ProcessType == nil {
		var ret ClusterContainersAvailableActions
		return ret
	}
	return *o.ProcessType
}

// GetProcessTypeOk returns a tuple with the ProcessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetProcessTypeOk() (*ClusterContainersAvailableActions, bool) {
	if o == nil || o.ProcessType == nil {
		return nil, false
	}
	return o.ProcessType, true
}

// HasProcessType returns a boolean if a field has been set.
func (o *ClusterHistory) HasProcessType() bool {
	if o != nil && o.ProcessType != nil {
		return true
	}

	return false
}

// SetProcessType gets a reference to the given ClusterContainersAvailableActions and assigns it to the ProcessType field.
func (o *ClusterHistory) SetProcessType(v ClusterContainersAvailableActions) {
	o.ProcessType = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterHistory) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterHistory) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterHistory) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterHistory) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterHistory) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterHistory) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterHistory) UnsetDescription() {
	o.Description.Unset()
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *ClusterHistory) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *ClusterHistory) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *ClusterHistory) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *ClusterHistory) UnsetSubType() {
	o.SubType.Unset()
}

// GetSubId returns the SubId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetSubId() string {
	if o == nil || o.SubId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubId.Get()
}

// GetSubIdOk returns a tuple with the SubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetSubIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubId.Get(), o.SubId.IsSet()
}

// HasSubId returns a boolean if a field has been set.
func (o *ClusterHistory) HasSubId() bool {
	if o != nil && o.SubId.IsSet() {
		return true
	}

	return false
}

// SetSubId gets a reference to the given NullableString and assigns it to the SubId field.
func (o *ClusterHistory) SetSubId(v string) {
	o.SubId.Set(&v)
}
// SetSubIdNil sets the value for SubId to be an explicit nil
func (o *ClusterHistory) SetSubIdNil() {
	o.SubId.Set(nil)
}

// UnsetSubId ensures that no value is present for SubId, not even an explicit nil
func (o *ClusterHistory) UnsetSubId() {
	o.SubId.Unset()
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetZoneId() int64 {
	if o == nil || o.ZoneId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ZoneId.Get()
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetZoneIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZoneId.Get(), o.ZoneId.IsSet()
}

// HasZoneId returns a boolean if a field has been set.
func (o *ClusterHistory) HasZoneId() bool {
	if o != nil && o.ZoneId.IsSet() {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given NullableInt64 and assigns it to the ZoneId field.
func (o *ClusterHistory) SetZoneId(v int64) {
	o.ZoneId.Set(&v)
}
// SetZoneIdNil sets the value for ZoneId to be an explicit nil
func (o *ClusterHistory) SetZoneIdNil() {
	o.ZoneId.Set(nil)
}

// UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
func (o *ClusterHistory) UnsetZoneId() {
	o.ZoneId.Unset()
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetIntegrationId() int64 {
	if o == nil || o.IntegrationId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IntegrationId.Get()
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetIntegrationIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IntegrationId.Get(), o.IntegrationId.IsSet()
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ClusterHistory) HasIntegrationId() bool {
	if o != nil && o.IntegrationId.IsSet() {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given NullableInt64 and assigns it to the IntegrationId field.
func (o *ClusterHistory) SetIntegrationId(v int64) {
	o.IntegrationId.Set(&v)
}
// SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil
func (o *ClusterHistory) SetIntegrationIdNil() {
	o.IntegrationId.Set(nil)
}

// UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
func (o *ClusterHistory) UnsetIntegrationId() {
	o.IntegrationId.Unset()
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetAppId() int64 {
	if o == nil || o.AppId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetAppIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// HasAppId returns a boolean if a field has been set.
func (o *ClusterHistory) HasAppId() bool {
	if o != nil && o.AppId.IsSet() {
		return true
	}

	return false
}

// SetAppId gets a reference to the given NullableInt64 and assigns it to the AppId field.
func (o *ClusterHistory) SetAppId(v int64) {
	o.AppId.Set(&v)
}
// SetAppIdNil sets the value for AppId to be an explicit nil
func (o *ClusterHistory) SetAppIdNil() {
	o.AppId.Set(nil)
}

// UnsetAppId ensures that no value is present for AppId, not even an explicit nil
func (o *ClusterHistory) UnsetAppId() {
	o.AppId.Unset()
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetInstanceId() int64 {
	if o == nil || o.InstanceId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetInstanceIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ClusterHistory) HasInstanceId() bool {
	if o != nil && o.InstanceId.IsSet() {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given NullableInt64 and assigns it to the InstanceId field.
func (o *ClusterHistory) SetInstanceId(v int64) {
	o.InstanceId.Set(&v)
}
// SetInstanceIdNil sets the value for InstanceId to be an explicit nil
func (o *ClusterHistory) SetInstanceIdNil() {
	o.InstanceId.Set(nil)
}

// UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
func (o *ClusterHistory) UnsetInstanceId() {
	o.InstanceId.Unset()
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetContainerId() int64 {
	if o == nil || o.ContainerId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId.Get()
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetContainerIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerId.Get(), o.ContainerId.IsSet()
}

// HasContainerId returns a boolean if a field has been set.
func (o *ClusterHistory) HasContainerId() bool {
	if o != nil && o.ContainerId.IsSet() {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given NullableInt64 and assigns it to the ContainerId field.
func (o *ClusterHistory) SetContainerId(v int64) {
	o.ContainerId.Set(&v)
}
// SetContainerIdNil sets the value for ContainerId to be an explicit nil
func (o *ClusterHistory) SetContainerIdNil() {
	o.ContainerId.Set(nil)
}

// UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
func (o *ClusterHistory) UnsetContainerId() {
	o.ContainerId.Unset()
}

// GetServerId returns the ServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetServerId() int64 {
	if o == nil || o.ServerId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ServerId.Get()
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetServerIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerId.Get(), o.ServerId.IsSet()
}

// HasServerId returns a boolean if a field has been set.
func (o *ClusterHistory) HasServerId() bool {
	if o != nil && o.ServerId.IsSet() {
		return true
	}

	return false
}

// SetServerId gets a reference to the given NullableInt64 and assigns it to the ServerId field.
func (o *ClusterHistory) SetServerId(v int64) {
	o.ServerId.Set(&v)
}
// SetServerIdNil sets the value for ServerId to be an explicit nil
func (o *ClusterHistory) SetServerIdNil() {
	o.ServerId.Set(nil)
}

// UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
func (o *ClusterHistory) UnsetServerId() {
	o.ServerId.Unset()
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetContainerName() string {
	if o == nil || o.ContainerName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContainerName.Get()
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetContainerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContainerName.Get(), o.ContainerName.IsSet()
}

// HasContainerName returns a boolean if a field has been set.
func (o *ClusterHistory) HasContainerName() bool {
	if o != nil && o.ContainerName.IsSet() {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given NullableString and assigns it to the ContainerName field.
func (o *ClusterHistory) SetContainerName(v string) {
	o.ContainerName.Set(&v)
}
// SetContainerNameNil sets the value for ContainerName to be an explicit nil
func (o *ClusterHistory) SetContainerNameNil() {
	o.ContainerName.Set(nil)
}

// UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
func (o *ClusterHistory) UnsetContainerName() {
	o.ContainerName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterHistory) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *ClusterHistory) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *ClusterHistory) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *ClusterHistory) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *ClusterHistory) UnsetReason() {
	o.Reason.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *ClusterHistory) GetPercent() float32 {
	if o == nil || o.Percent == nil {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetPercentOk() (*float32, bool) {
	if o == nil || o.Percent == nil {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *ClusterHistory) HasPercent() bool {
	if o != nil && o.Percent != nil {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *ClusterHistory) SetPercent(v float32) {
	o.Percent = &v
}

// GetStatusEta returns the StatusEta field value if set, zero value otherwise.
func (o *ClusterHistory) GetStatusEta() int64 {
	if o == nil || o.StatusEta == nil {
		var ret int64
		return ret
	}
	return *o.StatusEta
}

// GetStatusEtaOk returns a tuple with the StatusEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetStatusEtaOk() (*int64, bool) {
	if o == nil || o.StatusEta == nil {
		return nil, false
	}
	return o.StatusEta, true
}

// HasStatusEta returns a boolean if a field has been set.
func (o *ClusterHistory) HasStatusEta() bool {
	if o != nil && o.StatusEta != nil {
		return true
	}

	return false
}

// SetStatusEta gets a reference to the given int64 and assigns it to the StatusEta field.
func (o *ClusterHistory) SetStatusEta(v int64) {
	o.StatusEta = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterHistory) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *ClusterHistory) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ClusterHistory) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ClusterHistory) UnsetMessage() {
	o.Message.Unset()
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetOutput() string {
	if o == nil || o.Output.Get() == nil {
		var ret string
		return ret
	}
	return *o.Output.Get()
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetOutputOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Output.Get(), o.Output.IsSet()
}

// HasOutput returns a boolean if a field has been set.
func (o *ClusterHistory) HasOutput() bool {
	if o != nil && o.Output.IsSet() {
		return true
	}

	return false
}

// SetOutput gets a reference to the given NullableString and assigns it to the Output field.
func (o *ClusterHistory) SetOutput(v string) {
	o.Output.Set(&v)
}
// SetOutputNil sets the value for Output to be an explicit nil
func (o *ClusterHistory) SetOutputNil() {
	o.Output.Set(nil)
}

// UnsetOutput ensures that no value is present for Output, not even an explicit nil
func (o *ClusterHistory) UnsetOutput() {
	o.Output.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterHistory) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterHistory) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ClusterHistory) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ClusterHistory) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ClusterHistory) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ClusterHistory) UnsetError() {
	o.Error.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ClusterHistory) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ClusterHistory) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ClusterHistory) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ClusterHistory) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ClusterHistory) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ClusterHistory) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ClusterHistory) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ClusterHistory) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *ClusterHistory) SetDuration(v int64) {
	o.Duration = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ClusterHistory) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ClusterHistory) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ClusterHistory) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ClusterHistory) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ClusterHistory) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ClusterHistory) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ClusterHistory) GetCreatedBy() ClusterHistoryCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret ClusterHistoryCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetCreatedByOk() (*ClusterHistoryCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ClusterHistory) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ClusterHistoryCreatedBy and assigns it to the CreatedBy field.
func (o *ClusterHistory) SetCreatedBy(v ClusterHistoryCreatedBy) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ClusterHistory) GetUpdatedBy() ClusterHistoryCreatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret ClusterHistoryCreatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetUpdatedByOk() (*ClusterHistoryCreatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ClusterHistory) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given ClusterHistoryCreatedBy and assigns it to the UpdatedBy field.
func (o *ClusterHistory) SetUpdatedBy(v ClusterHistoryCreatedBy) {
	o.UpdatedBy = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ClusterHistory) GetEvents() []ClusterHistoryEvents {
	if o == nil || o.Events == nil {
		var ret []ClusterHistoryEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterHistory) GetEventsOk() (*[]ClusterHistoryEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ClusterHistory) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []ClusterHistoryEvents and assigns it to the Events field.
func (o *ClusterHistory) SetEvents(v []ClusterHistoryEvents) {
	o.Events = &v
}

func (o ClusterHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
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
	if o.AppId.IsSet() {
		toSerialize["appId"] = o.AppId.Get()
	}
	if o.InstanceId.IsSet() {
		toSerialize["instanceId"] = o.InstanceId.Get()
	}
	if o.ContainerId.IsSet() {
		toSerialize["containerId"] = o.ContainerId.Get()
	}
	if o.ServerId.IsSet() {
		toSerialize["serverId"] = o.ServerId.Get()
	}
	if o.ContainerName.IsSet() {
		toSerialize["containerName"] = o.ContainerName.Get()
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
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
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
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableClusterHistory struct {
	value *ClusterHistory
	isSet bool
}

func (v NullableClusterHistory) Get() *ClusterHistory {
	return v.value
}

func (v *NullableClusterHistory) Set(val *ClusterHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterHistory(val *ClusterHistory) *NullableClusterHistory {
	return &NullableClusterHistory{value: val, isSet: true}
}

func (v NullableClusterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


