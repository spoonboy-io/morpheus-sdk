# ClusterLayoutUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster layout name | [optional] 
**Description** | Pointer to **NullableString** | Cluster layout description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**ComputeVersion** | Pointer to **string** | Version of the cluster layout | [optional] 
**Creatable** | Pointer to **bool** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**HasAutoScale** | Pointer to **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**InstallContainerRuntime** | Pointer to **bool** | Install Docker (container runtime). | [optional] [default to false]
**MemoryRequirement** | Pointer to **int64** | Memory requirement in bytes | [optional] 
**GroupType** | Pointer to [**ClusterLayoutCreateGroupType**](clusterLayoutCreate_groupType.md) |  | [optional] 
**ProvisionType** | Pointer to [**ApiServicePlansServicePlanProvisionType**](_api_service_plans_servicePlan_provisionType.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types | [optional] 
**TaskSets** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets | [optional] 
**EnvironmentVariables** | Pointer to [**[]ClusterLayoutCreateEnvironmentVariables**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables | [optional] 
**Masters** | Pointer to [**[]ClusterLayoutCreateMasters**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes | [optional] 
**Workers** | Pointer to [**[]ClusterLayoutCreateMasters**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes | [optional] 

## Methods

### NewClusterLayoutUpdate

`func NewClusterLayoutUpdate() *ClusterLayoutUpdate`

NewClusterLayoutUpdate instantiates a new ClusterLayoutUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutUpdateWithDefaults

`func NewClusterLayoutUpdateWithDefaults() *ClusterLayoutUpdate`

NewClusterLayoutUpdateWithDefaults instantiates a new ClusterLayoutUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterLayoutUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterLayoutUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterLayoutUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterLayoutUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterLayoutUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterLayoutUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterLayoutUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *ClusterLayoutUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterLayoutUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterLayoutUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterLayoutUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClusterLayoutUpdate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClusterLayoutUpdate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetComputeVersion

`func (o *ClusterLayoutUpdate) GetComputeVersion() string`

GetComputeVersion returns the ComputeVersion field if non-nil, zero value otherwise.

### GetComputeVersionOk

`func (o *ClusterLayoutUpdate) GetComputeVersionOk() (*string, bool)`

GetComputeVersionOk returns a tuple with the ComputeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVersion

`func (o *ClusterLayoutUpdate) SetComputeVersion(v string)`

SetComputeVersion sets ComputeVersion field to given value.

### HasComputeVersion

`func (o *ClusterLayoutUpdate) HasComputeVersion() bool`

HasComputeVersion returns a boolean if a field has been set.

### GetCreatable

`func (o *ClusterLayoutUpdate) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ClusterLayoutUpdate) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ClusterLayoutUpdate) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ClusterLayoutUpdate) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *ClusterLayoutUpdate) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *ClusterLayoutUpdate) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *ClusterLayoutUpdate) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *ClusterLayoutUpdate) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ClusterLayoutUpdate) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ClusterLayoutUpdate) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ClusterLayoutUpdate) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ClusterLayoutUpdate) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *ClusterLayoutUpdate) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *ClusterLayoutUpdate) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *ClusterLayoutUpdate) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *ClusterLayoutUpdate) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetGroupType

`func (o *ClusterLayoutUpdate) GetGroupType() ClusterLayoutCreateGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ClusterLayoutUpdate) GetGroupTypeOk() (*ClusterLayoutCreateGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ClusterLayoutUpdate) SetGroupType(v ClusterLayoutCreateGroupType)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *ClusterLayoutUpdate) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetProvisionType

`func (o *ClusterLayoutUpdate) GetProvisionType() ApiServicePlansServicePlanProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ClusterLayoutUpdate) GetProvisionTypeOk() (*ApiServicePlansServicePlanProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ClusterLayoutUpdate) SetProvisionType(v ApiServicePlansServicePlanProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ClusterLayoutUpdate) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ClusterLayoutUpdate) GetOptionTypes() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ClusterLayoutUpdate) GetOptionTypesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ClusterLayoutUpdate) SetOptionTypes(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ClusterLayoutUpdate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTaskSets

`func (o *ClusterLayoutUpdate) GetTaskSets() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ClusterLayoutUpdate) GetTaskSetsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ClusterLayoutUpdate) SetTaskSets(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ClusterLayoutUpdate) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ClusterLayoutUpdate) GetEnvironmentVariables() []ClusterLayoutCreateEnvironmentVariables`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ClusterLayoutUpdate) GetEnvironmentVariablesOk() (*[]ClusterLayoutCreateEnvironmentVariables, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ClusterLayoutUpdate) SetEnvironmentVariables(v []ClusterLayoutCreateEnvironmentVariables)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ClusterLayoutUpdate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetMasters

`func (o *ClusterLayoutUpdate) GetMasters() []ClusterLayoutCreateMasters`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *ClusterLayoutUpdate) GetMastersOk() (*[]ClusterLayoutCreateMasters, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *ClusterLayoutUpdate) SetMasters(v []ClusterLayoutCreateMasters)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *ClusterLayoutUpdate) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetWorkers

`func (o *ClusterLayoutUpdate) GetWorkers() []ClusterLayoutCreateMasters`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *ClusterLayoutUpdate) GetWorkersOk() (*[]ClusterLayoutCreateMasters, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *ClusterLayoutUpdate) SetWorkers(v []ClusterLayoutCreateMasters)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *ClusterLayoutUpdate) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


