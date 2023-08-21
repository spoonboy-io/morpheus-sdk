# ClusterLayoutFile
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**ContentRef** | **String** |  | [optional] 
**ContentPath** | **String** |  | [optional] 
**Repository** | **String** |  | [optional] 
**Content** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutFile = Initialize-PSOpenAPIToolsClusterLayoutFile  -Id null `
 -SourceType null `
 -ContentRef null `
 -ContentPath null `
 -Repository null `
 -Content null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutFile | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

