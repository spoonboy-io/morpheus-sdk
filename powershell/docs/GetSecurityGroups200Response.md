# GetSecurityGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**SecurityGroup**](SecurityGroup.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetSecurityGroups200Response = Initialize-PSOpenAPIToolsGetSecurityGroups200Response  -SecurityGroup null
```

- Convert the resource to JSON
```powershell
$GetSecurityGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

