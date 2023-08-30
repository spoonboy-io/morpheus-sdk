# GetInvoiceLineItems200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItem** | [**LineItem**](LineItem.md) |  | [optional] 
**MasterAccount** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetInvoiceLineItems200ResponseAllOf = Initialize-PSOpenAPIToolsGetInvoiceLineItems200ResponseAllOf  -LineItem null `
 -MasterAccount null
```

- Convert the resource to JSON
```powershell
$GetInvoiceLineItems200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

