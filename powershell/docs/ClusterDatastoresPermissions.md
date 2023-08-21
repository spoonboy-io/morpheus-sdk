# ClusterDatastoresPermissions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | [**ClusterDatastoresPermissionsResourcePermissions**](ClusterDatastoresPermissionsResourcePermissions.md) |  | [optional] 
**Tenants** | [**InlineResponse20040AppDeployInstance[]**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

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

