# ClusterDatastoresPermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterDatastoresPermissionsResourcePermissions**](ClusterDatastoresPermissionsResourcePermissions.md) |  | [optional] 
**Tenants** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterDatastoresPermissions = Initialize-PSOpenAPIToolsClusterDatastoresPermissions  -ResourcePermissions null `
 -Tenants null
```

- Convert the resource to JSON
```powershell
$ClusterDatastoresPermissions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

