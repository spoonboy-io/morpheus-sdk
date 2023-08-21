# InstanceCreateVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the LV configuration being created. | [optional] [default to -1]
**RootVolume** | Pointer to **bool** | If set to false then a non-root LV will be created. | [optional] [default to true]
**Name** | Pointer to **string** | Name/type of the LV being created. | [optional] [default to "root"]
**Size** | Pointer to **int64** | Size of the LV to be created in GBs.  Uses default from service plan. | [optional] 
**SizeId** | Pointer to **NullableInt64** | Can be used to select pre-existing LV choices from Morpheus. | [optional] 
**StorageType** | Pointer to **NullableInt64** | Identifier for LV type | [optional] 
**DatastoreId** | Pointer to [**AnyOfstringlong**](anyOf&lt;string,long&gt;.md) | The ID of the specific datastore. Auto selection can be specified as auto or autoCluster (for clusters). | [optional] 

## Methods

### NewInstanceCreateVolume

`func NewInstanceCreateVolume() *InstanceCreateVolume`

NewInstanceCreateVolume instantiates a new InstanceCreateVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateVolumeWithDefaults

`func NewInstanceCreateVolumeWithDefaults() *InstanceCreateVolume`

NewInstanceCreateVolumeWithDefaults instantiates a new InstanceCreateVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceCreateVolume) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceCreateVolume) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceCreateVolume) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceCreateVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootVolume

`func (o *InstanceCreateVolume) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *InstanceCreateVolume) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *InstanceCreateVolume) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *InstanceCreateVolume) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *InstanceCreateVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCreateVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCreateVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCreateVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *InstanceCreateVolume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InstanceCreateVolume) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InstanceCreateVolume) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *InstanceCreateVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeId

`func (o *InstanceCreateVolume) GetSizeId() int64`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *InstanceCreateVolume) GetSizeIdOk() (*int64, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *InstanceCreateVolume) SetSizeId(v int64)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *InstanceCreateVolume) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### SetSizeIdNil

`func (o *InstanceCreateVolume) SetSizeIdNil(b bool)`

 SetSizeIdNil sets the value for SizeId to be an explicit nil

### UnsetSizeId
`func (o *InstanceCreateVolume) UnsetSizeId()`

UnsetSizeId ensures that no value is present for SizeId, not even an explicit nil
### GetStorageType

`func (o *InstanceCreateVolume) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InstanceCreateVolume) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InstanceCreateVolume) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *InstanceCreateVolume) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### SetStorageTypeNil

`func (o *InstanceCreateVolume) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *InstanceCreateVolume) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetDatastoreId

`func (o *InstanceCreateVolume) GetDatastoreId() AnyOfstringlong`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *InstanceCreateVolume) GetDatastoreIdOk() (*AnyOfstringlong, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *InstanceCreateVolume) SetDatastoreId(v AnyOfstringlong)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *InstanceCreateVolume) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


