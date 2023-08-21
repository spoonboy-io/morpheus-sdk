# ApiServersIdMakeManagedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | Pointer to **string** | SSH username to use when provisioning | [optional] 
**SshPassword** | Pointer to **string** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | Pointer to [**ApiServersIdInstallAgentServerServerOs**](_api_servers__id__install_agent_server_serverOs.md) |  | [optional] 
**Plan** | Pointer to [**ApiServersIdMakeManagedServerPlan**](_api_servers__id__make_managed_server_plan.md) |  | [optional] 
**Account** | Pointer to [**ApiServersIdMakeManagedServerAccount**](_api_servers__id__make_managed_server_account.md) |  | [optional] 
**ProvisionSiteId** | Pointer to **int64** | Specific group to assign the server | [optional] 
**Tags** | Pointer to [**[]ApiServersIdMakeManagedServerTags**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**Config** | Pointer to [**ApiServersIdMakeManagedServerConfig**](_api_servers__id__make_managed_server_config.md) |  | [optional] 

## Methods

### NewApiServersIdMakeManagedServer

`func NewApiServersIdMakeManagedServer() *ApiServersIdMakeManagedServer`

NewApiServersIdMakeManagedServer instantiates a new ApiServersIdMakeManagedServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServersIdMakeManagedServerWithDefaults

`func NewApiServersIdMakeManagedServerWithDefaults() *ApiServersIdMakeManagedServer`

NewApiServersIdMakeManagedServerWithDefaults instantiates a new ApiServersIdMakeManagedServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUsername

`func (o *ApiServersIdMakeManagedServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ApiServersIdMakeManagedServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ApiServersIdMakeManagedServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ApiServersIdMakeManagedServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ApiServersIdMakeManagedServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ApiServersIdMakeManagedServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ApiServersIdMakeManagedServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ApiServersIdMakeManagedServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetServerOs

`func (o *ApiServersIdMakeManagedServer) GetServerOs() ApiServersIdInstallAgentServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ApiServersIdMakeManagedServer) GetServerOsOk() (*ApiServersIdInstallAgentServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ApiServersIdMakeManagedServer) SetServerOs(v ApiServersIdInstallAgentServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ApiServersIdMakeManagedServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetPlan

`func (o *ApiServersIdMakeManagedServer) GetPlan() ApiServersIdMakeManagedServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ApiServersIdMakeManagedServer) GetPlanOk() (*ApiServersIdMakeManagedServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ApiServersIdMakeManagedServer) SetPlan(v ApiServersIdMakeManagedServerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ApiServersIdMakeManagedServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetAccount

`func (o *ApiServersIdMakeManagedServer) GetAccount() ApiServersIdMakeManagedServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApiServersIdMakeManagedServer) GetAccountOk() (*ApiServersIdMakeManagedServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApiServersIdMakeManagedServer) SetAccount(v ApiServersIdMakeManagedServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApiServersIdMakeManagedServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProvisionSiteId

`func (o *ApiServersIdMakeManagedServer) GetProvisionSiteId() int64`

GetProvisionSiteId returns the ProvisionSiteId field if non-nil, zero value otherwise.

### GetProvisionSiteIdOk

`func (o *ApiServersIdMakeManagedServer) GetProvisionSiteIdOk() (*int64, bool)`

GetProvisionSiteIdOk returns a tuple with the ProvisionSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionSiteId

`func (o *ApiServersIdMakeManagedServer) SetProvisionSiteId(v int64)`

SetProvisionSiteId sets ProvisionSiteId field to given value.

### HasProvisionSiteId

`func (o *ApiServersIdMakeManagedServer) HasProvisionSiteId() bool`

HasProvisionSiteId returns a boolean if a field has been set.

### GetTags

`func (o *ApiServersIdMakeManagedServer) GetTags() []ApiServersIdMakeManagedServerTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiServersIdMakeManagedServer) GetTagsOk() (*[]ApiServersIdMakeManagedServerTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiServersIdMakeManagedServer) SetTags(v []ApiServersIdMakeManagedServerTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiServersIdMakeManagedServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfig

`func (o *ApiServersIdMakeManagedServer) GetConfig() ApiServersIdMakeManagedServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiServersIdMakeManagedServer) GetConfigOk() (*ApiServersIdMakeManagedServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiServersIdMakeManagedServer) SetConfig(v ApiServersIdMakeManagedServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiServersIdMakeManagedServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


