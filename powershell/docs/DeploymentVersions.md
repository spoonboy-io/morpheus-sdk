# DeploymentVersions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**DeployType** | **String** |  | [optional] 
**FetchUrl** | **String** |  | [optional] 
**GitUrl** | **String** |  | [optional] 
**GitRef** | **String** |  | [optional] 
**UserVersion** | **String** |  | [optional] 
**Version** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DeploymentVersions = Initialize-PSOpenAPIToolsDeploymentVersions  -Id null `
 -DeployType null `
 -FetchUrl null `
 -GitUrl null `
 -GitRef null `
 -UserVersion null `
 -Version null `
 -Status null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$DeploymentVersions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

