# ClusterNamespaceCreateResourcePermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | **Boolean** | Pass true to allow access to all groups | [optional] 
**Sites** | [**ClusterUpdatePermissionsResourcePermissionsSites[]**](ClusterUpdatePermissionsResourcePermissionsSites.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | **Boolean** | Pass true to allow access to all plans | [optional] 
**Plans** | [**ClusterUpdatePermissionsResourcePermissionsSites[]**](ClusterUpdatePermissionsResourcePermissionsSites.md) | Array of plans that are allowed access | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespaceCreateResourcePermissions = Initialize-PSOpenAPIToolsClusterNamespaceCreateResourcePermissions  -All null `
 -Sites null `
 -AllPlans null `
 -Plans null
```

- Convert the resource to JSON
```powershell
$ClusterNamespaceCreateResourcePermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

