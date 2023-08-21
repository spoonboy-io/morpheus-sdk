# ClusterMastersStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedStorage** | **Int64** |  | [optional] 
**ReservedStorage** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**UsedMemory** | **Int64** |  | [optional] 
**ReservedMemory** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**CpuUsage** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterMastersStats = Initialize-PSOpenAPIToolsClusterMastersStats  -UsedStorage null `
 -ReservedStorage null `
 -MaxStorage null `
 -UsedMemory null `
 -ReservedMemory null `
 -MaxMemory null `
 -CpuUsage null
```

- Convert the resource to JSON
```powershell
$ClusterMastersStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

