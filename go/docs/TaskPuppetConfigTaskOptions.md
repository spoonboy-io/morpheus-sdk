# TaskPuppetConfigTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**PuppetEnvironment** | Pointer to **NullableString** |  | [optional] 
**PuppetNodeName** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskPuppetConfigTaskOptions

`func NewTaskPuppetConfigTaskOptions() *TaskPuppetConfigTaskOptions`

NewTaskPuppetConfigTaskOptions instantiates a new TaskPuppetConfigTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPuppetConfigTaskOptionsWithDefaults

`func NewTaskPuppetConfigTaskOptionsWithDefaults() *TaskPuppetConfigTaskOptions`

NewTaskPuppetConfigTaskOptionsWithDefaults instantiates a new TaskPuppetConfigTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *TaskPuppetConfigTaskOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskPuppetConfigTaskOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskPuppetConfigTaskOptions) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskPuppetConfigTaskOptions) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskPuppetConfigTaskOptions) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskPuppetConfigTaskOptions) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHost

`func (o *TaskPuppetConfigTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskPuppetConfigTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskPuppetConfigTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskPuppetConfigTaskOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskPuppetConfigTaskOptions) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskPuppetConfigTaskOptions) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskPuppetConfigTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskPuppetConfigTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskPuppetConfigTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskPuppetConfigTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskPuppetConfigTaskOptions) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskPuppetConfigTaskOptions) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPuppetEnvironment

`func (o *TaskPuppetConfigTaskOptions) GetPuppetEnvironment() string`

GetPuppetEnvironment returns the PuppetEnvironment field if non-nil, zero value otherwise.

### GetPuppetEnvironmentOk

`func (o *TaskPuppetConfigTaskOptions) GetPuppetEnvironmentOk() (*string, bool)`

GetPuppetEnvironmentOk returns a tuple with the PuppetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetEnvironment

`func (o *TaskPuppetConfigTaskOptions) SetPuppetEnvironment(v string)`

SetPuppetEnvironment sets PuppetEnvironment field to given value.

### HasPuppetEnvironment

`func (o *TaskPuppetConfigTaskOptions) HasPuppetEnvironment() bool`

HasPuppetEnvironment returns a boolean if a field has been set.

### SetPuppetEnvironmentNil

`func (o *TaskPuppetConfigTaskOptions) SetPuppetEnvironmentNil(b bool)`

 SetPuppetEnvironmentNil sets the value for PuppetEnvironment to be an explicit nil

### UnsetPuppetEnvironment
`func (o *TaskPuppetConfigTaskOptions) UnsetPuppetEnvironment()`

UnsetPuppetEnvironment ensures that no value is present for PuppetEnvironment, not even an explicit nil
### GetPuppetNodeName

`func (o *TaskPuppetConfigTaskOptions) GetPuppetNodeName() string`

GetPuppetNodeName returns the PuppetNodeName field if non-nil, zero value otherwise.

### GetPuppetNodeNameOk

`func (o *TaskPuppetConfigTaskOptions) GetPuppetNodeNameOk() (*string, bool)`

GetPuppetNodeNameOk returns a tuple with the PuppetNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuppetNodeName

`func (o *TaskPuppetConfigTaskOptions) SetPuppetNodeName(v string)`

SetPuppetNodeName sets PuppetNodeName field to given value.

### HasPuppetNodeName

`func (o *TaskPuppetConfigTaskOptions) HasPuppetNodeName() bool`

HasPuppetNodeName returns a boolean if a field has been set.

### SetPuppetNodeNameNil

`func (o *TaskPuppetConfigTaskOptions) SetPuppetNodeNameNil(b bool)`

 SetPuppetNodeNameNil sets the value for PuppetNodeName to be an explicit nil

### UnsetPuppetNodeName
`func (o *TaskPuppetConfigTaskOptions) UnsetPuppetNodeName()`

UnsetPuppetNodeName ensures that no value is present for PuppetNodeName, not even an explicit nil
### GetSshKey

`func (o *TaskPuppetConfigTaskOptions) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskPuppetConfigTaskOptions) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskPuppetConfigTaskOptions) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskPuppetConfigTaskOptions) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskPuppetConfigTaskOptions) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskPuppetConfigTaskOptions) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskPuppetConfigTaskOptions) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskPuppetConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskPuppetConfigTaskOptions) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskPuppetConfigTaskOptions) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskPuppetConfigTaskOptions) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskPuppetConfigTaskOptions) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskPuppetConfigTaskOptions) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskPuppetConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskPuppetConfigTaskOptions) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskPuppetConfigTaskOptions) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskPuppetConfigTaskOptions) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskPuppetConfigTaskOptions) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil
### GetPassword

`func (o *TaskPuppetConfigTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskPuppetConfigTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskPuppetConfigTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskPuppetConfigTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskPuppetConfigTaskOptions) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskPuppetConfigTaskOptions) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskPuppetConfigTaskOptions) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskPuppetConfigTaskOptions) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskPuppetConfigTaskOptions) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskPuppetConfigTaskOptions) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskPuppetConfigTaskOptions) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskPuppetConfigTaskOptions) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


