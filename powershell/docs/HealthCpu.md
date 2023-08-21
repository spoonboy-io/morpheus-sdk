# HealthCpu
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**CpuLoad** | **Int64** |  | [optional] 
**CpuTotalLoad** | **Int64** |  | [optional] 
**ProcessorCount** | **Int64** |  | [optional] 
**ProcessTime** | **Int64** |  | [optional] 
**SystemLoad** | **Decimal** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthCpu = Initialize-PSOpenAPIToolsHealthCpu  -Success null `
 -CpuLoad null `
 -CpuTotalLoad null `
 -ProcessorCount null `
 -ProcessTime null `
 -SystemLoad null `
 -Status null
```

- Convert the resource to JSON
```powershell
$HealthCpu | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

