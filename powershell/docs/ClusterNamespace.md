# ClusterNamespace
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Permissions** | [**ClusterNamespacePermissions**](ClusterNamespacePermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespace = Initialize-PSOpenAPIToolsClusterNamespace  -Id null `
 -Visibility null `
 -Name null `
 -Description null `
 -Status null `
 -ExternalId null `
 -Permissions null
```

- Convert the resource to JSON
```powershell
$ClusterNamespace | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

