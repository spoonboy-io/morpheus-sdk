# ListBudgets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Budgets** | [**Budgets[]**](Budgets.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBudgets200Response = Initialize-PSOpenAPIToolsListBudgets200Response  -Meta null `
 -Budgets null
```

- Convert the resource to JSON
```powershell
$ListBudgets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

