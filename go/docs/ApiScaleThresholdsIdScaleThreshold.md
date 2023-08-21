# ApiScaleThresholdsIdScaleThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the scale threshold | [optional] 
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

### NewApiScaleThresholdsIdScaleThreshold

`func NewApiScaleThresholdsIdScaleThreshold() *ApiScaleThresholdsIdScaleThreshold`

NewApiScaleThresholdsIdScaleThreshold instantiates a new ApiScaleThresholdsIdScaleThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiScaleThresholdsIdScaleThresholdWithDefaults

`func NewApiScaleThresholdsIdScaleThresholdWithDefaults() *ApiScaleThresholdsIdScaleThreshold`

NewApiScaleThresholdsIdScaleThresholdWithDefaults instantiates a new ApiScaleThresholdsIdScaleThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiScaleThresholdsIdScaleThreshold) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiScaleThresholdsIdScaleThreshold) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiScaleThresholdsIdScaleThreshold) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoUp

`func (o *ApiScaleThresholdsIdScaleThreshold) GetAutoUp() bool`

GetAutoUp returns the AutoUp field if non-nil, zero value otherwise.

### GetAutoUpOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetAutoUpOk() (*bool, bool)`

GetAutoUpOk returns a tuple with the AutoUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUp

`func (o *ApiScaleThresholdsIdScaleThreshold) SetAutoUp(v bool)`

SetAutoUp sets AutoUp field to given value.

### HasAutoUp

`func (o *ApiScaleThresholdsIdScaleThreshold) HasAutoUp() bool`

HasAutoUp returns a boolean if a field has been set.

### GetAutoDown

`func (o *ApiScaleThresholdsIdScaleThreshold) GetAutoDown() bool`

GetAutoDown returns the AutoDown field if non-nil, zero value otherwise.

### GetAutoDownOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetAutoDownOk() (*bool, bool)`

GetAutoDownOk returns a tuple with the AutoDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDown

`func (o *ApiScaleThresholdsIdScaleThreshold) SetAutoDown(v bool)`

SetAutoDown sets AutoDown field to given value.

### HasAutoDown

`func (o *ApiScaleThresholdsIdScaleThreshold) HasAutoDown() bool`

HasAutoDown returns a boolean if a field has been set.

### GetMinCount

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetMaxCount

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetCpuEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) GetCpuEnabled() bool`

GetCpuEnabled returns the CpuEnabled field if non-nil, zero value otherwise.

### GetCpuEnabledOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetCpuEnabledOk() (*bool, bool)`

GetCpuEnabledOk returns a tuple with the CpuEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) SetCpuEnabled(v bool)`

SetCpuEnabled sets CpuEnabled field to given value.

### HasCpuEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) HasCpuEnabled() bool`

HasCpuEnabled returns a boolean if a field has been set.

### GetMinCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinCpu() float32`

GetMinCpu returns the MinCpu field if non-nil, zero value otherwise.

### GetMinCpuOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinCpuOk() (*float32, bool)`

GetMinCpuOk returns a tuple with the MinCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMinCpu(v float32)`

SetMinCpu sets MinCpu field to given value.

### HasMinCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMinCpu() bool`

HasMinCpu returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMemoryEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMemoryEnabled() bool`

GetMemoryEnabled returns the MemoryEnabled field if non-nil, zero value otherwise.

### GetMemoryEnabledOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMemoryEnabledOk() (*bool, bool)`

GetMemoryEnabledOk returns a tuple with the MemoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMemoryEnabled(v bool)`

SetMemoryEnabled sets MemoryEnabled field to given value.

### HasMemoryEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMemoryEnabled() bool`

HasMemoryEnabled returns a boolean if a field has been set.

### GetMinMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinMemory() float32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinMemoryOk() (*float32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMinMemory(v float32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetDiskEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) GetDiskEnabled() bool`

GetDiskEnabled returns the DiskEnabled field if non-nil, zero value otherwise.

### GetDiskEnabledOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetDiskEnabledOk() (*bool, bool)`

GetDiskEnabledOk returns a tuple with the DiskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) SetDiskEnabled(v bool)`

SetDiskEnabled sets DiskEnabled field to given value.

### HasDiskEnabled

`func (o *ApiScaleThresholdsIdScaleThreshold) HasDiskEnabled() bool`

HasDiskEnabled returns a boolean if a field has been set.

### GetMinDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinDisk() float32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMinDiskOk() (*float32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMinDisk(v float32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMaxDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxDisk() float32`

GetMaxDisk returns the MaxDisk field if non-nil, zero value otherwise.

### GetMaxDiskOk

`func (o *ApiScaleThresholdsIdScaleThreshold) GetMaxDiskOk() (*float32, bool)`

GetMaxDiskOk returns a tuple with the MaxDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) SetMaxDisk(v float32)`

SetMaxDisk sets MaxDisk field to given value.

### HasMaxDisk

`func (o *ApiScaleThresholdsIdScaleThreshold) HasMaxDisk() bool`

HasMaxDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


