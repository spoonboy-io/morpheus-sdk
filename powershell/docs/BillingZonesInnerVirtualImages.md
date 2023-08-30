# BillingZonesInnerVirtualImages
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**VirtualImages** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Count** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingZonesInnerVirtualImages = Initialize-PSOpenAPIToolsBillingZonesInnerVirtualImages  -Price null `
 -Cost null `
 -VirtualImages null `
 -Count null
```

- Convert the resource to JSON
```powershell
$BillingZonesInnerVirtualImages | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
