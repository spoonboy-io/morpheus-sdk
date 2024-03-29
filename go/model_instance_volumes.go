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

// InstanceVolumes struct for InstanceVolumes
type InstanceVolumes struct {
	ControllerId NullableInt64 `json:"controllerId,omitempty"`
	DatastoreId NullableString `json:"datastoreId,omitempty"`
	DisplayOrder *int64 `json:"displayOrder,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	MaxIOPS NullableString `json:"maxIOPS,omitempty"`
	MaxStorage *int64 `json:"maxStorage,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	Resizeable *bool `json:"resizeable,omitempty"`
	PlanResizable *bool `json:"planResizable,omitempty"`
	RootVolume *bool `json:"rootVolume,omitempty"`
	Size *int64 `json:"size,omitempty"`
	StorageType *int64 `json:"storageType,omitempty"`
	UnitNumber NullableString `json:"unitNumber,omitempty"`
	ControllerMountPoint NullableString `json:"controllerMountPoint,omitempty"`
}

// NewInstanceVolumes instantiates a new InstanceVolumes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceVolumes() *InstanceVolumes {
	this := InstanceVolumes{}
	return &this
}

// NewInstanceVolumesWithDefaults instantiates a new InstanceVolumes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceVolumesWithDefaults() *InstanceVolumes {
	this := InstanceVolumes{}
	return &this
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceVolumes) GetControllerId() int64 {
	if o == nil || o.ControllerId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ControllerId.Get()
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceVolumes) GetControllerIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ControllerId.Get(), o.ControllerId.IsSet()
}

// HasControllerId returns a boolean if a field has been set.
func (o *InstanceVolumes) HasControllerId() bool {
	if o != nil && o.ControllerId.IsSet() {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given NullableInt64 and assigns it to the ControllerId field.
func (o *InstanceVolumes) SetControllerId(v int64) {
	o.ControllerId.Set(&v)
}
// SetControllerIdNil sets the value for ControllerId to be an explicit nil
func (o *InstanceVolumes) SetControllerIdNil() {
	o.ControllerId.Set(nil)
}

// UnsetControllerId ensures that no value is present for ControllerId, not even an explicit nil
func (o *InstanceVolumes) UnsetControllerId() {
	o.ControllerId.Unset()
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceVolumes) GetDatastoreId() string {
	if o == nil || o.DatastoreId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatastoreId.Get()
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceVolumes) GetDatastoreIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatastoreId.Get(), o.DatastoreId.IsSet()
}

// HasDatastoreId returns a boolean if a field has been set.
func (o *InstanceVolumes) HasDatastoreId() bool {
	if o != nil && o.DatastoreId.IsSet() {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given NullableString and assigns it to the DatastoreId field.
func (o *InstanceVolumes) SetDatastoreId(v string) {
	o.DatastoreId.Set(&v)
}
// SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil
func (o *InstanceVolumes) SetDatastoreIdNil() {
	o.DatastoreId.Set(nil)
}

// UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil
func (o *InstanceVolumes) UnsetDatastoreId() {
	o.DatastoreId.Unset()
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *InstanceVolumes) GetDisplayOrder() int64 {
	if o == nil || o.DisplayOrder == nil {
		var ret int64
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetDisplayOrderOk() (*int64, bool) {
	if o == nil || o.DisplayOrder == nil {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *InstanceVolumes) HasDisplayOrder() bool {
	if o != nil && o.DisplayOrder != nil {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int64 and assigns it to the DisplayOrder field.
func (o *InstanceVolumes) SetDisplayOrder(v int64) {
	o.DisplayOrder = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstanceVolumes) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstanceVolumes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InstanceVolumes) SetId(v int64) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *InstanceVolumes) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *InstanceVolumes) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *InstanceVolumes) SetUuid(v string) {
	o.Uuid = &v
}

// GetMaxIOPS returns the MaxIOPS field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceVolumes) GetMaxIOPS() string {
	if o == nil || o.MaxIOPS.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxIOPS.Get()
}

// GetMaxIOPSOk returns a tuple with the MaxIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceVolumes) GetMaxIOPSOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxIOPS.Get(), o.MaxIOPS.IsSet()
}

// HasMaxIOPS returns a boolean if a field has been set.
func (o *InstanceVolumes) HasMaxIOPS() bool {
	if o != nil && o.MaxIOPS.IsSet() {
		return true
	}

	return false
}

// SetMaxIOPS gets a reference to the given NullableString and assigns it to the MaxIOPS field.
func (o *InstanceVolumes) SetMaxIOPS(v string) {
	o.MaxIOPS.Set(&v)
}
// SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil
func (o *InstanceVolumes) SetMaxIOPSNil() {
	o.MaxIOPS.Set(nil)
}

// UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
func (o *InstanceVolumes) UnsetMaxIOPS() {
	o.MaxIOPS.Unset()
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *InstanceVolumes) GetMaxStorage() int64 {
	if o == nil || o.MaxStorage == nil {
		var ret int64
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetMaxStorageOk() (*int64, bool) {
	if o == nil || o.MaxStorage == nil {
		return nil, false
	}
	return o.MaxStorage, true
}

// HasMaxStorage returns a boolean if a field has been set.
func (o *InstanceVolumes) HasMaxStorage() bool {
	if o != nil && o.MaxStorage != nil {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given int64 and assigns it to the MaxStorage field.
func (o *InstanceVolumes) SetMaxStorage(v int64) {
	o.MaxStorage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceVolumes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceVolumes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceVolumes) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InstanceVolumes) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InstanceVolumes) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InstanceVolumes) SetShortName(v string) {
	o.ShortName = &v
}

// GetResizeable returns the Resizeable field value if set, zero value otherwise.
func (o *InstanceVolumes) GetResizeable() bool {
	if o == nil || o.Resizeable == nil {
		var ret bool
		return ret
	}
	return *o.Resizeable
}

// GetResizeableOk returns a tuple with the Resizeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetResizeableOk() (*bool, bool) {
	if o == nil || o.Resizeable == nil {
		return nil, false
	}
	return o.Resizeable, true
}

// HasResizeable returns a boolean if a field has been set.
func (o *InstanceVolumes) HasResizeable() bool {
	if o != nil && o.Resizeable != nil {
		return true
	}

	return false
}

// SetResizeable gets a reference to the given bool and assigns it to the Resizeable field.
func (o *InstanceVolumes) SetResizeable(v bool) {
	o.Resizeable = &v
}

// GetPlanResizable returns the PlanResizable field value if set, zero value otherwise.
func (o *InstanceVolumes) GetPlanResizable() bool {
	if o == nil || o.PlanResizable == nil {
		var ret bool
		return ret
	}
	return *o.PlanResizable
}

// GetPlanResizableOk returns a tuple with the PlanResizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetPlanResizableOk() (*bool, bool) {
	if o == nil || o.PlanResizable == nil {
		return nil, false
	}
	return o.PlanResizable, true
}

// HasPlanResizable returns a boolean if a field has been set.
func (o *InstanceVolumes) HasPlanResizable() bool {
	if o != nil && o.PlanResizable != nil {
		return true
	}

	return false
}

// SetPlanResizable gets a reference to the given bool and assigns it to the PlanResizable field.
func (o *InstanceVolumes) SetPlanResizable(v bool) {
	o.PlanResizable = &v
}

// GetRootVolume returns the RootVolume field value if set, zero value otherwise.
func (o *InstanceVolumes) GetRootVolume() bool {
	if o == nil || o.RootVolume == nil {
		var ret bool
		return ret
	}
	return *o.RootVolume
}

// GetRootVolumeOk returns a tuple with the RootVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetRootVolumeOk() (*bool, bool) {
	if o == nil || o.RootVolume == nil {
		return nil, false
	}
	return o.RootVolume, true
}

// HasRootVolume returns a boolean if a field has been set.
func (o *InstanceVolumes) HasRootVolume() bool {
	if o != nil && o.RootVolume != nil {
		return true
	}

	return false
}

// SetRootVolume gets a reference to the given bool and assigns it to the RootVolume field.
func (o *InstanceVolumes) SetRootVolume(v bool) {
	o.RootVolume = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InstanceVolumes) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InstanceVolumes) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *InstanceVolumes) SetSize(v int64) {
	o.Size = &v
}

// GetStorageType returns the StorageType field value if set, zero value otherwise.
func (o *InstanceVolumes) GetStorageType() int64 {
	if o == nil || o.StorageType == nil {
		var ret int64
		return ret
	}
	return *o.StorageType
}

// GetStorageTypeOk returns a tuple with the StorageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceVolumes) GetStorageTypeOk() (*int64, bool) {
	if o == nil || o.StorageType == nil {
		return nil, false
	}
	return o.StorageType, true
}

// HasStorageType returns a boolean if a field has been set.
func (o *InstanceVolumes) HasStorageType() bool {
	if o != nil && o.StorageType != nil {
		return true
	}

	return false
}

// SetStorageType gets a reference to the given int64 and assigns it to the StorageType field.
func (o *InstanceVolumes) SetStorageType(v int64) {
	o.StorageType = &v
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceVolumes) GetUnitNumber() string {
	if o == nil || o.UnitNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnitNumber.Get()
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceVolumes) GetUnitNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitNumber.Get(), o.UnitNumber.IsSet()
}

// HasUnitNumber returns a boolean if a field has been set.
func (o *InstanceVolumes) HasUnitNumber() bool {
	if o != nil && o.UnitNumber.IsSet() {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given NullableString and assigns it to the UnitNumber field.
func (o *InstanceVolumes) SetUnitNumber(v string) {
	o.UnitNumber.Set(&v)
}
// SetUnitNumberNil sets the value for UnitNumber to be an explicit nil
func (o *InstanceVolumes) SetUnitNumberNil() {
	o.UnitNumber.Set(nil)
}

// UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
func (o *InstanceVolumes) UnsetUnitNumber() {
	o.UnitNumber.Unset()
}

// GetControllerMountPoint returns the ControllerMountPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstanceVolumes) GetControllerMountPoint() string {
	if o == nil || o.ControllerMountPoint.Get() == nil {
		var ret string
		return ret
	}
	return *o.ControllerMountPoint.Get()
}

// GetControllerMountPointOk returns a tuple with the ControllerMountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceVolumes) GetControllerMountPointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ControllerMountPoint.Get(), o.ControllerMountPoint.IsSet()
}

// HasControllerMountPoint returns a boolean if a field has been set.
func (o *InstanceVolumes) HasControllerMountPoint() bool {
	if o != nil && o.ControllerMountPoint.IsSet() {
		return true
	}

	return false
}

// SetControllerMountPoint gets a reference to the given NullableString and assigns it to the ControllerMountPoint field.
func (o *InstanceVolumes) SetControllerMountPoint(v string) {
	o.ControllerMountPoint.Set(&v)
}
// SetControllerMountPointNil sets the value for ControllerMountPoint to be an explicit nil
func (o *InstanceVolumes) SetControllerMountPointNil() {
	o.ControllerMountPoint.Set(nil)
}

// UnsetControllerMountPoint ensures that no value is present for ControllerMountPoint, not even an explicit nil
func (o *InstanceVolumes) UnsetControllerMountPoint() {
	o.ControllerMountPoint.Unset()
}

func (o InstanceVolumes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ControllerId.IsSet() {
		toSerialize["controllerId"] = o.ControllerId.Get()
	}
	if o.DatastoreId.IsSet() {
		toSerialize["datastoreId"] = o.DatastoreId.Get()
	}
	if o.DisplayOrder != nil {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.MaxIOPS.IsSet() {
		toSerialize["maxIOPS"] = o.MaxIOPS.Get()
	}
	if o.MaxStorage != nil {
		toSerialize["maxStorage"] = o.MaxStorage
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.Resizeable != nil {
		toSerialize["resizeable"] = o.Resizeable
	}
	if o.PlanResizable != nil {
		toSerialize["planResizable"] = o.PlanResizable
	}
	if o.RootVolume != nil {
		toSerialize["rootVolume"] = o.RootVolume
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.StorageType != nil {
		toSerialize["storageType"] = o.StorageType
	}
	if o.UnitNumber.IsSet() {
		toSerialize["unitNumber"] = o.UnitNumber.Get()
	}
	if o.ControllerMountPoint.IsSet() {
		toSerialize["controllerMountPoint"] = o.ControllerMountPoint.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInstanceVolumes struct {
	value *InstanceVolumes
	isSet bool
}

func (v NullableInstanceVolumes) Get() *InstanceVolumes {
	return v.value
}

func (v *NullableInstanceVolumes) Set(val *InstanceVolumes) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceVolumes) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceVolumes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceVolumes(val *InstanceVolumes) *NullableInstanceVolumes {
	return &NullableInstanceVolumes{value: val, isSet: true}
}

func (v NullableInstanceVolumes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceVolumes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


