# GroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the group | 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Location** | Pointer to **string** | Optional location argument for your group | [optional] 
**Config** | Pointer to [**GroupCreateConfig**](groupCreate_config.md) |  | [optional] 

## Methods

### NewGroupUpdate

`func NewGroupUpdate(name string, ) *GroupUpdate`

NewGroupUpdate instantiates a new GroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateWithDefaults

`func NewGroupUpdateWithDefaults() *GroupUpdate`

NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *GroupUpdate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GroupUpdate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GroupUpdate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GroupUpdate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *GroupUpdate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GroupUpdate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GroupUpdate) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GroupUpdate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConfig

`func (o *GroupUpdate) GetConfig() GroupCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GroupUpdate) GetConfigOk() (*GroupCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GroupUpdate) SetConfig(v GroupCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GroupUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


