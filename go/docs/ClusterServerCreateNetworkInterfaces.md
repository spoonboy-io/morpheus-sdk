# ClusterServerCreateNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**ClusterServerCreateNetwork**](clusterServerCreate_network.md) |  | 
**NetworkInterfaceTypeId** | Pointer to **int64** | The id of type of the network interface. | [optional] 
**IpAddress** | Pointer to **string** | The ip address. Not applicable when using DHCP or IP Pools. | [optional] 

## Methods

### NewClusterServerCreateNetworkInterfaces

`func NewClusterServerCreateNetworkInterfaces(network ClusterServerCreateNetwork, ) *ClusterServerCreateNetworkInterfaces`

NewClusterServerCreateNetworkInterfaces instantiates a new ClusterServerCreateNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateNetworkInterfacesWithDefaults

`func NewClusterServerCreateNetworkInterfacesWithDefaults() *ClusterServerCreateNetworkInterfaces`

NewClusterServerCreateNetworkInterfacesWithDefaults instantiates a new ClusterServerCreateNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ClusterServerCreateNetworkInterfaces) GetNetwork() ClusterServerCreateNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterServerCreateNetworkInterfaces) GetNetworkOk() (*ClusterServerCreateNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterServerCreateNetworkInterfaces) SetNetwork(v ClusterServerCreateNetwork)`

SetNetwork sets Network field to given value.


### GetNetworkInterfaceTypeId

`func (o *ClusterServerCreateNetworkInterfaces) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *ClusterServerCreateNetworkInterfaces) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *ClusterServerCreateNetworkInterfaces) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *ClusterServerCreateNetworkInterfaces) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ClusterServerCreateNetworkInterfaces) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ClusterServerCreateNetworkInterfaces) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ClusterServerCreateNetworkInterfaces) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ClusterServerCreateNetworkInterfaces) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


