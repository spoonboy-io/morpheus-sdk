# CreateSubnetRequestResourcePermission
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Pass true to allow access all groups | [optional] 
**Sites** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner[]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of groups ID objects that are allowed access | [optional] 
**AllPlans** | **Boolean** |  | [optional] 
**Plans** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateSubnetRequestResourcePermission = Initialize-PSOpenAPIToolsCreateSubnetRequestResourcePermission  -All null `
 -Sites null `
 -AllPlans null `
 -Plans null
```

- Convert the resource to JSON
```powershell
$CreateSubnetRequestResourcePermission | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

