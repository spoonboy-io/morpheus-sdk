# GetAppState200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workloads** | [**AppStateWorkloadsInner[]**](AppStateWorkloadsInner.md) |  | [optional] 
**IacDrift** | **Boolean** |  | [optional] 
**PlanResources** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Specs** | [**AppStateSpecsInner[]**](AppStateSpecsInner.md) |  | [optional] 
**StateData** | **String** |  | [optional] 
**PlanData** | **String** |  | [optional] 
**VarInput** | [**AppStateInput**](AppStateInput.md) |  | [optional] 
**Output** | [**AppStateOutput**](AppStateOutput.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAppState200Response = Initialize-PSOpenAPIToolsGetAppState200Response  -Workloads null `
 -IacDrift null `
 -PlanResources null `
 -Specs null `
 -StateData null `
 -PlanData null `
 -VarInput null `
 -Output null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetAppState200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

