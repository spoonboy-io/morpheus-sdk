# GetBillingServersIdentifier200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**BillingServer**](BillingServer.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetBillingServersIdentifier200Response = Initialize-PSOpenAPIToolsGetBillingServersIdentifier200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$GetBillingServersIdentifier200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

