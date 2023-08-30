# BillingZonesInnerSnapshots
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Snapshots** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Count** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingZonesInnerSnapshots = Initialize-PSOpenAPIToolsBillingZonesInnerSnapshots  -Price null `
 -Cost null `
 -Snapshots null `
 -Count null
```

- Convert the resource to JSON
```powershell
$BillingZonesInnerSnapshots | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

