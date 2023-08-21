# TenantStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | Pointer to **int64** |  | [optional] 
**UserCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewTenantStats

`func NewTenantStats() *TenantStats`

NewTenantStats instantiates a new TenantStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantStatsWithDefaults

`func NewTenantStatsWithDefaults() *TenantStats`

NewTenantStatsWithDefaults instantiates a new TenantStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *TenantStats) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *TenantStats) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *TenantStats) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *TenantStats) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetUserCount

`func (o *TenantStats) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *TenantStats) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *TenantStats) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *TenantStats) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


