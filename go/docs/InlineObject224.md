# InlineObject224

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ApiServersIdResizeServer**](_api_servers__id__resize_server.md) |  | [optional] 
**ServicePlanOptions** | Pointer to [**ApiServersIdResizeServicePlanOptions**](_api_servers__id__resize_servicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]InstanceCreateVolume**](InstanceCreateVolume.md) | List of volumes with their new sizes. | [optional] 
**DeleteOriginalVolumes** | Pointer to **bool** | Delete the original volumes after resizing. (Amazon only) | [optional] [default to false]
**NetworkInterfaces** | Pointer to [**[]InstanceCreateNetwork**](InstanceCreateNetwork.md) | Key for network configurations. Include id to update an existing interface. | [optional] 

## Methods

### NewInlineObject224

`func NewInlineObject224() *InlineObject224`

NewInlineObject224 instantiates a new InlineObject224 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject224WithDefaults

`func NewInlineObject224WithDefaults() *InlineObject224`

NewInlineObject224WithDefaults instantiates a new InlineObject224 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *InlineObject224) GetServer() ApiServersIdResizeServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InlineObject224) GetServerOk() (*ApiServersIdResizeServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InlineObject224) SetServer(v ApiServersIdResizeServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *InlineObject224) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InlineObject224) GetServicePlanOptions() ApiServersIdResizeServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InlineObject224) GetServicePlanOptionsOk() (*ApiServersIdResizeServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InlineObject224) SetServicePlanOptions(v ApiServersIdResizeServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InlineObject224) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *InlineObject224) GetVolumes() []InstanceCreateVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InlineObject224) GetVolumesOk() (*[]InstanceCreateVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InlineObject224) SetVolumes(v []InstanceCreateVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InlineObject224) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetDeleteOriginalVolumes

`func (o *InlineObject224) GetDeleteOriginalVolumes() bool`

GetDeleteOriginalVolumes returns the DeleteOriginalVolumes field if non-nil, zero value otherwise.

### GetDeleteOriginalVolumesOk

`func (o *InlineObject224) GetDeleteOriginalVolumesOk() (*bool, bool)`

GetDeleteOriginalVolumesOk returns a tuple with the DeleteOriginalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOriginalVolumes

`func (o *InlineObject224) SetDeleteOriginalVolumes(v bool)`

SetDeleteOriginalVolumes sets DeleteOriginalVolumes field to given value.

### HasDeleteOriginalVolumes

`func (o *InlineObject224) HasDeleteOriginalVolumes() bool`

HasDeleteOriginalVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InlineObject224) GetNetworkInterfaces() []InstanceCreateNetwork`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InlineObject224) GetNetworkInterfacesOk() (*[]InstanceCreateNetwork, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InlineObject224) SetNetworkInterfaces(v []InstanceCreateNetwork)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InlineObject224) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


