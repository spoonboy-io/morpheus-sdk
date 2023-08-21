# DashboardInstanceStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCpu** | **Decimal** |  | [optional] 
**MaxCores** | **Decimal** |  | [optional] 
**CpuUsage** | **Decimal** |  | [optional] 
**CpuUsageAverage** | **Decimal** |  | [optional] 
**CpuUsagePeak** | **Decimal** |  | [optional] 
**UsedMemory** | **Decimal** |  | [optional] 
**MaxMemory** | **Decimal** |  | [optional] 
**UsedStorage** | **Decimal** |  | [optional] 
**MaxStorage** | **Decimal** |  | [optional] 
**Running** | **Decimal** |  | [optional] 
**Total** | **Decimal** |  | [optional] 
**TotalContainers** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardInstanceStats = Initialize-PSOpenAPIToolsDashboardInstanceStats  -MaxCpu null `
 -MaxCores null `
 -CpuUsage null `
 -CpuUsageAverage null `
 -CpuUsagePeak null `
 -UsedMemory null `
 -MaxMemory null `
 -UsedStorage null `
 -MaxStorage null `
 -Running null `
 -Total null `
 -TotalContainers null
```

- Convert the resource to JSON
```powershell
$DashboardInstanceStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

