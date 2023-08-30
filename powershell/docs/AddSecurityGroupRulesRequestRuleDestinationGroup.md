# AddSecurityGroupRulesRequestRuleDestinationGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | The destination Security Group ID. Required for &#x60;destinationType&#x60;&#x3D;group. | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupRulesRequestRuleDestinationGroup = Initialize-PSOpenAPIToolsAddSecurityGroupRulesRequestRuleDestinationGroup  -Id 56496
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupRulesRequestRuleDestinationGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

