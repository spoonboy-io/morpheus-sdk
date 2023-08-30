# GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Direction** | **String** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**DestinationType** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Policy** | **String** |  | [optional] 
**Priority** | **Int64** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**RuleGroup** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**GroupName** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Sources** | [**GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**Destinations** | [**GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**Applications** | [**GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**Scopes** | [**GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**Profiles** | [**GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**AppliedTargets** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner = Initialize-PSOpenAPIToolsGetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner  -Id null `
 -Direction null `
 -SourceType null `
 -DestinationType null `
 -Name null `
 -Policy null `
 -Priority null `
 -Enabled null `
 -RuleGroup null `
 -GroupName null `
 -Config null `
 -Sources null `
 -Destinations null `
 -Applications null `
 -Scopes null `
 -Profiles null `
 -AppliedTargets null
```

- Convert the resource to JSON
```powershell
$GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

