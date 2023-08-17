# AddBudgets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Budget** | [**Budgets**](Budgets.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBudgets200Response = Initialize-PSOpenAPIToolsAddBudgets200Response  -Success null `
 -Budget null
```

- Convert the resource to JSON
```powershell
$AddBudgets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

