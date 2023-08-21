# ContainerTypeContainerPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**LoadBalanceProtocol** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContainerTypeContainerPorts

`func NewContainerTypeContainerPorts() *ContainerTypeContainerPorts`

NewContainerTypeContainerPorts instantiates a new ContainerTypeContainerPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTypeContainerPortsWithDefaults

`func NewContainerTypeContainerPortsWithDefaults() *ContainerTypeContainerPorts`

NewContainerTypeContainerPortsWithDefaults instantiates a new ContainerTypeContainerPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerTypeContainerPorts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTypeContainerPorts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTypeContainerPorts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerTypeContainerPorts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ContainerTypeContainerPorts) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerTypeContainerPorts) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerTypeContainerPorts) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ContainerTypeContainerPorts) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLoadBalanceProtocol

`func (o *ContainerTypeContainerPorts) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *ContainerTypeContainerPorts) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *ContainerTypeContainerPorts) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *ContainerTypeContainerPorts) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.

### SetLoadBalanceProtocolNil

`func (o *ContainerTypeContainerPorts) SetLoadBalanceProtocolNil(b bool)`

 SetLoadBalanceProtocolNil sets the value for LoadBalanceProtocol to be an explicit nil

### UnsetLoadBalanceProtocol
`func (o *ContainerTypeContainerPorts) UnsetLoadBalanceProtocol()`

UnsetLoadBalanceProtocol ensures that no value is present for LoadBalanceProtocol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


