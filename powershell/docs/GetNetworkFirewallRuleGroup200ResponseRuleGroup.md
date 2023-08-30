# GetNetworkFirewallRuleGroup200ResponseRuleGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Priority** | **Int64** |  | [optional] 
**GroupLayer** | **String** |  | [optional] 
**Rules** | [**GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner[]**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkFirewallRuleGroup200ResponseRuleGroup = Initialize-PSOpenAPIToolsGetNetworkFirewallRuleGroup200ResponseRuleGroup  -Id null `
 -Name null `
 -Description null `
 -Priority null `
 -GroupLayer null `
 -Rules null
```

- Convert the resource to JSON
```powershell
$GetNetworkFirewallRuleGroup200ResponseRuleGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

