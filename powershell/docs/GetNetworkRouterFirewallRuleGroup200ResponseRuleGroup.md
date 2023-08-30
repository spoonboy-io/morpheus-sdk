# GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**IacId** | **String** |  | [optional] 
**Zone** | **String** |  | [optional] 
**ZonePool** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**Priority** | **Int64** |  | [optional] 
**GroupLayer** | **String** |  | [optional] 
**Rules** | [**GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner[]**](GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup = Initialize-PSOpenAPIToolsGetNetworkRouterFirewallRuleGroup200ResponseRuleGroup  -Id null `
 -Name null `
 -Description null `
 -ExternalId null `
 -IacId null `
 -Zone null `
 -ZonePool null `
 -Status null `
 -Priority null `
 -GroupLayer null `
 -Rules null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouterFirewallRuleGroup200ResponseRuleGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

