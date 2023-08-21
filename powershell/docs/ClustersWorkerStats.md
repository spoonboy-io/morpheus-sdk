# ClustersWorkerStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedStorage** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**UsedMemory** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**UsedCpu** | **Int64** |  | [optional] 
**CpuUsage** | **Decimal** |  | [optional] 
**CpuUsagePeak** | **Decimal** |  | [optional] 
**CpuUsageAvg** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClustersWorkerStats = Initialize-PSOpenAPIToolsClustersWorkerStats  -UsedStorage null `
 -MaxStorage null `
 -UsedMemory null `
 -MaxMemory null `
 -UsedCpu null `
 -CpuUsage null `
 -CpuUsagePeak null `
 -CpuUsageAvg null
```

- Convert the resource to JSON
```powershell
$ClustersWorkerStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

