# DeploymentVersionCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **String** | Version number (userVersion), a unique version identifier for the deployment version. | [optional] 
**UserVersion** | **String** | Alias for version | [optional] 
**DeployType** | **String** | Deploy Type, eg. file, git, fetch | [optional] 
**GitUrl** | **String** |  | [optional] 
**GitRef** | **String** |  | [optional] 
**FetchUrl** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DeploymentVersionCreate = Initialize-PSOpenAPIToolsDeploymentVersionCreate  -Version null `
 -UserVersion null `
 -DeployType null `
 -GitUrl null `
 -GitRef null `
 -FetchUrl null
```

- Convert the resource to JSON
```powershell
$DeploymentVersionCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

