# AddCredentials200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | [**Credential**](Credential.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCredentials200Response = Initialize-PSOpenAPIToolsAddCredentials200Response  -Credential null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCredentials200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

