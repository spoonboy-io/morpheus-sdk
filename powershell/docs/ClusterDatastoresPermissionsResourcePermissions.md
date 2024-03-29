# ClusterDatastoresPermissionsResourcePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | **Boolean** |  | [optional] 
**AllPlans** | **Boolean** |  | [optional] 
**DefaultTarget** | **Boolean** |  | [optional] 
**CanManage** | **Boolean** |  | [optional] 
**All** | **Boolean** |  | [optional] 
**Account** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Sites** | [**ResourcePermissionsSitesInner[]**](ResourcePermissionsSitesInner.md) |  | [optional] 
**Plans** | [**ResourcePermissionsSitesInner[]**](ResourcePermissionsSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterDatastoresPermissionsResourcePermissions = Initialize-PSOpenAPIToolsClusterDatastoresPermissionsResourcePermissions  -DefaultStore null `
 -AllPlans null `
 -DefaultTarget null `
 -CanManage null `
 -All null `
 -Account null `
 -Sites null `
 -Plans null
```

- Convert the resource to JSON
```powershell
$ClusterDatastoresPermissionsResourcePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

