# GetAppSecurityGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**SecurityGroups** | [**AppSecurityGroups[]**](AppSecurityGroups.md) |  | [optional] 
**FirewallEnabled** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAppSecurityGroups200Response = Initialize-PSOpenAPIToolsGetAppSecurityGroups200Response  -Success null `
 -SecurityGroups null `
 -FirewallEnabled null
```

- Convert the resource to JSON
```powershell
$GetAppSecurityGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

