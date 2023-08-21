# InstanceInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Network** | Pointer to [**InstanceNetwork**](instance_network.md) |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**NetworkInterfaceTypeId** | Pointer to **NullableInt64** |  | [optional] 
**IpMode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceInterfaces

`func NewInstanceInterfaces() *InstanceInterfaces`

NewInstanceInterfaces instantiates a new InstanceInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceInterfacesWithDefaults

`func NewInstanceInterfacesWithDefaults() *InstanceInterfaces`

NewInstanceInterfacesWithDefaults instantiates a new InstanceInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceInterfaces) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceInterfaces) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceInterfaces) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceInterfaces) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceInterfaces) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceInterfaces) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetNetwork

`func (o *InstanceInterfaces) GetNetwork() InstanceNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InstanceInterfaces) GetNetworkOk() (*InstanceNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InstanceInterfaces) SetNetwork(v InstanceNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InstanceInterfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIpAddress

`func (o *InstanceInterfaces) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *InstanceInterfaces) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *InstanceInterfaces) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *InstanceInterfaces) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *InstanceInterfaces) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *InstanceInterfaces) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetNetworkInterfaceTypeId

`func (o *InstanceInterfaces) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *InstanceInterfaces) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *InstanceInterfaces) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *InstanceInterfaces) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### SetNetworkInterfaceTypeIdNil

`func (o *InstanceInterfaces) SetNetworkInterfaceTypeIdNil(b bool)`

 SetNetworkInterfaceTypeIdNil sets the value for NetworkInterfaceTypeId to be an explicit nil

### UnsetNetworkInterfaceTypeId
`func (o *InstanceInterfaces) UnsetNetworkInterfaceTypeId()`

UnsetNetworkInterfaceTypeId ensures that no value is present for NetworkInterfaceTypeId, not even an explicit nil
### GetIpMode

`func (o *InstanceInterfaces) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *InstanceInterfaces) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *InstanceInterfaces) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *InstanceInterfaces) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### SetIpModeNil

`func (o *InstanceInterfaces) SetIpModeNil(b bool)`

 SetIpModeNil sets the value for IpMode to be an explicit nil

### UnsetIpMode
`func (o *InstanceInterfaces) UnsetIpMode()`

UnsetIpMode ensures that no value is present for IpMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


