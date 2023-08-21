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

// ApiVdiPoolsIdVdiPool struct for ApiVdiPoolsIdVdiPool
type ApiVdiPoolsIdVdiPool struct {
	// Virtual Desktop name
	Name *string `json:"name,omitempty"`
	// Virtual Desktop description
	Description *string `json:"description,omitempty"`
	// Owner (User) ID
	Owner *int64 `json:"owner,omitempty"`
	// Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby. 
	MinIdle *float32 `json:"minIdle,omitempty"`
	// The initial size of the pool to be allocated on creation
	InitialPoolSize *float32 `json:"initialPoolSize,omitempty"`
	// Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances 
	MaxIdle *float32 `json:"maxIdle,omitempty"`
	// Max limit on number of allocations and instances within the pool. 
	MaxPoolSize *float32 `json:"maxPoolSize,omitempty"`
	// Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence. 
	AllocationTimeoutMinutes *float32 `json:"allocationTimeoutMinutes,omitempty"`
	// Persistent Desktop Pool
	PersistentUser *bool `json:"persistentUser,omitempty"`
	// Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD)
	Recyclable *bool `json:"recyclable,omitempty"`
	// Allow copy from desktop
	AllowCopy *bool `json:"allowCopy,omitempty"`
	// Allow local printers from Desktop
	AllowPrinter *bool `json:"allowPrinter,omitempty"`
	// Allow File Share
	AllowFileshare *bool `json:"allowFileshare,omitempty"`
	// Allow hypervisor console
	AllowHypervisorConsole *bool `json:"allowHypervisorConsole,omitempty"`
	// Auto create local user upon reservation
	AutoCreateLocalUserOnReservation *bool `json:"autoCreateLocalUserOnReservation,omitempty"`
	// Can be used to enable or disable the VDI pool
	Enabled *bool `json:"enabled,omitempty"`
	// The relative location of an icon image
	IconPath *string `json:"iconPath,omitempty"`
	// Array of VDI App IDs
	Apps *[]int64 `json:"apps,omitempty"`
	// VDI Gateway ID
	Gateway *int64 `json:"gateway,omitempty"`
	// Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values.
	InstanceConfig *string `json:"instanceConfig,omitempty"`
	Config *ApiVdiPoolsIdVdiPoolConfig `json:"config,omitempty"`
	// Guest Console Jump Host
	GuestConsoleJumpHost *string `json:"guestConsoleJumpHost,omitempty"`
	// Guest Console Jump Port
	GuestConsoleJumpPort *int64 `json:"guestConsoleJumpPort,omitempty"`
	// Guest Console Jump Username
	GuestConsoleJumpUsername *string `json:"guestConsoleJumpUsername,omitempty"`
	// Guest Console Jump Password
	GuestConsoleJumpPassword *string `json:"guestConsoleJumpPassword,omitempty"`
	// Guest Console Jump Key Pair. see `Key Pair`
	GuestConsoleJumpKeypair *int64 `json:"guestConsoleJumpKeypair,omitempty"`
}

// NewApiVdiPoolsIdVdiPool instantiates a new ApiVdiPoolsIdVdiPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiVdiPoolsIdVdiPool() *ApiVdiPoolsIdVdiPool {
	this := ApiVdiPoolsIdVdiPool{}
	var persistentUser bool = false
	this.PersistentUser = &persistentUser
	var recyclable bool = false
	this.Recyclable = &recyclable
	var allowCopy bool = false
	this.AllowCopy = &allowCopy
	var allowPrinter bool = false
	this.AllowPrinter = &allowPrinter
	var allowFileshare bool = false
	this.AllowFileshare = &allowFileshare
	var allowHypervisorConsole bool = false
	this.AllowHypervisorConsole = &allowHypervisorConsole
	var autoCreateLocalUserOnReservation bool = false
	this.AutoCreateLocalUserOnReservation = &autoCreateLocalUserOnReservation
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewApiVdiPoolsIdVdiPoolWithDefaults instantiates a new ApiVdiPoolsIdVdiPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiVdiPoolsIdVdiPoolWithDefaults() *ApiVdiPoolsIdVdiPool {
	this := ApiVdiPoolsIdVdiPool{}
	var persistentUser bool = false
	this.PersistentUser = &persistentUser
	var recyclable bool = false
	this.Recyclable = &recyclable
	var allowCopy bool = false
	this.AllowCopy = &allowCopy
	var allowPrinter bool = false
	this.AllowPrinter = &allowPrinter
	var allowFileshare bool = false
	this.AllowFileshare = &allowFileshare
	var allowHypervisorConsole bool = false
	this.AllowHypervisorConsole = &allowHypervisorConsole
	var autoCreateLocalUserOnReservation bool = false
	this.AutoCreateLocalUserOnReservation = &autoCreateLocalUserOnReservation
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiVdiPoolsIdVdiPool) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiVdiPoolsIdVdiPool) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetOwner() int64 {
	if o == nil || o.Owner == nil {
		var ret int64
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetOwnerOk() (*int64, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given int64 and assigns it to the Owner field.
func (o *ApiVdiPoolsIdVdiPool) SetOwner(v int64) {
	o.Owner = &v
}

// GetMinIdle returns the MinIdle field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetMinIdle() float32 {
	if o == nil || o.MinIdle == nil {
		var ret float32
		return ret
	}
	return *o.MinIdle
}

// GetMinIdleOk returns a tuple with the MinIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetMinIdleOk() (*float32, bool) {
	if o == nil || o.MinIdle == nil {
		return nil, false
	}
	return o.MinIdle, true
}

// HasMinIdle returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasMinIdle() bool {
	if o != nil && o.MinIdle != nil {
		return true
	}

	return false
}

// SetMinIdle gets a reference to the given float32 and assigns it to the MinIdle field.
func (o *ApiVdiPoolsIdVdiPool) SetMinIdle(v float32) {
	o.MinIdle = &v
}

// GetInitialPoolSize returns the InitialPoolSize field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetInitialPoolSize() float32 {
	if o == nil || o.InitialPoolSize == nil {
		var ret float32
		return ret
	}
	return *o.InitialPoolSize
}

// GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetInitialPoolSizeOk() (*float32, bool) {
	if o == nil || o.InitialPoolSize == nil {
		return nil, false
	}
	return o.InitialPoolSize, true
}

// HasInitialPoolSize returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasInitialPoolSize() bool {
	if o != nil && o.InitialPoolSize != nil {
		return true
	}

	return false
}

// SetInitialPoolSize gets a reference to the given float32 and assigns it to the InitialPoolSize field.
func (o *ApiVdiPoolsIdVdiPool) SetInitialPoolSize(v float32) {
	o.InitialPoolSize = &v
}

// GetMaxIdle returns the MaxIdle field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetMaxIdle() float32 {
	if o == nil || o.MaxIdle == nil {
		var ret float32
		return ret
	}
	return *o.MaxIdle
}

// GetMaxIdleOk returns a tuple with the MaxIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetMaxIdleOk() (*float32, bool) {
	if o == nil || o.MaxIdle == nil {
		return nil, false
	}
	return o.MaxIdle, true
}

// HasMaxIdle returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasMaxIdle() bool {
	if o != nil && o.MaxIdle != nil {
		return true
	}

	return false
}

// SetMaxIdle gets a reference to the given float32 and assigns it to the MaxIdle field.
func (o *ApiVdiPoolsIdVdiPool) SetMaxIdle(v float32) {
	o.MaxIdle = &v
}

// GetMaxPoolSize returns the MaxPoolSize field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetMaxPoolSize() float32 {
	if o == nil || o.MaxPoolSize == nil {
		var ret float32
		return ret
	}
	return *o.MaxPoolSize
}

// GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetMaxPoolSizeOk() (*float32, bool) {
	if o == nil || o.MaxPoolSize == nil {
		return nil, false
	}
	return o.MaxPoolSize, true
}

// HasMaxPoolSize returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasMaxPoolSize() bool {
	if o != nil && o.MaxPoolSize != nil {
		return true
	}

	return false
}

// SetMaxPoolSize gets a reference to the given float32 and assigns it to the MaxPoolSize field.
func (o *ApiVdiPoolsIdVdiPool) SetMaxPoolSize(v float32) {
	o.MaxPoolSize = &v
}

// GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAllocationTimeoutMinutes() float32 {
	if o == nil || o.AllocationTimeoutMinutes == nil {
		var ret float32
		return ret
	}
	return *o.AllocationTimeoutMinutes
}

// GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAllocationTimeoutMinutesOk() (*float32, bool) {
	if o == nil || o.AllocationTimeoutMinutes == nil {
		return nil, false
	}
	return o.AllocationTimeoutMinutes, true
}

// HasAllocationTimeoutMinutes returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAllocationTimeoutMinutes() bool {
	if o != nil && o.AllocationTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetAllocationTimeoutMinutes gets a reference to the given float32 and assigns it to the AllocationTimeoutMinutes field.
func (o *ApiVdiPoolsIdVdiPool) SetAllocationTimeoutMinutes(v float32) {
	o.AllocationTimeoutMinutes = &v
}

// GetPersistentUser returns the PersistentUser field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetPersistentUser() bool {
	if o == nil || o.PersistentUser == nil {
		var ret bool
		return ret
	}
	return *o.PersistentUser
}

// GetPersistentUserOk returns a tuple with the PersistentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetPersistentUserOk() (*bool, bool) {
	if o == nil || o.PersistentUser == nil {
		return nil, false
	}
	return o.PersistentUser, true
}

// HasPersistentUser returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasPersistentUser() bool {
	if o != nil && o.PersistentUser != nil {
		return true
	}

	return false
}

// SetPersistentUser gets a reference to the given bool and assigns it to the PersistentUser field.
func (o *ApiVdiPoolsIdVdiPool) SetPersistentUser(v bool) {
	o.PersistentUser = &v
}

// GetRecyclable returns the Recyclable field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetRecyclable() bool {
	if o == nil || o.Recyclable == nil {
		var ret bool
		return ret
	}
	return *o.Recyclable
}

// GetRecyclableOk returns a tuple with the Recyclable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetRecyclableOk() (*bool, bool) {
	if o == nil || o.Recyclable == nil {
		return nil, false
	}
	return o.Recyclable, true
}

// HasRecyclable returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasRecyclable() bool {
	if o != nil && o.Recyclable != nil {
		return true
	}

	return false
}

// SetRecyclable gets a reference to the given bool and assigns it to the Recyclable field.
func (o *ApiVdiPoolsIdVdiPool) SetRecyclable(v bool) {
	o.Recyclable = &v
}

// GetAllowCopy returns the AllowCopy field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAllowCopy() bool {
	if o == nil || o.AllowCopy == nil {
		var ret bool
		return ret
	}
	return *o.AllowCopy
}

// GetAllowCopyOk returns a tuple with the AllowCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAllowCopyOk() (*bool, bool) {
	if o == nil || o.AllowCopy == nil {
		return nil, false
	}
	return o.AllowCopy, true
}

// HasAllowCopy returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAllowCopy() bool {
	if o != nil && o.AllowCopy != nil {
		return true
	}

	return false
}

// SetAllowCopy gets a reference to the given bool and assigns it to the AllowCopy field.
func (o *ApiVdiPoolsIdVdiPool) SetAllowCopy(v bool) {
	o.AllowCopy = &v
}

// GetAllowPrinter returns the AllowPrinter field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAllowPrinter() bool {
	if o == nil || o.AllowPrinter == nil {
		var ret bool
		return ret
	}
	return *o.AllowPrinter
}

// GetAllowPrinterOk returns a tuple with the AllowPrinter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAllowPrinterOk() (*bool, bool) {
	if o == nil || o.AllowPrinter == nil {
		return nil, false
	}
	return o.AllowPrinter, true
}

// HasAllowPrinter returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAllowPrinter() bool {
	if o != nil && o.AllowPrinter != nil {
		return true
	}

	return false
}

// SetAllowPrinter gets a reference to the given bool and assigns it to the AllowPrinter field.
func (o *ApiVdiPoolsIdVdiPool) SetAllowPrinter(v bool) {
	o.AllowPrinter = &v
}

// GetAllowFileshare returns the AllowFileshare field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAllowFileshare() bool {
	if o == nil || o.AllowFileshare == nil {
		var ret bool
		return ret
	}
	return *o.AllowFileshare
}

// GetAllowFileshareOk returns a tuple with the AllowFileshare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAllowFileshareOk() (*bool, bool) {
	if o == nil || o.AllowFileshare == nil {
		return nil, false
	}
	return o.AllowFileshare, true
}

// HasAllowFileshare returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAllowFileshare() bool {
	if o != nil && o.AllowFileshare != nil {
		return true
	}

	return false
}

// SetAllowFileshare gets a reference to the given bool and assigns it to the AllowFileshare field.
func (o *ApiVdiPoolsIdVdiPool) SetAllowFileshare(v bool) {
	o.AllowFileshare = &v
}

// GetAllowHypervisorConsole returns the AllowHypervisorConsole field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAllowHypervisorConsole() bool {
	if o == nil || o.AllowHypervisorConsole == nil {
		var ret bool
		return ret
	}
	return *o.AllowHypervisorConsole
}

// GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAllowHypervisorConsoleOk() (*bool, bool) {
	if o == nil || o.AllowHypervisorConsole == nil {
		return nil, false
	}
	return o.AllowHypervisorConsole, true
}

// HasAllowHypervisorConsole returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAllowHypervisorConsole() bool {
	if o != nil && o.AllowHypervisorConsole != nil {
		return true
	}

	return false
}

// SetAllowHypervisorConsole gets a reference to the given bool and assigns it to the AllowHypervisorConsole field.
func (o *ApiVdiPoolsIdVdiPool) SetAllowHypervisorConsole(v bool) {
	o.AllowHypervisorConsole = &v
}

// GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetAutoCreateLocalUserOnReservation() bool {
	if o == nil || o.AutoCreateLocalUserOnReservation == nil {
		var ret bool
		return ret
	}
	return *o.AutoCreateLocalUserOnReservation
}

// GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool) {
	if o == nil || o.AutoCreateLocalUserOnReservation == nil {
		return nil, false
	}
	return o.AutoCreateLocalUserOnReservation, true
}

// HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasAutoCreateLocalUserOnReservation() bool {
	if o != nil && o.AutoCreateLocalUserOnReservation != nil {
		return true
	}

	return false
}

// SetAutoCreateLocalUserOnReservation gets a reference to the given bool and assigns it to the AutoCreateLocalUserOnReservation field.
func (o *ApiVdiPoolsIdVdiPool) SetAutoCreateLocalUserOnReservation(v bool) {
	o.AutoCreateLocalUserOnReservation = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiVdiPoolsIdVdiPool) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIconPath returns the IconPath field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetIconPath() string {
	if o == nil || o.IconPath == nil {
		var ret string
		return ret
	}
	return *o.IconPath
}

// GetIconPathOk returns a tuple with the IconPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetIconPathOk() (*string, bool) {
	if o == nil || o.IconPath == nil {
		return nil, false
	}
	return o.IconPath, true
}

// HasIconPath returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasIconPath() bool {
	if o != nil && o.IconPath != nil {
		return true
	}

	return false
}

// SetIconPath gets a reference to the given string and assigns it to the IconPath field.
func (o *ApiVdiPoolsIdVdiPool) SetIconPath(v string) {
	o.IconPath = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetApps() []int64 {
	if o == nil || o.Apps == nil {
		var ret []int64
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetAppsOk() (*[]int64, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []int64 and assigns it to the Apps field.
func (o *ApiVdiPoolsIdVdiPool) SetApps(v []int64) {
	o.Apps = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGateway() int64 {
	if o == nil || o.Gateway == nil {
		var ret int64
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGatewayOk() (*int64, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given int64 and assigns it to the Gateway field.
func (o *ApiVdiPoolsIdVdiPool) SetGateway(v int64) {
	o.Gateway = &v
}

// GetInstanceConfig returns the InstanceConfig field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetInstanceConfig() string {
	if o == nil || o.InstanceConfig == nil {
		var ret string
		return ret
	}
	return *o.InstanceConfig
}

// GetInstanceConfigOk returns a tuple with the InstanceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetInstanceConfigOk() (*string, bool) {
	if o == nil || o.InstanceConfig == nil {
		return nil, false
	}
	return o.InstanceConfig, true
}

// HasInstanceConfig returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasInstanceConfig() bool {
	if o != nil && o.InstanceConfig != nil {
		return true
	}

	return false
}

// SetInstanceConfig gets a reference to the given string and assigns it to the InstanceConfig field.
func (o *ApiVdiPoolsIdVdiPool) SetInstanceConfig(v string) {
	o.InstanceConfig = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetConfig() ApiVdiPoolsIdVdiPoolConfig {
	if o == nil || o.Config == nil {
		var ret ApiVdiPoolsIdVdiPoolConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetConfigOk() (*ApiVdiPoolsIdVdiPoolConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ApiVdiPoolsIdVdiPoolConfig and assigns it to the Config field.
func (o *ApiVdiPoolsIdVdiPool) SetConfig(v ApiVdiPoolsIdVdiPoolConfig) {
	o.Config = &v
}

// GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpHost() string {
	if o == nil || o.GuestConsoleJumpHost == nil {
		var ret string
		return ret
	}
	return *o.GuestConsoleJumpHost
}

// GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpHostOk() (*string, bool) {
	if o == nil || o.GuestConsoleJumpHost == nil {
		return nil, false
	}
	return o.GuestConsoleJumpHost, true
}

// HasGuestConsoleJumpHost returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpHost() bool {
	if o != nil && o.GuestConsoleJumpHost != nil {
		return true
	}

	return false
}

// SetGuestConsoleJumpHost gets a reference to the given string and assigns it to the GuestConsoleJumpHost field.
func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpHost(v string) {
	o.GuestConsoleJumpHost = &v
}

// GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPort() int64 {
	if o == nil || o.GuestConsoleJumpPort == nil {
		var ret int64
		return ret
	}
	return *o.GuestConsoleJumpPort
}

// GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPortOk() (*int64, bool) {
	if o == nil || o.GuestConsoleJumpPort == nil {
		return nil, false
	}
	return o.GuestConsoleJumpPort, true
}

// HasGuestConsoleJumpPort returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpPort() bool {
	if o != nil && o.GuestConsoleJumpPort != nil {
		return true
	}

	return false
}

// SetGuestConsoleJumpPort gets a reference to the given int64 and assigns it to the GuestConsoleJumpPort field.
func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpPort(v int64) {
	o.GuestConsoleJumpPort = &v
}

// GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpUsername() string {
	if o == nil || o.GuestConsoleJumpUsername == nil {
		var ret string
		return ret
	}
	return *o.GuestConsoleJumpUsername
}

// GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool) {
	if o == nil || o.GuestConsoleJumpUsername == nil {
		return nil, false
	}
	return o.GuestConsoleJumpUsername, true
}

// HasGuestConsoleJumpUsername returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpUsername() bool {
	if o != nil && o.GuestConsoleJumpUsername != nil {
		return true
	}

	return false
}

// SetGuestConsoleJumpUsername gets a reference to the given string and assigns it to the GuestConsoleJumpUsername field.
func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpUsername(v string) {
	o.GuestConsoleJumpUsername = &v
}

// GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPassword() string {
	if o == nil || o.GuestConsoleJumpPassword == nil {
		var ret string
		return ret
	}
	return *o.GuestConsoleJumpPassword
}

// GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool) {
	if o == nil || o.GuestConsoleJumpPassword == nil {
		return nil, false
	}
	return o.GuestConsoleJumpPassword, true
}

// HasGuestConsoleJumpPassword returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpPassword() bool {
	if o != nil && o.GuestConsoleJumpPassword != nil {
		return true
	}

	return false
}

// SetGuestConsoleJumpPassword gets a reference to the given string and assigns it to the GuestConsoleJumpPassword field.
func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpPassword(v string) {
	o.GuestConsoleJumpPassword = &v
}

// GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field value if set, zero value otherwise.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpKeypair() int64 {
	if o == nil || o.GuestConsoleJumpKeypair == nil {
		var ret int64
		return ret
	}
	return *o.GuestConsoleJumpKeypair
}

// GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVdiPoolsIdVdiPool) GetGuestConsoleJumpKeypairOk() (*int64, bool) {
	if o == nil || o.GuestConsoleJumpKeypair == nil {
		return nil, false
	}
	return o.GuestConsoleJumpKeypair, true
}

// HasGuestConsoleJumpKeypair returns a boolean if a field has been set.
func (o *ApiVdiPoolsIdVdiPool) HasGuestConsoleJumpKeypair() bool {
	if o != nil && o.GuestConsoleJumpKeypair != nil {
		return true
	}

	return false
}

// SetGuestConsoleJumpKeypair gets a reference to the given int64 and assigns it to the GuestConsoleJumpKeypair field.
func (o *ApiVdiPoolsIdVdiPool) SetGuestConsoleJumpKeypair(v int64) {
	o.GuestConsoleJumpKeypair = &v
}

func (o ApiVdiPoolsIdVdiPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.MinIdle != nil {
		toSerialize["minIdle"] = o.MinIdle
	}
	if o.InitialPoolSize != nil {
		toSerialize["initialPoolSize"] = o.InitialPoolSize
	}
	if o.MaxIdle != nil {
		toSerialize["maxIdle"] = o.MaxIdle
	}
	if o.MaxPoolSize != nil {
		toSerialize["maxPoolSize"] = o.MaxPoolSize
	}
	if o.AllocationTimeoutMinutes != nil {
		toSerialize["allocationTimeoutMinutes"] = o.AllocationTimeoutMinutes
	}
	if o.PersistentUser != nil {
		toSerialize["persistentUser"] = o.PersistentUser
	}
	if o.Recyclable != nil {
		toSerialize["recyclable"] = o.Recyclable
	}
	if o.AllowCopy != nil {
		toSerialize["allowCopy"] = o.AllowCopy
	}
	if o.AllowPrinter != nil {
		toSerialize["allowPrinter"] = o.AllowPrinter
	}
	if o.AllowFileshare != nil {
		toSerialize["allowFileshare"] = o.AllowFileshare
	}
	if o.AllowHypervisorConsole != nil {
		toSerialize["allowHypervisorConsole"] = o.AllowHypervisorConsole
	}
	if o.AutoCreateLocalUserOnReservation != nil {
		toSerialize["autoCreateLocalUserOnReservation"] = o.AutoCreateLocalUserOnReservation
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.IconPath != nil {
		toSerialize["iconPath"] = o.IconPath
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	if o.InstanceConfig != nil {
		toSerialize["instanceConfig"] = o.InstanceConfig
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.GuestConsoleJumpHost != nil {
		toSerialize["guestConsoleJumpHost"] = o.GuestConsoleJumpHost
	}
	if o.GuestConsoleJumpPort != nil {
		toSerialize["guestConsoleJumpPort"] = o.GuestConsoleJumpPort
	}
	if o.GuestConsoleJumpUsername != nil {
		toSerialize["guestConsoleJumpUsername"] = o.GuestConsoleJumpUsername
	}
	if o.GuestConsoleJumpPassword != nil {
		toSerialize["guestConsoleJumpPassword"] = o.GuestConsoleJumpPassword
	}
	if o.GuestConsoleJumpKeypair != nil {
		toSerialize["guestConsoleJumpKeypair"] = o.GuestConsoleJumpKeypair
	}
	return json.Marshal(toSerialize)
}

type NullableApiVdiPoolsIdVdiPool struct {
	value *ApiVdiPoolsIdVdiPool
	isSet bool
}

func (v NullableApiVdiPoolsIdVdiPool) Get() *ApiVdiPoolsIdVdiPool {
	return v.value
}

func (v *NullableApiVdiPoolsIdVdiPool) Set(val *ApiVdiPoolsIdVdiPool) {
	v.value = val
	v.isSet = true
}

func (v NullableApiVdiPoolsIdVdiPool) IsSet() bool {
	return v.isSet
}

func (v *NullableApiVdiPoolsIdVdiPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiVdiPoolsIdVdiPool(val *ApiVdiPoolsIdVdiPool) *NullableApiVdiPoolsIdVdiPool {
	return &NullableApiVdiPoolsIdVdiPool{value: val, isSet: true}
}

func (v NullableApiVdiPoolsIdVdiPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiVdiPoolsIdVdiPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

