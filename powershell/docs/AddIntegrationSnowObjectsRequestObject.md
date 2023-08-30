# AddIntegrationSnowObjectsRequestObject
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name to display | [optional] 
**Type** | **String** | Integration Object Type Code | 
**CatalogItemType** | **Int64** | Catalog Item Type ID | 

## Examples

- Prepare the resource
```powershell
$AddIntegrationSnowObjectsRequestObject = Initialize-PSOpenAPIToolsAddIntegrationSnowObjectsRequestObject  -Name null `
 -Type catalog `
 -CatalogItemType 27
```

- Convert the resource to JSON
```powershell
$AddIntegrationSnowObjectsRequestObject | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

