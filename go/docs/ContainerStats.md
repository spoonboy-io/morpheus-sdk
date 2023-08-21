# ContainerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**UserCpuUsage** | Pointer to **float32** |  | [optional] 
**SystemCpuUsage** | Pointer to **float32** |  | [optional] 
**UsedMemory** | Pointer to **int32** |  | [optional] 
**MaxMemory** | Pointer to **int32** |  | [optional] 
**CacheMemory** | Pointer to **int32** |  | [optional] 
**MaxStorage** | Pointer to **int32** |  | [optional] 
**UsedStorage** | Pointer to **int32** |  | [optional] 
**ReadIOPS** | Pointer to **int32** |  | [optional] 
**WriteIOPS** | Pointer to **float32** |  | [optional] 
**TotalIOPS** | Pointer to **float32** |  | [optional] 
**NetTxUsage** | Pointer to **int32** |  | [optional] 
**NetRxUsage** | Pointer to **int32** |  | [optional] 

## Methods

### NewContainerStats

`func NewContainerStats() *ContainerStats`

NewContainerStats instantiates a new ContainerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStatsWithDefaults

`func NewContainerStatsWithDefaults() *ContainerStats`

NewContainerStatsWithDefaults instantiates a new ContainerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ContainerStats) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ContainerStats) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ContainerStats) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ContainerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetRunning

`func (o *ContainerStats) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ContainerStats) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ContainerStats) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ContainerStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUserCpuUsage

`func (o *ContainerStats) GetUserCpuUsage() float32`

GetUserCpuUsage returns the UserCpuUsage field if non-nil, zero value otherwise.

### GetUserCpuUsageOk

`func (o *ContainerStats) GetUserCpuUsageOk() (*float32, bool)`

GetUserCpuUsageOk returns a tuple with the UserCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCpuUsage

`func (o *ContainerStats) SetUserCpuUsage(v float32)`

SetUserCpuUsage sets UserCpuUsage field to given value.

### HasUserCpuUsage

`func (o *ContainerStats) HasUserCpuUsage() bool`

HasUserCpuUsage returns a boolean if a field has been set.

### GetSystemCpuUsage

`func (o *ContainerStats) GetSystemCpuUsage() float32`

GetSystemCpuUsage returns the SystemCpuUsage field if non-nil, zero value otherwise.

### GetSystemCpuUsageOk

`func (o *ContainerStats) GetSystemCpuUsageOk() (*float32, bool)`

GetSystemCpuUsageOk returns a tuple with the SystemCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemCpuUsage

`func (o *ContainerStats) SetSystemCpuUsage(v float32)`

SetSystemCpuUsage sets SystemCpuUsage field to given value.

### HasSystemCpuUsage

`func (o *ContainerStats) HasSystemCpuUsage() bool`

HasSystemCpuUsage returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ContainerStats) GetUsedMemory() int32`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ContainerStats) GetUsedMemoryOk() (*int32, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ContainerStats) SetUsedMemory(v int32)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ContainerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ContainerStats) GetMaxMemory() int32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ContainerStats) GetMaxMemoryOk() (*int32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ContainerStats) SetMaxMemory(v int32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ContainerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetCacheMemory

`func (o *ContainerStats) GetCacheMemory() int32`

GetCacheMemory returns the CacheMemory field if non-nil, zero value otherwise.

### GetCacheMemoryOk

`func (o *ContainerStats) GetCacheMemoryOk() (*int32, bool)`

GetCacheMemoryOk returns a tuple with the CacheMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMemory

`func (o *ContainerStats) SetCacheMemory(v int32)`

SetCacheMemory sets CacheMemory field to given value.

### HasCacheMemory

`func (o *ContainerStats) HasCacheMemory() bool`

HasCacheMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ContainerStats) GetMaxStorage() int32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ContainerStats) GetMaxStorageOk() (*int32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ContainerStats) SetMaxStorage(v int32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ContainerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ContainerStats) GetUsedStorage() int32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ContainerStats) GetUsedStorageOk() (*int32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ContainerStats) SetUsedStorage(v int32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ContainerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReadIOPS

`func (o *ContainerStats) GetReadIOPS() int32`

GetReadIOPS returns the ReadIOPS field if non-nil, zero value otherwise.

### GetReadIOPSOk

`func (o *ContainerStats) GetReadIOPSOk() (*int32, bool)`

GetReadIOPSOk returns a tuple with the ReadIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadIOPS

`func (o *ContainerStats) SetReadIOPS(v int32)`

SetReadIOPS sets ReadIOPS field to given value.

### HasReadIOPS

`func (o *ContainerStats) HasReadIOPS() bool`

HasReadIOPS returns a boolean if a field has been set.

### GetWriteIOPS

`func (o *ContainerStats) GetWriteIOPS() float32`

GetWriteIOPS returns the WriteIOPS field if non-nil, zero value otherwise.

### GetWriteIOPSOk

`func (o *ContainerStats) GetWriteIOPSOk() (*float32, bool)`

GetWriteIOPSOk returns a tuple with the WriteIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteIOPS

`func (o *ContainerStats) SetWriteIOPS(v float32)`

SetWriteIOPS sets WriteIOPS field to given value.

### HasWriteIOPS

`func (o *ContainerStats) HasWriteIOPS() bool`

HasWriteIOPS returns a boolean if a field has been set.

### GetTotalIOPS

`func (o *ContainerStats) GetTotalIOPS() float32`

GetTotalIOPS returns the TotalIOPS field if non-nil, zero value otherwise.

### GetTotalIOPSOk

`func (o *ContainerStats) GetTotalIOPSOk() (*float32, bool)`

GetTotalIOPSOk returns a tuple with the TotalIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIOPS

`func (o *ContainerStats) SetTotalIOPS(v float32)`

SetTotalIOPS sets TotalIOPS field to given value.

### HasTotalIOPS

`func (o *ContainerStats) HasTotalIOPS() bool`

HasTotalIOPS returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *ContainerStats) GetNetTxUsage() int32`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *ContainerStats) GetNetTxUsageOk() (*int32, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *ContainerStats) SetNetTxUsage(v int32)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *ContainerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *ContainerStats) GetNetRxUsage() int32`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *ContainerStats) GetNetRxUsageOk() (*int32, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *ContainerStats) SetNetRxUsage(v int32)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *ContainerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


