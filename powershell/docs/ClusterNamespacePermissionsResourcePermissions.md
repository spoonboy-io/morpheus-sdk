# ClusterNamespacePermissionsResourcePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllGroups** | **Boolean** |  | [optional] 
**DefaultStore** | **Boolean** |  | [optional] 
**AllPlans** | **Boolean** |  | [optional] 
**DefaultTarget** | **Boolean** |  | [optional] 
**MorpheusResourceType** | **String** |  | [optional] 
**MorpheusResourceId** | **Int64** |  | [optional] 
**CanManage** | **Boolean** |  | [optional] 
**All** | **Boolean** |  | [optional] 
**Account** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Sites** | [**ResourcePermissionsSitesInner[]**](ResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | [**ResourcePermissionsSitesInner[]**](ResourcePermissionsSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespacePermissionsResourcePermissions = Initialize-PSOpenAPIToolsClusterNamespacePermissionsResourcePermissions  -AllGroups null `
 -DefaultStore null `
 -AllPlans null `
 -DefaultTarget null `
 -MorpheusResourceType null `
 -MorpheusResourceId null `
 -CanManage null `
 -All null `
 -Account null `
 -Sites null `
 -Plans null
```

- Convert the resource to JSON
```powershell
$ClusterNamespacePermissionsResourcePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

