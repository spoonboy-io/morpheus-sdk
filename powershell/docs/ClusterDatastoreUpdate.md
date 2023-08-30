# ClusterDatastoreUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **Boolean** | Datastore active | [optional] [default to $true]
**Permissions** | [**ClusterUpdatePermissions**](ClusterUpdatePermissions.md) |  | [optional] 
**Visibility** | **String** | Visibility for datastore | [optional] [default to "private"]

## Examples

- Prepare the resource
```powershell
$ClusterDatastoreUpdate = Initialize-PSOpenAPIToolsClusterDatastoreUpdate  -Active null `
 -Permissions null `
 -Visibility null
```

- Convert the resource to JSON
```powershell
$ClusterDatastoreUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

