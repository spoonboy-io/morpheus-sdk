# HealthThreadsBusyThreads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CpuTime** | Pointer to **int64** |  | [optional] 
**BlockedTime** | Pointer to **int64** |  | [optional] 
**LockName** | Pointer to **NullableString** |  | [optional] 
**LockOwnerId** | Pointer to **int64** |  | [optional] 
**LockOwnerName** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**WaitedCount** | Pointer to **int64** |  | [optional] 
**WaitedTime** | Pointer to **int64** |  | [optional] 
**IsInNative** | Pointer to **bool** |  | [optional] 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**LockedMonitors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockedSynchronizers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**LockInfo** | Pointer to **NullableString** |  | [optional] 
**CurrentLines** | Pointer to **string** |  | [optional] 
**CpuPercent** | Pointer to **float32** |  | [optional] 

## Methods

### NewHealthThreadsBusyThreads

`func NewHealthThreadsBusyThreads() *HealthThreadsBusyThreads`

NewHealthThreadsBusyThreads instantiates a new HealthThreadsBusyThreads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthThreadsBusyThreadsWithDefaults

`func NewHealthThreadsBusyThreadsWithDefaults() *HealthThreadsBusyThreads`

NewHealthThreadsBusyThreadsWithDefaults instantiates a new HealthThreadsBusyThreads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *HealthThreadsBusyThreads) GetThreadId() int64`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *HealthThreadsBusyThreads) GetThreadIdOk() (*int64, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *HealthThreadsBusyThreads) SetThreadId(v int64)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *HealthThreadsBusyThreads) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetName

`func (o *HealthThreadsBusyThreads) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthThreadsBusyThreads) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthThreadsBusyThreads) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HealthThreadsBusyThreads) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpuTime

`func (o *HealthThreadsBusyThreads) GetCpuTime() int64`

GetCpuTime returns the CpuTime field if non-nil, zero value otherwise.

### GetCpuTimeOk

`func (o *HealthThreadsBusyThreads) GetCpuTimeOk() (*int64, bool)`

GetCpuTimeOk returns a tuple with the CpuTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTime

`func (o *HealthThreadsBusyThreads) SetCpuTime(v int64)`

SetCpuTime sets CpuTime field to given value.

### HasCpuTime

`func (o *HealthThreadsBusyThreads) HasCpuTime() bool`

HasCpuTime returns a boolean if a field has been set.

### GetBlockedTime

`func (o *HealthThreadsBusyThreads) GetBlockedTime() int64`

GetBlockedTime returns the BlockedTime field if non-nil, zero value otherwise.

### GetBlockedTimeOk

`func (o *HealthThreadsBusyThreads) GetBlockedTimeOk() (*int64, bool)`

GetBlockedTimeOk returns a tuple with the BlockedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTime

`func (o *HealthThreadsBusyThreads) SetBlockedTime(v int64)`

SetBlockedTime sets BlockedTime field to given value.

### HasBlockedTime

`func (o *HealthThreadsBusyThreads) HasBlockedTime() bool`

HasBlockedTime returns a boolean if a field has been set.

### GetLockName

`func (o *HealthThreadsBusyThreads) GetLockName() string`

GetLockName returns the LockName field if non-nil, zero value otherwise.

### GetLockNameOk

`func (o *HealthThreadsBusyThreads) GetLockNameOk() (*string, bool)`

GetLockNameOk returns a tuple with the LockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockName

`func (o *HealthThreadsBusyThreads) SetLockName(v string)`

SetLockName sets LockName field to given value.

### HasLockName

`func (o *HealthThreadsBusyThreads) HasLockName() bool`

HasLockName returns a boolean if a field has been set.

### SetLockNameNil

`func (o *HealthThreadsBusyThreads) SetLockNameNil(b bool)`

 SetLockNameNil sets the value for LockName to be an explicit nil

### UnsetLockName
`func (o *HealthThreadsBusyThreads) UnsetLockName()`

UnsetLockName ensures that no value is present for LockName, not even an explicit nil
### GetLockOwnerId

`func (o *HealthThreadsBusyThreads) GetLockOwnerId() int64`

GetLockOwnerId returns the LockOwnerId field if non-nil, zero value otherwise.

### GetLockOwnerIdOk

`func (o *HealthThreadsBusyThreads) GetLockOwnerIdOk() (*int64, bool)`

GetLockOwnerIdOk returns a tuple with the LockOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerId

`func (o *HealthThreadsBusyThreads) SetLockOwnerId(v int64)`

SetLockOwnerId sets LockOwnerId field to given value.

### HasLockOwnerId

`func (o *HealthThreadsBusyThreads) HasLockOwnerId() bool`

HasLockOwnerId returns a boolean if a field has been set.

### GetLockOwnerName

`func (o *HealthThreadsBusyThreads) GetLockOwnerName() string`

GetLockOwnerName returns the LockOwnerName field if non-nil, zero value otherwise.

### GetLockOwnerNameOk

`func (o *HealthThreadsBusyThreads) GetLockOwnerNameOk() (*string, bool)`

GetLockOwnerNameOk returns a tuple with the LockOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOwnerName

`func (o *HealthThreadsBusyThreads) SetLockOwnerName(v string)`

SetLockOwnerName sets LockOwnerName field to given value.

### HasLockOwnerName

`func (o *HealthThreadsBusyThreads) HasLockOwnerName() bool`

HasLockOwnerName returns a boolean if a field has been set.

### SetLockOwnerNameNil

`func (o *HealthThreadsBusyThreads) SetLockOwnerNameNil(b bool)`

 SetLockOwnerNameNil sets the value for LockOwnerName to be an explicit nil

### UnsetLockOwnerName
`func (o *HealthThreadsBusyThreads) UnsetLockOwnerName()`

UnsetLockOwnerName ensures that no value is present for LockOwnerName, not even an explicit nil
### GetState

`func (o *HealthThreadsBusyThreads) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HealthThreadsBusyThreads) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HealthThreadsBusyThreads) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HealthThreadsBusyThreads) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWaitedCount

`func (o *HealthThreadsBusyThreads) GetWaitedCount() int64`

GetWaitedCount returns the WaitedCount field if non-nil, zero value otherwise.

### GetWaitedCountOk

`func (o *HealthThreadsBusyThreads) GetWaitedCountOk() (*int64, bool)`

GetWaitedCountOk returns a tuple with the WaitedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedCount

`func (o *HealthThreadsBusyThreads) SetWaitedCount(v int64)`

SetWaitedCount sets WaitedCount field to given value.

### HasWaitedCount

`func (o *HealthThreadsBusyThreads) HasWaitedCount() bool`

HasWaitedCount returns a boolean if a field has been set.

### GetWaitedTime

`func (o *HealthThreadsBusyThreads) GetWaitedTime() int64`

GetWaitedTime returns the WaitedTime field if non-nil, zero value otherwise.

### GetWaitedTimeOk

`func (o *HealthThreadsBusyThreads) GetWaitedTimeOk() (*int64, bool)`

GetWaitedTimeOk returns a tuple with the WaitedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitedTime

`func (o *HealthThreadsBusyThreads) SetWaitedTime(v int64)`

SetWaitedTime sets WaitedTime field to given value.

### HasWaitedTime

`func (o *HealthThreadsBusyThreads) HasWaitedTime() bool`

HasWaitedTime returns a boolean if a field has been set.

### GetIsInNative

`func (o *HealthThreadsBusyThreads) GetIsInNative() bool`

GetIsInNative returns the IsInNative field if non-nil, zero value otherwise.

### GetIsInNativeOk

`func (o *HealthThreadsBusyThreads) GetIsInNativeOk() (*bool, bool)`

GetIsInNativeOk returns a tuple with the IsInNative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInNative

`func (o *HealthThreadsBusyThreads) SetIsInNative(v bool)`

SetIsInNative sets IsInNative field to given value.

### HasIsInNative

`func (o *HealthThreadsBusyThreads) HasIsInNative() bool`

HasIsInNative returns a boolean if a field has been set.

### GetIsSuspended

`func (o *HealthThreadsBusyThreads) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *HealthThreadsBusyThreads) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *HealthThreadsBusyThreads) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *HealthThreadsBusyThreads) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetLockedMonitors

`func (o *HealthThreadsBusyThreads) GetLockedMonitors() []map[string]interface{}`

GetLockedMonitors returns the LockedMonitors field if non-nil, zero value otherwise.

### GetLockedMonitorsOk

`func (o *HealthThreadsBusyThreads) GetLockedMonitorsOk() (*[]map[string]interface{}, bool)`

GetLockedMonitorsOk returns a tuple with the LockedMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedMonitors

`func (o *HealthThreadsBusyThreads) SetLockedMonitors(v []map[string]interface{})`

SetLockedMonitors sets LockedMonitors field to given value.

### HasLockedMonitors

`func (o *HealthThreadsBusyThreads) HasLockedMonitors() bool`

HasLockedMonitors returns a boolean if a field has been set.

### SetLockedMonitorsNil

`func (o *HealthThreadsBusyThreads) SetLockedMonitorsNil(b bool)`

 SetLockedMonitorsNil sets the value for LockedMonitors to be an explicit nil

### UnsetLockedMonitors
`func (o *HealthThreadsBusyThreads) UnsetLockedMonitors()`

UnsetLockedMonitors ensures that no value is present for LockedMonitors, not even an explicit nil
### GetLockedSynchronizers

`func (o *HealthThreadsBusyThreads) GetLockedSynchronizers() []map[string]interface{}`

GetLockedSynchronizers returns the LockedSynchronizers field if non-nil, zero value otherwise.

### GetLockedSynchronizersOk

`func (o *HealthThreadsBusyThreads) GetLockedSynchronizersOk() (*[]map[string]interface{}, bool)`

GetLockedSynchronizersOk returns a tuple with the LockedSynchronizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedSynchronizers

`func (o *HealthThreadsBusyThreads) SetLockedSynchronizers(v []map[string]interface{})`

SetLockedSynchronizers sets LockedSynchronizers field to given value.

### HasLockedSynchronizers

`func (o *HealthThreadsBusyThreads) HasLockedSynchronizers() bool`

HasLockedSynchronizers returns a boolean if a field has been set.

### SetLockedSynchronizersNil

`func (o *HealthThreadsBusyThreads) SetLockedSynchronizersNil(b bool)`

 SetLockedSynchronizersNil sets the value for LockedSynchronizers to be an explicit nil

### UnsetLockedSynchronizers
`func (o *HealthThreadsBusyThreads) UnsetLockedSynchronizers()`

UnsetLockedSynchronizers ensures that no value is present for LockedSynchronizers, not even an explicit nil
### GetLockInfo

`func (o *HealthThreadsBusyThreads) GetLockInfo() string`

GetLockInfo returns the LockInfo field if non-nil, zero value otherwise.

### GetLockInfoOk

`func (o *HealthThreadsBusyThreads) GetLockInfoOk() (*string, bool)`

GetLockInfoOk returns a tuple with the LockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInfo

`func (o *HealthThreadsBusyThreads) SetLockInfo(v string)`

SetLockInfo sets LockInfo field to given value.

### HasLockInfo

`func (o *HealthThreadsBusyThreads) HasLockInfo() bool`

HasLockInfo returns a boolean if a field has been set.

### SetLockInfoNil

`func (o *HealthThreadsBusyThreads) SetLockInfoNil(b bool)`

 SetLockInfoNil sets the value for LockInfo to be an explicit nil

### UnsetLockInfo
`func (o *HealthThreadsBusyThreads) UnsetLockInfo()`

UnsetLockInfo ensures that no value is present for LockInfo, not even an explicit nil
### GetCurrentLines

`func (o *HealthThreadsBusyThreads) GetCurrentLines() string`

GetCurrentLines returns the CurrentLines field if non-nil, zero value otherwise.

### GetCurrentLinesOk

`func (o *HealthThreadsBusyThreads) GetCurrentLinesOk() (*string, bool)`

GetCurrentLinesOk returns a tuple with the CurrentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLines

`func (o *HealthThreadsBusyThreads) SetCurrentLines(v string)`

SetCurrentLines sets CurrentLines field to given value.

### HasCurrentLines

`func (o *HealthThreadsBusyThreads) HasCurrentLines() bool`

HasCurrentLines returns a boolean if a field has been set.

### GetCpuPercent

`func (o *HealthThreadsBusyThreads) GetCpuPercent() float32`

GetCpuPercent returns the CpuPercent field if non-nil, zero value otherwise.

### GetCpuPercentOk

`func (o *HealthThreadsBusyThreads) GetCpuPercentOk() (*float32, bool)`

GetCpuPercentOk returns a tuple with the CpuPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPercent

`func (o *HealthThreadsBusyThreads) SetCpuPercent(v float32)`

SetCpuPercent sets CpuPercent field to given value.

### HasCpuPercent

`func (o *HealthThreadsBusyThreads) HasCpuPercent() bool`

HasCpuPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


