# ApiInstancesIdDeploysAppDeploy
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | **Int64** | Deployment ID. | [optional] 
**Version** | **String** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version | [optional] 
**VersionId** | **Int64** | Deployment Version ID. This can be passed instead of deploymentId and version. | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Map of configuration properties that vary by instance type. | [optional] 
**StageOnly** | **Boolean** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$ApiInstancesIdDeploysAppDeploy = Initialize-PSOpenAPIToolsApiInstancesIdDeploysAppDeploy  -DeploymentId null `
 -Version null `
 -VersionId null `
 -Config null `
 -StageOnly null
```

- Convert the resource to JSON
```powershell
$ApiInstancesIdDeploysAppDeploy | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

