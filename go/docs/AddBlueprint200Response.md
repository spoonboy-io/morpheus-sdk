# AddBlueprint200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Blueprint** | Pointer to [**BlueprintCreateSuccess**](BlueprintCreateSuccess.md) |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 
**Errors** | Pointer to **NullableString** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddBlueprint200Response

`func NewAddBlueprint200Response() *AddBlueprint200Response`

NewAddBlueprint200Response instantiates a new AddBlueprint200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprint200ResponseWithDefaults

`func NewAddBlueprint200ResponseWithDefaults() *AddBlueprint200Response`

NewAddBlueprint200ResponseWithDefaults instantiates a new AddBlueprint200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddBlueprint200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddBlueprint200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddBlueprint200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddBlueprint200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBlueprint

`func (o *AddBlueprint200Response) GetBlueprint() BlueprintCreateSuccess`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *AddBlueprint200Response) GetBlueprintOk() (*BlueprintCreateSuccess, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *AddBlueprint200Response) SetBlueprint(v BlueprintCreateSuccess)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *AddBlueprint200Response) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetMsg

`func (o *AddBlueprint200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AddBlueprint200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AddBlueprint200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AddBlueprint200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *AddBlueprint200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *AddBlueprint200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetErrors

`func (o *AddBlueprint200Response) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddBlueprint200Response) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddBlueprint200Response) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddBlueprint200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *AddBlueprint200Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *AddBlueprint200Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetErrorCode

`func (o *AddBlueprint200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AddBlueprint200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AddBlueprint200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AddBlueprint200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AddBlueprint200Response) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AddBlueprint200Response) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetInProgress

`func (o *AddBlueprint200Response) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *AddBlueprint200Response) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *AddBlueprint200Response) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *AddBlueprint200Response) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


