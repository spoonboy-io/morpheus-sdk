# InlineResponse20094Type

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Selectable** | Pointer to **bool** |  | [optional] 
**HasFirewall** | Pointer to **bool** |  | [optional] 
**HasDhcp** | Pointer to **bool** |  | [optional] 
**HasRouting** | Pointer to **bool** |  | [optional] 
**HasNetworkServer** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RuleOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInlineResponse20094Type

`func NewInlineResponse20094Type() *InlineResponse20094Type`

NewInlineResponse20094Type instantiates a new InlineResponse20094Type object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20094TypeWithDefaults

`func NewInlineResponse20094TypeWithDefaults() *InlineResponse20094Type`

NewInlineResponse20094TypeWithDefaults instantiates a new InlineResponse20094Type object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20094Type) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20094Type) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20094Type) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20094Type) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20094Type) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20094Type) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20094Type) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20094Type) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20094Type) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20094Type) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20094Type) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20094Type) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20094Type) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20094Type) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20094Type) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20094Type) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20094Type) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20094Type) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20094Type) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20094Type) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *InlineResponse20094Type) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InlineResponse20094Type) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InlineResponse20094Type) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InlineResponse20094Type) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetSelectable

`func (o *InlineResponse20094Type) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *InlineResponse20094Type) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *InlineResponse20094Type) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *InlineResponse20094Type) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetHasFirewall

`func (o *InlineResponse20094Type) GetHasFirewall() bool`

GetHasFirewall returns the HasFirewall field if non-nil, zero value otherwise.

### GetHasFirewallOk

`func (o *InlineResponse20094Type) GetHasFirewallOk() (*bool, bool)`

GetHasFirewallOk returns a tuple with the HasFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewall

`func (o *InlineResponse20094Type) SetHasFirewall(v bool)`

SetHasFirewall sets HasFirewall field to given value.

### HasHasFirewall

`func (o *InlineResponse20094Type) HasHasFirewall() bool`

HasHasFirewall returns a boolean if a field has been set.

### GetHasDhcp

`func (o *InlineResponse20094Type) GetHasDhcp() bool`

GetHasDhcp returns the HasDhcp field if non-nil, zero value otherwise.

### GetHasDhcpOk

`func (o *InlineResponse20094Type) GetHasDhcpOk() (*bool, bool)`

GetHasDhcpOk returns a tuple with the HasDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDhcp

`func (o *InlineResponse20094Type) SetHasDhcp(v bool)`

SetHasDhcp sets HasDhcp field to given value.

### HasHasDhcp

`func (o *InlineResponse20094Type) HasHasDhcp() bool`

HasHasDhcp returns a boolean if a field has been set.

### GetHasRouting

`func (o *InlineResponse20094Type) GetHasRouting() bool`

GetHasRouting returns the HasRouting field if non-nil, zero value otherwise.

### GetHasRoutingOk

`func (o *InlineResponse20094Type) GetHasRoutingOk() (*bool, bool)`

GetHasRoutingOk returns a tuple with the HasRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouting

`func (o *InlineResponse20094Type) SetHasRouting(v bool)`

SetHasRouting sets HasRouting field to given value.

### HasHasRouting

`func (o *InlineResponse20094Type) HasHasRouting() bool`

HasHasRouting returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *InlineResponse20094Type) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *InlineResponse20094Type) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *InlineResponse20094Type) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *InlineResponse20094Type) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InlineResponse20094Type) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InlineResponse20094Type) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InlineResponse20094Type) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InlineResponse20094Type) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRuleOptionTypes

`func (o *InlineResponse20094Type) GetRuleOptionTypes() []map[string]interface{}`

GetRuleOptionTypes returns the RuleOptionTypes field if non-nil, zero value otherwise.

### GetRuleOptionTypesOk

`func (o *InlineResponse20094Type) GetRuleOptionTypesOk() (*[]map[string]interface{}, bool)`

GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOptionTypes

`func (o *InlineResponse20094Type) SetRuleOptionTypes(v []map[string]interface{})`

SetRuleOptionTypes sets RuleOptionTypes field to given value.

### HasRuleOptionTypes

`func (o *InlineResponse20094Type) HasRuleOptionTypes() bool`

HasRuleOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


