# ApiStorageVolumesIdStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name scoped to your account for the storage volume | [optional] 
**Type** | Pointer to **string** | Storage Type Code or ID | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | Pointer to [**ApiStorageVolumesStorageVolumeStorageServer**](_api_storage_volumes_storageVolume_storageServer.md) |  | [optional] 
**StorageGroup** | Pointer to [**ApiStorageVolumesStorageVolumeStorageServer**](_api_storage_volumes_storageVolume_storageServer.md) |  | [optional] 

## Methods

### NewApiStorageVolumesIdStorageVolume

`func NewApiStorageVolumesIdStorageVolume() *ApiStorageVolumesIdStorageVolume`

NewApiStorageVolumesIdStorageVolume instantiates a new ApiStorageVolumesIdStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStorageVolumesIdStorageVolumeWithDefaults

`func NewApiStorageVolumesIdStorageVolumeWithDefaults() *ApiStorageVolumesIdStorageVolume`

NewApiStorageVolumesIdStorageVolumeWithDefaults instantiates a new ApiStorageVolumesIdStorageVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiStorageVolumesIdStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStorageVolumesIdStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStorageVolumesIdStorageVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiStorageVolumesIdStorageVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ApiStorageVolumesIdStorageVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStorageVolumesIdStorageVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStorageVolumesIdStorageVolume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiStorageVolumesIdStorageVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *ApiStorageVolumesIdStorageVolume) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiStorageVolumesIdStorageVolume) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiStorageVolumesIdStorageVolume) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiStorageVolumesIdStorageVolume) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *ApiStorageVolumesIdStorageVolume) GetStorageServer() ApiStorageVolumesStorageVolumeStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *ApiStorageVolumesIdStorageVolume) GetStorageServerOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *ApiStorageVolumesIdStorageVolume) SetStorageServer(v ApiStorageVolumesStorageVolumeStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *ApiStorageVolumesIdStorageVolume) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetStorageGroup

`func (o *ApiStorageVolumesIdStorageVolume) GetStorageGroup() ApiStorageVolumesStorageVolumeStorageServer`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *ApiStorageVolumesIdStorageVolume) GetStorageGroupOk() (*ApiStorageVolumesStorageVolumeStorageServer, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *ApiStorageVolumesIdStorageVolume) SetStorageGroup(v ApiStorageVolumesStorageVolumeStorageServer)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *ApiStorageVolumesIdStorageVolume) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


