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

// ServicePlan struct for ServicePlan
type ServicePlan struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Active *bool `json:"active,omitempty"`
	SortOrder *int64 `json:"sortOrder,omitempty"`
	Description *string `json:"description,omitempty"`
	MaxStorage *float32 `json:"maxStorage,omitempty"`
	MaxMemory *float32 `json:"maxMemory,omitempty"`
	MaxCpu NullableFloat32 `json:"maxCpu,omitempty"`
	MaxCores NullableFloat32 `json:"maxCores,omitempty"`
	MaxDisks NullableFloat32 `json:"maxDisks,omitempty"`
	CoresPerSocket *float32 `json:"coresPerSocket,omitempty"`
	CustomCpu *bool `json:"customCpu,omitempty"`
	CustomCores *bool `json:"customCores,omitempty"`
	CustomMaxStorage NullableBool `json:"customMaxStorage,omitempty"`
	CustomMaxDataStorage NullableBool `json:"customMaxDataStorage,omitempty"`
	CustomMaxMemory NullableBool `json:"customMaxMemory,omitempty"`
	AddVolumes NullableBool `json:"addVolumes,omitempty"`
	MemoryOptionSource NullableString `json:"memoryOptionSource,omitempty"`
	CpuOptionSource NullableString `json:"cpuOptionSource,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	RegionCode NullableString `json:"regionCode,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	ProvisionType *GuidanceVmwareSizingPlanBeforeActionProvisionType `json:"provisionType,omitempty"`
	Tenants *string `json:"tenants,omitempty"`
	PriceSets []GuidanceVmwareSizingPlanBeforeActionPriceSets `json:"priceSets,omitempty"`
	Config NullableServicePlanConfig `json:"config,omitempty"`
	Zones *[]InlineResponse20094Network `json:"zones,omitempty"`
	Permissions *ServicePlanPermissions `json:"permissions,omitempty"`
}

// NewServicePlan instantiates a new ServicePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlan() *ServicePlan {
	this := ServicePlan{}
	return &this
}

// NewServicePlanWithDefaults instantiates a new ServicePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanWithDefaults() *ServicePlan {
	this := ServicePlan{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServicePlan) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServicePlan) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ServicePlan) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServicePlan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServicePlan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServicePlan) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ServicePlan) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ServicePlan) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ServicePlan) SetCode(v string) {
	o.Code = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ServicePlan) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ServicePlan) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ServicePlan) SetActive(v bool) {
	o.Active = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *ServicePlan) GetSortOrder() int64 {
	if o == nil || o.SortOrder == nil {
		var ret int64
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetSortOrderOk() (*int64, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *ServicePlan) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int64 and assigns it to the SortOrder field.
func (o *ServicePlan) SetSortOrder(v int64) {
	o.SortOrder = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicePlan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicePlan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicePlan) SetDescription(v string) {
	o.Description = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *ServicePlan) GetMaxStorage() float32 {
	if o == nil || o.MaxStorage == nil {
		var ret float32
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetMaxStorageOk() (*float32, bool) {
	if o == nil || o.MaxStorage == nil {
		return nil, false
	}
	return o.MaxStorage, true
}

// HasMaxStorage returns a boolean if a field has been set.
func (o *ServicePlan) HasMaxStorage() bool {
	if o != nil && o.MaxStorage != nil {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given float32 and assigns it to the MaxStorage field.
func (o *ServicePlan) SetMaxStorage(v float32) {
	o.MaxStorage = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *ServicePlan) GetMaxMemory() float32 {
	if o == nil || o.MaxMemory == nil {
		var ret float32
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetMaxMemoryOk() (*float32, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *ServicePlan) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given float32 and assigns it to the MaxMemory field.
func (o *ServicePlan) SetMaxMemory(v float32) {
	o.MaxMemory = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetMaxCpu() float32 {
	if o == nil || o.MaxCpu.Get() == nil {
		var ret float32
		return ret
	}
	return *o.MaxCpu.Get()
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetMaxCpuOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxCpu.Get(), o.MaxCpu.IsSet()
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *ServicePlan) HasMaxCpu() bool {
	if o != nil && o.MaxCpu.IsSet() {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given NullableFloat32 and assigns it to the MaxCpu field.
func (o *ServicePlan) SetMaxCpu(v float32) {
	o.MaxCpu.Set(&v)
}
// SetMaxCpuNil sets the value for MaxCpu to be an explicit nil
func (o *ServicePlan) SetMaxCpuNil() {
	o.MaxCpu.Set(nil)
}

// UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
func (o *ServicePlan) UnsetMaxCpu() {
	o.MaxCpu.Unset()
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetMaxCores() float32 {
	if o == nil || o.MaxCores.Get() == nil {
		var ret float32
		return ret
	}
	return *o.MaxCores.Get()
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetMaxCoresOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxCores.Get(), o.MaxCores.IsSet()
}

// HasMaxCores returns a boolean if a field has been set.
func (o *ServicePlan) HasMaxCores() bool {
	if o != nil && o.MaxCores.IsSet() {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given NullableFloat32 and assigns it to the MaxCores field.
func (o *ServicePlan) SetMaxCores(v float32) {
	o.MaxCores.Set(&v)
}
// SetMaxCoresNil sets the value for MaxCores to be an explicit nil
func (o *ServicePlan) SetMaxCoresNil() {
	o.MaxCores.Set(nil)
}

// UnsetMaxCores ensures that no value is present for MaxCores, not even an explicit nil
func (o *ServicePlan) UnsetMaxCores() {
	o.MaxCores.Unset()
}

// GetMaxDisks returns the MaxDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetMaxDisks() float32 {
	if o == nil || o.MaxDisks.Get() == nil {
		var ret float32
		return ret
	}
	return *o.MaxDisks.Get()
}

// GetMaxDisksOk returns a tuple with the MaxDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetMaxDisksOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxDisks.Get(), o.MaxDisks.IsSet()
}

// HasMaxDisks returns a boolean if a field has been set.
func (o *ServicePlan) HasMaxDisks() bool {
	if o != nil && o.MaxDisks.IsSet() {
		return true
	}

	return false
}

// SetMaxDisks gets a reference to the given NullableFloat32 and assigns it to the MaxDisks field.
func (o *ServicePlan) SetMaxDisks(v float32) {
	o.MaxDisks.Set(&v)
}
// SetMaxDisksNil sets the value for MaxDisks to be an explicit nil
func (o *ServicePlan) SetMaxDisksNil() {
	o.MaxDisks.Set(nil)
}

// UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
func (o *ServicePlan) UnsetMaxDisks() {
	o.MaxDisks.Unset()
}

// GetCoresPerSocket returns the CoresPerSocket field value if set, zero value otherwise.
func (o *ServicePlan) GetCoresPerSocket() float32 {
	if o == nil || o.CoresPerSocket == nil {
		var ret float32
		return ret
	}
	return *o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetCoresPerSocketOk() (*float32, bool) {
	if o == nil || o.CoresPerSocket == nil {
		return nil, false
	}
	return o.CoresPerSocket, true
}

// HasCoresPerSocket returns a boolean if a field has been set.
func (o *ServicePlan) HasCoresPerSocket() bool {
	if o != nil && o.CoresPerSocket != nil {
		return true
	}

	return false
}

// SetCoresPerSocket gets a reference to the given float32 and assigns it to the CoresPerSocket field.
func (o *ServicePlan) SetCoresPerSocket(v float32) {
	o.CoresPerSocket = &v
}

// GetCustomCpu returns the CustomCpu field value if set, zero value otherwise.
func (o *ServicePlan) GetCustomCpu() bool {
	if o == nil || o.CustomCpu == nil {
		var ret bool
		return ret
	}
	return *o.CustomCpu
}

// GetCustomCpuOk returns a tuple with the CustomCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetCustomCpuOk() (*bool, bool) {
	if o == nil || o.CustomCpu == nil {
		return nil, false
	}
	return o.CustomCpu, true
}

// HasCustomCpu returns a boolean if a field has been set.
func (o *ServicePlan) HasCustomCpu() bool {
	if o != nil && o.CustomCpu != nil {
		return true
	}

	return false
}

// SetCustomCpu gets a reference to the given bool and assigns it to the CustomCpu field.
func (o *ServicePlan) SetCustomCpu(v bool) {
	o.CustomCpu = &v
}

// GetCustomCores returns the CustomCores field value if set, zero value otherwise.
func (o *ServicePlan) GetCustomCores() bool {
	if o == nil || o.CustomCores == nil {
		var ret bool
		return ret
	}
	return *o.CustomCores
}

// GetCustomCoresOk returns a tuple with the CustomCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetCustomCoresOk() (*bool, bool) {
	if o == nil || o.CustomCores == nil {
		return nil, false
	}
	return o.CustomCores, true
}

// HasCustomCores returns a boolean if a field has been set.
func (o *ServicePlan) HasCustomCores() bool {
	if o != nil && o.CustomCores != nil {
		return true
	}

	return false
}

// SetCustomCores gets a reference to the given bool and assigns it to the CustomCores field.
func (o *ServicePlan) SetCustomCores(v bool) {
	o.CustomCores = &v
}

// GetCustomMaxStorage returns the CustomMaxStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetCustomMaxStorage() bool {
	if o == nil || o.CustomMaxStorage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxStorage.Get()
}

// GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetCustomMaxStorageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomMaxStorage.Get(), o.CustomMaxStorage.IsSet()
}

// HasCustomMaxStorage returns a boolean if a field has been set.
func (o *ServicePlan) HasCustomMaxStorage() bool {
	if o != nil && o.CustomMaxStorage.IsSet() {
		return true
	}

	return false
}

// SetCustomMaxStorage gets a reference to the given NullableBool and assigns it to the CustomMaxStorage field.
func (o *ServicePlan) SetCustomMaxStorage(v bool) {
	o.CustomMaxStorage.Set(&v)
}
// SetCustomMaxStorageNil sets the value for CustomMaxStorage to be an explicit nil
func (o *ServicePlan) SetCustomMaxStorageNil() {
	o.CustomMaxStorage.Set(nil)
}

// UnsetCustomMaxStorage ensures that no value is present for CustomMaxStorage, not even an explicit nil
func (o *ServicePlan) UnsetCustomMaxStorage() {
	o.CustomMaxStorage.Unset()
}

// GetCustomMaxDataStorage returns the CustomMaxDataStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetCustomMaxDataStorage() bool {
	if o == nil || o.CustomMaxDataStorage.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxDataStorage.Get()
}

// GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetCustomMaxDataStorageOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomMaxDataStorage.Get(), o.CustomMaxDataStorage.IsSet()
}

// HasCustomMaxDataStorage returns a boolean if a field has been set.
func (o *ServicePlan) HasCustomMaxDataStorage() bool {
	if o != nil && o.CustomMaxDataStorage.IsSet() {
		return true
	}

	return false
}

// SetCustomMaxDataStorage gets a reference to the given NullableBool and assigns it to the CustomMaxDataStorage field.
func (o *ServicePlan) SetCustomMaxDataStorage(v bool) {
	o.CustomMaxDataStorage.Set(&v)
}
// SetCustomMaxDataStorageNil sets the value for CustomMaxDataStorage to be an explicit nil
func (o *ServicePlan) SetCustomMaxDataStorageNil() {
	o.CustomMaxDataStorage.Set(nil)
}

// UnsetCustomMaxDataStorage ensures that no value is present for CustomMaxDataStorage, not even an explicit nil
func (o *ServicePlan) UnsetCustomMaxDataStorage() {
	o.CustomMaxDataStorage.Unset()
}

// GetCustomMaxMemory returns the CustomMaxMemory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetCustomMaxMemory() bool {
	if o == nil || o.CustomMaxMemory.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxMemory.Get()
}

// GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetCustomMaxMemoryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomMaxMemory.Get(), o.CustomMaxMemory.IsSet()
}

// HasCustomMaxMemory returns a boolean if a field has been set.
func (o *ServicePlan) HasCustomMaxMemory() bool {
	if o != nil && o.CustomMaxMemory.IsSet() {
		return true
	}

	return false
}

// SetCustomMaxMemory gets a reference to the given NullableBool and assigns it to the CustomMaxMemory field.
func (o *ServicePlan) SetCustomMaxMemory(v bool) {
	o.CustomMaxMemory.Set(&v)
}
// SetCustomMaxMemoryNil sets the value for CustomMaxMemory to be an explicit nil
func (o *ServicePlan) SetCustomMaxMemoryNil() {
	o.CustomMaxMemory.Set(nil)
}

// UnsetCustomMaxMemory ensures that no value is present for CustomMaxMemory, not even an explicit nil
func (o *ServicePlan) UnsetCustomMaxMemory() {
	o.CustomMaxMemory.Unset()
}

// GetAddVolumes returns the AddVolumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetAddVolumes() bool {
	if o == nil || o.AddVolumes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AddVolumes.Get()
}

// GetAddVolumesOk returns a tuple with the AddVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetAddVolumesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddVolumes.Get(), o.AddVolumes.IsSet()
}

// HasAddVolumes returns a boolean if a field has been set.
func (o *ServicePlan) HasAddVolumes() bool {
	if o != nil && o.AddVolumes.IsSet() {
		return true
	}

	return false
}

// SetAddVolumes gets a reference to the given NullableBool and assigns it to the AddVolumes field.
func (o *ServicePlan) SetAddVolumes(v bool) {
	o.AddVolumes.Set(&v)
}
// SetAddVolumesNil sets the value for AddVolumes to be an explicit nil
func (o *ServicePlan) SetAddVolumesNil() {
	o.AddVolumes.Set(nil)
}

// UnsetAddVolumes ensures that no value is present for AddVolumes, not even an explicit nil
func (o *ServicePlan) UnsetAddVolumes() {
	o.AddVolumes.Unset()
}

// GetMemoryOptionSource returns the MemoryOptionSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetMemoryOptionSource() string {
	if o == nil || o.MemoryOptionSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryOptionSource.Get()
}

// GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetMemoryOptionSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryOptionSource.Get(), o.MemoryOptionSource.IsSet()
}

// HasMemoryOptionSource returns a boolean if a field has been set.
func (o *ServicePlan) HasMemoryOptionSource() bool {
	if o != nil && o.MemoryOptionSource.IsSet() {
		return true
	}

	return false
}

// SetMemoryOptionSource gets a reference to the given NullableString and assigns it to the MemoryOptionSource field.
func (o *ServicePlan) SetMemoryOptionSource(v string) {
	o.MemoryOptionSource.Set(&v)
}
// SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil
func (o *ServicePlan) SetMemoryOptionSourceNil() {
	o.MemoryOptionSource.Set(nil)
}

// UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
func (o *ServicePlan) UnsetMemoryOptionSource() {
	o.MemoryOptionSource.Unset()
}

// GetCpuOptionSource returns the CpuOptionSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetCpuOptionSource() string {
	if o == nil || o.CpuOptionSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuOptionSource.Get()
}

// GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetCpuOptionSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuOptionSource.Get(), o.CpuOptionSource.IsSet()
}

// HasCpuOptionSource returns a boolean if a field has been set.
func (o *ServicePlan) HasCpuOptionSource() bool {
	if o != nil && o.CpuOptionSource.IsSet() {
		return true
	}

	return false
}

// SetCpuOptionSource gets a reference to the given NullableString and assigns it to the CpuOptionSource field.
func (o *ServicePlan) SetCpuOptionSource(v string) {
	o.CpuOptionSource.Set(&v)
}
// SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil
func (o *ServicePlan) SetCpuOptionSourceNil() {
	o.CpuOptionSource.Set(nil)
}

// UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
func (o *ServicePlan) UnsetCpuOptionSource() {
	o.CpuOptionSource.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ServicePlan) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ServicePlan) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ServicePlan) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ServicePlan) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ServicePlan) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ServicePlan) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetRegionCode() string {
	if o == nil || o.RegionCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RegionCode.Get()
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetRegionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RegionCode.Get(), o.RegionCode.IsSet()
}

// HasRegionCode returns a boolean if a field has been set.
func (o *ServicePlan) HasRegionCode() bool {
	if o != nil && o.RegionCode.IsSet() {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given NullableString and assigns it to the RegionCode field.
func (o *ServicePlan) SetRegionCode(v string) {
	o.RegionCode.Set(&v)
}
// SetRegionCodeNil sets the value for RegionCode to be an explicit nil
func (o *ServicePlan) SetRegionCodeNil() {
	o.RegionCode.Set(nil)
}

// UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
func (o *ServicePlan) UnsetRegionCode() {
	o.RegionCode.Unset()
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ServicePlan) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ServicePlan) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ServicePlan) SetVisibility(v string) {
	o.Visibility = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *ServicePlan) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *ServicePlan) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *ServicePlan) SetEditable(v bool) {
	o.Editable = &v
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *ServicePlan) GetProvisionType() GuidanceVmwareSizingPlanBeforeActionProvisionType {
	if o == nil || o.ProvisionType == nil {
		var ret GuidanceVmwareSizingPlanBeforeActionProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetProvisionTypeOk() (*GuidanceVmwareSizingPlanBeforeActionProvisionType, bool) {
	if o == nil || o.ProvisionType == nil {
		return nil, false
	}
	return o.ProvisionType, true
}

// HasProvisionType returns a boolean if a field has been set.
func (o *ServicePlan) HasProvisionType() bool {
	if o != nil && o.ProvisionType != nil {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given GuidanceVmwareSizingPlanBeforeActionProvisionType and assigns it to the ProvisionType field.
func (o *ServicePlan) SetProvisionType(v GuidanceVmwareSizingPlanBeforeActionProvisionType) {
	o.ProvisionType = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ServicePlan) GetTenants() string {
	if o == nil || o.Tenants == nil {
		var ret string
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetTenantsOk() (*string, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ServicePlan) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given string and assigns it to the Tenants field.
func (o *ServicePlan) SetTenants(v string) {
	o.Tenants = &v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetPriceSets() []GuidanceVmwareSizingPlanBeforeActionPriceSets {
	if o == nil  {
		var ret []GuidanceVmwareSizingPlanBeforeActionPriceSets
		return ret
	}
	return o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetPriceSetsOk() (*[]GuidanceVmwareSizingPlanBeforeActionPriceSets, bool) {
	if o == nil || o.PriceSets == nil {
		return nil, false
	}
	return &o.PriceSets, true
}

// HasPriceSets returns a boolean if a field has been set.
func (o *ServicePlan) HasPriceSets() bool {
	if o != nil && o.PriceSets != nil {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []GuidanceVmwareSizingPlanBeforeActionPriceSets and assigns it to the PriceSets field.
func (o *ServicePlan) SetPriceSets(v []GuidanceVmwareSizingPlanBeforeActionPriceSets) {
	o.PriceSets = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServicePlan) GetConfig() ServicePlanConfig {
	if o == nil || o.Config.Get() == nil {
		var ret ServicePlanConfig
		return ret
	}
	return *o.Config.Get()
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServicePlan) GetConfigOk() (*ServicePlanConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Config.Get(), o.Config.IsSet()
}

// HasConfig returns a boolean if a field has been set.
func (o *ServicePlan) HasConfig() bool {
	if o != nil && o.Config.IsSet() {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NullableServicePlanConfig and assigns it to the Config field.
func (o *ServicePlan) SetConfig(v ServicePlanConfig) {
	o.Config.Set(&v)
}
// SetConfigNil sets the value for Config to be an explicit nil
func (o *ServicePlan) SetConfigNil() {
	o.Config.Set(nil)
}

// UnsetConfig ensures that no value is present for Config, not even an explicit nil
func (o *ServicePlan) UnsetConfig() {
	o.Config.Unset()
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *ServicePlan) GetZones() []InlineResponse20094Network {
	if o == nil || o.Zones == nil {
		var ret []InlineResponse20094Network
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetZonesOk() (*[]InlineResponse20094Network, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *ServicePlan) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given []InlineResponse20094Network and assigns it to the Zones field.
func (o *ServicePlan) SetZones(v []InlineResponse20094Network) {
	o.Zones = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ServicePlan) GetPermissions() ServicePlanPermissions {
	if o == nil || o.Permissions == nil {
		var ret ServicePlanPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlan) GetPermissionsOk() (*ServicePlanPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ServicePlan) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ServicePlanPermissions and assigns it to the Permissions field.
func (o *ServicePlan) SetPermissions(v ServicePlanPermissions) {
	o.Permissions = &v
}

func (o ServicePlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.MaxStorage != nil {
		toSerialize["maxStorage"] = o.MaxStorage
	}
	if o.MaxMemory != nil {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	if o.MaxCpu.IsSet() {
		toSerialize["maxCpu"] = o.MaxCpu.Get()
	}
	if o.MaxCores.IsSet() {
		toSerialize["maxCores"] = o.MaxCores.Get()
	}
	if o.MaxDisks.IsSet() {
		toSerialize["maxDisks"] = o.MaxDisks.Get()
	}
	if o.CoresPerSocket != nil {
		toSerialize["coresPerSocket"] = o.CoresPerSocket
	}
	if o.CustomCpu != nil {
		toSerialize["customCpu"] = o.CustomCpu
	}
	if o.CustomCores != nil {
		toSerialize["customCores"] = o.CustomCores
	}
	if o.CustomMaxStorage.IsSet() {
		toSerialize["customMaxStorage"] = o.CustomMaxStorage.Get()
	}
	if o.CustomMaxDataStorage.IsSet() {
		toSerialize["customMaxDataStorage"] = o.CustomMaxDataStorage.Get()
	}
	if o.CustomMaxMemory.IsSet() {
		toSerialize["customMaxMemory"] = o.CustomMaxMemory.Get()
	}
	if o.AddVolumes.IsSet() {
		toSerialize["addVolumes"] = o.AddVolumes.Get()
	}
	if o.MemoryOptionSource.IsSet() {
		toSerialize["memoryOptionSource"] = o.MemoryOptionSource.Get()
	}
	if o.CpuOptionSource.IsSet() {
		toSerialize["cpuOptionSource"] = o.CpuOptionSource.Get()
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.RegionCode.IsSet() {
		toSerialize["regionCode"] = o.RegionCode.Get()
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.Editable != nil {
		toSerialize["editable"] = o.Editable
	}
	if o.ProvisionType != nil {
		toSerialize["provisionType"] = o.ProvisionType
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	if o.PriceSets != nil {
		toSerialize["priceSets"] = o.PriceSets
	}
	if o.Config.IsSet() {
		toSerialize["config"] = o.Config.Get()
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableServicePlan struct {
	value *ServicePlan
	isSet bool
}

func (v NullableServicePlan) Get() *ServicePlan {
	return v.value
}

func (v *NullableServicePlan) Set(val *ServicePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlan(val *ServicePlan) *NullableServicePlan {
	return &NullableServicePlan{value: val, isSet: true}
}

func (v NullableServicePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


