# ServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolProviderType** | Pointer to **NullableString** |  | [optional] 
**IsVpcSelectable** | Pointer to **bool** |  | [optional] 
**SmbiosAssetTag** | Pointer to **NullableString** |  | [optional] 
**IsEC2** | Pointer to **bool** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**HostId** | Pointer to **NullableInt64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**NestedVirtualization** | Pointer to **NullableString** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**NoAgent** | Pointer to **bool** |  | [optional] 
**PowerScheduleType** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewServerConfig

`func NewServerConfig() *ServerConfig`

NewServerConfig instantiates a new ServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigWithDefaults

`func NewServerConfigWithDefaults() *ServerConfig`

NewServerConfigWithDefaults instantiates a new ServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolProviderType

`func (o *ServerConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *ServerConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *ServerConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *ServerConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### SetPoolProviderTypeNil

`func (o *ServerConfig) SetPoolProviderTypeNil(b bool)`

 SetPoolProviderTypeNil sets the value for PoolProviderType to be an explicit nil

### UnsetPoolProviderType
`func (o *ServerConfig) UnsetPoolProviderType()`

UnsetPoolProviderType ensures that no value is present for PoolProviderType, not even an explicit nil
### GetIsVpcSelectable

`func (o *ServerConfig) GetIsVpcSelectable() bool`

GetIsVpcSelectable returns the IsVpcSelectable field if non-nil, zero value otherwise.

### GetIsVpcSelectableOk

`func (o *ServerConfig) GetIsVpcSelectableOk() (*bool, bool)`

GetIsVpcSelectableOk returns a tuple with the IsVpcSelectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVpcSelectable

`func (o *ServerConfig) SetIsVpcSelectable(v bool)`

SetIsVpcSelectable sets IsVpcSelectable field to given value.

### HasIsVpcSelectable

`func (o *ServerConfig) HasIsVpcSelectable() bool`

HasIsVpcSelectable returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *ServerConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *ServerConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *ServerConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *ServerConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### SetSmbiosAssetTagNil

`func (o *ServerConfig) SetSmbiosAssetTagNil(b bool)`

 SetSmbiosAssetTagNil sets the value for SmbiosAssetTag to be an explicit nil

### UnsetSmbiosAssetTag
`func (o *ServerConfig) UnsetSmbiosAssetTag()`

UnsetSmbiosAssetTag ensures that no value is present for SmbiosAssetTag, not even an explicit nil
### GetIsEC2

`func (o *ServerConfig) GetIsEC2() bool`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *ServerConfig) GetIsEC2Ok() (*bool, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *ServerConfig) SetIsEC2(v bool)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *ServerConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ServerConfig) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ServerConfig) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ServerConfig) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ServerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *ServerConfig) GetHostId() int64`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ServerConfig) GetHostIdOk() (*int64, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ServerConfig) SetHostId(v int64)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ServerConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### SetHostIdNil

`func (o *ServerConfig) SetHostIdNil(b bool)`

 SetHostIdNil sets the value for HostId to be an explicit nil

### UnsetHostId
`func (o *ServerConfig) UnsetHostId()`

UnsetHostId ensures that no value is present for HostId, not even an explicit nil
### GetCreateUser

`func (o *ServerConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ServerConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ServerConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ServerConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *ServerConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *ServerConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *ServerConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *ServerConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### SetNestedVirtualizationNil

`func (o *ServerConfig) SetNestedVirtualizationNil(b bool)`

 SetNestedVirtualizationNil sets the value for NestedVirtualization to be an explicit nil

### UnsetNestedVirtualization
`func (o *ServerConfig) UnsetNestedVirtualization()`

UnsetNestedVirtualization ensures that no value is present for NestedVirtualization, not even an explicit nil
### GetVmwareFolderId

`func (o *ServerConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ServerConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ServerConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ServerConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetNoAgent

`func (o *ServerConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *ServerConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *ServerConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *ServerConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *ServerConfig) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *ServerConfig) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *ServerConfig) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *ServerConfig) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### SetPowerScheduleTypeNil

`func (o *ServerConfig) SetPowerScheduleTypeNil(b bool)`

 SetPowerScheduleTypeNil sets the value for PowerScheduleType to be an explicit nil

### UnsetPowerScheduleType
`func (o *ServerConfig) UnsetPowerScheduleType()`

UnsetPowerScheduleType ensures that no value is present for PowerScheduleType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


