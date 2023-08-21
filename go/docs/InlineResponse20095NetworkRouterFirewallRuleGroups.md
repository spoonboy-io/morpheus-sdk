# InlineResponse20095NetworkRouterFirewallRuleGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**IacId** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZonePool** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 
**GroupLayer** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse20095NetworkRouterFirewallRules**](InlineResponse20095NetworkRouterFirewallRules.md) |  | [optional] 

## Methods

### NewInlineResponse20095NetworkRouterFirewallRuleGroups

`func NewInlineResponse20095NetworkRouterFirewallRuleGroups() *InlineResponse20095NetworkRouterFirewallRuleGroups`

NewInlineResponse20095NetworkRouterFirewallRuleGroups instantiates a new InlineResponse20095NetworkRouterFirewallRuleGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095NetworkRouterFirewallRuleGroupsWithDefaults

`func NewInlineResponse20095NetworkRouterFirewallRuleGroupsWithDefaults() *InlineResponse20095NetworkRouterFirewallRuleGroups`

NewInlineResponse20095NetworkRouterFirewallRuleGroupsWithDefaults instantiates a new InlineResponse20095NetworkRouterFirewallRuleGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIacId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetIacId() string`

GetIacId returns the IacId field if non-nil, zero value otherwise.

### GetIacIdOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetIacIdOk() (*string, bool)`

GetIacIdOk returns a tuple with the IacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetIacId(v string)`

SetIacId sets IacId field to given value.

### HasIacId

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasIacId() bool`

HasIacId returns a boolean if a field has been set.

### SetIacIdNil

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetIacIdNil(b bool)`

 SetIacIdNil sets the value for IacId to be an explicit nil

### UnsetIacId
`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) UnsetIacId()`

UnsetIacId ensures that no value is present for IacId, not even an explicit nil
### GetZone

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZonePool

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetZonePool() string`

GetZonePool returns the ZonePool field if non-nil, zero value otherwise.

### GetZonePoolOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetZonePoolOk() (*string, bool)`

GetZonePoolOk returns a tuple with the ZonePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonePool

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetZonePool(v string)`

SetZonePool sets ZonePool field to given value.

### HasZonePool

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasZonePool() bool`

HasZonePool returns a boolean if a field has been set.

### SetZonePoolNil

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetZonePoolNil(b bool)`

 SetZonePoolNil sets the value for ZonePool to be an explicit nil

### UnsetZonePool
`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) UnsetZonePool()`

UnsetZonePool ensures that no value is present for ZonePool, not even an explicit nil
### GetStatus

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetGroupLayer

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetRules() []InlineResponse20095NetworkRouterFirewallRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) GetRulesOk() (*[]InlineResponse20095NetworkRouterFirewallRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) SetRules(v []InlineResponse20095NetworkRouterFirewallRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse20095NetworkRouterFirewallRuleGroups) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


