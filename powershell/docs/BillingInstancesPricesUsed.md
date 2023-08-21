# BillingInstancesPricesUsed
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** |  | [optional] 
**PricePerUnit** | **Decimal** |  | [optional] 
**CostPerUnit** | **Decimal** |  | [optional] 
**Quantity** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesPricesUsed = Initialize-PSOpenAPIToolsBillingInstancesPricesUsed  -Type null `
 -PricePerUnit null `
 -CostPerUnit null `
 -Quantity null
```

- Convert the resource to JSON
```powershell
$BillingInstancesPricesUsed | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

