# BudgetStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **String** |  | [optional] 
**ConversionRate** | **Int64** |  | [optional] 
**Intervals** | [**BudgetStatsIntervals[]**](BudgetStatsIntervals.md) |  | [optional] 
**Current** | [**BudgetStatsCurrent**](BudgetStatsCurrent.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BudgetStats = Initialize-PSOpenAPIToolsBudgetStats  -Currency null `
 -ConversionRate null `
 -Intervals null `
 -Current null
```

- Convert the resource to JSON
```powershell
$BudgetStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

