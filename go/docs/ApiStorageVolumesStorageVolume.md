# ApiStorageVolumesStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the storage volume | 
**Type** | **string** | Storage Type Code or ID | 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | [**ApiStorageVolumesStorageVolumeStorageServer**](_api_storage_volumes_storageVolume_storageServer.md) |  | 
**StorageGroup** | [**ApiStorageVolumesStorageVolumeStorageServer**](_api_storage_volumes_storageVolume_storageServer.md) |  | 

## Methods

### NewApiStorageVolumesStorageVolume

`func NewApiStorageVolumesStorageVolume(name string, type_ string, storageServer ApiStorageVolumesStorageVolumeStorageServer, storageGroup ApiStorageVolumesStorageVolumeStorageServer, ) *ApiStorageVolumesStorageVolume`

NewApiStorageVolumesStorageVolume instantiates a new ApiStorageVolumesStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStorageVolumesStorageVolumeWithDefaults

`func NewApiStorageVolumesStorageVolumeWithDefaults() *ApiStorageVolumesStorageVolume`

NewApiStorageVolumesStorageVolumeWithDefaults instantiates a new ApiStorageVolumesStorageVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiStorageVolumesStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStorageVolumesStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStorageVolumesStorageVolume) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApiStorageVolumesStorageVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStorageVolumesStorageVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStorageVolumesStorageVolume) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *ApiStorageVolumesStorageVolume) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiStorageVolumesStorageVolume) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiStorageVolumesStorageVolume) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiStorageVolumesStorageVolume) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *ApiStorageVolumesStorageVolume) GetStorageServer() ApiStorageVolumesStorageVolumeStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ApiStorageVolumesStorageVolume) GetStorageServerOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ApiStorageVolumesStorageVolume) SetStorageServer(v ApiStorageVolumesStorageVolumeStorageServer)`

SetStorageServer sets StorageServer field to given value.


### GetStorageGroup

`func (o *ApiStorageVolumesStorageVolume) GetStorageGroup() ApiStorageVolumesStorageVolumeStorageServer`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *ApiStorageVolumesStorageVolume) GetStorageGroupOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *ApiStorageVolumesStorageVolume) SetStorageGroup(v ApiStorageVolumesStorageVolumeStorageServer)`

SetStorageGroup sets StorageGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


