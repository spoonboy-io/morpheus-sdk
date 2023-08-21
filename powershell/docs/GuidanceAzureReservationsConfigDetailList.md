# GuidanceAzureReservationsConfigDetailList
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **String** |  | [optional] 
**ApiName** | **String** |  | [optional] 
**ApiType** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Period** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Size** | **String** |  | [optional] 
**Region** | **String** |  | [optional] 
**Term** | **String** |  | [optional] 
**MeterId** | **String** |  | [optional] 
**OnDemandCount** | **Int64** |  | [optional] 
**OnDemandCost** | **Decimal** |  | [optional] 
**ReservedCount** | **Int64** |  | [optional] 
**ReservedCost** | **Int64** |  | [optional] 
**RecommendedCount** | **Int64** |  | [optional] 
**RecommendedCost** | **Decimal** |  | [optional] 
**TotalSavings** | **Decimal** |  | [optional] 
**TotalSavingsPercent** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceAzureReservationsConfigDetailList = Initialize-PSOpenAPIToolsGuidanceAzureReservationsConfigDetailList  -Id null `
 -ApiName null `
 -ApiType null `
 -ExternalId null `
 -Period null `
 -Name null `
 -Type null `
 -Category null `
 -Size null `
 -Region null `
 -Term null `
 -MeterId null `
 -OnDemandCount null `
 -OnDemandCost null `
 -ReservedCount null `
 -ReservedCost null `
 -RecommendedCount null `
 -RecommendedCost null `
 -TotalSavings null `
 -TotalSavingsPercent null
```

- Convert the resource to JSON
```powershell
$GuidanceAzureReservationsConfigDetailList | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

