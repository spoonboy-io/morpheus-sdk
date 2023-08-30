# AddEnvironments200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**Environment**](Environment.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddEnvironments200Response = Initialize-PSOpenAPIToolsAddEnvironments200Response  -Environment null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddEnvironments200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

