# ImageBuildConfigNetworkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**ZoneVcenterConfigNetworkServer**](zoneVcenterConfig_networkServer.md) |  | [optional] 

## Methods

### NewImageBuildConfigNetworkInterfaces

`func NewImageBuildConfigNetworkInterfaces() *ImageBuildConfigNetworkInterfaces`

NewImageBuildConfigNetworkInterfaces instantiates a new ImageBuildConfigNetworkInterfaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildConfigNetworkInterfacesWithDefaults

`func NewImageBuildConfigNetworkInterfacesWithDefaults() *ImageBuildConfigNetworkInterfaces`

NewImageBuildConfigNetworkInterfacesWithDefaults instantiates a new ImageBuildConfigNetworkInterfaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryInterface

`func (o *ImageBuildConfigNetworkInterfaces) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ImageBuildConfigNetworkInterfaces) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ImageBuildConfigNetworkInterfaces) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ImageBuildConfigNetworkInterfaces) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetNetwork

`func (o *ImageBuildConfigNetworkInterfaces) GetNetwork() ZoneVcenterConfigNetworkServer`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ImageBuildConfigNetworkInterfaces) GetNetworkOk() (*ZoneVcenterConfigNetworkServer, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ImageBuildConfigNetworkInterfaces) SetNetwork(v ZoneVcenterConfigNetworkServer)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ImageBuildConfigNetworkInterfaces) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


