# SuccessMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Msg** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSuccessMessage

`func NewSuccessMessage() *SuccessMessage`

NewSuccessMessage instantiates a new SuccessMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessMessageWithDefaults

`func NewSuccessMessageWithDefaults() *SuccessMessage`

NewSuccessMessageWithDefaults instantiates a new SuccessMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SuccessMessage) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SuccessMessage) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SuccessMessage) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SuccessMessage) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *SuccessMessage) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SuccessMessage) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SuccessMessage) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SuccessMessage) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *SuccessMessage) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *SuccessMessage) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


