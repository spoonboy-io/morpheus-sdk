# BudgetPolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPrice** | **Decimal** |  | [optional] 
**MaxPriceCurrency** | **String** |  | [optional] 
**MaxPriceUnit** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BudgetPolicyTypeConfiguration = Initialize-PSOpenAPIToolsBudgetPolicyTypeConfiguration  -MaxPrice null `
 -MaxPriceCurrency null `
 -MaxPriceUnit null
```

- Convert the resource to JSON
```powershell
$BudgetPolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

