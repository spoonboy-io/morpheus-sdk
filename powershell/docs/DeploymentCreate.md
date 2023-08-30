# DeploymentCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name, a unique identifier for the deployment | [optional] 
**Description** | **String** | Description | [optional] 

## Examples

- Prepare the resource
```powershell
$DeploymentCreate = Initialize-PSOpenAPIToolsDeploymentCreate  -Name null `
 -Description null
```

- Convert the resource to JSON
```powershell
$DeploymentCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

