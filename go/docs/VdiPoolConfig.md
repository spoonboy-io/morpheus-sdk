# VdiPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**InstanceServicePlanAutoOptions**](instanceServicePlan_autoOptions.md) |  | [optional] 
**Cloud** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to [**VdiPoolConfigInstance**](vdiPool_config_instance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**VdiPoolConfigConfig**](vdiPool_config_config.md) |  | [optional] 
**Volumes** | Pointer to [**[]VdiPoolConfigVolumes**](VdiPoolConfigVolumes.md) |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Layout** | Pointer to [**InstanceConfigLayout**](instance_config_layout.md) |  | [optional] 
**StorageControllers** | Pointer to [**[]VdiPoolConfigStorageControllers**](VdiPoolConfigStorageControllers.md) |  | [optional] 
**Plan** | Pointer to [**InstanceConfigLayout**](instance_config_layout.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]VdiPoolConfigNetworkInterfaces**](VdiPoolConfigNetworkInterfaces.md) |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**Backup** | Pointer to [**InstanceConfigBackup**](instance_config_backup.md) |  | [optional] 
**LoadBalancer** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HideLock** | Pointer to **bool** |  | [optional] 
**HasNetworks** | Pointer to **bool** |  | [optional] 
**DisplayNetworks** | Pointer to [**[]VdiPoolConfigDisplayNetworks**](VdiPoolConfigDisplayNetworks.md) |  | [optional] 
**Copies** | Pointer to **int64** |  | [optional] 
**ShowScale** | Pointer to **bool** |  | [optional] 
**HasPreview** | Pointer to **bool** |  | [optional] 
**VolumesDisplay** | Pointer to [**[]VdiPoolConfigVolumesDisplay**](VdiPoolConfigVolumesDisplay.md) |  | [optional] 

## Methods

### NewVdiPoolConfig

`func NewVdiPoolConfig() *VdiPoolConfig`

NewVdiPoolConfig instantiates a new VdiPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolConfigWithDefaults

`func NewVdiPoolConfigWithDefaults() *VdiPoolConfig`

NewVdiPoolConfigWithDefaults instantiates a new VdiPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *VdiPoolConfig) GetGroup() InstanceServicePlanAutoOptions`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VdiPoolConfig) GetGroupOk() (*InstanceServicePlanAutoOptions, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VdiPoolConfig) SetGroup(v InstanceServicePlanAutoOptions)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VdiPoolConfig) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *VdiPoolConfig) GetCloud() InlineResponse20040AppDeployInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VdiPoolConfig) GetCloudOk() (*InlineResponse20040AppDeployInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VdiPoolConfig) SetCloud(v InlineResponse20040AppDeployInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *VdiPoolConfig) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *VdiPoolConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VdiPoolConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VdiPoolConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VdiPoolConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstance

`func (o *VdiPoolConfig) GetInstance() VdiPoolConfigInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *VdiPoolConfig) GetInstanceOk() (*VdiPoolConfigInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *VdiPoolConfig) SetInstance(v VdiPoolConfigInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *VdiPoolConfig) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *VdiPoolConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VdiPoolConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VdiPoolConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VdiPoolConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnvironment

`func (o *VdiPoolConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VdiPoolConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VdiPoolConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VdiPoolConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfig

`func (o *VdiPoolConfig) GetConfig() VdiPoolConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VdiPoolConfig) GetConfigOk() (*VdiPoolConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VdiPoolConfig) SetConfig(v VdiPoolConfigConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VdiPoolConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *VdiPoolConfig) GetVolumes() []VdiPoolConfigVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VdiPoolConfig) GetVolumesOk() (*[]VdiPoolConfigVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VdiPoolConfig) SetVolumes(v []VdiPoolConfigVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VdiPoolConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetHostName

`func (o *VdiPoolConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VdiPoolConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VdiPoolConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *VdiPoolConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLayout

`func (o *VdiPoolConfig) GetLayout() InstanceConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *VdiPoolConfig) GetLayoutOk() (*InstanceConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *VdiPoolConfig) SetLayout(v InstanceConfigLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *VdiPoolConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStorageControllers

`func (o *VdiPoolConfig) GetStorageControllers() []VdiPoolConfigStorageControllers`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *VdiPoolConfig) GetStorageControllersOk() (*[]VdiPoolConfigStorageControllers, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *VdiPoolConfig) SetStorageControllers(v []VdiPoolConfigStorageControllers)`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *VdiPoolConfig) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### SetStorageControllersNil

`func (o *VdiPoolConfig) SetStorageControllersNil(b bool)`

 SetStorageControllersNil sets the value for StorageControllers to be an explicit nil

### UnsetStorageControllers
`func (o *VdiPoolConfig) UnsetStorageControllers()`

UnsetStorageControllers ensures that no value is present for StorageControllers, not even an explicit nil
### GetPlan

`func (o *VdiPoolConfig) GetPlan() InstanceConfigLayout`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *VdiPoolConfig) GetPlanOk() (*InstanceConfigLayout, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *VdiPoolConfig) SetPlan(v InstanceConfigLayout)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *VdiPoolConfig) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetVersion

`func (o *VdiPoolConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VdiPoolConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VdiPoolConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VdiPoolConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *VdiPoolConfig) GetNetworkInterfaces() []VdiPoolConfigNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *VdiPoolConfig) GetNetworkInterfacesOk() (*[]VdiPoolConfigNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *VdiPoolConfig) SetNetworkInterfaces(v []VdiPoolConfigNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *VdiPoolConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetExecutionId

`func (o *VdiPoolConfig) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *VdiPoolConfig) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *VdiPoolConfig) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *VdiPoolConfig) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetBackup

`func (o *VdiPoolConfig) GetBackup() InstanceConfigBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *VdiPoolConfig) GetBackupOk() (*InstanceConfigBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *VdiPoolConfig) SetBackup(v InstanceConfigBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *VdiPoolConfig) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *VdiPoolConfig) GetLoadBalancer() []map[string]interface{}`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *VdiPoolConfig) GetLoadBalancerOk() (*[]map[string]interface{}, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *VdiPoolConfig) SetLoadBalancer(v []map[string]interface{})`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *VdiPoolConfig) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *VdiPoolConfig) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *VdiPoolConfig) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
### GetHideLock

`func (o *VdiPoolConfig) GetHideLock() bool`

GetHideLock returns the HideLock field if non-nil, zero value otherwise.

### GetHideLockOk

`func (o *VdiPoolConfig) GetHideLockOk() (*bool, bool)`

GetHideLockOk returns a tuple with the HideLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideLock

`func (o *VdiPoolConfig) SetHideLock(v bool)`

SetHideLock sets HideLock field to given value.

### HasHideLock

`func (o *VdiPoolConfig) HasHideLock() bool`

HasHideLock returns a boolean if a field has been set.

### GetHasNetworks

`func (o *VdiPoolConfig) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *VdiPoolConfig) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *VdiPoolConfig) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *VdiPoolConfig) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.

### GetDisplayNetworks

`func (o *VdiPoolConfig) GetDisplayNetworks() []VdiPoolConfigDisplayNetworks`

GetDisplayNetworks returns the DisplayNetworks field if non-nil, zero value otherwise.

### GetDisplayNetworksOk

`func (o *VdiPoolConfig) GetDisplayNetworksOk() (*[]VdiPoolConfigDisplayNetworks, bool)`

GetDisplayNetworksOk returns a tuple with the DisplayNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNetworks

`func (o *VdiPoolConfig) SetDisplayNetworks(v []VdiPoolConfigDisplayNetworks)`

SetDisplayNetworks sets DisplayNetworks field to given value.

### HasDisplayNetworks

`func (o *VdiPoolConfig) HasDisplayNetworks() bool`

HasDisplayNetworks returns a boolean if a field has been set.

### GetCopies

`func (o *VdiPoolConfig) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *VdiPoolConfig) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *VdiPoolConfig) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *VdiPoolConfig) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetShowScale

`func (o *VdiPoolConfig) GetShowScale() bool`

GetShowScale returns the ShowScale field if non-nil, zero value otherwise.

### GetShowScaleOk

`func (o *VdiPoolConfig) GetShowScaleOk() (*bool, bool)`

GetShowScaleOk returns a tuple with the ShowScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScale

`func (o *VdiPoolConfig) SetShowScale(v bool)`

SetShowScale sets ShowScale field to given value.

### HasShowScale

`func (o *VdiPoolConfig) HasShowScale() bool`

HasShowScale returns a boolean if a field has been set.

### GetHasPreview

`func (o *VdiPoolConfig) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *VdiPoolConfig) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *VdiPoolConfig) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.

### HasHasPreview

`func (o *VdiPoolConfig) HasHasPreview() bool`

HasHasPreview returns a boolean if a field has been set.

### GetVolumesDisplay

`func (o *VdiPoolConfig) GetVolumesDisplay() []VdiPoolConfigVolumesDisplay`

GetVolumesDisplay returns the VolumesDisplay field if non-nil, zero value otherwise.

### GetVolumesDisplayOk

`func (o *VdiPoolConfig) GetVolumesDisplayOk() (*[]VdiPoolConfigVolumesDisplay, bool)`

GetVolumesDisplayOk returns a tuple with the VolumesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesDisplay

`func (o *VdiPoolConfig) SetVolumesDisplay(v []VdiPoolConfigVolumesDisplay)`

SetVolumesDisplay sets VolumesDisplay field to given value.

### HasVolumesDisplay

`func (o *VdiPoolConfig) HasVolumesDisplay() bool`

HasVolumesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


