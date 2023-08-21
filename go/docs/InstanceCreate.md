# InstanceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**InstanceCreateInstance**](instanceCreate_instance.md) |  | 
**ZoneId** | Pointer to **int64** | The Cloud ID to provision the instance onto. | [optional] 
**Evars** | Pointer to [**[]ApiServersIdMakeManagedServerTags**](ApiServersIdMakeManagedServerTags.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**Copies** | Pointer to **int64** | Number of copies to provision. | [optional] [default to 1]
**LayoutSize** | Pointer to **int64** | Apply a multiply factor of containers/vms within the instance. | [optional] [default to 1]
**ServicePlanOptions** | Pointer to **map[string]interface{}** | Map of custom options depending on selected service plan. | [optional] 
**SecurityGroups** | Pointer to **[]map[string]interface{}** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**Volumes** | Pointer to [**[]InstanceCreateVolume**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | [optional] 
**NetworkInterfaces** | Pointer to [**[]InstanceCreateNetwork**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Config** | [**AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject**](anyOf&lt;instancesConfigAzure,instancesConfigVMWare,instancesConfigGCP,instancesConfigAWS,object&gt;.md) |  | 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]ApiServersIdMakeManagedServerTags**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]ApiServersIdMakeManagedServerTags**](ApiServersIdMakeManagedServerTags.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]InstanceCreatePorts**](InstanceCreatePorts.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewInstanceCreate

`func NewInstanceCreate(instance InstanceCreateInstance, config AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject, ) *InstanceCreate`

NewInstanceCreate instantiates a new InstanceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCreateWithDefaults

`func NewInstanceCreateWithDefaults() *InstanceCreate`

NewInstanceCreateWithDefaults instantiates a new InstanceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *InstanceCreate) GetInstance() InstanceCreateInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceCreate) GetInstanceOk() (*InstanceCreateInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceCreate) SetInstance(v InstanceCreateInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *InstanceCreate) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *InstanceCreate) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *InstanceCreate) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *InstanceCreate) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceCreate) GetEvars() []ApiServersIdMakeManagedServerTags`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceCreate) GetEvarsOk() (*[]ApiServersIdMakeManagedServerTags, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceCreate) SetEvars(v []ApiServersIdMakeManagedServerTags)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceCreate) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetCopies

`func (o *InstanceCreate) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *InstanceCreate) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *InstanceCreate) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *InstanceCreate) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetLayoutSize

`func (o *InstanceCreate) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *InstanceCreate) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *InstanceCreate) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *InstanceCreate) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceCreate) GetServicePlanOptions() map[string]interface{}`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceCreate) GetServicePlanOptionsOk() (*map[string]interface{}, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceCreate) SetServicePlanOptions(v map[string]interface{})`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceCreate) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceCreate) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceCreate) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceCreate) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceCreate) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *InstanceCreate) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *InstanceCreate) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetVolumes

`func (o *InstanceCreate) GetVolumes() []InstanceCreateVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceCreate) GetVolumesOk() (*[]InstanceCreateVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceCreate) SetVolumes(v []InstanceCreateVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InstanceCreate) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InstanceCreate) GetNetworkInterfaces() []InstanceCreateNetwork`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceCreate) GetNetworkInterfacesOk() (*[]InstanceCreateNetwork, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceCreate) SetNetworkInterfaces(v []InstanceCreateNetwork)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceCreate) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetConfig

`func (o *InstanceCreate) GetConfig() AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceCreate) GetConfigOk() (*AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceCreate) SetConfig(v AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject)`

SetConfig sets Config field to given value.


### GetLabels

`func (o *InstanceCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceCreate) GetTags() []ApiServersIdMakeManagedServerTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceCreate) GetTagsOk() (*[]ApiServersIdMakeManagedServerTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceCreate) SetTags(v []ApiServersIdMakeManagedServerTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceCreate) GetMetadata() []ApiServersIdMakeManagedServerTags`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceCreate) GetMetadataOk() (*[]ApiServersIdMakeManagedServerTags, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceCreate) SetMetadata(v []ApiServersIdMakeManagedServerTags)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceCreate) GetPorts() []InstanceCreatePorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceCreate) GetPortsOk() (*[]InstanceCreatePorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceCreate) SetPorts(v []InstanceCreatePorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceCreate) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *InstanceCreate) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *InstanceCreate) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *InstanceCreate) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *InstanceCreate) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *InstanceCreate) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *InstanceCreate) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *InstanceCreate) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *InstanceCreate) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


