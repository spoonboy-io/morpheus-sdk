# BillingServerPrices
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **String** |  | [optional] 
**PricePerUnit** | **Decimal** |  | [optional] 
**CostPerUnit** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Quantity** | **Int64** |  | [optional] 
**DatastoreId** | **String** |  | [optional] 
**VolumeType** | **String** |  | [optional] 
**Datastore** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServerPrices = Initialize-PSOpenAPIToolsBillingServerPrices  -Type null `
 -PricePerUnit null `
 -CostPerUnit null `
 -Cost null `
 -Price null `
 -Quantity null `
 -DatastoreId null `
 -VolumeType null `
 -Datastore null
```

- Convert the resource to JSON
```powershell
$BillingServerPrices | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

