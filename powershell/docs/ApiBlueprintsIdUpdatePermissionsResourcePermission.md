# ApiBlueprintsIdUpdatePermissionsResourcePermission
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Set to true to grant access to all groups | [optional] 
**Sites** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of objects identifying groups with access | [optional] 
**OwnerId** | **Int64** | User ID, can be used to change blueprint owner. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiBlueprintsIdUpdatePermissionsResourcePermission = Initialize-PSOpenAPIToolsApiBlueprintsIdUpdatePermissionsResourcePermission  -All null `
 -Sites null `
 -OwnerId null
```

- Convert the resource to JSON
```powershell
$ApiBlueprintsIdUpdatePermissionsResourcePermission | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

