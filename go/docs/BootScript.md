# BootScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**NullableArchiveBucketCreatedBy**](archiveBucket_createdBy.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewBootScript

`func NewBootScript() *BootScript`

NewBootScript instantiates a new BootScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootScriptWithDefaults

`func NewBootScriptWithDefaults() *BootScript`

NewBootScriptWithDefaults instantiates a new BootScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BootScript) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BootScript) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BootScript) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BootScript) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *BootScript) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BootScript) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BootScript) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BootScript) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFileName

`func (o *BootScript) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *BootScript) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *BootScript) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *BootScript) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetDescription

`func (o *BootScript) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BootScript) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BootScript) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BootScript) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BootScript) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BootScript) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContent

`func (o *BootScript) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BootScript) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BootScript) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BootScript) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BootScript) GetCreatedBy() ArchiveBucketCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BootScript) GetCreatedByOk() (*ArchiveBucketCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BootScript) SetCreatedBy(v ArchiveBucketCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BootScript) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *BootScript) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *BootScript) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetVisibility

`func (o *BootScript) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BootScript) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BootScript) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BootScript) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


