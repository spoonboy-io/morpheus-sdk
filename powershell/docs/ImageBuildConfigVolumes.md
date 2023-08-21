# ImageBuildConfigVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Size** | **Int64** |  | [optional] 
**MaxIOPS** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**RootVolume** | **Boolean** |  | [optional] 
**StorageType** | **Int64** |  | [optional] 
**DatastoreId** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuildConfigVolumes = Initialize-PSOpenAPIToolsImageBuildConfigVolumes  -Id null `
 -Size null `
 -MaxIOPS null `
 -Name null `
 -RootVolume null `
 -StorageType null `
 -DatastoreId null
```

- Convert the resource to JSON
```powershell
$ImageBuildConfigVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

