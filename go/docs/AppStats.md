# AppStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**Running** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**InstanceCount** | Pointer to **int64** |  | [optional] 
**InstanceDayCount** | Pointer to **[]int64** |  | [optional] 
**InstanceDayCountTotal** | Pointer to **int64** |  | [optional] 

## Methods

### NewAppStats

`func NewAppStats() *AppStats`

NewAppStats instantiates a new AppStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStatsWithDefaults

`func NewAppStatsWithDefaults() *AppStats`

NewAppStatsWithDefaults instantiates a new AppStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedMemory

`func (o *AppStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *AppStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *AppStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *AppStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AppStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AppStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AppStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AppStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedStorage

`func (o *AppStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *AppStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *AppStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *AppStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AppStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AppStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AppStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AppStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetRunning

`func (o *AppStats) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *AppStats) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *AppStats) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *AppStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTotal

`func (o *AppStats) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AppStats) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AppStats) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AppStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCpuUsage

`func (o *AppStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *AppStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *AppStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *AppStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetInstanceCount

`func (o *AppStats) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *AppStats) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *AppStats) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *AppStats) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceDayCount

`func (o *AppStats) GetInstanceDayCount() []int64`

GetInstanceDayCount returns the InstanceDayCount field if non-nil, zero value otherwise.

### GetInstanceDayCountOk

`func (o *AppStats) GetInstanceDayCountOk() (*[]int64, bool)`

GetInstanceDayCountOk returns a tuple with the InstanceDayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCount

`func (o *AppStats) SetInstanceDayCount(v []int64)`

SetInstanceDayCount sets InstanceDayCount field to given value.

### HasInstanceDayCount

`func (o *AppStats) HasInstanceDayCount() bool`

HasInstanceDayCount returns a boolean if a field has been set.

### GetInstanceDayCountTotal

`func (o *AppStats) GetInstanceDayCountTotal() int64`

GetInstanceDayCountTotal returns the InstanceDayCountTotal field if non-nil, zero value otherwise.

### GetInstanceDayCountTotalOk

`func (o *AppStats) GetInstanceDayCountTotalOk() (*int64, bool)`

GetInstanceDayCountTotalOk returns a tuple with the InstanceDayCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceDayCountTotal

`func (o *AppStats) SetInstanceDayCountTotal(v int64)`

SetInstanceDayCountTotal sets InstanceDayCountTotal field to given value.

### HasInstanceDayCountTotal

`func (o *AppStats) HasInstanceDayCountTotal() bool`

HasInstanceDayCountTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


