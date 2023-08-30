# HealthElasticStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **String** |  | [optional] 
**ClusterName** | **String** |  | [optional] 
**NodeTotal** | **String** |  | [optional] 
**NodeData** | **String** |  | [optional] 
**Shards** | **String** |  | [optional] 
**Primary** | **String** |  | [optional] 
**Relocating** | **String** |  | [optional] 
**Initializing** | **String** |  | [optional] 
**Unassigned** | **String** |  | [optional] 
**PendingTasks** | **String** |  | [optional] 
**ActivePercent** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthElasticStats = Initialize-PSOpenAPIToolsHealthElasticStats  -Status null `
 -ClusterName null `
 -NodeTotal null `
 -NodeData null `
 -Shards null `
 -Primary null `
 -Relocating null `
 -Initializing null `
 -Unassigned null `
 -PendingTasks null `
 -ActivePercent null
```

- Convert the resource to JSON
```powershell
$HealthElasticStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

