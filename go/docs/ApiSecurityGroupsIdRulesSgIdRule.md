# ApiSecurityGroupsIdRulesSgIdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the rule | [optional] 
**Direction** | Pointer to **string** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to "ingress"]
**SourceType** | Pointer to **string** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60;. | [optional] [default to "cidr"]
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

### NewApiSecurityGroupsIdRulesSgIdRule

`func NewApiSecurityGroupsIdRulesSgIdRule(protocol string, ruleType string, ) *ApiSecurityGroupsIdRulesSgIdRule`

NewApiSecurityGroupsIdRulesSgIdRule instantiates a new ApiSecurityGroupsIdRulesSgIdRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityGroupsIdRulesSgIdRuleWithDefaults

`func NewApiSecurityGroupsIdRulesSgIdRuleWithDefaults() *ApiSecurityGroupsIdRulesSgIdRule`

NewApiSecurityGroupsIdRulesSgIdRuleWithDefaults instantiates a new ApiSecurityGroupsIdRulesSgIdRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceGroup() ApiSecurityGroupsIdRulesRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceGroupOk() (*ApiSecurityGroupsIdRulesRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetSourceGroup(v ApiSecurityGroupsIdRulesRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceTier() ApiSecurityGroupsIdRulesRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetSourceTierOk() (*ApiSecurityGroupsIdRulesRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetSourceTier(v ApiSecurityGroupsIdRulesRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestinationType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationGroup() ApiSecurityGroupsIdRulesRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationGroupOk() (*ApiSecurityGroupsIdRulesRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetDestinationGroup(v ApiSecurityGroupsIdRulesRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationTier() ApiSecurityGroupsIdRulesRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetDestinationTierOk() (*ApiSecurityGroupsIdRulesRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetDestinationTier(v ApiSecurityGroupsIdRulesRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetRuleType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetPolicy

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetInstanceTypeId() int64`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *ApiSecurityGroupsIdRulesSgIdRule) GetInstanceTypeIdOk() (*int64, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesSgIdRule) SetInstanceTypeId(v int64)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *ApiSecurityGroupsIdRulesSgIdRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


