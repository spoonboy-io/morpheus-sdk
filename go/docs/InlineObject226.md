# InlineObject226

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | Pointer to **int64** | The cloud/zone ID we are moving the set of servers to | [optional] 
**Servers** | Pointer to [**[]ApiServersChangeCloudServers**](ApiServersChangeCloudServers.md) | A JSON array of source: and target: server ids to be moved. If the target is blank Morpheus will automatically try to match by the servers unique or externalId | [optional] 

## Methods

### NewInlineObject226

`func NewInlineObject226() *InlineObject226`

NewInlineObject226 instantiates a new InlineObject226 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject226WithDefaults

`func NewInlineObject226WithDefaults() *InlineObject226`

NewInlineObject226WithDefaults instantiates a new InlineObject226 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *InlineObject226) GetCloudId() int64`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *InlineObject226) GetCloudIdOk() (*int64, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *InlineObject226) SetCloudId(v int64)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *InlineObject226) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetServers

`func (o *InlineObject226) GetServers() []ApiServersChangeCloudServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineObject226) GetServersOk() (*[]ApiServersChangeCloudServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineObject226) SetServers(v []ApiServersChangeCloudServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InlineObject226) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


