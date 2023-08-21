# InlineObject223

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ApiServersIdMakeManagedServer**](_api_servers__id__make_managed_server.md) |  | [optional] 
**InstallAgent** | Pointer to **bool** | Install agent. Set to false to manually install agent instead. | [optional] [default to true]
**InstanceTypeId** | Pointer to **int64** | Instance Type ID for the new Instance | [optional] 
**Layout** | Pointer to **int64** | Layout ID for the new Instance | [optional] 

## Methods

### NewInlineObject223

`func NewInlineObject223() *InlineObject223`

NewInlineObject223 instantiates a new InlineObject223 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject223WithDefaults

`func NewInlineObject223WithDefaults() *InlineObject223`

NewInlineObject223WithDefaults instantiates a new InlineObject223 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *InlineObject223) GetServer() ApiServersIdMakeManagedServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *InlineObject223) GetServerOk() (*ApiServersIdMakeManagedServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *InlineObject223) SetServer(v ApiServersIdMakeManagedServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *InlineObject223) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInstallAgent

`func (o *InlineObject223) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *InlineObject223) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *InlineObject223) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *InlineObject223) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *InlineObject223) GetInstanceTypeId() int64`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *InlineObject223) GetInstanceTypeIdOk() (*int64, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *InlineObject223) SetInstanceTypeId(v int64)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *InlineObject223) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetLayout

`func (o *InlineObject223) GetLayout() int64`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InlineObject223) GetLayoutOk() (*int64, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InlineObject223) SetLayout(v int64)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *InlineObject223) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


