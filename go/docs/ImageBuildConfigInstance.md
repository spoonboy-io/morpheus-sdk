# ImageBuildConfigInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | Pointer to [**ImageBuildsConfigPlan**](imageBuilds_config_plan.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserGroup** | Pointer to [**ZoneVcenterConfigNetworkServer**](zoneVcenterConfig_networkServer.md) |  | [optional] 

## Methods

### NewImageBuildConfigInstance

`func NewImageBuildConfigInstance() *ImageBuildConfigInstance`

NewImageBuildConfigInstance instantiates a new ImageBuildConfigInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildConfigInstanceWithDefaults

`func NewImageBuildConfigInstanceWithDefaults() *ImageBuildConfigInstance`

NewImageBuildConfigInstanceWithDefaults instantiates a new ImageBuildConfigInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *ImageBuildConfigInstance) GetLayout() ImageBuildsConfigPlan`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ImageBuildConfigInstance) GetLayoutOk() (*ImageBuildsConfigPlan, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ImageBuildConfigInstance) SetLayout(v ImageBuildsConfigPlan)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ImageBuildConfigInstance) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetType

`func (o *ImageBuildConfigInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageBuildConfigInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageBuildConfigInstance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageBuildConfigInstance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserGroup

`func (o *ImageBuildConfigInstance) GetUserGroup() ZoneVcenterConfigNetworkServer`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ImageBuildConfigInstance) GetUserGroupOk() (*ZoneVcenterConfigNetworkServer, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ImageBuildConfigInstance) SetUserGroup(v ZoneVcenterConfigNetworkServer)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ImageBuildConfigInstance) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


