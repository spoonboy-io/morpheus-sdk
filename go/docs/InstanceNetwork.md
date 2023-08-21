# InstanceNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Subnet** | Pointer to **NullableString** |  | [optional] 
**DhcpServer** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Pool** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 

## Methods

### NewInstanceNetwork

`func NewInstanceNetwork() *InstanceNetwork`

NewInstanceNetwork instantiates a new InstanceNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceNetworkWithDefaults

`func NewInstanceNetworkWithDefaults() *InstanceNetwork`

NewInstanceNetworkWithDefaults instantiates a new InstanceNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InstanceNetwork) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstanceNetwork) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetGroup

`func (o *InstanceNetwork) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceNetwork) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceNetwork) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InstanceNetwork) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *InstanceNetwork) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *InstanceNetwork) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetSubnet

`func (o *InstanceNetwork) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InstanceNetwork) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InstanceNetwork) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InstanceNetwork) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### SetSubnetNil

`func (o *InstanceNetwork) SetSubnetNil(b bool)`

 SetSubnetNil sets the value for Subnet to be an explicit nil

### UnsetSubnet
`func (o *InstanceNetwork) UnsetSubnet()`

UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
### GetDhcpServer

`func (o *InstanceNetwork) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *InstanceNetwork) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *InstanceNetwork) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *InstanceNetwork) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### SetDhcpServerNil

`func (o *InstanceNetwork) SetDhcpServerNil(b bool)`

 SetDhcpServerNil sets the value for DhcpServer to be an explicit nil

### UnsetDhcpServer
`func (o *InstanceNetwork) UnsetDhcpServer()`

UnsetDhcpServer ensures that no value is present for DhcpServer, not even an explicit nil
### GetName

`func (o *InstanceNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceNetwork) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceNetwork) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPool

`func (o *InstanceNetwork) GetPool() InlineResponse20082LoadBalancerInstanceSslCert`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *InstanceNetwork) GetPoolOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *InstanceNetwork) SetPool(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetPool sets Pool field to given value.

### HasPool

`func (o *InstanceNetwork) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *InstanceNetwork) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *InstanceNetwork) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


