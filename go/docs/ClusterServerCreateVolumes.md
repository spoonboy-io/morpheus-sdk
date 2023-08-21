# ClusterServerCreateVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the LV configuration being created | [optional] [default to -1]
**RootVolume** | Pointer to **bool** | If set to false then a non-root LV will be created | [optional] [default to true]
**Name** | **string** | Name/type of the LV being created | [default to "root"]
**Size** | Pointer to **int64** | Size of the LV to be created in GBs  Default is from the service plan  | [optional] 
**SizeId** | Pointer to **NullableString** | Can be used to select pre-existing LV choices from Morpheus | [optional] 
**StorageType** | Pointer to **int64** | Identifier for LV type | [optional] 
**DatastoreId** | **NullableString** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 

## Methods

### NewClusterServerCreateVolumes

`func NewClusterServerCreateVolumes(name string, datastoreId NullableString, ) *ClusterServerCreateVolumes`

NewClusterServerCreateVolumes instantiates a new ClusterServerCreateVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateVolumesWithDefaults

`func NewClusterServerCreateVolumesWithDefaults() *ClusterServerCreateVolumes`

NewClusterServerCreateVolumesWithDefaults instantiates a new ClusterServerCreateVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterServerCreateVolumes) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterServerCreateVolumes) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterServerCreateVolumes) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterServerCreateVolumes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootVolume

`func (o *ClusterServerCreateVolumes) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ClusterServerCreateVolumes) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ClusterServerCreateVolumes) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ClusterServerCreateVolumes) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *ClusterServerCreateVolumes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterServerCreateVolumes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterServerCreateVolumes) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ClusterServerCreateVolumes) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ClusterServerCreateVolumes) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ClusterServerCreateVolumes) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ClusterServerCreateVolumes) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeId

`func (o *ClusterServerCreateVolumes) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *ClusterServerCreateVolumes) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *ClusterServerCreateVolumes) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *ClusterServerCreateVolumes) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### SetSizeIdNil

`func (o *ClusterServerCreateVolumes) SetSizeIdNil(b bool)`

 SetSizeIdNil sets the value for SizeId to be an explicit nil

### UnsetSizeId
`func (o *ClusterServerCreateVolumes) UnsetSizeId()`

UnsetSizeId ensures that no value is present for SizeId, not even an explicit nil
### GetStorageType

`func (o *ClusterServerCreateVolumes) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ClusterServerCreateVolumes) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ClusterServerCreateVolumes) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ClusterServerCreateVolumes) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ClusterServerCreateVolumes) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ClusterServerCreateVolumes) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ClusterServerCreateVolumes) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.


### SetDatastoreIdNil

`func (o *ClusterServerCreateVolumes) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ClusterServerCreateVolumes) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


