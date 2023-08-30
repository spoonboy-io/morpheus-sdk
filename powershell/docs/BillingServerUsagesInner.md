# BillingServerUsagesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**CreatedByUser** | **String** |  | [optional] 
**CreatedByUserId** | **Int64** |  | [optional] 
**SiteId** | **String** |  | [optional] 
**SiteName** | **String** |  | [optional] 
**SiteUUID** | **String** |  | [optional] 
**SiteCode** | **String** |  | [optional] 
**Currency** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Status** | **String** |  | [optional] 
**Tags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ApplicablePrices** | [**BillingServerUsagesInnerApplicablePricesInner[]**](BillingServerUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | **Int64** |  | [optional] 
**ServicePlanName** | **String** |  | [optional] 
**ResourcePoolId** | **Int64** |  | [optional] 
**ResourcePoolName** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServerUsagesInner = Initialize-PSOpenAPIToolsBillingServerUsagesInner  -Cost null `
 -Price null `
 -CreatedByUser null `
 -CreatedByUserId null `
 -SiteId null `
 -SiteName null `
 -SiteUUID null `
 -SiteCode null `
 -Currency null `
 -StartDate null `
 -EndDate null `
 -Status null `
 -Tags null `
 -ApplicablePrices null `
 -ServicePlanId null `
 -ServicePlanName null `
 -ResourcePoolId null `
 -ResourcePoolName null
```

- Convert the resource to JSON
```powershell
$BillingServerUsagesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

