# HealthElasticNodesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **String** |  | [optional] 
**HeapPercent** | **String** |  | [optional] 
**RamPercent** | **String** |  | [optional] 
**CpuCount** | **String** |  | [optional] 
**LoadOne** | **String** |  | [optional] 
**LoadFive** | **String** |  | [optional] 
**LoadFifteen** | **String** |  | [optional] 
**Role** | **String** |  | [optional] 
**Master** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthElasticNodesInner = Initialize-PSOpenAPIToolsHealthElasticNodesInner  -Ip null `
 -HeapPercent null `
 -RamPercent null `
 -CpuCount null `
 -LoadOne null `
 -LoadFive null `
 -LoadFifteen null `
 -Role null `
 -Master null `
 -Name null
```

- Convert the resource to JSON
```powershell
$HealthElasticNodesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

