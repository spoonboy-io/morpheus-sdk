# Approval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalName** | Pointer to **NullableString** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Approver** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**AccountIntegration** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RequestBy** | Pointer to **string** |  | [optional] 
**ApprovalItems** | Pointer to [**[]ApprovalItemApprovalItem**](ApprovalItemApprovalItem.md) |  | [optional] 

## Methods

### NewApproval

`func NewApproval() *Approval`

NewApproval instantiates a new Approval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWithDefaults

`func NewApprovalWithDefaults() *Approval`

NewApprovalWithDefaults instantiates a new Approval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Approval) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Approval) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Approval) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Approval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Approval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Approval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Approval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Approval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInternalId

`func (o *Approval) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *Approval) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *Approval) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *Approval) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *Approval) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *Approval) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *Approval) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Approval) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Approval) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Approval) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Approval) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Approval) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalName

`func (o *Approval) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *Approval) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *Approval) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *Approval) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### SetExternalNameNil

`func (o *Approval) SetExternalNameNil(b bool)`

 SetExternalNameNil sets the value for ExternalName to be an explicit nil

### UnsetExternalName
`func (o *Approval) UnsetExternalName()`

UnsetExternalName ensures that no value is present for ExternalName, not even an explicit nil
### GetRequestType

`func (o *Approval) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *Approval) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *Approval) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *Approval) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetAccount

`func (o *Approval) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Approval) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Approval) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Approval) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetApprover

`func (o *Approval) GetApprover() InlineResponse20040AppDeployInstance`

GetApprover returns the Approver field if non-nil, zero value otherwise.

### GetApproverOk

`func (o *Approval) GetApproverOk() (*InlineResponse20040AppDeployInstance, bool)`

GetApproverOk returns a tuple with the Approver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprover

`func (o *Approval) SetApprover(v InlineResponse20040AppDeployInstance)`

SetApprover sets Approver field to given value.

### HasApprover

`func (o *Approval) HasApprover() bool`

HasApprover returns a boolean if a field has been set.

### GetAccountIntegration

`func (o *Approval) GetAccountIntegration() string`

GetAccountIntegration returns the AccountIntegration field if non-nil, zero value otherwise.

### GetAccountIntegrationOk

`func (o *Approval) GetAccountIntegrationOk() (*string, bool)`

GetAccountIntegrationOk returns a tuple with the AccountIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegration

`func (o *Approval) SetAccountIntegration(v string)`

SetAccountIntegration sets AccountIntegration field to given value.

### HasAccountIntegration

`func (o *Approval) HasAccountIntegration() bool`

HasAccountIntegration returns a boolean if a field has been set.

### SetAccountIntegrationNil

`func (o *Approval) SetAccountIntegrationNil(b bool)`

 SetAccountIntegrationNil sets the value for AccountIntegration to be an explicit nil

### UnsetAccountIntegration
`func (o *Approval) UnsetAccountIntegration()`

UnsetAccountIntegration ensures that no value is present for AccountIntegration, not even an explicit nil
### GetStatus

`func (o *Approval) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Approval) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Approval) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Approval) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Approval) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Approval) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Approval) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Approval) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Approval) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Approval) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDateCreated

`func (o *Approval) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Approval) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Approval) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Approval) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Approval) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Approval) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Approval) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Approval) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRequestBy

`func (o *Approval) GetRequestBy() string`

GetRequestBy returns the RequestBy field if non-nil, zero value otherwise.

### GetRequestByOk

`func (o *Approval) GetRequestByOk() (*string, bool)`

GetRequestByOk returns a tuple with the RequestBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBy

`func (o *Approval) SetRequestBy(v string)`

SetRequestBy sets RequestBy field to given value.

### HasRequestBy

`func (o *Approval) HasRequestBy() bool`

HasRequestBy returns a boolean if a field has been set.

### GetApprovalItems

`func (o *Approval) GetApprovalItems() []ApprovalItemApprovalItem`

GetApprovalItems returns the ApprovalItems field if non-nil, zero value otherwise.

### GetApprovalItemsOk

`func (o *Approval) GetApprovalItemsOk() (*[]ApprovalItemApprovalItem, bool)`

GetApprovalItemsOk returns a tuple with the ApprovalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalItems

`func (o *Approval) SetApprovalItems(v []ApprovalItemApprovalItem)`

SetApprovalItems sets ApprovalItems field to given value.

### HasApprovalItems

`func (o *Approval) HasApprovalItems() bool`

HasApprovalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


