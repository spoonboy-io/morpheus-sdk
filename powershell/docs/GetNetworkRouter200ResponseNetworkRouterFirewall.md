# GetNetworkRouter200ResponseNetworkRouterFirewall
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **Boolean** |  | [optional] 
**Version** | **String** |  | [optional] 
**DefaultPolicy** | **String** |  | [optional] 
**Global** | **String** |  | [optional] 
**RuleGroups** | [**GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInner[]**](GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouter200ResponseNetworkRouterFirewall = Initialize-PSOpenAPIToolsGetNetworkRouter200ResponseNetworkRouterFirewall  -Enabled null `
 -Version null `
 -DefaultPolicy null `
 -Global null `
 -RuleGroups null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouter200ResponseNetworkRouterFirewall | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

