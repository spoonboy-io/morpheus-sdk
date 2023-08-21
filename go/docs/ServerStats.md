# ServerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** |  | [optional] 
**FreeSwap** | Pointer to **int64** |  | [optional] 
**UsedSwap** | Pointer to **int64** |  | [optional] 
**CpuIdleTime** | Pointer to **int64** |  | [optional] 
**CpuSystemTime** | Pointer to **int64** |  | [optional] 
**CpuUserTime** | Pointer to **int64** |  | [optional] 
**CpuTotalTime** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**UsedMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**UsedStorage** | Pointer to **int64** |  | [optional] 
**ReservedStorage** | Pointer to **int64** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**FreeMemory** | Pointer to **int64** |  | [optional] 
**NetTxUsage** | Pointer to **int64** |  | [optional] 
**NetRxUsage** | Pointer to **int64** |  | [optional] 
**NetworkBandwidth** | Pointer to **int64** |  | [optional] 

## Methods

### NewServerStats

`func NewServerStats() *ServerStats`

NewServerStats instantiates a new ServerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatsWithDefaults

`func NewServerStatsWithDefaults() *ServerStats`

NewServerStatsWithDefaults instantiates a new ServerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *ServerStats) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ServerStats) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ServerStats) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ServerStats) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetFreeSwap

`func (o *ServerStats) GetFreeSwap() int64`

GetFreeSwap returns the FreeSwap field if non-nil, zero value otherwise.

### GetFreeSwapOk

`func (o *ServerStats) GetFreeSwapOk() (*int64, bool)`

GetFreeSwapOk returns a tuple with the FreeSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSwap

`func (o *ServerStats) SetFreeSwap(v int64)`

SetFreeSwap sets FreeSwap field to given value.

### HasFreeSwap

`func (o *ServerStats) HasFreeSwap() bool`

HasFreeSwap returns a boolean if a field has been set.

### GetUsedSwap

`func (o *ServerStats) GetUsedSwap() int64`

GetUsedSwap returns the UsedSwap field if non-nil, zero value otherwise.

### GetUsedSwapOk

`func (o *ServerStats) GetUsedSwapOk() (*int64, bool)`

GetUsedSwapOk returns a tuple with the UsedSwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSwap

`func (o *ServerStats) SetUsedSwap(v int64)`

SetUsedSwap sets UsedSwap field to given value.

### HasUsedSwap

`func (o *ServerStats) HasUsedSwap() bool`

HasUsedSwap returns a boolean if a field has been set.

### GetCpuIdleTime

`func (o *ServerStats) GetCpuIdleTime() int64`

GetCpuIdleTime returns the CpuIdleTime field if non-nil, zero value otherwise.

### GetCpuIdleTimeOk

`func (o *ServerStats) GetCpuIdleTimeOk() (*int64, bool)`

GetCpuIdleTimeOk returns a tuple with the CpuIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTime

`func (o *ServerStats) SetCpuIdleTime(v int64)`

SetCpuIdleTime sets CpuIdleTime field to given value.

### HasCpuIdleTime

`func (o *ServerStats) HasCpuIdleTime() bool`

HasCpuIdleTime returns a boolean if a field has been set.

### GetCpuSystemTime

`func (o *ServerStats) GetCpuSystemTime() int64`

GetCpuSystemTime returns the CpuSystemTime field if non-nil, zero value otherwise.

### GetCpuSystemTimeOk

`func (o *ServerStats) GetCpuSystemTimeOk() (*int64, bool)`

GetCpuSystemTimeOk returns a tuple with the CpuSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTime

`func (o *ServerStats) SetCpuSystemTime(v int64)`

SetCpuSystemTime sets CpuSystemTime field to given value.

### HasCpuSystemTime

`func (o *ServerStats) HasCpuSystemTime() bool`

HasCpuSystemTime returns a boolean if a field has been set.

### GetCpuUserTime

`func (o *ServerStats) GetCpuUserTime() int64`

GetCpuUserTime returns the CpuUserTime field if non-nil, zero value otherwise.

### GetCpuUserTimeOk

`func (o *ServerStats) GetCpuUserTimeOk() (*int64, bool)`

GetCpuUserTimeOk returns a tuple with the CpuUserTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTime

`func (o *ServerStats) SetCpuUserTime(v int64)`

SetCpuUserTime sets CpuUserTime field to given value.

### HasCpuUserTime

`func (o *ServerStats) HasCpuUserTime() bool`

HasCpuUserTime returns a boolean if a field has been set.

### GetCpuTotalTime

`func (o *ServerStats) GetCpuTotalTime() int64`

GetCpuTotalTime returns the CpuTotalTime field if non-nil, zero value otherwise.

### GetCpuTotalTimeOk

`func (o *ServerStats) GetCpuTotalTimeOk() (*int64, bool)`

GetCpuTotalTimeOk returns a tuple with the CpuTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTime

`func (o *ServerStats) SetCpuTotalTime(v int64)`

SetCpuTotalTime sets CpuTotalTime field to given value.

### HasCpuTotalTime

`func (o *ServerStats) HasCpuTotalTime() bool`

HasCpuTotalTime returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ServerStats) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ServerStats) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ServerStats) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ServerStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedMemory

`func (o *ServerStats) GetUsedMemory() int64`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *ServerStats) GetUsedMemoryOk() (*int64, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *ServerStats) SetUsedMemory(v int64)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *ServerStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ServerStats) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ServerStats) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ServerStats) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ServerStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ServerStats) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ServerStats) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ServerStats) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ServerStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetReservedStorage

`func (o *ServerStats) GetReservedStorage() int64`

GetReservedStorage returns the ReservedStorage field if non-nil, zero value otherwise.

### GetReservedStorageOk

`func (o *ServerStats) GetReservedStorageOk() (*int64, bool)`

GetReservedStorageOk returns a tuple with the ReservedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStorage

`func (o *ServerStats) SetReservedStorage(v int64)`

SetReservedStorage sets ReservedStorage field to given value.

### HasReservedStorage

`func (o *ServerStats) HasReservedStorage() bool`

HasReservedStorage returns a boolean if a field has been set.

### GetCpuUsage

`func (o *ServerStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *ServerStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *ServerStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *ServerStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetFreeMemory

`func (o *ServerStats) GetFreeMemory() int64`

GetFreeMemory returns the FreeMemory field if non-nil, zero value otherwise.

### GetFreeMemoryOk

`func (o *ServerStats) GetFreeMemoryOk() (*int64, bool)`

GetFreeMemoryOk returns a tuple with the FreeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemory

`func (o *ServerStats) SetFreeMemory(v int64)`

SetFreeMemory sets FreeMemory field to given value.

### HasFreeMemory

`func (o *ServerStats) HasFreeMemory() bool`

HasFreeMemory returns a boolean if a field has been set.

### GetNetTxUsage

`func (o *ServerStats) GetNetTxUsage() int64`

GetNetTxUsage returns the NetTxUsage field if non-nil, zero value otherwise.

### GetNetTxUsageOk

`func (o *ServerStats) GetNetTxUsageOk() (*int64, bool)`

GetNetTxUsageOk returns a tuple with the NetTxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTxUsage

`func (o *ServerStats) SetNetTxUsage(v int64)`

SetNetTxUsage sets NetTxUsage field to given value.

### HasNetTxUsage

`func (o *ServerStats) HasNetTxUsage() bool`

HasNetTxUsage returns a boolean if a field has been set.

### GetNetRxUsage

`func (o *ServerStats) GetNetRxUsage() int64`

GetNetRxUsage returns the NetRxUsage field if non-nil, zero value otherwise.

### GetNetRxUsageOk

`func (o *ServerStats) GetNetRxUsageOk() (*int64, bool)`

GetNetRxUsageOk returns a tuple with the NetRxUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetRxUsage

`func (o *ServerStats) SetNetRxUsage(v int64)`

SetNetRxUsage sets NetRxUsage field to given value.

### HasNetRxUsage

`func (o *ServerStats) HasNetRxUsage() bool`

HasNetRxUsage returns a boolean if a field has been set.

### GetNetworkBandwidth

`func (o *ServerStats) GetNetworkBandwidth() int64`

GetNetworkBandwidth returns the NetworkBandwidth field if non-nil, zero value otherwise.

### GetNetworkBandwidthOk

`func (o *ServerStats) GetNetworkBandwidthOk() (*int64, bool)`

GetNetworkBandwidthOk returns a tuple with the NetworkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkBandwidth

`func (o *ServerStats) SetNetworkBandwidth(v int64)`

SetNetworkBandwidth sets NetworkBandwidth field to given value.

### HasNetworkBandwidth

`func (o *ServerStats) HasNetworkBandwidth() bool`

HasNetworkBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


