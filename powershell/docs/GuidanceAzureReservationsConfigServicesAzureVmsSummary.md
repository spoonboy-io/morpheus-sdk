# GuidanceAzureReservationsConfigServicesAzureVmsSummary
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSavings** | **Decimal** |  | [optional] 
**CurrencyCode** | **String** |  | [optional] 
**TotalSavingsPercent** | **Decimal** |  | [optional] 
**Term** | **String** |  | [optional] 
**PaymentOption** | **String** |  | [optional] 
**Service** | **String** |  | [optional] 
**OnDemandCount** | **Int64** |  | [optional] 
**OnDemandCost** | **Decimal** |  | [optional] 
**ReservedCount** | **Int64** |  | [optional] 
**ReservedCost** | **Int64** |  | [optional] 
**RecommendedCount** | **Int64** |  | [optional] 
**RecommendedCost** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceAzureReservationsConfigServicesAzureVmsSummary = Initialize-PSOpenAPIToolsGuidanceAzureReservationsConfigServicesAzureVmsSummary  -TotalSavings null `
 -CurrencyCode null `
 -TotalSavingsPercent null `
 -Term null `
 -PaymentOption null `
 -Service null `
 -OnDemandCount null `
 -OnDemandCost null `
 -ReservedCount null `
 -ReservedCost null `
 -RecommendedCount null `
 -RecommendedCost null
```

- Convert the resource to JSON
```powershell
$GuidanceAzureReservationsConfigServicesAzureVmsSummary | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

