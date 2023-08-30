# AddPriceSetsRequestPriceSet
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Price set name | 
**Code** | **String** | Price set code. Must be unique. | 
**RegionCode** | **String** | Price set region code | [optional] 
**Zone** | [**AddPriceSetsRequestPriceSetZone**](AddPriceSetsRequestPriceSetZone.md) |  | [optional] 
**ZonePool** | [**AddPriceSetsRequestPriceSetZonePool**](AddPriceSetsRequestPriceSetZonePool.md) |  | [optional] 
**PriceUnit** | **String** | Price Unit | 
**Type** | **String** | Price set type | 
**Prices** | **Int64[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPriceSetsRequestPriceSet = Initialize-PSOpenAPIToolsAddPriceSetsRequestPriceSet  -Name testName `
 -Code priceSet1 `
 -RegionCode region.code.1 `
 -Zone null `
 -ZonePool null `
 -PriceUnit null `
 -Type null `
 -Prices null
```

- Convert the resource to JSON
```powershell
$AddPriceSetsRequestPriceSet | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

