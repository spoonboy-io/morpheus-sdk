# ClusterContainersStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **System.DateTime** |  | [optional] 
**Running** | **Boolean** |  | [optional] 
**UserCpuUsage** | **Decimal** |  | [optional] 
**SystemCpuUsage** | **Decimal** |  | [optional] 
**UsedMemory** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**CacheMemory** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**UsedStorage** | **Int64** |  | [optional] 
**ReadIOPS** | **Int64** |  | [optional] 
**WriteIOPS** | **Int64** |  | [optional] 
**TotalIOPS** | **Int64** |  | [optional] 
**NetTxUsage** | **Int64** |  | [optional] 
**NetRxUsage** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterContainersStats = Initialize-PSOpenAPIToolsClusterContainersStats  -Ts null `
 -Running null `
 -UserCpuUsage null `
 -SystemCpuUsage null `
 -UsedMemory null `
 -MaxMemory null `
 -CacheMemory null `
 -MaxStorage null `
 -UsedStorage null `
 -ReadIOPS null `
 -WriteIOPS null `
 -TotalIOPS null `
 -NetTxUsage null `
 -NetRxUsage null
```

- Convert the resource to JSON
```powershell
$ClusterContainersStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

