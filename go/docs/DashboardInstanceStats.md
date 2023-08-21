# DashboardInstanceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCpu** | Pointer to **float32** |  | [optional] 
**MaxCores** | Pointer to **float32** |  | [optional] 
**CpuUsage** | Pointer to **float32** |  | [optional] 
**CpuUsageAverage** | Pointer to **float32** |  | [optional] 
**CpuUsagePeak** | Pointer to **float32** |  | [optional] 
**UsedMemory** | Pointer to **float32** |  | [optional] 
**MaxMemory** | Pointer to **float32** |  | [optional] 
**UsedStorage** | Pointer to **float32** |  | [optional] 
**MaxStorage** | Pointer to **float32** |  | [optional] 
**Running** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**TotalContainers** | Pointer to **float32** |  | [optional] 

## Methods

### NewDashboardInstanceStats

`func NewDashboardInstanceStats() *DashboardInstanceStats`

NewDashboardInstanceStats instantiates a new DashboardInstanceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardInstanceStatsWithDefaults

`func NewDashboardInstanceStatsWithDefaults() *DashboardInstanceStats`

NewDashboardInstanceStatsWithDefaults instantiates a new DashboardInstanceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCpu

`func (o *DashboardInstanceStats) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *DashboardInstanceStats) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *DashboardInstanceStats) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *DashboardInstanceStats) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *DashboardInstanceStats) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *DashboardInstanceStats) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *DashboardInstanceStats) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *DashboardInstanceStats) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCpuUsage

`func (o *DashboardInstanceStats) GetCpuUsage() float32`

GetCpuUsage returns the CpuUsage field if non-nil, zero value otherwise.

### GetCpuUsageOk

`func (o *DashboardInstanceStats) GetCpuUsageOk() (*float32, bool)`

GetCpuUsageOk returns a tuple with the CpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsage

`func (o *DashboardInstanceStats) SetCpuUsage(v float32)`

SetCpuUsage sets CpuUsage field to given value.

### HasCpuUsage

`func (o *DashboardInstanceStats) HasCpuUsage() bool`

HasCpuUsage returns a boolean if a field has been set.

### GetCpuUsageAverage

`func (o *DashboardInstanceStats) GetCpuUsageAverage() float32`

GetCpuUsageAverage returns the CpuUsageAverage field if non-nil, zero value otherwise.

### GetCpuUsageAverageOk

`func (o *DashboardInstanceStats) GetCpuUsageAverageOk() (*float32, bool)`

GetCpuUsageAverageOk returns a tuple with the CpuUsageAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageAverage

`func (o *DashboardInstanceStats) SetCpuUsageAverage(v float32)`

SetCpuUsageAverage sets CpuUsageAverage field to given value.

### HasCpuUsageAverage

`func (o *DashboardInstanceStats) HasCpuUsageAverage() bool`

HasCpuUsageAverage returns a boolean if a field has been set.

### GetCpuUsagePeak

`func (o *DashboardInstanceStats) GetCpuUsagePeak() float32`

GetCpuUsagePeak returns the CpuUsagePeak field if non-nil, zero value otherwise.

### GetCpuUsagePeakOk

`func (o *DashboardInstanceStats) GetCpuUsagePeakOk() (*float32, bool)`

GetCpuUsagePeakOk returns a tuple with the CpuUsagePeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsagePeak

`func (o *DashboardInstanceStats) SetCpuUsagePeak(v float32)`

SetCpuUsagePeak sets CpuUsagePeak field to given value.

### HasCpuUsagePeak

`func (o *DashboardInstanceStats) HasCpuUsagePeak() bool`

HasCpuUsagePeak returns a boolean if a field has been set.

### GetUsedMemory

`func (o *DashboardInstanceStats) GetUsedMemory() float32`

GetUsedMemory returns the UsedMemory field if non-nil, zero value otherwise.

### GetUsedMemoryOk

`func (o *DashboardInstanceStats) GetUsedMemoryOk() (*float32, bool)`

GetUsedMemoryOk returns a tuple with the UsedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemory

`func (o *DashboardInstanceStats) SetUsedMemory(v float32)`

SetUsedMemory sets UsedMemory field to given value.

### HasUsedMemory

`func (o *DashboardInstanceStats) HasUsedMemory() bool`

HasUsedMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *DashboardInstanceStats) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *DashboardInstanceStats) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *DashboardInstanceStats) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *DashboardInstanceStats) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetUsedStorage

`func (o *DashboardInstanceStats) GetUsedStorage() float32`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *DashboardInstanceStats) GetUsedStorageOk() (*float32, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *DashboardInstanceStats) SetUsedStorage(v float32)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *DashboardInstanceStats) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *DashboardInstanceStats) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *DashboardInstanceStats) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *DashboardInstanceStats) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *DashboardInstanceStats) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetRunning

`func (o *DashboardInstanceStats) GetRunning() float32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *DashboardInstanceStats) GetRunningOk() (*float32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *DashboardInstanceStats) SetRunning(v float32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *DashboardInstanceStats) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTotal

`func (o *DashboardInstanceStats) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DashboardInstanceStats) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DashboardInstanceStats) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DashboardInstanceStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalContainers

`func (o *DashboardInstanceStats) GetTotalContainers() float32`

GetTotalContainers returns the TotalContainers field if non-nil, zero value otherwise.

### GetTotalContainersOk

`func (o *DashboardInstanceStats) GetTotalContainersOk() (*float32, bool)`

GetTotalContainersOk returns a tuple with the TotalContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContainers

`func (o *DashboardInstanceStats) SetTotalContainers(v float32)`

SetTotalContainers sets TotalContainers field to given value.

### HasTotalContainers

`func (o *DashboardInstanceStats) HasTotalContainers() bool`

HasTotalContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


