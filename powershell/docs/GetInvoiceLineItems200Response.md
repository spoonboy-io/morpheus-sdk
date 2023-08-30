# GetInvoiceLineItems200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItem** | [**LineItem**](LineItem.md) |  | [optional] 
**MasterAccount** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetInvoiceLineItems200Response = Initialize-PSOpenAPIToolsGetInvoiceLineItems200Response  -LineItem null `
 -MasterAccount null
```

- Convert the resource to JSON
```powershell
$GetInvoiceLineItems200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

