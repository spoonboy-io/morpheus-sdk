# HealthThreads
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**ThreadList** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**BusyThreads** | [**HealthThreadsBusyThreads[]**](HealthThreadsBusyThreads.md) |  | [optional] 
**BlockedThreads** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**RunningThreads** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**TotalCpuTime** | **Int64** |  | [optional] 
**TotalThreads** | **Int64** |  | [optional] 
**RunningWebThreads** | **Int64** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthThreads = Initialize-PSOpenAPIToolsHealthThreads  -Success null `
 -ThreadList null `
 -BusyThreads null `
 -BlockedThreads null `
 -RunningThreads null `
 -TotalCpuTime null `
 -TotalThreads null `
 -RunningWebThreads null `
 -Status null
```

- Convert the resource to JSON
```powershell
$HealthThreads | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

