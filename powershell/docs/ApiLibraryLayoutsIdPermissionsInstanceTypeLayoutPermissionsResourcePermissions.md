# ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Set to true to grant access to all groups | [optional] 
**Sites** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of objects identifying groups with access | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions = Initialize-PSOpenAPIToolsApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions  -All null `
 -Sites null
```

- Convert the resource to JSON
```powershell
$ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissionsResourcePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

