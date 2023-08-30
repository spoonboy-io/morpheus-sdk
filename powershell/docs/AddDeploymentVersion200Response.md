# AddDeploymentVersion200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | [**DeploymentVersion**](DeploymentVersion.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddDeploymentVersion200Response = Initialize-PSOpenAPIToolsAddDeploymentVersion200Response  -Version null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddDeploymentVersion200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

