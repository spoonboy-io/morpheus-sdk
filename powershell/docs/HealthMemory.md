# HealthMemory
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**TotalMemory** | **Int64** |  | [optional] 
**FreeMemory** | **Int64** |  | [optional] 
**UsedMemory** | **Int64** |  | [optional] 
**SystemMemory** | **Int64** |  | [optional] 
**CommittedMemory** | **Int64** |  | [optional] 
**SystemFreeMemory** | **Int64** |  | [optional] 
**SystemSwap** | **Int64** |  | [optional] 
**SystemFreeSwap** | **Int64** |  | [optional] 
**SwapPercent** | **Int64** |  | [optional] 
**MemoryPercent** | **Decimal** |  | [optional] 
**SystemMemoryPercent** | **Decimal** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthMemory = Initialize-PSOpenAPIToolsHealthMemory  -Success null `
 -MaxMemory null `
 -TotalMemory null `
 -FreeMemory null `
 -UsedMemory null `
 -SystemMemory null `
 -CommittedMemory null `
 -SystemFreeMemory null `
 -SystemSwap null `
 -SystemFreeSwap null `
 -SwapPercent null `
 -MemoryPercent null `
 -SystemMemoryPercent null `
 -Status null
```

- Convert the resource to JSON
```powershell
$HealthMemory | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

