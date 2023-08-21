# BillingServersPricesUsed
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** |  | [optional] 
**PricePerUnit** | **Decimal** |  | [optional] 
**CostPerUnit** | **Decimal** |  | [optional] 
**Quantity** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServersPricesUsed = Initialize-PSOpenAPIToolsBillingServersPricesUsed  -Type null `
 -PricePerUnit null `
 -CostPerUnit null `
 -Quantity null
```

- Convert the resource to JSON
```powershell
$BillingServersPricesUsed | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
