# ApiSecurityGroupsIdRulesRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the rule | [optional] 
**Direction** | Pointer to **string** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to "ingress"]
**SourceType** | Pointer to **string** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60; | [optional] [default to "cidr"]
**Source** | Pointer to **string** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional] 
**SourceGroup** | Pointer to [**ApiSecurityGroupsIdRulesRuleSourceGroup**](_api_security_groups__id__rules_rule_sourceGroup.md) |  | [optional] 
**SourceTier** | Pointer to [**ApiSecurityGroupsIdRulesRuleSourceTier**](_api_security_groups__id__rules_rule_sourceTier.md) |  | [optional] 
**PortRange** | Pointer to **string** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  | [optional] 
**Protocol** | **string** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**DestinationType** | Pointer to **string** | Either cidr, group, tier, instance. | [optional] [default to "cidr"]
**Destination** | Pointer to **string** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional] 
**DestinationGroup** | Pointer to [**ApiSecurityGroupsIdRulesRuleDestinationGroup**](_api_security_groups__id__rules_rule_destinationGroup.md) |  | [optional] 
**DestinationTier** | Pointer to [**ApiSecurityGroupsIdRulesRuleDestinationTier**](_api_security_groups__id__rules_rule_destinationTier.md) |  | [optional] 
**RuleType** | **string** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | [default to "customRule"]
**Policy** | Pointer to **string** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional] 
**InstanceTypeId** | Pointer to **int64** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  | [optional] 

## Methods

### NewApiSecurityGroupsIdRulesRule

`func NewApiSecurityGroupsIdRulesRule(protocol string, ruleType string, ) *ApiSecurityGroupsIdRulesRule`

NewApiSecurityGroupsIdRulesRule instantiates a new ApiSecurityGroupsIdRulesRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityGroupsIdRulesRuleWithDefaults

`func NewApiSecurityGroupsIdRulesRuleWithDefaults() *ApiSecurityGroupsIdRulesRule`

NewApiSecurityGroupsIdRulesRuleWithDefaults instantiates a new ApiSecurityGroupsIdRulesRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityGroupsIdRulesRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityGroupsIdRulesRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityGroupsIdRulesRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiSecurityGroupsIdRulesRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *ApiSecurityGroupsIdRulesRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ApiSecurityGroupsIdRulesRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ApiSecurityGroupsIdRulesRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ApiSecurityGroupsIdRulesRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ApiSecurityGroupsIdRulesRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ApiSecurityGroupsIdRulesRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *ApiSecurityGroupsIdRulesRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiSecurityGroupsIdRulesRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiSecurityGroupsIdRulesRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceGroup

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceGroup() ApiSecurityGroupsIdRulesRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceGroupOk() (*ApiSecurityGroupsIdRulesRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *ApiSecurityGroupsIdRulesRule) SetSourceGroup(v ApiSecurityGroupsIdRulesRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *ApiSecurityGroupsIdRulesRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceTier() ApiSecurityGroupsIdRulesRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *ApiSecurityGroupsIdRulesRule) GetSourceTierOk() (*ApiSecurityGroupsIdRulesRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *ApiSecurityGroupsIdRulesRule) SetSourceTier(v ApiSecurityGroupsIdRulesRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *ApiSecurityGroupsIdRulesRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *ApiSecurityGroupsIdRulesRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ApiSecurityGroupsIdRulesRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ApiSecurityGroupsIdRulesRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ApiSecurityGroupsIdRulesRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *ApiSecurityGroupsIdRulesRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiSecurityGroupsIdRulesRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiSecurityGroupsIdRulesRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestinationType

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ApiSecurityGroupsIdRulesRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *ApiSecurityGroupsIdRulesRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *ApiSecurityGroupsIdRulesRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ApiSecurityGroupsIdRulesRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ApiSecurityGroupsIdRulesRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationGroup() ApiSecurityGroupsIdRulesRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationGroupOk() (*ApiSecurityGroupsIdRulesRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *ApiSecurityGroupsIdRulesRule) SetDestinationGroup(v ApiSecurityGroupsIdRulesRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *ApiSecurityGroupsIdRulesRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationTier() ApiSecurityGroupsIdRulesRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *ApiSecurityGroupsIdRulesRule) GetDestinationTierOk() (*ApiSecurityGroupsIdRulesRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *ApiSecurityGroupsIdRulesRule) SetDestinationTier(v ApiSecurityGroupsIdRulesRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *ApiSecurityGroupsIdRulesRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetRuleType

`func (o *ApiSecurityGroupsIdRulesRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ApiSecurityGroupsIdRulesRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ApiSecurityGroupsIdRulesRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetPolicy

`func (o *ApiSecurityGroupsIdRulesRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ApiSecurityGroupsIdRulesRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ApiSecurityGroupsIdRulesRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ApiSecurityGroupsIdRulesRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesRule) GetInstanceTypeId() int64`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *ApiSecurityGroupsIdRulesRule) GetInstanceTypeIdOk() (*int64, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesRule) SetInstanceTypeId(v int64)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


