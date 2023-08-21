# NetworkInterfaceUpdateSuccessServerInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Addresses** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**InterfaceId** | Pointer to **NullableString** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**NetworkPool** | Pointer to **map[string]interface{}** |  | [optional] 
**Dhcp** | Pointer to **bool** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**ReplaceHostRecord** | Pointer to **bool** |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**IpSubnet** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **NullableString** |  | [optional] 
**PublicIpAddress** | Pointer to **string** |  | [optional] 
**FabricId** | Pointer to **NullableString** |  | [optional] 
**Ipv6Subnet** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**PublicIpv6Address** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**NetworkGroup** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**NetworkPoolIPv6** | Pointer to **map[string]interface{}** |  | [optional] 
**Network** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**VlanId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**NetworkPosition** | Pointer to **NullableString** |  | [optional] 
**PoolAssigned** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkInterfaceUpdateSuccessServerInterfaces

`func NewNetworkInterfaceUpdateSuccessServerInterfaces() *NetworkInterfaceUpdateSuccessServerInterfaces`

NewNetworkInterfaceUpdateSuccessServerInterfaces instantiates a new NetworkInterfaceUpdateSuccessServerInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceUpdateSuccessServerInterfacesWithDefaults

`func NewNetworkInterfaceUpdateSuccessServerInterfacesWithDefaults() *NetworkInterfaceUpdateSuccessServerInterfaces`

NewNetworkInterfaceUpdateSuccessServerInterfacesWithDefaults instantiates a new NetworkInterfaceUpdateSuccessServerInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetAddresses() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetAddressesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetAddresses(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### SetInterfaceIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetInterfaceIdNil(b bool)`

 SetInterfaceIdNil sets the value for InterfaceId to be an explicit nil

### UnsetInterfaceId
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetInterfaceId()`

UnsetInterfaceId ensures that no value is present for InterfaceId, not even an explicit nil
### GetDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPool() map[string]interface{}`

GetNetworkPool returns the NetworkPool field if non-nil, zero value otherwise.

### GetNetworkPoolOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPoolOk() (*map[string]interface{}, bool)`

GetNetworkPoolOk returns a tuple with the NetworkPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPool(v map[string]interface{})`

SetNetworkPool sets NetworkPool field to given value.

### HasNetworkPool

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetworkPool() bool`

HasNetworkPool returns a boolean if a field has been set.

### SetNetworkPoolNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPoolNil(b bool)`

 SetNetworkPoolNil sets the value for NetworkPool to be an explicit nil

### UnsetNetworkPool
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetNetworkPool()`

UnsetNetworkPool ensures that no value is present for NetworkPool, not even an explicit nil
### GetDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDhcp() bool`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDhcpOk() (*bool, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetDhcp(v bool)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetReplaceHostRecord() bool`

GetReplaceHostRecord returns the ReplaceHostRecord field if non-nil, zero value otherwise.

### GetReplaceHostRecordOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetReplaceHostRecordOk() (*bool, bool)`

GetReplaceHostRecordOk returns a tuple with the ReplaceHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetReplaceHostRecord(v bool)`

SetReplaceHostRecord sets ReplaceHostRecord field to given value.

### HasReplaceHostRecord

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasReplaceHostRecord() bool`

HasReplaceHostRecord returns a boolean if a field has been set.

### GetIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpSubnet() string`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpSubnetOk() (*string, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetIpSubnet(v string)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### SetIpSubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetIpSubnetNil(b bool)`

 SetIpSubnetNil sets the value for IpSubnet to be an explicit nil

### UnsetIpSubnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetIpSubnet()`

UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
### GetConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPublicIpAddress() string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPublicIpAddressOk() (*string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetPublicIpAddress(v string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### GetFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### SetFabricIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetFabricIdNil(b bool)`

 SetFabricIdNil sets the value for FabricId to be an explicit nil

### UnsetFabricId
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetFabricId()`

UnsetFabricId ensures that no value is present for FabricId, not even an explicit nil
### GetIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpv6Subnet() string`

GetIpv6Subnet returns the Ipv6Subnet field if non-nil, zero value otherwise.

### GetIpv6SubnetOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetIpv6SubnetOk() (*string, bool)`

GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetIpv6Subnet(v string)`

SetIpv6Subnet sets Ipv6Subnet field to given value.

### HasIpv6Subnet

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasIpv6Subnet() bool`

HasIpv6Subnet returns a boolean if a field has been set.

### SetIpv6SubnetNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetIpv6SubnetNil(b bool)`

 SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil

### UnsetIpv6Subnet
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetIpv6Subnet()`

UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
### GetMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPublicIpv6Address() string`

GetPublicIpv6Address returns the PublicIpv6Address field if non-nil, zero value otherwise.

### GetPublicIpv6AddressOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPublicIpv6AddressOk() (*string, bool)`

GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetPublicIpv6Address(v string)`

SetPublicIpv6Address sets PublicIpv6Address field to given value.

### HasPublicIpv6Address

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasPublicIpv6Address() bool`

HasPublicIpv6Address returns a boolean if a field has been set.

### SetPublicIpv6AddressNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetPublicIpv6AddressNil(b bool)`

 SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil

### UnsetPublicIpv6Address
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetPublicIpv6Address()`

UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
### GetRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkGroup() string`

GetNetworkGroup returns the NetworkGroup field if non-nil, zero value otherwise.

### GetNetworkGroupOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkGroupOk() (*string, bool)`

GetNetworkGroupOk returns a tuple with the NetworkGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkGroup(v string)`

SetNetworkGroup sets NetworkGroup field to given value.

### HasNetworkGroup

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetworkGroup() bool`

HasNetworkGroup returns a boolean if a field has been set.

### SetNetworkGroupNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkGroupNil(b bool)`

 SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil

### UnsetNetworkGroup
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetNetworkGroup()`

UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
### GetRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetName

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPoolIPv6() map[string]interface{}`

GetNetworkPoolIPv6 returns the NetworkPoolIPv6 field if non-nil, zero value otherwise.

### GetNetworkPoolIPv6Ok

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPoolIPv6Ok() (*map[string]interface{}, bool)`

GetNetworkPoolIPv6Ok returns a tuple with the NetworkPoolIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPoolIPv6(v map[string]interface{})`

SetNetworkPoolIPv6 sets NetworkPoolIPv6 field to given value.

### HasNetworkPoolIPv6

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetworkPoolIPv6() bool`

HasNetworkPoolIPv6 returns a boolean if a field has been set.

### SetNetworkPoolIPv6Nil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPoolIPv6Nil(b bool)`

 SetNetworkPoolIPv6Nil sets the value for NetworkPoolIPv6 to be an explicit nil

### UnsetNetworkPoolIPv6
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetNetworkPoolIPv6()`

UnsetNetworkPoolIPv6 ensures that no value is present for NetworkPoolIPv6, not even an explicit nil
### GetNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetwork() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetwork(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### SetVlanIdNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetVlanIdNil(b bool)`

 SetVlanIdNil sets the value for VlanId to be an explicit nil

### UnsetVlanId
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetVlanId()`

UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
### GetType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetType() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetTypeOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetType(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPosition() string`

GetNetworkPosition returns the NetworkPosition field if non-nil, zero value otherwise.

### GetNetworkPositionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetNetworkPositionOk() (*string, bool)`

GetNetworkPositionOk returns a tuple with the NetworkPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPosition(v string)`

SetNetworkPosition sets NetworkPosition field to given value.

### HasNetworkPosition

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasNetworkPosition() bool`

HasNetworkPosition returns a boolean if a field has been set.

### SetNetworkPositionNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetNetworkPositionNil(b bool)`

 SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil

### UnsetNetworkPosition
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetNetworkPosition()`

UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
### GetPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPoolAssigned() bool`

GetPoolAssigned returns the PoolAssigned field if non-nil, zero value otherwise.

### GetPoolAssignedOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetPoolAssignedOk() (*bool, bool)`

GetPoolAssignedOk returns a tuple with the PoolAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetPoolAssigned(v bool)`

SetPoolAssigned sets PoolAssigned field to given value.

### HasPoolAssigned

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasPoolAssigned() bool`

HasPoolAssigned returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkInterfaceUpdateSuccessServerInterfaces) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


