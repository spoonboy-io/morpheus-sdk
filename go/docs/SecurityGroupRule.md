# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**CustomRule** | Pointer to **bool** |  | [optional] 
**InstanceTypeId** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceGroup** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**SourceTier** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**DestinationTier** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule() *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityGroupRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityGroupRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *SecurityGroupRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *SecurityGroupRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *SecurityGroupRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *SecurityGroupRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *SecurityGroupRule) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *SecurityGroupRule) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *SecurityGroupRule) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *SecurityGroupRule) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *SecurityGroupRule) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *SecurityGroupRule) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *SecurityGroupRule) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *SecurityGroupRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *SecurityGroupRule) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *SecurityGroupRule) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *SecurityGroupRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityGroupRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *SecurityGroupRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SecurityGroupRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SecurityGroupRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SecurityGroupRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *SecurityGroupRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SecurityGroupRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SecurityGroupRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *SecurityGroupRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *SecurityGroupRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SecurityGroupRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SecurityGroupRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SecurityGroupRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *SecurityGroupRule) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SecurityGroupRule) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *SecurityGroupRule) GetSourceGroup() InlineResponse20082LoadBalancerInstanceSslCert`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *SecurityGroupRule) GetSourceGroupOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *SecurityGroupRule) SetSourceGroup(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *SecurityGroupRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *SecurityGroupRule) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *SecurityGroupRule) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *SecurityGroupRule) GetSourceTier() InlineResponse20082LoadBalancerInstanceSslCert`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *SecurityGroupRule) GetSourceTierOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *SecurityGroupRule) SetSourceTier(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *SecurityGroupRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *SecurityGroupRule) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *SecurityGroupRule) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetPortRange

`func (o *SecurityGroupRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *SecurityGroupRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *SecurityGroupRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *SecurityGroupRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *SecurityGroupRule) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *SecurityGroupRule) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetProtocol

`func (o *SecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityGroupRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *SecurityGroupRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SecurityGroupRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SecurityGroupRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *SecurityGroupRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *SecurityGroupRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SecurityGroupRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SecurityGroupRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SecurityGroupRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *SecurityGroupRule) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *SecurityGroupRule) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *SecurityGroupRule) GetDestinationGroup() InlineResponse20082LoadBalancerInstanceSslCert`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *SecurityGroupRule) GetDestinationGroupOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *SecurityGroupRule) SetDestinationGroup(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *SecurityGroupRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### SetDestinationGroupNil

`func (o *SecurityGroupRule) SetDestinationGroupNil(b bool)`

 SetDestinationGroupNil sets the value for DestinationGroup to be an explicit nil

### UnsetDestinationGroup
`func (o *SecurityGroupRule) UnsetDestinationGroup()`

UnsetDestinationGroup ensures that no value is present for DestinationGroup, not even an explicit nil
### GetDestinationTier

`func (o *SecurityGroupRule) GetDestinationTier() InlineResponse20082LoadBalancerInstanceSslCert`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *SecurityGroupRule) GetDestinationTierOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *SecurityGroupRule) SetDestinationTier(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *SecurityGroupRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### SetDestinationTierNil

`func (o *SecurityGroupRule) SetDestinationTierNil(b bool)`

 SetDestinationTierNil sets the value for DestinationTier to be an explicit nil

### UnsetDestinationTier
`func (o *SecurityGroupRule) UnsetDestinationTier()`

UnsetDestinationTier ensures that no value is present for DestinationTier, not even an explicit nil
### GetExternalId

`func (o *SecurityGroupRule) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityGroupRule) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityGroupRule) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityGroupRule) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *SecurityGroupRule) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *SecurityGroupRule) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *SecurityGroupRule) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityGroupRule) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityGroupRule) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityGroupRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SecurityGroupRule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SecurityGroupRule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


