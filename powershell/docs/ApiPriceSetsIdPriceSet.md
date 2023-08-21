# ApiPriceSetsIdPriceSet
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Price set name | [optional] 
**Code** | **String** | Price set code. Must be unique. | [optional] 
**RegionCode** | **String** | Price set region code | [optional] 
**Zone** | [**ApiPriceSetsPriceSetZone**](ApiPriceSetsPriceSetZone.md) |  | [optional] 
**ZonePool** | [**ApiPriceSetsPriceSetZonePool**](ApiPriceSetsPriceSetZonePool.md) |  | [optional] 
**PriceUnit** | **String** | Price Unit | [optional] 
**Type** | **String** | Price set type | [optional] 
**Prices** | **Int64[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiPriceSetsIdPriceSet = Initialize-PSOpenAPIToolsApiPriceSetsIdPriceSet  -Name testName `
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
$ApiPriceSetsIdPriceSet | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

