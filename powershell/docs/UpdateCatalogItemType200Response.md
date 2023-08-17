# UpdateCatalogItemType200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**CatalogItemType** | [**CatalogItemType**](CatalogItemType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$UpdateCatalogItemType200Response = Initialize-PSOpenAPIToolsUpdateCatalogItemType200Response  -Success null `
 -CatalogItemType null
```

- Convert the resource to JSON
```powershell
$UpdateCatalogItemType200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

