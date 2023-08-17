# SetAppSecurityGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**SecurityGroups** | [**AppSecurityGroups[]**](AppSecurityGroups.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$SetAppSecurityGroups200Response = Initialize-PSOpenAPIToolsSetAppSecurityGroups200Response  -Success null `
 -SecurityGroups null
```

- Convert the resource to JSON
```powershell
$SetAppSecurityGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

