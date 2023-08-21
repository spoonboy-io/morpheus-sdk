# InlineResponse20095NetworkRouterType

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
**HasNat** | Pointer to **bool** |  | [optional] 
**HasNetworkServer** | Pointer to **bool** |  | [optional] 
**HasFirewallGroups** | Pointer to **bool** |  | [optional] 
**HasSecurityGroupPriority** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 
**RuleOptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 
**FirewallGroupOptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 
**NatOptionTypes** | Pointer to [**[]OptionType**](OptionType.md) |  | [optional] 

## Methods

### NewInlineResponse20095NetworkRouterType

`func NewInlineResponse20095NetworkRouterType() *InlineResponse20095NetworkRouterType`

NewInlineResponse20095NetworkRouterType instantiates a new InlineResponse20095NetworkRouterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095NetworkRouterTypeWithDefaults

`func NewInlineResponse20095NetworkRouterTypeWithDefaults() *InlineResponse20095NetworkRouterType`

NewInlineResponse20095NetworkRouterTypeWithDefaults instantiates a new InlineResponse20095NetworkRouterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095NetworkRouterType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095NetworkRouterType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095NetworkRouterType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095NetworkRouterType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20095NetworkRouterType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20095NetworkRouterType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20095NetworkRouterType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20095NetworkRouterType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095NetworkRouterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095NetworkRouterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095NetworkRouterType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095NetworkRouterType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20095NetworkRouterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20095NetworkRouterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20095NetworkRouterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20095NetworkRouterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20095NetworkRouterType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20095NetworkRouterType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20095NetworkRouterType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20095NetworkRouterType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *InlineResponse20095NetworkRouterType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InlineResponse20095NetworkRouterType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InlineResponse20095NetworkRouterType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InlineResponse20095NetworkRouterType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetSelectable

`func (o *InlineResponse20095NetworkRouterType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *InlineResponse20095NetworkRouterType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *InlineResponse20095NetworkRouterType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *InlineResponse20095NetworkRouterType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetHasFirewall

`func (o *InlineResponse20095NetworkRouterType) GetHasFirewall() bool`

GetHasFirewall returns the HasFirewall field if non-nil, zero value otherwise.

### GetHasFirewallOk

`func (o *InlineResponse20095NetworkRouterType) GetHasFirewallOk() (*bool, bool)`

GetHasFirewallOk returns a tuple with the HasFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewall

`func (o *InlineResponse20095NetworkRouterType) SetHasFirewall(v bool)`

SetHasFirewall sets HasFirewall field to given value.

### HasHasFirewall

`func (o *InlineResponse20095NetworkRouterType) HasHasFirewall() bool`

HasHasFirewall returns a boolean if a field has been set.

### GetHasDhcp

`func (o *InlineResponse20095NetworkRouterType) GetHasDhcp() bool`

GetHasDhcp returns the HasDhcp field if non-nil, zero value otherwise.

### GetHasDhcpOk

`func (o *InlineResponse20095NetworkRouterType) GetHasDhcpOk() (*bool, bool)`

GetHasDhcpOk returns a tuple with the HasDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDhcp

`func (o *InlineResponse20095NetworkRouterType) SetHasDhcp(v bool)`

SetHasDhcp sets HasDhcp field to given value.

### HasHasDhcp

`func (o *InlineResponse20095NetworkRouterType) HasHasDhcp() bool`

HasHasDhcp returns a boolean if a field has been set.

### GetHasRouting

`func (o *InlineResponse20095NetworkRouterType) GetHasRouting() bool`

GetHasRouting returns the HasRouting field if non-nil, zero value otherwise.

### GetHasRoutingOk

`func (o *InlineResponse20095NetworkRouterType) GetHasRoutingOk() (*bool, bool)`

GetHasRoutingOk returns a tuple with the HasRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouting

`func (o *InlineResponse20095NetworkRouterType) SetHasRouting(v bool)`

SetHasRouting sets HasRouting field to given value.

### HasHasRouting

`func (o *InlineResponse20095NetworkRouterType) HasHasRouting() bool`

HasHasRouting returns a boolean if a field has been set.

### GetHasNat

`func (o *InlineResponse20095NetworkRouterType) GetHasNat() bool`

GetHasNat returns the HasNat field if non-nil, zero value otherwise.

### GetHasNatOk

`func (o *InlineResponse20095NetworkRouterType) GetHasNatOk() (*bool, bool)`

GetHasNatOk returns a tuple with the HasNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNat

`func (o *InlineResponse20095NetworkRouterType) SetHasNat(v bool)`

SetHasNat sets HasNat field to given value.

### HasHasNat

`func (o *InlineResponse20095NetworkRouterType) HasHasNat() bool`

HasHasNat returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *InlineResponse20095NetworkRouterType) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *InlineResponse20095NetworkRouterType) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *InlineResponse20095NetworkRouterType) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *InlineResponse20095NetworkRouterType) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasFirewallGroups

`func (o *InlineResponse20095NetworkRouterType) GetHasFirewallGroups() bool`

GetHasFirewallGroups returns the HasFirewallGroups field if non-nil, zero value otherwise.

### GetHasFirewallGroupsOk

`func (o *InlineResponse20095NetworkRouterType) GetHasFirewallGroupsOk() (*bool, bool)`

GetHasFirewallGroupsOk returns a tuple with the HasFirewallGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewallGroups

`func (o *InlineResponse20095NetworkRouterType) SetHasFirewallGroups(v bool)`

SetHasFirewallGroups sets HasFirewallGroups field to given value.

### HasHasFirewallGroups

`func (o *InlineResponse20095NetworkRouterType) HasHasFirewallGroups() bool`

HasHasFirewallGroups returns a boolean if a field has been set.

### GetHasSecurityGroupPriority

`func (o *InlineResponse20095NetworkRouterType) GetHasSecurityGroupPriority() bool`

GetHasSecurityGroupPriority returns the HasSecurityGroupPriority field if non-nil, zero value otherwise.

### GetHasSecurityGroupPriorityOk

`func (o *InlineResponse20095NetworkRouterType) GetHasSecurityGroupPriorityOk() (*bool, bool)`

GetHasSecurityGroupPriorityOk returns a tuple with the HasSecurityGroupPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroupPriority

`func (o *InlineResponse20095NetworkRouterType) SetHasSecurityGroupPriority(v bool)`

SetHasSecurityGroupPriority sets HasSecurityGroupPriority field to given value.

### HasHasSecurityGroupPriority

`func (o *InlineResponse20095NetworkRouterType) HasHasSecurityGroupPriority() bool`

HasHasSecurityGroupPriority returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InlineResponse20095NetworkRouterType) GetOptionTypes() []OptionType`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InlineResponse20095NetworkRouterType) GetOptionTypesOk() (*[]OptionType, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InlineResponse20095NetworkRouterType) SetOptionTypes(v []OptionType)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InlineResponse20095NetworkRouterType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRuleOptionTypes

`func (o *InlineResponse20095NetworkRouterType) GetRuleOptionTypes() []OptionType`

GetRuleOptionTypes returns the RuleOptionTypes field if non-nil, zero value otherwise.

### GetRuleOptionTypesOk

`func (o *InlineResponse20095NetworkRouterType) GetRuleOptionTypesOk() (*[]OptionType, bool)`

GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOptionTypes

`func (o *InlineResponse20095NetworkRouterType) SetRuleOptionTypes(v []OptionType)`

SetRuleOptionTypes sets RuleOptionTypes field to given value.

### HasRuleOptionTypes

`func (o *InlineResponse20095NetworkRouterType) HasRuleOptionTypes() bool`

HasRuleOptionTypes returns a boolean if a field has been set.

### GetFirewallGroupOptionTypes

`func (o *InlineResponse20095NetworkRouterType) GetFirewallGroupOptionTypes() []OptionType`

GetFirewallGroupOptionTypes returns the FirewallGroupOptionTypes field if non-nil, zero value otherwise.

### GetFirewallGroupOptionTypesOk

`func (o *InlineResponse20095NetworkRouterType) GetFirewallGroupOptionTypesOk() (*[]OptionType, bool)`

GetFirewallGroupOptionTypesOk returns a tuple with the FirewallGroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallGroupOptionTypes

`func (o *InlineResponse20095NetworkRouterType) SetFirewallGroupOptionTypes(v []OptionType)`

SetFirewallGroupOptionTypes sets FirewallGroupOptionTypes field to given value.

### HasFirewallGroupOptionTypes

`func (o *InlineResponse20095NetworkRouterType) HasFirewallGroupOptionTypes() bool`

HasFirewallGroupOptionTypes returns a boolean if a field has been set.

### GetNatOptionTypes

`func (o *InlineResponse20095NetworkRouterType) GetNatOptionTypes() []OptionType`

GetNatOptionTypes returns the NatOptionTypes field if non-nil, zero value otherwise.

### GetNatOptionTypesOk

`func (o *InlineResponse20095NetworkRouterType) GetNatOptionTypesOk() (*[]OptionType, bool)`

GetNatOptionTypesOk returns a tuple with the NatOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOptionTypes

`func (o *InlineResponse20095NetworkRouterType) SetNatOptionTypes(v []OptionType)`

SetNatOptionTypes sets NatOptionTypes field to given value.

### HasNatOptionTypes

`func (o *InlineResponse20095NetworkRouterType) HasNatOptionTypes() bool`

HasNatOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


