/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AppCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCreateResponse{}

// AppCreateResponse struct for AppCreateResponse
type AppCreateResponse struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Environment *string `json:"environment,omitempty"`
	AccountId *int64 `json:"accountId,omitempty"`
	Account *ApplianceSettingsEnabledZoneTypesInner `json:"account,omitempty"`
	Owner *ActivityActivityInnerUser `json:"owner,omitempty"`
	SiteId *int64 `json:"siteId,omitempty"`
	Group *ApplianceSettingsEnabledZoneTypesInner `json:"group,omitempty"`
	Blueprint *AppBlueprint `json:"blueprint,omitempty"`
	Type *string `json:"type,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	RemovalDate NullableTime `json:"removalDate,omitempty"`
	AppContext *string `json:"appContext,omitempty"`
	Status *string `json:"status,omitempty"`
	AppStatus *string `json:"appStatus,omitempty"`
	InstanceCount *int64 `json:"instanceCount,omitempty"`
	ContainerCount *int64 `json:"containerCount,omitempty"`
	AppTiers []map[string]interface{} `json:"appTiers,omitempty"`
	Instances []ApplianceSettingsEnabledZoneTypesInner `json:"instances,omitempty"`
}

// NewAppCreateResponse instantiates a new AppCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCreateResponse() *AppCreateResponse {
	this := AppCreateResponse{}
	return &this
}

// NewAppCreateResponseWithDefaults instantiates a new AppCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCreateResponseWithDefaults() *AppCreateResponse {
	this := AppCreateResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppCreateResponse) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AppCreateResponse) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppCreateResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppCreateResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppCreateResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppCreateResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppCreateResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppCreateResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AppCreateResponse) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AppCreateResponse) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AppCreateResponse) SetLabels(v []string) {
	o.Labels = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AppCreateResponse) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AppCreateResponse) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *AppCreateResponse) SetEnvironment(v string) {
	o.Environment = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AppCreateResponse) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AppCreateResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *AppCreateResponse) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AppCreateResponse) GetAccount() ApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Account) {
		var ret ApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetAccountOk() (*ApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AppCreateResponse) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ApplianceSettingsEnabledZoneTypesInner and assigns it to the Account field.
func (o *AppCreateResponse) SetAccount(v ApplianceSettingsEnabledZoneTypesInner) {
	o.Account = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *AppCreateResponse) GetOwner() ActivityActivityInnerUser {
	if o == nil || IsNil(o.Owner) {
		var ret ActivityActivityInnerUser
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetOwnerOk() (*ActivityActivityInnerUser, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *AppCreateResponse) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ActivityActivityInnerUser and assigns it to the Owner field.
func (o *AppCreateResponse) SetOwner(v ActivityActivityInnerUser) {
	o.Owner = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *AppCreateResponse) GetSiteId() int64 {
	if o == nil || IsNil(o.SiteId) {
		var ret int64
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetSiteIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *AppCreateResponse) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int64 and assigns it to the SiteId field.
func (o *AppCreateResponse) SetSiteId(v int64) {
	o.SiteId = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AppCreateResponse) GetGroup() ApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Group) {
		var ret ApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetGroupOk() (*ApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AppCreateResponse) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given ApplianceSettingsEnabledZoneTypesInner and assigns it to the Group field.
func (o *AppCreateResponse) SetGroup(v ApplianceSettingsEnabledZoneTypesInner) {
	o.Group = &v
}

// GetBlueprint returns the Blueprint field value if set, zero value otherwise.
func (o *AppCreateResponse) GetBlueprint() AppBlueprint {
	if o == nil || IsNil(o.Blueprint) {
		var ret AppBlueprint
		return ret
	}
	return *o.Blueprint
}

// GetBlueprintOk returns a tuple with the Blueprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetBlueprintOk() (*AppBlueprint, bool) {
	if o == nil || IsNil(o.Blueprint) {
		return nil, false
	}
	return o.Blueprint, true
}

// HasBlueprint returns a boolean if a field has been set.
func (o *AppCreateResponse) HasBlueprint() bool {
	if o != nil && !IsNil(o.Blueprint) {
		return true
	}

	return false
}

// SetBlueprint gets a reference to the given AppBlueprint and assigns it to the Blueprint field.
func (o *AppCreateResponse) SetBlueprint(v AppBlueprint) {
	o.Blueprint = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AppCreateResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AppCreateResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AppCreateResponse) SetType(v string) {
	o.Type = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *AppCreateResponse) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *AppCreateResponse) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *AppCreateResponse) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AppCreateResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AppCreateResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AppCreateResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRemovalDate returns the RemovalDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppCreateResponse) GetRemovalDate() time.Time {
	if o == nil || IsNil(o.RemovalDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RemovalDate.Get()
}

// GetRemovalDateOk returns a tuple with the RemovalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppCreateResponse) GetRemovalDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovalDate.Get(), o.RemovalDate.IsSet()
}

// HasRemovalDate returns a boolean if a field has been set.
func (o *AppCreateResponse) HasRemovalDate() bool {
	if o != nil && o.RemovalDate.IsSet() {
		return true
	}

	return false
}

// SetRemovalDate gets a reference to the given NullableTime and assigns it to the RemovalDate field.
func (o *AppCreateResponse) SetRemovalDate(v time.Time) {
	o.RemovalDate.Set(&v)
}
// SetRemovalDateNil sets the value for RemovalDate to be an explicit nil
func (o *AppCreateResponse) SetRemovalDateNil() {
	o.RemovalDate.Set(nil)
}

// UnsetRemovalDate ensures that no value is present for RemovalDate, not even an explicit nil
func (o *AppCreateResponse) UnsetRemovalDate() {
	o.RemovalDate.Unset()
}

// GetAppContext returns the AppContext field value if set, zero value otherwise.
func (o *AppCreateResponse) GetAppContext() string {
	if o == nil || IsNil(o.AppContext) {
		var ret string
		return ret
	}
	return *o.AppContext
}

// GetAppContextOk returns a tuple with the AppContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetAppContextOk() (*string, bool) {
	if o == nil || IsNil(o.AppContext) {
		return nil, false
	}
	return o.AppContext, true
}

// HasAppContext returns a boolean if a field has been set.
func (o *AppCreateResponse) HasAppContext() bool {
	if o != nil && !IsNil(o.AppContext) {
		return true
	}

	return false
}

// SetAppContext gets a reference to the given string and assigns it to the AppContext field.
func (o *AppCreateResponse) SetAppContext(v string) {
	o.AppContext = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppCreateResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppCreateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppCreateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *AppCreateResponse) GetAppStatus() string {
	if o == nil || IsNil(o.AppStatus) {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetAppStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppStatus) {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *AppCreateResponse) HasAppStatus() bool {
	if o != nil && !IsNil(o.AppStatus) {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *AppCreateResponse) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *AppCreateResponse) GetInstanceCount() int64 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret int64
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetInstanceCountOk() (*int64, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *AppCreateResponse) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int64 and assigns it to the InstanceCount field.
func (o *AppCreateResponse) SetInstanceCount(v int64) {
	o.InstanceCount = &v
}

// GetContainerCount returns the ContainerCount field value if set, zero value otherwise.
func (o *AppCreateResponse) GetContainerCount() int64 {
	if o == nil || IsNil(o.ContainerCount) {
		var ret int64
		return ret
	}
	return *o.ContainerCount
}

// GetContainerCountOk returns a tuple with the ContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetContainerCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ContainerCount) {
		return nil, false
	}
	return o.ContainerCount, true
}

// HasContainerCount returns a boolean if a field has been set.
func (o *AppCreateResponse) HasContainerCount() bool {
	if o != nil && !IsNil(o.ContainerCount) {
		return true
	}

	return false
}

// SetContainerCount gets a reference to the given int64 and assigns it to the ContainerCount field.
func (o *AppCreateResponse) SetContainerCount(v int64) {
	o.ContainerCount = &v
}

// GetAppTiers returns the AppTiers field value if set, zero value otherwise.
func (o *AppCreateResponse) GetAppTiers() []map[string]interface{} {
	if o == nil || IsNil(o.AppTiers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.AppTiers
}

// GetAppTiersOk returns a tuple with the AppTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetAppTiersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AppTiers) {
		return nil, false
	}
	return o.AppTiers, true
}

// HasAppTiers returns a boolean if a field has been set.
func (o *AppCreateResponse) HasAppTiers() bool {
	if o != nil && !IsNil(o.AppTiers) {
		return true
	}

	return false
}

// SetAppTiers gets a reference to the given []map[string]interface{} and assigns it to the AppTiers field.
func (o *AppCreateResponse) SetAppTiers(v []map[string]interface{}) {
	o.AppTiers = v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *AppCreateResponse) GetInstances() []ApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Instances) {
		var ret []ApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCreateResponse) GetInstancesOk() ([]ApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *AppCreateResponse) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ApplianceSettingsEnabledZoneTypesInner and assigns it to the Instances field.
func (o *AppCreateResponse) SetInstances(v []ApplianceSettingsEnabledZoneTypesInner) {
	o.Instances = v
}

func (o AppCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.SiteId) {
		toSerialize["siteId"] = o.SiteId
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Blueprint) {
		toSerialize["blueprint"] = o.Blueprint
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.RemovalDate.IsSet() {
		toSerialize["removalDate"] = o.RemovalDate.Get()
	}
	if !IsNil(o.AppContext) {
		toSerialize["appContext"] = o.AppContext
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AppStatus) {
		toSerialize["appStatus"] = o.AppStatus
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instanceCount"] = o.InstanceCount
	}
	if !IsNil(o.ContainerCount) {
		toSerialize["containerCount"] = o.ContainerCount
	}
	if !IsNil(o.AppTiers) {
		toSerialize["appTiers"] = o.AppTiers
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableAppCreateResponse struct {
	value *AppCreateResponse
	isSet bool
}

func (v NullableAppCreateResponse) Get() *AppCreateResponse {
	return v.value
}

func (v *NullableAppCreateResponse) Set(val *AppCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCreateResponse(val *AppCreateResponse) *NullableAppCreateResponse {
	return &NullableAppCreateResponse{value: val, isSet: true}
}

func (v NullableAppCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


