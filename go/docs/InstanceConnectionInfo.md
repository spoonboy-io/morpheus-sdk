# InstanceConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInstanceConnectionInfo

`func NewInstanceConnectionInfo() *InstanceConnectionInfo`

NewInstanceConnectionInfo instantiates a new InstanceConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConnectionInfoWithDefaults

`func NewInstanceConnectionInfoWithDefaults() *InstanceConnectionInfo`

NewInstanceConnectionInfoWithDefaults instantiates a new InstanceConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InstanceConnectionInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceConnectionInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceConnectionInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceConnectionInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *InstanceConnectionInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InstanceConnectionInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InstanceConnectionInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *InstanceConnectionInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *InstanceConnectionInfo) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *InstanceConnectionInfo) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


