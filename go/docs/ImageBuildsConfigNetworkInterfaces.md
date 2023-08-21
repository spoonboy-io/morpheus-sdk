# ImageBuildsConfigNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpMode** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**ShowNetworkPoolLabel** | Pointer to **bool** |  | [optional] 
**ShowNetworkDhcpLabel** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**ImageBuildsConfigNetwork**](imageBuilds_config_network.md) |  | [optional] 
**NetworkInterfaceTypeId** | Pointer to **int64** |  | [optional] 
**NetworkInterfaceTypeIdName** | Pointer to **string** |  | [optional] 

## Methods

### NewImageBuildsConfigNetworkInterfaces

`func NewImageBuildsConfigNetworkInterfaces() *ImageBuildsConfigNetworkInterfaces`

NewImageBuildsConfigNetworkInterfaces instantiates a new ImageBuildsConfigNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildsConfigNetworkInterfacesWithDefaults

`func NewImageBuildsConfigNetworkInterfacesWithDefaults() *ImageBuildsConfigNetworkInterfaces`

NewImageBuildsConfigNetworkInterfacesWithDefaults instantiates a new ImageBuildsConfigNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpMode

`func (o *ImageBuildsConfigNetworkInterfaces) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *ImageBuildsConfigNetworkInterfaces) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *ImageBuildsConfigNetworkInterfaces) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *ImageBuildsConfigNetworkInterfaces) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ImageBuildsConfigNetworkInterfaces) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ImageBuildsConfigNetworkInterfaces) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetShowNetworkPoolLabel

`func (o *ImageBuildsConfigNetworkInterfaces) GetShowNetworkPoolLabel() bool`

GetShowNetworkPoolLabel returns the ShowNetworkPoolLabel field if non-nil, zero value otherwise.

### GetShowNetworkPoolLabelOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetShowNetworkPoolLabelOk() (*bool, bool)`

GetShowNetworkPoolLabelOk returns a tuple with the ShowNetworkPoolLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkPoolLabel

`func (o *ImageBuildsConfigNetworkInterfaces) SetShowNetworkPoolLabel(v bool)`

SetShowNetworkPoolLabel sets ShowNetworkPoolLabel field to given value.

### HasShowNetworkPoolLabel

`func (o *ImageBuildsConfigNetworkInterfaces) HasShowNetworkPoolLabel() bool`

HasShowNetworkPoolLabel returns a boolean if a field has been set.

### GetShowNetworkDhcpLabel

`func (o *ImageBuildsConfigNetworkInterfaces) GetShowNetworkDhcpLabel() bool`

GetShowNetworkDhcpLabel returns the ShowNetworkDhcpLabel field if non-nil, zero value otherwise.

### GetShowNetworkDhcpLabelOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetShowNetworkDhcpLabelOk() (*bool, bool)`

GetShowNetworkDhcpLabelOk returns a tuple with the ShowNetworkDhcpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkDhcpLabel

`func (o *ImageBuildsConfigNetworkInterfaces) SetShowNetworkDhcpLabel(v bool)`

SetShowNetworkDhcpLabel sets ShowNetworkDhcpLabel field to given value.

### HasShowNetworkDhcpLabel

`func (o *ImageBuildsConfigNetworkInterfaces) HasShowNetworkDhcpLabel() bool`

HasShowNetworkDhcpLabel returns a boolean if a field has been set.

### GetNetwork

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetwork() ImageBuildsConfigNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetworkOk() (*ImageBuildsConfigNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ImageBuildsConfigNetworkInterfaces) SetNetwork(v ImageBuildsConfigNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ImageBuildsConfigNetworkInterfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkInterfaceTypeId

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *ImageBuildsConfigNetworkInterfaces) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *ImageBuildsConfigNetworkInterfaces) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetNetworkInterfaceTypeIdName

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetworkInterfaceTypeIdName() string`

GetNetworkInterfaceTypeIdName returns the NetworkInterfaceTypeIdName field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdNameOk

`func (o *ImageBuildsConfigNetworkInterfaces) GetNetworkInterfaceTypeIdNameOk() (*string, bool)`

GetNetworkInterfaceTypeIdNameOk returns a tuple with the NetworkInterfaceTypeIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeIdName

`func (o *ImageBuildsConfigNetworkInterfaces) SetNetworkInterfaceTypeIdName(v string)`

SetNetworkInterfaceTypeIdName sets NetworkInterfaceTypeIdName field to given value.

### HasNetworkInterfaceTypeIdName

`func (o *ImageBuildsConfigNetworkInterfaces) HasNetworkInterfaceTypeIdName() bool`

HasNetworkInterfaceTypeIdName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


