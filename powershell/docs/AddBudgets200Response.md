# AddBudgets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | [**Budgets**](Budgets.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBudgets200Response = Initialize-PSOpenAPIToolsAddBudgets200Response  -Budget null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddBudgets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

