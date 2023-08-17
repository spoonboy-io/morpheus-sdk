# ListCatalogItemTypes200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**CatalogItemTypes** | [**CatalogItemType[]**](CatalogItemType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListCatalogItemTypes200Response = Initialize-PSOpenAPIToolsListCatalogItemTypes200Response  -Meta null `
 -CatalogItemTypes null
```

- Convert the resource to JSON
```powershell
$ListCatalogItemTypes200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

