# TaskType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Scriptable** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HasResults** | Pointer to **bool** |  | [optional] 
**AllowExecuteLocal** | Pointer to **NullableBool** |  | [optional] 
**AllowExecuteRemote** | Pointer to **NullableBool** |  | [optional] 
**AllowExecuteResource** | Pointer to **NullableBool** |  | [optional] 
**AllowLocalRepo** | Pointer to **NullableBool** |  | [optional] 
**AllowRemoteKeyAuth** | Pointer to **NullableBool** |  | [optional] 
**OptionTypes** | Pointer to [**[]TaskTypeOptionTypes**](TaskTypeOptionTypes.md) |  | [optional] 

## Methods

### NewTaskType

`func NewTaskType() *TaskType`

NewTaskType instantiates a new TaskType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTypeWithDefaults

`func NewTaskTypeWithDefaults() *TaskType`

NewTaskTypeWithDefaults instantiates a new TaskType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *TaskType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaskType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *TaskType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *TaskType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaskType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaskType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaskType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *TaskType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScriptable

`func (o *TaskType) GetScriptable() bool`

GetScriptable returns the Scriptable field if non-nil, zero value otherwise.

### GetScriptableOk

`func (o *TaskType) GetScriptableOk() (*bool, bool)`

GetScriptableOk returns a tuple with the Scriptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptable

`func (o *TaskType) SetScriptable(v bool)`

SetScriptable sets Scriptable field to given value.

### HasScriptable

`func (o *TaskType) HasScriptable() bool`

HasScriptable returns a boolean if a field has been set.

### GetEnabled

`func (o *TaskType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TaskType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TaskType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TaskType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasResults

`func (o *TaskType) GetHasResults() bool`

GetHasResults returns the HasResults field if non-nil, zero value otherwise.

### GetHasResultsOk

`func (o *TaskType) GetHasResultsOk() (*bool, bool)`

GetHasResultsOk returns a tuple with the HasResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResults

`func (o *TaskType) SetHasResults(v bool)`

SetHasResults sets HasResults field to given value.

### HasHasResults

`func (o *TaskType) HasHasResults() bool`

HasHasResults returns a boolean if a field has been set.

### GetAllowExecuteLocal

`func (o *TaskType) GetAllowExecuteLocal() bool`

GetAllowExecuteLocal returns the AllowExecuteLocal field if non-nil, zero value otherwise.

### GetAllowExecuteLocalOk

`func (o *TaskType) GetAllowExecuteLocalOk() (*bool, bool)`

GetAllowExecuteLocalOk returns a tuple with the AllowExecuteLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteLocal

`func (o *TaskType) SetAllowExecuteLocal(v bool)`

SetAllowExecuteLocal sets AllowExecuteLocal field to given value.

### HasAllowExecuteLocal

`func (o *TaskType) HasAllowExecuteLocal() bool`

HasAllowExecuteLocal returns a boolean if a field has been set.

### SetAllowExecuteLocalNil

`func (o *TaskType) SetAllowExecuteLocalNil(b bool)`

 SetAllowExecuteLocalNil sets the value for AllowExecuteLocal to be an explicit nil

### UnsetAllowExecuteLocal
`func (o *TaskType) UnsetAllowExecuteLocal()`

UnsetAllowExecuteLocal ensures that no value is present for AllowExecuteLocal, not even an explicit nil
### GetAllowExecuteRemote

`func (o *TaskType) GetAllowExecuteRemote() bool`

GetAllowExecuteRemote returns the AllowExecuteRemote field if non-nil, zero value otherwise.

### GetAllowExecuteRemoteOk

`func (o *TaskType) GetAllowExecuteRemoteOk() (*bool, bool)`

GetAllowExecuteRemoteOk returns a tuple with the AllowExecuteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteRemote

`func (o *TaskType) SetAllowExecuteRemote(v bool)`

SetAllowExecuteRemote sets AllowExecuteRemote field to given value.

### HasAllowExecuteRemote

`func (o *TaskType) HasAllowExecuteRemote() bool`

HasAllowExecuteRemote returns a boolean if a field has been set.

### SetAllowExecuteRemoteNil

`func (o *TaskType) SetAllowExecuteRemoteNil(b bool)`

 SetAllowExecuteRemoteNil sets the value for AllowExecuteRemote to be an explicit nil

### UnsetAllowExecuteRemote
`func (o *TaskType) UnsetAllowExecuteRemote()`

UnsetAllowExecuteRemote ensures that no value is present for AllowExecuteRemote, not even an explicit nil
### GetAllowExecuteResource

`func (o *TaskType) GetAllowExecuteResource() bool`

GetAllowExecuteResource returns the AllowExecuteResource field if non-nil, zero value otherwise.

### GetAllowExecuteResourceOk

`func (o *TaskType) GetAllowExecuteResourceOk() (*bool, bool)`

GetAllowExecuteResourceOk returns a tuple with the AllowExecuteResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExecuteResource

`func (o *TaskType) SetAllowExecuteResource(v bool)`

SetAllowExecuteResource sets AllowExecuteResource field to given value.

### HasAllowExecuteResource

`func (o *TaskType) HasAllowExecuteResource() bool`

HasAllowExecuteResource returns a boolean if a field has been set.

### SetAllowExecuteResourceNil

`func (o *TaskType) SetAllowExecuteResourceNil(b bool)`

 SetAllowExecuteResourceNil sets the value for AllowExecuteResource to be an explicit nil

### UnsetAllowExecuteResource
`func (o *TaskType) UnsetAllowExecuteResource()`

UnsetAllowExecuteResource ensures that no value is present for AllowExecuteResource, not even an explicit nil
### GetAllowLocalRepo

`func (o *TaskType) GetAllowLocalRepo() bool`

GetAllowLocalRepo returns the AllowLocalRepo field if non-nil, zero value otherwise.

### GetAllowLocalRepoOk

`func (o *TaskType) GetAllowLocalRepoOk() (*bool, bool)`

GetAllowLocalRepoOk returns a tuple with the AllowLocalRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLocalRepo

`func (o *TaskType) SetAllowLocalRepo(v bool)`

SetAllowLocalRepo sets AllowLocalRepo field to given value.

### HasAllowLocalRepo

`func (o *TaskType) HasAllowLocalRepo() bool`

HasAllowLocalRepo returns a boolean if a field has been set.

### SetAllowLocalRepoNil

`func (o *TaskType) SetAllowLocalRepoNil(b bool)`

 SetAllowLocalRepoNil sets the value for AllowLocalRepo to be an explicit nil

### UnsetAllowLocalRepo
`func (o *TaskType) UnsetAllowLocalRepo()`

UnsetAllowLocalRepo ensures that no value is present for AllowLocalRepo, not even an explicit nil
### GetAllowRemoteKeyAuth

`func (o *TaskType) GetAllowRemoteKeyAuth() bool`

GetAllowRemoteKeyAuth returns the AllowRemoteKeyAuth field if non-nil, zero value otherwise.

### GetAllowRemoteKeyAuthOk

`func (o *TaskType) GetAllowRemoteKeyAuthOk() (*bool, bool)`

GetAllowRemoteKeyAuthOk returns a tuple with the AllowRemoteKeyAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteKeyAuth

`func (o *TaskType) SetAllowRemoteKeyAuth(v bool)`

SetAllowRemoteKeyAuth sets AllowRemoteKeyAuth field to given value.

### HasAllowRemoteKeyAuth

`func (o *TaskType) HasAllowRemoteKeyAuth() bool`

HasAllowRemoteKeyAuth returns a boolean if a field has been set.

### SetAllowRemoteKeyAuthNil

`func (o *TaskType) SetAllowRemoteKeyAuthNil(b bool)`

 SetAllowRemoteKeyAuthNil sets the value for AllowRemoteKeyAuth to be an explicit nil

### UnsetAllowRemoteKeyAuth
`func (o *TaskType) UnsetAllowRemoteKeyAuth()`

UnsetAllowRemoteKeyAuth ensures that no value is present for AllowRemoteKeyAuth, not even an explicit nil
### GetOptionTypes

`func (o *TaskType) GetOptionTypes() []TaskTypeOptionTypes`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *TaskType) GetOptionTypesOk() (*[]TaskTypeOptionTypes, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *TaskType) SetOptionTypes(v []TaskTypeOptionTypes)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *TaskType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


