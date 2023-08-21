# NetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanIDs** | Pointer to **NullableString** |  | [optional] 
**ConnectedGateway** | Pointer to **string** |  | [optional] 
**SubnetIpManagementType** | Pointer to **string** |  | [optional] 
**SubnetIpServerId** | Pointer to **string** |  | [optional] 
**DhcpRange** | Pointer to **string** |  | [optional] 
**SubnetDhcpServerAddress** | Pointer to **string** |  | [optional] 
**SubnetDhcpLeaseTime** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkConfig

`func NewNetworkConfig() *NetworkConfig`

NewNetworkConfig instantiates a new NetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigWithDefaults

`func NewNetworkConfigWithDefaults() *NetworkConfig`

NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanIDs

`func (o *NetworkConfig) GetVlanIDs() string`

GetVlanIDs returns the VlanIDs field if non-nil, zero value otherwise.

### GetVlanIDsOk

`func (o *NetworkConfig) GetVlanIDsOk() (*string, bool)`

GetVlanIDsOk returns a tuple with the VlanIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIDs

`func (o *NetworkConfig) SetVlanIDs(v string)`

SetVlanIDs sets VlanIDs field to given value.

### HasVlanIDs

`func (o *NetworkConfig) HasVlanIDs() bool`

HasVlanIDs returns a boolean if a field has been set.

### SetVlanIDsNil

`func (o *NetworkConfig) SetVlanIDsNil(b bool)`

 SetVlanIDsNil sets the value for VlanIDs to be an explicit nil

### UnsetVlanIDs
`func (o *NetworkConfig) UnsetVlanIDs()`

UnsetVlanIDs ensures that no value is present for VlanIDs, not even an explicit nil
### GetConnectedGateway

`func (o *NetworkConfig) GetConnectedGateway() string`

GetConnectedGateway returns the ConnectedGateway field if non-nil, zero value otherwise.

### GetConnectedGatewayOk

`func (o *NetworkConfig) GetConnectedGatewayOk() (*string, bool)`

GetConnectedGatewayOk returns a tuple with the ConnectedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGateway

`func (o *NetworkConfig) SetConnectedGateway(v string)`

SetConnectedGateway sets ConnectedGateway field to given value.

### HasConnectedGateway

`func (o *NetworkConfig) HasConnectedGateway() bool`

HasConnectedGateway returns a boolean if a field has been set.

### GetSubnetIpManagementType

`func (o *NetworkConfig) GetSubnetIpManagementType() string`

GetSubnetIpManagementType returns the SubnetIpManagementType field if non-nil, zero value otherwise.

### GetSubnetIpManagementTypeOk

`func (o *NetworkConfig) GetSubnetIpManagementTypeOk() (*string, bool)`

GetSubnetIpManagementTypeOk returns a tuple with the SubnetIpManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpManagementType

`func (o *NetworkConfig) SetSubnetIpManagementType(v string)`

SetSubnetIpManagementType sets SubnetIpManagementType field to given value.

### HasSubnetIpManagementType

`func (o *NetworkConfig) HasSubnetIpManagementType() bool`

HasSubnetIpManagementType returns a boolean if a field has been set.

### GetSubnetIpServerId

`func (o *NetworkConfig) GetSubnetIpServerId() string`

GetSubnetIpServerId returns the SubnetIpServerId field if non-nil, zero value otherwise.

### GetSubnetIpServerIdOk

`func (o *NetworkConfig) GetSubnetIpServerIdOk() (*string, bool)`

GetSubnetIpServerIdOk returns a tuple with the SubnetIpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpServerId

`func (o *NetworkConfig) SetSubnetIpServerId(v string)`

SetSubnetIpServerId sets SubnetIpServerId field to given value.

### HasSubnetIpServerId

`func (o *NetworkConfig) HasSubnetIpServerId() bool`

HasSubnetIpServerId returns a boolean if a field has been set.

### GetDhcpRange

`func (o *NetworkConfig) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *NetworkConfig) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *NetworkConfig) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *NetworkConfig) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### GetSubnetDhcpServerAddress

`func (o *NetworkConfig) GetSubnetDhcpServerAddress() string`

GetSubnetDhcpServerAddress returns the SubnetDhcpServerAddress field if non-nil, zero value otherwise.

### GetSubnetDhcpServerAddressOk

`func (o *NetworkConfig) GetSubnetDhcpServerAddressOk() (*string, bool)`

GetSubnetDhcpServerAddressOk returns a tuple with the SubnetDhcpServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpServerAddress

`func (o *NetworkConfig) SetSubnetDhcpServerAddress(v string)`

SetSubnetDhcpServerAddress sets SubnetDhcpServerAddress field to given value.

### HasSubnetDhcpServerAddress

`func (o *NetworkConfig) HasSubnetDhcpServerAddress() bool`

HasSubnetDhcpServerAddress returns a boolean if a field has been set.

### GetSubnetDhcpLeaseTime

`func (o *NetworkConfig) GetSubnetDhcpLeaseTime() string`

GetSubnetDhcpLeaseTime returns the SubnetDhcpLeaseTime field if non-nil, zero value otherwise.

### GetSubnetDhcpLeaseTimeOk

`func (o *NetworkConfig) GetSubnetDhcpLeaseTimeOk() (*string, bool)`

GetSubnetDhcpLeaseTimeOk returns a tuple with the SubnetDhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpLeaseTime

`func (o *NetworkConfig) SetSubnetDhcpLeaseTime(v string)`

SetSubnetDhcpLeaseTime sets SubnetDhcpLeaseTime field to given value.

### HasSubnetDhcpLeaseTime

`func (o *NetworkConfig) HasSubnetDhcpLeaseTime() bool`

HasSubnetDhcpLeaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


