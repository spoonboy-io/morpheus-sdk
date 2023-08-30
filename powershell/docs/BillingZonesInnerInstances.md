# BillingZonesInnerInstances
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Instances** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Count** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingZonesInnerInstances = Initialize-PSOpenAPIToolsBillingZonesInnerInstances  -Price null `
 -Cost null `
 -Instances null `
 -Count null
```

- Convert the resource to JSON
```powershell
$BillingZonesInnerInstances | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

