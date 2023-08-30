# AddSecurityGroupRulesRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**AddSecurityGroupRulesRequestRule**](AddSecurityGroupRulesRequestRule.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupRulesRequest = Initialize-PSOpenAPIToolsAddSecurityGroupRulesRequest  -Rule null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupRulesRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

