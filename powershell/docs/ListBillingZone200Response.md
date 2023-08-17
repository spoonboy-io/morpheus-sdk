# ListBillingZone200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**BillingZones**](BillingZones.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBillingZone200Response = Initialize-PSOpenAPIToolsListBillingZone200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$ListBillingZone200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

