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

// GuidanceVmwareSizingPlanBeforeAction struct for GuidanceVmwareSizingPlanBeforeAction
type GuidanceVmwareSizingPlanBeforeAction struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Active *bool `json:"active,omitempty"`
	SortOrder *int64 `json:"sortOrder,omitempty"`
	Description NullableString `json:"description,omitempty"`
	MaxStorage *int64 `json:"maxStorage,omitempty"`
	MaxMemory *int64 `json:"maxMemory,omitempty"`
	MaxCpu NullableString `json:"maxCpu,omitempty"`
	MaxCores *int64 `json:"maxCores,omitempty"`
	MaxDisks NullableString `json:"maxDisks,omitempty"`
	CoresPerSocket *int64 `json:"coresPerSocket,omitempty"`
	CustomCpu *bool `json:"customCpu,omitempty"`
	CustomCores *bool `json:"customCores,omitempty"`
	CustomMaxStorage *bool `json:"customMaxStorage,omitempty"`
	CustomMaxDataStorage *bool `json:"customMaxDataStorage,omitempty"`
	CustomMaxMemory *bool `json:"customMaxMemory,omitempty"`
	AddVolumes *bool `json:"addVolumes,omitempty"`
	MemoryOptionSource NullableString `json:"memoryOptionSource,omitempty"`
	CpuOptionSource NullableString `json:"cpuOptionSource,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	RegionCode NullableString `json:"regionCode,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	ProvisionType *GuidanceVmwareSizingPlanBeforeActionProvisionType `json:"provisionType,omitempty"`
	Tenants *string `json:"tenants,omitempty"`
	PriceSets *[]GuidanceVmwareSizingPlanBeforeActionPriceSets `json:"priceSets,omitempty"`
	Config *GuidanceVmwareSizingPlanBeforeActionConfig `json:"config,omitempty"`
}

// NewGuidanceVmwareSizingPlanBeforeAction instantiates a new GuidanceVmwareSizingPlanBeforeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceVmwareSizingPlanBeforeAction() *GuidanceVmwareSizingPlanBeforeAction {
	this := GuidanceVmwareSizingPlanBeforeAction{}
	return &this
}

// NewGuidanceVmwareSizingPlanBeforeActionWithDefaults instantiates a new GuidanceVmwareSizingPlanBeforeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceVmwareSizingPlanBeforeActionWithDefaults() *GuidanceVmwareSizingPlanBeforeAction {
	this := GuidanceVmwareSizingPlanBeforeAction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCode(v string) {
	o.Code = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetActive(v bool) {
	o.Active = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetSortOrder() int64 {
	if o == nil || o.SortOrder == nil {
		var ret int64
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetSortOrderOk() (*int64, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int64 and assigns it to the SortOrder field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetSortOrder(v int64) {
	o.SortOrder = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetDescription() {
	o.Description.Unset()
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxStorage() int64 {
	if o == nil || o.MaxStorage == nil {
		var ret int64
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxStorageOk() (*int64, bool) {
	if o == nil || o.MaxStorage == nil {
		return nil, false
	}
	return o.MaxStorage, true
}

// HasMaxStorage returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxStorage() bool {
	if o != nil && o.MaxStorage != nil {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given int64 and assigns it to the MaxStorage field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxStorage(v int64) {
	o.MaxStorage = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxMemory() int64 {
	if o == nil || o.MaxMemory == nil {
		var ret int64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxMemoryOk() (*int64, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int64 and assigns it to the MaxMemory field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxMemory(v int64) {
	o.MaxMemory = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCpu() string {
	if o == nil || o.MaxCpu.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxCpu.Get()
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCpuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxCpu.Get(), o.MaxCpu.IsSet()
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxCpu() bool {
	if o != nil && o.MaxCpu.IsSet() {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given NullableString and assigns it to the MaxCpu field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCpu(v string) {
	o.MaxCpu.Set(&v)
}
// SetMaxCpuNil sets the value for MaxCpu to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCpuNil() {
	o.MaxCpu.Set(nil)
}

// UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMaxCpu() {
	o.MaxCpu.Unset()
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCores() int64 {
	if o == nil || o.MaxCores == nil {
		var ret int64
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCoresOk() (*int64, bool) {
	if o == nil || o.MaxCores == nil {
		return nil, false
	}
	return o.MaxCores, true
}

// HasMaxCores returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxCores() bool {
	if o != nil && o.MaxCores != nil {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given int64 and assigns it to the MaxCores field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCores(v int64) {
	o.MaxCores = &v
}

// GetMaxDisks returns the MaxDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxDisks() string {
	if o == nil || o.MaxDisks.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxDisks.Get()
}

// GetMaxDisksOk returns a tuple with the MaxDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxDisksOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxDisks.Get(), o.MaxDisks.IsSet()
}

// HasMaxDisks returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxDisks() bool {
	if o != nil && o.MaxDisks.IsSet() {
		return true
	}

	return false
}

// SetMaxDisks gets a reference to the given NullableString and assigns it to the MaxDisks field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxDisks(v string) {
	o.MaxDisks.Set(&v)
}
// SetMaxDisksNil sets the value for MaxDisks to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxDisksNil() {
	o.MaxDisks.Set(nil)
}

// UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMaxDisks() {
	o.MaxDisks.Unset()
}

// GetCoresPerSocket returns the CoresPerSocket field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCoresPerSocket() int64 {
	if o == nil || o.CoresPerSocket == nil {
		var ret int64
		return ret
	}
	return *o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil || o.CoresPerSocket == nil {
		return nil, false
	}
	return o.CoresPerSocket, true
}

// HasCoresPerSocket returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCoresPerSocket() bool {
	if o != nil && o.CoresPerSocket != nil {
		return true
	}

	return false
}

// SetCoresPerSocket gets a reference to the given int64 and assigns it to the CoresPerSocket field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = &v
}

// GetCustomCpu returns the CustomCpu field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCpu() bool {
	if o == nil || o.CustomCpu == nil {
		var ret bool
		return ret
	}
	return *o.CustomCpu
}

// GetCustomCpuOk returns a tuple with the CustomCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCpuOk() (*bool, bool) {
	if o == nil || o.CustomCpu == nil {
		return nil, false
	}
	return o.CustomCpu, true
}

// HasCustomCpu returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomCpu() bool {
	if o != nil && o.CustomCpu != nil {
		return true
	}

	return false
}

// SetCustomCpu gets a reference to the given bool and assigns it to the CustomCpu field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomCpu(v bool) {
	o.CustomCpu = &v
}

// GetCustomCores returns the CustomCores field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCores() bool {
	if o == nil || o.CustomCores == nil {
		var ret bool
		return ret
	}
	return *o.CustomCores
}

// GetCustomCoresOk returns a tuple with the CustomCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCoresOk() (*bool, bool) {
	if o == nil || o.CustomCores == nil {
		return nil, false
	}
	return o.CustomCores, true
}

// HasCustomCores returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomCores() bool {
	if o != nil && o.CustomCores != nil {
		return true
	}

	return false
}

// SetCustomCores gets a reference to the given bool and assigns it to the CustomCores field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomCores(v bool) {
	o.CustomCores = &v
}

// GetCustomMaxStorage returns the CustomMaxStorage field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxStorage() bool {
	if o == nil || o.CustomMaxStorage == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxStorage
}

// GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxStorageOk() (*bool, bool) {
	if o == nil || o.CustomMaxStorage == nil {
		return nil, false
	}
	return o.CustomMaxStorage, true
}

// HasCustomMaxStorage returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxStorage() bool {
	if o != nil && o.CustomMaxStorage != nil {
		return true
	}

	return false
}

// SetCustomMaxStorage gets a reference to the given bool and assigns it to the CustomMaxStorage field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxStorage(v bool) {
	o.CustomMaxStorage = &v
}

// GetCustomMaxDataStorage returns the CustomMaxDataStorage field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxDataStorage() bool {
	if o == nil || o.CustomMaxDataStorage == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxDataStorage
}

// GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxDataStorageOk() (*bool, bool) {
	if o == nil || o.CustomMaxDataStorage == nil {
		return nil, false
	}
	return o.CustomMaxDataStorage, true
}

// HasCustomMaxDataStorage returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxDataStorage() bool {
	if o != nil && o.CustomMaxDataStorage != nil {
		return true
	}

	return false
}

// SetCustomMaxDataStorage gets a reference to the given bool and assigns it to the CustomMaxDataStorage field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxDataStorage(v bool) {
	o.CustomMaxDataStorage = &v
}

// GetCustomMaxMemory returns the CustomMaxMemory field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxMemory() bool {
	if o == nil || o.CustomMaxMemory == nil {
		var ret bool
		return ret
	}
	return *o.CustomMaxMemory
}

// GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxMemoryOk() (*bool, bool) {
	if o == nil || o.CustomMaxMemory == nil {
		return nil, false
	}
	return o.CustomMaxMemory, true
}

// HasCustomMaxMemory returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxMemory() bool {
	if o != nil && o.CustomMaxMemory != nil {
		return true
	}

	return false
}

// SetCustomMaxMemory gets a reference to the given bool and assigns it to the CustomMaxMemory field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxMemory(v bool) {
	o.CustomMaxMemory = &v
}

// GetAddVolumes returns the AddVolumes field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetAddVolumes() bool {
	if o == nil || o.AddVolumes == nil {
		var ret bool
		return ret
	}
	return *o.AddVolumes
}

// GetAddVolumesOk returns a tuple with the AddVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetAddVolumesOk() (*bool, bool) {
	if o == nil || o.AddVolumes == nil {
		return nil, false
	}
	return o.AddVolumes, true
}

// HasAddVolumes returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasAddVolumes() bool {
	if o != nil && o.AddVolumes != nil {
		return true
	}

	return false
}

// SetAddVolumes gets a reference to the given bool and assigns it to the AddVolumes field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetAddVolumes(v bool) {
	o.AddVolumes = &v
}

// GetMemoryOptionSource returns the MemoryOptionSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMemoryOptionSource() string {
	if o == nil || o.MemoryOptionSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.MemoryOptionSource.Get()
}

// GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetMemoryOptionSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MemoryOptionSource.Get(), o.MemoryOptionSource.IsSet()
}

// HasMemoryOptionSource returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasMemoryOptionSource() bool {
	if o != nil && o.MemoryOptionSource.IsSet() {
		return true
	}

	return false
}

// SetMemoryOptionSource gets a reference to the given NullableString and assigns it to the MemoryOptionSource field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMemoryOptionSource(v string) {
	o.MemoryOptionSource.Set(&v)
}
// SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetMemoryOptionSourceNil() {
	o.MemoryOptionSource.Set(nil)
}

// UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMemoryOptionSource() {
	o.MemoryOptionSource.Unset()
}

// GetCpuOptionSource returns the CpuOptionSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCpuOptionSource() string {
	if o == nil || o.CpuOptionSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuOptionSource.Get()
}

// GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetCpuOptionSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CpuOptionSource.Get(), o.CpuOptionSource.IsSet()
}

// HasCpuOptionSource returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasCpuOptionSource() bool {
	if o != nil && o.CpuOptionSource.IsSet() {
		return true
	}

	return false
}

// SetCpuOptionSource gets a reference to the given NullableString and assigns it to the CpuOptionSource field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCpuOptionSource(v string) {
	o.CpuOptionSource.Set(&v)
}
// SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetCpuOptionSourceNil() {
	o.CpuOptionSource.Set(nil)
}

// UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetCpuOptionSource() {
	o.CpuOptionSource.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingPlanBeforeAction) GetRegionCode() string {
	if o == nil || o.RegionCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.RegionCode.Get()
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingPlanBeforeAction) GetRegionCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RegionCode.Get(), o.RegionCode.IsSet()
}

// HasRegionCode returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasRegionCode() bool {
	if o != nil && o.RegionCode.IsSet() {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given NullableString and assigns it to the RegionCode field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetRegionCode(v string) {
	o.RegionCode.Set(&v)
}
// SetRegionCodeNil sets the value for RegionCode to be an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) SetRegionCodeNil() {
	o.RegionCode.Set(nil)
}

// UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetRegionCode() {
	o.RegionCode.Unset()
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetVisibility(v string) {
	o.Visibility = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetEditable() bool {
	if o == nil || o.Editable == nil {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetEditableOk() (*bool, bool) {
	if o == nil || o.Editable == nil {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasEditable() bool {
	if o != nil && o.Editable != nil {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetEditable(v bool) {
	o.Editable = &v
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetProvisionType() GuidanceVmwareSizingPlanBeforeActionProvisionType {
	if o == nil || o.ProvisionType == nil {
		var ret GuidanceVmwareSizingPlanBeforeActionProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetProvisionTypeOk() (*GuidanceVmwareSizingPlanBeforeActionProvisionType, bool) {
	if o == nil || o.ProvisionType == nil {
		return nil, false
	}
	return o.ProvisionType, true
}

// HasProvisionType returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasProvisionType() bool {
	if o != nil && o.ProvisionType != nil {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given GuidanceVmwareSizingPlanBeforeActionProvisionType and assigns it to the ProvisionType field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetProvisionType(v GuidanceVmwareSizingPlanBeforeActionProvisionType) {
	o.ProvisionType = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetTenants() string {
	if o == nil || o.Tenants == nil {
		var ret string
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetTenantsOk() (*string, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given string and assigns it to the Tenants field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetTenants(v string) {
	o.Tenants = &v
}

// GetPriceSets returns the PriceSets field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetPriceSets() []GuidanceVmwareSizingPlanBeforeActionPriceSets {
	if o == nil || o.PriceSets == nil {
		var ret []GuidanceVmwareSizingPlanBeforeActionPriceSets
		return ret
	}
	return *o.PriceSets
}

// GetPriceSetsOk returns a tuple with the PriceSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetPriceSetsOk() (*[]GuidanceVmwareSizingPlanBeforeActionPriceSets, bool) {
	if o == nil || o.PriceSets == nil {
		return nil, false
	}
	return o.PriceSets, true
}

// HasPriceSets returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasPriceSets() bool {
	if o != nil && o.PriceSets != nil {
		return true
	}

	return false
}

// SetPriceSets gets a reference to the given []GuidanceVmwareSizingPlanBeforeActionPriceSets and assigns it to the PriceSets field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetPriceSets(v []GuidanceVmwareSizingPlanBeforeActionPriceSets) {
	o.PriceSets = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetConfig() GuidanceVmwareSizingPlanBeforeActionConfig {
	if o == nil || o.Config == nil {
		var ret GuidanceVmwareSizingPlanBeforeActionConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) GetConfigOk() (*GuidanceVmwareSizingPlanBeforeActionConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingPlanBeforeAction) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given GuidanceVmwareSizingPlanBeforeActionConfig and assigns it to the Config field.
func (o *GuidanceVmwareSizingPlanBeforeAction) SetConfig(v GuidanceVmwareSizingPlanBeforeActionConfig) {
	o.Config = &v
}

func (o GuidanceVmwareSizingPlanBeforeAction) MarshalJSON() ([]byte, error) {
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
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
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
	if o.MaxCores != nil {
		toSerialize["maxCores"] = o.MaxCores
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
	if o.CustomMaxStorage != nil {
		toSerialize["customMaxStorage"] = o.CustomMaxStorage
	}
	if o.CustomMaxDataStorage != nil {
		toSerialize["customMaxDataStorage"] = o.CustomMaxDataStorage
	}
	if o.CustomMaxMemory != nil {
		toSerialize["customMaxMemory"] = o.CustomMaxMemory
	}
	if o.AddVolumes != nil {
		toSerialize["addVolumes"] = o.AddVolumes
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
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceVmwareSizingPlanBeforeAction struct {
	value *GuidanceVmwareSizingPlanBeforeAction
	isSet bool
}

func (v NullableGuidanceVmwareSizingPlanBeforeAction) Get() *GuidanceVmwareSizingPlanBeforeAction {
	return v.value
}

func (v *NullableGuidanceVmwareSizingPlanBeforeAction) Set(val *GuidanceVmwareSizingPlanBeforeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceVmwareSizingPlanBeforeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceVmwareSizingPlanBeforeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceVmwareSizingPlanBeforeAction(val *GuidanceVmwareSizingPlanBeforeAction) *NullableGuidanceVmwareSizingPlanBeforeAction {
	return &NullableGuidanceVmwareSizingPlanBeforeAction{value: val, isSet: true}
}

func (v NullableGuidanceVmwareSizingPlanBeforeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceVmwareSizingPlanBeforeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


