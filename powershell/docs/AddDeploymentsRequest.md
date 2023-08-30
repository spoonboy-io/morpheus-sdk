# AddDeploymentsRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | [**DeploymentCreate**](DeploymentCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddDeploymentsRequest = Initialize-PSOpenAPIToolsAddDeploymentsRequest  -Deployment null
```

- Convert the resource to JSON
```powershell
$AddDeploymentsRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

