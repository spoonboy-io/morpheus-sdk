# ListBillingInstances200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**BillingInstances**](BillingInstances.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBillingInstances200Response = Initialize-PSOpenAPIToolsListBillingInstances200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$ListBillingInstances200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

