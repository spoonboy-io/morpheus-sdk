# AddSecurityGroupRulesRequestRuleSourceGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | The source Security Group ID. Required for &#x60;sourceType&#x60;&#x3D;group | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupRulesRequestRuleSourceGroup = Initialize-PSOpenAPIToolsAddSecurityGroupRulesRequestRuleSourceGroup  -Id 56496
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupRulesRequestRuleSourceGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

