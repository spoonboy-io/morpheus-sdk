# ClusterNamespaceCreateSuccess
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**RegionCode** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterNamespaceCreateSuccess = Initialize-PSOpenAPIToolsClusterNamespaceCreateSuccess  -Id null `
 -Name null `
 -Description null `
 -RegionCode null `
 -ExternalId null `
 -Status null
```

- Convert the resource to JSON
```powershell
$ClusterNamespaceCreateSuccess | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

