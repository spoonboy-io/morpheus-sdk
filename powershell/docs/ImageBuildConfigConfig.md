# ImageBuildConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | **Int64** |  | [optional] 
**VmwareFolderId** | **String** |  | [optional] 
**ResourcePoolId** | **Int64** |  | [optional] 
**NestedVirtualization** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuildConfigConfig = Initialize-PSOpenAPIToolsImageBuildConfigConfig  -Template null `
 -VmwareFolderId null `
 -ResourcePoolId null `
 -NestedVirtualization null
```

- Convert the resource to JSON
```powershell
$ImageBuildConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

