# InlineResponse20082LoadBalancerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancer**](inline_response_200_79_loadBalancerMonitor_loadBalancer.md) |  | [optional] 
**Instance** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Sticky** | Pointer to **bool** |  | [optional] 
**SslEnabled** | Pointer to **NullableString** |  | [optional] 
**ExternalAddress** | Pointer to **bool** |  | [optional] 
**BackendPort** | Pointer to **NullableString** |  | [optional] 
**VipType** | Pointer to **NullableString** |  | [optional] 
**VipAddress** | Pointer to **string** |  | [optional] 
**VipHostname** | Pointer to **NullableString** |  | [optional] 
**VipProtocol** | Pointer to **string** |  | [optional] 
**VipScheme** | Pointer to **NullableString** |  | [optional] 
**VipMode** | Pointer to **NullableString** |  | [optional] 
**VipName** | Pointer to **string** |  | [optional] 
**VipPort** | Pointer to **int64** |  | [optional] 
**VipSticky** | Pointer to **NullableString** |  | [optional] 
**VipBalance** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | Pointer to **NullableString** |  | [optional] 
**SourceAddress** | Pointer to **NullableString** |  | [optional] 
**SslCert** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**SslMode** | Pointer to **NullableString** |  | [optional] 
**SslRedirectMode** | Pointer to **NullableString** |  | [optional] 
**VipShared** | Pointer to **bool** |  | [optional] 
**VipDirectAddress** | Pointer to **NullableString** |  | [optional] 
**ServerName** | Pointer to **NullableString** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**Removing** | Pointer to **bool** |  | [optional] 
**VipSource** | Pointer to **string** |  | [optional] 
**ExtraConfig** | Pointer to **NullableString** |  | [optional] 
**ServiceAccess** | Pointer to **NullableString** |  | [optional] 
**NetworkId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**ExternalPortId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VipStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20082LoadBalancerInstance

`func NewInlineResponse20082LoadBalancerInstance() *InlineResponse20082LoadBalancerInstance`

NewInlineResponse20082LoadBalancerInstance instantiates a new InlineResponse20082LoadBalancerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082LoadBalancerInstanceWithDefaults

`func NewInlineResponse20082LoadBalancerInstanceWithDefaults() *InlineResponse20082LoadBalancerInstance`

NewInlineResponse20082LoadBalancerInstanceWithDefaults instantiates a new InlineResponse20082LoadBalancerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20082LoadBalancerInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20082LoadBalancerInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20082LoadBalancerInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *InlineResponse20082LoadBalancerInstance) GetLoadBalancer() InlineResponse20079LoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *InlineResponse20082LoadBalancerInstance) GetLoadBalancerOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *InlineResponse20082LoadBalancerInstance) SetLoadBalancer(v InlineResponse20079LoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *InlineResponse20082LoadBalancerInstance) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetInstance

`func (o *InlineResponse20082LoadBalancerInstance) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InlineResponse20082LoadBalancerInstance) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InlineResponse20082LoadBalancerInstance) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InlineResponse20082LoadBalancerInstance) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *InlineResponse20082LoadBalancerInstance) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *InlineResponse20082LoadBalancerInstance) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetDescription

`func (o *InlineResponse20082LoadBalancerInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20082LoadBalancerInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20082LoadBalancerInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20082LoadBalancerInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse20082LoadBalancerInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse20082LoadBalancerInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *InlineResponse20082LoadBalancerInstance) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20082LoadBalancerInstance) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20082LoadBalancerInstance) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20082LoadBalancerInstance) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20082LoadBalancerInstance) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDateCreated

`func (o *InlineResponse20082LoadBalancerInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20082LoadBalancerInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20082LoadBalancerInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20082LoadBalancerInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20082LoadBalancerInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20082LoadBalancerInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20082LoadBalancerInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20082LoadBalancerInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse20082LoadBalancerInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse20082LoadBalancerInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse20082LoadBalancerInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse20082LoadBalancerInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSticky

`func (o *InlineResponse20082LoadBalancerInstance) GetSticky() bool`

GetSticky returns the Sticky field if non-nil, zero value otherwise.

### GetStickyOk

`func (o *InlineResponse20082LoadBalancerInstance) GetStickyOk() (*bool, bool)`

GetStickyOk returns a tuple with the Sticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticky

`func (o *InlineResponse20082LoadBalancerInstance) SetSticky(v bool)`

SetSticky sets Sticky field to given value.

### HasSticky

`func (o *InlineResponse20082LoadBalancerInstance) HasSticky() bool`

HasSticky returns a boolean if a field has been set.

### GetSslEnabled

`func (o *InlineResponse20082LoadBalancerInstance) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *InlineResponse20082LoadBalancerInstance) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *InlineResponse20082LoadBalancerInstance) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetExternalAddress

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalAddress() bool`

GetExternalAddress returns the ExternalAddress field if non-nil, zero value otherwise.

### GetExternalAddressOk

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalAddressOk() (*bool, bool)`

GetExternalAddressOk returns a tuple with the ExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddress

`func (o *InlineResponse20082LoadBalancerInstance) SetExternalAddress(v bool)`

SetExternalAddress sets ExternalAddress field to given value.

### HasExternalAddress

`func (o *InlineResponse20082LoadBalancerInstance) HasExternalAddress() bool`

HasExternalAddress returns a boolean if a field has been set.

### GetBackendPort

`func (o *InlineResponse20082LoadBalancerInstance) GetBackendPort() string`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *InlineResponse20082LoadBalancerInstance) GetBackendPortOk() (*string, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *InlineResponse20082LoadBalancerInstance) SetBackendPort(v string)`

SetBackendPort sets BackendPort field to given value.

### HasBackendPort

`func (o *InlineResponse20082LoadBalancerInstance) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### SetBackendPortNil

`func (o *InlineResponse20082LoadBalancerInstance) SetBackendPortNil(b bool)`

 SetBackendPortNil sets the value for BackendPort to be an explicit nil

### UnsetBackendPort
`func (o *InlineResponse20082LoadBalancerInstance) UnsetBackendPort()`

UnsetBackendPort ensures that no value is present for BackendPort, not even an explicit nil
### GetVipType

`func (o *InlineResponse20082LoadBalancerInstance) GetVipType() string`

GetVipType returns the VipType field if non-nil, zero value otherwise.

### GetVipTypeOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipTypeOk() (*string, bool)`

GetVipTypeOk returns a tuple with the VipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipType

`func (o *InlineResponse20082LoadBalancerInstance) SetVipType(v string)`

SetVipType sets VipType field to given value.

### HasVipType

`func (o *InlineResponse20082LoadBalancerInstance) HasVipType() bool`

HasVipType returns a boolean if a field has been set.

### SetVipTypeNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipTypeNil(b bool)`

 SetVipTypeNil sets the value for VipType to be an explicit nil

### UnsetVipType
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipType()`

UnsetVipType ensures that no value is present for VipType, not even an explicit nil
### GetVipAddress

`func (o *InlineResponse20082LoadBalancerInstance) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *InlineResponse20082LoadBalancerInstance) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *InlineResponse20082LoadBalancerInstance) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipHostname

`func (o *InlineResponse20082LoadBalancerInstance) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *InlineResponse20082LoadBalancerInstance) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *InlineResponse20082LoadBalancerInstance) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### SetVipHostnameNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipHostnameNil(b bool)`

 SetVipHostnameNil sets the value for VipHostname to be an explicit nil

### UnsetVipHostname
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipHostname()`

UnsetVipHostname ensures that no value is present for VipHostname, not even an explicit nil
### GetVipProtocol

`func (o *InlineResponse20082LoadBalancerInstance) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *InlineResponse20082LoadBalancerInstance) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *InlineResponse20082LoadBalancerInstance) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipScheme

`func (o *InlineResponse20082LoadBalancerInstance) GetVipScheme() string`

GetVipScheme returns the VipScheme field if non-nil, zero value otherwise.

### GetVipSchemeOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipSchemeOk() (*string, bool)`

GetVipSchemeOk returns a tuple with the VipScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipScheme

`func (o *InlineResponse20082LoadBalancerInstance) SetVipScheme(v string)`

SetVipScheme sets VipScheme field to given value.

### HasVipScheme

`func (o *InlineResponse20082LoadBalancerInstance) HasVipScheme() bool`

HasVipScheme returns a boolean if a field has been set.

### SetVipSchemeNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipSchemeNil(b bool)`

 SetVipSchemeNil sets the value for VipScheme to be an explicit nil

### UnsetVipScheme
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipScheme()`

UnsetVipScheme ensures that no value is present for VipScheme, not even an explicit nil
### GetVipMode

`func (o *InlineResponse20082LoadBalancerInstance) GetVipMode() string`

GetVipMode returns the VipMode field if non-nil, zero value otherwise.

### GetVipModeOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipModeOk() (*string, bool)`

GetVipModeOk returns a tuple with the VipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipMode

`func (o *InlineResponse20082LoadBalancerInstance) SetVipMode(v string)`

SetVipMode sets VipMode field to given value.

### HasVipMode

`func (o *InlineResponse20082LoadBalancerInstance) HasVipMode() bool`

HasVipMode returns a boolean if a field has been set.

### SetVipModeNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipModeNil(b bool)`

 SetVipModeNil sets the value for VipMode to be an explicit nil

### UnsetVipMode
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipMode()`

UnsetVipMode ensures that no value is present for VipMode, not even an explicit nil
### GetVipName

`func (o *InlineResponse20082LoadBalancerInstance) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *InlineResponse20082LoadBalancerInstance) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *InlineResponse20082LoadBalancerInstance) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetVipPort

`func (o *InlineResponse20082LoadBalancerInstance) GetVipPort() int64`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipPortOk() (*int64, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *InlineResponse20082LoadBalancerInstance) SetVipPort(v int64)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *InlineResponse20082LoadBalancerInstance) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipSticky

`func (o *InlineResponse20082LoadBalancerInstance) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *InlineResponse20082LoadBalancerInstance) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *InlineResponse20082LoadBalancerInstance) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### SetVipStickyNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipStickyNil(b bool)`

 SetVipStickyNil sets the value for VipSticky to be an explicit nil

### UnsetVipSticky
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipSticky()`

UnsetVipSticky ensures that no value is present for VipSticky, not even an explicit nil
### GetVipBalance

`func (o *InlineResponse20082LoadBalancerInstance) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *InlineResponse20082LoadBalancerInstance) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *InlineResponse20082LoadBalancerInstance) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### SetVipBalanceNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipBalanceNil(b bool)`

 SetVipBalanceNil sets the value for VipBalance to be an explicit nil

### UnsetVipBalance
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipBalance()`

UnsetVipBalance ensures that no value is present for VipBalance, not even an explicit nil
### GetServicePort

`func (o *InlineResponse20082LoadBalancerInstance) GetServicePort() string`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *InlineResponse20082LoadBalancerInstance) GetServicePortOk() (*string, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *InlineResponse20082LoadBalancerInstance) SetServicePort(v string)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *InlineResponse20082LoadBalancerInstance) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *InlineResponse20082LoadBalancerInstance) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *InlineResponse20082LoadBalancerInstance) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetSourceAddress

`func (o *InlineResponse20082LoadBalancerInstance) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *InlineResponse20082LoadBalancerInstance) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *InlineResponse20082LoadBalancerInstance) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### SetSourceAddressNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetSslCert

`func (o *InlineResponse20082LoadBalancerInstance) GetSslCert() InlineResponse20082LoadBalancerInstanceSslCert`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSslCertOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *InlineResponse20082LoadBalancerInstance) SetSslCert(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *InlineResponse20082LoadBalancerInstance) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetSslMode

`func (o *InlineResponse20082LoadBalancerInstance) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *InlineResponse20082LoadBalancerInstance) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *InlineResponse20082LoadBalancerInstance) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### SetSslModeNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSslModeNil(b bool)`

 SetSslModeNil sets the value for SslMode to be an explicit nil

### UnsetSslMode
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSslMode()`

UnsetSslMode ensures that no value is present for SslMode, not even an explicit nil
### GetSslRedirectMode

`func (o *InlineResponse20082LoadBalancerInstance) GetSslRedirectMode() string`

GetSslRedirectMode returns the SslRedirectMode field if non-nil, zero value otherwise.

### GetSslRedirectModeOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSslRedirectModeOk() (*string, bool)`

GetSslRedirectModeOk returns a tuple with the SslRedirectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRedirectMode

`func (o *InlineResponse20082LoadBalancerInstance) SetSslRedirectMode(v string)`

SetSslRedirectMode sets SslRedirectMode field to given value.

### HasSslRedirectMode

`func (o *InlineResponse20082LoadBalancerInstance) HasSslRedirectMode() bool`

HasSslRedirectMode returns a boolean if a field has been set.

### SetSslRedirectModeNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSslRedirectModeNil(b bool)`

 SetSslRedirectModeNil sets the value for SslRedirectMode to be an explicit nil

### UnsetSslRedirectMode
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSslRedirectMode()`

UnsetSslRedirectMode ensures that no value is present for SslRedirectMode, not even an explicit nil
### GetVipShared

`func (o *InlineResponse20082LoadBalancerInstance) GetVipShared() bool`

GetVipShared returns the VipShared field if non-nil, zero value otherwise.

### GetVipSharedOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipSharedOk() (*bool, bool)`

GetVipSharedOk returns a tuple with the VipShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipShared

`func (o *InlineResponse20082LoadBalancerInstance) SetVipShared(v bool)`

SetVipShared sets VipShared field to given value.

### HasVipShared

`func (o *InlineResponse20082LoadBalancerInstance) HasVipShared() bool`

HasVipShared returns a boolean if a field has been set.

### GetVipDirectAddress

`func (o *InlineResponse20082LoadBalancerInstance) GetVipDirectAddress() string`

GetVipDirectAddress returns the VipDirectAddress field if non-nil, zero value otherwise.

### GetVipDirectAddressOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipDirectAddressOk() (*string, bool)`

GetVipDirectAddressOk returns a tuple with the VipDirectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipDirectAddress

`func (o *InlineResponse20082LoadBalancerInstance) SetVipDirectAddress(v string)`

SetVipDirectAddress sets VipDirectAddress field to given value.

### HasVipDirectAddress

`func (o *InlineResponse20082LoadBalancerInstance) HasVipDirectAddress() bool`

HasVipDirectAddress returns a boolean if a field has been set.

### SetVipDirectAddressNil

`func (o *InlineResponse20082LoadBalancerInstance) SetVipDirectAddressNil(b bool)`

 SetVipDirectAddressNil sets the value for VipDirectAddress to be an explicit nil

### UnsetVipDirectAddress
`func (o *InlineResponse20082LoadBalancerInstance) UnsetVipDirectAddress()`

UnsetVipDirectAddress ensures that no value is present for VipDirectAddress, not even an explicit nil
### GetServerName

`func (o *InlineResponse20082LoadBalancerInstance) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *InlineResponse20082LoadBalancerInstance) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *InlineResponse20082LoadBalancerInstance) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *InlineResponse20082LoadBalancerInstance) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *InlineResponse20082LoadBalancerInstance) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *InlineResponse20082LoadBalancerInstance) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetPoolName

`func (o *InlineResponse20082LoadBalancerInstance) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *InlineResponse20082LoadBalancerInstance) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *InlineResponse20082LoadBalancerInstance) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *InlineResponse20082LoadBalancerInstance) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *InlineResponse20082LoadBalancerInstance) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *InlineResponse20082LoadBalancerInstance) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetRemoving

`func (o *InlineResponse20082LoadBalancerInstance) GetRemoving() bool`

GetRemoving returns the Removing field if non-nil, zero value otherwise.

### GetRemovingOk

`func (o *InlineResponse20082LoadBalancerInstance) GetRemovingOk() (*bool, bool)`

GetRemovingOk returns a tuple with the Removing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoving

`func (o *InlineResponse20082LoadBalancerInstance) SetRemoving(v bool)`

SetRemoving sets Removing field to given value.

### HasRemoving

`func (o *InlineResponse20082LoadBalancerInstance) HasRemoving() bool`

HasRemoving returns a boolean if a field has been set.

### GetVipSource

`func (o *InlineResponse20082LoadBalancerInstance) GetVipSource() string`

GetVipSource returns the VipSource field if non-nil, zero value otherwise.

### GetVipSourceOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipSourceOk() (*string, bool)`

GetVipSourceOk returns a tuple with the VipSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSource

`func (o *InlineResponse20082LoadBalancerInstance) SetVipSource(v string)`

SetVipSource sets VipSource field to given value.

### HasVipSource

`func (o *InlineResponse20082LoadBalancerInstance) HasVipSource() bool`

HasVipSource returns a boolean if a field has been set.

### GetExtraConfig

`func (o *InlineResponse20082LoadBalancerInstance) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *InlineResponse20082LoadBalancerInstance) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *InlineResponse20082LoadBalancerInstance) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *InlineResponse20082LoadBalancerInstance) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *InlineResponse20082LoadBalancerInstance) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *InlineResponse20082LoadBalancerInstance) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetServiceAccess

`func (o *InlineResponse20082LoadBalancerInstance) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *InlineResponse20082LoadBalancerInstance) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *InlineResponse20082LoadBalancerInstance) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *InlineResponse20082LoadBalancerInstance) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### SetServiceAccessNil

`func (o *InlineResponse20082LoadBalancerInstance) SetServiceAccessNil(b bool)`

 SetServiceAccessNil sets the value for ServiceAccess to be an explicit nil

### UnsetServiceAccess
`func (o *InlineResponse20082LoadBalancerInstance) UnsetServiceAccess()`

UnsetServiceAccess ensures that no value is present for ServiceAccess, not even an explicit nil
### GetNetworkId

`func (o *InlineResponse20082LoadBalancerInstance) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20082LoadBalancerInstance) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *InlineResponse20082LoadBalancerInstance) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *InlineResponse20082LoadBalancerInstance) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *InlineResponse20082LoadBalancerInstance) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetSubnetId

`func (o *InlineResponse20082LoadBalancerInstance) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *InlineResponse20082LoadBalancerInstance) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *InlineResponse20082LoadBalancerInstance) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *InlineResponse20082LoadBalancerInstance) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *InlineResponse20082LoadBalancerInstance) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetExternalPortId

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *InlineResponse20082LoadBalancerInstance) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *InlineResponse20082LoadBalancerInstance) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *InlineResponse20082LoadBalancerInstance) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### SetExternalPortIdNil

`func (o *InlineResponse20082LoadBalancerInstance) SetExternalPortIdNil(b bool)`

 SetExternalPortIdNil sets the value for ExternalPortId to be an explicit nil

### UnsetExternalPortId
`func (o *InlineResponse20082LoadBalancerInstance) UnsetExternalPortId()`

UnsetExternalPortId ensures that no value is present for ExternalPortId, not even an explicit nil
### GetStatus

`func (o *InlineResponse20082LoadBalancerInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20082LoadBalancerInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20082LoadBalancerInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20082LoadBalancerInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVipStatus

`func (o *InlineResponse20082LoadBalancerInstance) GetVipStatus() string`

GetVipStatus returns the VipStatus field if non-nil, zero value otherwise.

### GetVipStatusOk

`func (o *InlineResponse20082LoadBalancerInstance) GetVipStatusOk() (*string, bool)`

GetVipStatusOk returns a tuple with the VipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipStatus

`func (o *InlineResponse20082LoadBalancerInstance) SetVipStatus(v string)`

SetVipStatus sets VipStatus field to given value.

### HasVipStatus

`func (o *InlineResponse20082LoadBalancerInstance) HasVipStatus() bool`

HasVipStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


