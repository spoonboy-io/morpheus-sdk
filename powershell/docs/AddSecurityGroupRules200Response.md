# AddSecurityGroupRules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**SecurityGroupRule**](SecurityGroupRule.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupRules200Response = Initialize-PSOpenAPIToolsAddSecurityGroupRules200Response  -Rule null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupRules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

