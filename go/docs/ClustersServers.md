# ClustersServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeSet** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ClustersComputeServerType**](clusters_computeServerType.md) |  | [optional] 

## Methods

### NewClustersServers

`func NewClustersServers() *ClustersServers`

NewClustersServers instantiates a new ClustersServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClustersServersWithDefaults

`func NewClustersServersWithDefaults() *ClustersServers`

NewClustersServersWithDefaults instantiates a new ClustersServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClustersServers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClustersServers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClustersServers) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClustersServers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClustersServers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClustersServers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClustersServers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClustersServers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeSet

`func (o *ClustersServers) GetTypeSet() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetTypeSet returns the TypeSet field if non-nil, zero value otherwise.

### GetTypeSetOk

`func (o *ClustersServers) GetTypeSetOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetTypeSetOk returns a tuple with the TypeSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSet

`func (o *ClustersServers) SetTypeSet(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetTypeSet sets TypeSet field to given value.

### HasTypeSet

`func (o *ClustersServers) HasTypeSet() bool`

HasTypeSet returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ClustersServers) GetComputeServerType() ClustersComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ClustersServers) GetComputeServerTypeOk() (*ClustersComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ClustersServers) SetComputeServerType(v ClustersComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ClustersServers) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


