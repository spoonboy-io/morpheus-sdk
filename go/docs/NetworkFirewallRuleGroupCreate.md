# NetworkFirewallRuleGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Network firewall rule group name | 
**Description** | Pointer to **string** | Network firewall rule group description | [optional] 
**Priority** | Pointer to **int64** | Network firewall rule group priority | [optional] 
**ExternalType** | **string** | Use SecurityPolicy | 
**GroupLayer** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkFirewallRuleGroupCreate

`func NewNetworkFirewallRuleGroupCreate(name string, externalType string, ) *NetworkFirewallRuleGroupCreate`

NewNetworkFirewallRuleGroupCreate instantiates a new NetworkFirewallRuleGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFirewallRuleGroupCreateWithDefaults

`func NewNetworkFirewallRuleGroupCreateWithDefaults() *NetworkFirewallRuleGroupCreate`

NewNetworkFirewallRuleGroupCreateWithDefaults instantiates a new NetworkFirewallRuleGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkFirewallRuleGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkFirewallRuleGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkFirewallRuleGroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkFirewallRuleGroupCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkFirewallRuleGroupCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkFirewallRuleGroupCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkFirewallRuleGroupCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkFirewallRuleGroupCreate) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkFirewallRuleGroupCreate) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkFirewallRuleGroupCreate) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkFirewallRuleGroupCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetExternalType

`func (o *NetworkFirewallRuleGroupCreate) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *NetworkFirewallRuleGroupCreate) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *NetworkFirewallRuleGroupCreate) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.


### GetGroupLayer

`func (o *NetworkFirewallRuleGroupCreate) GetGroupLayer() string`

GetGroupLayer returns the GroupLayer field if non-nil, zero value otherwise.

### GetGroupLayerOk

`func (o *NetworkFirewallRuleGroupCreate) GetGroupLayerOk() (*string, bool)`

GetGroupLayerOk returns a tuple with the GroupLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLayer

`func (o *NetworkFirewallRuleGroupCreate) SetGroupLayer(v string)`

SetGroupLayer sets GroupLayer field to given value.

### HasGroupLayer

`func (o *NetworkFirewallRuleGroupCreate) HasGroupLayer() bool`

HasGroupLayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


