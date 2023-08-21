# AppStateInput
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | [**AppStateInputVariables[]**](AppStateInputVariables.md) |  | [optional] 
**Providers** | [**AppStateInputProviders[]**](AppStateInputProviders.md) |  | [optional] 
**VarData** | [**AppStateInputData[]**](AppStateInputData.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppStateInput = Initialize-PSOpenAPIToolsAppStateInput  -Variables null `
 -Providers null `
 -VarData null
```

- Convert the resource to JSON
```powershell
$AppStateInput | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

