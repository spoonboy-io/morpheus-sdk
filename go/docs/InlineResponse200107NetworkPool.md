# InlineResponse200107NetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**NetworkPoolId** | Pointer to **int64** |  | [optional] 
**IpType** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**GatewayAddress** | Pointer to **NullableString** |  | [optional] 
**SubnetMask** | Pointer to **NullableString** |  | [optional] 
**DnsServer** | Pointer to **NullableString** |  | [optional] 
**InterfaceName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**StaticIp** | Pointer to **bool** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**PtrId** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**SubRefId** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**InlineResponse200107NetworkPoolCreatedBy**](inline_response_200_107_networkPool_createdBy.md) |  | [optional] 

## Methods

### NewInlineResponse200107NetworkPool

`func NewInlineResponse200107NetworkPool() *InlineResponse200107NetworkPool`

NewInlineResponse200107NetworkPool instantiates a new InlineResponse200107NetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200107NetworkPoolWithDefaults

`func NewInlineResponse200107NetworkPoolWithDefaults() *InlineResponse200107NetworkPool`

NewInlineResponse200107NetworkPoolWithDefaults instantiates a new InlineResponse200107NetworkPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200107NetworkPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200107NetworkPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200107NetworkPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200107NetworkPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkPoolId

`func (o *InlineResponse200107NetworkPool) GetNetworkPoolId() int64`

GetNetworkPoolId returns the NetworkPoolId field if non-nil, zero value otherwise.

### GetNetworkPoolIdOk

`func (o *InlineResponse200107NetworkPool) GetNetworkPoolIdOk() (*int64, bool)`

GetNetworkPoolIdOk returns a tuple with the NetworkPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolId

`func (o *InlineResponse200107NetworkPool) SetNetworkPoolId(v int64)`

SetNetworkPoolId sets NetworkPoolId field to given value.

### HasNetworkPoolId

`func (o *InlineResponse200107NetworkPool) HasNetworkPoolId() bool`

HasNetworkPoolId returns a boolean if a field has been set.

### GetIpType

`func (o *InlineResponse200107NetworkPool) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *InlineResponse200107NetworkPool) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *InlineResponse200107NetworkPool) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *InlineResponse200107NetworkPool) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpAddress

`func (o *InlineResponse200107NetworkPool) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InlineResponse200107NetworkPool) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InlineResponse200107NetworkPool) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InlineResponse200107NetworkPool) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetGatewayAddress

`func (o *InlineResponse200107NetworkPool) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *InlineResponse200107NetworkPool) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *InlineResponse200107NetworkPool) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *InlineResponse200107NetworkPool) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### SetGatewayAddressNil

`func (o *InlineResponse200107NetworkPool) SetGatewayAddressNil(b bool)`

 SetGatewayAddressNil sets the value for GatewayAddress to be an explicit nil

### UnsetGatewayAddress
`func (o *InlineResponse200107NetworkPool) UnsetGatewayAddress()`

UnsetGatewayAddress ensures that no value is present for GatewayAddress, not even an explicit nil
### GetSubnetMask

`func (o *InlineResponse200107NetworkPool) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *InlineResponse200107NetworkPool) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *InlineResponse200107NetworkPool) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *InlineResponse200107NetworkPool) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### SetSubnetMaskNil

`func (o *InlineResponse200107NetworkPool) SetSubnetMaskNil(b bool)`

 SetSubnetMaskNil sets the value for SubnetMask to be an explicit nil

### UnsetSubnetMask
`func (o *InlineResponse200107NetworkPool) UnsetSubnetMask()`

UnsetSubnetMask ensures that no value is present for SubnetMask, not even an explicit nil
### GetDnsServer

`func (o *InlineResponse200107NetworkPool) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *InlineResponse200107NetworkPool) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *InlineResponse200107NetworkPool) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *InlineResponse200107NetworkPool) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### SetDnsServerNil

`func (o *InlineResponse200107NetworkPool) SetDnsServerNil(b bool)`

 SetDnsServerNil sets the value for DnsServer to be an explicit nil

### UnsetDnsServer
`func (o *InlineResponse200107NetworkPool) UnsetDnsServer()`

UnsetDnsServer ensures that no value is present for DnsServer, not even an explicit nil
### GetInterfaceName

`func (o *InlineResponse200107NetworkPool) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InlineResponse200107NetworkPool) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InlineResponse200107NetworkPool) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InlineResponse200107NetworkPool) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### SetInterfaceNameNil

`func (o *InlineResponse200107NetworkPool) SetInterfaceNameNil(b bool)`

 SetInterfaceNameNil sets the value for InterfaceName to be an explicit nil

### UnsetInterfaceName
`func (o *InlineResponse200107NetworkPool) UnsetInterfaceName()`

UnsetInterfaceName ensures that no value is present for InterfaceName, not even an explicit nil
### GetDescription

`func (o *InlineResponse200107NetworkPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200107NetworkPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200107NetworkPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200107NetworkPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InlineResponse200107NetworkPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InlineResponse200107NetworkPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetActive

`func (o *InlineResponse200107NetworkPool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200107NetworkPool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200107NetworkPool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200107NetworkPool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStaticIp

`func (o *InlineResponse200107NetworkPool) GetStaticIp() bool`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *InlineResponse200107NetworkPool) GetStaticIpOk() (*bool, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *InlineResponse200107NetworkPool) SetStaticIp(v bool)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *InlineResponse200107NetworkPool) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### GetFqdn

`func (o *InlineResponse200107NetworkPool) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *InlineResponse200107NetworkPool) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *InlineResponse200107NetworkPool) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *InlineResponse200107NetworkPool) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetDomainName

`func (o *InlineResponse200107NetworkPool) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *InlineResponse200107NetworkPool) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *InlineResponse200107NetworkPool) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *InlineResponse200107NetworkPool) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *InlineResponse200107NetworkPool) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *InlineResponse200107NetworkPool) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetHostname

`func (o *InlineResponse200107NetworkPool) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *InlineResponse200107NetworkPool) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *InlineResponse200107NetworkPool) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *InlineResponse200107NetworkPool) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInternalId

`func (o *InlineResponse200107NetworkPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse200107NetworkPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse200107NetworkPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse200107NetworkPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InlineResponse200107NetworkPool) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InlineResponse200107NetworkPool) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *InlineResponse200107NetworkPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200107NetworkPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200107NetworkPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200107NetworkPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InlineResponse200107NetworkPool) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InlineResponse200107NetworkPool) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetPtrId

`func (o *InlineResponse200107NetworkPool) GetPtrId() string`

GetPtrId returns the PtrId field if non-nil, zero value otherwise.

### GetPtrIdOk

`func (o *InlineResponse200107NetworkPool) GetPtrIdOk() (*string, bool)`

GetPtrIdOk returns a tuple with the PtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrId

`func (o *InlineResponse200107NetworkPool) SetPtrId(v string)`

SetPtrId sets PtrId field to given value.

### HasPtrId

`func (o *InlineResponse200107NetworkPool) HasPtrId() bool`

HasPtrId returns a boolean if a field has been set.

### SetPtrIdNil

`func (o *InlineResponse200107NetworkPool) SetPtrIdNil(b bool)`

 SetPtrIdNil sets the value for PtrId to be an explicit nil

### UnsetPtrId
`func (o *InlineResponse200107NetworkPool) UnsetPtrId()`

UnsetPtrId ensures that no value is present for PtrId, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse200107NetworkPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse200107NetworkPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse200107NetworkPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse200107NetworkPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse200107NetworkPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse200107NetworkPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse200107NetworkPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse200107NetworkPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *InlineResponse200107NetworkPool) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InlineResponse200107NetworkPool) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InlineResponse200107NetworkPool) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InlineResponse200107NetworkPool) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InlineResponse200107NetworkPool) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InlineResponse200107NetworkPool) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InlineResponse200107NetworkPool) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InlineResponse200107NetworkPool) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRefType

`func (o *InlineResponse200107NetworkPool) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *InlineResponse200107NetworkPool) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *InlineResponse200107NetworkPool) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *InlineResponse200107NetworkPool) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *InlineResponse200107NetworkPool) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *InlineResponse200107NetworkPool) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *InlineResponse200107NetworkPool) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *InlineResponse200107NetworkPool) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *InlineResponse200107NetworkPool) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *InlineResponse200107NetworkPool) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *InlineResponse200107NetworkPool) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *InlineResponse200107NetworkPool) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetSubRefId

`func (o *InlineResponse200107NetworkPool) GetSubRefId() string`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *InlineResponse200107NetworkPool) GetSubRefIdOk() (*string, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *InlineResponse200107NetworkPool) SetSubRefId(v string)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *InlineResponse200107NetworkPool) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### SetSubRefIdNil

`func (o *InlineResponse200107NetworkPool) SetSubRefIdNil(b bool)`

 SetSubRefIdNil sets the value for SubRefId to be an explicit nil

### UnsetSubRefId
`func (o *InlineResponse200107NetworkPool) UnsetSubRefId()`

UnsetSubRefId ensures that no value is present for SubRefId, not even an explicit nil
### GetNetworkDomain

`func (o *InlineResponse200107NetworkPool) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *InlineResponse200107NetworkPool) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *InlineResponse200107NetworkPool) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *InlineResponse200107NetworkPool) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *InlineResponse200107NetworkPool) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *InlineResponse200107NetworkPool) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetCreatedBy

`func (o *InlineResponse200107NetworkPool) GetCreatedBy() InlineResponse200107NetworkPoolCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse200107NetworkPool) GetCreatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse200107NetworkPool) SetCreatedBy(v InlineResponse200107NetworkPoolCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse200107NetworkPool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


