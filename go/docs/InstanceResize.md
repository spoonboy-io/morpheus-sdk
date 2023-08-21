# InstanceResize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**InstanceResizeInstance**](instanceResize_instance.md) |  | [optional] 
**ServicePlanOptions** | Pointer to [**ServicePlanOptions**](servicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]InstanceCreateVolume**](InstanceCreateVolume.md) | Can be used to grow just the logical volume of the instance instead of choosing a plan | [optional] 
**DeleteOriginalVolumes** | Pointer to **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**NetworkInterfaces** | Pointer to [**[]InstanceCreateNetwork**](InstanceCreateNetwork.md) | Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. | [optional] 

## Methods

### NewInstanceResize

`func NewInstanceResize() *InstanceResize`

NewInstanceResize instantiates a new InstanceResize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResizeWithDefaults

`func NewInstanceResizeWithDefaults() *InstanceResize`

NewInstanceResizeWithDefaults instantiates a new InstanceResize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *InstanceResize) GetInstance() InstanceResizeInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceResize) GetInstanceOk() (*InstanceResizeInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceResize) SetInstance(v InstanceResizeInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceResize) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceResize) GetServicePlanOptions() ServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceResize) GetServicePlanOptionsOk() (*ServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceResize) SetServicePlanOptions(v ServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceResize) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *InstanceResize) GetVolumes() []InstanceCreateVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceResize) GetVolumesOk() (*[]InstanceCreateVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceResize) SetVolumes(v []InstanceCreateVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InstanceResize) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetDeleteOriginalVolumes

`func (o *InstanceResize) GetDeleteOriginalVolumes() bool`

GetDeleteOriginalVolumes returns the DeleteOriginalVolumes field if non-nil, zero value otherwise.

### GetDeleteOriginalVolumesOk

`func (o *InstanceResize) GetDeleteOriginalVolumesOk() (*bool, bool)`

GetDeleteOriginalVolumesOk returns a tuple with the DeleteOriginalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOriginalVolumes

`func (o *InstanceResize) SetDeleteOriginalVolumes(v bool)`

SetDeleteOriginalVolumes sets DeleteOriginalVolumes field to given value.

### HasDeleteOriginalVolumes

`func (o *InstanceResize) HasDeleteOriginalVolumes() bool`

HasDeleteOriginalVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InstanceResize) GetNetworkInterfaces() []InstanceCreateNetwork`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceResize) GetNetworkInterfacesOk() (*[]InstanceCreateNetwork, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceResize) SetNetworkInterfaces(v []InstanceCreateNetwork)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceResize) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


