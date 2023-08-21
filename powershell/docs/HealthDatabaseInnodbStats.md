# HealthDatabaseInnodbStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LargeMemory** | **Int64** |  | [optional] 
**DictionaryMemory** | **Int64** |  | [optional] 
**BufferPoolSize** | **Int64** |  | [optional] 
**FreeBuffers** | **Int64** |  | [optional] 
**DatabasePages** | **Int64** |  | [optional] 
**OldPages** | **Int64** |  | [optional] 
**PendingReads** | **Int64** |  | [optional] 
**InsertsPerSecond** | **Decimal** |  | [optional] 
**UpdatesPerSecond** | **Decimal** |  | [optional] 
**DeletesPerSecond** | **Decimal** |  | [optional] 
**ReadsPerSecond** | **Decimal** |  | [optional] 
**BufferHitRate** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthDatabaseInnodbStats = Initialize-PSOpenAPIToolsHealthDatabaseInnodbStats  -LargeMemory null `
 -DictionaryMemory null `
 -BufferPoolSize null `
 -FreeBuffers null `
 -DatabasePages null `
 -OldPages null `
 -PendingReads null `
 -InsertsPerSecond null `
 -UpdatesPerSecond null `
 -DeletesPerSecond null `
 -ReadsPerSecond null `
 -BufferHitRate null
```

- Convert the resource to JSON
```powershell
$HealthDatabaseInnodbStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

