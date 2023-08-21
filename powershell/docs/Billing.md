# Billing
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **Int64** |  | [optional] 
**AccountUUID** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**PriceUnit** | **String** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Zones** | [**BillingZones[]**](BillingZones.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Billing = Initialize-PSOpenAPIToolsBilling  -AccountId null `
 -AccountUUID null `
 -Name null `
 -StartDate null `
 -EndDate null `
 -PriceUnit null `
 -Price null `
 -Cost null `
 -Zones null
```

- Convert the resource to JSON
```powershell
$Billing | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

