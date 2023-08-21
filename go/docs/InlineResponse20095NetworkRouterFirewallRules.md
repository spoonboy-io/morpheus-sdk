# InlineResponse20095NetworkRouterFirewallRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **[]string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **[]string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Application** | Pointer to **NullableString** |  | [optional] 
**ApplicationType** | Pointer to **string** |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**SourceGroup** | Pointer to **NullableString** |  | [optional] 
**SourceTier** | Pointer to **NullableString** |  | [optional] 
**Applications** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Methods

### NewInlineResponse20095NetworkRouterFirewallRules

`func NewInlineResponse20095NetworkRouterFirewallRules() *InlineResponse20095NetworkRouterFirewallRules`

NewInlineResponse20095NetworkRouterFirewallRules instantiates a new InlineResponse20095NetworkRouterFirewallRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095NetworkRouterFirewallRulesWithDefaults

`func NewInlineResponse20095NetworkRouterFirewallRulesWithDefaults() *InlineResponse20095NetworkRouterFirewallRules`

NewInlineResponse20095NetworkRouterFirewallRulesWithDefaults instantiates a new InlineResponse20095NetworkRouterFirewallRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetEnabled

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupName

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetDirection

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRuleType

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceType

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestination

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDestination() []string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDestinationOk() (*[]string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetDestination(v []string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetProfiles

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetProfiles() []string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetProfilesOk() (*[]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetProfiles(v []string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProtocol

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetApplication

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetApplicationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetPortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetSourceGroup

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *InlineResponse20095NetworkRouterFirewallRules) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetApplications

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplications() []InlineResponse20040AppDeployInstance`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse20095NetworkRouterFirewallRules) GetApplicationsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse20095NetworkRouterFirewallRules) SetApplications(v []InlineResponse20040AppDeployInstance)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse20095NetworkRouterFirewallRules) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


