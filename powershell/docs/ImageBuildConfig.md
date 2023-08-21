# ImageBuildConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**ImageBuildConfigInstance**](ImageBuildConfigInstance.md) |  | [optional] 
**NetworkInterfaces** | [**ImageBuildConfigNetworkInterfaces[]**](ImageBuildConfigNetworkInterfaces.md) |  | [optional] 
**Volumes** | [**ImageBuildConfigVolumes[]**](ImageBuildConfigVolumes.md) |  | [optional] 
**StorageControllers** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ZoneId** | **Int64** |  | [optional] 
**Config** | [**ImageBuildConfigConfig**](ImageBuildConfigConfig.md) |  | [optional] 
**Plan** | [**ImageBuildsConfigPlan**](ImageBuildsConfigPlan.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuildConfig = Initialize-PSOpenAPIToolsImageBuildConfig  -Instance null `
 -NetworkInterfaces null `
 -Volumes null `
 -StorageControllers null `
 -ZoneId null `
 -Config null `
 -Plan null
```

- Convert the resource to JSON
```powershell
$ImageBuildConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

