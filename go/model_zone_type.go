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

// ZoneType struct for ZoneType
type ZoneType struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Provision *bool `json:"provision,omitempty"`
	AutoCapacity *bool `json:"autoCapacity,omitempty"`
	MigrationTarget *bool `json:"migrationTarget,omitempty"`
	HasDatastores *bool `json:"hasDatastores,omitempty"`
	HasNetworks *bool `json:"hasNetworks,omitempty"`
	HasResourcePools *bool `json:"hasResourcePools,omitempty"`
	HasSecurityGroups *bool `json:"hasSecurityGroups,omitempty"`
	HasContainers *bool `json:"hasContainers,omitempty"`
	HasBareMetal *bool `json:"hasBareMetal,omitempty"`
	HasServices *bool `json:"hasServices,omitempty"`
	HasFunctions *bool `json:"hasFunctions,omitempty"`
	HasJobs *bool `json:"hasJobs,omitempty"`
	HasDiscovery *bool `json:"hasDiscovery,omitempty"`
	HasCloudInit *bool `json:"hasCloudInit,omitempty"`
	HasFolders *bool `json:"hasFolders,omitempty"`
	HasFloatingIps *bool `json:"hasFloatingIps,omitempty"`
	HasMarketplace *bool `json:"hasMarketplace,omitempty"`
	CanCreateResourcePools *bool `json:"canCreateResourcePools,omitempty"`
	CanDeleteResourcePools *bool `json:"canDeleteResourcePools,omitempty"`
	CanCreateDatastores *bool `json:"canCreateDatastores,omitempty"`
	CanCreateNetworks *bool `json:"canCreateNetworks,omitempty"`
	CanChooseContainerMode *bool `json:"canChooseContainerMode,omitempty"`
	ProvisionRequiresResourcePool *bool `json:"provisionRequiresResourcePool,omitempty"`
	SupportsDistributedWorker *bool `json:"supportsDistributedWorker,omitempty"`
	Cloud *string `json:"cloud,omitempty"`
	ProvisionTypes *[]int64 `json:"provisionTypes,omitempty"`
	ZoneInstanceTypeLayoutId *int64 `json:"zoneInstanceTypeLayoutId,omitempty"`
	ServerTypes *[]ZoneTypeServerTypes `json:"serverTypes,omitempty"`
	OptionTypes *[]ZoneTypeOptionTypes1 `json:"optionTypes,omitempty"`
}

// NewZoneType instantiates a new ZoneType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneType() *ZoneType {
	this := ZoneType{}
	return &this
}

// NewZoneTypeWithDefaults instantiates a new ZoneType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTypeWithDefaults() *ZoneType {
	this := ZoneType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneType) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneType) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ZoneType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ZoneType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ZoneType) SetCode(v string) {
	o.Code = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ZoneType) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ZoneType) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ZoneType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProvision returns the Provision field value if set, zero value otherwise.
func (o *ZoneType) GetProvision() bool {
	if o == nil || o.Provision == nil {
		var ret bool
		return ret
	}
	return *o.Provision
}

// GetProvisionOk returns a tuple with the Provision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetProvisionOk() (*bool, bool) {
	if o == nil || o.Provision == nil {
		return nil, false
	}
	return o.Provision, true
}

// HasProvision returns a boolean if a field has been set.
func (o *ZoneType) HasProvision() bool {
	if o != nil && o.Provision != nil {
		return true
	}

	return false
}

// SetProvision gets a reference to the given bool and assigns it to the Provision field.
func (o *ZoneType) SetProvision(v bool) {
	o.Provision = &v
}

// GetAutoCapacity returns the AutoCapacity field value if set, zero value otherwise.
func (o *ZoneType) GetAutoCapacity() bool {
	if o == nil || o.AutoCapacity == nil {
		var ret bool
		return ret
	}
	return *o.AutoCapacity
}

// GetAutoCapacityOk returns a tuple with the AutoCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetAutoCapacityOk() (*bool, bool) {
	if o == nil || o.AutoCapacity == nil {
		return nil, false
	}
	return o.AutoCapacity, true
}

// HasAutoCapacity returns a boolean if a field has been set.
func (o *ZoneType) HasAutoCapacity() bool {
	if o != nil && o.AutoCapacity != nil {
		return true
	}

	return false
}

// SetAutoCapacity gets a reference to the given bool and assigns it to the AutoCapacity field.
func (o *ZoneType) SetAutoCapacity(v bool) {
	o.AutoCapacity = &v
}

// GetMigrationTarget returns the MigrationTarget field value if set, zero value otherwise.
func (o *ZoneType) GetMigrationTarget() bool {
	if o == nil || o.MigrationTarget == nil {
		var ret bool
		return ret
	}
	return *o.MigrationTarget
}

// GetMigrationTargetOk returns a tuple with the MigrationTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetMigrationTargetOk() (*bool, bool) {
	if o == nil || o.MigrationTarget == nil {
		return nil, false
	}
	return o.MigrationTarget, true
}

// HasMigrationTarget returns a boolean if a field has been set.
func (o *ZoneType) HasMigrationTarget() bool {
	if o != nil && o.MigrationTarget != nil {
		return true
	}

	return false
}

// SetMigrationTarget gets a reference to the given bool and assigns it to the MigrationTarget field.
func (o *ZoneType) SetMigrationTarget(v bool) {
	o.MigrationTarget = &v
}

// GetHasDatastores returns the HasDatastores field value if set, zero value otherwise.
func (o *ZoneType) GetHasDatastores() bool {
	if o == nil || o.HasDatastores == nil {
		var ret bool
		return ret
	}
	return *o.HasDatastores
}

// GetHasDatastoresOk returns a tuple with the HasDatastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasDatastoresOk() (*bool, bool) {
	if o == nil || o.HasDatastores == nil {
		return nil, false
	}
	return o.HasDatastores, true
}

// HasHasDatastores returns a boolean if a field has been set.
func (o *ZoneType) HasHasDatastores() bool {
	if o != nil && o.HasDatastores != nil {
		return true
	}

	return false
}

// SetHasDatastores gets a reference to the given bool and assigns it to the HasDatastores field.
func (o *ZoneType) SetHasDatastores(v bool) {
	o.HasDatastores = &v
}

// GetHasNetworks returns the HasNetworks field value if set, zero value otherwise.
func (o *ZoneType) GetHasNetworks() bool {
	if o == nil || o.HasNetworks == nil {
		var ret bool
		return ret
	}
	return *o.HasNetworks
}

// GetHasNetworksOk returns a tuple with the HasNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasNetworksOk() (*bool, bool) {
	if o == nil || o.HasNetworks == nil {
		return nil, false
	}
	return o.HasNetworks, true
}

// HasHasNetworks returns a boolean if a field has been set.
func (o *ZoneType) HasHasNetworks() bool {
	if o != nil && o.HasNetworks != nil {
		return true
	}

	return false
}

// SetHasNetworks gets a reference to the given bool and assigns it to the HasNetworks field.
func (o *ZoneType) SetHasNetworks(v bool) {
	o.HasNetworks = &v
}

// GetHasResourcePools returns the HasResourcePools field value if set, zero value otherwise.
func (o *ZoneType) GetHasResourcePools() bool {
	if o == nil || o.HasResourcePools == nil {
		var ret bool
		return ret
	}
	return *o.HasResourcePools
}

// GetHasResourcePoolsOk returns a tuple with the HasResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasResourcePoolsOk() (*bool, bool) {
	if o == nil || o.HasResourcePools == nil {
		return nil, false
	}
	return o.HasResourcePools, true
}

// HasHasResourcePools returns a boolean if a field has been set.
func (o *ZoneType) HasHasResourcePools() bool {
	if o != nil && o.HasResourcePools != nil {
		return true
	}

	return false
}

// SetHasResourcePools gets a reference to the given bool and assigns it to the HasResourcePools field.
func (o *ZoneType) SetHasResourcePools(v bool) {
	o.HasResourcePools = &v
}

// GetHasSecurityGroups returns the HasSecurityGroups field value if set, zero value otherwise.
func (o *ZoneType) GetHasSecurityGroups() bool {
	if o == nil || o.HasSecurityGroups == nil {
		var ret bool
		return ret
	}
	return *o.HasSecurityGroups
}

// GetHasSecurityGroupsOk returns a tuple with the HasSecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasSecurityGroupsOk() (*bool, bool) {
	if o == nil || o.HasSecurityGroups == nil {
		return nil, false
	}
	return o.HasSecurityGroups, true
}

// HasHasSecurityGroups returns a boolean if a field has been set.
func (o *ZoneType) HasHasSecurityGroups() bool {
	if o != nil && o.HasSecurityGroups != nil {
		return true
	}

	return false
}

// SetHasSecurityGroups gets a reference to the given bool and assigns it to the HasSecurityGroups field.
func (o *ZoneType) SetHasSecurityGroups(v bool) {
	o.HasSecurityGroups = &v
}

// GetHasContainers returns the HasContainers field value if set, zero value otherwise.
func (o *ZoneType) GetHasContainers() bool {
	if o == nil || o.HasContainers == nil {
		var ret bool
		return ret
	}
	return *o.HasContainers
}

// GetHasContainersOk returns a tuple with the HasContainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasContainersOk() (*bool, bool) {
	if o == nil || o.HasContainers == nil {
		return nil, false
	}
	return o.HasContainers, true
}

// HasHasContainers returns a boolean if a field has been set.
func (o *ZoneType) HasHasContainers() bool {
	if o != nil && o.HasContainers != nil {
		return true
	}

	return false
}

// SetHasContainers gets a reference to the given bool and assigns it to the HasContainers field.
func (o *ZoneType) SetHasContainers(v bool) {
	o.HasContainers = &v
}

// GetHasBareMetal returns the HasBareMetal field value if set, zero value otherwise.
func (o *ZoneType) GetHasBareMetal() bool {
	if o == nil || o.HasBareMetal == nil {
		var ret bool
		return ret
	}
	return *o.HasBareMetal
}

// GetHasBareMetalOk returns a tuple with the HasBareMetal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasBareMetalOk() (*bool, bool) {
	if o == nil || o.HasBareMetal == nil {
		return nil, false
	}
	return o.HasBareMetal, true
}

// HasHasBareMetal returns a boolean if a field has been set.
func (o *ZoneType) HasHasBareMetal() bool {
	if o != nil && o.HasBareMetal != nil {
		return true
	}

	return false
}

// SetHasBareMetal gets a reference to the given bool and assigns it to the HasBareMetal field.
func (o *ZoneType) SetHasBareMetal(v bool) {
	o.HasBareMetal = &v
}

// GetHasServices returns the HasServices field value if set, zero value otherwise.
func (o *ZoneType) GetHasServices() bool {
	if o == nil || o.HasServices == nil {
		var ret bool
		return ret
	}
	return *o.HasServices
}

// GetHasServicesOk returns a tuple with the HasServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasServicesOk() (*bool, bool) {
	if o == nil || o.HasServices == nil {
		return nil, false
	}
	return o.HasServices, true
}

// HasHasServices returns a boolean if a field has been set.
func (o *ZoneType) HasHasServices() bool {
	if o != nil && o.HasServices != nil {
		return true
	}

	return false
}

// SetHasServices gets a reference to the given bool and assigns it to the HasServices field.
func (o *ZoneType) SetHasServices(v bool) {
	o.HasServices = &v
}

// GetHasFunctions returns the HasFunctions field value if set, zero value otherwise.
func (o *ZoneType) GetHasFunctions() bool {
	if o == nil || o.HasFunctions == nil {
		var ret bool
		return ret
	}
	return *o.HasFunctions
}

// GetHasFunctionsOk returns a tuple with the HasFunctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasFunctionsOk() (*bool, bool) {
	if o == nil || o.HasFunctions == nil {
		return nil, false
	}
	return o.HasFunctions, true
}

// HasHasFunctions returns a boolean if a field has been set.
func (o *ZoneType) HasHasFunctions() bool {
	if o != nil && o.HasFunctions != nil {
		return true
	}

	return false
}

// SetHasFunctions gets a reference to the given bool and assigns it to the HasFunctions field.
func (o *ZoneType) SetHasFunctions(v bool) {
	o.HasFunctions = &v
}

// GetHasJobs returns the HasJobs field value if set, zero value otherwise.
func (o *ZoneType) GetHasJobs() bool {
	if o == nil || o.HasJobs == nil {
		var ret bool
		return ret
	}
	return *o.HasJobs
}

// GetHasJobsOk returns a tuple with the HasJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasJobsOk() (*bool, bool) {
	if o == nil || o.HasJobs == nil {
		return nil, false
	}
	return o.HasJobs, true
}

// HasHasJobs returns a boolean if a field has been set.
func (o *ZoneType) HasHasJobs() bool {
	if o != nil && o.HasJobs != nil {
		return true
	}

	return false
}

// SetHasJobs gets a reference to the given bool and assigns it to the HasJobs field.
func (o *ZoneType) SetHasJobs(v bool) {
	o.HasJobs = &v
}

// GetHasDiscovery returns the HasDiscovery field value if set, zero value otherwise.
func (o *ZoneType) GetHasDiscovery() bool {
	if o == nil || o.HasDiscovery == nil {
		var ret bool
		return ret
	}
	return *o.HasDiscovery
}

// GetHasDiscoveryOk returns a tuple with the HasDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasDiscoveryOk() (*bool, bool) {
	if o == nil || o.HasDiscovery == nil {
		return nil, false
	}
	return o.HasDiscovery, true
}

// HasHasDiscovery returns a boolean if a field has been set.
func (o *ZoneType) HasHasDiscovery() bool {
	if o != nil && o.HasDiscovery != nil {
		return true
	}

	return false
}

// SetHasDiscovery gets a reference to the given bool and assigns it to the HasDiscovery field.
func (o *ZoneType) SetHasDiscovery(v bool) {
	o.HasDiscovery = &v
}

// GetHasCloudInit returns the HasCloudInit field value if set, zero value otherwise.
func (o *ZoneType) GetHasCloudInit() bool {
	if o == nil || o.HasCloudInit == nil {
		var ret bool
		return ret
	}
	return *o.HasCloudInit
}

// GetHasCloudInitOk returns a tuple with the HasCloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasCloudInitOk() (*bool, bool) {
	if o == nil || o.HasCloudInit == nil {
		return nil, false
	}
	return o.HasCloudInit, true
}

// HasHasCloudInit returns a boolean if a field has been set.
func (o *ZoneType) HasHasCloudInit() bool {
	if o != nil && o.HasCloudInit != nil {
		return true
	}

	return false
}

// SetHasCloudInit gets a reference to the given bool and assigns it to the HasCloudInit field.
func (o *ZoneType) SetHasCloudInit(v bool) {
	o.HasCloudInit = &v
}

// GetHasFolders returns the HasFolders field value if set, zero value otherwise.
func (o *ZoneType) GetHasFolders() bool {
	if o == nil || o.HasFolders == nil {
		var ret bool
		return ret
	}
	return *o.HasFolders
}

// GetHasFoldersOk returns a tuple with the HasFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasFoldersOk() (*bool, bool) {
	if o == nil || o.HasFolders == nil {
		return nil, false
	}
	return o.HasFolders, true
}

// HasHasFolders returns a boolean if a field has been set.
func (o *ZoneType) HasHasFolders() bool {
	if o != nil && o.HasFolders != nil {
		return true
	}

	return false
}

// SetHasFolders gets a reference to the given bool and assigns it to the HasFolders field.
func (o *ZoneType) SetHasFolders(v bool) {
	o.HasFolders = &v
}

// GetHasFloatingIps returns the HasFloatingIps field value if set, zero value otherwise.
func (o *ZoneType) GetHasFloatingIps() bool {
	if o == nil || o.HasFloatingIps == nil {
		var ret bool
		return ret
	}
	return *o.HasFloatingIps
}

// GetHasFloatingIpsOk returns a tuple with the HasFloatingIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasFloatingIpsOk() (*bool, bool) {
	if o == nil || o.HasFloatingIps == nil {
		return nil, false
	}
	return o.HasFloatingIps, true
}

// HasHasFloatingIps returns a boolean if a field has been set.
func (o *ZoneType) HasHasFloatingIps() bool {
	if o != nil && o.HasFloatingIps != nil {
		return true
	}

	return false
}

// SetHasFloatingIps gets a reference to the given bool and assigns it to the HasFloatingIps field.
func (o *ZoneType) SetHasFloatingIps(v bool) {
	o.HasFloatingIps = &v
}

// GetHasMarketplace returns the HasMarketplace field value if set, zero value otherwise.
func (o *ZoneType) GetHasMarketplace() bool {
	if o == nil || o.HasMarketplace == nil {
		var ret bool
		return ret
	}
	return *o.HasMarketplace
}

// GetHasMarketplaceOk returns a tuple with the HasMarketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetHasMarketplaceOk() (*bool, bool) {
	if o == nil || o.HasMarketplace == nil {
		return nil, false
	}
	return o.HasMarketplace, true
}

// HasHasMarketplace returns a boolean if a field has been set.
func (o *ZoneType) HasHasMarketplace() bool {
	if o != nil && o.HasMarketplace != nil {
		return true
	}

	return false
}

// SetHasMarketplace gets a reference to the given bool and assigns it to the HasMarketplace field.
func (o *ZoneType) SetHasMarketplace(v bool) {
	o.HasMarketplace = &v
}

// GetCanCreateResourcePools returns the CanCreateResourcePools field value if set, zero value otherwise.
func (o *ZoneType) GetCanCreateResourcePools() bool {
	if o == nil || o.CanCreateResourcePools == nil {
		var ret bool
		return ret
	}
	return *o.CanCreateResourcePools
}

// GetCanCreateResourcePoolsOk returns a tuple with the CanCreateResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCanCreateResourcePoolsOk() (*bool, bool) {
	if o == nil || o.CanCreateResourcePools == nil {
		return nil, false
	}
	return o.CanCreateResourcePools, true
}

// HasCanCreateResourcePools returns a boolean if a field has been set.
func (o *ZoneType) HasCanCreateResourcePools() bool {
	if o != nil && o.CanCreateResourcePools != nil {
		return true
	}

	return false
}

// SetCanCreateResourcePools gets a reference to the given bool and assigns it to the CanCreateResourcePools field.
func (o *ZoneType) SetCanCreateResourcePools(v bool) {
	o.CanCreateResourcePools = &v
}

// GetCanDeleteResourcePools returns the CanDeleteResourcePools field value if set, zero value otherwise.
func (o *ZoneType) GetCanDeleteResourcePools() bool {
	if o == nil || o.CanDeleteResourcePools == nil {
		var ret bool
		return ret
	}
	return *o.CanDeleteResourcePools
}

// GetCanDeleteResourcePoolsOk returns a tuple with the CanDeleteResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCanDeleteResourcePoolsOk() (*bool, bool) {
	if o == nil || o.CanDeleteResourcePools == nil {
		return nil, false
	}
	return o.CanDeleteResourcePools, true
}

// HasCanDeleteResourcePools returns a boolean if a field has been set.
func (o *ZoneType) HasCanDeleteResourcePools() bool {
	if o != nil && o.CanDeleteResourcePools != nil {
		return true
	}

	return false
}

// SetCanDeleteResourcePools gets a reference to the given bool and assigns it to the CanDeleteResourcePools field.
func (o *ZoneType) SetCanDeleteResourcePools(v bool) {
	o.CanDeleteResourcePools = &v
}

// GetCanCreateDatastores returns the CanCreateDatastores field value if set, zero value otherwise.
func (o *ZoneType) GetCanCreateDatastores() bool {
	if o == nil || o.CanCreateDatastores == nil {
		var ret bool
		return ret
	}
	return *o.CanCreateDatastores
}

// GetCanCreateDatastoresOk returns a tuple with the CanCreateDatastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCanCreateDatastoresOk() (*bool, bool) {
	if o == nil || o.CanCreateDatastores == nil {
		return nil, false
	}
	return o.CanCreateDatastores, true
}

// HasCanCreateDatastores returns a boolean if a field has been set.
func (o *ZoneType) HasCanCreateDatastores() bool {
	if o != nil && o.CanCreateDatastores != nil {
		return true
	}

	return false
}

// SetCanCreateDatastores gets a reference to the given bool and assigns it to the CanCreateDatastores field.
func (o *ZoneType) SetCanCreateDatastores(v bool) {
	o.CanCreateDatastores = &v
}

// GetCanCreateNetworks returns the CanCreateNetworks field value if set, zero value otherwise.
func (o *ZoneType) GetCanCreateNetworks() bool {
	if o == nil || o.CanCreateNetworks == nil {
		var ret bool
		return ret
	}
	return *o.CanCreateNetworks
}

// GetCanCreateNetworksOk returns a tuple with the CanCreateNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCanCreateNetworksOk() (*bool, bool) {
	if o == nil || o.CanCreateNetworks == nil {
		return nil, false
	}
	return o.CanCreateNetworks, true
}

// HasCanCreateNetworks returns a boolean if a field has been set.
func (o *ZoneType) HasCanCreateNetworks() bool {
	if o != nil && o.CanCreateNetworks != nil {
		return true
	}

	return false
}

// SetCanCreateNetworks gets a reference to the given bool and assigns it to the CanCreateNetworks field.
func (o *ZoneType) SetCanCreateNetworks(v bool) {
	o.CanCreateNetworks = &v
}

// GetCanChooseContainerMode returns the CanChooseContainerMode field value if set, zero value otherwise.
func (o *ZoneType) GetCanChooseContainerMode() bool {
	if o == nil || o.CanChooseContainerMode == nil {
		var ret bool
		return ret
	}
	return *o.CanChooseContainerMode
}

// GetCanChooseContainerModeOk returns a tuple with the CanChooseContainerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCanChooseContainerModeOk() (*bool, bool) {
	if o == nil || o.CanChooseContainerMode == nil {
		return nil, false
	}
	return o.CanChooseContainerMode, true
}

// HasCanChooseContainerMode returns a boolean if a field has been set.
func (o *ZoneType) HasCanChooseContainerMode() bool {
	if o != nil && o.CanChooseContainerMode != nil {
		return true
	}

	return false
}

// SetCanChooseContainerMode gets a reference to the given bool and assigns it to the CanChooseContainerMode field.
func (o *ZoneType) SetCanChooseContainerMode(v bool) {
	o.CanChooseContainerMode = &v
}

// GetProvisionRequiresResourcePool returns the ProvisionRequiresResourcePool field value if set, zero value otherwise.
func (o *ZoneType) GetProvisionRequiresResourcePool() bool {
	if o == nil || o.ProvisionRequiresResourcePool == nil {
		var ret bool
		return ret
	}
	return *o.ProvisionRequiresResourcePool
}

// GetProvisionRequiresResourcePoolOk returns a tuple with the ProvisionRequiresResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetProvisionRequiresResourcePoolOk() (*bool, bool) {
	if o == nil || o.ProvisionRequiresResourcePool == nil {
		return nil, false
	}
	return o.ProvisionRequiresResourcePool, true
}

// HasProvisionRequiresResourcePool returns a boolean if a field has been set.
func (o *ZoneType) HasProvisionRequiresResourcePool() bool {
	if o != nil && o.ProvisionRequiresResourcePool != nil {
		return true
	}

	return false
}

// SetProvisionRequiresResourcePool gets a reference to the given bool and assigns it to the ProvisionRequiresResourcePool field.
func (o *ZoneType) SetProvisionRequiresResourcePool(v bool) {
	o.ProvisionRequiresResourcePool = &v
}

// GetSupportsDistributedWorker returns the SupportsDistributedWorker field value if set, zero value otherwise.
func (o *ZoneType) GetSupportsDistributedWorker() bool {
	if o == nil || o.SupportsDistributedWorker == nil {
		var ret bool
		return ret
	}
	return *o.SupportsDistributedWorker
}

// GetSupportsDistributedWorkerOk returns a tuple with the SupportsDistributedWorker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetSupportsDistributedWorkerOk() (*bool, bool) {
	if o == nil || o.SupportsDistributedWorker == nil {
		return nil, false
	}
	return o.SupportsDistributedWorker, true
}

// HasSupportsDistributedWorker returns a boolean if a field has been set.
func (o *ZoneType) HasSupportsDistributedWorker() bool {
	if o != nil && o.SupportsDistributedWorker != nil {
		return true
	}

	return false
}

// SetSupportsDistributedWorker gets a reference to the given bool and assigns it to the SupportsDistributedWorker field.
func (o *ZoneType) SetSupportsDistributedWorker(v bool) {
	o.SupportsDistributedWorker = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *ZoneType) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *ZoneType) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *ZoneType) SetCloud(v string) {
	o.Cloud = &v
}

// GetProvisionTypes returns the ProvisionTypes field value if set, zero value otherwise.
func (o *ZoneType) GetProvisionTypes() []int64 {
	if o == nil || o.ProvisionTypes == nil {
		var ret []int64
		return ret
	}
	return *o.ProvisionTypes
}

// GetProvisionTypesOk returns a tuple with the ProvisionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetProvisionTypesOk() (*[]int64, bool) {
	if o == nil || o.ProvisionTypes == nil {
		return nil, false
	}
	return o.ProvisionTypes, true
}

// HasProvisionTypes returns a boolean if a field has been set.
func (o *ZoneType) HasProvisionTypes() bool {
	if o != nil && o.ProvisionTypes != nil {
		return true
	}

	return false
}

// SetProvisionTypes gets a reference to the given []int64 and assigns it to the ProvisionTypes field.
func (o *ZoneType) SetProvisionTypes(v []int64) {
	o.ProvisionTypes = &v
}

// GetZoneInstanceTypeLayoutId returns the ZoneInstanceTypeLayoutId field value if set, zero value otherwise.
func (o *ZoneType) GetZoneInstanceTypeLayoutId() int64 {
	if o == nil || o.ZoneInstanceTypeLayoutId == nil {
		var ret int64
		return ret
	}
	return *o.ZoneInstanceTypeLayoutId
}

// GetZoneInstanceTypeLayoutIdOk returns a tuple with the ZoneInstanceTypeLayoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetZoneInstanceTypeLayoutIdOk() (*int64, bool) {
	if o == nil || o.ZoneInstanceTypeLayoutId == nil {
		return nil, false
	}
	return o.ZoneInstanceTypeLayoutId, true
}

// HasZoneInstanceTypeLayoutId returns a boolean if a field has been set.
func (o *ZoneType) HasZoneInstanceTypeLayoutId() bool {
	if o != nil && o.ZoneInstanceTypeLayoutId != nil {
		return true
	}

	return false
}

// SetZoneInstanceTypeLayoutId gets a reference to the given int64 and assigns it to the ZoneInstanceTypeLayoutId field.
func (o *ZoneType) SetZoneInstanceTypeLayoutId(v int64) {
	o.ZoneInstanceTypeLayoutId = &v
}

// GetServerTypes returns the ServerTypes field value if set, zero value otherwise.
func (o *ZoneType) GetServerTypes() []ZoneTypeServerTypes {
	if o == nil || o.ServerTypes == nil {
		var ret []ZoneTypeServerTypes
		return ret
	}
	return *o.ServerTypes
}

// GetServerTypesOk returns a tuple with the ServerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetServerTypesOk() (*[]ZoneTypeServerTypes, bool) {
	if o == nil || o.ServerTypes == nil {
		return nil, false
	}
	return o.ServerTypes, true
}

// HasServerTypes returns a boolean if a field has been set.
func (o *ZoneType) HasServerTypes() bool {
	if o != nil && o.ServerTypes != nil {
		return true
	}

	return false
}

// SetServerTypes gets a reference to the given []ZoneTypeServerTypes and assigns it to the ServerTypes field.
func (o *ZoneType) SetServerTypes(v []ZoneTypeServerTypes) {
	o.ServerTypes = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ZoneType) GetOptionTypes() []ZoneTypeOptionTypes1 {
	if o == nil || o.OptionTypes == nil {
		var ret []ZoneTypeOptionTypes1
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneType) GetOptionTypesOk() (*[]ZoneTypeOptionTypes1, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *ZoneType) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []ZoneTypeOptionTypes1 and assigns it to the OptionTypes field.
func (o *ZoneType) SetOptionTypes(v []ZoneTypeOptionTypes1) {
	o.OptionTypes = &v
}

func (o ZoneType) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Provision != nil {
		toSerialize["provision"] = o.Provision
	}
	if o.AutoCapacity != nil {
		toSerialize["autoCapacity"] = o.AutoCapacity
	}
	if o.MigrationTarget != nil {
		toSerialize["migrationTarget"] = o.MigrationTarget
	}
	if o.HasDatastores != nil {
		toSerialize["hasDatastores"] = o.HasDatastores
	}
	if o.HasNetworks != nil {
		toSerialize["hasNetworks"] = o.HasNetworks
	}
	if o.HasResourcePools != nil {
		toSerialize["hasResourcePools"] = o.HasResourcePools
	}
	if o.HasSecurityGroups != nil {
		toSerialize["hasSecurityGroups"] = o.HasSecurityGroups
	}
	if o.HasContainers != nil {
		toSerialize["hasContainers"] = o.HasContainers
	}
	if o.HasBareMetal != nil {
		toSerialize["hasBareMetal"] = o.HasBareMetal
	}
	if o.HasServices != nil {
		toSerialize["hasServices"] = o.HasServices
	}
	if o.HasFunctions != nil {
		toSerialize["hasFunctions"] = o.HasFunctions
	}
	if o.HasJobs != nil {
		toSerialize["hasJobs"] = o.HasJobs
	}
	if o.HasDiscovery != nil {
		toSerialize["hasDiscovery"] = o.HasDiscovery
	}
	if o.HasCloudInit != nil {
		toSerialize["hasCloudInit"] = o.HasCloudInit
	}
	if o.HasFolders != nil {
		toSerialize["hasFolders"] = o.HasFolders
	}
	if o.HasFloatingIps != nil {
		toSerialize["hasFloatingIps"] = o.HasFloatingIps
	}
	if o.HasMarketplace != nil {
		toSerialize["hasMarketplace"] = o.HasMarketplace
	}
	if o.CanCreateResourcePools != nil {
		toSerialize["canCreateResourcePools"] = o.CanCreateResourcePools
	}
	if o.CanDeleteResourcePools != nil {
		toSerialize["canDeleteResourcePools"] = o.CanDeleteResourcePools
	}
	if o.CanCreateDatastores != nil {
		toSerialize["canCreateDatastores"] = o.CanCreateDatastores
	}
	if o.CanCreateNetworks != nil {
		toSerialize["canCreateNetworks"] = o.CanCreateNetworks
	}
	if o.CanChooseContainerMode != nil {
		toSerialize["canChooseContainerMode"] = o.CanChooseContainerMode
	}
	if o.ProvisionRequiresResourcePool != nil {
		toSerialize["provisionRequiresResourcePool"] = o.ProvisionRequiresResourcePool
	}
	if o.SupportsDistributedWorker != nil {
		toSerialize["supportsDistributedWorker"] = o.SupportsDistributedWorker
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.ProvisionTypes != nil {
		toSerialize["provisionTypes"] = o.ProvisionTypes
	}
	if o.ZoneInstanceTypeLayoutId != nil {
		toSerialize["zoneInstanceTypeLayoutId"] = o.ZoneInstanceTypeLayoutId
	}
	if o.ServerTypes != nil {
		toSerialize["serverTypes"] = o.ServerTypes
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	return json.Marshal(toSerialize)
}

type NullableZoneType struct {
	value *ZoneType
	isSet bool
}

func (v NullableZoneType) Get() *ZoneType {
	return v.value
}

func (v *NullableZoneType) Set(val *ZoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneType(val *ZoneType) *NullableZoneType {
	return &NullableZoneType{value: val, isSet: true}
}

func (v NullableZoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

