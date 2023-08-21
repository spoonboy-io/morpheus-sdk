# MorpheusApi.ApiSecurityGroupsIdRulesSgIdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the rule | [optional] 
**direction** | **String** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to &#39;ingress&#39;]
**sourceType** | **String** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60;. | [optional] [default to &#39;cidr&#39;]
**source** | **String** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional] 
**sourceGroup** | [**ApiSecurityGroupsIdRulesRuleSourceGroup**](ApiSecurityGroupsIdRulesRuleSourceGroup.md) |  | [optional] 
**sourceTier** | [**ApiSecurityGroupsIdRulesRuleSourceTier**](ApiSecurityGroupsIdRulesRuleSourceTier.md) |  | [optional] 
**portRange** | **String** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  | [optional] 
**protocol** | **String** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**destinationType** | **String** | Either cidr, group, tier, instance. | [optional] [default to &#39;cidr&#39;]
**destination** | **String** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional] 
**destinationGroup** | [**ApiSecurityGroupsIdRulesRuleDestinationGroup**](ApiSecurityGroupsIdRulesRuleDestinationGroup.md) |  | [optional] 
**destinationTier** | [**ApiSecurityGroupsIdRulesRuleDestinationTier**](ApiSecurityGroupsIdRulesRuleDestinationTier.md) |  | [optional] 
**ruleType** | **String** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | [default to &#39;customRule&#39;]
**policy** | **String** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional] 
**instanceTypeId** | **Number** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  | [optional] 



## Enum: DirectionEnum


* `ingress` (value: `"ingress"`)

* `egress` (value: `"egress"`)





## Enum: SourceTypeEnum


* `cidr` (value: `"cidr"`)

* `group` (value: `"group"`)

* `tier` (value: `"tier"`)

* `all` (value: `"all"`)





## Enum: ProtocolEnum


* `tcp` (value: `"tcp"`)

* `udp` (value: `"udp"`)

* `icmp` (value: `"icmp"`)





## Enum: DestinationTypeEnum


* `cidr` (value: `"cidr"`)

* `group` (value: `"group"`)

* `tier` (value: `"tier"`)

* `instance` (value: `"instance"`)





## Enum: PolicyEnum


* `accept` (value: `"accept"`)

* `deny` (value: `"deny"`)




