# ListBillingAccount200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BillingInfo** | [**Billing**](Billing.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListBillingAccount200Response = Initialize-PSOpenAPIToolsListBillingAccount200Response  -Success null `
 -BillingInfo null
```

- Convert the resource to JSON
```powershell
$ListBillingAccount200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

