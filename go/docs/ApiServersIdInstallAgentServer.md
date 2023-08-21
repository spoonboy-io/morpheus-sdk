# ApiServersIdInstallAgentServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | Pointer to **string** | SSH username to use when provisioning | [optional] 
**SshPassword** | Pointer to **string** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | Pointer to [**ApiServersIdInstallAgentServerServerOs**](_api_servers__id__install_agent_server_serverOs.md) |  | [optional] 

## Methods

### NewApiServersIdInstallAgentServer

`func NewApiServersIdInstallAgentServer() *ApiServersIdInstallAgentServer`

NewApiServersIdInstallAgentServer instantiates a new ApiServersIdInstallAgentServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServersIdInstallAgentServerWithDefaults

`func NewApiServersIdInstallAgentServerWithDefaults() *ApiServersIdInstallAgentServer`

NewApiServersIdInstallAgentServerWithDefaults instantiates a new ApiServersIdInstallAgentServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUsername

`func (o *ApiServersIdInstallAgentServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ApiServersIdInstallAgentServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ApiServersIdInstallAgentServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ApiServersIdInstallAgentServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ApiServersIdInstallAgentServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ApiServersIdInstallAgentServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ApiServersIdInstallAgentServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ApiServersIdInstallAgentServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetServerOs

`func (o *ApiServersIdInstallAgentServer) GetServerOs() ApiServersIdInstallAgentServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ApiServersIdInstallAgentServer) GetServerOsOk() (*ApiServersIdInstallAgentServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ApiServersIdInstallAgentServer) SetServerOs(v ApiServersIdInstallAgentServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ApiServersIdInstallAgentServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


