# NetworkFirewallRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroup** | Pointer to [**NetworkFirewallRuleCreateRuleGroup**](networkFirewallRuleCreate_ruleGroup.md) |  | [optional] 
**Name** | **string** | Network firewall rule name | 
**Description** | Pointer to **NullableString** | Network firewall rule description | [optional] 
**Enabled** | Pointer to **bool** | Use this to set enabled state | [optional] 
**Priority** | Pointer to **NullableString** | Network firewall rule priority | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**NetworkFirewallRuleCreateSources**](networkFirewallRuleCreate_sources.md) |  | [optional] 
**Destinations** | Pointer to [**NetworkFirewallRuleCreateSources**](networkFirewallRuleCreate_sources.md) |  | [optional] 
**Config** | Pointer to [**NetworkFirewallRuleCreateConfig**](networkFirewallRuleCreate_config.md) |  | [optional] 
**Scopes** | Pointer to [**NetworkFirewallRuleCreateSources**](networkFirewallRuleCreate_sources.md) |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkFirewallRuleCreate

`func NewNetworkFirewallRuleCreate(name string, ) *NetworkFirewallRuleCreate`

NewNetworkFirewallRuleCreate instantiates a new NetworkFirewallRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFirewallRuleCreateWithDefaults

`func NewNetworkFirewallRuleCreateWithDefaults() *NetworkFirewallRuleCreate`

NewNetworkFirewallRuleCreateWithDefaults instantiates a new NetworkFirewallRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroup

`func (o *NetworkFirewallRuleCreate) GetRuleGroup() NetworkFirewallRuleCreateRuleGroup`

GetRuleGroup returns the RuleGroup field if non-nil, zero value otherwise.

### GetRuleGroupOk

`func (o *NetworkFirewallRuleCreate) GetRuleGroupOk() (*NetworkFirewallRuleCreateRuleGroup, bool)`

GetRuleGroupOk returns a tuple with the RuleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroup

`func (o *NetworkFirewallRuleCreate) SetRuleGroup(v NetworkFirewallRuleCreateRuleGroup)`

SetRuleGroup sets RuleGroup field to given value.

### HasRuleGroup

`func (o *NetworkFirewallRuleCreate) HasRuleGroup() bool`

HasRuleGroup returns a boolean if a field has been set.

### GetName

`func (o *NetworkFirewallRuleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFirewallRuleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFirewallRuleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkFirewallRuleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFirewallRuleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFirewallRuleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFirewallRuleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkFirewallRuleCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkFirewallRuleCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *NetworkFirewallRuleCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkFirewallRuleCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkFirewallRuleCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkFirewallRuleCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkFirewallRuleCreate) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkFirewallRuleCreate) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkFirewallRuleCreate) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkFirewallRuleCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *NetworkFirewallRuleCreate) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *NetworkFirewallRuleCreate) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetDirection

`func (o *NetworkFirewallRuleCreate) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkFirewallRuleCreate) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkFirewallRuleCreate) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkFirewallRuleCreate) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSources

`func (o *NetworkFirewallRuleCreate) GetSources() NetworkFirewallRuleCreateSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *NetworkFirewallRuleCreate) GetSourcesOk() (*NetworkFirewallRuleCreateSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *NetworkFirewallRuleCreate) SetSources(v NetworkFirewallRuleCreateSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *NetworkFirewallRuleCreate) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetDestinations

`func (o *NetworkFirewallRuleCreate) GetDestinations() NetworkFirewallRuleCreateSources`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *NetworkFirewallRuleCreate) GetDestinationsOk() (*NetworkFirewallRuleCreateSources, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *NetworkFirewallRuleCreate) SetDestinations(v NetworkFirewallRuleCreateSources)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *NetworkFirewallRuleCreate) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkFirewallRuleCreate) GetConfig() NetworkFirewallRuleCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkFirewallRuleCreate) GetConfigOk() (*NetworkFirewallRuleCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkFirewallRuleCreate) SetConfig(v NetworkFirewallRuleCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkFirewallRuleCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetScopes

`func (o *NetworkFirewallRuleCreate) GetScopes() NetworkFirewallRuleCreateSources`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NetworkFirewallRuleCreate) GetScopesOk() (*NetworkFirewallRuleCreateSources, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NetworkFirewallRuleCreate) SetScopes(v NetworkFirewallRuleCreateSources)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NetworkFirewallRuleCreate) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPolicy

`func (o *NetworkFirewallRuleCreate) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *NetworkFirewallRuleCreate) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *NetworkFirewallRuleCreate) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *NetworkFirewallRuleCreate) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


