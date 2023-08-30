# AddDeployments200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | [**DeploymentCreateSuccess**](DeploymentCreateSuccess.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddDeployments200Response = Initialize-PSOpenAPIToolsAddDeployments200Response  -Deployment null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddDeployments200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

