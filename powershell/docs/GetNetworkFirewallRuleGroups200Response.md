# GetNetworkFirewallRuleGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroups** | [**GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner[]**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkFirewallRuleGroups200Response = Initialize-PSOpenAPIToolsGetNetworkFirewallRuleGroups200Response  -RuleGroups null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkFirewallRuleGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

