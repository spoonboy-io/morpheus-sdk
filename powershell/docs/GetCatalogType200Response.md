# GetCatalogType200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItemTypes** | [**CatalogType[]**](CatalogType.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCatalogType200Response = Initialize-PSOpenAPIToolsGetCatalogType200Response  -CatalogItemTypes null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetCatalogType200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

