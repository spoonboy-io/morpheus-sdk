# AddSecurityGroupRulesRequestRule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the rule | [optional] 
**Direction** | **String** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to "ingress"]
**SourceType** | **String** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60; | [optional] [default to "cidr"]
**Source** | **String** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional] 
**SourceGroup** | [**AddSecurityGroupRulesRequestRuleSourceGroup**](AddSecurityGroupRulesRequestRuleSourceGroup.md) |  | [optional] 
**SourceTier** | [**AddSecurityGroupRulesRequestRuleSourceTier**](AddSecurityGroupRulesRequestRuleSourceTier.md) |  | [optional] 
**PortRange** | **String** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  | [optional] 
**Protocol** | **String** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**DestinationType** | **String** | Either cidr, group, tier, instance. | [optional] [default to "cidr"]
**Destination** | **String** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional] 
**DestinationGroup** | [**AddSecurityGroupRulesRequestRuleDestinationGroup**](AddSecurityGroupRulesRequestRuleDestinationGroup.md) |  | [optional] 
**DestinationTier** | [**AddSecurityGroupRulesRequestRuleDestinationTier**](AddSecurityGroupRulesRequestRuleDestinationTier.md) |  | [optional] 
**RuleType** | **String** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | [default to "customRule"]
**Policy** | **String** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional] 
**InstanceTypeId** | **Int64** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupRulesRequestRule = Initialize-PSOpenAPIToolsAddSecurityGroupRulesRequestRule  -Name null `
 -Direction null `
 -SourceType null `
 -Source 50.22.10.10/32 `
 -SourceGroup null `
 -SourceTier null `
 -PortRange 55-72 `
 -Protocol null `
 -DestinationType null `
 -Destination 50.22.10.10/32 `
 -DestinationGroup null `
 -DestinationTier null `
 -RuleType null `
 -Policy null `
 -InstanceTypeId 54568
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupRulesRequestRule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

