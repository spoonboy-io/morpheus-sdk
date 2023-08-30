# GetBudgets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | [**Budget**](Budget.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBudgets200Response = Initialize-PSOpenAPIToolsGetBudgets200Response  -Budget null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetBudgets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

