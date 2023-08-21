# VdiPoolConfigNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**VdiPoolConfigNetwork**](vdiPool_config_network.md) |  | [optional] 
**IpMode** | Pointer to **string** |  | [optional] 
**ShowNetworkPoolLabel** | Pointer to **bool** |  | [optional] 
**ShowNetworkDhcpLabel** | Pointer to **bool** |  | [optional] 

## Methods

### NewVdiPoolConfigNetworkInterfaces

`func NewVdiPoolConfigNetworkInterfaces() *VdiPoolConfigNetworkInterfaces`

NewVdiPoolConfigNetworkInterfaces instantiates a new VdiPoolConfigNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolConfigNetworkInterfacesWithDefaults

`func NewVdiPoolConfigNetworkInterfacesWithDefaults() *VdiPoolConfigNetworkInterfaces`

NewVdiPoolConfigNetworkInterfacesWithDefaults instantiates a new VdiPoolConfigNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryInterface

`func (o *VdiPoolConfigNetworkInterfaces) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *VdiPoolConfigNetworkInterfaces) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *VdiPoolConfigNetworkInterfaces) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *VdiPoolConfigNetworkInterfaces) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *VdiPoolConfigNetworkInterfaces) GetNetwork() VdiPoolConfigNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VdiPoolConfigNetworkInterfaces) GetNetworkOk() (*VdiPoolConfigNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VdiPoolConfigNetworkInterfaces) SetNetwork(v VdiPoolConfigNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VdiPoolConfigNetworkInterfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpMode

`func (o *VdiPoolConfigNetworkInterfaces) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *VdiPoolConfigNetworkInterfaces) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *VdiPoolConfigNetworkInterfaces) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *VdiPoolConfigNetworkInterfaces) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetShowNetworkPoolLabel

`func (o *VdiPoolConfigNetworkInterfaces) GetShowNetworkPoolLabel() bool`

GetShowNetworkPoolLabel returns the ShowNetworkPoolLabel field if non-nil, zero value otherwise.

### GetShowNetworkPoolLabelOk

`func (o *VdiPoolConfigNetworkInterfaces) GetShowNetworkPoolLabelOk() (*bool, bool)`

GetShowNetworkPoolLabelOk returns a tuple with the ShowNetworkPoolLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkPoolLabel

`func (o *VdiPoolConfigNetworkInterfaces) SetShowNetworkPoolLabel(v bool)`

SetShowNetworkPoolLabel sets ShowNetworkPoolLabel field to given value.

### HasShowNetworkPoolLabel

`func (o *VdiPoolConfigNetworkInterfaces) HasShowNetworkPoolLabel() bool`

HasShowNetworkPoolLabel returns a boolean if a field has been set.

### GetShowNetworkDhcpLabel

`func (o *VdiPoolConfigNetworkInterfaces) GetShowNetworkDhcpLabel() bool`

GetShowNetworkDhcpLabel returns the ShowNetworkDhcpLabel field if non-nil, zero value otherwise.

### GetShowNetworkDhcpLabelOk

`func (o *VdiPoolConfigNetworkInterfaces) GetShowNetworkDhcpLabelOk() (*bool, bool)`

GetShowNetworkDhcpLabelOk returns a tuple with the ShowNetworkDhcpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkDhcpLabel

`func (o *VdiPoolConfigNetworkInterfaces) SetShowNetworkDhcpLabel(v bool)`

SetShowNetworkDhcpLabel sets ShowNetworkDhcpLabel field to given value.

### HasShowNetworkDhcpLabel

`func (o *VdiPoolConfigNetworkInterfaces) HasShowNetworkDhcpLabel() bool`

HasShowNetworkDhcpLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


