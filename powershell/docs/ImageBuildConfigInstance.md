# ImageBuildConfigInstance
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | [**ImageBuildsConfigPlan**](ImageBuildsConfigPlan.md) |  | [optional] 
**Type** | **String** |  | [optional] 
**UserGroup** | [**ZoneVcenterConfigNetworkServer**](ZoneVcenterConfigNetworkServer.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuildConfigInstance = Initialize-PSOpenAPIToolsImageBuildConfigInstance  -Layout null `
 -Type null `
 -UserGroup null
```

- Convert the resource to JSON
```powershell
$ImageBuildConfigInstance | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

