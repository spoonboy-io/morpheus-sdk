# ImageBuildsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ImageBuildsConfigInstance**](imageBuilds_config_instance.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]ImageBuildsConfigNetworkInterfaces**](ImageBuildsConfigNetworkInterfaces.md) |  | [optional] 
**Volumes** | Pointer to [**[]ImageBuildsConfigVolumes**](ImageBuildsConfigVolumes.md) |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**ImageBuildsConfigConfig**](imageBuilds_config_config.md) |  | [optional] 
**Plan** | Pointer to [**ImageBuildsConfigPlan**](imageBuilds_config_plan.md) |  | [optional] 

## Methods

### NewImageBuildsConfig

`func NewImageBuildsConfig() *ImageBuildsConfig`

NewImageBuildsConfig instantiates a new ImageBuildsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildsConfigWithDefaults

`func NewImageBuildsConfigWithDefaults() *ImageBuildsConfig`

NewImageBuildsConfigWithDefaults instantiates a new ImageBuildsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ImageBuildsConfig) GetInstance() ImageBuildsConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ImageBuildsConfig) GetInstanceOk() (*ImageBuildsConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ImageBuildsConfig) SetInstance(v ImageBuildsConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ImageBuildsConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ImageBuildsConfig) GetNetworkInterfaces() []ImageBuildsConfigNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ImageBuildsConfig) GetNetworkInterfacesOk() (*[]ImageBuildsConfigNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ImageBuildsConfig) SetNetworkInterfaces(v []ImageBuildsConfigNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ImageBuildsConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVolumes

`func (o *ImageBuildsConfig) GetVolumes() []ImageBuildsConfigVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ImageBuildsConfig) GetVolumesOk() (*[]ImageBuildsConfigVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ImageBuildsConfig) SetVolumes(v []ImageBuildsConfigVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ImageBuildsConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ImageBuildsConfig) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ImageBuildsConfig) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ImageBuildsConfig) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ImageBuildsConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetZoneId

`func (o *ImageBuildsConfig) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ImageBuildsConfig) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ImageBuildsConfig) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ImageBuildsConfig) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetConfig

`func (o *ImageBuildsConfig) GetConfig() ImageBuildsConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ImageBuildsConfig) GetConfigOk() (*ImageBuildsConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ImageBuildsConfig) SetConfig(v ImageBuildsConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ImageBuildsConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPlan

`func (o *ImageBuildsConfig) GetPlan() ImageBuildsConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ImageBuildsConfig) GetPlanOk() (*ImageBuildsConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ImageBuildsConfig) SetPlan(v ImageBuildsConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ImageBuildsConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


