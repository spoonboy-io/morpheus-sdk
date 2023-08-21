# ApiInstancesIdThresholdInstanceThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUp** | Pointer to **bool** | Auto Upscale | [optional] [default to false]
**AutoDown** | Pointer to **bool** | Auto Downscale | [optional] [default to false]
**MinCount** | Pointer to **int32** | The minimum number of nodes to scale down to | [optional] 
**MaxCount** | Pointer to **int32** | The maximum number of nodes to scale up to | [optional] 
**CpuEnabled** | Pointer to **bool** | Enable CPU Threshold | [optional] [default to false]
**MinCpu** | Pointer to **float32** | Min CPU (%) | [optional] [default to 0]
**MaxCpu** | Pointer to **float32** | Max CPU (%) | [optional] [default to 0]
**MemoryEnabled** | Pointer to **bool** | Enable Memory Threshold | [optional] [default to false]
**MinMemory** | Pointer to **float32** | Min Memory (%) | [optional] [default to 0]
**MaxMemory** | Pointer to **float32** | Max Memory (%) | [optional] [default to 0]
**DiskEnabled** | Pointer to **bool** | Enable Disk Threshold | [optional] [default to false]
**MinDisk** | Pointer to **float32** | Min Disk (%) | [optional] [default to 0]
**MaxDisk** | Pointer to **float32** | Max Disk (%) | [optional] [default to 0]

## Methods

### NewApiInstancesIdThresholdInstanceThreshold

`func NewApiInstancesIdThresholdInstanceThreshold() *ApiInstancesIdThresholdInstanceThreshold`

NewApiInstancesIdThresholdInstanceThreshold instantiates a new ApiInstancesIdThresholdInstanceThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInstancesIdThresholdInstanceThresholdWithDefaults

`func NewApiInstancesIdThresholdInstanceThresholdWithDefaults() *ApiInstancesIdThresholdInstanceThreshold`

NewApiInstancesIdThresholdInstanceThresholdWithDefaults instantiates a new ApiInstancesIdThresholdInstanceThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUp

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCpu() float32`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinCpuOk() (*float32, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinCpu(v float32)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinMemory() float32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinMemoryOk() (*float32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinMemory(v float32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinDisk() float32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMinDiskOk() (*float32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMinDisk(v float32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxDisk() float32`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ApiInstancesIdThresholdInstanceThreshold) GetMaxDiskOk() (*float32, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) SetMaxDisk(v float32)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ApiInstancesIdThresholdInstanceThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


