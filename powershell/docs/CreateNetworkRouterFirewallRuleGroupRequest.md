# CreateNetworkRouterFirewallRuleGroupRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroup** | [**NetworkFirewallRuleGroupCreate**](NetworkFirewallRuleGroupCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkRouterFirewallRuleGroupRequest = Initialize-PSOpenAPIToolsCreateNetworkRouterFirewallRuleGroupRequest  -RuleGroup null
```

- Convert the resource to JSON
```powershell
$CreateNetworkRouterFirewallRuleGroupRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

