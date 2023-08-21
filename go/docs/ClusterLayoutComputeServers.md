# ClusterLayoutComputeServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PriorityOrder** | Pointer to **int64** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**MinNodeCount** | Pointer to **int64** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableString** |  | [optional] 
**DynamicCount** | Pointer to **bool** |  | [optional] 
**InstallContainerRuntime** | Pointer to **bool** |  | [optional] 
**InstallStorageRuntime** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**ContainerType** | Pointer to [**ClusterLayoutContainerType**](clusterLayout_containerType.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ClusterLayoutComputeServerType**](clusterLayout_computeServerType.md) |  | [optional] 
**ProvisionService** | Pointer to **NullableString** |  | [optional] 
**PlanCategory** | Pointer to **NullableString** |  | [optional] 
**NamePrefix** | Pointer to **NullableString** |  | [optional] 
**NameSuffix** | Pointer to **NullableString** |  | [optional] 
**ForceNameIndex** | Pointer to **bool** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterLayoutComputeServers

`func NewClusterLayoutComputeServers() *ClusterLayoutComputeServers`

NewClusterLayoutComputeServers instantiates a new ClusterLayoutComputeServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutComputeServersWithDefaults

`func NewClusterLayoutComputeServersWithDefaults() *ClusterLayoutComputeServers`

NewClusterLayoutComputeServersWithDefaults instantiates a new ClusterLayoutComputeServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterLayoutComputeServers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterLayoutComputeServers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterLayoutComputeServers) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterLayoutComputeServers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriorityOrder

`func (o *ClusterLayoutComputeServers) GetPriorityOrder() int64`

GetPriorityOrder returns the PriorityOrder field if non-nil, zero value otherwise.

### GetPriorityOrderOk

`func (o *ClusterLayoutComputeServers) GetPriorityOrderOk() (*int64, bool)`

GetPriorityOrderOk returns a tuple with the PriorityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityOrder

`func (o *ClusterLayoutComputeServers) SetPriorityOrder(v int64)`

SetPriorityOrder sets PriorityOrder field to given value.

### HasPriorityOrder

`func (o *ClusterLayoutComputeServers) HasPriorityOrder() bool`

HasPriorityOrder returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterLayoutComputeServers) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterLayoutComputeServers) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterLayoutComputeServers) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterLayoutComputeServers) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeType

`func (o *ClusterLayoutComputeServers) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ClusterLayoutComputeServers) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ClusterLayoutComputeServers) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ClusterLayoutComputeServers) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetMinNodeCount

`func (o *ClusterLayoutComputeServers) GetMinNodeCount() int64`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *ClusterLayoutComputeServers) GetMinNodeCountOk() (*int64, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *ClusterLayoutComputeServers) SetMinNodeCount(v int64)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *ClusterLayoutComputeServers) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### GetMaxNodeCount

`func (o *ClusterLayoutComputeServers) GetMaxNodeCount() string`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *ClusterLayoutComputeServers) GetMaxNodeCountOk() (*string, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *ClusterLayoutComputeServers) SetMaxNodeCount(v string)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *ClusterLayoutComputeServers) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *ClusterLayoutComputeServers) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *ClusterLayoutComputeServers) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetDynamicCount

`func (o *ClusterLayoutComputeServers) GetDynamicCount() bool`

GetDynamicCount returns the DynamicCount field if non-nil, zero value otherwise.

### GetDynamicCountOk

`func (o *ClusterLayoutComputeServers) GetDynamicCountOk() (*bool, bool)`

GetDynamicCountOk returns a tuple with the DynamicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCount

`func (o *ClusterLayoutComputeServers) SetDynamicCount(v bool)`

SetDynamicCount sets DynamicCount field to given value.

### HasDynamicCount

`func (o *ClusterLayoutComputeServers) HasDynamicCount() bool`

HasDynamicCount returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ClusterLayoutComputeServers) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ClusterLayoutComputeServers) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ClusterLayoutComputeServers) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ClusterLayoutComputeServers) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetInstallStorageRuntime

`func (o *ClusterLayoutComputeServers) GetInstallStorageRuntime() bool`

GetInstallStorageRuntime returns the InstallStorageRuntime field if non-nil, zero value otherwise.

### GetInstallStorageRuntimeOk

`func (o *ClusterLayoutComputeServers) GetInstallStorageRuntimeOk() (*bool, bool)`

GetInstallStorageRuntimeOk returns a tuple with the InstallStorageRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStorageRuntime

`func (o *ClusterLayoutComputeServers) SetInstallStorageRuntime(v bool)`

SetInstallStorageRuntime sets InstallStorageRuntime field to given value.

### HasInstallStorageRuntime

`func (o *ClusterLayoutComputeServers) HasInstallStorageRuntime() bool`

HasInstallStorageRuntime returns a boolean if a field has been set.

### GetName

`func (o *ClusterLayoutComputeServers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutComputeServers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutComputeServers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutComputeServers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ClusterLayoutComputeServers) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterLayoutComputeServers) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterLayoutComputeServers) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterLayoutComputeServers) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *ClusterLayoutComputeServers) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ClusterLayoutComputeServers) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ClusterLayoutComputeServers) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ClusterLayoutComputeServers) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterLayoutComputeServers) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterLayoutComputeServers) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterLayoutComputeServers) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterLayoutComputeServers) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ClusterLayoutComputeServers) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ClusterLayoutComputeServers) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerType

`func (o *ClusterLayoutComputeServers) GetContainerType() ClusterLayoutContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ClusterLayoutComputeServers) GetContainerTypeOk() (*ClusterLayoutContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ClusterLayoutComputeServers) SetContainerType(v ClusterLayoutContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ClusterLayoutComputeServers) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClusterLayoutComputeServers) GetComputeServerType() ClusterLayoutComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClusterLayoutComputeServers) GetComputeServerTypeOk() (*ClusterLayoutComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClusterLayoutComputeServers) SetComputeServerType(v ClusterLayoutComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClusterLayoutComputeServers) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetProvisionService

`func (o *ClusterLayoutComputeServers) GetProvisionService() string`

GetProvisionService returns the ProvisionService field if non-nil, zero value otherwise.

### GetProvisionServiceOk

`func (o *ClusterLayoutComputeServers) GetProvisionServiceOk() (*string, bool)`

GetProvisionServiceOk returns a tuple with the ProvisionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionService

`func (o *ClusterLayoutComputeServers) SetProvisionService(v string)`

SetProvisionService sets ProvisionService field to given value.

### HasProvisionService

`func (o *ClusterLayoutComputeServers) HasProvisionService() bool`

HasProvisionService returns a boolean if a field has been set.

### SetProvisionServiceNil

`func (o *ClusterLayoutComputeServers) SetProvisionServiceNil(b bool)`

 SetProvisionServiceNil sets the value for ProvisionService to be an explicit nil

### UnsetProvisionService
`func (o *ClusterLayoutComputeServers) UnsetProvisionService()`

UnsetProvisionService ensures that no value is present for ProvisionService, not even an explicit nil
### GetPlanCategory

`func (o *ClusterLayoutComputeServers) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *ClusterLayoutComputeServers) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *ClusterLayoutComputeServers) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *ClusterLayoutComputeServers) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### SetPlanCategoryNil

`func (o *ClusterLayoutComputeServers) SetPlanCategoryNil(b bool)`

 SetPlanCategoryNil sets the value for PlanCategory to be an explicit nil

### UnsetPlanCategory
`func (o *ClusterLayoutComputeServers) UnsetPlanCategory()`

UnsetPlanCategory ensures that no value is present for PlanCategory, not even an explicit nil
### GetNamePrefix

`func (o *ClusterLayoutComputeServers) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *ClusterLayoutComputeServers) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *ClusterLayoutComputeServers) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *ClusterLayoutComputeServers) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### SetNamePrefixNil

`func (o *ClusterLayoutComputeServers) SetNamePrefixNil(b bool)`

 SetNamePrefixNil sets the value for NamePrefix to be an explicit nil

### UnsetNamePrefix
`func (o *ClusterLayoutComputeServers) UnsetNamePrefix()`

UnsetNamePrefix ensures that no value is present for NamePrefix, not even an explicit nil
### GetNameSuffix

`func (o *ClusterLayoutComputeServers) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *ClusterLayoutComputeServers) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *ClusterLayoutComputeServers) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *ClusterLayoutComputeServers) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### SetNameSuffixNil

`func (o *ClusterLayoutComputeServers) SetNameSuffixNil(b bool)`

 SetNameSuffixNil sets the value for NameSuffix to be an explicit nil

### UnsetNameSuffix
`func (o *ClusterLayoutComputeServers) UnsetNameSuffix()`

UnsetNameSuffix ensures that no value is present for NameSuffix, not even an explicit nil
### GetForceNameIndex

`func (o *ClusterLayoutComputeServers) GetForceNameIndex() bool`

GetForceNameIndex returns the ForceNameIndex field if non-nil, zero value otherwise.

### GetForceNameIndexOk

`func (o *ClusterLayoutComputeServers) GetForceNameIndexOk() (*bool, bool)`

GetForceNameIndexOk returns a tuple with the ForceNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNameIndex

`func (o *ClusterLayoutComputeServers) SetForceNameIndex(v bool)`

SetForceNameIndex sets ForceNameIndex field to given value.

### HasForceNameIndex

`func (o *ClusterLayoutComputeServers) HasForceNameIndex() bool`

HasForceNameIndex returns a boolean if a field has been set.

### GetLoadBalance

`func (o *ClusterLayoutComputeServers) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *ClusterLayoutComputeServers) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *ClusterLayoutComputeServers) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *ClusterLayoutComputeServers) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


