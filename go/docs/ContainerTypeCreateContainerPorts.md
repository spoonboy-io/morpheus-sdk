# ContainerTypeCreateContainerPorts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Port** | **int64** |  | 
**LoadBalanceProtocol** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerTypeCreateContainerPorts

`func NewContainerTypeCreateContainerPorts(name string, port int64, ) *ContainerTypeCreateContainerPorts`

NewContainerTypeCreateContainerPorts instantiates a new ContainerTypeCreateContainerPorts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerTypeCreateContainerPortsWithDefaults

`func NewContainerTypeCreateContainerPortsWithDefaults() *ContainerTypeCreateContainerPorts`

NewContainerTypeCreateContainerPortsWithDefaults instantiates a new ContainerTypeCreateContainerPorts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerTypeCreateContainerPorts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTypeCreateContainerPorts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTypeCreateContainerPorts) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *ContainerTypeCreateContainerPorts) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerTypeCreateContainerPorts) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerTypeCreateContainerPorts) SetPort(v int64)`

SetPort sets Port field to given value.


### GetLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPorts) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *ContainerTypeCreateContainerPorts) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPorts) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPorts) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


