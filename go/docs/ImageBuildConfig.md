# ImageBuildConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ImageBuildConfigInstance**](imageBuild_config_instance.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]ImageBuildConfigNetworkInterfaces**](ImageBuildConfigNetworkInterfaces.md) |  | [optional] 
**Volumes** | Pointer to [**[]ImageBuildConfigVolumes**](ImageBuildConfigVolumes.md) |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**ImageBuildConfigConfig**](imageBuild_config_config.md) |  | [optional] 
**Plan** | Pointer to [**ImageBuildsConfigPlan**](imageBuilds_config_plan.md) |  | [optional] 

## Methods

### NewImageBuildConfig

`func NewImageBuildConfig() *ImageBuildConfig`

NewImageBuildConfig instantiates a new ImageBuildConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildConfigWithDefaults

`func NewImageBuildConfigWithDefaults() *ImageBuildConfig`

NewImageBuildConfigWithDefaults instantiates a new ImageBuildConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ImageBuildConfig) GetInstance() ImageBuildConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ImageBuildConfig) GetInstanceOk() (*ImageBuildConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ImageBuildConfig) SetInstance(v ImageBuildConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ImageBuildConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ImageBuildConfig) GetNetworkInterfaces() []ImageBuildConfigNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ImageBuildConfig) GetNetworkInterfacesOk() (*[]ImageBuildConfigNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ImageBuildConfig) SetNetworkInterfaces(v []ImageBuildConfigNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ImageBuildConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVolumes

`func (o *ImageBuildConfig) GetVolumes() []ImageBuildConfigVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ImageBuildConfig) GetVolumesOk() (*[]ImageBuildConfigVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ImageBuildConfig) SetVolumes(v []ImageBuildConfigVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ImageBuildConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *ImageBuildConfig) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *ImageBuildConfig) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *ImageBuildConfig) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *ImageBuildConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetZoneId

`func (o *ImageBuildConfig) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ImageBuildConfig) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ImageBuildConfig) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ImageBuildConfig) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetConfig

`func (o *ImageBuildConfig) GetConfig() ImageBuildConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ImageBuildConfig) GetConfigOk() (*ImageBuildConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ImageBuildConfig) SetConfig(v ImageBuildConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ImageBuildConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPlan

`func (o *ImageBuildConfig) GetPlan() ImageBuildsConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ImageBuildConfig) GetPlanOk() (*ImageBuildsConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ImageBuildConfig) SetPlan(v ImageBuildsConfigPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ImageBuildConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


