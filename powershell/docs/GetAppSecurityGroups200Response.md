# GetAppSecurityGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | [**AppSecurityGroups[]**](AppSecurityGroups.md) |  | [optional] 
**FirewallEnabled** | **Boolean** |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAppSecurityGroups200Response = Initialize-PSOpenAPIToolsGetAppSecurityGroups200Response  -SecurityGroups null `
 -FirewallEnabled null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetAppSecurityGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

