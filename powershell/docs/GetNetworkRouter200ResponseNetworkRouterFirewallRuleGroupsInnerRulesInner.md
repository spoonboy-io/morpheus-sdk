# GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**GroupName** | **String** |  | [optional] 
**Direction** | **String** |  | [optional] 
**RuleType** | **String** |  | [optional] 
**Policy** | **String** |  | [optional] 
**Source** | **String[]** |  | [optional] 
**SourceType** | **String** |  | [optional] 
**Destination** | **String[]** |  | [optional] 
**DestinationType** | **String** |  | [optional] 
**Profiles** | **String[]** |  | [optional] 
**Protocol** | **String** |  | [optional] 
**Application** | **String** |  | [optional] 
**ApplicationType** | **String** |  | [optional] 
**PortRange** | **String** |  | [optional] 
**SourcePortRange** | **String** |  | [optional] 
**SourceGroup** | **String** |  | [optional] 
**SourceTier** | **String** |  | [optional] 
**Applications** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner = Initialize-PSOpenAPIToolsGetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner  -Id null `
 -Name null `
 -Code null `
 -Enabled null `
 -GroupName null `
 -Direction null `
 -RuleType null `
 -Policy null `
 -Source null `
 -SourceType null `
 -Destination null `
 -DestinationType null `
 -Profiles null `
 -Protocol null `
 -Application null `
 -ApplicationType null `
 -PortRange null `
 -SourcePortRange null `
 -SourceGroup null `
 -SourceTier null `
 -Applications null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouter200ResponseNetworkRouterFirewallRuleGroupsInnerRulesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

