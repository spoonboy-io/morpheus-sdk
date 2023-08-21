# VdiPoolConfigVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootVolume** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**VolumeCustomizable** | Pointer to **bool** |  | [optional] 
**ReadonlyName** | Pointer to **bool** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StorageType** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**DatastoreId** | Pointer to **string** |  | [optional] 
**VId** | Pointer to **int64** |  | [optional] 

## Methods

### NewVdiPoolConfigVolumes

`func NewVdiPoolConfigVolumes() *VdiPoolConfigVolumes`

NewVdiPoolConfigVolumes instantiates a new VdiPoolConfigVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolConfigVolumesWithDefaults

`func NewVdiPoolConfigVolumesWithDefaults() *VdiPoolConfigVolumes`

NewVdiPoolConfigVolumesWithDefaults instantiates a new VdiPoolConfigVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootVolume

`func (o *VdiPoolConfigVolumes) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *VdiPoolConfigVolumes) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *VdiPoolConfigVolumes) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *VdiPoolConfigVolumes) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *VdiPoolConfigVolumes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdiPoolConfigVolumes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdiPoolConfigVolumes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdiPoolConfigVolumes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxStorage

`func (o *VdiPoolConfigVolumes) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *VdiPoolConfigVolumes) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *VdiPoolConfigVolumes) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *VdiPoolConfigVolumes) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetVolumeCustomizable

`func (o *VdiPoolConfigVolumes) GetVolumeCustomizable() bool`

GetVolumeCustomizable returns the VolumeCustomizable field if non-nil, zero value otherwise.

### GetVolumeCustomizableOk

`func (o *VdiPoolConfigVolumes) GetVolumeCustomizableOk() (*bool, bool)`

GetVolumeCustomizableOk returns a tuple with the VolumeCustomizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCustomizable

`func (o *VdiPoolConfigVolumes) SetVolumeCustomizable(v bool)`

SetVolumeCustomizable sets VolumeCustomizable field to given value.

### HasVolumeCustomizable

`func (o *VdiPoolConfigVolumes) HasVolumeCustomizable() bool`

HasVolumeCustomizable returns a boolean if a field has been set.

### GetReadonlyName

`func (o *VdiPoolConfigVolumes) GetReadonlyName() bool`

GetReadonlyName returns the ReadonlyName field if non-nil, zero value otherwise.

### GetReadonlyNameOk

`func (o *VdiPoolConfigVolumes) GetReadonlyNameOk() (*bool, bool)`

GetReadonlyNameOk returns a tuple with the ReadonlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyName

`func (o *VdiPoolConfigVolumes) SetReadonlyName(v bool)`

SetReadonlyName sets ReadonlyName field to given value.

### HasReadonlyName

`func (o *VdiPoolConfigVolumes) HasReadonlyName() bool`

HasReadonlyName returns a boolean if a field has been set.

### GetSize

`func (o *VdiPoolConfigVolumes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VdiPoolConfigVolumes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VdiPoolConfigVolumes) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VdiPoolConfigVolumes) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageType

`func (o *VdiPoolConfigVolumes) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *VdiPoolConfigVolumes) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *VdiPoolConfigVolumes) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *VdiPoolConfigVolumes) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *VdiPoolConfigVolumes) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *VdiPoolConfigVolumes) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *VdiPoolConfigVolumes) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *VdiPoolConfigVolumes) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *VdiPoolConfigVolumes) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *VdiPoolConfigVolumes) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetDatastoreId

`func (o *VdiPoolConfigVolumes) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *VdiPoolConfigVolumes) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *VdiPoolConfigVolumes) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *VdiPoolConfigVolumes) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.

### GetVId

`func (o *VdiPoolConfigVolumes) GetVId() int64`

GetVId returns the VId field if non-nil, zero value otherwise.

### GetVIdOk

`func (o *VdiPoolConfigVolumes) GetVIdOk() (*int64, bool)`

GetVIdOk returns a tuple with the VId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVId

`func (o *VdiPoolConfigVolumes) SetVId(v int64)`

SetVId sets VId field to given value.

### HasVId

`func (o *VdiPoolConfigVolumes) HasVId() bool`

HasVId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


