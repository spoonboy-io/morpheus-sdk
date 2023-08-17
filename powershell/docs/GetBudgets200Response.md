# GetBudgets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Budget** | [**Budget**](Budget.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBudgets200Response = Initialize-PSOpenAPIToolsGetBudgets200Response  -Success null `
 -Budget null
```

- Convert the resource to JSON
```powershell
$GetBudgets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

