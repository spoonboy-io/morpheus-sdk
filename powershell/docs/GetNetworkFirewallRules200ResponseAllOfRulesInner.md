# GetNetworkFirewallRules200ResponseAllOfRulesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Direction** | **String** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**DestinationType** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Policy** | **String** |  | [optional] 
**Priority** | **Int32** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**RuleGroup** | [**GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner**](GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner.md) |  | [optional] 
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
$GetNetworkFirewallRules200ResponseAllOfRulesInner = Initialize-PSOpenAPIToolsGetNetworkFirewallRules200ResponseAllOfRulesInner  -Id null `
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
$GetNetworkFirewallRules200ResponseAllOfRulesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

