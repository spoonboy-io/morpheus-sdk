# InlineResponse20077LoadBalancerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**OptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 
**VipOptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 

## Methods

### NewInlineResponse20077LoadBalancerType

`func NewInlineResponse20077LoadBalancerType() *InlineResponse20077LoadBalancerType`

NewInlineResponse20077LoadBalancerType instantiates a new InlineResponse20077LoadBalancerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20077LoadBalancerTypeWithDefaults

`func NewInlineResponse20077LoadBalancerTypeWithDefaults() *InlineResponse20077LoadBalancerType`

NewInlineResponse20077LoadBalancerTypeWithDefaults instantiates a new InlineResponse20077LoadBalancerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20077LoadBalancerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20077LoadBalancerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20077LoadBalancerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20077LoadBalancerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20077LoadBalancerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20077LoadBalancerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20077LoadBalancerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20077LoadBalancerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20077LoadBalancerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20077LoadBalancerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20077LoadBalancerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20077LoadBalancerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20077LoadBalancerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20077LoadBalancerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20077LoadBalancerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20077LoadBalancerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInternal

`func (o *InlineResponse20077LoadBalancerType) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *InlineResponse20077LoadBalancerType) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *InlineResponse20077LoadBalancerType) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *InlineResponse20077LoadBalancerType) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetCreatable

`func (o *InlineResponse20077LoadBalancerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InlineResponse20077LoadBalancerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InlineResponse20077LoadBalancerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InlineResponse20077LoadBalancerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetCreateType

`func (o *InlineResponse20077LoadBalancerType) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *InlineResponse20077LoadBalancerType) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *InlineResponse20077LoadBalancerType) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *InlineResponse20077LoadBalancerType) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InlineResponse20077LoadBalancerType) GetOptionTypes() []OptionType`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InlineResponse20077LoadBalancerType) GetOptionTypesOk() (*[]OptionType, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InlineResponse20077LoadBalancerType) SetOptionTypes(v []OptionType)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InlineResponse20077LoadBalancerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetVipOptionTypes

`func (o *InlineResponse20077LoadBalancerType) GetVipOptionTypes() []OptionType`

GetVipOptionTypes returns the VipOptionTypes field if non-nil, zero value otherwise.

### GetVipOptionTypesOk

`func (o *InlineResponse20077LoadBalancerType) GetVipOptionTypesOk() (*[]OptionType, bool)`

GetVipOptionTypesOk returns a tuple with the VipOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipOptionTypes

`func (o *InlineResponse20077LoadBalancerType) SetVipOptionTypes(v []OptionType)`

SetVipOptionTypes sets VipOptionTypes field to given value.

### HasVipOptionTypes

`func (o *InlineResponse20077LoadBalancerType) HasVipOptionTypes() bool`

HasVipOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


