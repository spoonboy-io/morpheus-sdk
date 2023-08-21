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

// ClusterDatastore struct for ClusterDatastore
type ClusterDatastore struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	StorageSize NullableInt64 `json:"storageSize,omitempty"`
	FreeSpace NullableInt64 `json:"freeSpace,omitempty"`
	DrsEnabled *bool `json:"drsEnabled,omitempty"`
	Active *bool `json:"active,omitempty"`
	AllowWrite *bool `json:"allowWrite,omitempty"`
	DefaultStore *bool `json:"defaultStore,omitempty"`
	Online *bool `json:"online,omitempty"`
	AllowRead *bool `json:"allowRead,omitempty"`
	AllowProvision *bool `json:"allowProvision,omitempty"`
	RefType *string `json:"refType,omitempty"`
	RefId *int64 `json:"refId,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Zone *ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"zone,omitempty"`
	ZonePool *ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"zonePool,omitempty"`
	Owner *ApiBlueprintsIdUpdatePermissionsResourcePermissionSites `json:"owner,omitempty"`
	Tenants *[]InlineResponse20040AppDeployInstance `json:"tenants,omitempty"`
	Permissions *ClusterDatastorePermissions `json:"permissions,omitempty"`
	Datastores *[]map[string]interface{} `json:"datastores,omitempty"`
}

// NewClusterDatastore instantiates a new ClusterDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDatastore() *ClusterDatastore {
	this := ClusterDatastore{}
	return &this
}

// NewClusterDatastoreWithDefaults instantiates a new ClusterDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDatastoreWithDefaults() *ClusterDatastore {
	this := ClusterDatastore{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterDatastore) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterDatastore) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterDatastore) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterDatastore) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterDatastore) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterDatastore) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterDatastore) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterDatastore) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ClusterDatastore) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ClusterDatastore) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ClusterDatastore) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ClusterDatastore) UnsetCode() {
	o.Code.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClusterDatastore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterDatastore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterDatastore) SetType(v string) {
	o.Type = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ClusterDatastore) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ClusterDatastore) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ClusterDatastore) SetVisibility(v string) {
	o.Visibility = &v
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterDatastore) GetStorageSize() int64 {
	if o == nil || o.StorageSize.Get() == nil {
		var ret int64
		return ret
	}
	return *o.StorageSize.Get()
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterDatastore) GetStorageSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageSize.Get(), o.StorageSize.IsSet()
}

// HasStorageSize returns a boolean if a field has been set.
func (o *ClusterDatastore) HasStorageSize() bool {
	if o != nil && o.StorageSize.IsSet() {
		return true
	}

	return false
}

// SetStorageSize gets a reference to the given NullableInt64 and assigns it to the StorageSize field.
func (o *ClusterDatastore) SetStorageSize(v int64) {
	o.StorageSize.Set(&v)
}
// SetStorageSizeNil sets the value for StorageSize to be an explicit nil
func (o *ClusterDatastore) SetStorageSizeNil() {
	o.StorageSize.Set(nil)
}

// UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
func (o *ClusterDatastore) UnsetStorageSize() {
	o.StorageSize.Unset()
}

// GetFreeSpace returns the FreeSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterDatastore) GetFreeSpace() int64 {
	if o == nil || o.FreeSpace.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FreeSpace.Get()
}

// GetFreeSpaceOk returns a tuple with the FreeSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterDatastore) GetFreeSpaceOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FreeSpace.Get(), o.FreeSpace.IsSet()
}

// HasFreeSpace returns a boolean if a field has been set.
func (o *ClusterDatastore) HasFreeSpace() bool {
	if o != nil && o.FreeSpace.IsSet() {
		return true
	}

	return false
}

// SetFreeSpace gets a reference to the given NullableInt64 and assigns it to the FreeSpace field.
func (o *ClusterDatastore) SetFreeSpace(v int64) {
	o.FreeSpace.Set(&v)
}
// SetFreeSpaceNil sets the value for FreeSpace to be an explicit nil
func (o *ClusterDatastore) SetFreeSpaceNil() {
	o.FreeSpace.Set(nil)
}

// UnsetFreeSpace ensures that no value is present for FreeSpace, not even an explicit nil
func (o *ClusterDatastore) UnsetFreeSpace() {
	o.FreeSpace.Unset()
}

// GetDrsEnabled returns the DrsEnabled field value if set, zero value otherwise.
func (o *ClusterDatastore) GetDrsEnabled() bool {
	if o == nil || o.DrsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DrsEnabled
}

// GetDrsEnabledOk returns a tuple with the DrsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetDrsEnabledOk() (*bool, bool) {
	if o == nil || o.DrsEnabled == nil {
		return nil, false
	}
	return o.DrsEnabled, true
}

// HasDrsEnabled returns a boolean if a field has been set.
func (o *ClusterDatastore) HasDrsEnabled() bool {
	if o != nil && o.DrsEnabled != nil {
		return true
	}

	return false
}

// SetDrsEnabled gets a reference to the given bool and assigns it to the DrsEnabled field.
func (o *ClusterDatastore) SetDrsEnabled(v bool) {
	o.DrsEnabled = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ClusterDatastore) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ClusterDatastore) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ClusterDatastore) SetActive(v bool) {
	o.Active = &v
}

// GetAllowWrite returns the AllowWrite field value if set, zero value otherwise.
func (o *ClusterDatastore) GetAllowWrite() bool {
	if o == nil || o.AllowWrite == nil {
		var ret bool
		return ret
	}
	return *o.AllowWrite
}

// GetAllowWriteOk returns a tuple with the AllowWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetAllowWriteOk() (*bool, bool) {
	if o == nil || o.AllowWrite == nil {
		return nil, false
	}
	return o.AllowWrite, true
}

// HasAllowWrite returns a boolean if a field has been set.
func (o *ClusterDatastore) HasAllowWrite() bool {
	if o != nil && o.AllowWrite != nil {
		return true
	}

	return false
}

// SetAllowWrite gets a reference to the given bool and assigns it to the AllowWrite field.
func (o *ClusterDatastore) SetAllowWrite(v bool) {
	o.AllowWrite = &v
}

// GetDefaultStore returns the DefaultStore field value if set, zero value otherwise.
func (o *ClusterDatastore) GetDefaultStore() bool {
	if o == nil || o.DefaultStore == nil {
		var ret bool
		return ret
	}
	return *o.DefaultStore
}

// GetDefaultStoreOk returns a tuple with the DefaultStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetDefaultStoreOk() (*bool, bool) {
	if o == nil || o.DefaultStore == nil {
		return nil, false
	}
	return o.DefaultStore, true
}

// HasDefaultStore returns a boolean if a field has been set.
func (o *ClusterDatastore) HasDefaultStore() bool {
	if o != nil && o.DefaultStore != nil {
		return true
	}

	return false
}

// SetDefaultStore gets a reference to the given bool and assigns it to the DefaultStore field.
func (o *ClusterDatastore) SetDefaultStore(v bool) {
	o.DefaultStore = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *ClusterDatastore) GetOnline() bool {
	if o == nil || o.Online == nil {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetOnlineOk() (*bool, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *ClusterDatastore) HasOnline() bool {
	if o != nil && o.Online != nil {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *ClusterDatastore) SetOnline(v bool) {
	o.Online = &v
}

// GetAllowRead returns the AllowRead field value if set, zero value otherwise.
func (o *ClusterDatastore) GetAllowRead() bool {
	if o == nil || o.AllowRead == nil {
		var ret bool
		return ret
	}
	return *o.AllowRead
}

// GetAllowReadOk returns a tuple with the AllowRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetAllowReadOk() (*bool, bool) {
	if o == nil || o.AllowRead == nil {
		return nil, false
	}
	return o.AllowRead, true
}

// HasAllowRead returns a boolean if a field has been set.
func (o *ClusterDatastore) HasAllowRead() bool {
	if o != nil && o.AllowRead != nil {
		return true
	}

	return false
}

// SetAllowRead gets a reference to the given bool and assigns it to the AllowRead field.
func (o *ClusterDatastore) SetAllowRead(v bool) {
	o.AllowRead = &v
}

// GetAllowProvision returns the AllowProvision field value if set, zero value otherwise.
func (o *ClusterDatastore) GetAllowProvision() bool {
	if o == nil || o.AllowProvision == nil {
		var ret bool
		return ret
	}
	return *o.AllowProvision
}

// GetAllowProvisionOk returns a tuple with the AllowProvision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetAllowProvisionOk() (*bool, bool) {
	if o == nil || o.AllowProvision == nil {
		return nil, false
	}
	return o.AllowProvision, true
}

// HasAllowProvision returns a boolean if a field has been set.
func (o *ClusterDatastore) HasAllowProvision() bool {
	if o != nil && o.AllowProvision != nil {
		return true
	}

	return false
}

// SetAllowProvision gets a reference to the given bool and assigns it to the AllowProvision field.
func (o *ClusterDatastore) SetAllowProvision(v bool) {
	o.AllowProvision = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *ClusterDatastore) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *ClusterDatastore) HasRefType() bool {
	if o != nil && o.RefType != nil {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *ClusterDatastore) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *ClusterDatastore) GetRefId() int64 {
	if o == nil || o.RefId == nil {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetRefIdOk() (*int64, bool) {
	if o == nil || o.RefId == nil {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *ClusterDatastore) HasRefId() bool {
	if o != nil && o.RefId != nil {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *ClusterDatastore) SetRefId(v int64) {
	o.RefId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ClusterDatastore) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ClusterDatastore) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ClusterDatastore) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ClusterDatastore) GetZone() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Zone == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetZoneOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ClusterDatastore) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Zone field.
func (o *ClusterDatastore) SetZone(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Zone = &v
}

// GetZonePool returns the ZonePool field value if set, zero value otherwise.
func (o *ClusterDatastore) GetZonePool() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.ZonePool == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.ZonePool
}

// GetZonePoolOk returns a tuple with the ZonePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetZonePoolOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.ZonePool == nil {
		return nil, false
	}
	return o.ZonePool, true
}

// HasZonePool returns a boolean if a field has been set.
func (o *ClusterDatastore) HasZonePool() bool {
	if o != nil && o.ZonePool != nil {
		return true
	}

	return false
}

// SetZonePool gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the ZonePool field.
func (o *ClusterDatastore) SetZonePool(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.ZonePool = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ClusterDatastore) GetOwner() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites {
	if o == nil || o.Owner == nil {
		var ret ApiBlueprintsIdUpdatePermissionsResourcePermissionSites
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetOwnerOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ClusterDatastore) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ApiBlueprintsIdUpdatePermissionsResourcePermissionSites and assigns it to the Owner field.
func (o *ClusterDatastore) SetOwner(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites) {
	o.Owner = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ClusterDatastore) GetTenants() []InlineResponse20040AppDeployInstance {
	if o == nil || o.Tenants == nil {
		var ret []InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetTenantsOk() (*[]InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ClusterDatastore) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []InlineResponse20040AppDeployInstance and assigns it to the Tenants field.
func (o *ClusterDatastore) SetTenants(v []InlineResponse20040AppDeployInstance) {
	o.Tenants = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *ClusterDatastore) GetPermissions() ClusterDatastorePermissions {
	if o == nil || o.Permissions == nil {
		var ret ClusterDatastorePermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetPermissionsOk() (*ClusterDatastorePermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *ClusterDatastore) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given ClusterDatastorePermissions and assigns it to the Permissions field.
func (o *ClusterDatastore) SetPermissions(v ClusterDatastorePermissions) {
	o.Permissions = &v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *ClusterDatastore) GetDatastores() []map[string]interface{} {
	if o == nil || o.Datastores == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDatastore) GetDatastoresOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *ClusterDatastore) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []map[string]interface{} and assigns it to the Datastores field.
func (o *ClusterDatastore) SetDatastores(v []map[string]interface{}) {
	o.Datastores = &v
}

func (o ClusterDatastore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.StorageSize.IsSet() {
		toSerialize["storageSize"] = o.StorageSize.Get()
	}
	if o.FreeSpace.IsSet() {
		toSerialize["freeSpace"] = o.FreeSpace.Get()
	}
	if o.DrsEnabled != nil {
		toSerialize["drsEnabled"] = o.DrsEnabled
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.AllowWrite != nil {
		toSerialize["allowWrite"] = o.AllowWrite
	}
	if o.DefaultStore != nil {
		toSerialize["defaultStore"] = o.DefaultStore
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.AllowRead != nil {
		toSerialize["allowRead"] = o.AllowRead
	}
	if o.AllowProvision != nil {
		toSerialize["allowProvision"] = o.AllowProvision
	}
	if o.RefType != nil {
		toSerialize["refType"] = o.RefType
	}
	if o.RefId != nil {
		toSerialize["refId"] = o.RefId
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.ZonePool != nil {
		toSerialize["zonePool"] = o.ZonePool
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Datastores != nil {
		toSerialize["datastores"] = o.Datastores
	}
	return json.Marshal(toSerialize)
}

type NullableClusterDatastore struct {
	value *ClusterDatastore
	isSet bool
}

func (v NullableClusterDatastore) Get() *ClusterDatastore {
	return v.value
}

func (v *NullableClusterDatastore) Set(val *ClusterDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDatastore(val *ClusterDatastore) *NullableClusterDatastore {
	return &NullableClusterDatastore{value: val, isSet: true}
}

func (v NullableClusterDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


