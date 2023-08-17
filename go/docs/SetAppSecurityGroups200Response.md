# SetAppSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]AppSecurityGroups**](AppSecurityGroups.md) |  | [optional] 

## Methods

### NewSetAppSecurityGroups200Response

`func NewSetAppSecurityGroups200Response() *SetAppSecurityGroups200Response`

NewSetAppSecurityGroups200Response instantiates a new SetAppSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAppSecurityGroups200ResponseWithDefaults

`func NewSetAppSecurityGroups200ResponseWithDefaults() *SetAppSecurityGroups200Response`

NewSetAppSecurityGroups200ResponseWithDefaults instantiates a new SetAppSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SetAppSecurityGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SetAppSecurityGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SetAppSecurityGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SetAppSecurityGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *SetAppSecurityGroups200Response) GetSecurityGroups() []AppSecurityGroups`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SetAppSecurityGroups200Response) GetSecurityGroupsOk() (*[]AppSecurityGroups, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SetAppSecurityGroups200Response) SetSecurityGroups(v []AppSecurityGroups)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *SetAppSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


