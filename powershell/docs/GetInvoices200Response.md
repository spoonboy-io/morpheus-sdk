# GetInvoices200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [**Invoice**](Invoice.md) |  | [optional] 
**MasterAccount** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetInvoices200Response = Initialize-PSOpenAPIToolsGetInvoices200Response  -Invoice null `
 -MasterAccount null
```

- Convert the resource to JSON
```powershell
$GetInvoices200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

