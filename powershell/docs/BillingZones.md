# BillingZones
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Zones** | [**BillingZonesInner[]**](BillingZonesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingZones = Initialize-PSOpenAPIToolsBillingZones  -Price null `
 -Cost null `
 -StartDate null `
 -EndDate null `
 -Zones null
```

- Convert the resource to JSON
```powershell
$BillingZones | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

