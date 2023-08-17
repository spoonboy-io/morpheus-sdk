# AddCatalogItemType200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**CatalogItemType** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCatalogItemType200Response = Initialize-PSOpenAPIToolsAddCatalogItemType200Response  -Success null `
 -CatalogItemType null
```

- Convert the resource to JSON
```powershell
$AddCatalogItemType200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

