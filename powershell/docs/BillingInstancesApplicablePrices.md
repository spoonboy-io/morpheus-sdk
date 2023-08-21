# BillingInstancesApplicablePrices
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**NumUnits** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Prices** | [**BillingInstancesPrices[]**](BillingInstancesPrices.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesApplicablePrices = Initialize-PSOpenAPIToolsBillingInstancesApplicablePrices  -StartDate null `
 -EndDate null `
 -NumUnits null `
 -Cost null `
 -Price null `
 -Currency null `
 -Prices null
```

- Convert the resource to JSON
```powershell
$BillingInstancesApplicablePrices | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

