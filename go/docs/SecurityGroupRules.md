# SecurityGroupRules

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
**SourceGroup** | Pointer to **NullableString** |  | [optional] 
**SourceTier** | Pointer to **NullableString** |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to **NullableString** |  | [optional] 
**DestinationTier** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityGroupRules

`func NewSecurityGroupRules() *SecurityGroupRules`

NewSecurityGroupRules instantiates a new SecurityGroupRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRulesWithDefaults

`func NewSecurityGroupRulesWithDefaults() *SecurityGroupRules`

NewSecurityGroupRulesWithDefaults instantiates a new SecurityGroupRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupRules) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupRules) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupRules) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecurityGroupRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupRules) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityGroupRules) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityGroupRules) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *SecurityGroupRules) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *SecurityGroupRules) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *SecurityGroupRules) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *SecurityGroupRules) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *SecurityGroupRules) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *SecurityGroupRules) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *SecurityGroupRules) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *SecurityGroupRules) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *SecurityGroupRules) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *SecurityGroupRules) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *SecurityGroupRules) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *SecurityGroupRules) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *SecurityGroupRules) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *SecurityGroupRules) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *SecurityGroupRules) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRules) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRules) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityGroupRules) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *SecurityGroupRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SecurityGroupRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SecurityGroupRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SecurityGroupRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *SecurityGroupRules) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *SecurityGroupRules) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *SecurityGroupRules) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *SecurityGroupRules) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *SecurityGroupRules) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SecurityGroupRules) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SecurityGroupRules) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SecurityGroupRules) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *SecurityGroupRules) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *SecurityGroupRules) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *SecurityGroupRules) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *SecurityGroupRules) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *SecurityGroupRules) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *SecurityGroupRules) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *SecurityGroupRules) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *SecurityGroupRules) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *SecurityGroupRules) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *SecurityGroupRules) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *SecurityGroupRules) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *SecurityGroupRules) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *SecurityGroupRules) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *SecurityGroupRules) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetPortRange

`func (o *SecurityGroupRules) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *SecurityGroupRules) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *SecurityGroupRules) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *SecurityGroupRules) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *SecurityGroupRules) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *SecurityGroupRules) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetProtocol

`func (o *SecurityGroupRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityGroupRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *SecurityGroupRules) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *SecurityGroupRules) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *SecurityGroupRules) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *SecurityGroupRules) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *SecurityGroupRules) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *SecurityGroupRules) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *SecurityGroupRules) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *SecurityGroupRules) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *SecurityGroupRules) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *SecurityGroupRules) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *SecurityGroupRules) GetDestinationGroup() string`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *SecurityGroupRules) GetDestinationGroupOk() (*string, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *SecurityGroupRules) SetDestinationGroup(v string)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *SecurityGroupRules) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### SetDestinationGroupNil

`func (o *SecurityGroupRules) SetDestinationGroupNil(b bool)`

 SetDestinationGroupNil sets the value for DestinationGroup to be an explicit nil

### UnsetDestinationGroup
`func (o *SecurityGroupRules) UnsetDestinationGroup()`

UnsetDestinationGroup ensures that no value is present for DestinationGroup, not even an explicit nil
### GetDestinationTier

`func (o *SecurityGroupRules) GetDestinationTier() string`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *SecurityGroupRules) GetDestinationTierOk() (*string, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *SecurityGroupRules) SetDestinationTier(v string)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *SecurityGroupRules) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### SetDestinationTierNil

`func (o *SecurityGroupRules) SetDestinationTierNil(b bool)`

 SetDestinationTierNil sets the value for DestinationTier to be an explicit nil

### UnsetDestinationTier
`func (o *SecurityGroupRules) UnsetDestinationTier()`

UnsetDestinationTier ensures that no value is present for DestinationTier, not even an explicit nil
### GetExternalId

`func (o *SecurityGroupRules) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityGroupRules) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityGroupRules) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityGroupRules) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *SecurityGroupRules) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityGroupRules) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityGroupRules) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityGroupRules) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SecurityGroupRules) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SecurityGroupRules) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


