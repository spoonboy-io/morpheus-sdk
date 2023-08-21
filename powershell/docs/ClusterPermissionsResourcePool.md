# ClusterPermissionsResourcePool
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterPermissionsResourcePool = Initialize-PSOpenAPIToolsClusterPermissionsResourcePool  -Id null `
 -Name null `
 -Visibility null
```

- Convert the resource to JSON
```powershell
$ClusterPermissionsResourcePool | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

