# BillingServersPrices
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** |  | [optional] 
**PricePerUnit** | **Decimal** |  | [optional] 
**CostPerUnit** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Quantity** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServersPrices = Initialize-PSOpenAPIToolsBillingServersPrices  -Type null `
 -PricePerUnit null `
 -CostPerUnit null `
 -Cost null `
 -Price null `
 -Quantity null
```

- Convert the resource to JSON
```powershell
$BillingServersPrices | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

