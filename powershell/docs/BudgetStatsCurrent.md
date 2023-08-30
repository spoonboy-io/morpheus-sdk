# BudgetStatsCurrent
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedCost** | **String** |  | [optional] 
**LastCost** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BudgetStatsCurrent = Initialize-PSOpenAPIToolsBudgetStatsCurrent  -EstimatedCost null `
 -LastCost null
```

- Convert the resource to JSON
```powershell
$BudgetStatsCurrent | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

