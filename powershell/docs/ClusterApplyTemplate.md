# ClusterApplyTemplate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**ExecutionId** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterApplyTemplate = Initialize-PSOpenAPIToolsClusterApplyTemplate  -Success null `
 -ExecutionId null
```

- Convert the resource to JSON
```powershell
$ClusterApplyTemplate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

