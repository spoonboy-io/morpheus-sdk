# BillingInstancesInstancesInnerContainersInnerUsagesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | [optional] 
**InstanceName** | **String** |  | [optional] 
**ZoneName** | **String** |  | [optional] 
**AccountName** | **String** |  | [optional] 
**Volumes** | [**BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner[]**](BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner.md) |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**MaxCpu** | **String** |  | [optional] 
**MaxCores** | **Int64** |  | [optional] 
**ServerExternalId** | **String** |  | [optional] 
**ServerInternalId** | **String** |  | [optional] 
**PlanName** | **String** |  | [optional] 
**HourlyPrice** | **Decimal** |  | [optional] 
**HourlyCost** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**PricesUsed** | [**BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner[]**](BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner.md) |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**CreatedByUser** | **String** |  | [optional] 
**CreatedByUserId** | **Int64** |  | [optional] 
**SiteId** | **Int64** |  | [optional] 
**SiteName** | **String** |  | [optional] 
**SiteUUID** | **String** |  | [optional] 
**SiteCode** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Status** | **String** |  | [optional] 
**Tags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ApplicablePrices** | [**BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner[]**](BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner.md) |  | [optional] 
**ServicePlanId** | **Int64** |  | [optional] 
**ServicePlanName** | **String** |  | [optional] 
**ResourcePoolId** | **Int64** |  | [optional] 
**ResourcePoolName** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesInstancesInnerContainersInnerUsagesInner = Initialize-PSOpenAPIToolsBillingInstancesInstancesInnerContainersInnerUsagesInner  -Name null `
 -InstanceName null `
 -ZoneName null `
 -AccountName null `
 -Volumes null `
 -MaxMemory null `
 -MaxCpu null `
 -MaxCores null `
 -ServerExternalId null `
 -ServerInternalId null `
 -PlanName null `
 -HourlyPrice null `
 -HourlyCost null `
 -Currency null `
 -PricesUsed null `
 -Cost null `
 -Price null `
 -CreatedByUser null `
 -CreatedByUserId null `
 -SiteId null `
 -SiteName null `
 -SiteUUID null `
 -SiteCode null `
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
$BillingInstancesInstancesInnerContainersInnerUsagesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

