# ImageBuildConfigVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **int64** |  | [optional] 
**DatastoreId** | Pointer to **string** |  | [optional] 

## Methods

### NewImageBuildConfigVolumes

`func NewImageBuildConfigVolumes() *ImageBuildConfigVolumes`

NewImageBuildConfigVolumes instantiates a new ImageBuildConfigVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildConfigVolumesWithDefaults

`func NewImageBuildConfigVolumesWithDefaults() *ImageBuildConfigVolumes`

NewImageBuildConfigVolumesWithDefaults instantiates a new ImageBuildConfigVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageBuildConfigVolumes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageBuildConfigVolumes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageBuildConfigVolumes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImageBuildConfigVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *ImageBuildConfigVolumes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageBuildConfigVolumes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageBuildConfigVolumes) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageBuildConfigVolumes) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *ImageBuildConfigVolumes) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ImageBuildConfigVolumes) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ImageBuildConfigVolumes) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ImageBuildConfigVolumes) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *ImageBuildConfigVolumes) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *ImageBuildConfigVolumes) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetName

`func (o *ImageBuildConfigVolumes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBuildConfigVolumes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBuildConfigVolumes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBuildConfigVolumes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootVolume

`func (o *ImageBuildConfigVolumes) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ImageBuildConfigVolumes) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ImageBuildConfigVolumes) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ImageBuildConfigVolumes) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetStorageType

`func (o *ImageBuildConfigVolumes) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ImageBuildConfigVolumes) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ImageBuildConfigVolumes) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ImageBuildConfigVolumes) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ImageBuildConfigVolumes) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ImageBuildConfigVolumes) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ImageBuildConfigVolumes) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *ImageBuildConfigVolumes) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


