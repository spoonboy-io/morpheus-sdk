# TaskPythonConfigTaskOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PythonArgs** | Pointer to **NullableString** |  | [optional] 
**PythonBinary** | Pointer to **NullableString** |  | [optional] 
**PythonAdditionalPackages** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitId** | Pointer to **NullableString** |  | [optional] 
**LocalScriptGitRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskPythonConfigTaskOptions

`func NewTaskPythonConfigTaskOptions() *TaskPythonConfigTaskOptions`

NewTaskPythonConfigTaskOptions instantiates a new TaskPythonConfigTaskOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPythonConfigTaskOptionsWithDefaults

`func NewTaskPythonConfigTaskOptionsWithDefaults() *TaskPythonConfigTaskOptions`

NewTaskPythonConfigTaskOptionsWithDefaults instantiates a new TaskPythonConfigTaskOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPythonArgs

`func (o *TaskPythonConfigTaskOptions) GetPythonArgs() string`

GetPythonArgs returns the PythonArgs field if non-nil, zero value otherwise.

### GetPythonArgsOk

`func (o *TaskPythonConfigTaskOptions) GetPythonArgsOk() (*string, bool)`

GetPythonArgsOk returns a tuple with the PythonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonArgs

`func (o *TaskPythonConfigTaskOptions) SetPythonArgs(v string)`

SetPythonArgs sets PythonArgs field to given value.

### HasPythonArgs

`func (o *TaskPythonConfigTaskOptions) HasPythonArgs() bool`

HasPythonArgs returns a boolean if a field has been set.

### SetPythonArgsNil

`func (o *TaskPythonConfigTaskOptions) SetPythonArgsNil(b bool)`

 SetPythonArgsNil sets the value for PythonArgs to be an explicit nil

### UnsetPythonArgs
`func (o *TaskPythonConfigTaskOptions) UnsetPythonArgs()`

UnsetPythonArgs ensures that no value is present for PythonArgs, not even an explicit nil
### GetPythonBinary

`func (o *TaskPythonConfigTaskOptions) GetPythonBinary() string`

GetPythonBinary returns the PythonBinary field if non-nil, zero value otherwise.

### GetPythonBinaryOk

`func (o *TaskPythonConfigTaskOptions) GetPythonBinaryOk() (*string, bool)`

GetPythonBinaryOk returns a tuple with the PythonBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonBinary

`func (o *TaskPythonConfigTaskOptions) SetPythonBinary(v string)`

SetPythonBinary sets PythonBinary field to given value.

### HasPythonBinary

`func (o *TaskPythonConfigTaskOptions) HasPythonBinary() bool`

HasPythonBinary returns a boolean if a field has been set.

### SetPythonBinaryNil

`func (o *TaskPythonConfigTaskOptions) SetPythonBinaryNil(b bool)`

 SetPythonBinaryNil sets the value for PythonBinary to be an explicit nil

### UnsetPythonBinary
`func (o *TaskPythonConfigTaskOptions) UnsetPythonBinary()`

UnsetPythonBinary ensures that no value is present for PythonBinary, not even an explicit nil
### GetPythonAdditionalPackages

`func (o *TaskPythonConfigTaskOptions) GetPythonAdditionalPackages() string`

GetPythonAdditionalPackages returns the PythonAdditionalPackages field if non-nil, zero value otherwise.

### GetPythonAdditionalPackagesOk

`func (o *TaskPythonConfigTaskOptions) GetPythonAdditionalPackagesOk() (*string, bool)`

GetPythonAdditionalPackagesOk returns a tuple with the PythonAdditionalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonAdditionalPackages

`func (o *TaskPythonConfigTaskOptions) SetPythonAdditionalPackages(v string)`

SetPythonAdditionalPackages sets PythonAdditionalPackages field to given value.

### HasPythonAdditionalPackages

`func (o *TaskPythonConfigTaskOptions) HasPythonAdditionalPackages() bool`

HasPythonAdditionalPackages returns a boolean if a field has been set.

### SetPythonAdditionalPackagesNil

`func (o *TaskPythonConfigTaskOptions) SetPythonAdditionalPackagesNil(b bool)`

 SetPythonAdditionalPackagesNil sets the value for PythonAdditionalPackages to be an explicit nil

### UnsetPythonAdditionalPackages
`func (o *TaskPythonConfigTaskOptions) UnsetPythonAdditionalPackages()`

UnsetPythonAdditionalPackages ensures that no value is present for PythonAdditionalPackages, not even an explicit nil
### GetPort

`func (o *TaskPythonConfigTaskOptions) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TaskPythonConfigTaskOptions) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TaskPythonConfigTaskOptions) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TaskPythonConfigTaskOptions) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *TaskPythonConfigTaskOptions) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *TaskPythonConfigTaskOptions) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHost

`func (o *TaskPythonConfigTaskOptions) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TaskPythonConfigTaskOptions) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TaskPythonConfigTaskOptions) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TaskPythonConfigTaskOptions) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *TaskPythonConfigTaskOptions) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *TaskPythonConfigTaskOptions) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *TaskPythonConfigTaskOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TaskPythonConfigTaskOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TaskPythonConfigTaskOptions) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TaskPythonConfigTaskOptions) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TaskPythonConfigTaskOptions) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TaskPythonConfigTaskOptions) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSshKey

`func (o *TaskPythonConfigTaskOptions) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *TaskPythonConfigTaskOptions) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *TaskPythonConfigTaskOptions) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *TaskPythonConfigTaskOptions) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *TaskPythonConfigTaskOptions) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *TaskPythonConfigTaskOptions) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetPassword

`func (o *TaskPythonConfigTaskOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TaskPythonConfigTaskOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TaskPythonConfigTaskOptions) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TaskPythonConfigTaskOptions) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *TaskPythonConfigTaskOptions) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *TaskPythonConfigTaskOptions) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *TaskPythonConfigTaskOptions) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *TaskPythonConfigTaskOptions) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *TaskPythonConfigTaskOptions) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *TaskPythonConfigTaskOptions) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *TaskPythonConfigTaskOptions) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *TaskPythonConfigTaskOptions) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetLocalScriptGitId

`func (o *TaskPythonConfigTaskOptions) GetLocalScriptGitId() string`

GetLocalScriptGitId returns the LocalScriptGitId field if non-nil, zero value otherwise.

### GetLocalScriptGitIdOk

`func (o *TaskPythonConfigTaskOptions) GetLocalScriptGitIdOk() (*string, bool)`

GetLocalScriptGitIdOk returns a tuple with the LocalScriptGitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitId

`func (o *TaskPythonConfigTaskOptions) SetLocalScriptGitId(v string)`

SetLocalScriptGitId sets LocalScriptGitId field to given value.

### HasLocalScriptGitId

`func (o *TaskPythonConfigTaskOptions) HasLocalScriptGitId() bool`

HasLocalScriptGitId returns a boolean if a field has been set.

### SetLocalScriptGitIdNil

`func (o *TaskPythonConfigTaskOptions) SetLocalScriptGitIdNil(b bool)`

 SetLocalScriptGitIdNil sets the value for LocalScriptGitId to be an explicit nil

### UnsetLocalScriptGitId
`func (o *TaskPythonConfigTaskOptions) UnsetLocalScriptGitId()`

UnsetLocalScriptGitId ensures that no value is present for LocalScriptGitId, not even an explicit nil
### GetLocalScriptGitRef

`func (o *TaskPythonConfigTaskOptions) GetLocalScriptGitRef() string`

GetLocalScriptGitRef returns the LocalScriptGitRef field if non-nil, zero value otherwise.

### GetLocalScriptGitRefOk

`func (o *TaskPythonConfigTaskOptions) GetLocalScriptGitRefOk() (*string, bool)`

GetLocalScriptGitRefOk returns a tuple with the LocalScriptGitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalScriptGitRef

`func (o *TaskPythonConfigTaskOptions) SetLocalScriptGitRef(v string)`

SetLocalScriptGitRef sets LocalScriptGitRef field to given value.

### HasLocalScriptGitRef

`func (o *TaskPythonConfigTaskOptions) HasLocalScriptGitRef() bool`

HasLocalScriptGitRef returns a boolean if a field has been set.

### SetLocalScriptGitRefNil

`func (o *TaskPythonConfigTaskOptions) SetLocalScriptGitRefNil(b bool)`

 SetLocalScriptGitRefNil sets the value for LocalScriptGitRef to be an explicit nil

### UnsetLocalScriptGitRef
`func (o *TaskPythonConfigTaskOptions) UnsetLocalScriptGitRef()`

UnsetLocalScriptGitRef ensures that no value is present for LocalScriptGitRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


