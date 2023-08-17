# BudgetStatsIntervalsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **Int64** |  | [optional] 
**Month** | **String** |  | [optional] 
**ShortName** | **String** |  | [optional] 
**ChartName** | **String** |  | [optional] 
**Budget** | **Int64** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BudgetStatsIntervalsInner = Initialize-PSOpenAPIToolsBudgetStatsIntervalsInner  -Index null `
 -Month null `
 -ShortName null `
 -ChartName null `
 -Budget null `
 -Cost null `
 -StartDate null `
 -EndDate null
```

- Convert the resource to JSON
```powershell
$BudgetStatsIntervalsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

