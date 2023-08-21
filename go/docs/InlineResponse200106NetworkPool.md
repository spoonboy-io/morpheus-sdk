# InlineResponse200106NetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**InlineResponse20094Network**](inline_response_200_94_network.md) |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DnsDomain** | Pointer to **NullableString** |  | [optional] 
**DnsSearchPath** | Pointer to **NullableString** |  | [optional] 
**HostPrefix** | Pointer to **NullableString** |  | [optional] 
**HttpProxy** | Pointer to **NullableString** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**DnsSuffixList** | Pointer to **[]string** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Netmask** | Pointer to **NullableString** |  | [optional] 
**SubnetAddress** | Pointer to **NullableString** |  | [optional] 
**IpCount** | Pointer to **int64** |  | [optional] 
**FreeCount** | Pointer to **int64** |  | [optional] 
**PoolEnabled** | Pointer to **bool** |  | [optional] 
**TftpServer** | Pointer to **NullableString** |  | [optional] 
**BootFile** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**ParentType** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**PoolGroup** | Pointer to **NullableString** |  | [optional] 
**IpRanges** | Pointer to [**[]InlineResponse200106NetworkPoolIpRanges**](InlineResponse200106NetworkPoolIpRanges.md) |  | [optional] 

## Methods

### NewInlineResponse200106NetworkPool

`func NewInlineResponse200106NetworkPool() *InlineResponse200106NetworkPool`

NewInlineResponse200106NetworkPool instantiates a new InlineResponse200106NetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200106NetworkPoolWithDefaults

`func NewInlineResponse200106NetworkPoolWithDefaults() *InlineResponse200106NetworkPool`

NewInlineResponse200106NetworkPoolWithDefaults instantiates a new InlineResponse200106NetworkPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200106NetworkPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200106NetworkPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200106NetworkPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200106NetworkPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse200106NetworkPool) GetType() InlineResponse20094Network`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200106NetworkPool) GetTypeOk() (*InlineResponse20094Network, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200106NetworkPool) SetType(v InlineResponse20094Network)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse200106NetworkPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *InlineResponse200106NetworkPool) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InlineResponse200106NetworkPool) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InlineResponse200106NetworkPool) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InlineResponse200106NetworkPool) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *InlineResponse200106NetworkPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse200106NetworkPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse200106NetworkPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse200106NetworkPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InlineResponse200106NetworkPool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InlineResponse200106NetworkPool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCode

`func (o *InlineResponse200106NetworkPool) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse200106NetworkPool) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse200106NetworkPool) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse200106NetworkPool) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InlineResponse200106NetworkPool) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InlineResponse200106NetworkPool) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *InlineResponse200106NetworkPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200106NetworkPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200106NetworkPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200106NetworkPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineResponse200106NetworkPool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineResponse200106NetworkPool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineResponse200106NetworkPool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineResponse200106NetworkPool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *InlineResponse200106NetworkPool) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *InlineResponse200106NetworkPool) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInternalId

`func (o *InlineResponse200106NetworkPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse200106NetworkPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse200106NetworkPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse200106NetworkPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InlineResponse200106NetworkPool) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InlineResponse200106NetworkPool) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *InlineResponse200106NetworkPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse200106NetworkPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse200106NetworkPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse200106NetworkPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InlineResponse200106NetworkPool) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InlineResponse200106NetworkPool) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDnsDomain

`func (o *InlineResponse200106NetworkPool) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *InlineResponse200106NetworkPool) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *InlineResponse200106NetworkPool) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *InlineResponse200106NetworkPool) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### SetDnsDomainNil

`func (o *InlineResponse200106NetworkPool) SetDnsDomainNil(b bool)`

 SetDnsDomainNil sets the value for DnsDomain to be an explicit nil

### UnsetDnsDomain
`func (o *InlineResponse200106NetworkPool) UnsetDnsDomain()`

UnsetDnsDomain ensures that no value is present for DnsDomain, not even an explicit nil
### GetDnsSearchPath

`func (o *InlineResponse200106NetworkPool) GetDnsSearchPath() string`

GetDnsSearchPath returns the DnsSearchPath field if non-nil, zero value otherwise.

### GetDnsSearchPathOk

`func (o *InlineResponse200106NetworkPool) GetDnsSearchPathOk() (*string, bool)`

GetDnsSearchPathOk returns a tuple with the DnsSearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSearchPath

`func (o *InlineResponse200106NetworkPool) SetDnsSearchPath(v string)`

SetDnsSearchPath sets DnsSearchPath field to given value.

### HasDnsSearchPath

`func (o *InlineResponse200106NetworkPool) HasDnsSearchPath() bool`

HasDnsSearchPath returns a boolean if a field has been set.

### SetDnsSearchPathNil

`func (o *InlineResponse200106NetworkPool) SetDnsSearchPathNil(b bool)`

 SetDnsSearchPathNil sets the value for DnsSearchPath to be an explicit nil

### UnsetDnsSearchPath
`func (o *InlineResponse200106NetworkPool) UnsetDnsSearchPath()`

UnsetDnsSearchPath ensures that no value is present for DnsSearchPath, not even an explicit nil
### GetHostPrefix

`func (o *InlineResponse200106NetworkPool) GetHostPrefix() string`

GetHostPrefix returns the HostPrefix field if non-nil, zero value otherwise.

### GetHostPrefixOk

`func (o *InlineResponse200106NetworkPool) GetHostPrefixOk() (*string, bool)`

GetHostPrefixOk returns a tuple with the HostPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPrefix

`func (o *InlineResponse200106NetworkPool) SetHostPrefix(v string)`

SetHostPrefix sets HostPrefix field to given value.

### HasHostPrefix

`func (o *InlineResponse200106NetworkPool) HasHostPrefix() bool`

HasHostPrefix returns a boolean if a field has been set.

### SetHostPrefixNil

`func (o *InlineResponse200106NetworkPool) SetHostPrefixNil(b bool)`

 SetHostPrefixNil sets the value for HostPrefix to be an explicit nil

### UnsetHostPrefix
`func (o *InlineResponse200106NetworkPool) UnsetHostPrefix()`

UnsetHostPrefix ensures that no value is present for HostPrefix, not even an explicit nil
### GetHttpProxy

`func (o *InlineResponse200106NetworkPool) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *InlineResponse200106NetworkPool) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *InlineResponse200106NetworkPool) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *InlineResponse200106NetworkPool) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *InlineResponse200106NetworkPool) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *InlineResponse200106NetworkPool) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetDnsServers

`func (o *InlineResponse200106NetworkPool) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *InlineResponse200106NetworkPool) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *InlineResponse200106NetworkPool) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *InlineResponse200106NetworkPool) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffixList

`func (o *InlineResponse200106NetworkPool) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *InlineResponse200106NetworkPool) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *InlineResponse200106NetworkPool) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *InlineResponse200106NetworkPool) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### GetDhcpServer

`func (o *InlineResponse200106NetworkPool) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *InlineResponse200106NetworkPool) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *InlineResponse200106NetworkPool) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *InlineResponse200106NetworkPool) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *InlineResponse200106NetworkPool) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *InlineResponse200106NetworkPool) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *InlineResponse200106NetworkPool) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *InlineResponse200106NetworkPool) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *InlineResponse200106NetworkPool) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *InlineResponse200106NetworkPool) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetGateway

`func (o *InlineResponse200106NetworkPool) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *InlineResponse200106NetworkPool) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *InlineResponse200106NetworkPool) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *InlineResponse200106NetworkPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *InlineResponse200106NetworkPool) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *InlineResponse200106NetworkPool) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *InlineResponse200106NetworkPool) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *InlineResponse200106NetworkPool) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *InlineResponse200106NetworkPool) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *InlineResponse200106NetworkPool) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *InlineResponse200106NetworkPool) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *InlineResponse200106NetworkPool) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetSubnetAddress

`func (o *InlineResponse200106NetworkPool) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *InlineResponse200106NetworkPool) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *InlineResponse200106NetworkPool) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *InlineResponse200106NetworkPool) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### SetSubnetAddressNil

`func (o *InlineResponse200106NetworkPool) SetSubnetAddressNil(b bool)`

 SetSubnetAddressNil sets the value for SubnetAddress to be an explicit nil

### UnsetSubnetAddress
`func (o *InlineResponse200106NetworkPool) UnsetSubnetAddress()`

UnsetSubnetAddress ensures that no value is present for SubnetAddress, not even an explicit nil
### GetIpCount

`func (o *InlineResponse200106NetworkPool) GetIpCount() int64`

GetIpCount returns the IpCount field if non-nil, zero value otherwise.

### GetIpCountOk

`func (o *InlineResponse200106NetworkPool) GetIpCountOk() (*int64, bool)`

GetIpCountOk returns a tuple with the IpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCount

`func (o *InlineResponse200106NetworkPool) SetIpCount(v int64)`

SetIpCount sets IpCount field to given value.

### HasIpCount

`func (o *InlineResponse200106NetworkPool) HasIpCount() bool`

HasIpCount returns a boolean if a field has been set.

### GetFreeCount

`func (o *InlineResponse200106NetworkPool) GetFreeCount() int64`

GetFreeCount returns the FreeCount field if non-nil, zero value otherwise.

### GetFreeCountOk

`func (o *InlineResponse200106NetworkPool) GetFreeCountOk() (*int64, bool)`

GetFreeCountOk returns a tuple with the FreeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCount

`func (o *InlineResponse200106NetworkPool) SetFreeCount(v int64)`

SetFreeCount sets FreeCount field to given value.

### HasFreeCount

`func (o *InlineResponse200106NetworkPool) HasFreeCount() bool`

HasFreeCount returns a boolean if a field has been set.

### GetPoolEnabled

`func (o *InlineResponse200106NetworkPool) GetPoolEnabled() bool`

GetPoolEnabled returns the PoolEnabled field if non-nil, zero value otherwise.

### GetPoolEnabledOk

`func (o *InlineResponse200106NetworkPool) GetPoolEnabledOk() (*bool, bool)`

GetPoolEnabledOk returns a tuple with the PoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEnabled

`func (o *InlineResponse200106NetworkPool) SetPoolEnabled(v bool)`

SetPoolEnabled sets PoolEnabled field to given value.

### HasPoolEnabled

`func (o *InlineResponse200106NetworkPool) HasPoolEnabled() bool`

HasPoolEnabled returns a boolean if a field has been set.

### GetTftpServer

`func (o *InlineResponse200106NetworkPool) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *InlineResponse200106NetworkPool) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *InlineResponse200106NetworkPool) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *InlineResponse200106NetworkPool) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *InlineResponse200106NetworkPool) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *InlineResponse200106NetworkPool) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *InlineResponse200106NetworkPool) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *InlineResponse200106NetworkPool) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *InlineResponse200106NetworkPool) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *InlineResponse200106NetworkPool) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *InlineResponse200106NetworkPool) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *InlineResponse200106NetworkPool) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetRefType

`func (o *InlineResponse200106NetworkPool) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *InlineResponse200106NetworkPool) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *InlineResponse200106NetworkPool) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *InlineResponse200106NetworkPool) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *InlineResponse200106NetworkPool) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *InlineResponse200106NetworkPool) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *InlineResponse200106NetworkPool) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *InlineResponse200106NetworkPool) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *InlineResponse200106NetworkPool) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *InlineResponse200106NetworkPool) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *InlineResponse200106NetworkPool) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *InlineResponse200106NetworkPool) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetParentType

`func (o *InlineResponse200106NetworkPool) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *InlineResponse200106NetworkPool) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *InlineResponse200106NetworkPool) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *InlineResponse200106NetworkPool) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### SetParentTypeNil

`func (o *InlineResponse200106NetworkPool) SetParentTypeNil(b bool)`

 SetParentTypeNil sets the value for ParentType to be an explicit nil

### UnsetParentType
`func (o *InlineResponse200106NetworkPool) UnsetParentType()`

UnsetParentType ensures that no value is present for ParentType, not even an explicit nil
### GetParentId

`func (o *InlineResponse200106NetworkPool) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *InlineResponse200106NetworkPool) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *InlineResponse200106NetworkPool) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *InlineResponse200106NetworkPool) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *InlineResponse200106NetworkPool) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *InlineResponse200106NetworkPool) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPoolGroup

`func (o *InlineResponse200106NetworkPool) GetPoolGroup() string`

GetPoolGroup returns the PoolGroup field if non-nil, zero value otherwise.

### GetPoolGroupOk

`func (o *InlineResponse200106NetworkPool) GetPoolGroupOk() (*string, bool)`

GetPoolGroupOk returns a tuple with the PoolGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolGroup

`func (o *InlineResponse200106NetworkPool) SetPoolGroup(v string)`

SetPoolGroup sets PoolGroup field to given value.

### HasPoolGroup

`func (o *InlineResponse200106NetworkPool) HasPoolGroup() bool`

HasPoolGroup returns a boolean if a field has been set.

### SetPoolGroupNil

`func (o *InlineResponse200106NetworkPool) SetPoolGroupNil(b bool)`

 SetPoolGroupNil sets the value for PoolGroup to be an explicit nil

### UnsetPoolGroup
`func (o *InlineResponse200106NetworkPool) UnsetPoolGroup()`

UnsetPoolGroup ensures that no value is present for PoolGroup, not even an explicit nil
### GetIpRanges

`func (o *InlineResponse200106NetworkPool) GetIpRanges() []InlineResponse200106NetworkPoolIpRanges`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *InlineResponse200106NetworkPool) GetIpRangesOk() (*[]InlineResponse200106NetworkPoolIpRanges, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *InlineResponse200106NetworkPool) SetIpRanges(v []InlineResponse200106NetworkPoolIpRanges)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *InlineResponse200106NetworkPool) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


