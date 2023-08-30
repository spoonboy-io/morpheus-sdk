# AddCatalogOrder200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**CatalogOrderCreateSuccess**](CatalogOrderCreateSuccess.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCatalogOrder200Response = Initialize-PSOpenAPIToolsAddCatalogOrder200Response  -Order null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCatalogOrder200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

