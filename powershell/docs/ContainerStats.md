# ContainerStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **String** |  | [optional] 
**Running** | **Boolean** |  | [optional] 
**UserCpuUsage** | **Decimal** |  | [optional] 
**SystemCpuUsage** | **Decimal** |  | [optional] 
**UsedMemory** | **Int32** |  | [optional] 
**MaxMemory** | **Int32** |  | [optional] 
**CacheMemory** | **Int32** |  | [optional] 
**MaxStorage** | **Int32** |  | [optional] 
**UsedStorage** | **Int32** |  | [optional] 
**ReadIOPS** | **Int32** |  | [optional] 
**WriteIOPS** | **Decimal** |  | [optional] 
**TotalIOPS** | **Decimal** |  | [optional] 
**NetTxUsage** | **Int32** |  | [optional] 
**NetRxUsage** | **Int32** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerStats = Initialize-PSOpenAPIToolsContainerStats  -Ts null `
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
$ContainerStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

