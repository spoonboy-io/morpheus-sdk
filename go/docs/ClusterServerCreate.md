# ClusterServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  | 
**ServerType** | Pointer to [**ClusterServerCreateServerType**](clusterServerCreate_serverType.md) |  | [optional] 
**Name** | **string** | Name to be used for host(s) created in the cluster | 
**Plan** | [**ClusterServerCreatePlan**](clusterServerCreate_plan.md) |  | 
**Volumes** | Pointer to [**[]ClusterServerCreateVolumes**](ClusterServerCreateVolumes.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**NetworkInterfaces** | Pointer to [**[]ClusterServerCreateNetworkInterfaces**](ClusterServerCreateNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**SecurityGroups** | Pointer to **[]string** | Key for security group configuration. | [optional] 
**Visibility** | Pointer to **string** | Visibility for server host | [optional] [default to "private"]
**UserGroup** | Pointer to [**ClusterServerCreateUserGroup**](clusterServerCreate_userGroup.md) |  | [optional] 
**NetworkDomain** | Pointer to **NullableString** | Network domain | [optional] 
**Hostname** | Pointer to **NullableString** | Hostname for server host | [optional] 
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**Tags** | Pointer to [**[]ApiServersIdMakeManagedServerTags**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 

## Methods

### NewClusterServerCreate

`func NewClusterServerCreate(config map[string]interface{}, name string, plan ClusterServerCreatePlan, ) *ClusterServerCreate`

NewClusterServerCreate instantiates a new ClusterServerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateWithDefaults

`func NewClusterServerCreateWithDefaults() *ClusterServerCreate`

NewClusterServerCreateWithDefaults instantiates a new ClusterServerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ClusterServerCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterServerCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterServerCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetServerType

`func (o *ClusterServerCreate) GetServerType() ClusterServerCreateServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ClusterServerCreate) GetServerTypeOk() (*ClusterServerCreateServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ClusterServerCreate) SetServerType(v ClusterServerCreateServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ClusterServerCreate) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *ClusterServerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterServerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterServerCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPlan

`func (o *ClusterServerCreate) GetPlan() ClusterServerCreatePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterServerCreate) GetPlanOk() (*ClusterServerCreatePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterServerCreate) SetPlan(v ClusterServerCreatePlan)`

SetPlan sets Plan field to given value.


### GetVolumes

`func (o *ClusterServerCreate) GetVolumes() []ClusterServerCreateVolumes`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ClusterServerCreate) GetVolumesOk() (*[]ClusterServerCreateVolumes, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ClusterServerCreate) SetVolumes(v []ClusterServerCreateVolumes)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ClusterServerCreate) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ClusterServerCreate) GetNetworkInterfaces() []ClusterServerCreateNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ClusterServerCreate) GetNetworkInterfacesOk() (*[]ClusterServerCreateNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ClusterServerCreate) SetNetworkInterfaces(v []ClusterServerCreateNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ClusterServerCreate) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ClusterServerCreate) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ClusterServerCreate) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ClusterServerCreate) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ClusterServerCreate) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterServerCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterServerCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterServerCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterServerCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetUserGroup

`func (o *ClusterServerCreate) GetUserGroup() ClusterServerCreateUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ClusterServerCreate) GetUserGroupOk() (*ClusterServerCreateUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ClusterServerCreate) SetUserGroup(v ClusterServerCreateUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ClusterServerCreate) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ClusterServerCreate) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ClusterServerCreate) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ClusterServerCreate) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ClusterServerCreate) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *ClusterServerCreate) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *ClusterServerCreate) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetHostname

`func (o *ClusterServerCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClusterServerCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClusterServerCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClusterServerCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *ClusterServerCreate) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *ClusterServerCreate) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetNodeCount

`func (o *ClusterServerCreate) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreate) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreate) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreate) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTags

`func (o *ClusterServerCreate) GetTags() []ApiServersIdMakeManagedServerTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterServerCreate) GetTagsOk() (*[]ApiServersIdMakeManagedServerTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterServerCreate) SetTags(v []ApiServersIdMakeManagedServerTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterServerCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterServerCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterServerCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterServerCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterServerCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


