# HealthThreadsBusyThreadsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**CpuTime** | **Int64** |  | [optional] 
**BlockedTime** | **Int64** |  | [optional] 
**LockName** | **String** |  | [optional] 
**LockOwnerId** | **Int64** |  | [optional] 
**LockOwnerName** | **String** |  | [optional] 
**State** | **String** |  | [optional] 
**WaitedCount** | **Int64** |  | [optional] 
**WaitedTime** | **Int64** |  | [optional] 
**IsInNative** | **Boolean** |  | [optional] 
**IsSuspended** | **Boolean** |  | [optional] 
**LockedMonitors** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**LockedSynchronizers** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**LockInfo** | **String** |  | [optional] 
**CurrentLines** | **String** |  | [optional] 
**CpuPercent** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthThreadsBusyThreadsInner = Initialize-PSOpenAPIToolsHealthThreadsBusyThreadsInner  -ThreadId null `
 -Name null `
 -CpuTime null `
 -BlockedTime null `
 -LockName null `
 -LockOwnerId null `
 -LockOwnerName null `
 -State null `
 -WaitedCount null `
 -WaitedTime null `
 -IsInNative null `
 -IsSuspended null `
 -LockedMonitors null `
 -LockedSynchronizers null `
 -LockInfo null `
 -CurrentLines null `
 -CpuPercent null
```

- Convert the resource to JSON
```powershell
$HealthThreadsBusyThreadsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

