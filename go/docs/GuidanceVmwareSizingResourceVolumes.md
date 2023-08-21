# GuidanceVmwareSizingResourceVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ControllerId** | Pointer to **int64** |  | [optional] 
**ControllerMountPoint** | Pointer to **string** |  | [optional] 
**Resizeable** | Pointer to **bool** |  | [optional] 
**PlanResizable** | Pointer to **bool** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**UnitNumber** | Pointer to **string** |  | [optional] 
**TypeId** | Pointer to **int64** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**DatastoreId** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewGuidanceVmwareSizingResourceVolumes

`func NewGuidanceVmwareSizingResourceVolumes() *GuidanceVmwareSizingResourceVolumes`

NewGuidanceVmwareSizingResourceVolumes instantiates a new GuidanceVmwareSizingResourceVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceVmwareSizingResourceVolumesWithDefaults

`func NewGuidanceVmwareSizingResourceVolumesWithDefaults() *GuidanceVmwareSizingResourceVolumes`

NewGuidanceVmwareSizingResourceVolumesWithDefaults instantiates a new GuidanceVmwareSizingResourceVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuidanceVmwareSizingResourceVolumes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuidanceVmwareSizingResourceVolumes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GuidanceVmwareSizingResourceVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GuidanceVmwareSizingResourceVolumes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuidanceVmwareSizingResourceVolumes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuidanceVmwareSizingResourceVolumes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetControllerId

`func (o *GuidanceVmwareSizingResourceVolumes) GetControllerId() int64`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetControllerIdOk() (*int64, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *GuidanceVmwareSizingResourceVolumes) SetControllerId(v int64)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *GuidanceVmwareSizingResourceVolumes) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerMountPoint

`func (o *GuidanceVmwareSizingResourceVolumes) GetControllerMountPoint() string`

GetControllerMountPoint returns the ControllerMountPoint field if non-nil, zero value otherwise.

### GetControllerMountPointOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetControllerMountPointOk() (*string, bool)`

GetControllerMountPointOk returns a tuple with the ControllerMountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerMountPoint

`func (o *GuidanceVmwareSizingResourceVolumes) SetControllerMountPoint(v string)`

SetControllerMountPoint sets ControllerMountPoint field to given value.

### HasControllerMountPoint

`func (o *GuidanceVmwareSizingResourceVolumes) HasControllerMountPoint() bool`

HasControllerMountPoint returns a boolean if a field has been set.

### GetResizeable

`func (o *GuidanceVmwareSizingResourceVolumes) GetResizeable() bool`

GetResizeable returns the Resizeable field if non-nil, zero value otherwise.

### GetResizeableOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetResizeableOk() (*bool, bool)`

GetResizeableOk returns a tuple with the Resizeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeable

`func (o *GuidanceVmwareSizingResourceVolumes) SetResizeable(v bool)`

SetResizeable sets Resizeable field to given value.

### HasResizeable

`func (o *GuidanceVmwareSizingResourceVolumes) HasResizeable() bool`

HasResizeable returns a boolean if a field has been set.

### GetPlanResizable

`func (o *GuidanceVmwareSizingResourceVolumes) GetPlanResizable() bool`

GetPlanResizable returns the PlanResizable field if non-nil, zero value otherwise.

### GetPlanResizableOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetPlanResizableOk() (*bool, bool)`

GetPlanResizableOk returns a tuple with the PlanResizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanResizable

`func (o *GuidanceVmwareSizingResourceVolumes) SetPlanResizable(v bool)`

SetPlanResizable sets PlanResizable field to given value.

### HasPlanResizable

`func (o *GuidanceVmwareSizingResourceVolumes) HasPlanResizable() bool`

HasPlanResizable returns a boolean if a field has been set.

### GetRootVolume

`func (o *GuidanceVmwareSizingResourceVolumes) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *GuidanceVmwareSizingResourceVolumes) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *GuidanceVmwareSizingResourceVolumes) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetUnitNumber

`func (o *GuidanceVmwareSizingResourceVolumes) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *GuidanceVmwareSizingResourceVolumes) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *GuidanceVmwareSizingResourceVolumes) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### GetTypeId

`func (o *GuidanceVmwareSizingResourceVolumes) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *GuidanceVmwareSizingResourceVolumes) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *GuidanceVmwareSizingResourceVolumes) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetDatastoreId

`func (o *GuidanceVmwareSizingResourceVolumes) GetDatastoreId() int64`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetDatastoreIdOk() (*int64, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *GuidanceVmwareSizingResourceVolumes) SetDatastoreId(v int64)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *GuidanceVmwareSizingResourceVolumes) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GuidanceVmwareSizingResourceVolumes) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GuidanceVmwareSizingResourceVolumes) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GuidanceVmwareSizingResourceVolumes) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *GuidanceVmwareSizingResourceVolumes) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GuidanceVmwareSizingResourceVolumes) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GuidanceVmwareSizingResourceVolumes) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *GuidanceVmwareSizingResourceVolumes) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *GuidanceVmwareSizingResourceVolumes) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *GuidanceVmwareSizingResourceVolumes) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetUuid

`func (o *GuidanceVmwareSizingResourceVolumes) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GuidanceVmwareSizingResourceVolumes) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GuidanceVmwareSizingResourceVolumes) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GuidanceVmwareSizingResourceVolumes) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


