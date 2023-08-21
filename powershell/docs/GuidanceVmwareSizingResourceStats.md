# GuidanceVmwareSizingResourceStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | **String** |  | [optional] 
**FreeMemory** | **Int64** |  | [optional] 
**UsedMemory** | **Int64** |  | [optional] 
**FreeSwap** | **Int64** |  | [optional] 
**UsedSwap** | **Int64** |  | [optional] 
**CpuIdleTime** | **Int64** |  | [optional] 
**CpuSystemTime** | **Int64** |  | [optional] 
**CpuUserTime** | **Int64** |  | [optional] 
**CpuTotalTime** | **Int64** |  | [optional] 
**CpuUsage** | **Decimal** |  | [optional] 
**UsedStorage** | **Int64** |  | [optional] 
**ReservedStorage** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**NetTxUsage** | **Int64** |  | [optional] 
**NetRxUsage** | **Int64** |  | [optional] 
**NetworkBandwidth** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingResourceStats = Initialize-PSOpenAPIToolsGuidanceVmwareSizingResourceStats  -Ts null `
 -FreeMemory null `
 -UsedMemory null `
 -FreeSwap null `
 -UsedSwap null `
 -CpuIdleTime null `
 -CpuSystemTime null `
 -CpuUserTime null `
 -CpuTotalTime null `
 -CpuUsage null `
 -UsedStorage null `
 -ReservedStorage null `
 -MaxStorage null `
 -NetTxUsage null `
 -NetRxUsage null `
 -NetworkBandwidth null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingResourceStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

