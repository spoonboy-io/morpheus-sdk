# GetBillingZoneIdentifier200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**BillingZone**](BillingZone.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBillingZoneIdentifier200Response = Initialize-PSOpenAPIToolsGetBillingZoneIdentifier200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$GetBillingZoneIdentifier200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

