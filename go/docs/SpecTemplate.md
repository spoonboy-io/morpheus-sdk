# SpecTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**DeploymentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**ClusterLayoutFile**](clusterLayout_file.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSpecTemplate

`func NewSpecTemplate() *SpecTemplate`

NewSpecTemplate instantiates a new SpecTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecTemplateWithDefaults

`func NewSpecTemplateWithDefaults() *SpecTemplate`

NewSpecTemplateWithDefaults instantiates a new SpecTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpecTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpecTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpecTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SpecTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *SpecTemplate) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SpecTemplate) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SpecTemplate) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SpecTemplate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *SpecTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpecTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpecTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpecTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *SpecTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *SpecTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *SpecTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *SpecTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *SpecTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpecTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpecTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpecTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *SpecTemplate) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SpecTemplate) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetType

`func (o *SpecTemplate) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpecTemplate) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpecTemplate) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *SpecTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *SpecTemplate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SpecTemplate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SpecTemplate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SpecTemplate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *SpecTemplate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *SpecTemplate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalType

`func (o *SpecTemplate) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *SpecTemplate) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *SpecTemplate) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *SpecTemplate) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *SpecTemplate) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *SpecTemplate) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetDeploymentId

`func (o *SpecTemplate) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *SpecTemplate) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *SpecTemplate) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *SpecTemplate) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *SpecTemplate) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *SpecTemplate) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetStatus

`func (o *SpecTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpecTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpecTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpecTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SpecTemplate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SpecTemplate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFile

`func (o *SpecTemplate) GetFile() ClusterLayoutFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SpecTemplate) GetFileOk() (*ClusterLayoutFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SpecTemplate) SetFile(v ClusterLayoutFile)`

SetFile sets File field to given value.

### HasFile

`func (o *SpecTemplate) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *SpecTemplate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SpecTemplate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SpecTemplate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SpecTemplate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SpecTemplate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpecTemplate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpecTemplate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpecTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SpecTemplate) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SpecTemplate) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SpecTemplate) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SpecTemplate) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *SpecTemplate) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *SpecTemplate) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetDateCreated

`func (o *SpecTemplate) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SpecTemplate) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SpecTemplate) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SpecTemplate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SpecTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SpecTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SpecTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SpecTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


