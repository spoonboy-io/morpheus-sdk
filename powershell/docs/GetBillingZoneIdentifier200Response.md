# GetBillingZoneIdentifier200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | [**BillingZone**](BillingZone.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBillingZoneIdentifier200Response = Initialize-PSOpenAPIToolsGetBillingZoneIdentifier200Response  -BillingInfo null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetBillingZoneIdentifier200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

