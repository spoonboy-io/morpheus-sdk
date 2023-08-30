# AddPolicies200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**Policy**](Policy.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPolicies200Response = Initialize-PSOpenAPIToolsAddPolicies200Response  -Policy null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddPolicies200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

