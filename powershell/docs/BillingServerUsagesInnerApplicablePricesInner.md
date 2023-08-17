# BillingServerUsagesInnerApplicablePricesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**NumUnits** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Prices** | [**BillingServerUsagesInnerApplicablePricesInnerPricesInner[]**](BillingServerUsagesInnerApplicablePricesInnerPricesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServerUsagesInnerApplicablePricesInner = Initialize-PSOpenAPIToolsBillingServerUsagesInnerApplicablePricesInner  -StartDate null `
 -EndDate null `
 -NumUnits null `
 -Cost null `
 -Price null `
 -Currency null `
 -Prices null
```

- Convert the resource to JSON
```powershell
$BillingServerUsagesInnerApplicablePricesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

