# HealthElastic
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**Master** | [**HealthElasticMaster**](HealthElasticMaster.md) |  | [optional] 
**Nodes** | [**HealthElasticNodes[]**](HealthElasticNodes.md) |  | [optional] 
**Stats** | [**HealthElasticStats**](HealthElasticStats.md) |  | [optional] 
**Indices** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**BadIndices** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthElastic = Initialize-PSOpenAPIToolsHealthElastic  -Success null `
 -Status null `
 -Master null `
 -Nodes null `
 -Stats null `
 -Indices null `
 -BadIndices null
```

- Convert the resource to JSON
```powershell
$HealthElastic | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

