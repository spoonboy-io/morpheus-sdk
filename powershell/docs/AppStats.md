# AppStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedMemory** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**UsedStorage** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**Running** | **Int64** |  | [optional] 
**Total** | **Int64** |  | [optional] 
**CpuUsage** | **Decimal** |  | [optional] 
**InstanceCount** | **Int64** |  | [optional] 
**InstanceDayCount** | **Int64[]** |  | [optional] 
**InstanceDayCountTotal** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppStats = Initialize-PSOpenAPIToolsAppStats  -UsedMemory null `
 -MaxMemory null `
 -UsedStorage null `
 -MaxStorage null `
 -Running null `
 -Total null `
 -CpuUsage null `
 -InstanceCount null `
 -InstanceDayCount null `
 -InstanceDayCountTotal null
```

- Convert the resource to JSON
```powershell
$AppStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

