# ListBillingServers200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**BillingServers**](BillingServers.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBillingServers200Response = Initialize-PSOpenAPIToolsListBillingServers200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$ListBillingServers200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

