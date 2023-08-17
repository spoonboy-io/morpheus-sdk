# AppStateInput
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variables** | [**AppStateInputVariablesInner[]**](AppStateInputVariablesInner.md) |  | [optional] 
**Providers** | [**AppStateInputProvidersInner[]**](AppStateInputProvidersInner.md) |  | [optional] 
**VarData** | [**AppStateInputDataInner[]**](AppStateInputDataInner.md) |  | [optional] 

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

